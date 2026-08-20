package sfc

import (
	"bytes"
	"testing"
)

func TestValidatePatch10Bytecode_RejectsPlaceholder(t *testing.T) {
	placeholder := append([]byte{}, patch10DeadbeefSentinel...)
	placeholder = append(placeholder, bytes.Repeat([]byte{0x01}, minRealPatch10BytecodeLen)...)
	if err := validatePatch10Bytecode(placeholder); err == nil {
		t.Fatal("validatePatch10Bytecode accepted a deadbeef-sentinel placeholder")
	}
}

func TestValidatePatch10Bytecode_RejectsEmpty(t *testing.T) {
	if err := validatePatch10Bytecode(nil); err == nil {
		t.Fatal("validatePatch10Bytecode accepted nil input")
	}
	if err := validatePatch10Bytecode([]byte{}); err == nil {
		t.Fatal("validatePatch10Bytecode accepted empty input")
	}
}

func TestValidatePatch10Bytecode_RejectsTooShort(t *testing.T) {
	short := bytes.Repeat([]byte{0x60}, 400)
	if err := validatePatch10Bytecode(short); err == nil {
		t.Fatal("validatePatch10Bytecode accepted 400-byte input")
	}
}

func TestValidatePatch10Bytecode_RejectsAllZero(t *testing.T) {
	zeros := make([]byte, minRealPatch10BytecodeLen+1000)
	if err := validatePatch10Bytecode(zeros); err == nil {
		t.Fatal("validatePatch10Bytecode accepted all-zero input")
	}
}

// Unlike the prior patches' validators (plausibility-only), Patch10 pins the
// exact compile digest, so a plausible-but-wrong blob must be REJECTED — the
// wrong-blob failure mode becomes a refused boot instead of a silent
// consensus divergence.
func TestValidatePatch10Bytecode_RejectsPlausibleButWrong(t *testing.T) {
	plausible := make([]byte, minRealPatch10BytecodeLen+5000)
	for i := range plausible {
		plausible[i] = byte((i % 251) + 1)
	}
	err := validatePatch10Bytecode(plausible)
	if err == nil {
		t.Fatal("validatePatch10Bytecode accepted a plausible-but-wrong blob; the sha256 pin must reject it")
	}
	if err != errPatch10WrongDigest {
		t.Fatalf("expected errPatch10WrongDigest, got: %v", err)
	}
}

func TestValidatePatch10Bytecode_RejectsPatch8(t *testing.T) {
	err := validatePatch10Bytecode(GetPatch8ContractBin())
	if err == nil {
		t.Fatal("validatePatch10Bytecode accepted Patch8 bytecode")
	}
	if err != errPatch10EqualsPatch8 {
		t.Fatalf("expected errPatch10EqualsPatch8, got: %v", err)
	}
}

func TestValidatePatch10Bytecode_RejectsPatch9(t *testing.T) {
	err := validatePatch10Bytecode(GetPatch9ContractBin())
	if err == nil {
		t.Fatal("validatePatch10Bytecode accepted Patch9 bytecode")
	}
	if err != errPatch10EqualsPatch9 {
		t.Fatalf("expected errPatch10EqualsPatch9, got: %v", err)
	}
}

// TestPatch10ContractBin_PassesEnforce asserts the shipped Patch10 asset is the
// real compiled Cycle-165 SFC runtime bytecode: it passes
// validatePatch10Bytecode (so EnforcePatch10StartupCheck lets the binary boot)
// and GetPatch10ContractBin returns it verbatim. The bytecode MUST still be
// reproduced byte-identical in the official release build and pinned in
// deployment-log.md before any live flash.
func TestPatch10ContractBin_PassesEnforce(t *testing.T) {
	if err := validatePatch10Bytecode(patch10ContractBin); err != nil {
		t.Fatalf("shipped Patch10 bytecode failed validation: %v", err)
	}
	if !bytes.Equal(GetPatch10ContractBin(), patch10ContractBin) {
		t.Fatal("GetPatch10ContractBin does not return patch10ContractBin verbatim")
	}
	if len(patch10ContractBin) < 40000 {
		t.Fatalf("Patch10 bytecode unexpectedly small (%d bytes) — not a real SFC compile", len(patch10ContractBin))
	}
}

func TestPatch10DiffersFromPatch8AndPatch9(t *testing.T) {
	if bytes.Equal(patch10ContractBin, GetPatch8ContractBin()) {
		t.Fatal("Patch10 bytecode must differ from Patch8")
	}
	if bytes.Equal(patch10ContractBin, GetPatch9ContractBin()) {
		t.Fatal("Patch10 bytecode must differ from Patch9")
	}
}

// TestLatestContractBinMatchesPatch10 pins the fresh-activation bytecode:
// mainnet's first SfcV2 activation (scheduled 2026-08-29) installs
// GetLatestContractBin(), which must be the Cycle-165 lockup-preservation
// bytecode so mainnet never runs the Cycle-164 chunked-settlement lockup
// destruction at all.
func TestLatestContractBinMatchesPatch10(t *testing.T) {
	if !bytes.Equal(GetLatestContractBin(), GetPatch10ContractBin()) {
		t.Fatal("GetLatestContractBin must return the Patch10 lockup-preservation bytecode for fresh SfcV2 activations")
	}
}
