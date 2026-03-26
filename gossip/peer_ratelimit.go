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
