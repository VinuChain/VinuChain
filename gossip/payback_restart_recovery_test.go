package gossip

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// This file is the consensus-layer structural pin set for audit findings A1/T2.
// It complements payback/payback_restart_determinism_test.go, which proves at the
// cache layer that a warmed cache converges with a never-restarted node.
//
// A1/T2 — RESOLVED:
//   A validator that restarts mid-epoch constructs an empty svc.paybackCache
//   (service.go). RecoverEVM re-derives only trailing EVM state into a LOCAL
//   reexecCache. The fix is a forward-sealing warm-up: Start() now calls
//   s.WarmUpPaybackCache() AFTER RecoverEVM and BEFORE the engine
//   (engine.Bootstrap → GetConsensusCallbacks) hands s.paybackCache to the live
//   block processor. WarmUpPaybackCache replays the current epoch's already-sealed
//   (tx, receipt) pairs through s.paybackCache via PrepareForBlock/AddTransaction/
//   FinishBlock, mirroring the live derivation (the per-block epoch comes from
//   GetHistoryEpochState(FindBlockEpoch(b)).Epoch, exactly as ReexecuteBlocks).
//   The warmed PaybackUsedMap therefore matches a never-restarted peer at head, so
//   the first new block seals an identical FeeRefund / block.Root.
//
// WHAT THESE TESTS PIN:
//   Group A (fix wiring): Start() warms the live cache after RecoverEVM and before
//     forward sealing; the warm-up replays via s.paybackCache.AddTransaction; the
//     epoch derivation mirrors the live path.
//   Group B (still-correct facts): old receipts are never re-sealed by the
//     recovery path; ReexecuteBlocks still uses its own local reexecCache.
//
// These tests use the source-structural pin idiom from the existing test files
// (gossip/payback_v2_activation_test.go, gossip/block_processor_test.go).
// A refactor that removes or reorders the warm-up turns these tests red.

// ── Group A: Fix-wiring structural facts ──────────────────────────────────

// TestServiceWarmsPaybackCacheBeforeForwardSealing pins that:
//  1. svc.paybackCache is still constructed via NewPaybackCache (empty) at
//     service construction time.
//  2. Start() now calls s.WarmUpPaybackCache() AFTER s.RecoverEVM() and before
//     the block processor is started — installing the current epoch's accumulated
//     quotaUsed into the live cache before any new block is sealed. This is the
//     A1 fix.
func TestServiceWarmsPaybackCacheBeforeForwardSealing(t *testing.T) {
	src, err := os.ReadFile("service.go")
	require.NoError(t, err)
	s := string(src)

	// The live cache is built once, empty, at service construction time.
	require.Contains(t, s, "svc.paybackCache, err = payback.NewPaybackCache(paybackStore,",
		"svc.paybackCache must be constructed via NewPaybackCache (empty) at service construction")

	// Locate Start() body.
	startIdx := strings.Index(s, "func (s *Service) Start()")
	require.GreaterOrEqual(t, startIdx, 0, "Service.Start must exist")
	nextFunc := strings.Index(s[startIdx+1:], "\nfunc ")
	var startBody string
	if nextFunc >= 0 {
		startBody = s[startIdx : startIdx+1+nextFunc]
	} else {
		startBody = s[startIdx:]
	}

	recoverIdx := strings.Index(startBody, "s.RecoverEVM()")
	require.GreaterOrEqual(t, recoverIdx, 0, "Start() must call RecoverEVM()")

	warmIdx := strings.Index(startBody, "s.WarmUpPaybackCache()")
	require.GreaterOrEqual(t, warmIdx, 0,
		"Start() must call s.WarmUpPaybackCache() — this is the A1 fix; if it is "+
			"removed, mid-epoch restarts diverge from non-restarted peers")

	// The warm-up must run AFTER state recovery (so trailing EVM state exists).
	require.Greater(t, warmIdx, recoverIdx,
		"WarmUpPaybackCache() must be called after RecoverEVM()")

	// The warm-up must run BEFORE the block processor starts forward sealing.
	blockProcIdx := strings.Index(startBody, "s.blockProcTasks.Start(")
	require.GreaterOrEqual(t, blockProcIdx, 0, "Start() must start the block processor")
	require.Less(t, warmIdx, blockProcIdx,
		"WarmUpPaybackCache() must run before the block processor starts so the "+
			"first sealed block reads the warmed (correct) quotaUsed")
}

// TestWarmUpReplaysCurrentEpochIntoLiveCache pins that WarmUpPaybackCache replays
// already-sealed (tx, receipt) pairs into the LIVE s.paybackCache using the same
// lifecycle the live block processor uses, and derives each block's epoch via the
// same path as ReexecuteBlocks. This is what makes the warmed cache identical to
// a never-restarted node's.
func TestWarmUpReplaysCurrentEpochIntoLiveCache(t *testing.T) {
	src, err := os.ReadFile("c_block_callbacks.go")
	require.NoError(t, err)
	s := string(src)

	warmStart := strings.Index(s, "func (s *Service) WarmUpPaybackCache(")
	require.GreaterOrEqual(t, warmStart, 0, "WarmUpPaybackCache must exist")
	warmBody := s[warmStart:]

	// It must feed the LIVE cache (not a local copy) through the live lifecycle.
	require.Contains(t, warmBody, "s.paybackCache.PrepareForBlock(",
		"WarmUpPaybackCache must PrepareForBlock on the live s.paybackCache")
	require.Contains(t, warmBody, "s.paybackCache.AddTransaction(",
		"WarmUpPaybackCache must replay txs into the live s.paybackCache via AddTransaction")
	require.Contains(t, warmBody, "s.paybackCache.FinishBlock()",
		"WarmUpPaybackCache must FinishBlock on the live s.paybackCache")

	// The per-block epoch must be derived exactly as the live re-execution path:
	// GetHistoryEpochState(FindBlockEpoch(b)).Epoch.
	require.Contains(t, warmBody, "s.store.FindBlockEpoch(b)",
		"WarmUpPaybackCache must derive the block epoch via FindBlockEpoch, mirroring ReexecuteBlocks")
	require.Contains(t, warmBody, "s.store.GetHistoryEpochState(",
		"WarmUpPaybackCache must resolve epoch rules via GetHistoryEpochState, mirroring ReexecuteBlocks")
	require.Contains(t, warmBody, "es.Epoch, es.Rules,",
		"WarmUpPaybackCache must PrepareForBlock with the history epoch+rules, mirroring the live processor")
}

// TestGetConsensusCallbacksPassesLiveWarmedCache pins that GetConsensusCallbacks
// passes s.paybackCache (the live cache that Start() has warmed) into the block
// processor. Combined with the warm-up tests above, this establishes the complete
// FIXED path: warmed live cache → GetConsensusCallbacks → block processor →
// GetAvailablePaybackByAddress returns the correct quotaUsed → matching AddBalance
// → matching state root.
func TestGetConsensusCallbacksPassesLiveWarmedCache(t *testing.T) {
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
			"Start() warms this cache before sealing so it carries the epoch's quotaUsed")
}

// TestReexecCacheIsLocalAndNotInstalledIntoLiveCache pins that ReexecuteBlocks
// still uses a local reexecCache that is NEVER assigned to s.paybackCache. The A1
// fix lives in WarmUpPaybackCache, NOT in ReexecuteBlocks: re-execution rebuilds
// EVM state under the already-sealed block.Root and must keep using a fresh local
// cache to avoid double-counting quota. (The warm-up is the separate, dedicated
// path that feeds the live cache.)
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
	// — is fine; we check for mutation signals only. The live-cache warm-up is
	// WarmUpPaybackCache's job, deliberately kept separate from re-execution.
	require.NotContains(t, reexecBody, "s.paybackCache = ",
		"ReexecuteBlocks must NOT assign to s.paybackCache; warm-up belongs in WarmUpPaybackCache")
	for _, mergeSignal := range []string{
		"s.paybackCache.AddTransaction",
		"s.paybackCache.PrepareForBlock",
		"s.paybackCache.SetUsedMap",
		"s.paybackCache.LoadEpoch",
		"s.paybackCache.WarmUp",
		"s.paybackCache.Replay",
	} {
		require.NotContains(t, reexecBody, mergeSignal,
			"ReexecuteBlocks must NOT write into s.paybackCache via %s; warm-up belongs in WarmUpPaybackCache",
			mergeSignal)
	}
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
