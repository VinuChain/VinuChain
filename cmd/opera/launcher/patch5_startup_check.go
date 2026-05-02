package launcher

import (
	"testing"

	"github.com/Fantom-foundation/go-opera/opera/contracts/sfc"
)

// init wires the SfcV2Patch5 bytecode validity check into the opera binary's
// startup so an invalid Patch5 asset (placeholder sentinel, all-zero, byte-
// identical to Patch4, or below the minimum SFC size) refuses to start the
// node rather than only crashing at the on-chain re-flash block.
//
// The check is skipped inside `go test` runs via testing.Testing() so that
// the placeholder-still-present state of the v2.0.13-elemont scaffolding
// branch does not break every test that transitively imports this package.
// At ship time, when the real Cycle-161 bytecode lands in opera/contracts/sfc,
// the check is a no-op regardless of context.
//
// Critically, this guard ALSO refuses to start a v2.0.13 binary that
// accidentally has SfcV2Patch5 flipped on in any rule constructor while
// the bytecode is still the placeholder — without it, the binary would
// boot, sync up to the activation block, and only crash at re-flash time,
// stranding the validator. Keeping the check at startup makes the failure
// fast and obvious.
func init() {
	if testing.Testing() {
		return
	}
	sfc.EnforcePatch5StartupCheck()
}
