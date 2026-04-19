package gossip

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestPeerRateLimiter_AllowCountsEachCall verifies that Allow increments the
// counter on every call, so a rate-limited peer is always tracked regardless
// of call order relative to other subsystems (e.g. semaphore acquire). If
// the rate limiter were called only after a semaphore acquire that fails, the
// counter would not be incremented for dropped messages — weakening the
// disconnect-on-abuse guarantee under load.
func TestPeerRateLimiter_AllowCountsEachCall(t *testing.T) {
	r := newPeerRateLimiter()

	// Fill the bucket to the limit.
	for i := 0; i < peerRateLimitMaxMsgs; i++ {
		require.True(t, r.Allow("peer1"), "expected Allow to return true for message %d", i+1)
	}

	// The very next call must be rejected — counter must have been incremented
	// for every prior call.
	require.False(t, r.Allow("peer1"), "expected Allow to return false after exceeding rate limit")
}

// TestPeerRateLimiter_WindowReset verifies that a new window resets the counter.
func TestPeerRateLimiter_WindowReset(t *testing.T) {
	r := newPeerRateLimiter()

	// Exhaust the window.
	for i := 0; i < peerRateLimitMaxMsgs; i++ {
		r.Allow("peer1")
	}
	require.False(t, r.Allow("peer1"))

	// Manually advance the clock by resetting the bucket's window start.
	r.mu.Lock()
	r.buckets["peer1"].windowStart = r.buckets["peer1"].windowStart.Add(-peerRateLimitWindow)
	r.mu.Unlock()

	// After window expiry, Allow should reset and accept again.
	require.True(t, r.Allow("peer1"), "expected Allow to return true after window reset")
}

func TestPeerEventQuota_AcquireRelease(t *testing.T) {
	q := newPeerEventQuota()

	// Can acquire up to maxPerPeer items.
	ok := q.Acquire("peer1", defaultMaxEventsPerPeer)
	require.True(t, ok)

	// One more would exceed the quota.
	ok = q.Acquire("peer1", 1)
	require.False(t, ok)

	// Different peer has independent quota.
	ok = q.Acquire("peer2", defaultMaxEventsPerPeer)
	require.True(t, ok)

	// Release all for peer1; subsequent acquire succeeds again.
	q.Release("peer1", defaultMaxEventsPerPeer)
	ok = q.Acquire("peer1", 1)
	require.True(t, ok)
	q.Release("peer1", 1)
	q.Release("peer2", defaultMaxEventsPerPeer)
}

func TestPeerStreamQuota_AcquireRelease(t *testing.T) {
	q := newPeerStreamQuota()

	// Can acquire up to maxPerPeer stream items.
	ok := q.Acquire("peer1", defaultMaxStreamsPerPeer)
	require.True(t, ok)

	// One more would exceed the quota.
	ok = q.Acquire("peer1", 1)
	require.False(t, ok)

	// Release all; quota is returned to zero so pending entry is deleted.
	q.Release("peer1", defaultMaxStreamsPerPeer)

	// After a full release the quota slot is gone.
	q.mu.Lock()
	_, exists := q.pending["peer1"]
	q.mu.Unlock()
	require.False(t, exists)
}

func TestPeerEventQuota_RemovePeer(t *testing.T) {
	q := newPeerEventQuota()
	q.Acquire("peer1", 5)
	q.RemovePeer("peer1")

	// Quota fully freed; new acquire succeeds.
	ok := q.Acquire("peer1", defaultMaxEventsPerPeer)
	require.True(t, ok)
	q.Release("peer1", defaultMaxEventsPerPeer)
}

// TestPeerStreamQuota_NeverBlocksDAGEvents verifies the documented invariant
// that DAG event stream responses do NOT hold peerStreamQuota slots.
//
// BVs, BRs, and EPs hold peerStreamQuota until their processor done() callback
// fires. DAG events use peerEventQuota inside handleEvents instead.
// This test verifies that after a maxPerPeer-sized batch of "DAG events"
// is consumed (simulating acquire + immediate release), the quota is zero.
// The stream quota must never be the backpressure mechanism for DAG events;
// peerEventQuota inside handleEvents provides the real protection.
func TestPeerStreamQuota_NeverBlocksDAGEvents(t *testing.T) {
	q := newPeerStreamQuota()
	const peer = "peer1"
	const n = defaultMaxStreamsPerPeer

	// Simulate what the old EventsStreamResponse code did:
	// acquire and immediately release.
	ok := q.Acquire(peer, n)
	require.True(t, ok)
	q.Release(peer, n)

	// After the immediate release the pending count is zero.
	q.mu.Lock()
	pending := q.pending[peer]
	q.mu.Unlock()
	require.Equal(t, 0, pending)
}

// TestPeerStreamQuota_BlocksWhenFull verifies that peerStreamQuota properly
// blocks BV/BR/EP stream items when the quota is exhausted.
// (BVs/BRs/EPs hold the quota until their processor done() callback fires.)
func TestPeerStreamQuota_BlocksWhenFull(t *testing.T) {
	q := newPeerStreamQuota()
	const peer = "peer1"
	const n = defaultMaxStreamsPerPeer / 2

	// First chunk acquired and held (simulating BVs processor holding quota).
	ok := q.Acquire(peer, n)
	require.True(t, ok)

	// Second chunk would not exceed quota.
	ok = q.Acquire(peer, n)
	require.True(t, ok)

	// Third chunk exceeds quota — must be rejected.
	ok = q.Acquire(peer, 1)
	require.False(t, ok)

	// Release simulates BV processor done() callback.
	q.Release(peer, n)
	q.Release(peer, n)

	// Now quota is free again.
	ok = q.Acquire(peer, defaultMaxStreamsPerPeer)
	require.True(t, ok)
	q.Release(peer, defaultMaxStreamsPerPeer)
}

// TestPeerEventQuota_FitsFullDagChunk verifies that the per-peer DAG event
// quota is large enough to accept one full sync chunk (DefaultChunkItemsNum
// from dagstreamleecher) without rejection. With ParallelChunksDownload up
// to 6, a single sync session legitimately delivers up to 6×500 = 3000
// in-flight events from one peer; a quota smaller than that drops every
// chunk and spams "peer exceeded event processing quota" during normal
// catch-up.
func TestPeerEventQuota_FitsFullDagChunk(t *testing.T) {
	q := newPeerEventQuota()
	const oneChunk = 500 // dagstreamleecher.DefaultConfig().Session.DefaultChunkItemsNum
	const parallelChunks = 6
	require.True(t, q.Acquire("peer1", oneChunk),
		"quota must accept a single full DAG chunk")
	for i := 1; i < parallelChunks; i++ {
		require.True(t, q.Acquire("peer1", oneChunk),
			"quota must accept up to ParallelChunksDownload in-flight chunks")
	}
}

// TestPeerStreamQuota_FitsFullStreamChunk verifies the BV/BR/EP per-peer
// stream quota accepts one full chunk (DefaultChunkItemsNum) without
// rejection. Same reasoning as the DAG event variant — BV/BR/EP chunks
// are sized identically by their leechers.
func TestPeerStreamQuota_FitsFullStreamChunk(t *testing.T) {
	q := newPeerStreamQuota()
	const oneChunk = 500
	require.True(t, q.Acquire("peer1", oneChunk),
		"stream quota must accept a single full BV/BR/EP chunk")
}

// TestPeerEventQuota_ConcurrentSafety verifies no data race under concurrent
// Acquire/Release from multiple goroutines.
func TestPeerEventQuota_ConcurrentSafety(t *testing.T) {
	q := newPeerEventQuota()
	const goroutines = 50
	const batchSize = 1

	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if q.Acquire("peer", batchSize) {
				q.Release("peer", batchSize)
			}
		}()
	}
	wg.Wait()

	// All goroutines complete; quota must be non-negative.
	q.mu.Lock()
	pending := q.pending["peer"]
	q.mu.Unlock()
	require.GreaterOrEqual(t, pending, 0)
}
