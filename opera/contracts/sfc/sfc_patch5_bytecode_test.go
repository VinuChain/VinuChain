package sfc

import (
	"bytes"
	"testing"
)

// TestValidatePatch5Bytecode_RejectsPlaceholder asserts the guard rejects a
// fresh 256-byte deadbeef-sentinel placeholder — the exact shape this
// scaffolding branch ships. Kept as a regression guard so a future Cycle-N
// scaffolding cycle still trips the validator.
func TestValidatePatch5Bytecode_RejectsPlaceholder(t *testing.T) {
	placeholder := bytes.Repeat([]byte{0xde, 0xad, 0xbe, 0xef}, 64)
	if err := validatePatch5Bytecode(placeholder); err == nil {
		t.Fatal("validatePatch5Bytecode accepted a deadbeef-sentinel placeholder — guard is broken")
	}
}

// TestValidatePatch5Bytecode_RejectsEmpty asserts the guard rejects nil/empty
// input. A future refactor that accidentally returns nil from
// GetPatch5ContractBin must still trigger log.Crit at re-flash time.
func TestValidatePatch5Bytecode_RejectsEmpty(t *testing.T) {
	if err := validatePatch5Bytecode(nil); err == nil {
		t.Fatal("validatePatch5Bytecode accepted nil input")
	}
	if err := validatePatch5Bytecode([]byte{}); err == nil {
		t.Fatal("validatePatch5Bytecode accepted empty input")
	}
}

// TestValidatePatch5Bytecode_RejectsTooShort asserts the minimum-length floor
// fires even when the sentinel prefix is absent.
func TestValidatePatch5Bytecode_RejectsTooShort(t *testing.T) {
	short := bytes.Repeat([]byte{0x60, 0x80, 0x60, 0x40}, 100) // 400 bytes
	if err := validatePatch5Bytecode(short); err == nil {
		t.Fatal("validatePatch5Bytecode accepted 400-byte input; minimum is 10000")
	}
}

// TestValidatePatch5Bytecode_RejectsAllZero asserts the all-zero paranoia
// guard catches a plausible miscompile output.
func TestValidatePatch5Bytecode_RejectsAllZero(t *testing.T) {
	zeros := make([]byte, minRealPatch5BytecodeLen+1000)
	if err := validatePatch5Bytecode(zeros); err == nil {
		t.Fatal("validatePatch5Bytecode accepted all-zero input")
	}
}

// TestValidatePatch5Bytecode_AcceptsPlausibleReal asserts the guard does not
// over-reject a fabricated bytecode that looks plausibly like a real SFC
// compile — starts with the Solidity dispatcher prefix, length above floor,
// contains non-zero mixed bytes.
func TestValidatePatch5Bytecode_AcceptsPlausibleReal(t *testing.T) {
	plausible := make([]byte, minRealPatch5BytecodeLen+5000)
	copy(plausible, []byte{0x60, 0x80, 0x60, 0x40, 0x52, 0x34, 0x80, 0x15, 0x61})
	for i := 9; i < len(plausible); i++ {
		plausible[i] = byte(i % 251)
	}
	if err := validatePatch5Bytecode(plausible); err != nil {
		t.Fatalf("validatePatch5Bytecode rejected plausible real bytecode: %v", err)
	}
}

// TestValidatePatch5Bytecode_RejectsPatch4 asserts that shipping the Patch4
// (Cycle-160) bytecode in the Patch5 slot is rejected. A byte-identical
// re-flash would be a no-op on chain and the pubkey-validation fix would
// not deploy.
func TestValidatePatch5Bytecode_RejectsPatch4(t *testing.T) {
	err := validatePatch5Bytecode(GetPatch4ContractBin())
	if err == nil {
		t.Fatal("validatePatch5Bytecode accepted Patch4 bytecode — a byte-identical re-flash would no-op on chain")
	}
	if err != errPatch5EqualsPatch4 {
		t.Fatalf("expected errPatch5EqualsPatch4, got: %v", err)
	}
}

// TestPatch5ContractBin_IsCurrentlyPlaceholder pins the v2.0.13-elemont
// scaffolding state: the compiled-in patch5ContractBin IS the deadbeef
// placeholder. When Cycle-161 truffle artefact lands and replaces this
// blob, this test will start failing — at which point delete this test and
// verify TestValidatePatch5Bytecode_RejectsPatch4 still asserts the guard
// against accidental Patch4=Patch5 shipping. Until then, this test prevents
// the v2.0.13 binary from accidentally shipping with a flipped activation
// flag (which would log.Crit at startup via EnforcePatch5StartupCheck) by
// keeping the placeholder shape verifiable.
func TestPatch5ContractBin_IsCurrentlyPlaceholder(t *testing.T) {
	if !bytes.HasPrefix(patch5ContractBin, patch5DeadbeefSentinel) {
		t.Fatal("patch5ContractBin no longer begins with deadbeef sentinel — when the real Cycle-161 blob lands, delete TestPatch5ContractBin_IsCurrentlyPlaceholder")
	}
}
