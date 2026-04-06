package eventcheck_test

import (
	"testing"

	"github.com/Fantom-foundation/go-opera/eventcheck"
	"github.com/Fantom-foundation/go-opera/eventcheck/heavycheck"
)

// TestIsBan_UnknownEpochEventLocator verifies that ErrUnknownEpochEventLocator
// is NOT classified as a ban-worthy error. A node receiving an event with a
// MisbehaviourProof whose locator epoch is unknown may simply lack the historical
// epoch data — this is a transient sync condition, not peer misbehaviour.
func TestIsBan_UnknownEpochEventLocator(t *testing.T) {
	if eventcheck.IsBan(heavycheck.ErrUnknownEpochEventLocator) {
		t.Error("ErrUnknownEpochEventLocator should not be a ban-worthy error: " +
			"a node may legitimately lack data for old epochs referenced in MPs")
	}
}

// TestIsBan_UnknownEpochBVsNotBan verifies the existing non-ban classification for BVs.
func TestIsBan_UnknownEpochBVsNotBan(t *testing.T) {
	if eventcheck.IsBan(eventcheck.ErrUnknownEpochBVs) {
		t.Error("ErrUnknownEpochBVs should not be a ban-worthy error")
	}
}

// TestIsBan_UnknownEpochEVNotBan verifies the existing non-ban classification for EV.
func TestIsBan_UnknownEpochEVNotBan(t *testing.T) {
	if eventcheck.IsBan(eventcheck.ErrUnknownEpochEV) {
		t.Error("ErrUnknownEpochEV should not be a ban-worthy error")
	}
}

// TestIsBan_BanWorthyErrors verifies that genuinely bad signatures ARE ban-worthy.
func TestIsBan_BanWorthyErrors(t *testing.T) {
	banWorthy := []error{
		heavycheck.ErrWrongEventSig,
		heavycheck.ErrMalformedTxSig,
		heavycheck.ErrWrongPayloadHash,
	}
	for _, err := range banWorthy {
		if !eventcheck.IsBan(err) {
			t.Errorf("expected %v to be ban-worthy", err)
		}
	}
}
