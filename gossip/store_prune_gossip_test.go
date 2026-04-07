package gossip

import (
	"testing"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-opera/inter"
)

func TestPruneEpochData_DeletesEvents(t *testing.T) {
	s := NewMemStore()
	defer s.Close()

	e := fakeEventForEpoch(1)
	s.SetEvent(e)

	id := e.ID()
	require.NotNil(t, s.GetEventPayload(id), "event must exist before pruning")

	count, err := s.PruneEpochData(idx.Epoch(1))
	require.NoError(t, err)
	require.Positive(t, count, "must report at least one deleted entry")

	require.Nil(t, s.GetEventPayload(id), "event must be gone after pruning epoch 1")
}

func TestPruneEpochData_KeepsNewerEpoch(t *testing.T) {
	s := NewMemStore()
	defer s.Close()

	e1 := fakeEventForEpoch(1)
	e2 := fakeEventForEpoch(2)
	s.SetEvent(e1)
	s.SetEvent(e2)

	id1, id2 := e1.ID(), e2.ID()

	_, err := s.PruneEpochData(idx.Epoch(1))
	require.NoError(t, err)

	require.Nil(t, s.GetEventPayload(id1), "epoch 1 event must be gone")
	require.NotNil(t, s.GetEventPayload(id2), "epoch 2 event must still be present")
}

// fakeEventForEpoch builds a minimal EventPayload for the given epoch.
func fakeEventForEpoch(epoch idx.Epoch) *inter.EventPayload {
	var e inter.MutableEventPayload
	e.SetVersion(1)
	e.SetEpoch(epoch)
	e.SetLamport(idx.Lamport(epoch)) // unique lamport per epoch
	e.SetParents(hash.Events{})
	return e.Build()
}
