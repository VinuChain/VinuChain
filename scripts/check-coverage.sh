#!/usr/bin/env bash
#
# Fail if total statement coverage drops below the regression floor.
#
# The floor is the CI coverage baseline at introduction (35.5%, measured by
# `go tool cover` over `make coverage`) minus a 1-point slack for normal
# fluctuation. Raise COVERAGE_FLOOR (or this default) deliberately as coverage
# climbs -- never lower it just to make a red build pass.
#
# Usage: scripts/check-coverage.sh [cover.prof]
set -euo pipefail

PROF="${1:-cover.prof}"
FLOOR="${COVERAGE_FLOOR:-34.5}"

if [[ ! -f "$PROF" ]]; then
  echo "check-coverage: profile '$PROF' not found (run 'make coverage' first)" >&2
  exit 2
fi

total="$(go tool cover -func="$PROF" | tail -n1 | grep -oE '[0-9]+(\.[0-9]+)?%' | tr -d '%')"
if [[ -z "$total" ]]; then
  echo "check-coverage: could not parse total coverage from '$PROF'" >&2
  exit 2
fi

awk -v t="$total" -v f="$FLOOR" 'BEGIN {
  if (t+0 < f+0) { printf "check-coverage: FAIL -- total %.1f%% is below floor %.1f%%\n", t, f; exit 1 }
  printf "check-coverage: OK -- total %.1f%% >= floor %.1f%%\n", t, f
}'
