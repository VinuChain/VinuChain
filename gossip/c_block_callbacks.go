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
	reexecCache, err := payback.NewPaybackCache(s.paybackCache.GetStore())
	if err != nil {
		log.Crit("Failed to create re-execution PaybackCache", "err", err)
	}

	prev := s.store.GetBlock(from)
	for b := from + 1; b <= to; b++ {
		block := s.store.GetBlock(b)
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
		es := s.store.GetHistoryEpochState(s.store.FindBlockEpoch(b))

		evmProcessor := blockProc.EVMModule.Start(blockCtx, statedb, evmStateReader, func(t *types.Log) {}, es.Rules, es.Rules.EvmChainConfig(upgradeHeights), reexecCache)
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
		return b
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
