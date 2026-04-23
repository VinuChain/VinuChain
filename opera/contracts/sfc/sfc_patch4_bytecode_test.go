package sfc

import (
	"bytes"
	"testing"
)

// TestValidatePatch4Bytecode_RejectsPlaceholder asserts the guard rejects the
// scaffolding sentinel bytes so a release cannot ship them by accident.
func TestValidatePatch4Bytecode_RejectsPlaceholder(t *testing.T) {
	if err := validatePatch4Bytecode(patch4PlaceholderBytecode); err == nil {
		t.Fatal("validatePatch4Bytecode accepted the placeholder sentinel — guard is broken")
	}
}

// TestValidatePatch4Bytecode_RejectsEmpty asserts the guard rejects nil/empty
// input. A future refactor that accidentally returns nil from
// GetPatch4ContractBin must still trigger log.Crit at re-flash time.
func TestValidatePatch4Bytecode_RejectsEmpty(t *testing.T) {
	if err := validatePatch4Bytecode(nil); err == nil {
		t.Fatal("validatePatch4Bytecode accepted nil input")
	}
	if err := validatePatch4Bytecode([]byte{}); err == nil {
		t.Fatal("validatePatch4Bytecode accepted empty input")
	}
}

// TestValidatePatch4Bytecode_RejectsTooShort asserts the minimum-length floor
// fires even when the sentinel prefix is absent.
func TestValidatePatch4Bytecode_RejectsTooShort(t *testing.T) {
	// Non-sentinel bytes but well below minRealPatch4BytecodeLen.
	short := bytes.Repeat([]byte{0x60, 0x80, 0x60, 0x40}, 100) // 400 bytes
	if err := validatePatch4Bytecode(short); err == nil {
		t.Fatal("validatePatch4Bytecode accepted 400-byte input; minimum is 10000")
	}
}

// TestValidatePatch4Bytecode_RejectsAllZero asserts the all-zero paranoia
// guard catches a plausible miscompile output.
func TestValidatePatch4Bytecode_RejectsAllZero(t *testing.T) {
	zeros := make([]byte, minRealPatch4BytecodeLen+1000)
	if err := validatePatch4Bytecode(zeros); err == nil {
		t.Fatal("validatePatch4Bytecode accepted all-zero input")
	}
}

// TestValidatePatch4Bytecode_AcceptsPlausibleReal asserts the guard does not
// over-reject a fabricated bytecode that looks plausibly like a real SFC
// compile — starts with the Solidity dispatcher prefix, length above floor,
// contains non-zero mixed bytes.
func TestValidatePatch4Bytecode_AcceptsPlausibleReal(t *testing.T) {
	plausible := make([]byte, minRealPatch4BytecodeLen+5000)
	// Solidity dispatcher prefix `0x60806040...`
	copy(plausible, []byte{0x60, 0x80, 0x60, 0x40, 0x52, 0x34, 0x80, 0x15, 0x61})
	// Fill the rest with varied bytes so the all-zero guard is not tripped.
	for i := 9; i < len(plausible); i++ {
		plausible[i] = byte(i % 251)
	}
	if err := validatePatch4Bytecode(plausible); err != nil {
		t.Fatalf("validatePatch4Bytecode rejected plausible real bytecode: %v", err)
	}
}

// TestValidatePatch4Bytecode_RejectsPatch3 asserts that shipping the Patch3
// (Cycle-159) bytecode in the Patch4 slot is rejected. A byte-identical
// re-flash would be a no-op on chain and the lock-end-time fix would not
// deploy.
func TestValidatePatch4Bytecode_RejectsPatch3(t *testing.T) {
	err := validatePatch4Bytecode(GetContractBin())
	if err == nil {
		t.Fatal("validatePatch4Bytecode accepted Patch3 bytecode — a byte-identical re-flash would no-op on chain")
	}
	if err != errPatch4EqualsPatch3 {
		t.Fatalf("expected errPatch4EqualsPatch3, got: %v", err)
	}
}
