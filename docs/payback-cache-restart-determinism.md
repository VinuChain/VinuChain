# PaybackCache Restart Determinism — Analysis & Confirmed Bug

**Audit refs:** A1 (HIGH), T2 (HIGH) — `reports/vinuchain-audit-2026-06-10/01-VinuChain.md`
**Status:** hypothesis "a mid-epoch restart can change consensus output" — **CONFIRMED as a real consensus bug.**

> **Revision note.** An earlier version of this document (commit 92abc8a) incorrectly
> concluded the hypothesis was falsified. It proved correctly that already-sealed
> receipts are never recomputed post-restart. It missed that the confirmed bug lies
> in **forward sealing of new blocks** after restart, not in re-sealing old ones.
> This document supersedes that analysis.

---

## The concern (from audit A1)

`receipt.FeeRefund` is consensus-sealed state. It is derived from the per-address
available payback quota, which subtracts the accumulated `quotaUsed` held in the
**volatile, in-memory** `PaybackCache.PaybackUsedMap`
(`payback/payback_cache.go`). That map is never persisted to disk and is reset to
empty at every epoch boundary (`cleanupOldEpochsLocked`).

The question: does a mid-epoch restart cause a validator to compute different
`FeeRefund` values than its non-restarted peers?

---

## Confirmed bug — exact causal chain

All file:line citations are to the `audit-impl` worktree.

### Step 1 — Live cache is constructed empty on startup

`gossip/service.go:488`:
```go
svc.paybackCache, err = payback.NewPaybackCache(paybackStore, ...)
```
`NewPaybackCache` returns a cache with empty `PaybackUsedMap` and `StakesMap`.
There is no replay of the current epoch's already-sealed blocks into this cache.

### Step 2 — RecoverEVM uses a LOCAL cache that is never installed back

`gossip/service.go:651` calls `s.RecoverEVM()`, which calls
`s.ReexecuteBlocks(b, start)` (`c_block_callbacks.go:133`).

`ReexecuteBlocks` (`c_block_callbacks.go:68-121`) builds:
```go
reexecCache, err := payback.NewPaybackCache(s.paybackCache.GetStore(), ...)
```
This `reexecCache` is a local variable. It accumulates `quotaUsed` across the
re-derived trailing blocks — but it is **never assigned to `s.paybackCache`** and
goes out of scope when `ReexecuteBlocks` returns. `s.paybackCache` remains empty.

### Step 3 — GetConsensusCallbacks passes the still-empty live cache to every block processor

`gossip/c_block_callbacks.go:61`:
```go
s.paybackCache,   // ← this is the empty live cache
```
is passed to `newBlockProcessor`. This is called once from
`engine.Bootstrap(svc.GetConsensusCallbacks())` (`cmd/opera/launcher/launcher.go:398`),
which happens **after** `RecoverEVM` completes in `Start()` — so by the time the
engine begins delivering new blocks for sealing, the live cache is still empty.

### Step 4 — Empty cache produces excess FeeRefund in forward sealing

For the first new (not-yet-sealed) block delivered after recovery:

`evmcore/state_processor.go:155`:
```go
availablePayback := paybackCache.GetAvailablePaybackByAddress(msg.From(), evm)
```

`GetAvailablePaybackByAddress` returns `paybackSum - quotaUsed`
(`payback/payback_cache.go:561`).

With a **warm** cache (non-restarted node): `quotaUsed = X` (accumulated over
intra-epoch blocks) → `availablePayback = paybackSum - X`.

With an **empty** cache (restarted node): `quotaUsed = 0` →
`availablePayback = paybackSum` (full epoch quota, unreduced).

### Step 5 — The excess availablePayback mutates the EVM state trie

`evmcore/state_transition.go:488-493`:
```go
if feeRefund.Sign() > 0 {
    st.feeRefund = feeRefund
    remaining = remaining.Add(remaining, feeRefund)
}
st.state.AddBalance(st.msg.From(), remaining)   // ← state trie mutation
```

A larger `availablePayback` → larger `feeRefund` → larger `remaining` →
`AddBalance` credits **more wei** to the sender's account.

### Step 6 — The state trie difference is consensus-sealed

`gossip/blockproc/evmmodule/evm.go:190`:
```go
newStateHash, err := p.statedb.Commit(true)
```
All `AddBalance` mutations are hashed into `newStateHash` → `evmBlock.Root` →
`block.Root` (`gossip/block_processor.go:610`).

`block.Root` is consensus-sealed. A different `AddBalance` → different
`block.Root` → **receipts-root mismatch → full consensus split** (per CLAUDE.md).

### Secondary effect — validator fee accounting also diverges

`gossip/blockproc/drivermodule/driver_txs.go:184-190`:
```go
feeRefund := r.FeeRefund
validatorFee := new(big.Int).Sub(txFee, feeRefund)
```
A larger `FeeRefund` → smaller `validatorFee` → different
`ValidatorStates[originatorIdx].Originated` → different epoch-end SFC staking
reward distribution. This is also consensus state.

---

## What is NOT broken (still correct from earlier analysis)

The recovery path (`ReexecuteBlocks`) never re-seals already-committed receipts.
It does not call `SetReceipts` and commits trailing state under the already-sealed
`block.Root`. Already-sealed `FeeRefund` values from the current epoch are
therefore never overwritten. The bug is exclusively in **forward sealing** of
new blocks after restart.

At epoch boundaries all caches (warm or fresh) converge to zero
(`cleanupOldEpochsLocked`), so a restart that lands on or immediately after an
epoch boundary does not diverge.

---

## Blast radius

| Dimension | Assessment |
|---|---|
| **Scope** | State-root-affecting (not RPC-only) |
| **Mechanism** | `AddBalance` before `statedb.Commit` → enters Merkle trie → `block.Root` |
| **Consequence** | Chain split: restarted validator seals a different `block.Root` than peers |
| **Preconditions** | (a) `Upgrades.Podgorica` active (FeeRefund enabled); (b) validator restarts mid-epoch; (c) an address that accumulated FeeRefund earlier in the epoch submits another tx post-restart |
| **Epoch boundary** | No divergence if restart lands at or after an epoch boundary (PaybackUsedMap resets to zero for all nodes) |
| **Practical frequency** | Mid-epoch restarts occur on every upgrade, crash recovery, or routine maintenance; condition (c) is met whenever any staking address transacts |

---

## Required fix (owner action — NOT implemented here)

Before the first new block is sealed after startup, `s.paybackCache` must reflect
the accumulated `quotaUsed` of the current epoch. Two viable approaches:

**Option A — Replay on startup (no schema change):**
After `RecoverEVM` returns, iterate the already-sealed blocks of the current
epoch (from epoch-start block to `s.store.GetLatestBlockIndex()`), retrieve each
block's transactions and receipts, and call `s.paybackCache.AddTransaction(tx,
receipt)` for each. This warm-up must complete before `engine.Bootstrap` delivers
the first new block for sealing.

**Option B — Persist/reload PaybackUsedMap (schema change):**
Persist `PaybackUsedMap` to the gossip store at the end of each block (or epoch)
and reload it on startup. More robust, higher storage overhead.

Either fix closes A1. **Do not implement unilaterally on a live mainnet without
a coordinated upgrade and thorough testing** — both options touch the consensus
path.

---

## Tests pinning this finding

**`payback/payback_restart_determinism_test.go`**
- `TestForwardSealingDivergesAfterMidEpochRestart` — characterisation test for
  the confirmed bug. Asserts that a fresh cache and a warm cache return different
  `quotaUsed` for the same mid-epoch block, pinning the current broken behaviour.
  This test **must be deleted or rewritten once the fix is applied.**
- `TestPaybackCacheIsVolatileWithinEpoch` — proves the cache IS volatile
  mid-epoch (the premise is real).
- `TestPaybackUsedMapResetsAtEpochBoundary` — proves convergence at epoch
  boundaries (bounds the blast radius to within one epoch).

**`gossip/payback_restart_recovery_test.go`**
- Group A pins: `TestServiceCacheIsEmptyAtStartAndHasNoWarmUpPath`,
  `TestReexecCacheIsLocalAndNotInstalledIntoLiveCache`,
  `TestGetConsensusCallbacksPassesLiveCache` — the three structural facts that
  together constitute the bug path.
- Group B pins: `TestReexecutionDoesNotReseal`,
  `TestRecoverEVMOnlyReexecutesTrailingUnpersistedBlocks`,
  `TestLiveBlockPathSealsReceiptsExactlyOnce` — facts that remain correct;
  a refactor that makes `ReexecuteBlocks` call `SetReceipts` would compound the bug.
