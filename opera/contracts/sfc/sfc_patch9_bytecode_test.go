package sfc

import (
	"bytes"
	"testing"
)

func TestValidatePatch9Bytecode_RejectsPlaceholder(t *testing.T) {
	placeholder := append([]byte{}, patch9DeadbeefSentinel...)
	placeholder = append(placeholder, bytes.Repeat([]byte{0x01}, minRealPatch9BytecodeLen)...)
	if err := validatePatch9Bytecode(placeholder); err == nil {
		t.Fatal("validatePatch9Bytecode accepted a deadbeef-sentinel placeholder")
	}
}

func TestValidatePatch9Bytecode_RejectsEmpty(t *testing.T) {
	if err := validatePatch9Bytecode(nil); err == nil {
		t.Fatal("validatePatch9Bytecode accepted nil input")
	}
	if err := validatePatch9Bytecode([]byte{}); err == nil {
		t.Fatal("validatePatch9Bytecode accepted empty input")
	}
}

func TestValidatePatch9Bytecode_RejectsTooShort(t *testing.T) {
	short := bytes.Repeat([]byte{0x60}, 400)
	if err := validatePatch9Bytecode(short); err == nil {
		t.Fatal("validatePatch9Bytecode accepted 400-byte input")
	}
}

func TestValidatePatch9Bytecode_RejectsAllZero(t *testing.T) {
	zeros := make([]byte, minRealPatch9BytecodeLen+1000)
	if err := validatePatch9Bytecode(zeros); err == nil {
		t.Fatal("validatePatch9Bytecode accepted all-zero input")
	}
}

func TestValidatePatch9Bytecode_AcceptsPlausibleReal(t *testing.T) {
	plausible := make([]byte, minRealPatch9BytecodeLen+5000)
	for i := range plausible {
		plausible[i] = byte((i % 251) + 1)
	}
	if err := validatePatch9Bytecode(plausible); err != nil {
		t.Fatalf("validatePatch9Bytecode rejected plausible real bytecode: %v", err)
	}
}

func TestValidatePatch9Bytecode_RejectsPatch7(t *testing.T) {
	err := validatePatch9Bytecode(GetPatch7ContractBin())
	if err == nil {
		t.Fatal("validatePatch9Bytecode accepted Patch7 bytecode")
	}
	if err != errPatch9EqualsPatch7 {
		t.Fatalf("expected errPatch9EqualsPatch7, got: %v", err)
	}
}

func TestValidatePatch9Bytecode_RejectsPatch8(t *testing.T) {
	err := validatePatch9Bytecode(GetPatch8ContractBin())
	if err == nil {
		t.Fatal("validatePatch9Bytecode accepted Patch8 bytecode")
	}
	if err != errPatch9EqualsPatch8 {
		t.Fatalf("expected errPatch9EqualsPatch8, got: %v", err)
	}
}

// TestPatch9ContractBin_PassesEnforce asserts the shipped Patch9 asset is the
// real compiled Cycle-164 SFC runtime bytecode: it passes validatePatch9Bytecode
// (so EnforcePatch9StartupCheck lets the binary boot) and GetPatch9ContractBin
// returns it verbatim. The bytecode MUST still be reproduced byte-identical in the
// official release build and pinned in deployment-log.md before any live flash.
func TestPatch9ContractBin_PassesEnforce(t *testing.T) {
	if err := validatePatch9Bytecode(patch9ContractBin); err != nil {
		t.Fatalf("shipped Patch9 bytecode failed validation: %v", err)
	}
	if !bytes.Equal(GetPatch9ContractBin(), patch9ContractBin) {
		t.Fatal("GetPatch9ContractBin does not return patch9ContractBin verbatim")
	}
	// Sanity: a real compiled SFC runtime is ~48 KB.
	if len(patch9ContractBin) < 40000 {
		t.Fatalf("Patch9 bytecode unexpectedly small (%d bytes) — not a real SFC compile", len(patch9ContractBin))
	}
}

func TestPatch9DiffersFromPatch7AndPatch8(t *testing.T) {
	if bytes.Equal(patch9ContractBin, GetPatch7ContractBin()) {
		t.Fatal("Patch9 bytecode must differ from Patch7")
	}
	if bytes.Equal(patch9ContractBin, GetPatch8ContractBin()) {
		t.Fatal("Patch9 bytecode must differ from Patch8")
	}
}

func TestLatestContractBinMatchesPatch9(t *testing.T) {
	if !bytes.Equal(GetLatestContractBin(), GetPatch9ContractBin()) {
		t.Fatal("GetLatestContractBin must return the Patch9 two-reward-fix bytecode for fresh SfcV2 activations")
	}
}
