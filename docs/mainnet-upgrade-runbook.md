# Mainnet ELEMONT Protocol Runbook

**Upgrade window:** 2026-08-29 10:00 UTC

**Network:** VinuChain mainnet, chain ID `207`

**Status checked:** 2026-08-25

This is the protocol-coordination runbook for the ELEMONT mainnet activation.
Host-specific commands belong in the public `VinuChain-Docs` Mainnet Upgrade
Guide; instance access, snapshot storage, alarms, and key custody remain in the
private operations runbook.

## Outcome

Existing validators and RPC nodes make one in-place binary swap from
`v2.0.0-rc.1` to `v2.0.49-elemont`. The upgrade is complete when:

- all five activation seals have completed through `VinuLatestEVM`;
- the SFC runs the expected Cycle-165 bytecode and reports version `305`;
- `Economy.QuotaCacheAddress` is the mainnet QuotaContractV2 address;
- upgraded nodes agree with the public RPC at the same block height; and
- an identity-free post-activation snapshot and its verified manifest are
  published for recovery and new nodes.

This release includes Payback V2. There is no second Payback activation release.
The five seals are stages of this one upgrade, not five operator restarts.

## Fixed release values

| Item | Required value |
| --- | --- |
| Release | `v2.0.49-elemont` |
| Tag commit | `8b88cc49d11e56635385413fe8f9eaec1969c1ac` |
| linux/amd64 binary | `opera-v2.0.49-elemont-linux-amd64` |
| Published SHA256 | `678040e9f88a98331a8cc32b7bf5b9e0ae4acdf84919390465eeee584b7f56c1` |
| Minimum glibc | `2.34` |
| SFC address | `0xFC00FACE00000000000000000000000000000000` |
| SFC version after seal 1 | `305` |
| SFC runtime length after seal 1 | `48,757` bytes |
| SFC runtime keccak after seal 1 | `0x29b88152209fe22bef409376aa7f137d0e0f571f46afa1385f32320765e49e50` |
| QuotaContractV2 after seal 1 | `0x5D989A2d65d049e2198D91d8ddc31C918f2544AB` |

## Before the window

Complete every gate. An unchecked gate is a no-go, not an item to improvise
during the cutover.

1. **Prove the release artifact.** Download it from the tagged GitHub release on
   each target host. Verify the published SHA256, `opera version`, tag
   commit, architecture, and glibc compatibility before stopping a node.
2. **Finish the code and state gates.** Confirm the QuotaContractV2 address is
   baked for mainnet and staging, the startup guard passes, testnet has completed
   the same feature path, and the SFC delegation backfill list has been re-derived
   against a block near activation. Record the compiled pair count separately
   from the state-dependent expected `Appended` and `Repaired` results; never
   equate list length with either seal-log field.
3. **Prove every node's effective configuration.** Following the public guide,
   record the live process user, `HOME`, executable, working directory, arguments,
   service, effective `DataDir`, `IPCPath`, P2P configuration, and peer baseline.
   Run both binaries' `dumpconfig` with the same user, `HOME`, and recorded
   configuration, datadir, IPC, HTTP, and WebSocket flags. The effective datadir
   and RPC configuration must match.
4. **Prove historical indexing.** Require `TxIndex = true` and confirm transaction
   `0xfa3cbe1ec4220bee33a30d7f922ff4274489503f6c48729abce40e71589988f0`
   resolves at block `14,551,915`. Enabling indexing now does not restore missing
   historical receipts; a failing node needs a coordinator-approved indexed
   snapshot before the upgrade.
5. **Prepare recovery.** Preserve the old binary and its checksum. Verify that
   every validator's keystore and P2P `nodekey` have protected offline backups.
   Name the recovery owner and the destination for stage-matched snapshot
   manifests, checksums, and restore commands.
6. **Publish the fresh-node policy.** Freeze every fresh node start during
   activation. After seal 5, permit snapshot-based onboarding only when the
   post-activation artifact is verified and the coordinator opens it.
   Original-genesis replay under `v2.0.49-elemont` remains prohibited until a
   compatible maintenance binary and regenerated genesis are published.

## Why fresh replay is frozen

The distributed 2024 genesis predates the ELEMONT-era activations. A new binary
replaying it can stage persisted rule changes at different epochs from the live
chain and fail with `wrong event epoch hash`. Existing nodes continuing from
canonical chaindata do not have this replay problem.

The safe route is therefore:

- upgrade existing nodes in place on their verified current datadir;
- forbid original-genesis replay during and after activation; and
- onboard or recover nodes from a verified post-activation snapshot.

## Upgrade-day procedure

### 1. Obtain the final GO

Immediately before the first validator stop, the named coordinator must publish
a timestamped **GO** that records:

- ready and online validator IDs representing more than two-thirds of active
  stake;
- the release tag, commit, and binary checksum;
- the recovery owner and manifest destination; and
- confirmation that the fresh-node freeze and original-genesis prohibition are
  public.

Save the message link. If quorum or any release identity check fails, stop.

### 2. Swap existing nodes in place

Stagger validators one at a time. For each node:

1. Stop it cleanly and prove the process exited. Never use `SIGKILL`.
2. Replace the executable atomically with the verified binary.
3. Start it through the recorded service or script with the recorded working
   directory and configuration. The only permitted launch change is an explicit
   `--datadir` pin already proven necessary in preflight.
4. Confirm `v2.0.49-elemont`, chain ID `207`, advancing height, normal peers,
   unchanged P2P configuration, and the same block hash as the public RPC at a
   fixed height.

Do not start the next validator until the current validator passes its immediate
canonical-chain checkpoint.

### 3. Verify all five seals

Activation follows epoch seals, not a fixed clock. Under normal progress the
complete sequence can take up to roughly 20 hours after the swap.

| Seal | Newly active flags | Required checkpoint |
| --- | --- | --- |
| 1 | `SfcV2`, `Elemont`, `ElemontPubkeyValidation`, `Shanghai`, `PaybackV2` | SFC version/runtime and QuotaContractV2 address match the fixed values |
| 2 | `Cancun` | `vc_getRules` reports `Cancun = true` |
| 3 | `Prague` | `vc_getRules` reports `Prague = true` |
| 4 | `VinuBLS12381` | `vc_getRules` reports `VinuBLS12381 = true` |
| 5 | `VinuLatestEVM` | `vc_getRules` reports `VinuLatestEVM = true` |

Keep the new process running between seals. In particular, validators must not
restart between seal 1 and seal 3. If a validator fails in that interval, keep
it from emitting and use coordinator-approved, stage-matched recovery.

Record the actual activation epoch and block for every flag from live rules and
seal logs. Do not turn projected times into historical facts.

### 4. Publish the final recovery snapshot

Only after seal 5 and the final canonical-chain check:

1. Cleanly stop the non-validator source node and prove the process exited. A
   file-level copy of a live LevelDB datadir is not valid. A tested storage-level
   snapshot is acceptable only while the node is stopped.
2. Build the snapshot without `keystore/`, `nodekey`, `opera.ipc`,
   `static-nodes.json`, or `trusted-nodes.json`. Fail closed if the finished
   archive contains any excluded path.
3. Publish a manifest containing the activation stage, capture block and epoch,
   client version, download URL, size, SHA256, datadir root layout, expected
   ownership, extraction commands, and post-restore verification.
4. Independently download and verify the checksum and archive layout before
   announcing it.
5. Open snapshot-based onboarding only after the verified artifact and recovery
   instructions are public. Keep the original-genesis prohibition in force.

## Final verification

Query `vc_getRules("latest")` locally and through
`https://rpc.vinuchain.org`. All of these must be `true`:

```text
Berlin London Shanghai Cancun Prague VinuBLS12381 VinuLatestEVM
Llr Podgorica SfcV2 Elemont ElemontPubkeyValidation PaybackV2
```

All `SfcV2Patch*` flags and `PaybackV2Patch` must remain absent or `false` on
mainnet. They are repair flags, not missing features.

After seal 1, verify:

- the SFC `version()` result is bytes32 `305`;
- the SFC runtime keccak is
  `0x29b88152209fe22bef409376aa7f137d0e0f571f46afa1385f32320765e49e50`;
- `Economy.QuotaCacheAddress` is
  `0x5D989A2d65d049e2198D91d8ddc31C918f2544AB`; and
- `feeRefundBlockCount()` on that contract returns `75`.

After seal 5, compare the same non-null block hash locally and publicly, and
confirm the node continues advancing with its normal peer count. Do not declare
completion from flags alone.

## Rollback and recovery

### Before seal 1

Prove from the public chain that `SfcV2` is still `false`. A coordinated rollback
may then restore the checksummed `v2.0.0-rc.1` binary and restart on the same
current datadir. Re-run the version, height, peers, and same-height block-hash
checks. Validators must not restore an older datadir copy.

### After seal 1

There is no binary downgrade or pre-seal-datadir rollback. Stop the affected
node and recover with `v2.0.49-elemont` plus a coordinator-approved snapshot
whose manifest matches the node's activation stage. Never replay the original
genesis. Keep a recovered validator from emitting until chain ID, version,
height, peers, and a same-height canonical hash all verify and the coordinator
authorizes it.

## Post-activation hardening

After seal 5, record the historical activation epochs and blocks, regenerate the
distributed mainnet genesis, and prepare the matching `AllowedOperaGenesis` and
`StoredStateRequirements` entries. These facts cannot be pinned safely in the
activation binary before they exist: doing so would reject the pre-seal fleet.

This is defense-in-depth for future stale-genesis and stale-datadir starts, not
another consensus activation. Existing operators complete the August window
with the single `v2.0.49-elemont` swap; the supported fresh-node route remains
the verified post-activation snapshot until a maintenance binary carries the
historical pins.

## Failure paths

- **`wrong event epoch hash` -> stale genesis or wrong-stage data ->** stop the
  node, restore a stage-matched snapshot on `v2.0.49-elemont`, then compare a
  fixed-height hash before enabling validation.
- **Payback cache or transaction-index startup failure -> missing historical
  receipts ->** stop retrying, restore an indexed stage-matched snapshot, then
  recheck the known transaction and startup warm-up.
- **New version but stale height -> wrong datadir or lost P2P configuration ->**
  stop the node, restore the recorded launch values, then verify advancing height,
  peers, and a canonical hash.
- **A seal does not complete -> insufficient upgraded stake or unhealthy
  validators ->** freeze further restarts and escalate to the named coordinator;
  do not restart validators speculatively.

## References

- `opera/rules.go` — staged mainnet rules and the original Quota proxy.
- `opera/payback_v2_address.go` — baked QuotaContractV2 addresses and startup
  guard.
- `gossip/service.go` — five-seal staging dependencies.
- `cmd/opera/launcher/params.go` — allowed genesis and stored-state guards.
- `VinuChain-Docs/technical-docs/vinuchain-mainnet/chain-upgrade-guide.md` —
  public operator commands.
- `vinuchain-ops-docs/ops/mainnet-upgrade-day-runbook-20260829.md` — controlled
  host execution and monitoring.
