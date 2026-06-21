package sfc

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto"
)

func newCursorBackfillState(t *testing.T) *state.StateDB {
	t.Helper()
	statedb, err := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	if err != nil {
		t.Fatalf("failed to create state: %v", err)
	}
	return statedb
}

func setCursorSealedEpoch(statedb *state.StateDB, epoch uint64) {
	statedb.SetState(ContractAddress, cursorUint64Hash(sfcCurrentSealedEpochSlot), common.BigToHash(new(big.Int).SetUint64(epoch)))
}

// TestCursorSlotDerivationMatchesKnownKey pins the nested-mapping storage key
// for stashedRewardsUntilEpoch[delegator][validatorID] against an independently
// computed keccak. The base slot 114 and the key formula were verified on live
// testnet (NetworkID 206) before being trusted: for (0x8406413be0…,17) the
// stored cursor read 300 via both eth_call and eth_getStorageAt at this key.
func TestCursorSlotDerivationMatchesKnownKey(t *testing.T) {
	delegator := common.HexToAddress("0x8406413be017b3b7c72416a39cb8b3eeb7b9c2ef")
	validatorID := uint64(17)

	// Independent recomputation: keccak(validatorID || keccak(delegator || 114)).
	innerPre := make([]byte, 64)
	copy(innerPre[12:32], delegator.Bytes())
	copy(innerPre[32:64], common.BigToHash(new(big.Int).SetUint64(114)).Bytes())
	inner := crypto.Keccak256Hash(innerPre)
	outerPre := make([]byte, 64)
	copy(outerPre[0:32], common.BigToHash(new(big.Int).SetUint64(validatorID)).Bytes())
	copy(outerPre[32:64], inner.Bytes())
	want := crypto.Keccak256Hash(outerPre)

	if got := cursorSlot(delegator, validatorID); got != want {
		t.Fatalf("cursorSlot mismatch: got %s, want %s", got.Hex(), want.Hex())
	}

	// The base slot must be 114 (verified against the live chain).
	if sfcStashedRewardsUntilEpochSlot != 114 {
		t.Fatalf("stashedRewardsUntilEpoch base slot must be 114, got %d", sfcStashedRewardsUntilEpochSlot)
	}
	// currentSealedEpoch slot must be 103.
	if sfcCurrentSealedEpochSlot != 103 {
		t.Fatalf("currentSealedEpoch slot must be 103, got %d", sfcCurrentSealedEpochSlot)
	}
}

// TestCursorGetSetRoundtrip proves cursorGet/cursorSet read and write the same
// derived storage slot.
func TestCursorGetSetRoundtrip(t *testing.T) {
	statedb := newCursorBackfillState(t)
	d := common.HexToAddress("0x031a75844f399de13f70b9be4828cb040614d5be")
	v := uint64(15)
	if got := cursorGet(statedb, d, v); got.Sign() != 0 {
		t.Fatalf("expected zero cursor on fresh state, got %s", got)
	}
	cursorSet(statedb, d, v, big.NewInt(5643))
	if got := cursorGet(statedb, d, v); got.Uint64() != 5643 {
		t.Fatalf("cursorGet after set: got %s, want 5643", got)
	}
	// The raw storage slot must hold the value.
	if raw := statedb.GetState(ContractAddress, cursorSlot(d, v)).Big(); raw.Uint64() != 5643 {
		t.Fatalf("raw slot mismatch: got %s, want 5643", raw)
	}
}

// TestBackfillTestnetRewardCursors_RaisesStuckCursors verifies the happy path:
// every corrupted pair whose corrected epoch is above the stored value is raised
// to the corrected epoch.
func TestBackfillTestnetRewardCursors_RaisesStuckCursors(t *testing.T) {
	statedb := newCursorBackfillState(t)
	setCursorSealedEpoch(statedb, 6014)

	// Seed each pair's stored cursor below its corrected epoch.
	for _, p := range testnetCursorBackfill {
		cursorSet(statedb, p.delegator, p.validatorID, big.NewInt(int64(p.correctedEpoch-1)))
	}

	stats := BackfillTestnetRewardCursors(statedb)
	if stats.Raised != uint64(len(testnetCursorBackfill)) {
		t.Fatalf("expected %d raises, got %d", len(testnetCursorBackfill), stats.Raised)
	}
	for _, p := range testnetCursorBackfill {
		if got := cursorGet(statedb, p.delegator, p.validatorID).Uint64(); got != p.correctedEpoch {
			t.Fatalf("pair (%s,%d): got cursor %d, want %d", p.delegator.Hex(), p.validatorID, got, p.correctedEpoch)
		}
	}
}

// TestBackfillTestnetRewardCursors_NeverLowers proves the raise-only invariant:
// a stored cursor already at or above the corrected epoch is never lowered.
func TestBackfillTestnetRewardCursors_NeverLowers(t *testing.T) {
	statedb := newCursorBackfillState(t)
	setCursorSealedEpoch(statedb, 6014)

	// Set every cursor far above its corrected epoch.
	for _, p := range testnetCursorBackfill {
		cursorSet(statedb, p.delegator, p.validatorID, big.NewInt(6013))
	}

	stats := BackfillTestnetRewardCursors(statedb)
	if stats.Changed() {
		t.Fatalf("expected no raises when all cursors already above corrected, got %d", stats.Raised)
	}
	for _, p := range testnetCursorBackfill {
		if got := cursorGet(statedb, p.delegator, p.validatorID).Uint64(); got != 6013 {
			t.Fatalf("pair (%s,%d): cursor was modified to %d, must stay 6013", p.delegator.Hex(), p.validatorID, got)
		}
	}
}

// TestBackfillTestnetRewardCursors_EqualIsNoOp proves a cursor already exactly
// at the corrected epoch is a no-op (raise-only requires strict increase).
func TestBackfillTestnetRewardCursors_EqualIsNoOp(t *testing.T) {
	statedb := newCursorBackfillState(t)
	setCursorSealedEpoch(statedb, 6014)
	p := testnetCursorBackfill[0]
	cursorSet(statedb, p.delegator, p.validatorID, new(big.Int).SetUint64(p.correctedEpoch))
	// All others left at zero, so they will be raised; isolate the equal pair.
	stats := backfillRewardCursors(statedb, []cursorBackfillPair{p})
	if stats.Changed() {
		t.Fatalf("equal-cursor pair must be a no-op, got %d raises", stats.Raised)
	}
}

// TestBackfillTestnetRewardCursors_CapsAtSealedEpoch proves no cursor can ever be
// raised past currentSealedEpoch even if the corrected list says otherwise.
func TestBackfillTestnetRewardCursors_CapsAtSealedEpoch(t *testing.T) {
	statedb := newCursorBackfillState(t)
	const sealed = uint64(5800)
	setCursorSealedEpoch(statedb, sealed)

	// Use a single pair whose corrected epoch (6012) exceeds the sealed epoch.
	p := cursorPair("0x2ea34893c3c7513a9457f782ea35ac4390619992", 17, 6012)
	cursorSet(statedb, p.delegator, p.validatorID, big.NewInt(200))

	stats := backfillRewardCursors(statedb, []cursorBackfillPair{p})
	if stats.Raised != 1 {
		t.Fatalf("expected 1 raise, got %d", stats.Raised)
	}
	if got := cursorGet(statedb, p.delegator, p.validatorID).Uint64(); got != sealed {
		t.Fatalf("cursor must be capped at sealed epoch %d, got %d", sealed, got)
	}
}

// TestMainnetCursorBackfillEmpty pins the testnet-only gating at the data layer:
// the mainnet correction list is empty so the activation cannot mint phantom
// rewards on mainnet.
func TestMainnetCursorBackfillEmpty(t *testing.T) {
	if len(mainnetCursorBackfill) != 0 {
		t.Fatalf("mainnetCursorBackfill must be empty (mainnet ships the fix via GetLatestContractBin), got %d", len(mainnetCursorBackfill))
	}
	statedb := newCursorBackfillState(t)
	setCursorSealedEpoch(statedb, 6014)
	if stats := backfillRewardCursors(statedb, mainnetCursorBackfill); stats.Changed() {
		t.Fatalf("empty list must not write state, got %d raises", stats.Raised)
	}
}

// TestTestnetCursorBackfillListPinned pins the exact 12-pair correction list so a
// future edit that adds/removes a pair (or flips an excluded near-miss/val16
// pair into the list) is caught at build time. Validators 6/10/11/14 (near-miss)
// and 16 (phantom) must NOT appear.
func TestTestnetCursorBackfillListPinned(t *testing.T) {
	if got := len(testnetCursorBackfill); got != 12 {
		t.Fatalf("testnetCursorBackfill length: got %d, want 12", got)
	}
	forbiddenValidators := map[uint64]bool{6: true, 10: true, 11: true, 14: true, 16: true}
	allowedValidators := map[uint64]bool{15: true, 17: true, 18: true}
	seen := make(map[string]struct{}, len(testnetCursorBackfill))
	for _, p := range testnetCursorBackfill {
		if p.delegator == (common.Address{}) {
			t.Fatalf("zero delegator for validator %d", p.validatorID)
		}
		if p.correctedEpoch == 0 {
			t.Fatalf("zero corrected epoch for (%s,%d)", p.delegator.Hex(), p.validatorID)
		}
		if forbiddenValidators[p.validatorID] {
			t.Fatalf("excluded validator %d must not be in the backfill list (near-miss/val16)", p.validatorID)
		}
		if !allowedValidators[p.validatorID] {
			t.Fatalf("unexpected validator %d in backfill list", p.validatorID)
		}
		key := p.delegator.Hex() + ":" + new(big.Int).SetUint64(p.validatorID).String()
		if _, ok := seen[key]; ok {
			t.Fatalf("duplicate backfill pair %s", key)
		}
		seen[key] = struct{}{}
	}
}
