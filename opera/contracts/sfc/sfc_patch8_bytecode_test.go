package sfc

import (
	"bytes"
	"testing"
)

func TestValidatePatch8Bytecode_RejectsPlaceholder(t *testing.T) {
	placeholder := append([]byte{}, patch8DeadbeefSentinel...)
	placeholder = append(placeholder, bytes.Repeat([]byte{0x01}, minRealPatch8BytecodeLen)...)
	if err := validatePatch8Bytecode(placeholder); err == nil {
		t.Fatal("validatePatch8Bytecode accepted a deadbeef-sentinel placeholder")
	}
}

func TestValidatePatch8Bytecode_RejectsEmpty(t *testing.T) {
	if err := validatePatch8Bytecode(nil); err == nil {
		t.Fatal("validatePatch8Bytecode accepted nil input")
	}
	if err := validatePatch8Bytecode([]byte{}); err == nil {
		t.Fatal("validatePatch8Bytecode accepted empty input")
	}
}

func TestValidatePatch8Bytecode_RejectsTooShort(t *testing.T) {
	short := bytes.Repeat([]byte{0x60}, 400)
	if err := validatePatch8Bytecode(short); err == nil {
		t.Fatal("validatePatch8Bytecode accepted 400-byte input")
	}
}

func TestValidatePatch8Bytecode_RejectsAllZero(t *testing.T) {
	zeros := make([]byte, minRealPatch8BytecodeLen+1000)
	if err := validatePatch8Bytecode(zeros); err == nil {
		t.Fatal("validatePatch8Bytecode accepted all-zero input")
	}
}

func TestValidatePatch8Bytecode_AcceptsPlausibleReal(t *testing.T) {
	plausible := make([]byte, minRealPatch8BytecodeLen+5000)
	for i := range plausible {
		plausible[i] = byte((i % 251) + 1)
	}
	if err := validatePatch8Bytecode(plausible); err != nil {
		t.Fatalf("validatePatch8Bytecode rejected plausible real bytecode: %v", err)
	}
}

func TestValidatePatch8Bytecode_RejectsPatch6(t *testing.T) {
	err := validatePatch8Bytecode(GetPatch6ContractBin())
	if err == nil {
		t.Fatal("validatePatch8Bytecode accepted Patch6 bytecode")
	}
	if err != errPatch8EqualsPatch6 {
		t.Fatalf("expected errPatch8EqualsPatch6, got: %v", err)
	}
}

func TestValidatePatch8Bytecode_RejectsPatch7(t *testing.T) {
	err := validatePatch8Bytecode(GetPatch7ContractBin())
	if err == nil {
		t.Fatal("validatePatch8Bytecode accepted Patch7 bytecode")
	}
	if err != errPatch8EqualsPatch7 {
		t.Fatalf("expected errPatch8EqualsPatch7, got: %v", err)
	}
}

// TestPatch8ContractBin_PassesEnforce asserts the shipped Patch8 asset is the
// real compiled Cycle-163 SFC runtime bytecode: it passes validatePatch8Bytecode
// (so EnforcePatch8StartupCheck lets the binary boot) and GetPatch8ContractBin
// returns it verbatim. The bytecode MUST still be reproduced byte-identical in the
// official release build and pinned in deployment-log.md before any live flash.
func TestPatch8ContractBin_PassesEnforce(t *testing.T) {
	if err := validatePatch8Bytecode(patch8ContractBin); err != nil {
		t.Fatalf("shipped Patch8 bytecode failed validation: %v", err)
	}
	if !bytes.Equal(GetPatch8ContractBin(), patch8ContractBin) {
		t.Fatal("GetPatch8ContractBin does not return patch8ContractBin verbatim")
	}
	// Sanity: a real compiled SFC runtime is ~48 KB.
	if len(patch8ContractBin) < 40000 {
		t.Fatalf("Patch8 bytecode unexpectedly small (%d bytes) — not a real SFC compile", len(patch8ContractBin))
	}
}

func TestPatch8DiffersFromPatch6AndPatch7(t *testing.T) {
	if bytes.Equal(patch8ContractBin, GetPatch6ContractBin()) {
		t.Fatal("Patch8 bytecode must differ from Patch6")
	}
	if bytes.Equal(patch8ContractBin, GetPatch7ContractBin()) {
		t.Fatal("Patch8 bytecode must differ from Patch7")
	}
}

func TestLatestContractBinMatchesPatch8(t *testing.T) {
	if !bytes.Equal(GetLatestContractBin(), GetPatch8ContractBin()) {
		t.Fatal("GetLatestContractBin must return the Patch8 reactivation bytecode for fresh SfcV2 activations")
	}
}
