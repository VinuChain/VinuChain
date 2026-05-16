package sfc

import (
	"bytes"
	"testing"
)

func TestValidatePatch6Bytecode_RejectsPlaceholder(t *testing.T) {
	placeholder := append([]byte{}, patch6DeadbeefSentinel...)
	placeholder = append(placeholder, bytes.Repeat([]byte{0x01}, minRealPatch6BytecodeLen)...)
	if err := validatePatch6Bytecode(placeholder); err == nil {
		t.Fatal("validatePatch6Bytecode accepted a deadbeef-sentinel placeholder")
	}
}

func TestValidatePatch6Bytecode_RejectsEmpty(t *testing.T) {
	if err := validatePatch6Bytecode(nil); err == nil {
		t.Fatal("validatePatch6Bytecode accepted nil input")
	}
	if err := validatePatch6Bytecode([]byte{}); err == nil {
		t.Fatal("validatePatch6Bytecode accepted empty input")
	}
}

func TestValidatePatch6Bytecode_RejectsTooShort(t *testing.T) {
	short := bytes.Repeat([]byte{0x60}, 400)
	if err := validatePatch6Bytecode(short); err == nil {
		t.Fatal("validatePatch6Bytecode accepted 400-byte input")
	}
}

func TestValidatePatch6Bytecode_RejectsAllZero(t *testing.T) {
	zeros := make([]byte, minRealPatch6BytecodeLen+1000)
	if err := validatePatch6Bytecode(zeros); err == nil {
		t.Fatal("validatePatch6Bytecode accepted all-zero input")
	}
}

func TestValidatePatch6Bytecode_AcceptsPlausibleReal(t *testing.T) {
	plausible := make([]byte, minRealPatch6BytecodeLen+5000)
	for i := range plausible {
		plausible[i] = byte((i % 251) + 1)
	}
	if err := validatePatch6Bytecode(plausible); err != nil {
		t.Fatalf("validatePatch6Bytecode rejected plausible real bytecode: %v", err)
	}
}

func TestValidatePatch6Bytecode_RejectsPatch5(t *testing.T) {
	err := validatePatch6Bytecode(GetPatch5ContractBin())
	if err == nil {
		t.Fatal("validatePatch6Bytecode accepted Patch5 bytecode")
	}
	if err != errPatch6EqualsPatch5 {
		t.Fatalf("expected errPatch6EqualsPatch5, got: %v", err)
	}
}

func TestPatch6ContractBin_PassesEnforce(t *testing.T) {
	if err := validatePatch6Bytecode(patch6ContractBin); err != nil {
		t.Fatalf("validatePatch6Bytecode rejected the compiled-in Cycle-162 bytecode: %v", err)
	}
	if !bytes.Equal(GetPatch6ContractBin(), patch6ContractBin) {
		t.Fatal("GetPatch6ContractBin does not return patch6ContractBin verbatim")
	}
}

func TestLatestContractBinMatchesPatch6(t *testing.T) {
	if !bytes.Equal(GetLatestContractBin(), GetPatch6ContractBin()) {
		t.Fatal("GetLatestContractBin must return the Patch6 bytecode for fresh SfcV2 activations")
	}
}
