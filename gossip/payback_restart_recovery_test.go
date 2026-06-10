package gossip

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// This file is the consensus-layer structural pin set for audit findings A1/T2.
// It complements payback/payback_restart_determinism_test.go, which proves the
// cache divergence at the cache layer and characterises the confirmed bug.
//
// CONFIRMED BUG — A1 is a real consensus issue:
//   A validator that restarts mid-epoch constructs an empty svc.paybackCache
//   (service.go:488). ReexecuteBlocks uses a LOCAL reexecCache that is never
//   installed back into svc.paybackCache. GetConsensusCallbacks then passes the
//   still-empty svc.paybackCache into every new block processor. For any address
//   that used FeeRefund quota earlier in the current epoch, the restarted node
//   computes a LARGER availableQuota → LARGER AddBalance → DIFFERENT state root.
//   This is state-root-affecting (not RPC-only): AddBalance in refundGas feeds
//   statedb.Commit() → block.Root → consensus-sealed.
//
// WHAT THESE TESTS PIN:
//   Group A (confirmed-bug structural facts): the wiring that causes the bug —
//     live cache is empty post-restart, reexecCache is local, no warm-up bridge.
//   Group B (still-correct facts): old receipts are never re-sealed by the
//     recovery path. This part of the previous analysis remains true.
//
// These tests use the source-structural pin idiom from the existing test files
// (gossip/payback_v2_activation_test.go, gossip/block_processor_test.go).
// A refactor that changes the wiring turns these tests red.

// ── Group A: Confirmed-bug structural facts ───────────────────────────────

// TestServiceCacheIsEmptyAtStartAndHasNoWarmUpPath pins that:
//  1. svc.paybackCache is constructed via NewPaybackCache with no replay of
//     the current epoch's history (empty construction at service.go:488).
//  2. There is no call that installs the epoch's accumulated quotaUsed into the
//     live cache before the engine starts sealing blocks.
//
// If a future fix adds a warm-up/replay path, this test should be updated to
// document the new wiring — and the A1 bug is likely resolved.
func TestServiceCacheIsEmptyAtStartAndHasNoWarmUpPath(t *testing.T) {
	src, err := os.ReadFile("service.go")
	require.NoError(t, err)
	s := string(src)

	// The live cache is built once, empty, at service construction time.
	require.Contains(t, s, "svc.paybackCache, err = payback.NewPaybackCache(paybackStore,",
		"svc.paybackCache must be constructed via NewPaybackCache (empty); "+
			"if warm-up/reload is added here the A1 bug may be fixed — update this test")

	// Locate Start() body to check for warm-up signals.
	startIdx := strings.Index(s, "func (s *Service) Start()")
	require.GreaterOrEqual(t, startIdx, 0, "Service.Start must exist")
	nextFunc := strings.Index(s[startIdx+1:], "\nfunc ")
	var startBody string
	if nextFunc >= 0 {
		startBody = s[startIdx : startIdx+1+nextFunc]
	} else {
		startBody = s[startIdx:]
	}

	require.Contains(t, startBody, "s.RecoverEVM()",
		"Start() must call RecoverEVM()")

	// These substrings would indicate a cache warm-up path. If any appear in
	// Start(), the A1 bug may be resolved — verify and update these tests.
	warmUpSignals := []string{
		"paybackCache.WarmUp",
		"paybackCache.Replay",
		"paybackCache.SetUsedMap",
		"paybackCache.LoadEpoch",
		"paybackCache.AddTransaction",
		"ReplayEpochIntoCache",
		"WarmPaybackCache",
	}
	for _, signal := range warmUpSignals {
		require.NotContains(t, startBody, signal,
			"Start() contains a payback cache warm-up call (%s) not present "+
				"when A1 was confirmed — the bug may be fixed; verify and update "+
				"these tests accordingly", signal)
	}
}

// TestReexecCacheIsLocalAndNotInstalledIntoLiveCache pins that ReexecuteBlocks
// uses a local reexecCache that is NEVER assigned to s.paybackCache. This is
// the structural reason the live cache remains empty after recovery.
//
// If a fix makes ReexecuteBlocks (or its caller) install accumulated state from
// reexecCache into s.paybackCache, A1 would be resolved and this test must be
// updated to document the new wiring.
func TestReexecCacheIsLocalAndNotInstalledIntoLiveCache(t *testing.T) {
	src, err := os.ReadFile("c_block_callbacks.go")
	require.NoError(t, err)
	s := string(src)

	reexecStart := strings.Index(s, "func (s *Service) ReexecuteBlocks(")
	require.GreaterOrEqual(t, reexecStart, 0, "ReexecuteBlocks must exist")
	recoverStart := strings.Index(s, "func (s *Service) RecoverEVM(")
	require.Greater(t, recoverStart, reexecStart, "RecoverEVM must follow ReexecuteBlocks")
	reexecBody := s[reexecStart:recoverStart]

	// A local reexecCache is built inside ReexecuteBlocks.
	require.Contains(t, reexecBody, "reexecCache, err := payback.NewPaybackCache(",
		"ReexecuteBlocks must build a local reexecCache")

	// The local cache is passed to EVMModule.Start (not s.paybackCache).
	require.Contains(t, reexecBody, "reexecCache, es.Epoch)",
		"ReexecuteBlocks must pass reexecCache (not s.paybackCache) to EVMModule.Start")

	// s.paybackCache must NOT be written inside ReexecuteBlocks. The one
	// legitimate read — s.paybackCache.GetStore() used to construct reexecCache
	// — is fine; we check for mutation signals only.
	require.NotContains(t, reexecBody, "s.paybackCache = ",
		"ReexecuteBlocks must NOT assign to s.paybackCache; if it does, "+
			"this is a potential fix for A1 — verify and update these tests")
	// These would indicate the local cache is being merged back into the live one.
	for _, mergeSignal := range []string{
		"s.paybackCache.AddTransaction",
		"s.paybackCache.PrepareForBlock",
		"s.paybackCache.SetUsedMap",
		"s.paybackCache.LoadEpoch",
		"s.paybackCache.WarmUp",
		"s.paybackCache.Replay",
	} {
		require.NotContains(t, reexecBody, mergeSignal,
			"ReexecuteBlocks must NOT write into s.paybackCache via %s; if it does, "+
				"this is a potential fix for A1 — verify and update these tests",
			mergeSignal)
	}
}

// TestGetConsensusCallbacksPassesLiveCache pins that GetConsensusCallbacks
// passes s.paybackCache (the live, post-restart empty cache) into the block
// processor — not a purpose-built warm cache. Combined with the tests above,
// this establishes the complete confirmed-bug path:
//   empty live cache → GetConsensusCallbacks → block processor →
//   GetAvailablePaybackByAddress returns quotaUsed=0 → excess AddBalance →
//   wrong state root.
func TestGetConsensusCallbacksPassesLiveCache(t *testing.T) {
	src, err := os.ReadFile("c_block_callbacks.go")
	require.NoError(t, err)
	s := string(src)

	cbStart := strings.Index(s, "func (s *Service) GetConsensusCallbacks()")
	require.GreaterOrEqual(t, cbStart, 0, "GetConsensusCallbacks must exist")
	reexecStart := strings.Index(s, "func (s *Service) ReexecuteBlocks(")
	require.Greater(t, reexecStart, cbStart, "ReexecuteBlocks must follow GetConsensusCallbacks")
	cbBody := s[cbStart:reexecStart]

	require.Contains(t, cbBody, "s.paybackCache,",
		"GetConsensusCallbacks must pass s.paybackCache to the block processor; "+
			"this is the live (post-restart: empty) cache that reaches forward sealing")
}

// ── Group B: Still-correct facts (old receipts not re-sealed) ────────────

// TestReexecutionDoesNotReseal pins that ReexecuteBlocks does NOT call
// SetReceipts. Already-sealed receipts from the current epoch are never
// overwritten by the recovery path. The A1 bug is only in FORWARD sealing.
func TestReexecutionDoesNotReseal(t *testing.T) {
	src, err := os.ReadFile("c_block_callbacks.go")
	require.NoError(t, err)
	s := string(src)

	reexecStart := strings.Index(s, "func (s *Service) ReexecuteBlocks(")
	require.GreaterOrEqual(t, reexecStart, 0)
	recoverStart := strings.Index(s, "func (s *Service) RecoverEVM(")
	require.Greater(t, recoverStart, reexecStart)
	reexecBody := s[reexecStart:recoverStart]

	require.NotContains(t, reexecBody, "SetReceipts",
		"ReexecuteBlocks must NOT call SetReceipts; re-sealing old receipts from "+
			"a fresh cache would compound the A1 bug further")

	require.Contains(t, reexecBody, "s.store.evm.Commit(b, block.Root, false)",
		"ReexecuteBlocks must commit under the already-sealed block.Root")
}

// TestRecoverEVMOnlyReexecutesTrailingUnpersistedBlocks pins that recovery is
// bounded to the not-yet-persisted trailing window. The A1 bug therefore cannot
// affect blocks whose EVM state was already flushed to disk; it is limited to
// blocks in the trailing in-memory window within the current epoch.
func TestRecoverEVMOnlyReexecutesTrailingUnpersistedBlocks(t *testing.T) {
	src, err := os.ReadFile("c_block_callbacks.go")
	require.NoError(t, err)
	s := string(src)

	require.Contains(t, s, "s.store.evm.HasStateDB(block.Root)",
		"RecoverEVM must locate the recovery anchor via HasStateDB")
	require.Contains(t, s, "s.ReexecuteBlocks(b, start)",
		"RecoverEVM must re-derive only the trailing [anchor, head] window")
}

// TestLiveBlockPathSealsReceiptsExactlyOnce pins that the live block processor
// is the sole caller of SetReceipts. The A1 bug manifests here: when the cache
// is empty post-restart, this forward-sealing path produces the wrong FeeRefund,
// wrong AddBalance, and consequently the wrong state root.
func TestLiveBlockPathSealsReceiptsExactlyOnce(t *testing.T) {
	src, err := os.ReadFile("block_processor.go")
	require.NoError(t, err)
	s := string(src)

	require.Contains(t, s, "bp.store.evm.SetReceipts(bp.blockCtx.Idx, allReceipts)",
		"the live block processor must be the receipt sealer; the A1 bug is that "+
			"this path uses an empty cache post-restart, producing wrong FeeRefund")
}
