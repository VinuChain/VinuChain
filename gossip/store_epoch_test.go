package gossip

import (
	"testing"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/stretchr/testify/require"
)

// TestGetEpochStoreNilGuard verifies that getEpochStore returns nil without
// panicking when the epoch store has not yet been initialized. Before the fix,
// getAnyEpochStore() returned nil and es.epoch then caused a nil pointer panic.
func TestGetEpochStoreNilGuard(t *testing.T) {
	var s Store
	// epochStore atomic.Value starts as nil (zero value).
	// getEpochStore must return nil, not panic.
	es := s.getEpochStore(idx.Epoch(1))
	require.Nil(t, es)
}
