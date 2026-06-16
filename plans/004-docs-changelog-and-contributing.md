# Plan 004: Backfill CHANGELOG and fix CONTRIBUTING / Sonar staleness

> **Executor instructions**: Follow this plan step by step. Run every
> verification command and confirm the expected result before moving on. If
> anything in "STOP conditions" occurs, stop and report — do not improvise.
> When done, update the status row for this plan in `plans/README.md`.
>
> **Drift check (run first)**: `git diff --stat af41ca7..HEAD -- CHANGELOG.md CONTRIBUTING.md sonar-project.properties`
> If any changed since this plan was written, read the live files and compare
> against the "Current state" excerpts before proceeding.

## Status

- **Priority**: P2
- **Effort**: S
- **Risk**: LOW
- **Depends on**: none
- **Category**: docs
- **Planned at**: commit `af41ca7`, 2026-06-16

## Why this matters

VinuChain ships public releases (the `vX.Y.Z-elemont` tag series, 30+ releases)
that external node operators consume — and operators *have* been bitten by
release regressions (e.g. a v2.0.14 default-bootnodes break). But the repo's
public release history is effectively invisible:

- `CHANGELOG.md` contains **only** upstream Fantom entries (`v0.3.0`, `v0.4.0`)
  and a note explicitly punting VinuChain's own history to "operational tracking
  outside this repository." A reader of the repo cannot tell what changed between
  `v2.0.30` and `v2.0.39`.
- `CONTRIBUTING.md` is the unmodified upstream template: it links to a
  goreportcard URL with literal `**github_user** / **github_repo**` placeholders
  and references a `MAINTAINERS` file that does not exist in the repo.
- `sonar-project.properties` still identifies the project as
  `Fantom-foundation_go-opera` / `go-opera`.

None of this is dangerous, but it is stale-and-wrong public-facing metadata.
This plan backfills a real changelog from the git tag history (which is the
authoritative, always-present source) and corrects the template/identity drift.

## Current state

- `CHANGELOG.md` — opens with a multi-paragraph note explaining the file holds
  only upstream-inherited history, then `## v0.4.0` and `## v0.3.0` Fantom
  sections. The note to KEEP (do not delete it) begins:

  > **Note on this file.** VinuChain is a fork of Fantom go-opera ...
  > VinuChain's own release versioning is the `Elemont` series ...

- Release tags: there are **39** elemont tags (confirm exact count with
  `git tag | grep -c elemont`). Oldest→newest they run `v1.0.0-elemont`,
  `v1.0.1-elemont`, `v1.0.2-elemont`, then `v2.0.2-elemont` … `v2.0.39-elemont`,
  **with gaps** (no `v2.0.1`, `v2.0.13`, or `v2.0.25` tag). The oldest tag is
  `v1.0.0-elemont` — NOT `v2.0.6`. Do not assume a contiguous v2-only range.
  Commits follow Conventional Commits (`feat(scope):`, `fix(scope):`, `ci:`,
  `chore(release):`), so tag-to-tag `git log` subjects are descriptive.

- `CONTRIBUTING.md:96` (the placeholder line):

  ```
  To be accepted, the PR should adhere to these quality standards (https://goreportcard.com/report/github.com/ **github_user** / **github_repo**):
  ```

- `CONTRIBUTING.md:105` (the dead reference):

  ```
  To make sure every PR is checked, we have [team maintainers](MAINTAINERS). Every PR MUST be reviewed by at least two maintainers before it can get merged.
  ```

  `ls MAINTAINERS` → no such file.

- `sonar-project.properties` (entire file):

  ```
  sonar.projectKey=Fantom-foundation_go-opera
  sonar.projectName=go-opera

  sonar.sources=.
  sonar.exclusions=**/*_test.go,**/vendor/**

  sonar.tests=.
  sonar.test.inclusions=**/*_test.go
  sonar.test.exclusions=**/vendor/**
  ```

  There is **no** Sonar scanner step in `.github/workflows/ci.yml` (CI runs only
  build/vet/test/race), so this file is currently dormant — changing the
  identity strings is safe and cosmetic.

## Commands you will need

| Purpose                              | Command | Expected |
|--------------------------------------|---------|----------|
| List elemont tags oldest→newest      | `git tag --sort=creatordate \| grep elemont` | the ordered tag list |
| Commits between two tags             | `git log --no-merges --pretty='%s' <prevTag>..<tag>` | conventional-commit subjects |
| Tag date                             | `git log -1 --format=%as <tag>` | `YYYY-MM-DD` |
| Confirm MAINTAINERS absence          | `ls MAINTAINERS 2>/dev/null \|\| echo absent` | `absent` |

## Scope

**In scope** (modify):
- `CHANGELOG.md`
- `CONTRIBUTING.md`
- `sonar-project.properties`

**Out of scope** (do NOT touch):
- Any `.go` file, CI config, or the gitignored `.claude/rules/deployment-log.md`
  (it may not exist in a clean checkout — do not depend on it; if it *is* present
  it can be used as a richer source for entry prose, but git is the source of truth).
- Do NOT create a `MAINTAINERS` file (you do not know the maintainer list);
  reword the reference instead (Step 4).
- Do NOT delete the upstream-inherited note or the Fantom `v0.3.0`/`v0.4.0`
  sections in CHANGELOG.md — they are retained for provenance.

## Git workflow

- Branch: `advisor/004-docs-hygiene`.
- Commits: one per file is fine, Conventional Commits style, e.g.
  `docs: backfill VinuChain elemont release changelog`.
- Do NOT push or open a PR unless instructed.

## Steps

### Step 1: Generate the raw per-release data

For each elemont tag from newest to oldest, collect the date and the
`feat`/`fix`/`perf`/`ci`/`refactor` commit subjects since the previous tag:

```sh
prev=""
for t in $(git tag --sort=creatordate | grep elemont); do
  if [ -n "$prev" ]; then
    echo "=== $t ($(git log -1 --format=%as "$t")) ==="
    git log --no-merges --pretty='- %s' "$prev..$t" | grep -E '^- (feat|fix|perf|ci|refactor|chore\(release\))'
  fi
  prev="$t"
done
```

This produces the source material. **The loop's `if [ -n "$prev" ]` guard skips
the OLDEST tag** (`v1.0.0-elemont`) because it has no predecessor — so the loop
prints 38 blocks, not 39. You must still write a CHANGELOG entry for that oldest
tag in Step 2: a one-line "Initial elemont release" stub (date from
`git log -1 --format=%as v1.0.0-elemont`). That brings the heading count to the
full 39.

### Step 2: Write the VinuChain release section into CHANGELOG.md

PREPEND a new top section **above** the existing upstream note (keep the note and
Fantom sections intact below). Use this shape, one block per tag, newest first:

```markdown
# Changelog

## VinuChain (Elemont series)

> Generated from the `vX.Y.Z-elemont` git tag history. Each entry lists the
> user- and operator-facing changes shipped in that release.

### v2.0.39-elemont — 2026-06-16

- fix(payback): warm PaybackCache (epochs E-1 and E) on startup ...
- chore(release): v2.0.39-elemont — PaybackCache restart warm-up

### v2.0.38-elemont — <date>

- ...

<... continue down through v2.0.2-elemont, then the three v1 tags ...>

### v1.0.0-elemont — <date>

- Initial elemont release.

---

<!-- existing upstream-inherited note and v0.4.0/v0.3.0 sections stay here, unchanged -->
```

Write **exactly one `### <tag> — <date>` heading per elemont tag, all 39**,
newest first, ending at `v1.0.0-elemont`. Note the headings are `### v1.0.0-elemont`
etc. — the version regex must match both `v1.` and `v2.` tags. Keep entries
concise — one line per meaningful commit, dropping pure version-bump noise except
the `chore(release)` line that names the release theme. You do not need to invent
prose beyond the commit subjects; faithful transcription is the goal.

**Verify**: `grep -c '^### v' CHANGELOG.md` returns **39** (matches
`git tag | grep -c elemont`; note `^### v` — NOT `^### v2\.`, which would miss the
three v1 tags), and `grep -c 'Upstream-inherited history' CHANGELOG.md` still
returns 1 (the provenance note survived).

### Step 3: Fix the CONTRIBUTING goreportcard placeholder

Replace the `**github_user** / **github_repo**` placeholder with the real path:

```
(https://goreportcard.com/report/github.com/VinuChain/VinuChain)
```

**Verify**: `grep -c 'github_user\|github_repo' CONTRIBUTING.md` → 0.

### Step 4: Fix the CONTRIBUTING MAINTAINERS reference

Reword line 105 so it does not link a non-existent file. Replace the
`[team maintainers](MAINTAINERS)` markdown link with plain text, e.g.:

```
To make sure every PR is checked, it must be reviewed by the VinuChain
maintainers. Every PR MUST be reviewed by at least two maintainers before it
can get merged.
```

**Verify**: `grep -c '](MAINTAINERS)' CONTRIBUTING.md` → 0.

### Step 5: Update the Sonar project identity

In `sonar-project.properties` set the VinuChain identity:

```
sonar.projectKey=VinuChain_VinuChain
sonar.projectName=VinuChain
```

Leave the `sonar.sources` / `sonar.exclusions` / `sonar.tests` lines unchanged.

**Verify**: `grep -c 'Fantom-foundation\|go-opera' sonar-project.properties` → 0.

## Test plan

This is a docs/metadata change — no Go tests apply. Sanity only:

- `go build ./...` → exit 0 (proves nothing structural was touched).
- The per-file `grep` verifications in Steps 2–5.
- `git status --porcelain -- CHANGELOG.md CONTRIBUTING.md sonar-project.properties`
  shows only the three in-scope files modified (untracked `plans/` is expected).

## Done criteria

ALL must hold:

- [ ] `grep -c '^### v' CHANGELOG.md` equals `git tag | grep -c elemont` (i.e. 39).
      (Use `^### v`, NOT `^### v2\.` — the latter misses the three `v1.x` tags and
      can never reach 39.)
- [ ] `grep -c 'Upstream-inherited history' CHANGELOG.md` == 1 (provenance kept).
- [ ] `grep -c 'github_user\|github_repo' CONTRIBUTING.md` == 0.
- [ ] `grep -c '](MAINTAINERS)' CONTRIBUTING.md` == 0.
- [ ] `grep -c 'Fantom-foundation\|go-opera' sonar-project.properties` == 0.
- [ ] `git status --porcelain -- CHANGELOG.md CONTRIBUTING.md sonar-project.properties`
      lists only those three. (An unscoped `git status` also showing `?? plans/` is
      expected — untracked, leave it.)
- [ ] `go build ./...` exits 0.
- [ ] `plans/README.md` status row for 004 updated (allowed index write).

## STOP conditions

Stop and report back (do not improvise) if:

- You discover a live Sonar integration after all (e.g. a CI step or a
  `.github` workflow that runs `sonar-scanner`, or a `SONAR_TOKEN` secret usage).
  Changing `sonar.projectKey` would orphan its history — confirm with the
  operator before changing the key.
- The tag history is unexpectedly empty (`git tag | grep elemont` returns
  nothing) — the checkout may be shallow; report so it can be fetched with tags.
- `git tag | grep -c elemont` does NOT return 39, or the oldest tag is not
  `v1.0.0-elemont` — the tag history has changed since this plan was written.
  Re-derive the actual tag list and count, write one heading per real tag, and
  note the discrepancy in your report rather than forcing the number 39.

## Maintenance notes

- Going forward, add a CHANGELOG entry as part of each `chore(release)` commit so
  this never goes stale again. Plan 006 (release runbook) is the natural place to
  bake that step in; if both land, cross-reference them.
- A future improvement is to automate this with a release-notes generator
  (e.g. from Conventional Commits) — explicitly out of scope here to keep the
  change small and reviewable.
- Reviewer should spot-check 2–3 CHANGELOG entries against `git log` for the
  matching tag range to confirm faithful transcription.
