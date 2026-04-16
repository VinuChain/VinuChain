package evmmodule

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"

	"github.com/Fantom-foundation/go-opera/evmcore"
	"github.com/Fantom-foundation/go-opera/gossip/blockproc"
	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/inter/iblockproc"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/contracts/evmwriter"
	"github.com/Fantom-foundation/go-opera/payback"
	"github.com/Fantom-foundation/go-opera/utils"
)

type EVMModule struct{}

func New() *EVMModule {
	return &EVMModule{}
}

func (p *EVMModule) Start(block iblockproc.BlockCtx, statedb *state.StateDB, reader evmcore.DummyChain, onNewLog func(*types.Log), net opera.Rules, vmConfig vm.Config, evmCfg *params.ChainConfig, paybackCache *payback.PaybackCache, epoch idx.Epoch) blockproc.EVMProcessor {
	var prevBlockHash common.Hash
	if block.Idx != 0 {
		prevBlockHash = reader.GetHeader(common.Hash{}, uint64(block.Idx-1)).Hash
	}
	return &OperaEVMProcessor{
		block:         block,
		reader:        reader,
		statedb:       statedb,
		onNewLog:      onNewLog,
		net:           net,
		vmConfig:      vmConfig,
		evmCfg:        evmCfg,
		epoch:         epoch,
		blockIdx:      utils.U64toBig(uint64(block.Idx)),
		prevBlockHash: prevBlockHash,
		paybackCache:  paybackCache,
	}
}

type OperaEVMProcessor struct {
	block    iblockproc.BlockCtx
	reader   evmcore.DummyChain
	statedb  *state.StateDB
	onNewLog func(*types.Log)
	net      opera.Rules
	vmConfig vm.Config
	evmCfg   *params.ChainConfig
	epoch    idx.Epoch

	blockIdx      *big.Int
	prevBlockHash common.Hash

	gasUsed uint64

	cachedBaseFee *big.Int // computed once from parent header, reused in Finalize

	incomingTxs types.Transactions
	skippedTxs  []uint32
	receipts    types.Receipts

	paybackCache *payback.PaybackCache
}

func (p *OperaEVMProcessor) evmBlockWith(txs types.Transactions) *evmcore.EvmBlock {
	if p.cachedBaseFee == nil && p.net.Upgrades.London {
		p.cachedBaseFee = p.net.Economy.MinGasPrice
		if p.block.Idx != 0 {
			parentHeader := p.reader.GetHeader(p.prevBlockHash, uint64(p.block.Idx-1))
			if parentHeader != nil {
				// Build the ETH-formatted parent header for CalcBaseFee, overriding
				// GasLimit to MaxBlockGas to handle old stored blocks that have
				// math.MaxUint64 as their gas limit.
				parentEth := parentHeader.EthHeader()
				parentEth.GasLimit = p.net.Blocks.MaxBlockGas
				computed := misc.CalcBaseFee(p.evmCfg, parentEth)
				if computed.Cmp(p.net.Economy.MinGasPrice) < 0 {
					computed = new(big.Int).Set(p.net.Economy.MinGasPrice)
				}
				p.cachedBaseFee = computed
			}
		}
	}
	baseFee := p.cachedBaseFee
	h := &evmcore.EvmHeader{
		Number:     p.blockIdx,
		Hash:       common.Hash(p.block.Atropos),
		ParentHash: p.prevBlockHash,
		Root:       common.Hash{},
		Time:       p.block.Time,
		Coinbase:   common.Address{},
		GasLimit:   p.net.Blocks.MaxBlockGas,
		GasUsed:    p.gasUsed,
		BaseFee:    baseFee,
	}

	return evmcore.NewEvmBlock(h, txs)
}

func (p *OperaEVMProcessor) Execute(txs types.Transactions) types.Receipts {
	evmProcessor := evmcore.NewStateProcessor(p.evmCfg, p.reader)
	txsOffset := uint(len(p.incomingTxs))

	// Prepare payback cache with the epoch known at block processor level,
	// not from a live store query (which can return a different epoch on
	// sealing blocks or during re-execution).
	p.paybackCache.PrepareForBlock(p.epoch, p.net, p.block.Time.Time())

	// Update the EvmWriter's payback proxy address from current rules so
	// isSystemContract protects the proxy from swapCode/setStorage.
	if pc, ok := opera.DefaultVMConfig.StatePrecompiles[evmwriter.ContractAddress].(*evmwriter.PreCompiledContract); ok {
		pc.SetPaybackProxyAddr(p.net.Economy.QuotaCacheAddress)
		pc.SetElemont(p.net.Upgrades.Elemont)
	}

	// Process txs
	evmBlock := p.evmBlockWith(txs)
	receipts, _, skipped, err := evmProcessor.Process(evmBlock, p.statedb, p.vmConfig, &p.gasUsed, func(l *types.Log, _ *state.StateDB) {
		// Note: l.Index is properly set before
		l.TxIndex += txsOffset
		p.onNewLog(l)
	}, p.paybackCache, p.net.Economy.MinGasPrice)
	if err != nil {
		// log.Crit exits without flush — acceptable here because an EVM
		// internal error (distinct from a reverted tx) signals a consensus-
		// breaking fault; the node must halt to avoid a chain split.
		log.Crit("EVM internal error", "err", err)
	}

	if txsOffset > 0 {
		for i, n := range skipped {
			skipped[i] = n + uint32(txsOffset)
		}
		for _, r := range receipts {
			r.TransactionIndex += txsOffset
		}
	}

	p.incomingTxs = append(p.incomingTxs, txs...)
	p.skippedTxs = append(p.skippedTxs, skipped...)
	p.receipts = append(p.receipts, receipts...)

	return receipts
}

func (p *OperaEVMProcessor) Finalize() (evmBlock *evmcore.EvmBlock, skippedTxs []uint32, receipts types.Receipts) {
	// FinishBlock is called exactly once per logical block at Finalize time,
	// not per Execute() call, to prevent blkCtx teardown between phases.
	if p.paybackCache != nil && p.paybackCache.GetStore() != nil {
		p.paybackCache.FinishBlock()
	}
	evmBlock = p.evmBlockWith(
		// Filter skipped transactions. Receipts are filtered already
		inter.FilterSkippedTxs(p.incomingTxs, p.skippedTxs),
	)
	skippedTxs = p.skippedTxs
	receipts = p.receipts

	// Get state root
	newStateHash, err := p.statedb.Commit(true)
	if err != nil {
		// log.Crit exits without flush — acceptable here because a failed
		// state commit means the trie is inconsistent; persisting partial
		// state would corrupt the database.
		log.Crit("Failed to commit state", "err", err)
	}
	evmBlock.Root = newStateHash

	return
}
