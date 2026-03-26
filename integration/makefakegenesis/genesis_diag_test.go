package makefakegenesis

import (
	"math/big"
	"testing"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"

	"github.com/Fantom-foundation/go-opera/inter/ibr"
	"github.com/Fantom-foundation/go-opera/inter/ier"
)

func TestFakeGenesisStoreDoesNotPanic(t *testing.T) {
	balance := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1000000))
	stake := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(100000))

	store := FakeGenesisStore(1, balance, stake)
	if store == nil {
		t.Fatal("FakeGenesisStore returned nil")
	}

	header := store.Header()
	if header.GenesisID == (hash.Hash{}) {
		t.Fatal("genesis hash is zero")
	}
	if header.NetworkID == 0 {
		t.Fatal("network ID is zero")
	}

	g := store.Genesis()

	var epochCount int
	g.Epochs.ForEach(func(er ier.LlrIdxFullEpochRecord) bool {
		epochCount++
		if er.EpochState.Epoch < 1 {
			t.Errorf("epoch state has invalid epoch number: %d", er.EpochState.Epoch)
		}
		if er.EpochState.Validators == nil {
			t.Error("epoch state has nil validators")
		}
		return true
	})
	if epochCount == 0 {
		t.Fatal("genesis has no epoch records")
	}

	var blockCount int
	g.Blocks.ForEach(func(_ ibr.LlrIdxFullBlockRecord) bool {
		blockCount++
		return true
	})
	if blockCount == 0 {
		t.Fatal("genesis has no block records")
	}
}

func TestFakeGenesisStoreMultiValidator(t *testing.T) {
	const numValidators = 3
	balance := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1000000))
	stake := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(100000))

	store := FakeGenesisStore(idx.Validator(numValidators), balance, stake)
	if store == nil {
		t.Fatal("FakeGenesisStore returned nil")
	}

	header := store.Header()
	if header.GenesisID == (hash.Hash{}) {
		t.Fatal("genesis hash is zero")
	}

	g := store.Genesis()

	var epochCount int
	g.Epochs.ForEach(func(er ier.LlrIdxFullEpochRecord) bool {
		epochCount++
		if er.EpochState.Epoch < 1 {
			t.Errorf("epoch state has invalid epoch number: %d", er.EpochState.Epoch)
		}
		if er.EpochState.Validators == nil {
			t.Error("epoch state has nil validators")
		}
		return true
	})
	if epochCount == 0 {
		t.Fatal("genesis has no epoch records")
	}

	var blockCount int
	g.Blocks.ForEach(func(_ ibr.LlrIdxFullBlockRecord) bool {
		blockCount++
		return true
	})
	if blockCount == 0 {
		t.Fatal("genesis has no block records")
	}
}
