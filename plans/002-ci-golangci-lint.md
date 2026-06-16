# Plan 002: Add a `golangci-lint` static-analysis job to CI

> **Executor instructions**: Follow this plan step by step. Run every
> verification command and confirm the expected result before moving on. If
> anything in "STOP conditions" occurs, stop and report — do not improvise.
> When done, update the status row for this plan in `plans/README.md`.
>
> **Drift check (run first)**: `git diff --stat af41ca7..HEAD -- .github/workflows/ci.yml`
> This plan edits the same file as plans 001 and 003. If 001 has already landed,
> the file will already contain a `vulncheck` job — that is fine; add the lint
> job alongside it. Only treat a STOP-worthy mismatch if the `build-test`/`race`
> jobs were renamed or the checkout/setup-go pattern changed.

## Status

- **Priority**: P1
- **Effort**: M
- **Risk**: LOW
- **Depends on**: plans/001 (same file — execute after, or bundle into one CI PR)
- **Category**: dx
- **Planned at**: commit `af41ca7`, 2026-06-16

## Why this matters

CI's only static analysis today is `go vet ./...`, which catches a deliberately
narrow set of issues. `golangci-lint` aggregates linters (`staticcheck`,
`ineffassign`, `errcheck`, `unconvert`, `unused`, …) that catch a much larger
class of real bugs — error-shadowing, ineffectual assignments, unreachable
guards, mishandled errors — on a consensus node where those matter. The repo
has **no** `.golangci.yml` and no lint job.

The catch: this is a fork of go-ethereum with ~90K LOC of inherited code, so a
full-tree lint will report hundreds of pre-existing issues. The realistic way
to introduce a linter to a legacy codebase is **only-new-issues** mode: the job
reports only issues introduced by the PR under review, so it is immediately
actionable without a giant cleanup, and the tree can be ratcheted toward
clean over time.

## Current state

- `.github/workflows/ci.yml` — jobs `build-test` and `race` (plus `vulncheck`
  if plan 001 landed). No lint job. Same checkout/setup-go pattern described in
  plan 001's "Current state".
- No `.golangci.yml`, `.golangci.yaml`, or `.golangci.toml` exists at repo root
  (`ls .golangci.*` → not found).
- Generated/inherited code that must be excluded from linting:
  - `gossip/contract/sfc100/contract.go`, `payback/contract/sfc/contract.go`
    (~4730 LOC each — abigen output)
  - `gossip/contract/driver100/contract.go`, `*_legacy.go` siblings,
    `gossip/contract/driverauth100/contract.go`, `gossip/contract/netinit100/`
  - i.e. everything under `gossip/contract/**` and `payback/contract/**` is
    generated bindings.
- `go vet ./...` currently passes clean (so the baseline has no vet-level noise).

## Commands you will need

| Purpose                | Command | Expected on success |
|------------------------|---------|---------------------|
| Install golangci-lint  | `go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest` (see Step 1 for version handling) | binary in `$(go env GOPATH)/bin` |
| Lint one package       | `$(go env GOPATH)/bin/golangci-lint run ./version/...` | runs to completion (may print issues) |
| Lint whole tree (noise survey) | `$(go env GOPATH)/bin/golangci-lint run ./... --issues-exit-code=0` | exits 0, prints the inherited-issue count |
| Validate config        | `$(go env GOPATH)/bin/golangci-lint config verify` (v2) | "Configuration is valid" |
| YAML sanity            | `python3 -c "import yaml; yaml.safe_load(open('.github/workflows/ci.yml'))"` | exit 0 |

## Scope

**In scope** (the only files you should modify/create):
- `.golangci.yml` (create)
- `.github/workflows/ci.yml` (add one job)

**Out of scope** (do NOT touch):
- **Any `.go` file.** Do NOT fix lint findings in existing code under this plan —
  the whole point of only-new-issues mode is to avoid a sweeping code change.
  Cleanups of inherited findings are separate, reviewed work.
- `go.mod` / `go.sum`.

## Git workflow

- Branch: `advisor/002-ci-golangci-lint` (or continue the CI-hardening branch
  from plan 001).
- One commit; Conventional Commits style, e.g. `ci: add golangci-lint (new-issues only)`.
- Do NOT push or open a PR unless instructed.

## Steps

### Step 1: Install golangci-lint, then PIN the exact version everywhere

golangci-lint v2 is the current major line and supports Go 1.25. Install it, then
**record the exact version** — you will pin that same version in the CI action so
local and CI never diverge (a floating `@latest` in CI can resolve to a different
build than you validated against, silently breaking the config-vs-tool match):

```sh
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
$(go env GOPATH)/bin/golangci-lint version   # e.g. "golangci-lint has version 2.1.6"
$(go env GOPATH)/bin/golangci-lint run ./version/...   # small leaf package
```

Write down the printed version (e.g. `v2.1.6`). You will use it verbatim in
Step 4's `version:` field — NOT `latest`.

**Verify**: the `version` command prints a `2.x` version, and the `run` command
completes without a tool/internal error (it may print zero or a few lint findings
— that is fine here; lint findings are not a tool error).

STOP and report (do not improvise) if:
- `go install` errors that it requires a Go newer than 1.25.8, OR it silently
  triggers a Go-toolchain auto-download to a newer Go (look for a `go: downloading
  go1.XX` line) — either means a version mismatch you should not paper over.
- `golangci-lint version` prints a `1.x` or `3.x` version rather than `2.x` — the
  config schema in Step 3 and the action tag in Step 4 are written for v2; a
  different major needs the operator's call. Report it.

### Step 2: Survey the inherited-issue noise floor (informational)

```sh
$(go env GOPATH)/bin/golangci-lint run ./... --issues-exit-code=0 2>&1 | tail -5
```

Record the total issue count in your report. This is the backlog the
only-new-issues gate intentionally does not block on. (Expect it to be large —
that is the inherited go-ethereum code, and is why this plan does not gate the
whole tree.)

### Step 3: Create `.golangci.yml`

Use the golangci-lint **v2** config schema below. It enables a conservative,
low-false-positive linter set on top of the v2 defaults, and excludes generated
bindings. If your installed version is v1.x, see the STOP/adapt note at the end
of this step.

```yaml
version: "2"

linters:
  enable:
    - errcheck
    - govet
    - ineffassign
    - staticcheck
    - unused
    - unconvert
    - misspell
  exclusions:
    # Generated abigen bindings — never hand-edited, will never be "clean".
    # NOTE: `paths` entries are REGULAR EXPRESSIONS (not literal globs); the bare
    # strings below match the intended directories as substrings of file paths.
    generated: lax
    paths:
      - gossip/contract
      - payback/contract

issues:
  # Cap so a single noisy file cannot hide everything else.
  max-issues-per-linter: 0
  max-same-issues: 0
```

**Verify**: `$(go env GOPATH)/bin/golangci-lint config verify` → "Configuration is valid".

If `config verify` fails, distinguish the two failure modes (do not conflate them):

1. **Unknown linter name** — the error names a specific linter (e.g. `unknown
   linters: unconvert`). FIX: delete only that one linter from the `enable:` list
   and re-run `config verify`. Note the removal in your report. Do NOT touch the
   schema or the other linters.
2. **Schema rejected wholesale** — errors about unknown top-level keys
   (`version`, `linters`, `exclusions`, `issues`) or "unsupported version". This
   means your installed tool is NOT v2 — but Step 1 already STOPs on a non-v2
   version, so if you reach here, STOP and report rather than guessing a schema.

Do not add linters beyond the list above, and do not invent schema keys.

### Step 4: Add the `lint` job to `.github/workflows/ci.yml`

Use the official action with `only-new-issues: true` so PRs are gated on *newly
introduced* issues only. The action needs full history to compute the diff base,
hence `fetch-depth: 0`.

```yaml
  lint:
    name: golangci-lint
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4
        with:
          fetch-depth: 0   # required for only-new-issues diff base

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version-file: go.mod
          cache: false

      # only-new-issues: report only issues the PR introduces, so the large
      # inherited go-ethereum backlog does not block contributors. Ratchet
      # toward full-tree by dropping this flag once the backlog is burned down.
      - name: golangci-lint
        uses: golangci/golangci-lint-action@v7
        with:
          version: v2.1.6          # ← REPLACE with the exact version from Step 1
          only-new-issues: true
```

> **Pin `version:` to the exact golangci-lint version you recorded in Step 1**
> (e.g. `v2.1.6`), NOT `latest` — this is what keeps the config you validated
> locally in lockstep with what CI runs. The action tag `@v7` is the action major
> that drives golangci-lint **v2**; keep `@v7` as long as Step 1 gave a `2.x`
> version. (If Step 1 gave a non-2.x version you should already have STOPped.) The
> action installs and caches golangci-lint itself — the VinuChain org's
> GitHub-hosted Actions support this.

**Verify**: `python3 -c "import yaml; yaml.safe_load(open('.github/workflows/ci.yml'))"` → exit 0.
(Needs PyYAML; if `python3 -c "import yaml"` raises `ModuleNotFoundError`, skip
this check — CI itself rejects malformed YAML on push.)

### Step 5: Prove the linter actually works AND the gate is green

Two distinct checks — the first proves the **linter+config actually run** (a
non-tautological functional test); the second proves **no new issues** on the
unchanged tree.

**5a — Functional check (the linter runs and the config is honored):**

```sh
$(go env GOPATH)/bin/golangci-lint run --issues-exit-code=0 ./version/... ./payback/... ./opera/...
```

**Verify**: exits 0 and the tool prints either no findings or real lint findings
(NOT a config/tool error like "can't load config" or "unknown linter"). This
confirms the enabled linters execute against real packages with your config.
Capture whether it printed findings — that is fine; the point is the linter ran.

**5b — New-issues check (degenerate at the planned-at commit — read the caveat):**

```sh
$(go env GOPATH)/bin/golangci-lint run --new-from-rev=af41ca7 ./...
```

**Verify**: exit 0.

> **Caveat — do not over-trust 5b.** When you run this, if `af41ca7` is still the
> current `HEAD` and you have changed no `.go` files, the diff base is empty, so
> this command exits 0 *no matter what* — it does NOT prove the config works
> (that is what 5a is for). 5b only proves you did not accidentally introduce new
> Go-source findings. If it reports issues, you edited Go source — revert it (this
> plan is config-only). The real gate is exercised by the CI `only-new-issues`
> action on the first actual PR, which neither 5a nor 5b can fully reproduce
> locally; that is expected.

## Test plan

No Go unit tests apply. Verification is:

- `golangci-lint config verify` passes (Step 3).
- The functional run 5a exits 0 and shows the linter actually executing (Step 5a).
- The new-issues run 5b exits 0 (Step 5b — degenerate at planned-at; see its caveat).
- CI YAML parses (Step 4), and the action `version:` is pinned (not `latest`).
- `git status --porcelain -- .golangci.yml .github/workflows/ci.yml` shows only those two.

## Done criteria

ALL must hold:

- [ ] `.golangci.yml` exists and `golangci-lint config verify` passes.
- [ ] Step 5a (`golangci-lint run --issues-exit-code=0 ./version/... ./payback/... ./opera/...`)
      exits 0 with the linter running (not a config/tool error). [This is the real proof.]
- [ ] Step 5b (`golangci-lint run --new-from-rev=af41ca7 ./...`) exits 0.
- [ ] `.github/workflows/ci.yml` has a `lint` job (`grep -n "golangci" .github/workflows/ci.yml` → ≥1)
      and its `version:` is the pinned 2.x from Step 1, NOT `latest`
      (`grep -n "version: latest" .github/workflows/ci.yml` → 0 lines).
- [ ] CI YAML parses (Step 4) — or PyYAML is unavailable and you noted it.
- [ ] `git status --porcelain -- .golangci.yml .github/workflows/ci.yml` lists only those two.
      (An unscoped `git status` also showing `?? plans/` is expected — untracked, leave it.)
- [ ] No `.go` file is modified (`git diff --name-only af41ca7 -- '*.go'` is empty).
- [ ] `plans/README.md` status row for 002 updated (allowed index write), and you
      reported the Step 2 noise-floor count.

## STOP conditions

Stop and report back (do not improvise) if:

- `go install ...golangci-lint@latest` requires a Go newer than 1.25.8.
- `golangci-lint config verify` cannot be satisfied for the installed version
  even after minimal schema adaptation (keeping the same linter set + exclusions).
- The only way to make `--new-from-rev=af41ca7` exit 0 appears to require editing
  Go source (it should not — you changed no source).
- The live `ci.yml` structure no longer matches plan 001's "Current state" excerpt.

## Maintenance notes

- **Ratcheting plan**: once the inherited backlog (Step 2 count) is burned down
  in follow-up cleanup PRs, drop `only-new-issues: true` to gate the full tree.
  Until then, every new PR keeps its own changes clean — net-improving by
  construction.
- Adding a linter to the `enable` list will surface new issues on the next PR's
  changed lines; do that deliberately, one linter at a time.
- Reviewer should confirm the `lint` check is added to branch protection
  (GitHub settings, outside this repo).
- Keep the `gossip/contract` / `payback/contract` exclusions — regenerating
  those bindings (`gossip/sfc_test.go` `go:generate`) would otherwise reintroduce
  generated-code noise.
