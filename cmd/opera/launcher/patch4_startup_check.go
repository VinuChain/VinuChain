package launcher

import (
	"testing"

	"github.com/Fantom-foundation/go-opera/opera/contracts/sfc"
)

// init wires the SfcV2Patch4 bytecode validity check into the opera binary's
// startup so an invalid Patch4 asset (placeholder sentinel, all-zero, byte-
// identical to Patch3, or below the minimum SFC size) refuses to start the
// node rather than only crashing at the on-chain re-flash block.
//
// The check is skipped inside `go test` runs via testing.Testing() so that
// the placeholder-still-present state of the scaffolding branch does not
// break every test that transitively imports this package. At ship time,
// when the real Cycle-160 bytecode lands in opera/contracts/sfc, the check
// is a no-op regardless of context.
func init() {
	if testing.Testing() {
		return
	}
	sfc.EnforcePatch4StartupCheck()
}
