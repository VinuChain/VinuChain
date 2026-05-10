#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
VINUCHAIN_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
WORKSPACE_DIR="$(cd "$VINUCHAIN_DIR/.." && pwd)"

QUOTA_CONTRACT_DIR="${QUOTA_CONTRACT_DIR:-$WORKSPACE_DIR/vinu-quotacontract}"
INTERVAL_SECONDS=15
TIMEOUT_SECONDS=0
COMMIT_CHANGES=false
PUSH_CHANGES=false
RUN_FINAL_AUDIT=true

usage() {
  cat <<'EOF'
Usage:
  scripts/wait-finalize-payback-receiver-rollout.sh [options]

Options:
  --interval-seconds <n>  Poll interval for the quota upgrade watcher. Default: 15.
  --timeout-seconds <n>   Stop waiting after this many seconds. Default: 0 (no timeout).
  --commit                After dry-run succeeds, finalize and commit generated changes.
  --push                  Push generated commits. Requires --commit.
  --no-final-audit        Skip the final strict rollout audit after --commit.
  -h, --help              Show this help.

Uses npm run watch:testnet:quota-upgrade to wait for the testnet Quota proxy to
point at the verified receiver-capable implementation, runs the finalizer
dry-run first, then optionally runs the commit/push finalizer and strict rollout
audit. This does not submit the ProxyAdmin owner transaction.
EOF
}

parse_non_negative_int() {
  local value="$1"
  local label="$2"
  if [[ ! "$value" =~ ^(0|[1-9][0-9]*)$ ]]; then
    printf '%s requires a non-negative integer.\n' "$label" >&2
    exit 1
  fi
}

parse_positive_int() {
  local value="$1"
  local label="$2"
  if [[ ! "$value" =~ ^[1-9][0-9]*$ ]]; then
    printf '%s requires a positive integer.\n' "$label" >&2
    exit 1
  fi
}

while [ "$#" -gt 0 ]; do
  case "$1" in
    --interval-seconds)
      INTERVAL_SECONDS="${2:-}"
      parse_positive_int "$INTERVAL_SECONDS" "$1"
      shift 2
      ;;
    --timeout-seconds)
      TIMEOUT_SECONDS="${2:-}"
      parse_non_negative_int "$TIMEOUT_SECONDS" "$1"
      shift 2
      ;;
    --commit)
      COMMIT_CHANGES=true
      shift
      ;;
    --push)
      PUSH_CHANGES=true
      shift
      ;;
    --no-final-audit)
      RUN_FINAL_AUDIT=false
      shift
      ;;
    -h | --help)
      usage
      exit 0
      ;;
    *)
      printf 'unknown argument: %s\n' "$1" >&2
      usage >&2
      exit 1
      ;;
  esac
done

if [ "$PUSH_CHANGES" = true ] && [ "$COMMIT_CHANGES" != true ]; then
  printf '%s\n' '--push requires --commit.' >&2
  exit 1
fi

if [ ! -d "$QUOTA_CONTRACT_DIR" ]; then
  printf 'missing quota contract checkout: %s\n' "$QUOTA_CONTRACT_DIR" >&2
  exit 1
fi

watch_args=(--interval-seconds "$INTERVAL_SECONDS")
if [ "$TIMEOUT_SECONDS" != 0 ]; then
  watch_args+=(--timeout-seconds "$TIMEOUT_SECONDS")
fi

printf 'Waiting for Quota proxy receiver upgrade...\n'
printf 'dir: %s\n' "$QUOTA_CONTRACT_DIR"
printf 'cmd: npm run watch:testnet:quota-upgrade -- %s\n' "${watch_args[*]}"

set +e
watch_output="$(
  cd "$QUOTA_CONTRACT_DIR" || exit 1
  npm run watch:testnet:quota-upgrade -- "${watch_args[@]}" 2>&1
)"
watch_status=$?
set -e

printf '%s\n' "$watch_output"
if [ "$watch_status" -ne 0 ]; then
  exit "$watch_status"
fi

ready_json="$(awk '/^\{/{line=$0} END{print line}' <<<"$watch_output")"
if [ -z "$ready_json" ]; then
  printf 'watcher did not print a JSON readiness result.\n' >&2
  exit 1
fi

if [ "$(jq -r '.ready // false' <<<"$ready_json")" != true ]; then
  printf 'watcher exited successfully without ready=true.\n' >&2
  exit 1
fi

upgrade_tx="$(jq -r '.upgradeTransactionHash // "auto"' <<<"$ready_json")"
finalizer_args=(--upgrade-tx "$upgrade_tx")

printf '\nRunning finalizer dry-run first...\n'
"$VINUCHAIN_DIR/scripts/finalize-payback-receiver-rollout.sh" \
  --dry-run "${finalizer_args[@]}"

if [ "$COMMIT_CHANGES" != true ]; then
  printf '\nDry-run finalization completed. To finalize generated repo changes, run:\n'
  printf '  %s/scripts/wait-finalize-payback-receiver-rollout.sh --commit --push\n' \
    "$VINUCHAIN_DIR"
  exit 0
fi

commit_args=(--commit)
if [ "$PUSH_CHANGES" = true ]; then
  commit_args+=(--push)
fi

printf '\nRunning committing finalizer...\n'
"$VINUCHAIN_DIR/scripts/finalize-payback-receiver-rollout.sh" \
  "${commit_args[@]}" "${finalizer_args[@]}"

if [ "$RUN_FINAL_AUDIT" = true ]; then
  printf '\nRunning strict rollout audit...\n'
  "$VINUCHAIN_DIR/scripts/audit-payback-receiver-rollout.sh"
fi
