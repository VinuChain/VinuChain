package sfc

import (
	"bytes"
	"testing"
)

func TestValidatePatch7Bytecode_RejectsPlaceholder(t *testing.T) {
	placeholder := append([]byte{}, patch7DeadbeefSentinel...)
	placeholder = append(placeholder, bytes.Repeat([]byte{0x01}, minRealPatch7BytecodeLen)...)
	if err := validatePatch7Bytecode(placeholder); err == nil {
		t.Fatal("validatePatch7Bytecode accepted a deadbeef-sentinel placeholder")
	}
}

func TestValidatePatch7Bytecode_RejectsEmpty(t *testing.T) {
	if err := validatePatch7Bytecode(nil); err == nil {
		t.Fatal("validatePatch7Bytecode accepted nil input")
	}
	if err := validatePatch7Bytecode([]byte{}); err == nil {
		t.Fatal("validatePatch7Bytecode accepted empty input")
	}
}

func TestValidatePatch7Bytecode_RejectsTooShort(t *testing.T) {
	short := bytes.Repeat([]byte{0x60}, 400)
	if err := validatePatch7Bytecode(short); err == nil {
		t.Fatal("validatePatch7Bytecode accepted 400-byte input")
	}
}

func TestValidatePatch7Bytecode_RejectsAllZero(t *testing.T) {
	zeros := make([]byte, minRealPatch7BytecodeLen+1000)
	if err := validatePatch7Bytecode(zeros); err == nil {
		t.Fatal("validatePatch7Bytecode accepted all-zero input")
	}
}

func TestValidatePatch7Bytecode_AcceptsPlausibleReal(t *testing.T) {
	plausible := make([]byte, minRealPatch7BytecodeLen+5000)
	for i := range plausible {
		plausible[i] = byte((i % 251) + 1)
	}
	if err := validatePatch7Bytecode(plausible); err != nil {
		t.Fatalf("validatePatch7Bytecode rejected plausible real bytecode: %v", err)
	}
}

func TestValidatePatch7Bytecode_RejectsPatch5(t *testing.T) {
	err := validatePatch7Bytecode(GetPatch5ContractBin())
	if err == nil {
		t.Fatal("validatePatch7Bytecode accepted Patch5 bytecode")
	}
	if err != errPatch7EqualsPatch5 {
		t.Fatalf("expected errPatch7EqualsPatch5, got: %v", err)
	}
}

func TestValidatePatch7Bytecode_RejectsPatch6(t *testing.T) {
	err := validatePatch7Bytecode(GetPatch6ContractBin())
	if err == nil {
		t.Fatal("validatePatch7Bytecode accepted Patch6 bytecode")
	}
	if err != errPatch7EqualsPatch6 {
		t.Fatalf("expected errPatch7EqualsPatch6, got: %v", err)
	}
}

func TestPatch7ContractBin_PassesEnforce(t *testing.T) {
	if err := validatePatch7Bytecode(patch7ContractBin); err != nil {
		t.Fatalf("validatePatch7Bytecode rejected the compiled-in Cycle-162 reward-cursor bytecode: %v", err)
	}
	if !bytes.Equal(GetPatch7ContractBin(), patch7ContractBin) {
		t.Fatal("GetPatch7ContractBin does not return patch7ContractBin verbatim")
	}
}

func TestPatch7DiffersFromPatch5AndPatch6(t *testing.T) {
	if bytes.Equal(patch7ContractBin, GetPatch5ContractBin()) {
		t.Fatal("Patch7 bytecode must differ from Patch5")
	}
	if bytes.Equal(patch7ContractBin, GetPatch6ContractBin()) {
		t.Fatal("Patch7 bytecode must differ from Patch6")
	}
}

// TestLatestContractBinSupersedesPatch7 documents that Patch7 is no longer the
// newest SFC bytecode: SfcV2Patch8 (Cycle-163 self-service-reactivation) superseded
// it, so GetLatestContractBin() must NOT return the Patch7 bytecode anymore. The
// positive assertion (latest == Patch8) lives in
// sfc_patch8_bytecode_test.go::TestLatestContractBinMatchesPatch8.
func TestLatestContractBinSupersedesPatch7(t *testing.T) {
	if bytes.Equal(GetLatestContractBin(), GetPatch7ContractBin()) {
		t.Fatal("GetLatestContractBin must no longer return the Patch7 bytecode — Patch8 (self-service reactivation) supersedes it for fresh SfcV2 activations")
	}
}

// TestEnforcePatch7StartupCheck_PassesWithRealBytecode proves the binary startup
// check does not log.Crit (panic) with the compiled-in Cycle-162 reward-cursor
// asset in place — i.e. this build is shippable.
func TestEnforcePatch7StartupCheck_PassesWithRealBytecode(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("EnforcePatch7StartupCheck panicked with the real bytecode in place: %v", r)
		}
	}()
	EnforcePatch7StartupCheck()
}
