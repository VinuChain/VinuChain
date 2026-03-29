package gossip

import (
	"sync"
	"time"
)

const (
	peerRateLimitWindow  = 10 * time.Second
	peerRateLimitMaxMsgs = 500
)

type peerRateLimiter struct {
	mu      sync.Mutex
	buckets map[string]*rateBucket
}

type rateBucket struct {
	count       int
	windowStart time.Time
}

func newPeerRateLimiter() *peerRateLimiter {
	return &peerRateLimiter{
		buckets: make(map[string]*rateBucket),
	}
}

func (r *peerRateLimiter) Allow(peerID string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	b, ok := r.buckets[peerID]
	if !ok || now.Sub(b.windowStart) >= peerRateLimitWindow {
		r.buckets[peerID] = &rateBucket{count: 1, windowStart: now}
		return true
	}
	b.count++
	return b.count <= peerRateLimitMaxMsgs
}

func (r *peerRateLimiter) RemovePeer(peerID string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.buckets, peerID)
}

// peerEventQuota tracks how many items from each peer are currently pending
// in the heavy check queue. This prevents a single peer from monopolizing
// the validation pipeline.
type peerEventQuota struct {
	mu       sync.Mutex
	pending  map[string]int
	maxPerPeer int
}

const (
	defaultMaxEventsPerPeer  = 200
	defaultMaxStreamsPerPeer = 100
)

func newPeerEventQuota() *peerEventQuota {
	return &peerEventQuota{
		pending:    make(map[string]int),
		maxPerPeer: defaultMaxEventsPerPeer,
	}
}

func newPeerStreamQuota() *peerEventQuota {
	return &peerEventQuota{
		pending:    make(map[string]int),
		maxPerPeer: defaultMaxStreamsPerPeer,
	}
}

// Acquire increments the pending count for a peer. Returns false if the
// peer has exceeded its quota (caller should drop the items).
func (q *peerEventQuota) Acquire(peerID string, n int) bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.pending[peerID]+n > q.maxPerPeer {
		return false
	}
	q.pending[peerID] += n
	return true
}

// Release decrements the pending count when processing completes.
func (q *peerEventQuota) Release(peerID string, n int) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.pending[peerID] -= n
	if q.pending[peerID] <= 0 {
		delete(q.pending, peerID)
	}
}

// RemovePeer cleans up quota tracking when a peer disconnects.
func (q *peerEventQuota) RemovePeer(peerID string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	delete(q.pending, peerID)
}
