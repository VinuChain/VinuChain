package makefakegenesis

import (
	"math/big"
	"testing"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"

	"github.com/Fantom-foundation/go-opera/inter/ibr"
	"github.com/Fantom-foundation/go-opera/inter/ier"
	"github.com/Fantom-foundation/go-opera/opera/genesisstore"
)

func assertGenesisStructure(t *testing.T, store *genesisstore.Store, expectedValidators int) {
	t.Helper()

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
		if er.EpochState.EpochStart == 0 {
			t.Error("epoch state has zero EpochStart timestamp")
		}
		if er.EpochState.Rules.NetworkID == 0 {
			t.Error("epoch state has zero NetworkID in rules")
		}
		if expectedValidators > 0 {
			profileCount := len(er.EpochState.ValidatorProfiles)
			if profileCount != expectedValidators {
				t.Errorf("expected %d validator profiles in epoch state, got %d", expectedValidators, profileCount)
			}
		}
		return true
	})
	if epochCount == 0 {
		t.Fatal("genesis has no epoch records")
	}

	var blockCount int
	g.Blocks.ForEach(func(br ibr.LlrIdxFullBlockRecord) bool {
		blockCount++
		if br.Idx == 0 {
			t.Error("block record has zero index")
		}
		if br.Time == 0 {
			t.Error("block record has zero timestamp")
		}
		return true
	})
	if blockCount == 0 {
		t.Fatal("genesis has no block records")
	}

	var evmItemCount int
	g.RawEvmItems.ForEach(func(key, value []byte) bool {
		evmItemCount++
		return true
	})
	if evmItemCount == 0 {
		t.Fatal("genesis has no EVM state items (expected deployed contract code including SFC)")
	}
}

func TestFakeGenesisStoreDoesNotPanic(t *testing.T) {
	balance := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1000000))
	stake := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(100000))

	store := FakeGenesisStore(1, balance, stake)
	assertGenesisStructure(t, store, 1)

	header := store.Header()
	if header.NetworkID == 0 {
		t.Fatal("network ID is zero")
	}
}

func TestFakeGenesisStoreMultiValidator(t *testing.T) {
	const numValidators = 3
	balance := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1000000))
	stake := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(100000))

	store := FakeGenesisStore(idx.Validator(numValidators), balance, stake)
	assertGenesisStructure(t, store, numValidators)
}
