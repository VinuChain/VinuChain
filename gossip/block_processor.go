package gossip

import (
	"encoding/json"
	"fmt"
	"runtime/debug"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/dag"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/inter/pos"
	"github.com/Fantom-foundation/lachesis-base/lachesis"
	"github.com/Fantom-foundation/lachesis-base/utils/workers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Fantom-foundation/go-opera/evmcore"
	"github.com/Fantom-foundation/go-opera/gossip/blockproc"
	"github.com/Fantom-foundation/go-opera/gossip/blockproc/verwatcher"
	"github.com/Fantom-foundation/go-opera/gossip/emitter"
	"github.com/Fantom-foundation/go-opera/gossip/evmstore"
	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/inter/iblockproc"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/contracts/sfc"
	"github.com/Fantom-foundation/go-opera/payback"
	"github.com/Fantom-foundation/go-opera/txtrace"
	"github.com/Fantom-foundation/go-opera/utils"
)

// BlockProcessor encapsulates the per-block processing state and injected
// dependencies that were previously captured by the closure returned from
// consensusCallbackBeginBlockFn.
type BlockProcessor struct {
	// Injected dependencies (immutable after construction)
	parallelTasks *workers.Workers
	wg            *sync.WaitGroup
	blockBusyFlag *uint32
	store         *Store
	blockProc     BlockProc
	txIndex       bool
	feed          *ServiceFeed
	emitters      *[]*emitter.Emitter
	verWatcher    *verwatcher.VerWarcher
	pc            *payback.PaybackCache

	// Per-block mutable state (set during Begin/processing)
	cBlock            *lachesis.Block
	bs                iblockproc.BlockState
	es                iblockproc.EpochState
	statedb           *state.StateDB
	evmStateReader    *EvmStateReader
	eventProcessor    blockproc.ConfirmedEventsProcessor
	start             time.Time
	atroposTime       inter.Timestamp
	atroposDegenerate bool
	confirmedEvents   hash.OrderedEvents
	mpsCheatersMap    map[idx.ValidatorID]struct{}
	blockCtx          iblockproc.BlockCtx
	skipBlock         bool
	sealer            blockproc.SealerProcessor
	sealing           bool
	txListener        blockproc.TxListener
	evmProcessor      blockproc.EVMProcessor
	executionStart    time.Time
	preInternalTxs    types.Transactions
	newValidators     *pos.Validators
}

func newBlockProcessor(
	parallelTasks *workers.Workers,
	wg *sync.WaitGroup,
	blockBusyFlag *uint32,
	store *Store,
	blockProc BlockProc,
	txIndex bool,
	feed *ServiceFeed,
	emitters *[]*emitter.Emitter,
	verWatcher *verwatcher.VerWarcher,
	pc *payback.PaybackCache,
) *BlockProcessor {
	return &BlockProcessor{
		parallelTasks: parallelTasks,
		wg:            wg,
		blockBusyFlag: blockBusyFlag,
		store:         store,
		blockProc:     blockProc,
		txIndex:       txIndex,
		feed:          feed,
		emitters:      emitters,
		verWatcher:    verWatcher,
		pc:            pc,
	}
}

// reset zeroes all per-block mutable state so that no field from the previous
// block leaks into the next invocation. Fields that Begin() unconditionally
// reassigns (e.g. cBlock, bs, es, statedb) are still reset here for defence
// in depth — a future refactor that reorders Begin() must not silently inherit
// stale values.
func (bp *BlockProcessor) reset() {
	bp.cBlock = nil
	bp.bs = iblockproc.BlockState{}
	bp.es = iblockproc.EpochState{}
	bp.statedb = nil
	bp.evmStateReader = nil
	bp.eventProcessor = nil
	bp.start = time.Time{}
	bp.atroposTime = 0
	bp.atroposDegenerate = false
	bp.confirmedEvents = nil
	bp.mpsCheatersMap = nil
	bp.blockCtx = iblockproc.BlockCtx{}
	bp.skipBlock = false
	bp.sealer = nil
	bp.sealing = false
	bp.txListener = nil
	bp.evmProcessor = nil
	bp.executionStart = time.Time{}
	bp.preInternalTxs = nil
	bp.newValidators = nil
}

// Begin initializes per-block state and returns the ApplyEvent/EndBlock callbacks.
func (bp *BlockProcessor) Begin(cBlock *lachesis.Block) lachesis.BlockCallbacks {
	bp.wg.Wait()
	bp.reset()
	bp.start = time.Now()
	bp.cBlock = cBlock

	// Note: take copies to avoid race conditions with API calls
	bp.bs = bp.store.GetBlockState().Copy()
	bp.es = bp.store.GetEpochState().Copy()

	// merge cheaters to ensure that every cheater will get punished even if only previous (not current) Atropos observed a doublesign
	// this feature is needed because blocks may be skipped even if cheaters list isn't empty
	// otherwise cheaters would get punished after a first block where cheaters were observed
	bp.bs.EpochCheaters = mergeCheaters(bp.bs.EpochCheaters, cBlock.Cheaters)

	// Get stateDB
	statedb, err := bp.store.evm.StateDB(bp.bs.FinalizedStateRoot)
	if err != nil {
		// log.Crit exits without flush — acceptable here because a missing
		// state root means the on-disk state is corrupted and no further
		// blocks can be processed; an unclean exit is the safest outcome.
		log.Crit("Failed to open StateDB", "err", err)
	}
	bp.statedb = statedb
	bp.evmStateReader = &EvmStateReader{
		ServiceFeed: bp.feed,
		store:       bp.store,
	}

	bp.eventProcessor = bp.blockProc.EventsModule.Start(bp.bs, bp.es)

	bp.atroposTime = bp.bs.LastBlock.Time + 1
	bp.atroposDegenerate = true
	// events with txs
	bp.confirmedEvents = make(hash.OrderedEvents, 0, 3*bp.es.Validators.Len())

	bp.mpsCheatersMap = make(map[idx.ValidatorID]struct{})
	bp.newValidators = nil

	return lachesis.BlockCallbacks{
		ApplyEvent: bp.applyEvent,
		EndBlock:   bp.endBlock,
	}
}

func (bp *BlockProcessor) applyEvent(_e dag.Event) {
	e, ok := _e.(inter.EventI)
	if !ok {
		log.Crit("applyEvent received non-inter.EventI event", "type", fmt.Sprintf("%T", _e))
		return
	}
	if bp.cBlock.Atropos == e.ID() {
		bp.atroposTime = e.MedianTime()
		bp.atroposDegenerate = false
	}
	if e.AnyTxs() {
		bp.confirmedEvents = append(bp.confirmedEvents, e.ID())
	}
	if e.AnyMisbehaviourProofs() {
		payload := bp.store.GetEventPayload(e.ID())
		if payload == nil {
			log.Crit("Event payload not found for confirmed event with misbehaviour proofs", "id", e.ID())
			return
		}
		reportCheater := func(reporter, cheater idx.ValidatorID) {
			bp.mpsCheatersMap[cheater] = struct{}{}
		}
		mps := payload.MisbehaviourProofs()
		for _, mp := range mps {
			// self-contained parts of proofs are already checked by the checkers
			if proof := mp.BlockVoteDoublesign; proof != nil {
				reportCheater(e.Creator(), proof.Pair[0].Signed.Locator.Creator)
			}
			if proof := mp.EpochVoteDoublesign; proof != nil {
				reportCheater(e.Creator(), proof.Pair[0].Signed.Locator.Creator)
			}
			if proof := mp.EventsDoublesign; proof != nil {
				reportCheater(e.Creator(), proof.Pair[0].Locator.Creator)
			}
			if proof := mp.WrongBlockVote; proof != nil {
				// all other votes are the same, see MinAccomplicesForProof
				if proof.WrongEpoch {
					actualBlockEpoch := bp.store.FindBlockEpoch(proof.Block)
					if actualBlockEpoch != 0 && actualBlockEpoch != proof.Pals[0].Val.Epoch {
						for _, pal := range proof.Pals {
							reportCheater(e.Creator(), pal.Signed.Locator.Creator)
						}
					}
				} else {
					actualRecord := bp.store.GetFullBlockRecord(proof.Block)
					if actualRecord != nil && proof.GetVote(0) != actualRecord.Hash() {
						for _, pal := range proof.Pals {
							reportCheater(e.Creator(), pal.Signed.Locator.Creator)
						}
					}
				}
			}
			if proof := mp.WrongEpochVote; proof != nil {
				// all other votes are the same, see MinAccomplicesForProof
				vote := proof.Pals[0]
				actualRecord := bp.store.GetFullEpochRecord(vote.Val.Epoch)
				if actualRecord == nil {
					continue
				}
				if vote.Val.Vote != actualRecord.Hash() {
					for _, pal := range proof.Pals {
						reportCheater(e.Creator(), pal.Signed.Locator.Creator)
					}
				}
			}
		}
	}
	bp.eventProcessor.ProcessConfirmedEvent(e)
	for _, em := range *bp.emitters {
		em.OnEventConfirmed(e)
	}
}

// buildBlockContext creates the block context and decides whether to skip the block.
// Returns true if the block should be skipped (and has already been persisted).
func (bp *BlockProcessor) buildBlockContext() bool {
	if bp.atroposTime <= bp.bs.LastBlock.Time {
		bp.atroposTime = bp.bs.LastBlock.Time + 1
	}
	bp.blockCtx = iblockproc.BlockCtx{
		Idx:     bp.bs.LastBlock.Idx + 1,
		Time:    bp.atroposTime,
		Atropos: bp.cBlock.Atropos,
	}
	// Note:
	// it's possible that a previous Atropos observes current Atropos (1)
	// (even stronger statement is true - it's possible that current Atropos is equal to a previous Atropos).
	// (1) is true when and only when ApplyEvent wasn't called.
	// In other words, we should assume that every non-cheater root may be elected as an Atropos in any order,
	// even if typically every previous Atropos happened-before current Atropos
	// We have to skip block in case (1) to ensure that every block ID is unique.
	// If Atropos ID wasn't used as a block ID, it wouldn't be required.
	bp.skipBlock = bp.atroposDegenerate
	// Check if empty block should be pruned
	emptyBlock := bp.confirmedEvents.Len() == 0 && bp.cBlock.Cheaters.Len() == 0
	bp.skipBlock = bp.skipBlock || (emptyBlock && bp.blockCtx.Time < bp.bs.LastBlock.Time+bp.es.Rules.Blocks.MaxEmptyBlockSkipPeriod)
	// Finalize the progress of eventProcessor
	bp.bs = bp.eventProcessor.Finalize(bp.blockCtx, bp.skipBlock)
	{ // CONSENSUS-CRITICAL: sort deterministically after map iteration.
		// Go map iteration is non-deterministic; without this sort, different
		// nodes would produce different cheater orderings, causing state divergence.
		mpsCheaters := make(lachesis.Cheaters, 0, len(bp.mpsCheatersMap))
		for vid := range bp.mpsCheatersMap {
			mpsCheaters = append(mpsCheaters, vid)
		}
		sort.Slice(mpsCheaters, func(i, j int) bool {
			a, b := mpsCheaters[i], mpsCheaters[j]
			return a < b
		})
		bp.bs.EpochCheaters = mergeCheaters(bp.bs.EpochCheaters, mpsCheaters)
	}
	if bp.skipBlock {
		// save the latest block state even if block is skipped
		bp.store.SetBlockEpochState(bp.bs, bp.es)
		log.Debug("Frame is skipped", "atropos", bp.cBlock.Atropos.String())
		return true
	}
	return false
}

// initProcessors sets up the sealer, txListener, and evmProcessor for the current block.
func (bp *BlockProcessor) initProcessors() {
	bp.sealer = bp.blockProc.SealerModule.Start(bp.blockCtx, bp.bs, bp.es)
	bp.sealing = bp.sealer.EpochSealing()
	bp.txListener = bp.blockProc.TxListenerModule.Start(bp.blockCtx, bp.bs, bp.es, bp.statedb)
	onNewLogAll := func(l *types.Log) {
		bp.txListener.OnNewLog(l)
		// Note: it's possible for logs to get indexed twice by BR and block processing
		if bp.verWatcher != nil {
			bp.verWatcher.OnNewLog(l)
		}
	}

	// skip LLR block/epoch deciding if not activated
	if !bp.es.Rules.Upgrades.Llr {
		bp.store.ModifyLlrState(func(llrs *LlrState) {
			if llrs.LowestBlockToDecide == bp.blockCtx.Idx {
				llrs.LowestBlockToDecide++
			}
			if bp.sealing && bp.es.Epoch+1 == llrs.LowestEpochToDecide {
				llrs.LowestEpochToDecide++
			}
		})
	}

	vmCfg := opera.DefaultVMConfig
	if bp.store.TxTraceStore() != nil {
		vmCfg.Debug = true
		vmCfg.Tracer = txtrace.NewTraceStructLogger(bp.store.TxTraceStore())
	}
	bp.evmProcessor = bp.blockProc.EVMModule.Start(bp.blockCtx, bp.statedb, bp.evmStateReader, onNewLogAll, bp.es.Rules, vmCfg, bp.es.Rules.EvmChainConfig(bp.store.GetUpgradeHeights()), bp.pc, bp.es.Epoch)
	bp.executionStart = time.Now()
}

// executePreInternalTxs executes pre-internal transactions and finalizes the tx listener.
func (bp *BlockProcessor) executePreInternalTxs() {
	bp.preInternalTxs = bp.blockProc.PreTxTransactor.PopInternalTxs(bp.blockCtx, bp.bs, bp.es, bp.sealing, bp.statedb)
	preInternalReceipts := bp.evmProcessor.Execute(bp.preInternalTxs)
	bp.bs = bp.txListener.Finalize()
	for _, r := range preInternalReceipts {
		if r.Status == 0 {
			log.Crit("Pre-internal transaction reverted", "txid", r.TxHash.String(), "block", bp.blockCtx.Idx)
		}
	}
}

// sealEpochIfNeeded conditionally seals the epoch if the sealer indicates it.
func (bp *BlockProcessor) sealEpochIfNeeded() {
	if !bp.sealing {
		return
	}
	bp.sealer.Update(bp.bs, bp.es)
	prevUpg := bp.es.Rules.Upgrades
	bp.bs, bp.es = bp.sealer.SealEpoch()
	if bp.es.Rules.Upgrades != prevUpg {
		bp.store.AddUpgradeHeight(opera.UpgradeHeight{
			Upgrades: bp.es.Rules.Upgrades,
			Height:   bp.blockCtx.Idx + 1,
		})
	}
	// SFC V2 bytecode upgrade ordering
	//
	// This upgrade runs after pre-internal txs (which executed against V1)
	// but before post-internal txs and user txs in the same block. Within
	// a single epoch-sealing block the execution timeline is:
	//
	//   1. Pre-internal txs   → executed against SFC V1 bytecode
	//   2. SealEpoch          → new rules take effect, V2 bytecode installed here
	//   3. Post-internal txs  → executed against SFC V2 bytecode
	//   4. User txs           → executed against SFC V2 bytecode
	//
	// This is safe provided V2 is backward-compatible with the state left by
	// V1 pre-internal calls in step 1. Specifically, V2 must not reinterpret
	// storage slots written by V1, must accept the same delegation/staking
	// state shape, and must not change function selectors used by the driver
	// in post-internal transactions. Any V2 candidate that breaks these
	// invariants would corrupt in-block state.
	if bp.es.Rules.Upgrades.SfcV2 && !prevUpg.SfcV2 {
		log.Info("Applying SFC V2 bytecode upgrade", "block", bp.blockCtx.Idx)
		bp.statedb.SetCode(sfc.ContractAddress, sfc.GetContractBin())
	}
	bp.store.SetBlockEpochState(bp.bs, bp.es)
	bp.newValidators = bp.es.Validators
	bp.txListener.Update(bp.bs, bp.es)
}

// processBlock executes post-internal transactions, finalizes EVM state,
// writes the block to the store, updates metrics, and broadcasts notifications.
func (bp *BlockProcessor) processBlock() {
	// Execute post-internal transactions
	internalTxs := bp.blockProc.PostTxTransactor.PopInternalTxs(bp.blockCtx, bp.bs, bp.es, bp.sealing, bp.statedb)
	internalReceipts := bp.evmProcessor.Execute(internalTxs)
	for _, r := range internalReceipts {
		if r.Status == 0 {
			log.Crit("Internal transaction reverted", "txid", r.TxHash.String(), "block", bp.blockCtx.Idx)
		}
	}

	// INVARIANT: lachesis-base delivers events to applyEvent in undefined
	// order. This sort ensures deterministic event ordering (by Lamport time)
	// across all validators regardless of delivery order.
	sort.Sort(bp.confirmedEvents)

	// new block
	var block = &inter.Block{
		Time:    bp.blockCtx.Time,
		Atropos: bp.cBlock.Atropos,
		Events:  hash.Events(bp.confirmedEvents),
	}
	for _, tx := range append(bp.preInternalTxs, internalTxs...) {
		block.Txs = append(block.Txs, tx.Hash())
	}

	block, blockEvents := spillBlockEvents(bp.store, block, bp.es.Rules)
	txs := make(types.Transactions, 0, blockEvents.Len()*10)
	for _, e := range blockEvents {
		txs = append(txs, e.Txs()...)
	}

	bp.evmProcessor.Execute(txs)

	evmBlock, skippedTxs, allReceipts := bp.evmProcessor.Finalize()
	block.SkippedTxs = skippedTxs
	block.Root = hash.Hash(evmBlock.Root)
	block.GasUsed = evmBlock.GasUsed

	// memorize event position of each tx
	txPositions := make(map[common.Hash]ExtendedTxPosition)
	for _, e := range blockEvents {
		for i, tx := range e.Txs() {
			// If tx was met in multiple events, then assign to first ordered event
			if _, ok := txPositions[tx.Hash()]; ok {
				continue
			}
			txPositions[tx.Hash()] = ExtendedTxPosition{
				TxPosition: evmstore.TxPosition{
					Event:       e.ID(),
					EventOffset: uint32(i),
				},
				EventCreator: e.Creator(),
			}
		}
	}
	// memorize block position of each tx
	for i, tx := range evmBlock.Transactions {
		// not skipped txs only
		position := txPositions[tx.Hash()]
		position.Block = bp.blockCtx.Idx
		position.BlockOffset = uint32(i)
		txPositions[tx.Hash()] = position
	}

	// When txs are skipped, the TransactionPosition recorded by the tracer
	// (via statedb.TxIndex) no longer matches the final BlockOffset. Correct
	// the stored trace bytes for every executed tx so the RPC returns the right
	// position.
	if len(skippedTxs) > 0 && bp.store.TxTraceStore() != nil {
		for _, tx := range evmBlock.Transactions {
			txBytes := bp.store.TxTraceStore().GetTx(tx.Hash())
			if txBytes == nil {
				continue
			}
			var traces []txtrace.ActionTrace
			if err := json.Unmarshal(txBytes, &traces); err != nil {
				continue
			}
			pos := uint64(txPositions[tx.Hash()].BlockOffset)
			for i := range traces {
				traces[i].TransactionPosition = pos
			}
			if corrected, err := json.Marshal(traces); err == nil {
				bp.store.TxTraceStore().SetTxTrace(tx.Hash(), corrected)
			}
		}
	}

	// call OnNewReceipt
	for i, r := range allReceipts {
		creator := txPositions[r.TxHash].EventCreator
		if creator != 0 && bp.es.Validators.Get(creator) == 0 {
			creator = 0
		}
		bp.txListener.OnNewReceipt(evmBlock.Transactions[i], r, creator)
	}
	bp.bs = bp.txListener.Finalize()
	bp.bs.FinalizedStateRoot = block.Root
	// At this point, block state is finalized

	// Build index for not skipped txs
	if bp.txIndex {
		for _, tx := range evmBlock.Transactions {
			// not skipped txs only
			bp.store.evm.SetTxPosition(tx.Hash(), txPositions[tx.Hash()].TxPosition)
		}

		// Index receipts
		// Note: it's possible for receipts to get indexed twice by BR and block processing
		if allReceipts.Len() != 0 {
			bp.store.evm.SetReceipts(bp.blockCtx.Idx, allReceipts)
			for _, r := range allReceipts {
				bp.store.evm.IndexLogs(r.Logs...)
			}
		}
	}
	for _, tx := range append(bp.preInternalTxs, internalTxs...) {
		bp.store.evm.SetTx(tx.Hash(), tx)
	}

	bp.bs.LastBlock = bp.blockCtx
	bp.bs.CheatersWritten = uint32(bp.bs.EpochCheaters.Len())
	if bp.sealing {
		bp.store.SetHistoryBlockEpochState(bp.es.Epoch, bp.bs, bp.es)
		bp.store.SetEpochBlock(bp.blockCtx.Idx+1, bp.es.Epoch)
	}
	bp.store.SetBlock(bp.blockCtx.Idx, block)
	bp.store.SetBlockIndex(block.Atropos, bp.blockCtx.Idx)
	bp.store.SetBlockEpochState(bp.bs, bp.es)
	bp.store.EvmStore().SetCachedEvmBlock(bp.blockCtx.Idx, evmBlock)
	updateLowestBlockToFill(bp.blockCtx.Idx, bp.store)
	updateLowestEpochToFill(bp.es.Epoch, bp.store)

	// Update the metrics touched during block processing
	accountReadTimer.Update(bp.statedb.AccountReads)
	storageReadTimer.Update(bp.statedb.StorageReads)
	accountUpdateTimer.Update(bp.statedb.AccountUpdates)
	storageUpdateTimer.Update(bp.statedb.StorageUpdates)
	snapshotAccountReadTimer.Update(bp.statedb.SnapshotAccountReads)
	snapshotStorageReadTimer.Update(bp.statedb.SnapshotStorageReads)
	accountHashTimer.Update(bp.statedb.AccountHashes)
	storageHashTimer.Update(bp.statedb.StorageHashes)
	triehash := bp.statedb.AccountHashes + bp.statedb.StorageHashes
	trieproc := bp.statedb.SnapshotAccountReads + bp.statedb.AccountReads + bp.statedb.AccountUpdates
	trieproc += bp.statedb.SnapshotStorageReads + bp.statedb.StorageReads + bp.statedb.StorageUpdates
	blockExecutionTimer.Update(time.Since(bp.executionStart) - trieproc - triehash)

	// Update the metrics touched by new block
	headBlockGauge.Update(int64(bp.blockCtx.Idx))
	headHeaderGauge.Update(int64(bp.blockCtx.Idx))
	headFastBlockGauge.Update(int64(bp.blockCtx.Idx))

	// Notify about new block
	if bp.feed != nil {
		bp.feed.newBlock.Send(evmcore.ChainHeadNotify{Block: evmBlock})
		var logs []*types.Log
		for _, r := range allReceipts {
			for _, l := range r.Logs {
				logs = append(logs, l)
			}
		}
		bp.feed.newLogs.Send(logs)
	}

	commitStart := time.Now()
	bp.store.commitEVM(false)

	// Update the metrics touched during block commit
	accountCommitTimer.Update(bp.statedb.AccountCommits)
	storageCommitTimer.Update(bp.statedb.StorageCommits)
	snapshotCommitTimer.Update(bp.statedb.SnapshotCommits)
	blockWriteTimer.Update(time.Since(commitStart) - bp.statedb.AccountCommits - bp.statedb.StorageCommits - bp.statedb.SnapshotCommits)
	blockInsertTimer.UpdateSince(bp.start)

	now := time.Now()
	blockAge := now.Sub(block.Time.Time())
	log.Info("New block", "index", bp.blockCtx.Idx, "id", block.Atropos, "gas_used",
		evmBlock.GasUsed, "txs", fmt.Sprintf("%d/%d", len(evmBlock.Transactions), len(block.SkippedTxs)),
		"age", utils.PrettyDuration(blockAge), "t", utils.PrettyDuration(now.Sub(bp.start)))
	blockAgeGauge.Update(int64(blockAge.Nanoseconds()))
}

// dispatchBlock runs processBlock either asynchronously (if there are confirmed
// events to process) or synchronously.
func (bp *BlockProcessor) dispatchBlock() {
	if bp.confirmedEvents.Len() != 0 {
		atomic.StoreUint32(bp.blockBusyFlag, 1)
		bp.wg.Add(1)
		err := bp.parallelTasks.Enqueue(func() {
			defer atomic.StoreUint32(bp.blockBusyFlag, 0)
			defer bp.wg.Done()
			defer func() {
				if r := recover(); r != nil {
					// log.Crit exits without flush — acceptable here because a panic
					// in block processing indicates corrupted consensus state; the
					// node cannot safely continue.
					log.Crit("Panic in block processing", "err", r, "stack", string(debug.Stack()))
				}
			}()
			bp.processBlock()
		})
		if err != nil {
			// log.Crit exits without flush — acceptable here because the
			// worker pool is full or shut down; block processing cannot
			// proceed and the node must terminate.
			log.Crit("Failed to enqueue block processing", "err", err)
		}
	} else {
		defer func() {
			if r := recover(); r != nil {
				log.Crit("Panic in block processing", "err", r, "stack", string(debug.Stack()))
			}
		}()
		bp.processBlock()
	}
}

func (bp *BlockProcessor) endBlock() *pos.Validators {
	if bp.buildBlockContext() {
		return nil
	}
	bp.initProcessors()
	bp.executePreInternalTxs()
	bp.sealEpochIfNeeded()
	bp.dispatchBlock()
	return bp.newValidators
}
