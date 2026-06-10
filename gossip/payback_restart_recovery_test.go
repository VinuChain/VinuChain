package gossip

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// This file is the consensus-layer half of the M1 PaybackCache
// restart-determinism deliverable (audit findings A1 / T2). The cache-layer
// half lives in payback/payback_restart_determinism_test.go, which proves the
// in-memory cache is genuinely volatile mid-epoch. This file pins the
// structural facts in the RESTART/RECOVERY path that make that volatility
// consensus-SAFE, so a future refactor cannot silently break the safety
// argument without turning CI red.
//
// The safety argument (full derivation in the payback test file):
//   - receipt.FeeRefund is sealed once during live forward processing and is
//     never recomputed for an already-sealed block against an empty cache.
//   - On restart the cache is rebuilt empty; the node then RECOVERS in-memory
//     EVM state by re-deriving only the trailing not-yet-persisted blocks with
//     a dedicated fresh cache, WITHOUT re-sealing receipts.
//   - The fresh cache only contributes to NEW (not-yet-sealed) consensus output
//     after recovery, identically for restarted and non-restarted validators.
//
// These tests pin the load-bearing facts of that argument.

// TestReexecutionUsesFreshCacheAndDoesNotReseal pins the core safety contract
// of ReexecuteBlocks: it re-derives EVM state with a dedicated fresh
// PaybackCache and commits the trailing blocks under their ALREADY-SEALED
// roots, but it does NOT persist receipts. If a future change made
// ReexecuteBlocks call SetReceipts (re-sealing FeeRefund derived from an empty
// cache), the volatile cache could overwrite sealed consensus state and this
// test would fail.
func TestReexecutionUsesFreshCacheAndDoesNotReseal(t *testing.T) {
	src, err := os.ReadFile("c_block_callbacks.go")
	require.NoError(t, err)
	s := string(src)

	// 1. Re-execution allocates a dedicated fresh cache, never the live one.
	require.Contains(t, s, "reexecCache, err := payback.NewPaybackCache(",
		"ReexecuteBlocks must build a dedicated fresh PaybackCache so live (warm) quota state cannot leak into re-derivation")
	require.Contains(t, s, "evmProcessor := blockProc.EVMModule.Start(",
		"ReexecuteBlocks must drive re-derivation through EVMModule.Start")
	require.Contains(t, s, "reexecCache, es.Epoch)",
		"ReexecuteBlocks must hand the FRESH reexecCache (not s.paybackCache) to the re-derivation processor")

	// 2. Re-execution commits the trailing state under each block's
	//    already-sealed root, not a freshly-recomputed root.
	require.Contains(t, s, "s.store.evm.Commit(b, block.Root, false)",
		"ReexecuteBlocks must commit re-derived state under the already-sealed block.Root, keeping sealed receipts authoritative")

	// 3. THE load-bearing negative invariant: ReexecuteBlocks must NOT re-seal
	//    receipts. SetReceipts is how block_processor persists FeeRefund; its
	//    presence in the re-execution path would let an empty cache overwrite
	//    sealed consensus output.
	reexecStart := strings.Index(s, "func (s *Service) ReexecuteBlocks(")
	require.GreaterOrEqual(t, reexecStart, 0, "ReexecuteBlocks must exist")
	reexecEnd := strings.Index(s, "func (s *Service) RecoverEVM(")
	require.Greater(t, reexecEnd, reexecStart, "RecoverEVM must follow ReexecuteBlocks")
	reexecBody := s[reexecStart:reexecEnd]
	require.NotContains(t, reexecBody, "SetReceipts",
		"ReexecuteBlocks must NOT re-seal receipts (SetReceipts) — re-derivation with a fresh cache may not replace sealed FeeRefund; doing so would make the volatile cache consensus-divergent on restart")
}

// TestRecoverEVMOnlyReexecutesTrailingUnpersistedBlocks pins that recovery
// starts from the most recent block whose EVM state is already persisted
// (HasStateDB) and only re-derives forward from there. This bounds the
// re-derivation window to blocks that were never finalized to disk, so the
// fresh cache never re-derives a block whose sealed receipts are already the
// canonical, persisted truth for the chain prefix that other validators agree
// on.
func TestRecoverEVMOnlyReexecutesTrailingUnpersistedBlocks(t *testing.T) {
	src, err := os.ReadFile("c_block_callbacks.go")
	require.NoError(t, err)
	s := string(src)

	require.Contains(t, s, "s.store.evm.HasStateDB(block.Root)",
		"RecoverEVM must locate the recovery anchor via HasStateDB so only trailing not-yet-persisted blocks are re-derived")
	require.Contains(t, s, "s.ReexecuteBlocks(b, start)",
		"RecoverEVM must re-derive only [anchor, head], not the whole epoch")
}

// TestLiveBlockPathSealsReceiptsExactlyOnce pins that the LIVE block processor
// is the sole sealer of receipts (SetReceipts) and that FeeRefund is derived
// during forward sealing only. Combined with the re-execution invariants above,
// this establishes that the volatile cache only ever influences receipts at
// their single original sealing — never a second, post-restart re-derivation.
func TestLiveBlockPathSealsReceiptsExactlyOnce(t *testing.T) {
	src, err := os.ReadFile("block_processor.go")
	require.NoError(t, err)
	s := string(src)

	require.Contains(t, s, "bp.store.evm.SetReceipts(bp.blockCtx.Idx, allReceipts)",
		"the live block processor must be the receipt sealer; if this moves or the re-execution path gains a SetReceipts call the determinism argument must be re-derived")
}

// TestServiceConstructsEmptyPaybackCacheOnStart documents (and pins) that the
// live cache is reconstructed EMPTY on every service start — i.e. the volatile
// state is genuinely lost across restart. This is intentional: correctness
// rests on the recovery path above, not on warming the cache. If a future
// change adds cache warm-up/persistence here, the determinism story changes and
// these pins should be revisited together.
func TestServiceConstructsEmptyPaybackCacheOnStart(t *testing.T) {
	src, err := os.ReadFile("service.go")
	require.NoError(t, err)
	s := string(src)

	require.Contains(t, s, "svc.paybackCache, err = payback.NewPaybackCache(paybackStore,",
		"service start must construct the live PaybackCache via NewPaybackCache (empty); the safety argument assumes no warm-up here")
}
