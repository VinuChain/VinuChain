# Mainnet Hard-Fork Release Runbook

This is the protocol-level, publishable procedure for activating a staged
hard-fork on VinuChain mainnet (chain 207). It captures the order-dependent
prerequisites and verification steps so the upgrade is repeatable and external
validators know what to expect.

Box-specific execution detail (instance access, snapshot storage paths, key
custody) deliberately lives in the internal deployment-log and is **not**
reproduced here. This document is self-sufficient for the protocol-level
procedure; the internal log is required only for the operator running the boxes.

## Scope & current state

- **Mainnet (chain 207)** still runs `v2.0.0-rc.1`, which pre-dates the
  Podgorica / SfcV2 / Elemont era. The flags are already **staged in code** --
  `VinuChainMainNetRules()` (`opera/rules.go`) carries the full upgrade flag set
  (`Berlin`, `London`, `Shanghai`, `Cancun`, `Prague`, `VinuBLS12381`,
  `VinuLatestEVM`, `Llr`, `Podgorica`, `SfcV2`, `Elemont`,
  `ElemontPubkeyValidation`).
- **Scope decision 2026-08-19: FULL PARITY.** The 2026-08-29 10:00 UTC release
  activates the entire testnet feature set in one window, including `PaybackV2`.
  This supersedes the two-release sequencing previously recommended below.
- **The one remaining code gap is `PaybackV2`.** It cannot be flipped until
  `QuotaContractV2` is deployed on mainnet and both `paybackV2MainnetAddress` and
  `paybackV2StagingAddress` are baked into `opera/payback_v2_address.go` --
  `EnforcePaybackV2StartupCheck()` panics at process init otherwise, on every
  network. See `vinuchain-ops-docs/ops/paybackv2-mainnet-deploy-runbook.md`.
  Until then `Economy.QuotaCacheAddress` stays pinned to the live V1 Quota proxy
  `0x1c4269fbbd4a8254f69383eef6af720bcd0acda6`.
- **Testnet (chain 206)** has already activated the full flag set plus PaybackV2
  and serves as the dress rehearsal for every step below. Treat a clean testnet
  rollout as the precondition for starting the mainnet window.
- **PaybackV2 is intentionally NOT yet enabled on mainnet rules** -- it ships as a
  separate, later release (see [Release sequencing](#release-sequencing)).

## The core hazard: stale-genesis replay divergence

The distributed mainnet genesis (May 2024) pre-dates every Elemont-era flag.
When a new mainnet binary first boots, **all staged flags fire at the first epoch
seal after boot**. This is safe for a node continuing from existing chaindata,
but it is a **chain-splitting trap for any node replaying from the 2024
genesis**:

- A fresh-install node replaying from the 2024 genesis seals each flag at a
  *different* first-replay block than the live chain did.
- Because `receipt.FeeRefund` (introduced by Podgorica) is a **persisted receipt
  field**, this difference produces a receipts-root mismatch. That is a full
  consensus split, surfacing on the diverging node as a
  `wrong event epoch hash` error -- the exact failure that bit testnet operators
  before.
- The multi-fork staging is **sequential, not simultaneous**: Cancun only stages
  after Shanghai is *active*, and Prague only after Cancun. The EVM forks
  therefore activate across **consecutive epoch seals**, not all in one seal.
  Plan for several seals to elapse before the chain has settled into the final
  rule set.

The mitigation is the three prerequisites in the next section: produce a
post-upgrade snapshot and a regenerated genesis so that fresh installs never have
to replay the pre-fork history, and forbid fresh installs during the window.

## Prerequisites -- all three, same day as the binary (none optional)

All three belong to the **same upgrade**, but their timing differs: the operator
announcement (3) must be published **before** the window opens, while the snapshot
(1) and the regenerated genesis (2) are produced **after** the final seal. Skipping
or mis-ordering any one re-opens the stale-genesis divergence described above.

1. **Post-upgrade chaindata snapshot.** Take the snapshot **only after the RPC
   node has sealed *every* staged flag** -- not merely the first post-boot seal.
   (This is the single detail most likely to be gotten wrong.)

   **Under full parity the last fork is `VinuLatestEVM`, not Prague.** The EVM
   forks stage sequentially (`gossip/service.go`: Cancun waits for Shanghai to be
   active, Prague for Cancun, VinuBLS12381 for Prague, VinuLatestEVM for
   VinuBLS12381), so the activation crosses **five consecutive epoch seals**.
   Mainnet epochs seal at the 4h `MaxEpochDuration` cap -- measured median 240.3
   min over the twelve seals before 2026-08-19 -- so the sequence takes roughly
   20 hours from the first seal. Snapshotting after Prague (seal 3) would publish
   an artifact missing two activations and re-open exactly the divergence these
   prerequisites exist to close. Confirm from `vc_getRules` that
   `VinuLatestEVM = true` before snapshotting. When tarring the datadir, **exclude**
   `nodekey`, the keystore, the IPC socket, and the static-/trusted-nodes files,
   so the snapshot is identity-free and safe to distribute.
2. **Regenerated distributed genesis + `AllowedOperaGenesis` update.** Export a
   fresh distributed genesis from a node that has **already sealed the flags**,
   then update `cmd/opera/launcher/params.go` `AllowedOperaGenesis` with the new
   section hashes. This lets fresh installs adopt the new genesis **without**
   needing `--genesis.allowExperimental`.
3. **Operator announcement.** Publish an explicit rule: **no fresh validator
   installs during, or within 24 h of, the upgrade window.** In-place binary
   swaps on existing datadirs are safe. A fresh install booted from the *stale*
   genesis during the window will diverge and require a chaindata wipe plus a
   snapshot restore to recover. **Publish this announcement BEFORE the swap
   window opens** (see the Execution outline) -- not after the seal: a freeze
   announced once the window has already begun cannot stop a fresh install that
   has already started diverging.

## Release sequencing

> **SUPERSEDED 2026-08-19.** The guidance below (two releases, PaybackV2 >= 2
> weeks after the consensus flags) was the prior recommendation. The operator has
> decided to ship **full parity in a single 2026-08-29 window**, PaybackV2
> included. Keep reading for *what PaybackV2 requires* -- every prerequisite it
> lists still applies, they just apply to the same activation day as the
> consensus flags rather than a later one. Because PaybackV2 is a
> persisted-state change, the snapshot / regenerated-genesis / no-fresh-installs
> prerequisites cover it too.
>
> Consequence worth stating plainly: combining them does compound the blast
> radius, and mainnet is jumping from `v2.0.0-rc.1` with no intermediate release,
> so **after the first seal there is no binary to roll back to** -- recovery is
> snapshot-based only.

The original two-release recommendation follows.

- **Release 1 -- consensus flags.** Ship Podgorica + SfcV2 + Elemont + the EVM
  forks (Berlin/London/Shanghai/Cancun/Prague) as one release. This is the set
  currently staged in `VinuChainMainNetRules()`.
- **Release 2 -- PaybackV2 (separate, later).** Stage PaybackV2 as its own
  release, **>= 2 weeks after** Release 1 has been live without surprises.
  PaybackV2 is itself a **persisted-state change**: at its activation seal it
  swaps `Economy.QuotaCacheAddress` to the freshly-deployed QuotaContractV2 (see
  `opera/payback_v2_address.go`). Because that is a persisted-state change, **all
  three prerequisites (snapshot-after-final-seal, regenerated genesis,
  no-fresh-installs announcement) apply again on PaybackV2's own activation day.**
  A startup check (`EnforcePaybackV2StartupCheck`) refuses to boot a binary that
  enables PaybackV2 while the matching network's V2 address is still the zero
  sentinel, so the address must be deployed and recorded before that release
  ships.

## Execution outline

Protocol-level steps. Box-specific steps -- instance access, snapshot upload/
download paths, and operator key custody -- are in the internal deployment-log and
are not reproduced here.

1. **Publish the fresh-install freeze first (prerequisite 3).** BEFORE any binary
   swap, announce and put into effect the **no fresh validator installs during, or
   within 24 h of, the upgrade window** rule. This must lead the window: a fresh
   install that boots from the stale 2024 genesis during the swap/sealing window
   diverges *before* any later-published warning could reach operators, so a freeze
   announced after the seal is too late to prevent the divergence it exists to stop.
2. **Pre-build off-box.** Build the new binary on a dedicated build host, never on
   a production validator/RPC box during the window.
3. **Cross-verify the binary.** Compute the `sha256` of the binary on **more than
   one build host** and confirm the hashes match before trusting it. Distribute
   only a hash-verified binary.
4. **Swap and restart in place.** On each box, replace the binary on the existing
   datadir and restart. In-place swaps on existing chaindata are the safe path.
5. **Never SIGKILL a validator.** A hard kill risks LevelDB corruption. Stop nodes
   only with a clean `SIGINT` / `systemctl restart`.
6. **Wait for the final seal, then snapshot.** Let the RPC node seal through to
   Prague (see prerequisite 1), then take the post-upgrade snapshot and regenerate
   the distributed genesis (prerequisites 1 and 2). The fresh-install freeze from
   step 1 stays in effect until at least 24 h after the final seal.

## Post-activation follow-up: the stale-datadir guard (do NOT do this before the seals)

`checkStoredChainState` (`cmd/opera/launcher/config.go`) refuses to boot a datadir of a known
public-network genesis lineage whose persisted activation history disagrees with the live chain's.
Testnet has an entry pinning `SfcV2Patch7/8/9` at epochs 6017/6118/6119 — the guard that would
have prevented testnet validators 17 and 18 forking on 2026-06-21.

**Mainnet deliberately has no entry, and must not get one before the activation.**
`StoredStateRequirements` in `cmd/opera/launcher/params.go` says why:

> Mainnet and staging carry no requirement: their ELEMONT-era activations have not rolled out
> yet, so a pre-activation datadir is legitimately below every seal. Add an entry for those
> lineages as part of the mainnet upgrade rollout, once the activation epochs are historical
> fact -- **never at the release that first stages them, which would refuse the whole pre-seal
> fleet.**

Adding it to the 2026-08-29 release would brick every mainnet node on boot. So this is a
**follow-up release** task, after the fifth seal:

1. Record the actual activation epoch and block for each flag as it seals -- `SfcV2`, `Elemont`,
   `ElemontPubkeyValidation`, `Shanghai`, `PaybackV2` (seal 1), then `Cancun`, `Prague`,
   `VinuBLS12381`, `VinuLatestEVM` (seals 2-5). Read them from `vc_getRules` and the seal-time
   logs, not from projections.
2. Add a `StoredStateRequirement` for `vinuChainMainnetHeader.GenesisID` with those activations,
   and set `Bootstrap` to the regenerated post-activation mainnet genesis URL.
3. Add the matching entry for the staging lineage (`vinuChainTestMainnetHeader.GenesisID`) if
   staging is run against the same activations.
4. Set `SupersededBy` on the two existing mainnet presets ("VinuChain mainnet without history",
   "VinuChain mainnet with deployed contracts") to the regenerated genesis URL, so fresh installs
   are pushed off the pre-activation genesis instead of replaying it and diverging.

Until that follow-up ships, mainnet has **no** automatic protection against a stale-datadir or
stale-genesis boot -- which is exactly why the fresh-install freeze (prerequisite 3) is an
operational rule for this window rather than a code guard.

## Post-activation verification checklist

After the seal, confirm the upgrade took using only the **public RPC**. These are
copy-pasteable so anyone can independently verify.

1. **Rules reflect the staged flags and address.** Query `vc_getRules` and confirm
   the expected `Upgrades.*` flags are `true` and `Economy.QuotaCacheAddress`
   matches the intended value.

   ```sh
   curl -s -X POST https://rpc.vinuchain.org \
     -H 'Content-Type: application/json' \
     -d '{"jsonrpc":"2.0","id":1,"method":"vc_getRules","params":["latest"]}' | jq .result.Upgrades
   ```

   Expected (Release 1): `Podgorica`, `SfcV2`, `Elemont`, `Shanghai`, `Cancun`,
   and `Prague` all `true`. Then confirm the address:

   ```sh
   curl -s -X POST https://rpc.vinuchain.org \
     -H 'Content-Type: application/json' \
     -d '{"jsonrpc":"2.0","id":1,"method":"vc_getRules","params":["latest"]}' \
     | jq -r .result.Economy.QuotaCacheAddress
   ```

   Expected: `0x1c4269fbbd4a8254f69383eef6af720bcd0acda6` (the live mainnet Quota
   proxy pinned in `VinuChainMainNetRules()`). After Release 2 this value changes
   to the recorded QuotaContractV2 address.

2. **SFC bytecode matches the intended V2 cycle.** Fetch the deployed SFC bytecode
   via the public RPC and confirm it matches the intended V2 cycle build.

   ```sh
   curl -s -X POST https://rpc.vinuchain.org \
     -H 'Content-Type: application/json' \
     -d '{"jsonrpc":"2.0","id":1,"method":"eth_getCode","params":["0xFC00FACE00000000000000000000000000000000","latest"]}' | jq -r .result
   ```

3. **(Release 2 / PaybackV2 only) `feeRefundBlockCount()` returns the expected
   value.** `eth_call` `feeRefundBlockCount()` on the new contract and confirm the
   returned value matches the intended configuration.

   ```sh
   # 0x<selector> = keccak256("feeRefundBlockCount()")[:4]; to = QuotaContractV2 address
   curl -s -X POST https://rpc.vinuchain.org \
     -H 'Content-Type: application/json' \
     -d '{"jsonrpc":"2.0","id":1,"method":"eth_call","params":[{"to":"0x<QuotaContractV2>","data":"0x<selector>"},"latest"]}' | jq -r .result
   ```

All three (1 and 2 for Release 1; add 3 for Release 2) must pass before declaring
the upgrade complete.

## Rollback / divergence response

- **An existing-datadir node fails to keep up.** In-place binary swaps on existing
  chaindata are unaffected by the stale-genesis hazard. If such a node misbehaves,
  treat it as an ordinary node-health incident (clean restart, resync from a
  healthy peer); it is not a genesis-divergence case.
- **A fresh-install node shows `wrong event epoch hash`.** This is the
  stale-genesis divergence. Recover by **wiping its chaindata and restoring from
  the post-seal snapshot** (the one taken after Prague sealed). Do **not** attempt
  to replay it forward from the 2024 genesis -- that is what caused the divergence.
- **Prevention beats rollback.** The no-fresh-installs-in-the-window announcement
  (prerequisite 3) exists precisely to avoid this state. Honor it.

## References

- `opera/rules.go` -- `VinuChainMainNetRules()`: the staged mainnet flag set and
  the pinned `Economy.QuotaCacheAddress`.
- `opera/payback_v2_address.go` -- per-network PaybackV2 / QuotaContractV2 address
  slots, the zero-sentinel convention, and the `EnforcePaybackV2StartupCheck`
  boot guard.
- `cmd/opera/launcher/params.go` -- `AllowedOperaGenesis`, updated with the
  regenerated genesis section hashes (prerequisite 2).
- Internal deployment-log -- box-specific execution detail (instance access,
  snapshot storage, key custody). Not part of this repository's public tree;
  available to the operator only.
