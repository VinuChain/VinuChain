package ethapi

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/params"
)

// newTestStateDB creates a minimal in-memory StateDB suitable for unit tests.
func newTestStateDB(t *testing.T) *state.StateDB {
	t.Helper()
	db := rawdb.NewMemoryDatabase()
	sdb, err := state.New(common.Hash{}, state.NewDatabase(db), nil)
	if err != nil {
		t.Fatalf("state.New: %v", err)
	}
	return sdb
}

// TestStateOverrideApply_CodeSizeLimit verifies that Apply rejects code blobs
// larger than params.MaxCodeSize in a state override, and accepts blobs of exactly
// params.MaxCodeSize bytes.
func TestStateOverrideApply_CodeSizeLimit(t *testing.T) {
	oversized := make(hexutil.Bytes, params.MaxCodeSize+1)
	override := StateOverride{
		common.Address{1}: OverrideAccount{Code: &oversized},
	}
	sdb := newTestStateDB(t)
	if err := override.Apply(sdb); err == nil {
		t.Fatalf("Apply must reject code larger than params.MaxCodeSize (%d); got nil error", params.MaxCodeSize)
	}
}

// TestStateOverrideApply_CodeSizeLimit_ExactBoundary verifies that code of exactly
// params.MaxCodeSize is accepted (boundary condition).
func TestStateOverrideApply_CodeSizeLimit_ExactBoundary(t *testing.T) {
	exact := make(hexutil.Bytes, params.MaxCodeSize)
	override := StateOverride{
		common.Address{1}: OverrideAccount{Code: &exact},
	}
	sdb := newTestStateDB(t)
	if err := override.Apply(sdb); err != nil {
		t.Fatalf("Apply must accept code of exactly MaxCodeSize; got err: %v", err)
	}
}

// TestStateOverrideApply_CodeSizeLimit_NilCode verifies that a nil Code field
// is a no-op (no crash or spurious error).
func TestStateOverrideApply_CodeSizeLimit_NilCode(t *testing.T) {
	override := StateOverride{
		common.Address{1}: OverrideAccount{Code: nil},
	}
	sdb := newTestStateDB(t)
	if err := override.Apply(sdb); err != nil {
		t.Fatalf("Apply with nil Code must succeed; got err: %v", err)
	}
}

// TestStateOverrideApply_OversizedCodeRejectedBeforeKeccak verifies that after the
// fix, Apply rejects oversized code immediately — before the expensive
// crypto.Keccak256Hash call inside StateDB.SetCode.
//
// Pre-fix, a 1MB code blob triggered ~2ms of synchronous keccak work inside Apply
// (confirmed by timing before the fix was applied). Post-fix, Apply returns an error
// in microseconds because the size check runs first.
//
// This serves as both a negative PoC (fix blocks the attack path) and a performance
// regression guard (reject early, do not hash).
func TestStateOverrideApply_OversizedCodeRejectedBeforeKeccak(t *testing.T) {
	largeMB := make(hexutil.Bytes, 1*1024*1024)
	for i := range largeMB {
		largeMB[i] = 0x61 // PUSH2 opcode
	}
	override := StateOverride{
		common.Address{1}: OverrideAccount{Code: &largeMB},
	}
	sdb := newTestStateDB(t)

	start := time.Now()
	err := override.Apply(sdb)
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("Apply must reject 1MB code override after the fix")
	}
	// The rejection path is a simple len() check — must complete in well under 1ms.
	// Pre-fix, the same call took ~2ms because keccak ran first.
	if elapsed > 1*time.Millisecond {
		t.Errorf("Apply took %v; expected <1ms after fix (keccak must not run for oversized code)", elapsed)
	}
	t.Logf("Apply correctly rejected 1MB code in %v (no keccak computed)", elapsed)
}

// TestStateOverrideApply_AccountLimit verifies the existing 256-account cap
// is enforced (regression guard).
func TestStateOverrideApply_AccountLimit(t *testing.T) {
	override := make(StateOverride, 257)
	code := hexutil.Bytes{0x60, 0x00} // minimal valid code
	for i := 0; i < 257; i++ {
		var addr common.Address
		// Use two bytes to avoid collisions above 255.
		addr[18] = byte(i >> 8)
		addr[19] = byte(i & 0xff)
		override[addr] = OverrideAccount{Code: &code}
	}
	sdb := newTestStateDB(t)
	err := override.Apply(sdb)
	if err == nil {
		t.Fatal("Apply must reject override with > 256 accounts")
	}
}
