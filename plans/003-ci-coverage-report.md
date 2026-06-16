# Plan 003: Add a non-blocking coverage report to CI

> **Executor instructions**: Follow this plan step by step. Run every
> verification command and confirm the expected result before moving on. If
> anything in "STOP conditions" occurs, stop and report — do not improvise.
> When done, update the status row for this plan in `plans/README.md`.
>
> **Drift check (run first)**: `git diff --stat af41ca7..HEAD -- .github/workflows/ci.yml Makefile`
> This plan edits `.github/workflows/ci.yml` (shared with plans 001/002). If
> those landed, the file already has `vulncheck`/`lint` jobs — add the coverage
> step/job alongside them. STOP only if `build-test`/`race` were restructured.

## Status

- **Priority**: P2
- **Effort**: S
- **Risk**: LOW
- **Depends on**: plans/001, 002 (same file — order after them or bundle)
- **Category**: tests / dx
- **Planned at**: commit `af41ca7`, 2026-06-16

## Why this matters

The repo's audit loop has manually driven up coverage on consensus-critical
packages (e.g. `vecmt` to 86%), and `CONTRIBUTING.md` states a target of
"≥80% for non-data-related packages and ≥90% for data related packages." But
**nothing measures or surfaces coverage in CI** — the `make coverage` target
exists but is never run there, and the only coverage artifact in the tree
(`cover.prof`) is a stale snapshot from 2026-03-27. Coverage gains are therefore
invisible and can silently regress.

This plan adds a **non-blocking** coverage job: it computes total coverage on
every push/PR and surfaces the number (job summary + uploaded artifact), without
failing the build. It is deliberately not a hard gate — a coverage threshold on
an inherited codebase would be arbitrary and brittle. Making the number visible
is the high-leverage, low-risk first step; a gate can come later.

## Current state

- `Makefile` `coverage` target (the canonical invocation — note it excludes the
  generated bindings and the emitter mock, which would otherwise dilute the
  number):

  ```makefile
  .PHONY: coverage
  coverage:
  	go test -coverprofile=cover.prof $$(go list ./... | grep -v '/gossip/contract/' | grep -v '/gossip/emitter/mock')
  	go tool cover -func cover.prof | tail -n 1
  ```

  The final `tail -n 1` prints a `total:	(statements)	NN.N%` line.

- `cover.prof` is **gitignored** (`.gitignore` lists `cover.prof`), so the CI
  artifact will not be accidentally committed. Good — do not change that.

- `.github/workflows/ci.yml` — jobs `build-test`, `race` (and `vulncheck`/`lint`
  if 001/002 landed). Same checkout/setup-go pattern as plan 001.

- The full test suite passes (`go test ./...` exits 0), so a coverage run will
  complete; expect it to be slower than a plain `go test` because of
  instrumentation.

## Commands you will need

| Purpose             | Command | Expected on success |
|---------------------|---------|---------------------|
| Run coverage        | `make coverage` | prints a `total:\t...\tNN.N%` line |
| Extract total only  | `go tool cover -func=cover.prof \| tail -n 1` | `total:	(statements)	NN.N%` |
| YAML sanity         | `python3 -c "import yaml; yaml.safe_load(open('.github/workflows/ci.yml'))"` | exit 0 |

## Scope

**In scope**:
- `.github/workflows/ci.yml` (add one job)

**Out of scope** (do NOT touch):
- `Makefile` — the `coverage` target already does exactly what we need; reuse it.
- Any `.go` file or test. Do NOT add/modify tests to chase a number under this plan.
- Do NOT commit `cover.prof` (it is gitignored; keep it that way).
- Do NOT introduce a coverage *threshold/gate* — non-blocking only.

## Git workflow

- Branch: `advisor/003-ci-coverage` (or continue the CI-hardening branch).
- One commit; Conventional Commits, e.g. `ci: add non-blocking coverage report`.
- Do NOT push or open a PR unless instructed.

## Steps

### Step 1: Confirm `make coverage` works locally and note the number

```sh
make coverage
```

**Verify**: the command finishes and the last line matches
`total:\s+\(statements\)\s+[0-9.]+%`. Record that percentage in your report —
it is the current baseline the CI job will start surfacing.

(This can take several minutes — it runs the whole suite with instrumentation.
That is expected.)

> Two gotchas:
> - There is a **stale** `cover.prof` already in the tree (from 2026-03-27) that
>   references source files which no longer exist. Running `go tool cover
>   -func=cover.prof` *before* `make coverage` will error with "no such file or
>   directory" — ignore it. `make coverage` rewrites `cover.prof` from scratch
>   (`-coverprofile=cover.prof`), so always run `make coverage` first.
> - This local `make coverage` run is **load-bearing**: the CI job (Step 2) uses
>   `continue-on-error: true`, so if a test only fails under coverage
>   instrumentation, CI will go green anyway. The local run here is the one place
>   that failure surfaces — see the STOP condition.

### Step 2: Add a `coverage` job to `.github/workflows/ci.yml`

Add this job under `jobs:` (same indentation as the others). It runs the
existing Make target, writes the total to the GitHub Step Summary, and uploads
the profile as an artifact — and never fails the build:

```yaml
  coverage:
    name: coverage (non-blocking)
    runs-on: ubuntu-latest
    # Informational only — never blocks merges.
    continue-on-error: true
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version-file: go.mod
          cache: false

      - name: Compute coverage
        run: make coverage

      - name: Publish total to job summary
        run: |
          echo "### Test coverage" >> "$GITHUB_STEP_SUMMARY"
          echo '```' >> "$GITHUB_STEP_SUMMARY"
          go tool cover -func=cover.prof | tail -n 1 >> "$GITHUB_STEP_SUMMARY"
          echo '```' >> "$GITHUB_STEP_SUMMARY"

      - name: Upload coverage profile
        uses: actions/upload-artifact@v4
        with:
          name: cover-prof
          path: cover.prof
          if-no-files-found: error
```

**Verify**: `python3 -c "import yaml; yaml.safe_load(open('.github/workflows/ci.yml'))"` → exit 0.

### Step 3: Confirm the job's commands reproduce locally end-to-end

```sh
make coverage && go tool cover -func=cover.prof | tail -n 1
```

**Verify**: prints the `total: ... NN.N%` line and `cover.prof` exists in the
working tree (and is gitignored — `git status --porcelain cover.prof` prints
nothing).

## Test plan

No Go unit tests apply. Verification is:

- `make coverage` produces a total percentage (Step 1).
- CI YAML parses (Step 2) — needs PyYAML; if `python3 -c "import yaml"` raises
  `ModuleNotFoundError`, skip this check and rely on CI to reject bad YAML on push.
- `cover.prof` is generated but not staged (Step 3).
- `git status --porcelain -- .github/workflows/ci.yml` shows only that file modified.

## Done criteria

ALL must hold:

- [ ] `make coverage` exits 0 and prints a `total: ... %` line.
- [ ] `.github/workflows/ci.yml` has a `coverage` job with `continue-on-error: true`
      (`grep -n "continue-on-error" .github/workflows/ci.yml` → ≥1).
- [ ] `python3 -c "import yaml; yaml.safe_load(open('.github/workflows/ci.yml'))"` exits 0.
- [ ] `git status --porcelain -- .github/workflows/ci.yml` shows that file modified,
      and `git status --porcelain -- cover.prof` prints nothing (cover.prof stays gitignored).
      (An unscoped `git status` will also show `?? plans/` — untracked and expected; do not "clean" it.)
- [ ] `plans/README.md` status row for 003 updated, with the baseline % recorded.
      (Editing this index file is expected and is the one allowed write outside the in-scope list.)

## STOP conditions

Stop and report back (do not improvise) if:

- `make coverage` fails (a test that passes under `go test ./...` but fails under
  `-coverprofile` instrumentation would be a real, separate finding worth surfacing).
- `cover.prof` is NOT gitignored in the live tree (the upload step would risk
  committing a large profile) — report rather than work around.
- The live `ci.yml` structure no longer matches plan 001's "Current state".

## Maintenance notes

- This is intentionally non-blocking. If the team later wants a floor, the
  natural next step is a small script that parses the `total:` line and fails if
  it drops below a recorded baseline — add that as a separate, deliberate change.
- The `make coverage` exclusions (`gossip/contract/`, `gossip/emitter/mock`)
  keep generated/mocked code out of the denominator; preserve them.
- If CI runtime becomes a concern, consider running the coverage job only on
  `push` to the default branch (not every PR) — it is the slowest job because of
  instrumentation.
