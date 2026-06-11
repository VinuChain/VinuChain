package gossip

import (
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/lachesis"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"

	"github.com/Fantom-foundation/go-opera/gossip/evmstore"
	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/inter/iblockproc"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/payback"
	"github.com/Fantom-foundation/go-opera/utils/signers/internaltx"
)

var (
	// Ethereum compatible metrics set (see go-ethereum/core)

	headBlockGauge     = metrics.GetOrRegisterGauge("chain/head/block", nil)
	headHeaderGauge    = metrics.GetOrRegisterGauge("chain/head/header", nil)
	headFastBlockGauge = metrics.GetOrRegisterGauge("chain/head/receipt", nil)

	accountReadTimer   = metrics.GetOrRegisterTimer("chain/account/reads", nil)
	accountHashTimer   = metrics.GetOrRegisterTimer("chain/account/hashes", nil)
	accountUpdateTimer = metrics.GetOrRegisterTimer("chain/account/updates", nil)
	accountCommitTimer = metrics.GetOrRegisterTimer("chain/account/commits", nil)

	storageReadTimer   = metrics.GetOrRegisterTimer("chain/storage/reads", nil)
	storageHashTimer   = metrics.GetOrRegisterTimer("chain/storage/hashes", nil)
	storageUpdateTimer = metrics.GetOrRegisterTimer("chain/storage/updates", nil)
	storageCommitTimer = metrics.GetOrRegisterTimer("chain/storage/commits", nil)

	snapshotAccountReadTimer = metrics.GetOrRegisterTimer("chain/snapshot/account/reads", nil)
	snapshotStorageReadTimer = metrics.GetOrRegisterTimer("chain/snapshot/storage/reads", nil)
	snapshotCommitTimer      = metrics.GetOrRegisterTimer("chain/snapshot/commits", nil)

	blockInsertTimer    = metrics.GetOrRegisterTimer("chain/inserts", nil)
	blockExecutionTimer = metrics.GetOrRegisterTimer("chain/execution", nil)
	blockWriteTimer     = metrics.GetOrRegisterTimer("chain/write", nil)
	blockAgeGauge       = metrics.GetOrRegisterGauge("chain/block/age", nil)
)

type ExtendedTxPosition struct {
	evmstore.TxPosition
	EventCreator idx.ValidatorID
}

// GetConsensusCallbacks returns single (for Service) callback instance.
func (s *Service) GetConsensusCallbacks() lachesis.ConsensusCallbacks {
	bp := newBlockProcessor(
		s.blockProcTasks,
		&s.blockProcWg,
		&s.blockBusyFlag,
		s.store,
		s.blockProcModules,
		s.config.TxIndex,
		&s.feed,
		&s.emitters,
		s.verWatcher,
		s.paybackCache,
	)
	return lachesis.ConsensusCallbacks{
		BeginBlock: bp.Begin,
	}
}

func (s *Service) ReexecuteBlocks(from, to idx.Block) {
	blockProc := s.blockProcModules
	upgradeHeights := s.store.GetUpgradeHeights()
	evmStateReader := s.GetEvmStateReader()

	// Use a fresh PaybackCache for re-execution to avoid residual quota
	// data from the live cache producing incorrect FeeRefund values.
	reexecCache, err := payback.NewPaybackCache(s.paybackCache.GetStore(), s.store.GetRules().Economy.QuotaCacheMaxAddresses)
	if err != nil {
		log.Crit("Failed to create re-execution PaybackCache", "err", err)
	}

	prev := s.store.GetBlock(from)
	if prev == nil {
		log.Crit("Re-execution start block not found", "block", from)
	}
	for b := from + 1; b <= to; b++ {
		block := s.store.GetBlock(b)
		if block == nil {
			log.Crit("Re-execution block not found", "block", b)
		}
		blockCtx := iblockproc.BlockCtx{
			Idx:     b,
			Time:    block.Time,
			Atropos: block.Atropos,
		}
		statedb, err := s.store.evm.StateDB(prev.Root)
		if err != nil {
			// log.Crit exits without flush — acceptable here because a
			// corrupt state root during re-execution means the database is
			// unrecoverable; continuing would produce wrong chain state.
			log.Crit("Failure to re-execute blocks", "err", err)
		}
		epoch := s.store.FindBlockEpoch(b)
		if epoch == 0 {
			log.Crit("Block epoch mapping not found", "block", b)
		}
		es := s.store.GetHistoryEpochState(epoch)
		if es == nil {
			log.Crit("Epoch state not found for re-execution", "block", b, "epoch", epoch)
		}

		evmProcessor := blockProc.EVMModule.Start(blockCtx, statedb, evmStateReader, func(t *types.Log) {}, es.Rules, opera.DefaultVMConfig, es.Rules.EvmChainConfig(upgradeHeights), reexecCache, es.Epoch)
		txs := s.store.GetBlockTxs(b, block)
		evmProcessor.Execute(txs)
		evmProcessor.Finalize()
		if err := s.store.evm.Commit(b, block.Root, false); err != nil {
			log.Crit("Failed to commit EVM state during re-execution", "block", b, "err", err)
		}
		s.store.evm.Cap()
		s.mayCommit(false)
		prev = block
	}
}

func (s *Service) RecoverEVM() {
	start := s.store.GetLatestBlockIndex()
	for b := start; b >= 1 && b > start-20000; b-- {
		block := s.store.GetBlock(b)
		if block == nil {
			break
		}
		if s.store.evm.HasStateDB(block.Root) {
			if b != start {
				s.Log.Warn("Reexecuting blocks after abrupt stopping", "from", b, "to", start)
				s.ReexecuteBlocks(b, start)
			}
			break
		}
	}
}

// warmUpProgressInterval bounds how often WarmUpPaybackCache emits a progress
// log during a long replay (e.g. a genesis-imported node with a large
// trailing window). Keeps the fallback path observable without spamming logs.
const warmUpProgressInterval = 10000

// WarmUpPaybackCache rebuilds s.paybackCache (PaybackUsedMap for the current
// epoch E, and StakesMap for the previous epoch E-1 and the current epoch E) to
// the exact accumulated state a never-restarted node carries at head, by
// replaying the already-sealed blocks of epochs E-1 and E through the live cache.
//
// CONSENSUS-CRITICAL (audit A1/T2): PaybackCache.PaybackUsedMap and
// PaybackCache.StakesMap are volatile, in-memory, and never persisted.
// PaybackUsedMap is reset at every epoch boundary; StakesMap retains entries for
// the current epoch E, the previous epoch E-1, and E-2 (cleanupOldEpochsLocked
// deletes only epochs < E-2). On restart s.paybackCache is constructed empty
// (service.go) and RecoverEVM only re-derives trailing EVM state into a LOCAL
// reexecCache — it never installs quotaUsed/stakes back into s.paybackCache.
//
// Two consensus-relevant maps must be reconstructed:
//   - PaybackUsedMap[E] (A1): the first NEW block sealed after a mid-epoch
//     restart reads quotaUsed=0 for addresses that used FeeRefund earlier in
//     epoch E, yielding a larger availablePayback → larger FeeRefund → larger
//     AddBalance → a different block.Root than non-restarted peers.
//   - StakesMap[E-1] and StakesMap[E]: calculateFullDurationLocked branches on
//     getSumStakeByAddressSplitLocked(addr, E, E-1) — only epochs E and E-1 feed
//     consensus (E-2 is retained by cleanup but never read by the duration math,
//     so it need not be reconstructed). An address that staked in BOTH E-1 and E
//     would take a different duration branch if StakesMap[E-1] were empty →
//     different fullDuration → different paybackSum → different FeeRefund → split.
//
// The replay starts at the FIRST BLOCK OF EPOCH E-1 (so StakesMap[E-1] is
// reconstructed) and runs to head. PrepareForBlock's per-epoch cleanup zeroes
// PaybackUsedMap at the E-1→E boundary during the replay while StakesMap[E-1]
// survives, reproducing exactly the consensus-relevant state a live node holds at
// head: PaybackUsedMap for E only, StakesMap for E-1 and E.
//
// NOTE on the start block: the original fix already started here. The audit
// rework finding P1 (claiming the start was the first block of epoch E, leaving
// StakesMap[E-1] empty) was REFUTED — GetHistoryBlockEpochState(E-1).LastBlock.Idx
// + 1 is the first block of E-1, not E (see paybackWarmUpFirstBlock and
// docs/payback-cache-restart-determinism.md). The genuine reworks below are the
// raw-receipt read (P2a/P2b) and the bounded fallback + progress logs (P3).
//
// The replay mirrors the live forward-sealing derivation exactly:
//   - the per-block epoch is GetHistoryEpochState(FindBlockEpoch(b)).Epoch, the
//     same derivation ReexecuteBlocks uses;
//   - PrepareForBlock/AddTransaction/FinishBlock are driven in the same order as
//     OperaEVMProcessor.Execute/Finalize and evmcore.StateProcessor.Process.
//
// Receipts are read from RAW storage (GetRawReceipts) and the minimal fields
// AddTransaction needs — Status, FeeRefund, and per-tx GasUsed (derived from the
// stored CumulativeGasUsed deltas, exactly as DeriveFields does) — are computed
// locally. The warm-up deliberately does NOT call evmstore.GetReceipts: that
// path runs DeriveFields with a zero BlockHash and POISONS the shared receipts
// LRU served by RPC (GetReceiptsByNumber). The raw path touches no shared cache,
// so RPC output is never corrupted, and it never derives through DeriveFields, so
// it never trips the cryptic "transaction and receipt count mismatch" Crit on a
// normally-configured (TxIndex-enabled) node.
//
// FAIL CLOSED on missing consensus material: if a tx-bearing block's receipts
// cannot be read (the dominant cause is running with TxIndex DISABLED, under which
// receipts are never persisted — block_processor.go gates SetReceipts on
// bp.txIndex), the payback cache would be reconstructed INCOMPLETELY, and sealing
// forward from it would produce a different FeeRefund / block.Root than peers (a
// consensus split). Rather than warn-and-continue into a divergent state, the
// warm-up calls s.Log.Crit with an actionable message: validators MUST run with
// TxIndex enabled. (This is strictly safer than the pre-rework behaviour, which
// crashed via the receipt/tx count mismatch with a confusing message and, on the
// way, polluted the RPC receipts LRU.)
//
// AddTransaction relies on tx.From() returning the cached sender. Store-loaded
// txs are RLP-decoded with an empty sender cache, so the sender is recovered and
// cached here for non-internal txs via types.Sender — exactly as the live path's
// tx.AsMessage does. Internal txs are left uncached so AddTransaction skips them,
// matching the live path (TxAsMessage builds internal-tx messages without caching
// tx.from, so AddTransaction's tx.From()==zero short-circuit drops them).
//
// KNOWN LIMITATION (epoch-seal rules drift): on an epoch-sealing block the live
// processor refreshes rules MID-BLOCK (sealEpochIfNeeded → SetRules), so txs in
// the post-seal portion of a block whose seal changes Economy.QuotaCacheAddress
// are classified against the NEW quota address live, while this replay processes
// the whole block under the OLD epoch's rules. If the replay window contains such
// an activation block AND that block carries stake/unstake txs to either address
// post-seal, the reconstructed StakesMap can differ from a never-restarted peer's.
// This is at most one block per QuotaCacheAddress-changing upgrade. Operational
// rule: do NOT restart validators within the two-epoch window after such an
// upgrade activates (see docs/payback-cache-restart-determinism.md).
func (s *Service) WarmUpPaybackCache() {
	if s.paybackCache == nil {
		return
	}

	currentEpoch := s.store.GetEpoch()
	if currentEpoch < 1 {
		return
	}

	head := s.store.GetLatestBlockIndex()
	if head < 1 {
		return
	}

	// Replay from the first block of epoch E-1 so StakesMap[E-1] is reconstructed.
	// See paybackWarmUpFirstBlock for the epoch-index convention and the bounded
	// fallback used on pruned / genesis-imported nodes. PrepareForBlock self-resets
	// PaybackUsedMap at each epoch boundary, so an earlier start still converges to
	// the correct accumulation — at the cost of extra per-tx ECDSA recovery, which
	// the fallback bounds and logs.
	firstBlock := s.paybackWarmUpFirstBlock(currentEpoch, head)
	if firstBlock < 1 {
		firstBlock = 1
	}
	if firstBlock > head {
		return
	}

	signer := types.LatestSignerForChainID(s.store.GetEvmChainConfig().ChainID)

	warmed := 0
	// Leading-gap tolerance: on genesis-imported or pruned nodes the replay
	// window can begin before the earliest block the store actually holds
	// (the bounded fallback may select block 1). Blocks absent BEFORE the
	// first readable block are a well-understood leading gap and are skipped
	// (with a Warn below); any hole AFTER replay material has started is
	// structural corruption and fails closed.
	seenAny := false
	leadingGap := 0
	for b := firstBlock; b <= head; b++ {
		if (b-firstBlock)%warmUpProgressInterval == 0 && b != firstBlock {
			s.Log.Info("Warming payback cache (replaying current epoch window)",
				"block", b, "fromBlock", firstBlock, "toBlock", head)
		}

		// FAIL CLOSED on structurally missing chain data once replay material
		// has started, mirroring ReexecuteBlocks (which Crits on the same
		// conditions): a hole in the middle of the window means the cache
		// would be reconstructed incompletely and forward sealing would
		// diverge — the exact failure mode the missing-receipts path below
		// refuses to start under. Blocks absent before the first readable
		// block are a leading gap (genesis import / pruning) and are
		// tolerated; see seenAny/leadingGap above.
		block := s.store.GetBlock(b)
		if block == nil {
			if !seenAny {
				leadingGap++
				continue
			}
			s.Log.Crit("WarmUpPaybackCache: block missing inside the replay window — "+
				"the payback cache cannot be reconstructed completely; refusing to start "+
				"rather than seal divergent FeeRefund/block.Root",
				"block", b, "fromBlock", firstBlock, "toBlock", head)
		}

		// Mirror ReexecuteBlocks / the live block processor: the epoch the cache
		// sees for this block is GetHistoryEpochState(FindBlockEpoch(b)).Epoch
		// with that epoch's rules.
		epoch := s.store.FindBlockEpoch(b)
		if epoch == 0 {
			if !seenAny {
				leadingGap++
				continue
			}
			s.Log.Crit("WarmUpPaybackCache: cannot resolve epoch for a block inside the replay window",
				"block", b, "fromBlock", firstBlock, "toBlock", head)
		}
		es := s.store.GetHistoryEpochState(epoch)
		if es == nil {
			if !seenAny {
				leadingGap++
				continue
			}
			s.Log.Crit("WarmUpPaybackCache: missing history epoch state for a block inside the replay window",
				"block", b, "epoch", epoch, "fromBlock", firstBlock, "toBlock", head)
		}
		if !seenAny {
			seenAny = true
			if leadingGap > 0 {
				s.Log.Warn("WarmUpPaybackCache: skipped leading blocks absent from the store "+
					"(genesis import or pruning); replay effectively starts here",
					"skipped", leadingGap, "firstReadableBlock", b, "fromBlock", firstBlock, "toBlock", head)
			}
		}

		txs := s.store.GetBlockTxs(b, block)

		// Read RAW stored receipts (no LRU pollution; never derives through
		// DeriveFields, so no cryptic count-mismatch Crit). receipts is nil when
		// this block's receipts were never persisted (e.g. TxIndex=false) or their
		// count does not match the tx count.
		receipts := s.rawWarmUpReceipts(b, len(txs))
		if receipts == nil {
			if len(txs) > 0 {
				// FAIL CLOSED: this block carried txs but its receipts cannot be
				// read, so PaybackUsedMap[E] / StakesMap would be reconstructed
				// incompletely. Sealing forward from an incomplete payback cache
				// produces a different FeeRefund (and block.Root) than peers — a
				// consensus split. Refuse to start rather than diverge silently.
				//
				// The dominant cause is running with TxIndex disabled, under which
				// receipts are not persisted (block_processor.go gates SetReceipts
				// on bp.txIndex). Validators MUST run with TxIndex enabled.
				if !s.config.TxIndex {
					s.Log.Crit("WarmUpPaybackCache: cannot warm payback cache because TxIndex is disabled — "+
						"receipts for the replay window (epochs E-1..E) were never persisted, so this node "+
						"would seal divergent FeeRefund/block.Root. Merely enabling TxIndex does NOT "+
						"backfill receipts for already-sealed blocks: re-sync this node (snapshot or "+
						"genesis re-import) with TxIndex enabled.",
						"block", b, "epoch", currentEpoch)
				}
				s.Log.Crit("WarmUpPaybackCache: receipts missing for a tx-bearing block in the replay window — "+
					"the payback cache cannot be reconstructed and forward sealing would diverge. "+
					"Causes: receipt-store corruption, pruning of recent receipts, or a previous run "+
					"with TxIndex disabled (enabling TxIndex does NOT backfill already-sealed blocks); "+
					"re-sync the node (snapshot or genesis re-import).",
					"block", b, "epoch", currentEpoch, "txs", len(txs))
			}
			continue
		}

		s.paybackCache.PrepareForBlock(es.Epoch, es.Rules, block.Time.Time())
		for i, tx := range txs {
			if i >= len(receipts) || receipts[i] == nil {
				continue
			}
			// Populate the sender cache exactly as the live tx.AsMessage path
			// does for non-internal txs. Internal txs are intentionally left
			// uncached so AddTransaction's tx.From()==zero guard drops them,
			// matching live behaviour.
			if !internaltx.IsInternal(tx) {
				if _, err := types.Sender(signer, tx); err != nil {
					continue
				}
			}
			if err := s.paybackCache.AddTransaction(tx, receipts[i]); err != nil {
				s.Log.Error("WarmUpPaybackCache: failed to replay tx into payback cache",
					"block", b, "tx", tx.Hash(), "err", err)
			}
		}
		s.paybackCache.FinishBlock()
		warmed++
	}

	if !seenAny {
		// The entire window was a leading gap (no readable block at all):
		// inherent to a genesis cut / pruning shape that predates this node's
		// data, not introduced here — but it must not pass silently, because
		// the node proceeds with an unwarmed (empty) cache.
		s.Log.Warn("WarmUpPaybackCache: no readable block in the replay window; payback cache NOT warmed",
			"skipped", leadingGap, "fromBlock", firstBlock, "toBlock", head, "epoch", currentEpoch)
		return
	}

	s.Log.Info("Warmed payback cache for current epoch",
		"epoch", currentEpoch, "fromBlock", firstBlock, "toBlock", head, "blocks", warmed)
}

// paybackWarmUpFirstBlock returns the first block of epoch E-1 — the earliest
// block the warm-up must replay so the consensus-relevant maps a never-restarted
// node holds at head (PaybackUsedMap for E; StakesMap for E-1 and E) are
// reconstructed exactly. It bounds the fallback when epoch-boundary history is
// unavailable. See WarmUpPaybackCache.
//
// Epoch-index convention (verified against block_processor.go sealing +
// store_block.go FindBlockEpoch): GetHistoryBlockEpochState(N) is written when
// epoch N is sealed, with LastBlock.Idx = the last block of the blocks whose
// FindBlockEpoch == N-1. Therefore GetHistoryBlockEpochState(N).LastBlock.Idx + 1
// == the FIRST block whose FindBlockEpoch == N. With currentEpoch == E ==
// FindBlockEpoch(head), GetHistoryBlockEpochState(E-1).LastBlock.Idx + 1 is the
// first block of epoch E-1 — NOT epoch E. (This is why the original
// `currentEpoch - 1` lookup already covered StakesMap[E-1]; see the refuted
// audit finding P1 in docs/payback-cache-restart-determinism.md.)
func (s *Service) paybackWarmUpFirstBlock(currentEpoch idx.Epoch, head idx.Block) idx.Block {
	// First block of epoch E-1 = (last block of E-2) + 1 =
	// GetHistoryBlockEpochState(E-1).LastBlock.Idx + 1.
	if currentEpoch > 1 {
		if prevBs, _ := s.store.GetHistoryBlockEpochState(currentEpoch - 1); prevBs != nil {
			return prevBs.LastBlock.Idx + 1
		}
	} else if currentEpoch == 1 {
		// Epoch 1: no previous epoch exists; the whole chain is epoch 1.
		return 1
	}

	// Fallback (pruned / genesis-imported node, or epoch-boundary state missing):
	// walk FindBlockEpoch backwards from head to find the first block of E-1. This
	// scan only reads per-block epoch mappings (no ECDSA recovery), so it is cheap;
	// it is naturally bounded by the length of epochs E-1 and E (epochs are time-
	// and gas-bounded). We MUST find the true start of E-1 or fall back to block 1:
	// returning a mid-epoch block would replay only a suffix of epoch E, leaving
	// PaybackUsedMap[E] incomplete → excess FeeRefund → consensus split. A progress
	// log every warmUpProgressInterval blocks keeps a long scan observable.
	targetEpoch := currentEpoch - 1
	firstOfPrev := idx.Block(0)
	scanned := 0
	for b := head; b >= 1; b-- {
		e := s.store.FindBlockEpoch(b)
		if e == targetEpoch {
			firstOfPrev = b
		} else if e != 0 && e < targetEpoch && firstOfPrev != 0 {
			// Walked past the start of E-1 into an older epoch: firstOfPrev now
			// holds the first block of E-1.
			break
		}
		scanned++
		if scanned%warmUpProgressInterval == 0 {
			s.Log.Warn("WarmUpPaybackCache: scanning backwards for epoch-boundary start (epoch history unavailable)",
				"currentEpoch", currentEpoch, "scannedFromHead", scanned, "atBlock", b)
		}
	}
	if firstOfPrev > 0 {
		return firstOfPrev
	}
	// No block maps to epoch E-1 (e.g. the whole available history is a single
	// epoch, or FindBlockEpoch returns 0 throughout). Replay from block 1: complete
	// and deterministic. PrepareForBlock self-resets PaybackUsedMap per boundary so
	// the final accumulation is still exactly the current epoch's.
	s.Log.Warn("WarmUpPaybackCache: epoch-boundary history unavailable and no E-1 block found; replaying from block 1",
		"currentEpoch", currentEpoch, "head", head)
	return 1
}

// rawWarmUpReceipts loads the RAW stored receipts for block b and projects each
// into a minimal *types.Receipt carrying only the fields AddTransaction reads:
// Status, FeeRefund, and per-tx GasUsed. GasUsed is derived from the stored
// CumulativeGasUsed deltas exactly as Receipts.DeriveFields computes it
// (GasUsed[0] = CumulativeGasUsed[0]; GasUsed[i] = Cum[i] - Cum[i-1]). This
// avoids evmstore.GetReceipts entirely so the shared receipts LRU is never
// poisoned with zero-BlockHash receipts, and it never derives through
// DeriveFields so a normally-configured node never hits the cryptic count-mismatch
// Crit. Returns nil when the block's receipts are absent (e.g. TxIndex=false) or
// their count does not match the tx count; the caller fails closed on a nil result
// for a tx-bearing block (see WarmUpPaybackCache).
func (s *Service) rawWarmUpReceipts(b idx.Block, txCount int) types.Receipts {
	if txCount == 0 {
		return types.Receipts{}
	}
	stored, _ := s.store.evm.GetRawReceipts(b)
	if len(stored) == 0 {
		return nil
	}
	if len(stored) != txCount {
		// Mismatch — do not guess; the caller fails closed rather than
		// misattribute FeeRefund (a wrong-but-converging warm-up is still a
		// consensus risk).
		s.Log.Warn("WarmUpPaybackCache: stored receipt count does not match tx count",
			"block", b, "receipts", len(stored), "txs", txCount)
		return nil
	}
	receipts := make(types.Receipts, len(stored))
	prevCumulative := uint64(0)
	for i, r := range stored {
		rc := (*types.Receipt)(r)
		gasUsed := rc.CumulativeGasUsed
		if i > 0 {
			gasUsed = rc.CumulativeGasUsed - prevCumulative
		}
		prevCumulative = rc.CumulativeGasUsed
		receipts[i] = &types.Receipt{
			Status:    rc.Status,
			GasUsed:   gasUsed,
			FeeRefund: rc.FeeRefund,
		}
	}
	return receipts
}

// spillBlockEvents excludes first events which exceed MaxBlockGas
func spillBlockEvents(store *Store, block *inter.Block, network opera.Rules) (*inter.Block, inter.EventPayloads) {
	fullEvents := make(inter.EventPayloads, len(block.Events))
	if len(block.Events) == 0 {
		return block, fullEvents
	}
	gasPowerUsedSum := uint64(0)
	// iterate in reversed order
	for i := len(block.Events) - 1; ; i-- {
		id := block.Events[i]
		e := store.GetEventPayload(id)
		if e == nil {
			// log.Crit exits without flush — acceptable here because a missing
			// confirmed event indicates store corruption; block assembly cannot
			// continue with incomplete data.
			log.Crit("Block event not found", "event", id.String())
		}
		fullEvents[i] = e
		gasPowerUsedSum += e.GasPowerUsed()
		// stop if limit is exceeded, erase [:i] events
		if gasPowerUsedSum > network.Blocks.MaxBlockGas {
			// spill
			block.Events = block.Events[i+1:]
			fullEvents = fullEvents[i+1:]
			break
		}
		if i == 0 {
			break
		}
	}
	return block, fullEvents
}

func mergeCheaters(a, b lachesis.Cheaters) lachesis.Cheaters {
	if len(b) == 0 {
		return a
	}
	if len(a) == 0 {
		cp := make(lachesis.Cheaters, len(b))
		copy(cp, b)
		return cp
	}
	aSet := a.Set()
	merged := make(lachesis.Cheaters, 0, len(b)+len(a))
	for _, v := range a {
		merged = append(merged, v)
	}
	for _, v := range b {
		if _, ok := aSet[v]; !ok {
			merged = append(merged, v)
		}
	}
	return merged
}
