#!/usr/bin/env bash
set -u -o pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
VINUCHAIN_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
WORKSPACE_DIR="$(cd "$VINUCHAIN_DIR/.." && pwd)"

QUOTA_CONTRACT_DIR="${QUOTA_CONTRACT_DIR:-$WORKSPACE_DIR/vinu-quotacontract}"
VINUCHAIN_LISTS_DIR="${VINUCHAIN_LISTS_DIR:-$WORKSPACE_DIR/vinuchain-lists}"
VINUSCAN_FRONTEND_DIR="${VINUSCAN_FRONTEND_DIR:-$WORKSPACE_DIR/vinuscan-frontend}"
VINUCHAIN_DOCS_DIR="${VINUCHAIN_DOCS_DIR:-$WORKSPACE_DIR/VinuChain-Docs}"

failures=()

run_step() {
  local label="$1"
  local dir="$2"
  shift 2

  printf '\n==> %s\n' "$label"
  printf 'dir: %s\n' "$dir"
  printf 'cmd: %s\n' "$*"

  if [ ! -d "$dir" ]; then
    printf 'missing directory: %s\n' "$dir" >&2
    failures+=("$label: missing directory $dir")
    return 1
  fi

  (
    cd "$dir" || exit 1
    "$@"
  )
  local status=$?
  if [ "$status" -ne 0 ]; then
    failures+=("$label: exit $status")
  fi
  return "$status"
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
    failures+=("$label: missing directory $dir")
    return 1
  fi

  (
    cd "$dir" || exit 1
    bash -lc "$command"
  )
  local status=$?
  if [ "$status" -ne 0 ]; then
    failures+=("$label: exit $status")
  fi
  return "$status"
}

run_shell_step \
  "VinuChain node/rules/proxy receiver readiness" \
  "$VINUCHAIN_DIR" \
  "REQUIRE_PAYBACK_RECEIVER_READY=true scripts/audit-payback-receiver-testnet.sh"

run_shell_step \
  "Quota receiver implementation VinuExplorer verification" \
  "$QUOTA_CONTRACT_DIR" \
  "REQUIRE_QUOTA_RECEIVER_READY=true npm run audit:testnet:quota"

run_shell_step \
  "Quota upgrade dispatch secret gate" \
  "$QUOTA_CONTRACT_DIR" \
  "npm run dispatch:testnet:quota-upgrade:sequence -- --dry-run"

run_shell_step \
  "Quota contract proxy upgrade and VinuExplorer verification" \
  "$QUOTA_CONTRACT_DIR" \
  "REQUIRE_QUOTA_UPGRADED=true REQUIRE_QUOTA_VERIFIED=true npm run audit:testnet:quota"

run_shell_step \
  "vinuchain-lists exact contract registry" \
  "$VINUCHAIN_LISTS_DIR" \
  "REQUIRE_QUOTA_LISTS_CURRENT=true npm run audit:vinuchain-quota"

run_step \
  "vinuscan frontend receiver readiness" \
  "$VINUSCAN_FRONTEND_DIR" \
  npm run finalize:quota-testnet

run_shell_step \
  "VinuChain docs rollout instructions" \
  "$VINUCHAIN_DOCS_DIR" \
  "set -euo pipefail
  guide=technical-docs/vinuchain-testnet/chain-upgrade-guide.md
  rg -q '0x80DA5f5e78c94EE5125Be515Ad4cd248469B57ba' \"\$guide\"
  rg -q 'dispatch:testnet:quota-upgrade' \"\$guide\"
  rg -q 'dispatch:testnet:quota-upgrade:sequence' \"\$guide\"
  rg -q 'finalize:vinuchain-quota' \"\$guide\"
  rg -q 'finalize:quota-testnet' \"\$guide\"
  rg -q 'finalize-payback-receiver-docs.sh' \"\$guide\"
  readiness=\"\$(../VinuChain/scripts/audit-payback-receiver-testnet.sh | jq -r '.receiverReadiness')\"
  if [ \"\$readiness\" = ready ]; then
    rg -q '\\*\\*Payback receiver rollout status:\\*\\* Complete' \"\$guide\"
    ! rg -q 'still points at the verified pre-receiver implementation|remaining mutating step|proxy upgrade pending|live Quota proxy upgrade is still pending|pending proxy-upgrade' \"\$guide\"
  else
    rg -q 'still points at the verified pre-receiver implementation' \"\$guide\"
    rg -q 'Quota proxy upgrade pending' \"\$guide\"
  fi"

for repo in \
  "$VINUCHAIN_DIR" \
  "$QUOTA_CONTRACT_DIR" \
  "$VINUCHAIN_LISTS_DIR" \
  "$VINUSCAN_FRONTEND_DIR" \
  "$VINUCHAIN_DOCS_DIR"; do
  run_shell_step "git clean status: $(basename "$repo")" "$repo" "test -z \"\$(git status --porcelain=v1 -uall)\""
done

printf '\n==> Rollout audit summary\n'
if [ "${#failures[@]}" -eq 0 ]; then
  printf 'All Payback receiver rollout gates passed.\n'
  exit 0
fi

printf 'Payback receiver rollout is not complete. Failed gates:\n'
for failure in "${failures[@]}"; do
  printf -- '- %s\n' "$failure"
done
exit 1
