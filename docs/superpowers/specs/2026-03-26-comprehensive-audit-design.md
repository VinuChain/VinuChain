# Comprehensive Audit Design: VinuChain + go-vinu

**Date:** 2026-03-26
**Scope:** Full security audit of VinuChain (go-opera fork, ~74K LoC) and go-vinu (go-ethereum fork, delta + critical paths)
**Branch:** `elemont`
**Status:** Pre-mainnet — both Podgorica and SFC V2 are staging/testnet only

## Goals

1. Verify hard fork readiness for Podgorica + SFC V2 deployment
2. Confirm all 60+ prior audit remediation fixes are solid with no regressions
3. Assess full security posture across every subsystem
4. Deliver findings report, code fixes (all severities), and regression tests

## Audit Structure

Three sequential rounds, each using a 4-phase structure:

```
Phase 1: Context Building     — shared knowledge base for specialists
Phase 2: Domain Audits        — parallel specialist agents per domain
Phase 3: Cross-Cutting        — inter-domain analysis, variant hunting, FP triage
Phase 4: Remediation          — fix every finding, write regression test, code review
```

Rounds are gated — each must complete before the next begins. Findings cascade forward.

## Severity Scale

| Severity | Definition | Examples |
|----------|-----------|---------|
| Critical | Funds at risk, consensus break, chain halt | Incorrect FeeRefund allowing theft, SFC reward drain, fork-inducing bug |
| High | Incorrect state, economic manipulation, privilege escalation | Staking threshold bypass, governance takeover, incorrect epoch transition |
| Medium | DoS vector, data leak, degraded security property | RPC query DoS, P2P flooding, timing side-channel in key handling |
| Low | Code quality, best practice violation, minor edge case | Missing error check, suboptimal gas usage, unchecked type assertion |
| Info | Observation, documentation gap, style | Missing godoc, dead code, naming inconsistency |

**All severities get fixed.** No finding is left unaddressed.

## Deliverables

| Artifact | Path |
|----------|------|
| Round A findings | `docs/audit/round-a-findings.md` |
| Round B findings | `docs/audit/round-b-findings.md` |
| Round C findings | `docs/audit/round-c-findings.md` |
| Final consolidated report | `docs/audit/final-report.md` |
| Code fixes | On `elemont` branch, one commit per finding or logical group |
| Regression tests | Co-located with fixes, one test per finding minimum |

## Priority Tiers

| Tier | Domains | Covered In |
|------|---------|-----------|
| Tier 1 (highest) | Payback/fee refund, SFC V2 Solidity, consensus integration, go-vinu FeeRefund | Round A (primary), Round B (verification) |
| Tier 2 | P2P, RPC/API, event validation, emitter, storage, key management, gas oracle | Round C (primary) |

---

## Round A: Hard Fork Readiness (Podgorica + SFC V2)

### A.1 — Context Building (parallel)

| Agent/Tool | Output |
|-----------|--------|
| `audit-context-building:audit-context` | Deep architectural context of payback, SFC, consensus integration |
| `x-ray` | Threat model, entry points, git-weighted attack surfaces, readiness verdict |
| `entry-point-analyzer:entry-points` | All state-changing functions by access level |
| go-vinu delta extraction (`git diff` against upstream go-ethereum) | Isolated VinuChain-specific changes |

### A.2 — Domain Audits (parallel specialists)

**Payback Auditor**
- Scope: `payback/`, refund math in `evmcore/`, `receipt.FeeRefund`, quota logic
- Focus: Refund calculation correctness, epoch boundary edge cases, staking threshold manipulation, overflow/underflow, static call failure handling

**SFC/Solidity Auditor**
- Scope: `opera/contracts/sfc/`, driver, evmwriter, payback proxy, `sfc.sol`, `sfc_fixed.sol`
- Focus: SFC V2 staking/delegation/rewards, 30% base fee burn correctness, governance attacks, reentrancy, access control, event emission accuracy
- Tools: `solidity-auditor` (DEEP mode), `semantic-guard-analysis`, `state-invariant-detection`

**Consensus Integration Auditor**
- Scope: `gossip/blockproc/`, `vecmt/`, epoch transitions, upgrade flag activation
- Focus: Podgorica activation safety, block processing pipeline with new rules, event validation correctness

**go-vinu Fork Auditor**
- Scope: Delta files from A.1 + surrounding critical paths (state processor, receipt handling, EVM call chain)
- Focus: FeeRefund field integrity through full tx lifecycle, receipt serialization/deserialization, state transition correctness

### A.3 — Cross-Cutting Analysis

| Analysis | Tool | Purpose |
|---------|------|---------|
| Dimensional Analysis | `dimensional-analysis` | Wei/gwei mismatches, epoch/block confusion in payback and fee math |
| Semantic Guard Analysis | `semantic-guard-analysis` | Functions bypassing guards applied elsewhere in SFC/driver |
| State Invariant Detection | `state-invariant-detection` | Conservation laws in SFC V2 (totalSupply, delegation sums, reward pools) |
| Variant Analysis | `variant-analysis:variants` | Hunt variants of any Phase 2 bug patterns |
| FP Triage | `fp-check:fp-check` | Eliminate false positives before committing to fixes |

### A.4 — Remediation

For every confirmed finding:
1. Implement fix on `elemont` branch
2. Write regression test proving the fix
3. Code review agent validates the fix
4. Record in `docs/audit/round-a-findings.md`

**Gate:** All Round A findings fixed and tested before Round B begins.

---

## Round B: Remediation Verification

### B.1 — Context Building

| Agent/Tool | Purpose |
|-----------|---------|
| Prior Findings Loader | Parse 60+ findings from prior audits (commits `db0fa2e`, `3710df7`, `ff00aa7`) |
| Round A Findings Loader | Load Round A findings and fixes as verification targets |
| `differential-review:diff-review` | Full `elemont` vs `main` diff — blast radius of all changes |

### B.2 — Verification Audits (parallel)

**Fix Completeness Verifier**
- Checks each of the 60+ prior fixes and all Round A fixes actually address root cause, not just symptom

**Regression Hunter**
- All files modified by prior fixes — looks for side effects, logic inversions, off-by-one errors introduced during remediation

**Test Coverage Auditor**
- Evaluates whether existing tests cover the fixed code paths — identifies gaps

**Commit Archaeology**
- `git log` + `git blame` on fixed files — checks whether fixes were reverted, overwritten, or conflicted

### B.3 — Cross-Cutting Analysis

| Analysis | Tool | Purpose |
|---------|------|---------|
| Static Analysis | `static-analysis:semgrep` | Full codebase Semgrep scan for patterns missed in manual review |
| Constant-Time | `constant-time-analysis:ct-check` | Timing side-channels in validator key handling, signature verification |
| Insecure Defaults | `insecure-defaults` | Hardcoded secrets, weak auth, fail-open patterns |

### B.4 — Remediation

Fix all findings, regression tests, code review. Record in `docs/audit/round-b-findings.md`.

**Gate:** All prior fixes verified. All new findings fixed. No regressions detected.

---

## Round C: Full Security Posture

### C.1 — Context Building

| Agent/Tool | Purpose |
|-----------|---------|
| Attack Surface Mapper | Enumerate every external interface: RPC endpoints, P2P handlers, CLI flags, genesis config |
| `supply-chain-risk-auditor` | Dependency health audit on `go.mod` — abandoned packages, known CVEs |
| go-vinu Full Delta | Complete diff of go-vinu v1.20.1-quota against upstream go-ethereum — all changes, not just payback |

### C.2 — Domain Audits (parallel — broadest sweep)

| Specialist | Scope |
|-----------|-------|
| **P2P Security** | `gossip/handler.go`, `gossip/protocols/`, peer selection, message validation — eclipse attacks, flooding, malicious peer handling |
| **RPC/API Security** | `ethapi/`, `vcclient/`, all exposed endpoints — auth, rate limiting, expensive query DoS, data leaks |
| **Event Validation** | `eventcheck/` (all 6 sub-checkers) — bypass vectors, malformed event acceptance |
| **Emitter Security** | `gossip/emitter/` — tx selection manipulation, event creation, action file handling |
| **State & Storage** | `gossip/evmstore/`, `topicsdb/`, `utils/iodb/` — LevelDB access patterns, pruner safety, data corruption vectors |
| **Key Management** | `valkeystore/` — encryption correctness, cache safety, memory handling of validator private keys |
| **Gas Oracle** | `gossip/gasprice/` — price manipulation, stale price exploitation |
| **Genesis & Integration** | `integration/`, `gossip/apply_genesis.go` — genesis validation, chain init safety |

### C.3 — Cross-Cutting Analysis

| Analysis | Tool | Purpose |
|---------|------|---------|
| Variant Analysis | `variant-analysis:variants` | Hunt variants of ALL findings from Rounds A, B, and C.2 |
| Sharp Edges | `sharp-edges` | Dangerous API patterns, footgun configs across codebase |
| Dimensional Analysis | `dimensional-analysis` | Full codebase sweep — all numeric operations |
| Custom Semgrep Rules | `semgrep-rule-creator:semgrep-rule` | Rules built from Round A/B/C finding patterns, then run via `static-analysis:semgrep` |

### C.4 — Remediation

Fix all findings, regression tests, code review. Record in `docs/audit/round-c-findings.md`.

### Final Consolidation

After C.4:
- Write `docs/audit/final-report.md` consolidating all findings across all 3 rounds
- Summary statistics: findings by severity, by domain, fix status
- Residual risk assessment: design trade-offs and known limitations that can't be fixed in code
- Verification that the build passes and all tests (existing + new regression tests) are green

---

## Tools & Skills Matrix

| Category | Tools Used |
|----------|----------|
| Pre-Audit Recon | `x-ray`, `audit-context-building:audit-context`, `entry-point-analyzer:entry-points` |
| Solidity Audit | `solidity-auditor` (DEEP), `semantic-guard-analysis`, `state-invariant-detection` |
| Go Static Analysis | `static-analysis:semgrep`, `insecure-defaults`, `sharp-edges` |
| Numeric Correctness | `dimensional-analysis` (full suite: scanner, annotator, discoverer, propagator, validator) |
| Crypto Safety | `constant-time-analysis:ct-check` |
| Finding Quality | `fp-check:fp-check`, `variant-analysis:variants` |
| Diff Review | `differential-review:diff-review` |
| Supply Chain | `supply-chain-risk-auditor` |
| Custom Rules | `semgrep-rule-creator:semgrep-rule` |
| Code Review | Code Reviewer agent on every fix |
| Verification | `make test`, `make opera`, build + test green gate |

## Agent Allocation

| Agent Type | Round A | Round B | Round C |
|-----------|---------|---------|---------|
| Context builders (parallel) | 4 | 3 | 3 |
| Domain specialists (parallel) | 4 | 4 | 8 |
| Cross-cutting analyzers (parallel) | 5 | 3 | 4 |
| Remediators | As needed per finding | As needed | As needed |
| Code reviewers | 1 per fix | 1 per fix | 1 per fix |

## Constraints

- All work stays on the `elemont` branch
- No worktrees — parallel agents operate in the same working directory
- Build must compile (`make opera`) after each remediation phase
- Tests must pass (`make test`) after each remediation phase
- go-vinu fork is read-only (module cache) — findings in go-vinu are reported but fixes require a separate go-vinu PR
