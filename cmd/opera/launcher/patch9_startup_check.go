package launcher

import (
	"testing"

	"github.com/Fantom-foundation/go-opera/opera/contracts/sfc"
)

// init wires the SfcV2Patch9 bytecode validity check into the opera binary's
// startup so an invalid Patch9 asset (placeholder sentinel, all-zero, byte-
// identical to Patch7/Patch8, or below the minimum SFC size) refuses to start
// the node rather than only crashing at the on-chain re-flash block.
func init() {
	if testing.Testing() {
		return
	}
	sfc.EnforcePatch9StartupCheck()
}
