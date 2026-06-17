#!/usr/bin/env bash
#
# Keep the "VinuChain (Elemont series)" section of CHANGELOG.md current.
#
# For every vX.Y.Z-elemont release tag that does NOT already have a
# "### <tag> — <date>" heading, generate an entry from the git history since
# the previous tag and insert it at the top of the section (newest first).
# Idempotent: with no new tags it is a no-op. Existing (possibly hand-curated)
# entries are never modified.
#
#   scripts/gen-changelog.sh           update CHANGELOG.md in place
#   scripts/gen-changelog.sh --check    exit 1 if any tag lacks an entry (CI/dev)
set -euo pipefail

CHANGELOG="${CHANGELOG:-CHANGELOG.md}"
DASH=$'—'            # em-dash used in "### <tag> — <date>"
MODE="${1:-update}"

mapfile -t desc < <(git tag --sort=-creatordate | grep -E '^v[0-9].*-elemont$' || true)
mapfile -t asc  < <(git tag --sort=creatordate  | grep -E '^v[0-9].*-elemont$' || true)

if (( ${#desc[@]} == 0 )); then
  echo "gen-changelog: no *-elemont tags found (shallow clone? fetch tags)" >&2
  exit 2
fi

missing=()
for t in "${desc[@]}"; do
  grep -q "^### ${t} " "$CHANGELOG" || missing+=("$t")
done

if [[ "$MODE" == "--check" ]]; then
  if (( ${#missing[@]} )); then
    echo "gen-changelog: CHANGELOG.md missing entries for: ${missing[*]}" >&2
    exit 1
  fi
  echo "gen-changelog: up to date (${#desc[@]} elemont tags all present)"
  exit 0
fi

if (( ${#missing[@]} == 0 )); then
  echo "gen-changelog: CHANGELOG.md is up to date; nothing to add."
  exit 0
fi

prev_of() { local cur="$1" prev=""; local x; for x in "${asc[@]}"; do [[ "$x" == "$cur" ]] && { printf '%s' "$prev"; return; }; prev="$x"; done; }

block=""
for t in "${desc[@]}"; do                 # newest-first so inserted block is ordered
  printf '%s\n' "${missing[@]}" | grep -qx "$t" || continue
  date="$(git log -1 --format=%as "$t")"
  prev="$(prev_of "$t")"
  if [[ -n "$prev" ]]; then
    subs="$(git log --no-merges --pretty='- %s' "${prev}..${t}" | grep -E '^- (feat|fix|perf|ci|refactor|chore\(release\))' || true)"
  else
    subs="- Initial elemont release."
  fi
  [[ -z "$subs" ]] && subs="- (no user-facing changes)"
  block+="### ${t} ${DASH} ${date}"$'\n\n'"${subs}"$'\n\n'
done

CHANGELOG="$CHANGELOG" BLOCK="$block" python3 - <<'PY'
import os, re
p = os.environ["CHANGELOG"]; block = os.environ["BLOCK"]
s = open(p, encoding="utf-8").read()
m = re.search(r'^### ', s, flags=re.M)
assert m, "no existing '### ' heading to anchor insertion"
i = m.start()
open(p, "w", encoding="utf-8").write(s[:i] + block + s[i:])
PY
echo "gen-changelog: added ${#missing[@]} entr(y/ies): ${missing[*]}"
