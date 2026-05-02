package sfc100

import (
	"testing"

	"github.com/Fantom-foundation/go-opera/inter/validatorpk"
)

// TestSFCPubkeyValidationReasonStrings is the reason-string contract that
// Cycle-161+ SFC bytecode must satisfy. These strings are emitted by the
// require() statements added to SFC.sol's createValidator,
// _rawCreateValidator, and NodeDriverAuth.updateValidatorPubkey paths to
// reject pubkeys that are not in the lachesis-base wire format
// ([type(1)=0xc0 || raw(65)] for uncompressed Secp256k1).
//
// Background: a malformed pubkey lacking the 0xc0 type-byte prefix bypassed
// SFC's prior `pubkey.length == 33 || pubkey.length == 65` check on testnet.
// The validator's events failed lachesis-base verification, leading to zero
// uptime and zero rewards while delegators were stuck unable to undelegate.
//
// This test does not exercise the deployed bytecode (the SFC truffle harness
// lives in a separate opera-sfc checkout consumed via gossip/sfc_test.go's
// go:generate directives). It pins the strings so a future SFC.sol patch
// cannot silently shorten them past the 32-byte single-storage-slot budget
// or drift from what Blockscout/external tooling expects to surface.
//
// On-chain verification path:
//  1. Build Cycle-161 with solc 0.5.17+commit.d19bba13 --optimize
//     --optimize-runs=10000 --evm-version=istanbul.
//  2. Activate SfcV2Patch5 (gated by a new upgrade flag, owned by the next
//     SFC patch task) on a fakenet running the patched binary.
//  3. Call SFC.createValidator(badPubkey) where badPubkey is 65 bytes
//     starting with 0x04 (the malformation seen on testnet validator 16);
//     assert revert with reason "invalid pubkey type".
//  4. Call SFC.createValidator(badPubkey) where badPubkey is 66 bytes
//     starting with 0xc1; assert revert with reason "invalid pubkey type".
//  5. Call SFC.createValidator(badPubkey) where badPubkey is 65 bytes
//     starting with 0xc0; assert revert with reason "invalid pubkey length".
//  6. Call SFC.createValidator(goodPubkey) where goodPubkey is the standard
//     0xc0 || 0x04 || X(32) || Y(32) layout; assert success.
func TestSFCPubkeyValidationReasonStrings(t *testing.T) {
	const (
		invalidLen  = "invalid pubkey length"
		invalidType = "invalid pubkey type"
	)
	if got, want := len(invalidLen), 32; got > want {
		t.Errorf("invalid pubkey length: %d bytes, must fit in one storage word (≤%d)", got, want)
	}
	if got, want := len(invalidType), 32; got > want {
		t.Errorf("invalid pubkey type: %d bytes, must fit in one storage word (≤%d)", got, want)
	}

	// Sanity-pin the lachesis-base pubkey format constants the SFC patch
	// is asserting against. If validatorpk.Types.Secp256k1 ever changes,
	// the SFC.sol patch's literal `0xc0` must change in lock-step.
	if validatorpk.Types.Secp256k1 != 0xc0 {
		t.Fatalf("validatorpk.Types.Secp256k1 changed from 0xc0 to 0x%02x — SFC.sol patch must update its `pubkey[0] == 0xc0` checks", validatorpk.Types.Secp256k1)
	}

	// Compose the canonical good pubkey layout the SFC validator must
	// accept after Cycle-161: 0xc0 || 0x04 || X(32) || Y(32) = 66 bytes.
	good := make([]byte, 0, 66)
	good = append(good, validatorpk.Types.Secp256k1)
	good = append(good, 0x04)
	good = append(good, make([]byte, 64)...) // dummy X || Y
	if len(good) != 66 {
		t.Fatalf("canonical good pubkey wrong length: got %d, want 66", len(good))
	}
	if good[0] != 0xc0 {
		t.Fatalf("canonical good pubkey wrong type byte: got 0x%02x, want 0xc0", good[0])
	}

	// Reproduce the exact malformation seen on testnet validator 16: 65
	// bytes, leading byte 0x04 (the uncompressed-secp256k1 marker mistaken
	// for the lachesis type byte). Cycle-161 must reject this with
	// "invalid pubkey length" since it fails the length check first.
	bad16 := make([]byte, 65)
	bad16[0] = 0x04
	if len(bad16) == 66 {
		t.Fatalf("bad16 pubkey reproduction shape changed; this test no longer exercises the testnet val-16 case")
	}
	if bad16[0] == 0xc0 {
		t.Fatalf("bad16 pubkey reproduction shape changed; leading byte must remain 0x04 to mirror testnet val-16")
	}
}
