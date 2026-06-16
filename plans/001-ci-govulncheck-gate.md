# Plan 001: Add a `govulncheck` supply-chain gate to CI

> **Executor instructions**: Follow this plan step by step. Run every
> verification command and confirm the expected result before moving to the
> next step. If anything in the "STOP conditions" section occurs, stop and
> report — do not improvise. When done, update the status row for this plan in
> `plans/README.md`.
>
> **Drift check (run first)**: `git diff --stat af41ca7..HEAD -- .github/workflows/ci.yml Makefile`
> If `.github/workflows/ci.yml` changed since this plan was written, read the
> live file and compare against the "Current state" excerpt below before
> proceeding; on a structural mismatch, treat it as a STOP condition.

## Status

- **Priority**: P1
- **Effort**: S
- **Risk**: LOW
- **Depends on**: none
- **Category**: dx / security
- **Planned at**: commit `af41ca7`, 2026-06-16

## Why this matters

VinuChain depends on a large transitive module graph that includes known-vulnerable
modules (a `govulncheck` run during recon found **21 vulnerabilities in required
modules, 0 of which the compiled binary actually calls**). The team currently
manages this risk by hand — e.g. `go.mod` carries a `replace
github.com/sirupsen/logrus v1.2.0 => github.com/sirupsen/logrus v1.8.3` directive
with a comment explaining that nothing in the binary imports the vulnerable
version. That manual,
call-path-aware analysis is exactly what `govulncheck` automates. Without it in
CI, a future change that starts *calling* one of those 21 already-present vulns
would ship unflagged, and Dependabot alone (which flags module presence, not
call paths) produces noise that trains maintainers to ignore it.

This plan adds a `govulncheck` job that fails CI only when vulnerable code is
**actually reachable** — the signal the maintainers already act on, made
automatic. It touches no Go source.

## Current state

- `.github/workflows/ci.yml` — the entire CI definition. Two jobs today:
  `build-test` (checkout → setup-go → `go build ./...` → `go vet ./...` →
  `make opera` → `go test ./...`) and `race` (checkout → setup-go →
  `go test -race` on fork-delta packages). There is **no** vulnerability,
  lint, fuzz, or coverage job.

  The top of the file (match this style — `runs-on: ubuntu-latest`,
  `actions/checkout@v4`, `actions/setup-go@v5` with `go-version-file: go.mod`,
  `cache: false`):

  ```yaml
  name: CI

  on:
    push:
    pull_request:

  concurrency:
    group: ci-${{ github.ref }}
    cancel-in-progress: true

  permissions:
    contents: read

  jobs:
    build-test:
      name: build + vet + test
      runs-on: ubuntu-latest
      steps:
        - name: Checkout
          uses: actions/checkout@v4
        - name: Set up Go
          uses: actions/setup-go@v5
          with:
            go-version-file: go.mod
            cache: false
        - name: go build ./...
          run: go build ./...
        ...
  ```

- `Makefile` — has `.PHONY` targets `all`, `opera`, `opera-image`, `test`,
  `coverage`, `fuzz`, `clean`. There is no `vulncheck` target. The `coverage`
  target shows the existing convention for invoking Go tooling:

  ```makefile
  .PHONY: coverage
  coverage:
  	go test -coverprofile=cover.prof $$(go list ./... | grep -v '/gossip/contract/' | grep -v '/gossip/emitter/mock')
  	go tool cover -func cover.prof | tail -n 1
  ```

- `go.mod` — module path `github.com/Fantom-foundation/go-opera`, `go 1.25.8`.
  The `logrus` replace directive documents the manual CVE-triage habit this
  plan automates.

## Commands you will need

| Purpose            | Command                                              | Expected on success |
|--------------------|------------------------------------------------------|---------------------|
| Run govulncheck    | `go run golang.org/x/vuln/cmd/govulncheck@latest ./...` | ends (on stdout) with "Your code is affected by 0 vulnerabilities." |
| YAML sanity (opt.) | `python3 -c "import yaml; yaml.safe_load(open('.github/workflows/ci.yml'))"` | exits 0, no output |
| Build still OK     | `go build ./...`                                     | exit 0 |

> Notes:
> - `govulncheck@latest` resolves a tool that requires Go ≥ 1.25; the repo is on
>   1.25.8 so this is fine. On the first run you will see `go: downloading ...`
>   and possibly a `switching to go1.25.x` line on **stderr** — that is normal
>   toolchain/DB fetching, not a failure. The "0 vulnerabilities" string lands on
>   **stdout**. The tool needs network access (the existing CI already fetches
>   modules).
> - The YAML sanity command needs PyYAML (`python3 -c "import yaml"`). If that
>   import raises `ModuleNotFoundError`, PyYAML is not installed in your
>   environment — skip the YAML-sanity checks (they are marked optional) and rely
>   on CI itself to reject malformed YAML on push. Do NOT treat a missing-PyYAML
>   error as a plan failure.

## Scope

**In scope** (the only files you should modify):
- `.github/workflows/ci.yml` (add one job)
- `Makefile` (add one convenience target — optional but recommended)

**Out of scope** (do NOT touch):
- Any `.go` file. If `govulncheck` reports a *called* vulnerability, do NOT fix
  the code under this plan — STOP and report (see STOP conditions); fixing
  reachable vulns is a separate, reviewed change.
- `go.mod` / `go.sum` — do not bump dependencies here.

## Git workflow

- Branch: `advisor/001-ci-govulncheck` (or your environment's convention).
- One commit; message style is Conventional Commits with a scope (see `git log`
  — e.g. `ci: add govulncheck supply-chain gate`).
- Do NOT push or open a PR unless the operator instructed it.

## Steps

### Step 1: Confirm the gate passes locally before wiring CI

Run the tool exactly as CI will:

**Verify**: `go run golang.org/x/vuln/cmd/govulncheck@latest ./...`
→ output ends with `Your code is affected by 0 vulnerabilities.`
(It will also note "N vulnerabilities in modules you require, but your code
doesn't appear to call these" — that is informational and must NOT fail the
build.)

If instead it reports `Your code is affected by N vulnerabilities` with N>0 and
a `Vulnerability #...` call-stack section, **STOP** — the baseline is not clean
and wiring a hard gate would immediately break CI. Report the finding.

### Step 2: Add a `vulncheck` job to `.github/workflows/ci.yml`

Add a third job, parallel to `build-test` and `race`, using the same
checkout/setup-go pattern. Append this job under the existing `jobs:` map
(keep two-space indentation consistent with the file):

```yaml
  vulncheck:
    name: govulncheck
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version-file: go.mod
          cache: false

      # Call-path-aware vulnerability scan. Fails ONLY when a known-vulnerable
      # symbol is actually reachable from this module's code — not merely
      # present in the dependency graph. This automates the manual CVE triage
      # documented by the logrus replace directive in go.mod.
      - name: govulncheck
        run: go run golang.org/x/vuln/cmd/govulncheck@latest ./...
```

**Verify**: `python3 -c "import yaml; yaml.safe_load(open('.github/workflows/ci.yml'))"`
→ exit 0 (file is valid YAML and the new job parses).

### Step 3 (recommended): Add a `make vulncheck` convenience target

So contributors can run the same check locally. Add to `Makefile`:

```makefile
.PHONY: vulncheck
vulncheck:
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
```

**Verify**: `make vulncheck` → ends with `Your code is affected by 0 vulnerabilities.`

## Test plan

No Go unit tests apply (this is CI/tooling config). Verification is the
commands above plus a final confirmation that nothing else broke:

- `go build ./...` → exit 0 (unchanged source).
- The YAML parses (Step 2 verify).
- `git status --porcelain -- .github/workflows/ci.yml Makefile` shows only those
  two files modified. (An unscoped `git status` will ALSO show `?? plans/` —
  that directory is untracked and expected; leave it alone.)

## Done criteria

Machine-checkable. ALL must hold:

- [ ] `go run golang.org/x/vuln/cmd/govulncheck@latest ./...` exits 0 with
      "0 vulnerabilities" in called code.
- [ ] `.github/workflows/ci.yml` contains a `vulncheck` job (`grep -n "govulncheck" .github/workflows/ci.yml` returns ≥1 line).
- [ ] `python3 -c "import yaml; yaml.safe_load(open('.github/workflows/ci.yml'))"` exits 0.
- [ ] `Makefile` has a `vulncheck` target (if Step 3 done): `grep -n "^vulncheck:" Makefile` returns 1 line.
- [ ] `git status --porcelain -- .github/workflows/ci.yml Makefile` lists only those
      two files. (An unscoped `git status` additionally showing `?? plans/` is expected —
      `plans/` is untracked; do not remove it.)
- [ ] `plans/README.md` status row for 001 updated. (Editing this index file is
      expected and is the one allowed write outside the in-scope source list above.)

## STOP conditions

Stop and report back (do not improvise) if:

- Step 1 reports any **called** vulnerability (N>0 affecting your code). Fixing
  it means a dependency bump or source change that needs review — out of scope here.
- The live `.github/workflows/ci.yml` no longer matches the "Current state"
  excerpt structurally (e.g. jobs renamed, `setup-go` replaced).
- `go run ...govulncheck@latest` fails to fetch the tool or the vuln DB (network
  policy on the runner) — report so the operator can decide on a pinned/vendored
  alternative instead.

## Maintenance notes

- If a future dependency change makes a graph vuln *reachable*, this job turns
  red — that is the intended signal. The fix is a targeted dependency bump (or a
  `replace`, matching the logrus precedent), not disabling the gate.
- `govulncheck@latest` floats the tool version. If reproducibility becomes a
  concern, pin it (e.g. `@v1.3.0`) — but `latest` keeps the vuln-DB client
  current, which is usually what you want for a security gate.
- Reviewer should confirm the job name appears in branch-protection required
  checks (an operator/GitHub-settings step, outside this repo).
