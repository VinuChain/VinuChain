package evmcore

import (
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
)

// minimalDummyChain implements DummyChain for state transition tests.
type minimalDummyChain struct{}

func (minimalDummyChain) GetHeader(_ common.Hash, _ uint64) *EvmHeader { return nil }

func shanghaiTestChainConfig() *params.ChainConfig {
	cfg := *params.TestChainConfig
	cfg.HomesteadBlock = common.Big0
	cfg.IstanbulBlock = common.Big0
	cfg.BerlinBlock = common.Big0
	cfg.LondonBlock = common.Big0
	cfg.ShanghaiBlock = common.Big0
	cfg.CancunBlock = nil
	return &cfg
}

func TestIntrinsicGasEIP3860AddsInitcodeWordCost(t *testing.T) {
	data := make([]byte, 33)

	preShanghai, err := IntrinsicGas(data, nil, true, true, true, false)
	if err != nil {
		t.Fatal(err)
	}
	shanghai, err := IntrinsicGas(data, nil, true, true, true, true)
	if err != nil {
		t.Fatal(err)
	}

	want := preShanghai + 2*params.InitCodeWordGas
	if shanghai != want {
		t.Fatalf("Shanghai intrinsic gas = %d, want %d", shanghai, want)
	}
}

func TestTransitionDbShanghaiRejectsOversizedInitcode(t *testing.T) {
	const gasLimit = uint64(1_000_000)
	sender := common.HexToAddress("0x1111")
	startBalance := big.NewInt(1_000_000_000)

	statedb, err := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	if err != nil {
		t.Fatal(err)
	}
	statedb.SetBalance(sender, startBalance)

	header := &EvmHeader{
		Number:   big.NewInt(1),
		GasLimit: gasLimit,
		BaseFee:  big.NewInt(0),
	}
	evm := vm.NewEVM(NewEVMBlockContext(header, minimalDummyChain{}, nil), vm.TxContext{}, statedb, shanghaiTestChainConfig(), vm.Config{})
	msg := types.NewMessage(
		sender,
		nil,
		0,
		big.NewInt(0),
		gasLimit,
		big.NewInt(1),
		nil,
		nil,
		make([]byte, params.MaxInitCodeSize+1),
		nil,
		true,
	)

	gp := new(GasPool).AddGas(gasLimit)
	_, err = ApplyMessage(evm, msg, gp, nil)
	if !errors.Is(err, ErrMaxInitCodeSizeExceeded) {
		t.Fatalf("ApplyMessage error = %v, want %v", err, ErrMaxInitCodeSizeExceeded)
	}
	if got := statedb.GetBalance(sender); got.Cmp(startBalance) != 0 {
		t.Fatalf("sender balance mutated on skipped oversized initcode: got %v, want %v", got, startBalance)
	}
	if got := uint64(*gp); got != gasLimit {
		t.Fatalf("gas pool mutated on skipped oversized initcode: got %d, want %d", got, gasLimit)
	}
}

func TestTransitionDbShanghaiIntrinsicGasErrorDoesNotDebit(t *testing.T) {
	data := make([]byte, 33)
	gasLimit, err := IntrinsicGas(data, nil, true, true, true, false)
	if err != nil {
		t.Fatal(err)
	}
	sender := common.HexToAddress("0x1111")
	startBalance := big.NewInt(1_000_000_000)

	statedb, err := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	if err != nil {
		t.Fatal(err)
	}
	statedb.SetBalance(sender, startBalance)

	header := &EvmHeader{
		Number:   big.NewInt(1),
		GasLimit: gasLimit,
		BaseFee:  big.NewInt(0),
	}
	evm := vm.NewEVM(NewEVMBlockContext(header, minimalDummyChain{}, nil), vm.TxContext{}, statedb, shanghaiTestChainConfig(), vm.Config{})
	msg := types.NewMessage(
		sender,
		nil,
		0,
		big.NewInt(0),
		gasLimit,
		big.NewInt(1),
		nil,
		nil,
		data,
		nil,
		true,
	)

	gp := new(GasPool).AddGas(gasLimit)
	_, err = ApplyMessage(evm, msg, gp, nil)
	if !errors.Is(err, ErrIntrinsicGas) {
		t.Fatalf("ApplyMessage error = %v, want %v", err, ErrIntrinsicGas)
	}
	if got := statedb.GetBalance(sender); got.Cmp(startBalance) != 0 {
		t.Fatalf("sender balance mutated on skipped intrinsic gas error: got %v, want %v", got, startBalance)
	}
	if got := uint64(*gp); got != gasLimit {
		t.Fatalf("gas pool mutated on skipped intrinsic gas error: got %d, want %d", got, gasLimit)
	}
}

func TestTransitionDbShanghaiWarmsCoinbase(t *testing.T) {
	const gasLimit = uint64(100_000)
	sender := common.HexToAddress("0x1111")
	receiver := common.HexToAddress("0x2222")
	coinbase := common.HexToAddress("0xc0de")

	statedb, err := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	if err != nil {
		t.Fatal(err)
	}
	statedb.SetBalance(sender, big.NewInt(1_000_000_000))

	header := &EvmHeader{
		Number:   big.NewInt(1),
		GasLimit: gasLimit,
		BaseFee:  big.NewInt(0),
		Coinbase: coinbase,
	}
	evm := vm.NewEVM(NewEVMBlockContext(header, minimalDummyChain{}, nil), vm.TxContext{}, statedb, shanghaiTestChainConfig(), vm.Config{})
	msg := types.NewMessage(
		sender,
		&receiver,
		0,
		big.NewInt(0),
		gasLimit,
		big.NewInt(1),
		nil,
		nil,
		nil,
		nil,
		true,
	)

	if _, err := ApplyMessage(evm, msg, new(GasPool).AddGas(gasLimit), nil); err != nil {
		t.Fatalf("ApplyMessage: %v", err)
	}
	if !statedb.AddressInAccessList(coinbase) {
		t.Fatalf("coinbase %s was not warmed under Shanghai rules", coinbase)
	}
}

// TestFeeRefundBoundaries verifies that refundGas computes
// feeRefund = min(fee, availableQuota) across all boundary cases.
//
// A simple ETH transfer costs 21000 gas; with gasPrice=2 the fee is 42000 wei.
func TestFeeRefundBoundaries(t *testing.T) {
	const (
		gasLimit = uint64(21000)
		gasPrice = int64(2)
		fee      = int64(gasLimit) * gasPrice // 42000 wei
	)

	sender := common.HexToAddress("0x1111")
	receiver := common.HexToAddress("0x2222")

	cases := []struct {
		name           string
		availableQuota *big.Int
		wantRefund     int64
	}{
		{"nil quota (pre-Podgorica)", nil, 0},
		{"zero quota", big.NewInt(0), 0},
		{"quota < fee (partial refund)", big.NewInt(10000), 10000},
		{"quota == fee (exact refund)", big.NewInt(fee), fee},
		{"quota > fee (full refund)", big.NewInt(fee * 2), fee},
	}

	cfg := params.TestChainConfig

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			db := rawdb.NewMemoryDatabase()
			statedb, err := state.New(common.Hash{}, state.NewDatabase(db), nil)
			if err != nil {
				t.Fatalf("state.New: %v", err)
			}
			// Fund sender: enough for gas cost (no value transfer in this test).
			statedb.SetBalance(sender, big.NewInt(1e18))

			header := &EvmHeader{
				Number:   big.NewInt(1),
				GasLimit: gasLimit * 2,
				BaseFee:  big.NewInt(0),
			}
			blockCtx := NewEVMBlockContext(header, minimalDummyChain{}, nil)
			evm := vm.NewEVM(blockCtx, vm.TxContext{}, statedb, cfg, vm.Config{})

			msg := types.NewMessage(
				sender,
				&receiver,
				0,             // nonce (isFake skips nonce check)
				big.NewInt(0), // value
				gasLimit,
				big.NewInt(gasPrice),
				nil,  // gasFeeCap (legacy tx)
				nil,  // gasTipCap (legacy tx)
				nil,  // data
				nil,  // accessList
				true, // isFake: skip nonce/signature validation
			)

			gp := new(GasPool).AddGas(gasLimit)
			// The test header has no BaseFeeFloor set, so the context's BaseFeeFloor
			// is nil and the congestion guard is inert. This keeps the test focused
			// on refund-boundary logic.
			result, err := ApplyMessage(evm, msg, gp, tc.availableQuota)
			if err != nil {
				t.Fatalf("ApplyMessage: %v", err)
			}

			got := result.FeeRefund
			if got == nil {
				got = big.NewInt(0)
			}
			want := big.NewInt(tc.wantRefund)
			if got.Cmp(want) != 0 {
				t.Errorf("FeeRefund = %s, want %s", got, want)
			}
		})
	}
}
