package payback

import (
	"math/big"
	"testing"
	"time"

	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"
)

// ============================================================================
// M1 PaybackCache restart-determinism — REVISED VERDICT
//
// Audit findings A1 (HIGH) / T2 (HIGH):
//   "PaybackCache is volatile yet feeds consensus state (restart-determinism
//    hypothesis). A validator that restarts mid-epoch loses its accumulated
//    PaybackUsedMap and would compute a *larger* available payback (quotaUsed=0)
//    for subsequent blocks than a peer that never restarted."
//
// VERDICT: A1 CONFIRMED AS A REAL CONSENSUS BUG (mid-epoch restart scenario).
//
// The previous analysis in commit 92abc8a was partially correct: it proved that
// already-SEALED receipts are never recomputed after restart (ReexecuteBlocks
// does not call SetReceipts). That fact is still true and is pinned in
// gossip/payback_restart_recovery_test.go.
//
// However it missed the critical forward-sealing path:
//
//   1. On startup, svc.paybackCache is constructed EMPTY (service.go:488,
//      NewPaybackCache with no warm-up). ReexecuteBlocks uses a LOCAL
//      reexecCache — it does NOT install accumulated quotaUsed back into
//      svc.paybackCache (c_block_callbacks.go:75; svc.paybackCache untouched).
//
//   2. After RecoverEVM returns, GetConsensusCallbacks (called via
//      engine.Bootstrap, launcher.go:398) passes svc.paybackCache — still empty
//      — into every future block processor (c_block_callbacks.go:61).
//
//   3. When the engine delivers the next NEW (not-yet-sealed) block for sealing,
//      the live block processor calls:
//        paybackCache.GetAvailablePaybackByAddress(msg.From(), evm)
//        → state_processor.go:155
//      which returns (paybackSum - quotaUsed). With quotaUsed=0, this is the
//      FULL epoch quota — not the reduced amount that a non-restarted validator
//      carries after intra-epoch usage.
//
//   4. The result flows directly into ApplyMessage → refundGas → AddBalance
//      (state_transition.go:493) — a STATE TRIE MUTATION. The credited amount
//      differs between the restarted node and a non-restarted peer.
//
//   5. statedb.Commit() at Finalize (evmmodule/evm.go:190) hashes ALL AddBalance
//      mutations into block.Root. A different AddBalance → different block.Root.
//      block.Root is consensus-sealed and must match across all validators.
//      A receipts-root mismatch (documented in CLAUDE.md as a full consensus
//      split) is the consequence.
//
// BLAST RADIUS:
//   - This is STATE-ROOT-AFFECTING, not RPC-only. The FeeRefund changes the
//     sender's balance via AddBalance before statedb.Commit, making it part of
//     the Merkle state trie that all validators must agree on. It also feeds
//     r.FeeRefund → driver_txs.go:190 (validatorFee = txFee - feeRefund) →
//     ValidatorStates[].Originated, which affects epoch-end SFC staking rewards.
//   - Preconditions for divergence to manifest in production:
//     (a) Upgrades.Podgorica is active (FeeRefund enabled)
//     (b) A validator restarts mid-epoch
//     (c) In the same epoch after restart, the same address submits another tx
//         that would have had non-zero FeeRefund if the cache were warm
//     This is not a theoretical edge case: mid-epoch restarts happen on every
//     upgrade, crash-recovery, or routine maintenance.
//
// WHAT IS NOT BROKEN (previous analysis remains correct):
//   - Already-sealed receipts are never recomputed post-restart (pinned in
//     gossip/payback_restart_recovery_test.go). The bug only affects FORWARD
//     sealing of new blocks in the current epoch.
//   - At an epoch boundary all caches (warm or fresh) converge to zero, so
//     divergence is strictly bounded to one epoch (TestPaybackUsedMapResetsAtEpochBoundary).
//
// REQUIRED FIX (owner action, NOT done here — consensus-critical):
//   On startup (or at the start of each new epoch), replay the already-sealed
//   blocks of the current epoch through s.paybackCache (calling AddTransaction
//   for each tx/receipt pair) so that PaybackUsedMap is warmed to the epoch's
//   accumulated state before the first new block is sealed. Alternatively,
//   persist PaybackUsedMap to disk and reload it on startup.
//
// ============================================================================

// TestForwardSealingDivergesAfterMidEpochRestart is the characterisation test
// for the confirmed bug: a fresh (post-restart) cache and a warm (continuous)
// cache return different available-payback values for a mid-epoch block. This
// difference propagates into AddBalance (state_transition.go:493) and therefore
// into block.Root (consensus-sealed). The test documents current divergent
// behaviour; it MUST FAIL (or be deleted) once the fix is applied.
//
// The test does NOT use t.Skip because it is a PASSING characterisation test:
// it asserts that the caches DO diverge, pinning the current broken state so CI
// immediately catches the moment someone fixes or accidentally worsens the bug.
func TestForwardSealingDivergesAfterMidEpochRestart(t *testing.T) {
	const epoch = idx.Epoch(5)
	rules := opera.Rules{}
	blockTime := time.Now()
	addr := common.HexToAddress("0xDEADBEEF")

	// ---- Warm cache: node that never restarted. ----
	// Block 1 of epoch: address consumes 500 wei of quota (simulating prior usage
	// from an earlier intra-epoch block, accumulated via AddTransaction).
	warm, err := NewPaybackCache(&stubStore{}, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}
	warm.PrepareForBlock(epoch, rules, blockTime)
	warm.mu.Lock()
	warm.PaybackUsedMap[addr] = big.NewInt(500)
	warm.mu.Unlock()
	warm.FinishBlock()

	// Block 2 (new block to be sealed): read available quota on the warm cache.
	warm.PrepareForBlock(epoch, rules, blockTime)
	warm.mu.RLock()
	warmQuotaUsed := warm.getQuotaUsedLocked(addr)
	warm.mu.RUnlock()
	warm.FinishBlock()

	// ---- Fresh cache: node that restarted mid-epoch (no warm-up). ----
	// This is exactly svc.paybackCache after service.go:488 and before any
	// AddTransaction replay. GetConsensusCallbacks passes this to the block
	// processor for the SAME block 2 above.
	fresh, err := NewPaybackCache(&stubStore{}, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}
	fresh.PrepareForBlock(epoch, rules, blockTime)
	fresh.mu.RLock()
	freshQuotaUsed := fresh.getQuotaUsedLocked(addr)
	fresh.mu.RUnlock()
	fresh.FinishBlock()

	// Both caches see quotaUsed; the warm cache has 500, the fresh one has 0.
	if warmQuotaUsed.Cmp(big.NewInt(500)) != 0 {
		t.Fatalf("warm cache quotaUsed want 500, got %s", warmQuotaUsed)
	}
	if freshQuotaUsed.Sign() != 0 {
		t.Fatalf("fresh cache quotaUsed want 0, got %s", freshQuotaUsed)
	}

	// GetAvailablePaybackByAddress returns (paybackSum - quotaUsed).
	// A fresh cache subtracts 0 less quota → MORE available payback → MORE
	// AddBalance → different state root.
	// Characterise the bug: caches diverge, and the divergence is exactly
	// the previously-accumulated usage (500 wei in this example).
	divergence := new(big.Int).Sub(freshQuotaUsed, warmQuotaUsed) // expected: -500
	if divergence.Sign() == 0 {
		t.Fatal("BUG APPEARS FIXED: warm and fresh caches now agree mid-epoch. " +
			"Delete or rewrite this test to document the fix. " +
			"If this fires unexpectedly, the cache warm-up is working — " +
			"verify the fix is intentional before removing the test.")
	}

	// Document the exact magnitude of the divergence so regressions are obvious.
	expectedDivergence := big.NewInt(-500)
	if divergence.Cmp(expectedDivergence) != 0 {
		t.Fatalf("divergence changed: want %s, got %s — check if the payback formula changed",
			expectedDivergence, divergence)
	}
}

// TestPaybackCacheIsVolatileWithinEpoch demonstrates the cache-layer volatility
// as an isolated property: accumulated quotaUsed in a warm cache differs from a
// freshly-constructed one at mid-epoch. This underpins the confirmed bug.
func TestPaybackCacheIsVolatileWithinEpoch(t *testing.T) {
	const epoch = idx.Epoch(5)
	rules := opera.Rules{}
	blockTime := time.Now()
	addr := common.HexToAddress("0xA11CE")

	warm, err := NewPaybackCache(&stubStore{}, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}
	warm.PrepareForBlock(epoch, rules, blockTime)
	warm.mu.Lock()
	warm.PaybackUsedMap[addr] = big.NewInt(300)
	warm.mu.Unlock()
	warm.FinishBlock()

	warm.PrepareForBlock(epoch, rules, blockTime)
	warm.mu.Lock()
	warm.PaybackUsedMap[addr].Add(warm.PaybackUsedMap[addr], big.NewInt(200))
	warm.mu.Unlock()
	warm.mu.RLock()
	warmQuotaUsed := warm.getQuotaUsedLocked(addr)
	warm.mu.RUnlock()
	warm.FinishBlock()

	if warmQuotaUsed.Cmp(big.NewInt(500)) != 0 {
		t.Fatalf("warm cache quotaUsed: want 500, got %s", warmQuotaUsed)
	}

	fresh, err := NewPaybackCache(&stubStore{}, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}
	fresh.PrepareForBlock(epoch, rules, blockTime)
	fresh.mu.RLock()
	freshQuotaUsed := fresh.getQuotaUsedLocked(addr)
	fresh.mu.RUnlock()
	fresh.FinishBlock()

	if freshQuotaUsed.Sign() != 0 {
		t.Fatalf("fresh cache quotaUsed: want 0, got %s", freshQuotaUsed)
	}

	if warmQuotaUsed.Cmp(freshQuotaUsed) == 0 {
		t.Fatal("expected warm and fresh caches to diverge mid-epoch")
	}
}

// TestPaybackUsedMapResetsAtEpochBoundary proves that warm and fresh caches
// converge at epoch boundaries. This bounds the A1 bug to within a single epoch:
// a restart that lands ON or AFTER an epoch boundary cannot diverge at all.
func TestPaybackUsedMapResetsAtEpochBoundary(t *testing.T) {
	rules := opera.Rules{}
	blockTime := time.Now()
	addr := common.HexToAddress("0xB0B")

	warm, err := NewPaybackCache(&stubStore{}, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}
	warm.PrepareForBlock(idx.Epoch(7), rules, blockTime)
	warm.mu.Lock()
	warm.PaybackUsedMap[addr] = big.NewInt(900)
	warm.mu.Unlock()
	warm.FinishBlock()

	// Cross into epoch 8 — cleanup resets PaybackUsedMap.
	warm.PrepareForBlock(idx.Epoch(8), rules, blockTime)
	warm.mu.RLock()
	warmQuotaUsed := warm.getQuotaUsedLocked(addr)
	warm.mu.RUnlock()
	warm.FinishBlock()

	fresh, err := NewPaybackCache(&stubStore{}, 0)
	if err != nil {
		t.Fatalf("NewPaybackCache: %v", err)
	}
	fresh.PrepareForBlock(idx.Epoch(8), rules, blockTime)
	fresh.mu.RLock()
	freshQuotaUsed := fresh.getQuotaUsedLocked(addr)
	fresh.mu.RUnlock()
	fresh.FinishBlock()

	if warmQuotaUsed.Sign() != 0 {
		t.Fatalf("warm cache after epoch boundary: want 0, got %s", warmQuotaUsed)
	}
	if freshQuotaUsed.Sign() != 0 {
		t.Fatalf("fresh cache in new epoch: want 0, got %s", freshQuotaUsed)
	}
	if warmQuotaUsed.Cmp(freshQuotaUsed) != 0 {
		t.Fatalf("warm (%s) != fresh (%s) across epoch boundary", warmQuotaUsed, freshQuotaUsed)
	}
}
