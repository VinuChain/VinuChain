package evmmodule

import (
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-opera/evmcore"
	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/inter/iblockproc"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/payback"
)

// stubChain implements evmcore.DummyChain for tests that don't need a
// parent header lookup. It returns a minimal EvmHeader keyed by the
// requested block index.
type stubChain struct {
	headers map[uint64]*evmcore.EvmHeader
}

func newStubChain() *stubChain {
	return &stubChain{headers: map[uint64]*evmcore.EvmHeader{}}
}

func (s *stubChain) GetHeader(_ common.Hash, num uint64) *evmcore.EvmHeader {
	return s.headers[num]
}

func newMemStateDB(t *testing.T) *state.StateDB {
	t.Helper()
	db := rawdb.NewMemoryDatabase()
	sdb, err := state.New(common.Hash{}, state.NewDatabase(db), nil)
	require.NoError(t, err)
	return sdb
}

func newPaybackCache() *payback.PaybackCache {
	return &payback.PaybackCache{
		PaybackUsedMap: make(map[common.Address]*big.Int),
		StakesMap:      make(map[idx.Epoch]*payback.EpochStakes),
	}
}

func baseRules() opera.Rules {
	r := opera.Rules{
		NetworkID: 27,
		Economy: opera.EconomyRules{
			MinGasPrice: big.NewInt(1_000_000_000), // 1 gwei
		},
		Blocks: opera.BlocksRules{
			MaxBlockGas:             20_000_000,
			MaxEmptyBlockSkipPeriod: inter.Timestamp(3 * time.Second),
		},
	}
	return r
}

// newStartedProcessor sets up an OperaEVMProcessor with a fresh state DB
// and an empty payback cache. If blockIdx > 0, a placeholder parent header
// is registered in the chain so Start() can look it up without panicking.
func newStartedProcessor(t *testing.T, blockIdx idx.Block, rules opera.Rules, chain *stubChain) (*OperaEVMProcessor, *state.StateDB) {
	t.Helper()
	if blockIdx > 0 {
		parentIdx := uint64(blockIdx - 1)
		if _, ok := chain.headers[parentIdx]; !ok {
			chain.headers[parentIdx] = &evmcore.EvmHeader{
				Hash:   common.Hash{},
				Number: big.NewInt(int64(parentIdx)),
			}
		}
	}
	sdb := newMemStateDB(t)
	pc := newPaybackCache()
	upgradeHeights := []opera.UpgradeHeight{{Upgrades: rules.Upgrades, Height: 0}}
	evmCfg := rules.EvmChainConfig(upgradeHeights)

	mod := New()
	p := mod.Start(
		iblockproc.BlockCtx{Idx: blockIdx, Time: inter.Timestamp(2_000_000_000)},
		sdb,
		chain,
		func(*types.Log) {},
		rules,
		opera.DefaultVMConfig,
		evmCfg,
		pc,
		idx.Epoch(10),
	)
	require.NotNil(t, p)
	op, ok := p.(*OperaEVMProcessor)
	require.True(t, ok, "Start must return an *OperaEVMProcessor")
	return op, sdb
}

// --- Module construction --------------------------------------------------

func TestNew_ReturnsNonNilModule(t *testing.T) {
	m := New()
	require.NotNil(t, m)
}

func TestStart_PopulatesProcessorFields(t *testing.T) {
	rules := baseRules()
	p, _ := newStartedProcessor(t, idx.Block(0), rules, newStubChain())
	require.Equal(t, idx.Block(0), p.block.Idx)
	require.Equal(t, idx.Epoch(10), p.epoch)
	require.Equal(t, rules.NetworkID, p.net.NetworkID)
	require.Equal(t, big.NewInt(0), p.blockIdx)
}

func TestStart_NonZeroBlockLooksUpParentHash(t *testing.T) {
	rules := baseRules()
	chain := newStubChain()
	parentHash := common.HexToHash("0xdeadbeef")
	chain.headers[6] = &evmcore.EvmHeader{Hash: parentHash, Number: big.NewInt(6)}

	p, _ := newStartedProcessor(t, idx.Block(7), rules, chain)
	require.Equal(t, parentHash, p.prevBlockHash,
		"prevBlockHash must be taken from GetHeader(_, blockIdx-1)")
}

func TestStart_BlockZeroSkipsParentLookup(t *testing.T) {
	p, _ := newStartedProcessor(t, idx.Block(0), baseRules(), newStubChain())
	require.Equal(t, common.Hash{}, p.prevBlockHash,
		"block 0 must leave prevBlockHash zero (no parent to look up)")
}

// --- evmBlockWith header construction ------------------------------------

func TestEvmBlockWith_LondonOff_NilBaseFee(t *testing.T) {
	rules := baseRules()
	rules.Upgrades.London = false
	p, _ := newStartedProcessor(t, idx.Block(5), rules, newStubChain())

	block := p.evmBlockWith(nil)
	require.Nil(t, block.BaseFee, "pre-London must produce nil BaseFee")
	require.Equal(t, uint64(math.MaxUint64), block.GasLimit)
	require.Equal(t, big.NewInt(5), block.Number)
}

func TestEvmBlockWith_LondonOn_BlockZero_BaseFeeEqualsMinGasPrice(t *testing.T) {
	rules := baseRules()
	rules.Upgrades.London = true
	p, _ := newStartedProcessor(t, idx.Block(0), rules, newStubChain())

	block := p.evmBlockWith(nil)
	require.NotNil(t, block.BaseFee)
	require.Equal(t, 0, block.BaseFee.Cmp(rules.Economy.MinGasPrice),
		"block 0 under London must default BaseFee to MinGasPrice")
	// Defensive copy: mutating the block's BaseFee must not mutate the rule.
	block.BaseFee.Add(block.BaseFee, big.NewInt(1))
	require.Equal(t, int64(1_000_000_000), rules.Economy.MinGasPrice.Int64(),
		"MinGasPrice rule must be defensively copied, not shared")
}

func TestEvmBlockWith_LondonOn_CachesBaseFeeAcrossCalls(t *testing.T) {
	rules := baseRules()
	rules.Upgrades.London = true
	p, _ := newStartedProcessor(t, idx.Block(0), rules, newStubChain())

	first := p.evmBlockWith(nil).BaseFee
	second := p.evmBlockWith(nil).BaseFee
	// cachedBaseFee is reused; both blocks carry the same *big.Int.
	require.True(t, first == second,
		"cachedBaseFee must be computed once and reused")
}

func TestEvmBlockWith_UsesCurrentGasUsed(t *testing.T) {
	p, _ := newStartedProcessor(t, idx.Block(1), baseRules(), newStubChain())
	p.gasUsed = 55_555
	block := p.evmBlockWith(nil)
	require.Equal(t, uint64(55_555), block.GasUsed)
}

// --- Finalize empty-block path -------------------------------------------

func TestFinalize_NoTxs_ReturnsEmptyBlock(t *testing.T) {
	p, _ := newStartedProcessor(t, idx.Block(1), baseRules(), newStubChain())

	block, skipped, receipts := p.Finalize()

	require.NotNil(t, block)
	require.Empty(t, block.Transactions)
	require.Empty(t, skipped)
	require.Empty(t, receipts)
	require.NotEqual(t, common.Hash{}, block.Root,
		"empty state still has a computed trie root")
}

func TestFinalize_PopulatesRootFromStateCommit(t *testing.T) {
	rules := baseRules()
	p, sdb := newStartedProcessor(t, idx.Block(1), rules, newStubChain())

	// Touch state so Commit produces a non-empty trie root.
	addr := common.HexToAddress("0xabc")
	sdb.SetBalance(addr, big.NewInt(123))

	block, _, _ := p.Finalize()
	require.NotEqual(t, common.Hash{}, block.Root)
}

// --- Execute with zero txs — should be a no-op pass-through --------------

func TestExecute_EmptyTxList_DoesNotPanic(t *testing.T) {
	p, _ := newStartedProcessor(t, idx.Block(1), baseRules(), newStubChain())
	require.NotPanics(t, func() {
		receipts := p.Execute(nil)
		require.Empty(t, receipts)
	})
}
