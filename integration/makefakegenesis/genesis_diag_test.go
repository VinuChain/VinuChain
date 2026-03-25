package makefakegenesis

import (
	"math/big"
	"testing"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
)

func TestFakeGenesisStoreDoesNotPanic(t *testing.T) {
	balance := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1000000))
	stake := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(100000))

	store := FakeGenesisStore(1, balance, stake)
	if store == nil {
		t.Fatal("FakeGenesisStore returned nil")
	}
	t.Log("FakeGenesisStore(1 validator) succeeded")
}

func TestFakeGenesisStoreMultiValidator(t *testing.T) {
	balance := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1000000))
	stake := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(100000))

	store := FakeGenesisStore(idx.Validator(3), balance, stake)
	if store == nil {
		t.Fatal("FakeGenesisStore returned nil")
	}
	t.Log("FakeGenesisStore(3 validators) succeeded")
}
