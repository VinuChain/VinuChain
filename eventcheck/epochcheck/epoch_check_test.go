package epochcheck

import (
	"errors"
	"math"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"

	"github.com/Fantom-foundation/go-opera/eventcheck/basiccheck"
	"github.com/Fantom-foundation/go-opera/evmcore"
	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/opera"
)

func shanghaiCheckRules() opera.Rules {
	rules := opera.VinuChainTestNetRules()
	rules.Economy.MinGasPrice = big.NewInt(0)
	rules.Upgrades.Shanghai = true
	return rules
}

func TestCheckTxsShanghaiRejectsOversizedInitcode(t *testing.T) {
	rules := shanghaiCheckRules()
	tx := types.NewContractCreation(0, big.NewInt(0), 1_000_000, big.NewInt(1), make([]byte, params.MaxInitCodeSize+1))

	if err := CheckTxs(types.Transactions{tx}, rules); !errors.Is(err, evmcore.ErrMaxInitCodeSizeExceeded) {
		t.Fatalf("CheckTxs error = %v, want %v", err, evmcore.ErrMaxInitCodeSizeExceeded)
	}
}

func TestCheckTxsShanghaiMetersInitcodeGas(t *testing.T) {
	rules := shanghaiCheckRules()
	data := make([]byte, 33)
	preShanghaiGas, err := evmcore.IntrinsicGas(data, nil, true, true, true, false)
	if err != nil {
		t.Fatal(err)
	}
	tx := types.NewContractCreation(0, big.NewInt(0), preShanghaiGas, big.NewInt(1), data)

	if err := CheckTxs(types.Transactions{tx}, rules); !errors.Is(err, basiccheck.ErrIntrinsicGas) {
		t.Fatalf("CheckTxs error = %v, want %v", err, basiccheck.ErrIntrinsicGas)
	}

	rules.Upgrades.Shanghai = false
	if err := CheckTxs(types.Transactions{tx}, rules); err != nil {
		t.Fatalf("pre-Shanghai CheckTxs error = %v", err)
	}
}

func TestCalcGasPowerUsedSaturatesTxGasOverflow(t *testing.T) {
	rules := shanghaiCheckRules()
	e := &inter.MutableEventPayload{}
	e.SetTxs(types.Transactions{
		types.NewTransaction(0, common.Address{}, big.NewInt(0), math.MaxUint64-1, big.NewInt(0), nil),
		types.NewTransaction(1, common.Address{}, big.NewInt(0), 10, big.NewInt(0), nil),
	})

	if got := CalcGasPowerUsed(e, rules); got != math.MaxUint64 {
		t.Fatalf("CalcGasPowerUsed = %d, want saturated %d", got, uint64(math.MaxUint64))
	}
}
