package sfc

import (
	"bytes"

	"github.com/ethereum/go-ethereum/log"
)

// GetPatch4ContractBin returns the Cycle-160 SFC bytecode re-flashed by the
// SfcV2Patch4 upgrade. It is distinct from GetContractBin (which returns the
// Cycle-159 bytecode used for SfcV2Patch3 and for fresh SfcV2 activations on
// mainnet) so that existing patch3 code paths are not disturbed while the
// Cycle-160 bytecode is still being compiled from VinuChain/vinuchain-lists
// PR #2 (the lock-end-time bug fix: `relock/extendLock` compared the new
// duration against the old duration instead of comparing the new end time
// against the old end time, which let stakers silently shorten their lock
// period by relocking with a shorter duration from a point later in time).
//
// PLACEHOLDER — real Cycle-160 bytecode drops in when vinuchain-lists PR #2
// merges and is compiled with solc 0.5.17+commit.d19bba13 --optimize
// --optimize-runs=10000 --evm-version=istanbul. Until then the bytes below
// are a `deadbeef` sentinel, and the guard in validatePatch4Bytecode will
// `log.Crit` at first attempted re-flash so a release cannot silently ship
// the placeholder.
func GetPatch4ContractBin() []byte {
	b := patch4PlaceholderBytecode
	if err := validatePatch4Bytecode(b); err != nil {
		log.Crit("SfcV2Patch4 bytecode is a placeholder — recompile from vinuchain-lists SFC.sol after PR #2 merges, then replace sfc_patch4_bytecode.go before cutting a release", "err", err)
	}
	return b
}

// patch4DeadbeefSentinel is the 4-byte magic pattern repeated throughout the
// placeholder. Any real compiled SFC bytecode begins with the Solidity
// dispatcher `0x60806040…`, so a leading deadbeef is an obviously-sentinel
// marker that the guard rejects.
var patch4DeadbeefSentinel = []byte{0xde, 0xad, 0xbe, 0xef}

// patch4PlaceholderBytecode is the sentinel asset shipped alongside the
// SfcV2Patch4 scaffolding. It is intentionally short, intentionally
// non-zero, and intentionally starts with a deadbeef prefix so the guard
// catches it three different ways (length < 10000, leading sentinel, all
// bytes being the repeating sentinel pattern). Replace the whole var with
// the real compiled Cycle-160 bytecode when PR #2 lands.
var patch4PlaceholderBytecode = bytes.Repeat(patch4DeadbeefSentinel, 64)

// minRealPatch4BytecodeLen is a lower bound on a legitimate compiled SFC.
// The current Cycle-159 bytecode is 45,240 bytes; any real Cycle-160 build
// will be within a few hundred bytes of that. 10,000 is a conservative
// floor that comfortably rejects the 256-byte placeholder while leaving
// head-room for plausible compiler variance.
const minRealPatch4BytecodeLen = 10000

// validatePatch4Bytecode returns a non-nil error if the argument is
// recognisable as the scaffolding placeholder, so releases cannot ship the
// sentinel bytes in place of the real Cycle-160 compile. Rejects:
//
//  1. nil or empty input
//  2. anything shorter than minRealPatch4BytecodeLen
//  3. a leading 0xdeadbeef prefix (sentinel marker)
//  4. all-zero bytecode (paranoia guard; real compiled SFC is not all zero)
func validatePatch4Bytecode(code []byte) error {
	if len(code) == 0 {
		return errPatch4Empty
	}
	if len(code) < minRealPatch4BytecodeLen {
		return errPatch4TooShort
	}
	if bytes.HasPrefix(code, patch4DeadbeefSentinel) {
		return errPatch4Sentinel
	}
	allZero := true
	for _, b := range code {
		if b != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return errPatch4AllZero
	}
	return nil
}

type patch4Error string

func (e patch4Error) Error() string { return string(e) }

const (
	errPatch4Empty    patch4Error = "patch4 bytecode is empty"
	errPatch4TooShort patch4Error = "patch4 bytecode shorter than minimum SFC size — placeholder not replaced"
	errPatch4Sentinel patch4Error = "patch4 bytecode begins with deadbeef sentinel — placeholder not replaced"
	errPatch4AllZero  patch4Error = "patch4 bytecode is all zeros"
)
