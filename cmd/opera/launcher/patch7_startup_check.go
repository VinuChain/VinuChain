package launcher

import (
	"testing"

	"github.com/Fantom-foundation/go-opera/opera/contracts/sfc"
)

// init wires the SfcV2Patch7 bytecode validity check into the opera binary's
// startup so an invalid Patch7 asset (placeholder sentinel, all-zero, byte-
// identical to Patch5/Patch6, or below the minimum SFC size) refuses to start
// the node rather than only crashing at the on-chain re-flash block.
func init() {
	if testing.Testing() {
		return
	}
	sfc.EnforcePatch7StartupCheck()
}
