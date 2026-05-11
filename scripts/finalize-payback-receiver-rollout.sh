#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
VINUCHAIN_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
WORKSPACE_DIR="$(cd "$VINUCHAIN_DIR/.." && pwd)"

QUOTA_CONTRACT_DIR="${QUOTA_CONTRACT_DIR:-$WORKSPACE_DIR/vinu-quotacontract}"
VINUCHAIN_LISTS_DIR="${VINUCHAIN_LISTS_DIR:-$WORKSPACE_DIR/vinuchain-lists}"
VINUSCAN_FRONTEND_DIR="${VINUSCAN_FRONTEND_DIR:-$WORKSPACE_DIR/vinuscan-frontend}"
VINUCHAIN_DOCS_DIR="${VINUCHAIN_DOCS_DIR:-$WORKSPACE_DIR/VinuChain-Docs}"
RPC_URL="${RPC_URL:-https://vinufoundation-rpc.com}"
EXPECTED_QUOTA_PROXY="${EXPECTED_QUOTA_PROXY:-0x824B93dE7221cf8a35FBd29d5202f6eFa3A29C5D}"
EXPECTED_RECEIVER_IMPLEMENTATION="${EXPECTED_RECEIVER_IMPLEMENTATION:-0x80DA5f5e78c94EE5125Be515Ad4cd248469B57ba}"
UPGRADED_EVENT_TOPIC="0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b"
UPGRADE_TX="${QUOTA_UPGRADE_TX:-}"
DRY_RUN=false
COMMIT_CHANGES=false
PUSH_CHANGES=false

usage() {
  cat <<'EOF'
Usage:
  scripts/finalize-payback-receiver-rollout.sh [options] <upgrade-tx>

Options:
  --upgrade-tx <hash>  Quota proxy upgrade transaction hash. May also be set as QUOTA_UPGRADE_TX.
                       Use "auto" to discover it from the proxy Upgraded event.
  --dry-run            Run readiness checks and pass dry-run to write-producing finalizers.
  --commit             Commit generated vinuchain-lists and docs changes.
  --push               Push generated commits. Requires --commit.
  -h, --help           Show this help.

Run this after the testnet Quota proxy has been upgraded to the verified
receiver-capable implementation. The script coordinates existing repo-level
finalizers and refuses to continue if the live proxy is not receiver-ready.
EOF
}

while [ "$#" -gt 0 ]; do
  case "$1" in
    --upgrade-tx)
      UPGRADE_TX="${2:-}"
      if [ -z "$UPGRADE_TX" ]; then
        printf '%s\n' '--upgrade-tx requires a hash' >&2
        exit 1
      fi
      shift 2
      ;;
    --dry-run)
      DRY_RUN=true
      shift
      ;;
    --commit)
      COMMIT_CHANGES=true
      shift
      ;;
    --push)
      PUSH_CHANGES=true
      shift
      ;;
    -h | --help)
      usage
      exit 0
      ;;
    0x*)
      if [ -n "$UPGRADE_TX" ]; then
        printf 'upgrade transaction hash specified more than once\n' >&2
        exit 1
      fi
      UPGRADE_TX="$1"
      shift
      ;;
    *)
      printf 'unknown argument: %s\n' "$1" >&2
      usage >&2
      exit 1
      ;;
  esac
done

if [ "$UPGRADE_TX" != "auto" ] && [[ ! "$UPGRADE_TX" =~ ^0x[0-9a-fA-F]{64}$ ]]; then
  printf 'Set QUOTA_UPGRADE_TX or pass a 32-byte upgrade tx hash.\n' >&2
  exit 1
fi

if [ "$DRY_RUN" = true ] && { [ "$COMMIT_CHANGES" = true ] || [ "$PUSH_CHANGES" = true ]; }; then
  printf '%s\n' '--dry-run cannot be combined with --commit or --push.' >&2
  exit 1
fi

if [ "$PUSH_CHANGES" = true ] && [ "$COMMIT_CHANGES" != true ]; then
  printf '%s\n' '--push requires --commit.' >&2
  exit 1
fi

run_step() {
  local label="$1"
  local dir="$2"
  shift 2

  printf '\n==> %s\n' "$label"
  printf 'dir: %s\n' "$dir"
  printf 'cmd: %s\n' "$*"

  if [ ! -d "$dir" ]; then
    printf 'missing directory: %s\n' "$dir" >&2
    exit 1
  fi

  (
    cd "$dir" || exit 1
    "$@"
  )
}

run_shell_step() {
  local label="$1"
  local dir="$2"
  local command="$3"

  printf '\n==> %s\n' "$label"
  printf 'dir: %s\n' "$dir"
  printf 'cmd: %s\n' "$command"

  if [ ! -d "$dir" ]; then
    printf 'missing directory: %s\n' "$dir" >&2
    exit 1
  fi

  (
    cd "$dir" || exit 1
    bash -lc "$command"
  )
}

rpc() {
  local method="$1"
  local params="${2:-[]}"
  local payload response

  payload="$(jq -nc --arg method "$method" --argjson params "$params" \
    '{jsonrpc:"2.0",id:1,method:$method,params:$params}')"

  if ! response="$(curl -sS --fail -H 'content-type: application/json' --data "$payload" "$RPC_URL")"; then
    printf 'RPC request failed: %s\n' "$method" >&2
    exit 1
  fi

  if jq -e '.error' >/dev/null <<<"$response"; then
    printf 'RPC returned error for %s: %s\n' "$method" "$(jq -c '.error' <<<"$response")" >&2
    exit 1
  fi

  jq -c '.result' <<<"$response"
}

hex_block() {
  printf '0x%x' "$1"
}

discover_upgrade_tx() {
  local implementation_topic latest_hex latest_block from_block to_block logs log
  local tx_hash="" tx_block="" tx_log_index="" range_size=100000

  implementation_topic="0x000000000000000000000000$(printf '%s' "${EXPECTED_RECEIVER_IMPLEMENTATION#0x}" | tr '[:upper:]' '[:lower:]')"
  latest_hex="$(rpc eth_blockNumber | jq -r '.')"
  latest_block=$((16#${latest_hex#0x}))

  for ((from_block = 0; from_block <= latest_block; from_block += range_size)); do
    to_block=$((from_block + range_size - 1))
    if [ "$to_block" -gt "$latest_block" ]; then
      to_block="$latest_block"
    fi

    logs="$(rpc eth_getLogs "$(jq -nc \
      --arg address "$EXPECTED_QUOTA_PROXY" \
      --arg fromBlock "$(hex_block "$from_block")" \
      --arg toBlock "$(hex_block "$to_block")" \
      --arg topic0 "$UPGRADED_EVENT_TOPIC" \
      --arg topic1 "$implementation_topic" \
      '[{address:$address,fromBlock:$fromBlock,toBlock:$toBlock,topics:[$topic0,$topic1]}]')")"

    while IFS= read -r log; do
      [ -n "$log" ] || continue
      tx_hash="$(jq -r '.transactionHash' <<<"$log")"
      tx_block="$(jq -r '.blockNumber' <<<"$log")"
      tx_log_index="$(jq -r '.logIndex' <<<"$log")"
    done < <(jq -c '.[]' <<<"$logs")
  done

  if [ -z "$tx_hash" ]; then
    printf 'Could not find proxy Upgraded event for %s -> %s through block %s.\n' \
      "$EXPECTED_QUOTA_PROXY" "$EXPECTED_RECEIVER_IMPLEMENTATION" "$latest_block" >&2
    exit 1
  fi

  printf 'Discovered Quota proxy upgrade tx %s at block %s logIndex %s.\n' \
    "$tx_hash" "$tx_block" "$tx_log_index" >&2
  printf '%s\n' "$tx_hash"
}

require_clean_repo() {
  local dir="$1"
  local status

  if [ ! -d "$dir" ]; then
    printf 'missing directory: %s\n' "$dir" >&2
    exit 1
  fi
  status="$(git -C "$dir" status --porcelain=v1 -uall)"
  if [ -n "$status" ]; then
    printf 'Refusing to finalize with dirty repo: %s\n%s\n' "$dir" "$status" >&2
    exit 1
  fi
}

commit_repo_changes() {
  local label="$1"
  local dir="$2"
  local message="$3"
  shift 3

  run_step "stage $label changes" "$dir" git add "$@"
  if git -C "$dir" diff --cached --quiet; then
    printf '\n==> %s commit\nNo staged changes.\n' "$label"
    return
  fi

  run_step "commit $label changes" "$dir" git commit -m "$message"

  local status
  status="$(git -C "$dir" status --porcelain=v1 -uall)"
  if [ -n "$status" ]; then
    printf 'Unexpected dirty state after committing %s changes:\n%s\n' "$label" "$status" >&2
    exit 1
  fi
}

push_repo() {
  local label="$1"
  local dir="$2"

  run_step "push $label" "$dir" git push
}

dry_arg=()
if [ "$DRY_RUN" = true ]; then
  dry_arg=(--dry-run)
else
  require_clean_repo "$VINUCHAIN_DIR"
  require_clean_repo "$QUOTA_CONTRACT_DIR"
  require_clean_repo "$VINUCHAIN_LISTS_DIR"
  require_clean_repo "$VINUSCAN_FRONTEND_DIR"
  require_clean_repo "$VINUCHAIN_DOCS_DIR"
fi

run_shell_step \
  "VinuChain receiver readiness" \
  "$VINUCHAIN_DIR" \
  "REQUIRE_PAYBACK_RECEIVER_READY=true scripts/audit-payback-receiver-testnet.sh"

run_shell_step \
  "AWS testnet RPC and validator binary readiness" \
  "$VINUCHAIN_DIR" \
  "scripts/audit-testnet-aws-opera.sh"

run_shell_step \
  "Quota contract proxy and explorer audit" \
  "$QUOTA_CONTRACT_DIR" \
  "REQUIRE_QUOTA_UPGRADED=true REQUIRE_QUOTA_VERIFIED=true npm run audit:testnet:quota"

if [ "$UPGRADE_TX" = "auto" ]; then
  UPGRADE_TX="$(discover_upgrade_tx)"
fi

run_step \
  "vinuchain-lists quota finalizer" \
  "$VINUCHAIN_LISTS_DIR" \
  npm run finalize:vinuchain-quota -- "${dry_arg[@]}"

if [ "$DRY_RUN" != true ]; then
  run_step \
    "vinuchain-lists validation" \
    "$VINUCHAIN_LISTS_DIR" \
    npm run validate

  run_shell_step \
    "vinuchain-lists strict quota audit" \
    "$VINUCHAIN_LISTS_DIR" \
    "REQUIRE_QUOTA_LISTS_CURRENT=true npm run audit:vinuchain-quota"
fi

run_step \
  "vinuscan frontend finalizer" \
  "$VINUSCAN_FRONTEND_DIR" \
  npm run finalize:quota-testnet -- "${dry_arg[@]}"

run_step \
  "VinuChain docs finalizer" \
  "$VINUCHAIN_DOCS_DIR" \
  scripts/finalize-payback-receiver-docs.sh --upgrade-tx "$UPGRADE_TX" "${dry_arg[@]}"

if [ "$DRY_RUN" = true ]; then
  printf '\nDry run complete. No files were updated.\n'
else
  printf '\nPayback receiver rollout finalizers completed.\n'
  if [ "$COMMIT_CHANGES" = true ]; then
    commit_repo_changes \
      "vinuchain-lists" \
      "$VINUCHAIN_LISTS_DIR" \
      "chore(vinuchain): finalize quota receiver implementation" \
      contracts/vinuchain/info.json

    commit_repo_changes \
      "VinuChain-Docs" \
      "$VINUCHAIN_DOCS_DIR" \
      "docs(testnet): finalize payback receiver rollout" \
      technical-docs/vinuchain-testnet/chain-upgrade-guide.md
  else
    printf 'Review, commit, and push generated changes in vinuchain-lists and VinuChain-Docs, then run:\n'
  fi

  if [ "$PUSH_CHANGES" = true ]; then
    push_repo "vinuchain-lists" "$VINUCHAIN_LISTS_DIR"
    push_repo "VinuChain-Docs" "$VINUCHAIN_DOCS_DIR"
  fi

  printf '  %s/scripts/audit-payback-receiver-rollout.sh\n' "$VINUCHAIN_DIR"
fi
