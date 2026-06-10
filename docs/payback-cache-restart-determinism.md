# PaybackCache Restart Determinism — Analysis, Confirmed Bug & Fix

**Audit refs:** A1 (HIGH), T2 (HIGH) — `reports/vinuchain-audit-2026-06-10/01-VinuChain.md`
**Status:** hypothesis "a mid-epoch restart can change consensus output" — **CONFIRMED, then RESOLVED (Option A — startup replay, no schema change).**

> **Revision note.** An earlier version of this document (commit 92abc8a) incorrectly
> concluded the hypothesis was falsified. It proved correctly that already-sealed
> receipts are never recomputed post-restart. It missed that the confirmed bug lies
> in **forward sealing of new blocks** after restart, not in re-sealing old ones.
> This document supersedes that analysis.
>
> **Fix note.** The forward-sealing divergence described below is now fixed by a
> startup cache warm-up (`gossip/service.go` → `gossip/c_block_callbacks.go`
> `WarmUpPaybackCache`). See **Resolution** at the bottom. The fix is committed on
> branch `audit-impl` but **not pushed**; it requires owner fakenet/testnet
> validation and a coordinated upgrade before any deploy.

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

## Fix (Option A — replay on startup, IMPLEMENTED)

Before the first new block is sealed after startup, `s.paybackCache` must reflect
the accumulated `quotaUsed` of the current epoch. Two approaches were considered:

**Option A — Replay on startup (no schema change) — CHOSEN & IMPLEMENTED.**
After `RecoverEVM` returns and before `engine.Bootstrap` delivers the first new
block for sealing, iterate the already-sealed blocks of the current epoch,
retrieve each block's transactions and receipts, and replay them through
`s.paybackCache.AddTransaction(tx, receipt)` — driven by the same
`PrepareForBlock`/`FinishBlock` lifecycle the live block processor uses.

**Option B — Persist/reload PaybackUsedMap (schema change):** rejected for this
pass (higher storage overhead and a DB-format change). Option A is sufficient,
self-correcting at epoch boundaries, and touches no on-disk schema.

**Do not deploy unilaterally on a live mainnet without a coordinated upgrade and
fakenet/testnet validation** — this is consensus-path code.

---

## Resolution

**Status: RESOLVED on branch `audit-impl` (committed, not pushed).**

### The seam

`gossip/service.go` `Service.Start()`:

```go
s.RecoverEVM()
s.WarmUpPaybackCache()   // ← added: A1 fix
root := s.store.GetBlockState().FinalizedStateRoot
```

`WarmUpPaybackCache` runs **after** `RecoverEVM` (so trailing EVM state exists)
and **before** `s.blockProcTasks.Start(1)` and the engine's first delivered block
(via `engine.Bootstrap(svc.GetConsensusCallbacks())` in `cmd/opera/launcher`).

### The warm-up (`gossip/c_block_callbacks.go` `WarmUpPaybackCache`)

The volatile cache carries **two** consensus-relevant maps, not one:

| Map | Lifetime at head (epoch E) | Consensus-relevant epochs |
|---|---|---|
| `PaybackUsedMap` | epoch **E only** (reset every boundary) | **E** |
| `StakesMap` | epochs **E, E-1, E-2** (`cleanupOldEpochsLocked` deletes only `< E-2`) | **E and E-1** (E-2 retained but never read by the duration math) |

`StakesMap[E-1]` is consensus-relevant because `calculateFullDurationLocked`
(`payback/payback_cache.go`) branches on `getSumStakeByAddressSplitLocked(addr, E, E-1)`:
an address with `sumCurrent>0 && sumPrev==0` uses time-since-last-stake; otherwise it
uses `prevEpochState.Duration()`. An address that staked in BOTH E-1 and E therefore
takes a **different duration branch** depending on whether `StakesMap[E-1]` is
populated → different `fullDuration` → different `paybackSum` → different `FeeRefund`.

Accordingly the replay starts at the **first block of epoch E-1**, not epoch E:
`GetHistoryBlockEpochState(E-1).LastBlock.Idx + 1`.

> **Epoch-index convention (and the refuted finding P1).** The audit rework finding
> P1 claimed `GetHistoryBlockEpochState(E-1).LastBlock.Idx + 1` is the *first block
> of epoch E* (so the original warm-up would leave `StakesMap[E-1]` empty). This was
> **REFUTED**. `GetHistoryBlockEpochState(N)` is written when epoch N is sealed
> (`gossip/block_processor.go:739`, alongside `SetEpochBlock(idx+1, N)` at line 740),
> with `LastBlock.Idx` = the last block of the `FindBlockEpoch==N-1` group.
> Therefore `GetHistoryBlockEpochState(N).LastBlock.Idx + 1` == the **first block of
> the `FindBlockEpoch==N` group**. With `currentEpoch == E == FindBlockEpoch(head)`,
> the original `GetHistoryBlockEpochState(E-1).LastBlock.Idx + 1` is the **first
> block of epoch E-1** — so the original fix **already** covered `StakesMap[E-1]`.
> This was verified empirically (the lookup returns exactly the first block whose
> `FindBlockEpoch == E-1`, and forcing a *true* epoch-E-only start leaves
> `StakesMap[E-1]` empty, failing the StakesMap pin). `StakesMap[E-2]` is retained by
> `cleanupOldEpochsLocked` but never read by `getSumStakeByAddressSplitLocked`
> (which only sums epochs E and E-1), so it is consensus-irrelevant and is not
> reconstructed.

For every already-sealed block `b` in `[firstBlock(E-1), GetLatestBlockIndex()]`:

1. Derive the block's epoch and rules **exactly as `ReexecuteBlocks` does**:
   `es := GetHistoryEpochState(FindBlockEpoch(b))`.
2. `s.paybackCache.PrepareForBlock(es.Epoch, es.Rules, block.Time.Time())` —
   identical to `OperaEVMProcessor.Execute` (`evmmodule/evm.go:137`).
3. For each `(tx, receipt)` from `GetBlockTxs(b, block)` + the **raw** stored
   receipts (see "No receipts-LRU pollution" below): cache the sender for
   non-internal txs via `types.Sender(signer, tx)` (mirroring the live
   `tx.AsMessage` path) — internal txs are left uncached so `AddTransaction`'s
   `tx.From()==zero` guard drops them, exactly as live — then
   `s.paybackCache.AddTransaction(tx, receipt)`.
4. `s.paybackCache.FinishBlock()`.

`PrepareForBlock`'s `cleanupOldEpochsLocked` zeroes `PaybackUsedMap` at the
E-1→E boundary **during the replay** while `StakesMap[E-1]` survives (cutoff is
`E-2`). The warmed cache therefore ends holding **exactly** what a never-restarted
node holds at head: `PaybackUsedMap` for **E only**, `StakesMap` for **E-1 and E**.

### Why this matches the live derivation bit-for-bit

The warmed cache equals a never-restarted node across **both** maps and the
relevant epochs:

- **`PaybackUsedMap[E]`** — same accumulation function: `AddTransaction` is the
  **sole** writer on the live path (`evmcore/state_processor.go`), gated on the
  same `receipt.Status == Successful` and `receipt.FeeRefund > 0` conditions, with
  the same `min(feeRefund, txFee)` cap.
- **`StakesMap[E-1]` and `StakesMap[E]`** — replaying from the first block of E-1
  records the same stake() / stakeFor() entries the live processor recorded as
  those blocks were originally sealed, keyed by the same epoch; the E-1→E boundary
  cleanup preserves E-1 exactly as live.
- Same epoch keying: `GetHistoryEpochState(FindBlockEpoch(b)).Epoch`, the existing
  re-execution precedent.
- Same sender semantics: non-internal txs recover the **same** EIP-155 sender
  address regardless of signer variant; internal txs are excluded on both paths.

### No receipts-LRU pollution (no RPC corruption)

The warm-up reads receipts via `evm.GetRawReceipts(b)` and locally projects each
into the minimal fields `AddTransaction` needs — `Status`, `FeeRefund`, and per-tx
`GasUsed` (derived from the stored `CumulativeGasUsed` deltas exactly as
`Receipts.DeriveFields` computes them). It deliberately does **not** call
`evmstore.GetReceipts`: that path runs `DeriveFields` with a zero `BlockHash` and
inserts the result into the shared receipts LRU keyed by block number. Because
`EthAPIBackend.GetReceiptsByNumber` serves from that LRU, the old warm-up would
have made every RPC query for a current-epoch block return `blockHash 0x0`
receipts/logs after each restart until eviction. The raw-receipt path touches no
shared cache, so RPC output is unaffected.

### TxIndex=false: fail closed (no silent divergence)

Receipts are only persisted when `bp.txIndex` is true
(`gossip/block_processor.go`). On a node running with `TxIndex` disabled, the
trailing receipts do not exist. The pre-rework warm-up called
`evmstore.GetReceipts`, whose `DeriveFields` errored ("transaction and receipt
count mismatch") and tripped `s.Log.Crit` with a confusing message — and, on the
way, polluted the RPC receipts LRU.

The reworked warm-up reads **raw** receipts (no LRU pollution, no `DeriveFields`),
so a normally-configured (TxIndex-enabled) node never hits a spurious crash. When
a tx-bearing block's receipts genuinely cannot be read, the payback cache would be
reconstructed **incompletely**, and sealing forward from it would compute a
different FeeRefund / `block.Root` than peers. The warm-up therefore **fails
closed**: it calls `s.Log.Crit` with an actionable message naming the TxIndex
requirement, refusing to start rather than seal divergent state. (An earlier
rework draft warned-and-continued here; that was rejected in review because a
validator must not proceed with incomplete consensus material. Fail-closed is
strictly safer and is what ships.)

> **Operational requirement:** validators MUST run with `TxIndex` enabled. With
> `TxIndex=false` the payback cache cannot be warmed from persisted receipts;
> the node now refuses to start (fail closed) instead of diverging on the
> FeeRefund/StakesMap consensus path.
>
> **Release-note callout:** this converts `TxIndex=false` from a tolerated
> configuration into a refuse-to-start configuration on the next restart, and
> merely enabling `TxIndex` does **not** backfill receipts for already-sealed
> blocks — a node that previously ran with `TxIndex=false` must be re-synced
> (snapshot or genesis re-import) with `TxIndex` enabled. The Crit messages say
> exactly this.

The same fail-closed policy applies to structurally missing chain data inside
the replay window — with one deliberate carve-out. Blocks absent **before the
first readable block** are a *leading gap*: on genesis-imported or pruned
nodes the window's start (the bounded fallback may select block 1) can predate
the earliest block the store holds, which is well-understood and tolerated
(skipped with a single `Warn` naming the skipped count and the first readable
block). Once replay material has started, a missing block, an unresolvable
`FindBlockEpoch`, or a missing history epoch state each `Crit` (matching
`ReexecuteBlocks`, which treats the identical conditions as fatal) instead of
silently skipping — a hole in the middle of the window would reconstruct
exactly the incomplete cache this design refuses to run with. Mid-window holes
are unreachable on a healthy node.

### Known limitation — epoch-seal rules drift (QuotaCacheAddress upgrades)

On an epoch-sealing block, the **live** processor refreshes rules mid-block
(`sealEpochIfNeeded` → `SetRules`), so txs in the post-seal portion of a block
whose seal changes `Economy.QuotaCacheAddress` are classified against the NEW
quota address. The warm-up replays each block under a single rules snapshot
(`GetHistoryEpochState(FindBlockEpoch(b)).Rules` — the pre-seal rules), so for
that one block it classifies stake/unstake txs against the OLD address. If a
restart's replay window contains such an activation block AND that block
carries quota-contract stake txs post-seal, the reconstructed `StakesMap` can
differ from a never-restarted peer's.

The trigger is narrow — at most one block per `QuotaCacheAddress`-changing
upgrade, and only when that exact block carries stake-selector txs — but
upgrade windows are precisely when validators restart, so it is handled
operationally rather than left implicit:

> **Operational rule:** do NOT restart validators within the two-epoch window
> after a `QuotaCacheAddress`-changing upgrade activates. If a restart in that
> window is unavoidable, restart from a state snapshot taken after the window
> closes, or accept that the node must be compared against a healthy peer
> before re-joining as an emitter.

Splitting the replay of a sealing block at the seal point (re-deriving the
mid-block rules switch) was considered and rejected for now: it would
re-implement `executePreInternalTxs`/seal ordering inside the warm-up, whose
divergence risk outweighs the narrow window it closes. Revisit if
`QuotaCacheAddress` upgrades become frequent.

### Deterministic fallback (pruned / genesis-imported nodes)

When the epoch-boundary history needed to locate the first block of E-1 is
unavailable (`GetHistoryBlockEpochState(E-1)` returns nil — pruned history or a
genesis-imported node), the warm-up walks `FindBlockEpoch` backwards from head to
find the true first block of epoch E-1. This scan reads only per-block epoch
mappings (no ECDSA recovery), so it is cheap and is naturally bounded by the
length of epochs E-1 and E (epochs are time- and gas-bounded). It emits a
progress `Warn` every `warmUpProgressInterval` (10000) scanned blocks so a long
scan is observable.

Crucially, the fallback **never returns a mid-epoch block**: it returns either the
true first block of E-1, or — if no block maps to E-1 at all (e.g. the whole
available history is one epoch, or `FindBlockEpoch` returns 0) — block 1. Both
choices replay the *complete* current epoch. Returning a bounded window *inside*
epoch E would leave `PaybackUsedMap[E]` missing refunds consumed earlier in the
epoch → excess FeeRefund after restart → consensus split; that incorrect
shortcut was rejected in review and is explicitly avoided. A progress log is also
emitted every 10000 blocks during the replay itself.

### Verification

- `go build ./...` — clean.
- `go vet ./...` — clean.
- `payback` + `gossip` packages — green, including `-race`.
- Convergence (cache layer, PaybackUsedMap): `payback/.../TestForwardSealingConvergesAfterMidEpochRestart`
  — a restarted-then-warmed cache reports the same per-address `quotaUsed` as a
  never-restarted cache (real signed txs + FeeRefund receipts).
- Convergence (cache layer, StakesMap[E-1]): `payback/.../TestWarmUpReconstructsStakesMapForPreviousEpoch`
  — replaying epochs E-1 and E reconstructs `StakesMap[E-1]` (an address staking
  in both E-1 and E); a cache-layer property check for the StakesMap[E-1] math.
- Convergence (end-to-end, PaybackUsedMap): `gossip/.../TestWarmUpPaybackCacheConvergesWithNeverRestartedNode`
  — seals real blocks via the live processor, injects crafted non-zero FeeRefund
  into the **raw** stored receipts (the fakenet genesis has no quota contract, so
  this avoids the earlier **vacuous** all-zero comparison), swaps in a fresh empty
  `paybackCache` (simulating restart), runs `WarmUpPaybackCache`, and asserts the
  rebuilt `PaybackUsedMap` is **non-empty** and equals the never-restarted snapshot.
- Convergence (end-to-end, StakesMap[E-1]): `gossip/.../TestWarmUpReconstructsStakesMapForPreviousEpochEndToEnd`
  — seals real `stake()` txs to the fakenet `QuotaCacheAddress` across an epoch
  boundary, simulates a restart, and asserts the warmed `StakesMap[E-1]` matches
  the never-restarted node's (fails against the old epoch-E-only warm-up).

---

## Tests pinning this finding (flipped: divergence → convergence)

**`payback/payback_restart_determinism_test.go`**
- `TestForwardSealingConvergesAfterMidEpochRestart` — **success-criterion test
  for the fix** (replaces the former `…Diverges…` characterisation test). Drives
  real signed txs + non-zero-FeeRefund receipts: a restarted cache that replays
  the epoch's already-sealed `(tx, receipt)` pairs via `AddTransaction` (exactly
  what `WarmUpPaybackCache` does) reports the **same** per-address `quotaUsed` as
  a never-restarted cache. Fails loudly if the warm-up is removed or warms to the
  wrong value.
- `TestPaybackCacheIsVolatileWithinEpoch` — proves the cache IS volatile
  mid-epoch (the premise the fix addresses is real).
- `TestPaybackUsedMapResetsAtEpochBoundary` — proves convergence at epoch
  boundaries (the warm-up's per-epoch reset relies on this).

**`gossip/payback_restart_recovery_test.go`** (structural pins, updated to the
NEW correct wiring)
- Group A (fix wiring): `TestServiceWarmsPaybackCacheBeforeForwardSealing`
  (Start() calls `WarmUpPaybackCache` after `RecoverEVM` and before the block
  processor starts), `TestWarmUpReplaysCurrentEpochIntoLiveCache` (replays into
  the live `s.paybackCache` via `PrepareForBlock`/`AddTransaction`/`FinishBlock`
  with the `GetHistoryEpochState(FindBlockEpoch(b))` epoch derivation),
  `TestGetConsensusCallbacksPassesLiveWarmedCache`.
- Group B (still-correct facts): `TestReexecCacheIsLocalAndNotInstalledIntoLiveCache`
  (re-execution keeps its own local cache; warm-up is the separate dedicated
  path), `TestReexecutionDoesNotReseal`,
  `TestRecoverEVMOnlyReexecutesTrailingUnpersistedBlocks`,
  `TestLiveBlockPathSealsReceiptsExactlyOnce`.

**`gossip/payback_restart_simulation_test.go`** (new, end-to-end)
- `TestWarmUpPaybackCacheConvergesWithNeverRestartedNode` — seals real blocks
  through the live block processor, simulates a mid-epoch restart by replacing
  `svc.paybackCache` with a fresh empty cache, runs `WarmUpPaybackCache`, and
  asserts the rebuilt `PaybackUsedMap` equals the never-restarted snapshot.
