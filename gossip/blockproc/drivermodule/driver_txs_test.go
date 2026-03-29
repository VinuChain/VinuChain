package drivermodule

import (
	"math/big"
	"testing"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/inter/pos"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-opera/inter/iblockproc"
	"github.com/Fantom-foundation/go-opera/opera"
)

func newTestStateDB() *state.StateDB {
	db := rawdb.NewMemoryDatabase()
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(db), nil)
	return statedb
}

func newTestListener(rules opera.Rules) (*DriverTxListener, *state.StateDB) {
	statedb := newTestStateDB()
	validators := pos.EqualWeightValidators([]idx.ValidatorID{1}, 100)
	bs := iblockproc.BlockState{
		ValidatorStates: []iblockproc.ValidatorBlockState{
			{Originated: new(big.Int)},
		},
	}
	es := iblockproc.EpochState{
		Validators: validators,
		Rules:      rules,
	}
	var minGasPrice *big.Int
	if rules.Economy.MinGasPrice != nil {
		minGasPrice = new(big.Int).Set(rules.Economy.MinGasPrice)
	}
	return &DriverTxListener{
		es:               es,
		bs:               bs,
		statedb:          statedb,
		blockMinGasPrice: minGasPrice,
	}, statedb
}

func legacyTx(gasPrice *big.Int, gasLimit uint64) *types.Transaction {
	return types.NewTransaction(0, common.HexToAddress("0x1"), common.Big0, gasLimit, gasPrice, nil)
}

func eip1559Tx(gasTipCap, gasFeeCap *big.Int, gasLimit uint64) *types.Transaction {
	to := common.HexToAddress("0x1")
	return types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(206),
		Nonce:     0,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &to,
		Value:     common.Big0,
	})
}

func receipt(gasUsed uint64, feeRefund *big.Int) *types.Receipt {
	return &types.Receipt{
		GasUsed:   gasUsed,
		FeeRefund: feeRefund,
	}
}

func originated(l *DriverTxListener) *big.Int {
	return l.bs.ValidatorStates[0].Originated
}

func TestOnNewReceipt_LegacyTx_NoBurn_NoSfcV2(t *testing.T) {
	require := require.New(t)
	rules := opera.Rules{
		Upgrades: opera.Upgrades{Berlin: true, London: true, Llr: true},
		Economy:  opera.EconomyRules{MinGasPrice: big.NewInt(1e9)},
	}
	l, _ := newTestListener(rules)

	gasPrice := big.NewInt(2e9)
	tx := legacyTx(gasPrice, 100000)
	r := receipt(21000, nil)

	l.OnNewReceipt(tx, r, 1)

	// Full fee goes to validator: 21000 * 2e9 = 42000e9
	expected := new(big.Int).Mul(big.NewInt(21000), gasPrice)
	require.Equal(expected, originated(l))
}

func TestOnNewReceipt_EIP1559_NoBurn_SfcV2Disabled(t *testing.T) {
	require := require.New(t)
	rules := opera.Rules{
		Upgrades: opera.Upgrades{Berlin: true, London: true, Llr: true},
		Economy:  opera.EconomyRules{MinGasPrice: big.NewInt(1e9)},
	}
	l, _ := newTestListener(rules)

	tx := eip1559Tx(big.NewInt(1e9), big.NewInt(2e9), 100000)
	r := receipt(21000, nil)

	l.OnNewReceipt(tx, r, 1)

	// effectiveGasPrice = min(1e9 + 1e9, 2e9) = 2e9
	// No burn (SfcV2 disabled), full fee to validator
	expected := new(big.Int).Mul(big.NewInt(21000), big.NewInt(2e9))
	require.Equal(expected, originated(l))
}

func TestOnNewReceipt_EIP1559_30PercentBurn_SfcV2Enabled(t *testing.T) {
	require := require.New(t)
	baseFee := big.NewInt(1e9)
	rules := opera.Rules{
		Upgrades: opera.Upgrades{Berlin: true, London: true, Llr: true, SfcV2: true},
		Economy:  opera.EconomyRules{MinGasPrice: baseFee},
	}
	l, statedb := newTestListener(rules)

	tx := eip1559Tx(big.NewInt(1e9), big.NewInt(2e9), 100000)
	r := receipt(21000, nil)

	l.OnNewReceipt(tx, r, 1)

	// effectiveGasPrice = min(1e9 + 1e9, 2e9) = 2e9
	// txFee = 21000 * 2e9 = 42000e9
	// baseFeeUsed = 21000 * 1e9 = 21000e9
	// burnAmount = 21000e9 * 30 / 100 = 6300e9
	// originated = 42000e9 - 6300e9 = 35700e9
	txFee := new(big.Int).Mul(big.NewInt(21000), big.NewInt(2e9))
	burnExpected := new(big.Int).Mul(big.NewInt(21000), baseFee)
	burnExpected.Mul(burnExpected, big.NewInt(30))
	burnExpected.Div(burnExpected, big.NewInt(100))
	expectedOriginated := new(big.Int).Sub(txFee, burnExpected)

	require.Equal(expectedOriginated, originated(l))

	// Burn sent to 0x0
	require.Equal(burnExpected, statedb.GetBalance(common.Address{}))
}

func TestOnNewReceipt_WithRefund_BurnFromValidatorShare(t *testing.T) {
	require := require.New(t)
	baseFee := big.NewInt(1e9)
	rules := opera.Rules{
		Upgrades: opera.Upgrades{Berlin: true, London: true, Llr: true, SfcV2: true},
		Economy:  opera.EconomyRules{MinGasPrice: baseFee},
	}
	l, statedb := newTestListener(rules)

	tx := eip1559Tx(big.NewInt(0), big.NewInt(1e9), 100000)
	refund := big.NewInt(10000e9) // partial refund
	r := receipt(21000, refund)

	l.OnNewReceipt(tx, r, 1)

	// effectiveGasPrice = min(0 + 1e9, 1e9) = 1e9
	// txFee = 21000 * 1e9 = 21000e9
	// validatorFee = 21000e9 - 10000e9 = 11000e9
	// baseFeeUsed = 21000 * 1e9 = 21000e9
	// burnAmount = 21000e9 * 30 / 100 = 6300e9
	// burn capped to validatorFee? 6300e9 < 11000e9 → no cap
	// originated = 11000e9 - 6300e9 = 4700e9
	txFee := new(big.Int).Mul(big.NewInt(21000), big.NewInt(1e9))
	validatorFee := new(big.Int).Sub(txFee, refund)
	burnExpected := new(big.Int).Mul(big.NewInt(21000), baseFee)
	burnExpected.Mul(burnExpected, big.NewInt(30))
	burnExpected.Div(burnExpected, big.NewInt(100))
	expectedOriginated := new(big.Int).Sub(validatorFee, burnExpected)

	require.Equal(expectedOriginated, originated(l))
	require.Equal(burnExpected, statedb.GetBalance(common.Address{}))
}

func TestOnNewReceipt_BurnExceedsValidatorFee_Capped(t *testing.T) {
	require := require.New(t)
	baseFee := big.NewInt(1e9)
	rules := opera.Rules{
		Upgrades: opera.Upgrades{Berlin: true, London: true, Llr: true, SfcV2: true},
		Economy:  opera.EconomyRules{MinGasPrice: baseFee},
	}
	l, statedb := newTestListener(rules)

	tx := eip1559Tx(big.NewInt(0), big.NewInt(1e9), 100000)
	// Refund almost the entire fee, leaving very little for validators
	txFee := new(big.Int).Mul(big.NewInt(21000), big.NewInt(1e9))
	refund := new(big.Int).Sub(txFee, big.NewInt(100)) // validatorFee = 100 wei
	r := receipt(21000, refund)

	l.OnNewReceipt(tx, r, 1)

	// validatorFee = 100 wei
	// burnAmount = 6300e9 (30% of baseFee portion) >> 100 → capped to 100
	// originated = 100 - 100 = 0
	require.Equal(new(big.Int), originated(l))

	// Burn is capped to validatorFee (100 wei)
	require.Equal(big.NewInt(100), statedb.GetBalance(common.Address{}))
}

func TestOnNewReceipt_NilFeeRefund_NoPanic(t *testing.T) {
	require := require.New(t)
	rules := opera.Rules{
		Upgrades: opera.Upgrades{Berlin: true, London: true, Llr: true, SfcV2: true},
		Economy:  opera.EconomyRules{MinGasPrice: big.NewInt(1e9)},
	}
	l, _ := newTestListener(rules)

	tx := legacyTx(big.NewInt(1e9), 100000)
	r := receipt(21000, nil) // nil FeeRefund

	// Should not panic
	require.NotPanics(func() {
		l.OnNewReceipt(tx, r, 1)
	})

	// Full fee minus burn
	txFee := new(big.Int).Mul(big.NewInt(21000), big.NewInt(1e9))
	burnExpected := new(big.Int).Mul(txFee, big.NewInt(30))
	burnExpected.Div(burnExpected, big.NewInt(100))
	expectedOriginated := new(big.Int).Sub(txFee, burnExpected)
	require.Equal(expectedOriginated, originated(l))
}

func TestOnNewReceipt_ZeroMinGasPrice_NoBurn(t *testing.T) {
	require := require.New(t)
	rules := opera.Rules{
		Upgrades: opera.Upgrades{Berlin: true, London: true, Llr: true, SfcV2: true},
		Economy:  opera.EconomyRules{MinGasPrice: big.NewInt(0)},
	}
	l, statedb := newTestListener(rules)

	tx := legacyTx(big.NewInt(1e9), 100000)
	r := receipt(21000, nil)

	l.OnNewReceipt(tx, r, 1)

	// MinGasPrice is 0, so burn guard (Sign() > 0) prevents burn
	expected := new(big.Int).Mul(big.NewInt(21000), big.NewInt(1e9))
	require.Equal(expected, originated(l))
	require.Equal(new(big.Int), statedb.GetBalance(common.Address{}))
}

func TestOnNewReceipt_RefundExceedsFee_OriginatedZero(t *testing.T) {
	require := require.New(t)
	rules := opera.Rules{
		Upgrades: opera.Upgrades{Berlin: true, London: true, Llr: true, SfcV2: true},
		Economy:  opera.EconomyRules{MinGasPrice: big.NewInt(1e9)},
	}
	l, statedb := newTestListener(rules)

	tx := legacyTx(big.NewInt(1e9), 100000)
	// Refund exceeds fee (edge case)
	hugeRefund := new(big.Int).Mul(big.NewInt(100000), big.NewInt(1e9))
	r := receipt(21000, hugeRefund)

	l.OnNewReceipt(tx, r, 1)

	// validatorFee clamps to 0, burn is 0, originated is 0
	require.Equal(new(big.Int), originated(l))
	require.Equal(new(big.Int), statedb.GetBalance(common.Address{}))
}

func TestOnNewReceipt_ZeroOriginator_Skipped(t *testing.T) {
	require := require.New(t)
	rules := opera.Rules{
		Upgrades: opera.Upgrades{Berlin: true, London: true, Llr: true, SfcV2: true},
		Economy:  opera.EconomyRules{MinGasPrice: big.NewInt(1e9)},
	}
	l, _ := newTestListener(rules)

	tx := legacyTx(big.NewInt(1e9), 100000)
	r := receipt(21000, nil)

	l.OnNewReceipt(tx, r, 0) // originator 0 = skip

	require.Equal(new(big.Int), originated(l))
}

func TestOnNewReceipt_NilMinGasPrice_NoBurn(t *testing.T) {
	require := require.New(t)
	rules := opera.Rules{
		Upgrades: opera.Upgrades{Berlin: true, London: true, Llr: true, SfcV2: true},
		Economy:  opera.EconomyRules{MinGasPrice: nil},
	}
	l, statedb := newTestListener(rules)

	tx := legacyTx(big.NewInt(1e9), 100000)
	r := receipt(21000, nil)

	require.NotPanics(func() {
		l.OnNewReceipt(tx, r, 1)
	})

	// Nil MinGasPrice: effectiveGasPrice falls back to tx.GasPrice(), no burn
	expected := new(big.Int).Mul(big.NewInt(21000), big.NewInt(1e9))
	require.Equal(expected, originated(l))
	require.Equal(new(big.Int), statedb.GetBalance(common.Address{}))
}

func TestOnNewReceipt_MultipleReceipts_Accumulate(t *testing.T) {
	require := require.New(t)
	baseFee := big.NewInt(1e9)
	rules := opera.Rules{
		Upgrades: opera.Upgrades{Berlin: true, London: true, Llr: true, SfcV2: true},
		Economy:  opera.EconomyRules{MinGasPrice: baseFee},
	}
	l, statedb := newTestListener(rules)

	tx1 := eip1559Tx(big.NewInt(1e9), big.NewInt(2e9), 100000)
	r1 := receipt(21000, nil)
	l.OnNewReceipt(tx1, r1, 1)

	tx2 := eip1559Tx(big.NewInt(1e9), big.NewInt(2e9), 100000)
	r2 := receipt(42000, nil)
	l.OnNewReceipt(tx2, r2, 1)

	// tx1: effectiveGasPrice=2e9, txFee=21000*2e9=42000e9, burn=21000*1e9*30/100=6300e9, originated=35700e9
	// tx2: effectiveGasPrice=2e9, txFee=42000*2e9=84000e9, burn=42000*1e9*30/100=12600e9, originated=71400e9
	// total originated = 35700e9 + 71400e9 = 107100e9
	// total burn = 6300e9 + 12600e9 = 18900e9
	expectedOriginated := big.NewInt(107100e9)
	expectedBurn := big.NewInt(18900e9)

	require.Equal(expectedOriginated, originated(l))
	require.Equal(expectedBurn, statedb.GetBalance(common.Address{}))
}

func TestOnNewReceipt_LondonFalse_SfcV2True_NoBurn(t *testing.T) {
	require := require.New(t)
	rules := opera.Rules{
		Upgrades: opera.Upgrades{Berlin: true, London: false, Llr: true, SfcV2: true},
		Economy:  opera.EconomyRules{MinGasPrice: big.NewInt(1e9)},
	}
	l, statedb := newTestListener(rules)

	tx := legacyTx(big.NewInt(2e9), 100000)
	r := receipt(21000, nil)

	l.OnNewReceipt(tx, r, 1)

	// SfcV2 is true but London is false: burn guard requires both, no burn
	expected := new(big.Int).Mul(big.NewInt(21000), big.NewInt(2e9))
	require.Equal(expected, originated(l))
	require.Equal(new(big.Int), statedb.GetBalance(common.Address{}))
}
