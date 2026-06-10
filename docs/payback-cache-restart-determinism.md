# PaybackCache Restart Determinism — Invariant Note

**Audit refs:** A1 (HIGH), T2 (HIGH) — `reports/vinuchain-audit-2026-06-10/01-VinuChain.md`
**Status:** hypothesis "a mid-epoch restart can change consensus output" — **falsified** at the consensus layer; pinned by tests.

## The concern

`receipt.FeeRefund` is consensus-sealed state (it feeds the receipts root; a
mismatch is a full chain split). It is derived from the per-address available
payback quota, which is reduced by the accumulated `quotaUsed` held in the
**volatile, in-memory** `PaybackCache.PaybackUsedMap`
(`payback/payback_cache.go`). That map:

- is never persisted to disk,
- is reset to empty at every epoch boundary (`cleanupOldEpochsLocked`),
- is reconstructed **empty** on every service start (`gossip/service.go`).

So a node that restarts mid-epoch loses its accumulated `quotaUsed` and, in
isolation, would compute a **larger** available payback (`quotaUsed = 0`) for an
address than a peer that never restarted. If that re-derivation were allowed to
**replace** an already-sealed `FeeRefund`, validators would diverge within an
epoch. This volatility is real and is proven by
`TestPaybackCacheIsVolatileWithinEpoch`.

## Why it is nevertheless deterministic across restart

The cache's volatility does not reach sealed consensus output, because of how
the node recovers in-memory EVM state on startup:

1. **Receipts are sealed once, forward-only.** During live block processing the
   block processor computes `FeeRefund`, builds the block, and persists receipts
   exactly once (`gossip/block_processor.go`: `Finalize()` →
   `evm.SetReceipts(...)`). A sealed receipt is never recomputed for that block
   during normal forward operation.

2. **On restart the live cache is empty — but the node recovers, it does not
   re-seal.** `RecoverEVM` (`gossip/c_block_callbacks.go`, called from
   `Service` start) walks back to the most recent block whose EVM **state root
   is already persisted** (`HasStateDB(block.Root)`) and re-executes **only the
   trailing, not-yet-persisted blocks** via `ReexecuteBlocks`.

3. **Re-execution uses a dedicated fresh cache and does not re-seal receipts.**
   `ReexecuteBlocks` builds a separate `reexecCache` (so live warm state cannot
   leak in), re-derives forward, commits the trailing state under each block's
   **already-sealed `block.Root`** (it discards the recomputed root), and
   **never calls `SetReceipts`.** Its sole product is a warm in-memory state
   trie so the node can resume forward sealing. The already-sealed `FeeRefund`s
   of the current epoch are therefore never overwritten by an empty-cache
   re-derivation.

4. **The fresh cache only affects NEW consensus output.** After recovery
   completes, every validator — restarted or not — seals the same
   not-yet-sealed blocks from a cache whose state derives deterministically from
   the same persisted chain prefix. At an epoch boundary all caches converge
   exactly (`TestPaybackUsedMapResetsAtEpochBoundary`).

## The load-bearing invariant

> **Re-derivation with a fresh `PaybackCache` may never replace an
> already-sealed receipt.** Only the live, forward-only block processor seals
> `FeeRefund` (via `SetReceipts`); the restart/recovery path re-derives state
> for warm-up only and must never call `SetReceipts`.

If a future change violates this — e.g. `ReexecuteBlocks` gains a `SetReceipts`
call, or the live block path re-derives already-sealed receipts from a fresh
cache — the volatile cache would become consensus-divergent on restart. **At
that point the correct fix is to persist/rebuild `PaybackUsedMap` across
restart, not to silence the tests.**

## Tests pinning this invariant

- `payback/payback_restart_determinism_test.go`
  - `TestPaybackCacheIsVolatileWithinEpoch` — proves the cache IS volatile
    mid-epoch (the premise is real, not hand-waved away).
  - `TestPaybackUsedMapResetsAtEpochBoundary` — proves warm/fresh caches
    converge across an epoch boundary.
- `gossip/payback_restart_recovery_test.go`
  - `TestReexecutionUsesFreshCacheAndDoesNotReseal` — pins that re-execution
    uses a fresh cache, commits under the sealed root, and does **not** call
    `SetReceipts`.
  - `TestRecoverEVMOnlyReexecutesTrailingUnpersistedBlocks` — pins that recovery
    is bounded to trailing not-yet-persisted blocks.
  - `TestLiveBlockPathSealsReceiptsExactlyOnce` — pins the live path as the sole
    receipt sealer.
  - `TestServiceConstructsEmptyPaybackCacheOnStart` — documents that the live
    cache is intentionally empty on start (no warm-up assumed).

## Residual risk / owner action

The argument above is a **structural proof**, validated by source-pin tests, not
a full end-to-end fork test that boots two nodes, restarts one mid-epoch, and
byte-compares receipts roots. That end-to-end test (in `integration/` or a
fakenet harness) remains the strongest possible confirmation and is recommended
as a follow-up (audit Task 4 / Milestone M0 fakenet smoke). It was deferred here
because it requires a fully-wired multi-node harness beyond the scope of the
unit-test deliverable. Maintainers who can run that harness should treat it as
the final sign-off on A1.
