package evmmodule

import (
	"crypto/ecdsa"
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
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

func forkTestRules(shanghai, cancun bool) opera.Rules {
	rules := opera.VinuChainTestNetRules()
	rules.Upgrades.Shanghai = shanghai
	rules.Upgrades.Cancun = cancun
	rules.Economy.MinGasPrice = big.NewInt(1)
	return rules
}

func fundedSender(t *testing.T, sdb *state.StateDB) (*ecdsa.PrivateKey, common.Address) {
	t.Helper()
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	addr := crypto.PubkeyToAddress(key.PublicKey)
	sdb.SetBalance(addr, big.NewInt(1_000_000_000_000_000_000))
	return key, addr
}

func signCreationTx(t *testing.T, p *OperaEVMProcessor, key *ecdsa.PrivateKey, nonce uint64, code []byte) *types.Transaction {
	t.Helper()
	tx := types.NewContractCreation(nonce, big.NewInt(0), 1_000_000, p.net.Economy.MinGasPrice, code)
	signed, err := types.SignTx(tx, types.MakeSigner(p.evmCfg, p.blockIdx), key)
	require.NoError(t, err)
	return signed
}

func signCallTx(t *testing.T, p *OperaEVMProcessor, key *ecdsa.PrivateKey, nonce uint64, to common.Address) *types.Transaction {
	t.Helper()
	tx := types.NewTransaction(nonce, to, big.NewInt(0), 1_000_000, p.net.Economy.MinGasPrice, nil)
	signed, err := types.SignTx(tx, types.MakeSigner(p.evmCfg, p.blockIdx), key)
	require.NoError(t, err)
	return signed
}

func executeSingleTx(t *testing.T, p *OperaEVMProcessor, tx *types.Transaction) *types.Receipt {
	t.Helper()
	receipts := p.Execute(types.Transactions{tx})
	_, skipped, finalReceipts := p.Finalize()
	require.Empty(t, skipped)
	require.Len(t, receipts, 1)
	require.Len(t, finalReceipts, 1)
	return finalReceipts[0]
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

func TestExecuteShanghaiPush0ThroughProcessor(t *testing.T) {
	initcode := []byte{byte(vm.PUSH0), byte(vm.PUSH0), byte(vm.RETURN)}

	preShanghai, preState := newStartedProcessor(t, idx.Block(0), forkTestRules(false, false), newStubChain())
	preKey, _ := fundedSender(t, preState)
	preReceipt := executeSingleTx(t, preShanghai, signCreationTx(t, preShanghai, preKey, 0, initcode))
	require.Equal(t, types.ReceiptStatusFailed, preReceipt.Status,
		"PUSH0 must fail before Shanghai when executed through the VinuChain processor")

	shanghai, shanghaiState := newStartedProcessor(t, idx.Block(0), forkTestRules(true, false), newStubChain())
	shanghaiKey, _ := fundedSender(t, shanghaiState)
	shanghaiReceipt := executeSingleTx(t, shanghai, signCreationTx(t, shanghai, shanghaiKey, 0, initcode))
	require.Equal(t, types.ReceiptStatusSuccessful, shanghaiReceipt.Status,
		"PUSH0 must execute after Shanghai through the VinuChain processor")
}

func TestExecuteCancunTransientStorageThroughProcessor(t *testing.T) {
	code := []byte{
		byte(vm.PUSH1), 0x2a,
		byte(vm.PUSH1), 0x01,
		byte(vm.TSTORE),
		byte(vm.PUSH1), 0x01,
		byte(vm.TLOAD),
		byte(vm.PUSH0),
		byte(vm.MSTORE),
		byte(vm.STOP),
	}

	preCancun, preState := newStartedProcessor(t, idx.Block(0), forkTestRules(true, false), newStubChain())
	preKey, _ := fundedSender(t, preState)
	contract := common.HexToAddress("0x100")
	preState.SetCode(contract, code)
	preReceipt := executeSingleTx(t, preCancun, signCallTx(t, preCancun, preKey, 0, contract))
	require.Equal(t, types.ReceiptStatusFailed, preReceipt.Status,
		"TSTORE/TLOAD must fail before Cancun when executed through the VinuChain processor")

	cancun, cancunState := newStartedProcessor(t, idx.Block(0), forkTestRules(true, true), newStubChain())
	cancunKey, _ := fundedSender(t, cancunState)
	cancunState.SetCode(contract, code)
	cancunReceipt := executeSingleTx(t, cancun, signCallTx(t, cancun, cancunKey, 0, contract))
	require.Equal(t, types.ReceiptStatusSuccessful, cancunReceipt.Status,
		"TSTORE/TLOAD must execute after Cancun through the VinuChain processor")
}

func TestExecuteCancunMcopyThroughProcessor(t *testing.T) {
	value := []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
	}
	code := []byte{byte(vm.PUSH32)}
	code = append(code, value...)
	code = append(code,
		byte(vm.PUSH1), 0x20,
		byte(vm.MSTORE),
		byte(vm.PUSH1), 0x20,
		byte(vm.PUSH1), 0x20,
		byte(vm.PUSH0),
		byte(vm.MCOPY),
		byte(vm.STOP),
	)

	preCancun, preState := newStartedProcessor(t, idx.Block(0), forkTestRules(true, false), newStubChain())
	preKey, _ := fundedSender(t, preState)
	contract := common.HexToAddress("0x101")
	preState.SetCode(contract, code)
	preReceipt := executeSingleTx(t, preCancun, signCallTx(t, preCancun, preKey, 0, contract))
	require.Equal(t, types.ReceiptStatusFailed, preReceipt.Status,
		"MCOPY must fail before Cancun when executed through the VinuChain processor")

	cancun, cancunState := newStartedProcessor(t, idx.Block(0), forkTestRules(true, true), newStubChain())
	cancunKey, _ := fundedSender(t, cancunState)
	cancunState.SetCode(contract, code)
	cancunReceipt := executeSingleTx(t, cancun, signCallTx(t, cancun, cancunKey, 0, contract))
	require.Equal(t, types.ReceiptStatusSuccessful, cancunReceipt.Status,
		"MCOPY must execute after Cancun through the VinuChain processor")
}

func TestExecuteCancunSelfdestructThroughProcessor(t *testing.T) {
	cancun, sdb := newStartedProcessor(t, idx.Block(0), forkTestRules(true, true), newStubChain())
	key, _ := fundedSender(t, sdb)
	contract := common.HexToAddress("0x102")
	beneficiary := common.HexToAddress("0x202")
	code := append([]byte{byte(vm.PUSH20)}, beneficiary.Bytes()...)
	code = append(code, byte(vm.SELFDESTRUCT))
	sdb.SetCode(contract, code)
	sdb.SetBalance(contract, big.NewInt(100))

	receipt := executeSingleTx(t, cancun, signCallTx(t, cancun, key, 0, contract))
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	require.Equal(t, 0, sdb.GetBalance(contract).Sign(),
		"Cancun SELFDESTRUCT must transfer an existing contract balance")
	require.Equal(t, 0, sdb.GetBalance(beneficiary).Cmp(big.NewInt(100)),
		"Cancun SELFDESTRUCT must credit the beneficiary")
	require.NotZero(t, sdb.GetCodeSize(contract),
		"Cancun SELFDESTRUCT must preserve existing contract code")
}
