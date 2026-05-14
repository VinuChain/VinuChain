package launcher

import (
	"testing"

	"github.com/Fantom-foundation/go-opera/opera"
)

// init wires the PaybackV2 sentinel check into the opera binary's startup
// so a build that flips PaybackV2=true on a rule constructor without
// recording the deployed V2 contract address in opera/payback_v2_address.go
// refuses to start rather than booting, syncing up to the activation block,
// and only failing at the seal-time switch (which would log.Crit the node
// at the worst possible moment).
//
// Skipped under `go test` runs so the placeholder-still-present scaffold
// state does not break tests that transitively import this package — the
// scaffold ships with PaybackV2=false in all real rule constructors, so
// the test-time skip does not mask any production risk. When the real V2
// address lands in opera/payback_v2_address.go and PaybackV2 is enabled
// in VinuChainTestNetRules / VinuChainMainNetRules, the check is a no-op
// regardless of context.
func init() {
	if testing.Testing() {
		return
	}
	opera.EnforcePaybackV2StartupCheck()
}
