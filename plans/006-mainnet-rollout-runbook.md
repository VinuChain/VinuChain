# Plan 006: Write the mainnet hard-fork release runbook

> **Executor instructions**: This plan produces ONE documentation file. Follow
> the outline, fill it from the inlined source material below, and run the
> verification commands. Do not touch code. When done, update the status row
> for this plan in `plans/README.md`.
>
> **Drift check (run first)**: `git diff --stat af41ca7..HEAD -- opera/rules.go opera/payback_v2_address.go`
> These are referenced for exact, current addresses/flags. If they changed,
> read the live files for the values rather than trusting any older copy.

## Status

- **Priority**: P3
- **Effort**: M
- **Risk**: LOW (documentation only)
- **Depends on**: none
- **Category**: docs / direction
- **Planned at**: commit `af41ca7`, 2026-06-16

## Why this matters

VinuChain's mainnet (chain 207) still runs `v2.0.0-rc.1`, while the codebase has
**already staged** the full mainnet hard-fork: `VinuChainMainNetRules()` carries
the full-parity upgrade flag set (`Podgorica`, `SfcV2`, `Elemont`, plus the EVM
forks) and is pinned by tests. The gap to activating it on mainnet is **not
code — it is a high-stakes operational sequence** with hard, order-dependent
prerequisites that, if missed, cause a chain-splitting `wrong event epoch hash`
divergence for fresh-install nodes (the exact failure that bit testnet operators
before). That sequence currently lives only in an internal, gitignored ops log.

This plan captures the **publishable** upgrade procedure and prerequisites in a
tracked runbook so the upgrade is repeatable, reviewable, and external validators
know what to expect. Box-specific execution detail (instance IDs, SSM, key
custody) deliberately stays out of the public doc.

## Current state

- Mainnet runs `v2.0.0-rc.1` (pre-Podgorica/SfcV2/Elemont). Testnet (chain 206)
  has already activated the full flag set + PaybackV2 and serves as the dress
  rehearsal.
- `opera/rules.go` `VinuChainMainNetRules()` — stages the mainnet flags and sets
  `Economy.QuotaCacheAddress`. **Read this file for the exact current flag set
  and address** rather than hardcoding from memory.
- `opera/payback_v2_address.go` — per-network PaybackV2 contract address slots
  and the sentinel/startup-check logic.
- There is no tracked upgrade runbook (`ls docs/*runbook* docs/*upgrade* 2>/dev/null`
  → only `docs/upgrade-guide-genesis-validators.md`, which is a validator setup
  guide, not the hard-fork rollout sequence).
- An internal, **gitignored** `.claude/rules/deployment-log.md` holds the full
  ops detail. The runbook should reference it for box-specific steps but must be
  self-sufficient for the protocol-level procedure.

## The source material to encode (publishable subset)

The runbook must capture these invariants and steps (paraphrase into clean prose;
do not invent beyond this):

**A. Why a mainnet upgrade is dangerous (the core hazard).** The distributed
mainnet genesis (May 2024) pre-dates every elemont-era flag. When a new mainnet
binary first boots, all staged flags fire at the first epoch seal after boot. A
node replaying from the 2024 genesis seals each flag at a *different* first-replay
block than the live chain did — and because `receipt.FeeRefund` (Podgorica) is a
persisted field, the receipts-root mismatch is a full consensus split. The
multi-fork staging is sequential: Cancun stages only after Shanghai is *active*,
Prague only after Cancun — so they activate across *consecutive* epoch seals, not
one.

**B. The three same-day-as-binary prerequisites (none optional).**
1. **Post-upgrade chaindata snapshot** — taken only after the RPC has sealed
   **every** staged flag (wait until the *last* EVM fork, Prague, has sealed —
   not just the first post-boot seal). Excludes nodekey/keystore/ipc/static- and
   trusted-nodes when tarring.
2. **Regenerated distributed genesis** — exported from a node that already sealed
   the flags; update `cmd/opera/launcher/params.go` `AllowedOperaGenesis` with the
   new section hashes so fresh installs don't need `--genesis.allowExperimental`.
3. **Operator announcement** — explicit "no fresh validator installs during or
   within 24 h of the upgrade window" rule. In-place binary swaps on existing
   datadirs are safe; fresh installs from the stale genesis during the window
   diverge and need a chaindata wipe + snapshot restore.

**C. Recommended sequencing.** Ship the consensus-flag rollout
(Podgorica + SfcV2 + Elemont + EVM forks) as one release; stage **PaybackV2** as a
*separate* later release, ≥ 2 weeks after the first has been live without
surprises (combining them compounds blast radius). PaybackV2 is itself a
persisted-state change (it swaps `Economy.QuotaCacheAddress` at its activation
seal), so prerequisites B1–B3 apply to it again on its own day.

**D. Post-activation verification (publishable eth_call checks).** After the
seal, confirm via the public RPC: `vc_getRules` shows the expected `Upgrades.*`
flags true and the expected `Economy.QuotaCacheAddress`; the live SFC bytecode
matches the intended V2 cycle; and (for PaybackV2) `eth_call feeRefundBlockCount()`
on the new contract returns the expected value. Provide these as a copy-pasteable
verification checklist so anyone can confirm the upgrade took.

**E. Binary-swap safety (publishable).** Pre-build off-box, sha256-verify the
binary across build hosts, then swap + restart — never rebuild on a production
box during the window. Never SIGKILL a validator (LevelDB corruption); use clean
SIGINT/`systemctl restart`.

> **Do NOT put in the public runbook**: AWS instance IDs, internal IPs, S3 bucket
> names, SSM document names, or any private-key custody detail (which EOA holds
> which key). Reference "the internal deployment-log" for those. Public addresses
> already in tracked source (`opera/rules.go`) are fine to cite.

## Commands you will need

| Purpose | Command | Expected |
|---------|---------|----------|
| Confirm no runbook exists | `ls docs/*runbook* 2>/dev/null \|\| echo none` | `none` |
| Read current mainnet flags | `sed -n '/func VinuChainMainNetRules/,/^}/p' opera/rules.go` | the staged flag set + address |
| Markdown link sanity (opt.) | `grep -n '](' docs/mainnet-upgrade-runbook.md` | links resolve to real paths |
| Build sanity | `go build ./...` | exit 0 (nothing structural touched) |

## Scope

**In scope** (create):
- `docs/mainnet-upgrade-runbook.md`

**Out of scope** (do NOT touch):
- Any `.go` file, CI, or other docs.
- The gitignored `.claude/rules/deployment-log.md` (read-reference only; do not
  copy its infra/key specifics into the public doc).
- Do NOT actually perform any upgrade step — this plan writes the runbook only.

> Note: `docs/` is gitignored **except** explicitly allowlisted files (see
> `.gitignore`: `docs/*` then `!docs/payback-cache-restart-determinism.md`).
> To make the runbook a tracked file you must ALSO add a `!docs/mainnet-upgrade-runbook.md`
> allowlist line to `.gitignore` (add it next to the existing `!docs/...` line).
> If the operator prefers the runbook stay untracked, skip the `.gitignore` edit
> and note that in your report. Treat `.gitignore` as in-scope for this one line.

## Git workflow

- Branch: `advisor/006-mainnet-runbook`.
- Commit: `docs: add mainnet hard-fork release runbook`.
- Do NOT push or open a PR unless instructed.

## Steps

### Step 1: Confirm the target file does not exist and read current flags

```sh
ls docs/*runbook* 2>/dev/null || echo none
sed -n '/func VinuChainMainNetRules/,/^}/p' opera/rules.go
```

**Verify**: prints `none`, and you can see the current staged mainnet flag set
and `QuotaCacheAddress` to cite accurately.

### Step 2: Write `docs/mainnet-upgrade-runbook.md`

Use this outline, filling each section from source material A–E above:

```markdown
# Mainnet Hard-Fork Release Runbook

## Scope & current state
(Mainnet on v2.0.0-rc.1; flags staged in VinuChainMainNetRules; testnet is the rehearsal.)

## The core hazard: stale-genesis replay divergence
(Source A.)

## Prerequisites — all three, same day as the binary (none optional)
1. Post-upgrade chaindata snapshot (after Prague seals)
2. Regenerated distributed genesis + AllowedOperaGenesis update
3. Operator announcement (no fresh installs in the window)
(Source B.)

## Release sequencing
(Consensus-flag release first; PaybackV2 as a separate later release ≥2 weeks on. Source C.)

## Execution outline
(Pre-build off-box, sha256-verify, swap, clean SIGINT restart. Source E. For
box-specific steps — instance access, SSM, S3 paths — see the internal
deployment-log; not reproduced here.)

## Post-activation verification checklist
(Copy-pasteable vc_getRules / SFC bytecode / feeRefundBlockCount checks. Source D.)

## Rollback / divergence response
(If a fresh-install node shows `wrong event epoch hash`: wipe chaindata, restore
from the post-seal snapshot. Existing-datadir in-place swaps are unaffected.)

## References
- opera/rules.go (staged flags)
- opera/payback_v2_address.go (PaybackV2 address slots)
- internal deployment-log (box-specific execution; not in this repo's public tree)
```

> **Reference hygiene**: only link files that are actually **tracked** (would exist
> in a fresh clone). `opera/rules.go` and `opera/payback_v2_address.go` are tracked.
> Do NOT cite `docs/upgrade-guide-genesis-validators.md` — despite existing on disk
> it is gitignored (caught by `docs/*`, not allowlisted), so a public reader cloning
> the repo would hit a dead link. Confirm with `git ls-files <path>` (non-empty =
> tracked) before adding any new cross-reference to the runbook.

**Verify**: `test -f docs/mainnet-upgrade-runbook.md && echo ok` → `ok`; and
`grep -c '^## ' docs/mainnet-upgrade-runbook.md` → ≥ 6 (all sections present).

### Step 3: Make it tracked (if operator wants it public)

Add to `.gitignore`, immediately after the existing `!docs/payback-cache-restart-determinism.md` line:

```
!docs/mainnet-upgrade-runbook.md
```

**Verify**: `git check-ignore docs/mainnet-upgrade-runbook.md` prints **nothing**
(file is no longer ignored), and `git status --porcelain` shows the new file as
untracked/added.

### Step 4: Confirm no secrets/infra leaked into the public doc

This grep targets ONLY mechanical, unambiguous leaks. It deliberately does **not**
scan for `0x…` addresses — the public `QuotaCacheAddress` from `opera/rules.go` is
*meant* to appear in the verification checklist (it is a contract address, not a
secret), so an address regex would false-positive on intended content. The
key-custody guard (no EOA *key-holder* hints) is enforced by the textual rule
below, not by this scan.

```sh
grep -nE 'i-[0-9a-f]{17}|AKIA[0-9A-Z]{16}|s3://|10\.0\.[0-9]|PRIVATE_(MAIN|TEST)|VinuChain-RunAsUbuntu' docs/mainnet-upgrade-runbook.md
```

**Verify**: the grep prints **nothing** (zero matches, exit 1) — no AWS instance
IDs, access keys, S3 URIs, internal `10.0.x` IPs, private-key env-var names, or the
internal SSM document name leaked in.

Then do a one-time manual read for the one thing a regex cannot catch: the doc must
not state *which EOA holds which key* (key-custody). Naming a contract address for
verification is fine; naming a key-holder address is not. If you find such a
sentence, remove it and reference "the internal deployment-log" instead.

## Test plan

Documentation only — no Go tests. Verification is the per-step `grep`/`test`
checks plus `go build ./...` exit 0 (proves no code was touched).

## Done criteria

ALL must hold:

- [ ] `docs/mainnet-upgrade-runbook.md` exists with ≥ 6 `## ` sections covering
      hazard, prerequisites, sequencing, execution, verification, rollback.
- [ ] The Step 4 secret/infra grep prints **zero matches** (exits non-zero), and
      you did the one-time manual key-custody read.
- [ ] If tracked: `git check-ignore docs/mainnet-upgrade-runbook.md` prints nothing,
      and `.gitignore` has the new `!docs/...` allowlist line.
- [ ] Every cross-reference in the runbook's "References" section is a tracked file
      (`git ls-files <path>` non-empty) — no dangling links to gitignored docs.
- [ ] `go build ./...` exits 0.
- [ ] `git status --porcelain -- docs/mainnet-upgrade-runbook.md .gitignore` lists
      only the runbook and (if Step 3 done) `.gitignore`. (An unscoped `git status`
      also showing `?? plans/` is expected — untracked, leave it.)
- [ ] `plans/README.md` status row for 006 updated (allowed index write).

## STOP conditions

Stop and report back (do not improvise) if:

- `VinuChainMainNetRules()` no longer exists or its staged flags differ
  materially from source material A–C (the upgrade plan-of-record may have
  changed — the runbook must reflect reality, so confirm before writing).
- You are tempted to include box-specific infra or key-custody detail to make the
  runbook "complete" — that belongs in the internal log, not here.

## Maintenance notes

- This runbook is the protocol-level companion to the internal deployment-log;
  when the mainnet upgrade actually executes, fold the dated, as-run record into
  the internal log and keep this doc as the reusable template.
- If plan 004 (CHANGELOG) also lands, add a runbook step to update CHANGELOG.md as
  part of the release.
- Reviewer should confirm the doc contains no operational secrets and that the
  prerequisite ordering (snapshot after *Prague* seal, not first seal) is stated
  unambiguously — that exact detail is the one most likely to be gotten wrong.
```
