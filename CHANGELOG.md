# Changelog

## VinuChain (Elemont series)

> Generated from the `vX.Y.Z-elemont` git tag history. Each entry lists the
> user- and operator-facing changes shipped in that release.

### Unreleased

- fix(launcher): keep a populated legacy `.opera` data directory in use until
  `.vinuchain` independently carries chain state. Previously the mere existence
  of `.vinuchain` ended the fallback, so an empty directory left behind by a
  failed flagless start — or the `go-opera/`+`keystore/` shell that
  `opera account list` creates — silently moved the node onto empty state and
  resynced it from genesis.
- fix(launcher): emit the legacy-data-directory notice through the node logger,
  once, and only when the default data directory is actually in use. It
  previously printed twice per invocation via the standard library logger, on
  every command including `opera version`, even when `--datadir` pointed
  somewhere else entirely.
- test(launcher): cover `DefaultDataDir` resolution across platforms and across
  the legacy-migration states, including the two regressions above.

### v2.0.39-elemont — 2026-06-11

- chore(release): v2.0.39-elemont — PaybackCache restart warm-up
- fix(payback): warm PaybackCache (epochs E-1 and E) on startup to fix mid-epoch restart consensus split (A1/T2)
- refactor(payback): drop dead address params; truthful CHANGELOG header
- ci: add GitHub Actions build/vet/test gate and remove dead appveyor.yml

### v2.0.38-elemont — 2026-06-04

- fix(evmwriter): suppress SFC balance warning noise

### v2.0.37-elemont — 2026-06-03

- chore(release): consume go-vinu precompile vector tests

### v2.0.36-elemont — 2026-06-03

- fix(evm): enforce VinuLatestEVM transaction gas cap

### v2.0.35-elemont — 2026-06-03

- chore(release): bump Elemont to 2.0.35

### v2.0.34-elemont — 2026-06-03

- feat(evm): add Vinu latest-EVM fork and eth_config

### v2.0.33-elemont — 2026-06-03

- feat(evm): stage Vinu BLS12-381 fork

### v2.0.32-elemont — 2026-06-03

- feat(config): stage full-parity mainnet upgrade flags + fix Quota proxy

### v2.0.31-elemont — 2026-05-29

- chore(release): bump Elemont version to 2.0.31
- refactor(gossip,ethapi): extract and pin unprotected-tx mainnet guard

### v2.0.30-elemont — 2026-05-28

- feat(ethapi): allowlist the canonical Arachnid deployer tx on mainnet

### v2.0.29-elemont — 2026-05-28

- feat(launcher): expose --rpc.allow-unprotected-txs CLI flag
- fix(ethapi): keep accepted Prague tx submission successful

### v2.0.28-elemont — 2026-05-19

- feat(evm): enable Prague set-code transactions

### v2.0.27-elemont — 2026-05-19

- chore(release): bump Elemont version to 2.0.27
- fix(evm): harden Elemont EIP activation

### v2.0.26-elemont — 2026-05-18

- fix(evm): close EIP audit gaps
- fix(evm): wire Shanghai transaction rules

### v2.0.24-elemont — 2026-05-18

- feat(evm): stage Cancun selfdestruct rules

### v2.0.23-elemont — 2026-05-18

- feat(evm): stage Cancun opcode support

### v2.0.22-elemont — 2026-05-18

- feat(evm): activate Shanghai on testnet

### v2.0.21-elemont — 2026-05-17

- fix(sfc): auto-backfill delegation rows on activation

### v2.0.20-elemont — 2026-05-17

- feat(sfc): wire SfcV2Patch6 backfill bytecode

### v2.0.19-elemont — 2026-05-16

- chore(release): bump Elemont version to 2.0.19
- fix(payback): rebind testnet PaybackV2 to corrected contract

### v2.0.18-elemont — 2026-05-15

- feat(payback,config): activate PaybackV2 on testnet
- fix(payback,config): apply PaybackV2 scaffold review fixes
- feat(payback,config): scaffold PaybackV2 upgrade flag
- fix(payback): correct quota upgrade event topic
- fix(testnet): use current public RPC in SFC helper

### v2.0.17-elemont — 2026-05-10

- fix(testnet): align payback proxy release target

### v2.0.16-elemont — 2026-05-10

- fix(payback): support receiver-funded quota stakes

### v2.0.15-elemont — 2026-05-06

- chore(release): v2.0.15-elemont
- fix(consensus+payback): preserve latest activation and tracked refunds
- fix(launcher): keep default bootnodes available for named VinuChain networks

### v2.0.14-elemont — 2026-05-02

- chore(release): v2.0.14-elemont
- refactor(utils): rename fast buffer Read/WriteByte to clear go vet warnings
- feat(consensus): activate SfcV2Patch5 + ElemontPubkeyValidation on testnet
- feat(sfc): install Cycle-161 SFC bytecode
- chore(release): v2.0.13-elemont
- feat(consensus): scaffold malformed-pubkey skip + SfcV2Patch5 Cycle-161 placeholder
- fix(rpc): compute gasUsedRatio from real block usage in eth_feeHistory
- fix(ops): add persisted opera-swap script with trace-node safety guard

### v2.0.12-elemont — 2026-04-25

- chore(release): v2.0.12-elemont — multi-SfcV2Patch divergence warn
- refactor(gossip): drop unused migration consts and launder typed nil
- feat(scripts): chaindata snapshot producer with SNAPSHOT_INFO.txt
- feat(gossip): warn on multi-SfcV2Patch co-activation at single seal

### v2.0.11-elemont — 2026-04-23

- chore(release): v2.0.11-elemont — SfcV2Patch4 Cycle-160 lock-end-time fix
- feat(sfc): inline real Cycle-160 bytecode for SfcV2Patch4
- fix(sfc): harden Patch4 bytecode validator with equality check and startup-time enforcement (#12)
- feat(sfc): add SfcV2Patch4 scaffolding + deadbeef placeholder guard

### v2.0.10-elemont — 2026-04-19

- chore(release): v2.0.10-elemont — SfcV2Patch3 reentrancy guard rollup
- feat(sfc): add SfcV2Patch3 flag for Cycle-159 bytecode re-flash

### v2.0.9-elemont — 2026-04-19

- chore(release): v2.0.9-elemont — trusted preset for regenerated testnet genesis
- feat(config): add trusted preset for 2026-04-19 testnet genesis with history

### v2.0.8-elemont — 2026-04-19

- chore(release): v2.0.8-elemont — PeerProgress drift caps hotfix
- fix(gossip): remove peer progress drift caps blocking stale-node rejoin

### v2.0.7-elemont — 2026-04-19

- chore(release): v2.0.7-elemont — per-peer quota sizing hotfix
- fix(gossip): size per-peer in-flight quota for full sync chunks

### v2.0.6-elemont — 2026-04-18

- feat(rpc): add vc_getPaybackBalance with concurrency cap

### v2.0.5-elemont — 2026-04-18

- feat(sfc): add SfcV2Patch2 flag to re-flash testnet SFC bytecode

### v2.0.4-elemont — 2026-04-17

- chore(release): v2.0.4-elemont — lachesis-base v0.1.6 rollup

### v2.0.3-elemont — 2026-04-17

- chore(release): v2.0.3-elemont — Cycle 152-158 audit hardening rollup
- fix(evm): clamp prevGasPowerLeft to maxGasPower in CalcValidatorGasPower
- fix(gossip): saturate DirtyGasRefund additions in gas price oracle backend
- fix(gossip): saturate GasRefund addition to prevent uint64 overflow in gas power check
- fix(evm): restore math.MaxUint64 gas pool for internal transactions in evmBlockWith
- fix(config): remove exported FakePassword constant from validatorpk public API
- refactor(evm): dedupe BaseFeeFloor construction and tighten EvmHeader consistency
- fix(payback): correct refundGas comment re: simulation guard behavior
- fix(sfc): correct reentrancyguard check for post-upgrade counter state
- refactor(evm): move baseFeeFloor from ApplyMessage param into BlockContext
- fix(payback): use chain-configured MinGasPrice as congestion threshold
- refactor(evm): cache base fee per block, eliminate per-tx bigint alloc
- fix(evm): enable dynamic EIP-1559 base fee with correct block gas limit
- fix(rpc): compute pending base fee from CalcBaseFee in tx pool
- fix(payback): suppress quota refunds when network is congested

### v2.0.2-elemont — 2026-04-11

- refactor(payback): hoist DB read out of pc.mu in getPaybackData
- fix(gossip): restore optimistic pre-check in createEvent

### v1.0.2-elemont — 2026-04-11

- feat(config): add SfcV2Patch flag to re-flash SFC V2 bytecode on activation
- fix(sfc): patch restakeRewards to isolate newly-stashed rewards before rescale

### v1.0.1-elemont — 2026-04-10

- fix(gossip): stage upgrade flags via DirtyRules so SFC V2 bytecode swap fires

### v1.0.0-elemont — 2026-04-07

- Initial elemont release.

---

## Upstream-inherited history (Fantom go-opera / Lachesis)

## v0.4.0 (October 14, 2018)

SECURITY:

* keygen: write keys to files instead of tty. 

FEATURES:

* proxy: Introduced in-memory proxy.
* cmd: Enable reading config from file (lachesis.toml, .json, or .yaml)

IMPROVEMENTS:

* node: major refactoring of configuration and initialization of Lachesis node.
* node: Node ID is calculated from public key rather than from sorting the 
peers.json file.

## v0.3.0 (September 4, 2018)

FEATURES:

* poset: Replaced Leemon Baird's original "Fair" ordering method with 
Lamport timestamps.
* poset: Introduced the concept of Frames and Roots to enable initializing a
poset from a "non-zero" state.
* node: Added FastSync protocol to enable nodes to catch up with other nodes 
without downloading the entire poset. 
* proxy: Introduce Snapshot/Restore functionality.

IMPROVEMENTS:

* poset: Refactored the consensus methods around the concept of Frames.
* poset: Removed special case for "initial" Events, and make use of Roots 
instead. 
* docs: Added sections on Lachesis and FastSync.
