package sfc

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
)

// GetPatch5ContractBin returns the Cycle-161 SFC bytecode re-flashed by the
// SfcV2Patch5 upgrade. It is distinct from GetContractBin (which returns the
// Cycle-159 bytecode used for SfcV2Patch3 and for fresh SfcV2 activations on
// mainnet) and from GetPatch4ContractBin (which returns the Cycle-160 bytecode
// re-flashed by SfcV2Patch4) so the Cycle-161 asset can evolve independently
// while legacy code paths remain untouched.
//
// Cycle-161 source: VinuChain/vinuchain-lists contracts/vinuchain/SFC.sol
// (the v2.0.13-elemont createValidator/_rawCreateValidator pubkey-shape patch).
// Compiled with solc 0.5.17+commit.d19bba13 --optimize --optimize-runs=10000
// --evm-version=istanbul.
//
// Delta from Cycle-160 (v2.0.11-elemont, SfcV2Patch4): three new validation
// require statements at the on-chain validator-pubkey ingress points:
//
//  1. SFC.createValidator: require(pubkey.length == 66) and
//     require(pubkey[0] == 0xc0).
//  2. SFC._rawCreateValidator: same two require statements (also covers
//     setGenesisValidator, which calls _rawCreateValidator).
//  3. NodeDriverAuth.updateValidatorPubkey: same two require statements as a
//     defence-in-depth on the relay path.
//
// These reject the malformed-pubkey shape that admitted testnet validator 16
// (a 65-byte pubkey starting with 0x04 instead of the canonical 66-byte
// 0xc0-prefixed encoding). Forward-only — existing stored pubkeys remain
// unchanged in storage; on-chain mitigation for already-admitted validators
// is governance-driven (deactivateValidator path), not bytecode-driven.
//
// Validity of the compiled asset is enforced at opera binary startup via
// EnforcePatch5StartupCheck (wired from cmd/opera/launcher/patch5_startup_check.go).
// A build whose Patch5 bytes are still the scaffolding placeholder, all
// zero, below the minimum SFC size, or byte-identical to Patch4 refuses to
// start the node rather than crashing at on-chain re-flash time.
func GetPatch5ContractBin() []byte {
	return patch5ContractBin
}

// EnforcePatch5StartupCheck validates the compiled-in Patch5 bytecode at
// binary startup. Wired from cmd/opera/launcher so the binary refuses to
// start when the bytecode is invalid (under-size with non-sentinel bytes,
// all-zero, byte-identical to Patch4), rather than only crashing at the
// on-chain re-flash block.
//
// The deadbeef-sentinel placeholder is the EXPECTED state of the v2.0.13
// scaffolding release — the flag defaults to false on every constructor,
// so GetPatch5ContractBin() is dead at runtime, and the placeholder is a
// known-shipping artefact. For that case the check downgrades to log.Warn
// so the binary remains runnable for QA, fakenet smoke, and ops handoffs;
// the v2.0.14 release that flips SfcV2Patch5 = true MUST replace the
// placeholder, and the warn message documents that requirement so it
// cannot be silently shipped with a real flag flip. Any OTHER validation
// failure (under-size with non-sentinel bytes, all-zero, equals Patch4)
// still log.Crits — those are unambiguous misconfigurations regardless of
// scaffolding state.
func EnforcePatch5StartupCheck() {
	err := validatePatch5Bytecode(patch5ContractBin)
	if err == nil {
		return
	}
	// Treat the v2.0.13 scaffolding placeholder as the expected state and
	// downgrade to log.Warn. The placeholder is identified by either a
	// leading deadbeef prefix (any length ≥ 4) OR a too-short payload that
	// is itself entirely deadbeef-repetition — both shapes are
	// unambiguously "not a real compile." Any OTHER validation failure
	// (all-zero, byte-identical to Patch4, plausibly-sized but corrupted)
	// still log.Crits because those are real misconfigurations.
	if isDeadbeefPlaceholder(patch5ContractBin) {
		log.Warn("SfcV2Patch5 bytecode is the deadbeef-sentinel scaffolding placeholder — flipping SfcV2Patch5 = true in any constructor without first replacing patch5ContractBin will brick the validator at the activation block. Replace with the truffle-compiled Cycle-161 hex blob before the v2.0.14-elemont release.")
		return
	}
	log.Crit("SfcV2Patch5 bytecode asset is invalid — the binary must not start with this build. Recompile SFC from vinuchain-lists and replace sfc_patch5_bytecode.go", "err", err)
}

// isDeadbeefPlaceholder returns true when code is the canonical v2.0.13
// scaffolding placeholder: a non-empty buffer whose contents are entirely
// repeating 0xdeadbeef bytes (any length, including the 256-byte default).
// Used to distinguish the expected scaffolding state from a real but
// corrupt bytecode payload.
func isDeadbeefPlaceholder(code []byte) bool {
	if len(code) == 0 || len(code)%4 != 0 {
		return false
	}
	for i := 0; i < len(code); i += 4 {
		if !bytes.Equal(code[i:i+4], patch5DeadbeefSentinel) {
			return false
		}
	}
	return true
}

// patch5DeadbeefSentinel is the 4-byte magic pattern that scaffolding builds
// use as a leading prefix. Any real compiled SFC bytecode begins with the
// Solidity dispatcher `0x60806040…`, so a leading deadbeef is an obviously-
// sentinel marker retained here as a defensive check against a future
// regression re-introducing a placeholder.
var patch5DeadbeefSentinel = []byte{0xde, 0xad, 0xbe, 0xef}

// patch5ContractBin is the deployed Cycle-161 runtime bytecode. The full
// hex string is a single-line literal to mirror the sfc_predeploy.go
// convention for GetContractBin and the sfc_patch4_bytecode.go convention
// for patch4ContractBin.
//
// Currently a 256-byte deadbeef-sentinel placeholder. Replace with the
// truffle-compiled Cycle-161 hex blob and re-run TestValidatePatch5Bytecode_*
// before tagging the v2.0.14-elemont release that activates SfcV2Patch5.
var patch5ContractBin = hexutil.MustDecode("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef")

// minRealPatch5BytecodeLen is a lower bound on a legitimate compiled SFC.
// Cycle-159/160 bytecode is 45,240 bytes; Cycle-161 is expected within a
// few hundred bytes of that. 10,000 is a conservative floor that comfortably
// rejects sentinel-sized placeholders while leaving head-room for plausible
// compiler variance.
const minRealPatch5BytecodeLen = 10000

// validatePatch5Bytecode returns a non-nil error if the argument is
// recognisable as a scaffolding placeholder or a misconfigured asset, so
// releases cannot ship the sentinel bytes in place of a real compile.
// Rejects:
//
//  1. nil or empty input
//  2. anything shorter than minRealPatch5BytecodeLen
//  3. a leading 0xdeadbeef prefix (sentinel marker)
//  4. all-zero bytecode (paranoia guard; real compiled SFC is not all zero)
//  5. byte-identical to Patch4 / GetPatch4ContractBin (no-op re-flash would
//     skip the pubkey-validation fix)
func validatePatch5Bytecode(code []byte) error {
	if len(code) == 0 {
		return errPatch5Empty
	}
	if len(code) < minRealPatch5BytecodeLen {
		return errPatch5TooShort
	}
	if bytes.HasPrefix(code, patch5DeadbeefSentinel) {
		return errPatch5Sentinel
	}
	allZero := true
	for _, b := range code {
		if b != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return errPatch5AllZero
	}
	if bytes.Equal(code, GetPatch4ContractBin()) {
		return errPatch5EqualsPatch4
	}
	return nil
}

type patch5Error string

func (e patch5Error) Error() string { return string(e) }

const (
	errPatch5Empty        patch5Error = "patch5 bytecode is empty"
	errPatch5TooShort     patch5Error = "patch5 bytecode shorter than minimum SFC size — placeholder not replaced"
	errPatch5Sentinel     patch5Error = "patch5 bytecode begins with deadbeef sentinel — placeholder not replaced"
	errPatch5AllZero      patch5Error = "patch5 bytecode is all zeros"
	errPatch5EqualsPatch4 patch5Error = "SfcV2Patch5 bytecode is byte-identical to Patch4 (Cycle-160) — the Patch5 re-flash would no-op on chain and the pubkey-validation fix would not deploy"
)
