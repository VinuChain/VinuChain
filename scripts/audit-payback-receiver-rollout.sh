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
  "VinuChain core receiver stake accounting" \
  "$VINUCHAIN_DIR" \
  "set -euo pipefail
  live_commit=\"\${EXPECTED_CLIENT_COMMIT:-bbc34e6}\"
  git merge-base --is-ancestor \"\$live_commit\" HEAD
  git tag --contains \"\$live_commit\" | rg -q '^v2\\.0\\.17-elemont$'
  git show \"\$live_commit\":payback/payback_cache.go | rg -q 'stakeForSelector = methodSelector\\(\"stakeFor\\(address\\)\"\\)'
  git show \"\$live_commit\":payback/payback_cache.go | rg -q 'decodeStakeForAddress\\(data\\)'
  git show \"\$live_commit\":payback/payback_cache.go | rg -q 'return TxTypeStake, stakeAddress'
  git show \"\$live_commit\":payback/payback_cache_test.go | rg -q 'TestAddTransaction_StakeForRecordsReceiver'
  git show \"\$live_commit\":payback/payback_cache_test.go | rg -q 'stakeFor must not record the payer as the payback stake owner'
  rg -q 'stakeForSelector = methodSelector\\(\"stakeFor\\(address\\)\"\\)' payback/payback_cache.go
  rg -q 'decodeStakeForAddress\\(data\\)' payback/payback_cache.go
  rg -q 'return TxTypeStake, stakeAddress' payback/payback_cache.go
  rg -q 'GetAvailablePaybackByAddress\\(msg.From\\(\\), evm\\)' evmcore/state_processor.go
  rg -q 'TestAddTransaction_StakeForRecordsReceiver' payback/payback_cache_test.go
  rg -q 'stakeFor must not record the payer as the payback stake owner' payback/payback_cache_test.go
  rg -q 'receiver payback stake sums' payback/payback_cache_test.go
  go test ./payback -run 'TestAddTransaction_(StakeRecordsSender|StakeForRecordsReceiver)$' -count=1"

run_shell_step \
  "VinuChain node/rules/proxy receiver readiness" \
  "$VINUCHAIN_DIR" \
  "REQUIRE_PAYBACK_RECEIVER_READY=true scripts/audit-payback-receiver-testnet.sh"

run_shell_step \
  "AWS testnet RPC and validator binary readiness" \
  "$VINUCHAIN_DIR" \
  "scripts/audit-testnet-aws-opera.sh"

run_shell_step \
  "Quota receiver implementation VinuExplorer verification" \
  "$QUOTA_CONTRACT_DIR" \
  "REQUIRE_QUOTA_RECEIVER_READY=true npm run audit:testnet:quota"

run_shell_step \
  "Quota owner secret helper validation" \
  "$QUOTA_CONTRACT_DIR" \
  "set -euo pipefail
  helper=scripts/configure-quota-testnet-owner-secret.js
  tx_helper=scripts/prepare-quota-testnet-upgrade-tx.js
  sign_helper=scripts/sign-quota-testnet-upgrade-tx.js
  prepare_dispatch_helper=scripts/dispatch-quota-testnet-prepare-upgrade-tx.js
  download_prepare_helper=scripts/download-quota-testnet-prepared-tx.js
  broadcast_helper=scripts/broadcast-quota-testnet-upgrade-tx.js
  signed_dispatch_helper=scripts/dispatch-quota-testnet-signed-broadcast.js
  prepare_tx_workflow=.github/workflows/quota-testnet-prepare-upgrade-tx.yml
  signed_tx_workflow=.github/workflows/quota-testnet-broadcast-signed-tx.yml
  test -f \"\$helper\"
  test -f \"\$tx_helper\"
  test -f \"\$sign_helper\"
  test -f \"\$prepare_dispatch_helper\"
  test -f \"\$download_prepare_helper\"
  test -f \"\$broadcast_helper\"
  test -f \"\$signed_dispatch_helper\"
  test -f \"\$prepare_tx_workflow\"
  test -f \"\$signed_tx_workflow\"
  node --check \"\$helper\"
  node --check \"\$tx_helper\"
  node --check \"\$sign_helper\"
  node --check \"\$prepare_dispatch_helper\"
  node --check \"\$download_prepare_helper\"
  node --check \"\$broadcast_helper\"
  node --check \"\$signed_dispatch_helper\"
  rg -q 'configure:testnet:quota-upgrade-secret' package.json README.md
  rg -q 'audit:testnet:quota-owner-action' package.json README.md
  rg -q 'prepare:testnet:quota-upgrade-tx' package.json README.md
  rg -q 'sign:testnet:quota-upgrade-tx' package.json README.md
  rg -q 'dispatch:testnet:quota-prepare-upgrade-tx' package.json README.md
  rg -q 'download:testnet:quota-prepared-tx' package.json README.md
  rg -q 'broadcast:testnet:quota-upgrade-tx' package.json README.md
  rg -q 'dispatch:testnet:quota-signed-broadcast' package.json README.md
  rg -q -- '--skip-secret-check' README.md
  rg -q 'Check deployer secret' README.md
  rg -q 'live proxy implementation' README.md
  rg -q 'latest prepared artifact download' README.md
  rg -q 'prints private keys or signed raw transaction bytes' README.md
  rg -q 'dispatch-dry-run' README.md \"\$signed_dispatch_helper\"
  rg -q 'suggestedLegacyTransaction' README.md
  rg -q 'is_fully_verified=true' README.md
  rg -q 'is_partially_verified=false' README.md
  rg -q 'local .* artifact bytecode' README.md
  rg -q 'Quota Testnet Prepare Upgrade Tx' README.md \"\$prepare_tx_workflow\"
  rg -q 'quota-prepared-upgrade-testnet' README.md \"\$prepare_tx_workflow\"
  rg -q 'GitHub artifact API helper' README.md
  rg -q 'npm run download:testnet:quota-prepared-tx -- <run-id>' README.md
  rg -q 'actions/artifacts' \"\$download_prepare_helper\"
  npm run download:testnet:quota-prepared-tx -- --help | rg -q -- '--dir'
  npm run dispatch:testnet:quota-prepare-upgrade-tx -- --dry-run | rg -q 'quota-testnet-prepare-upgrade-tx.yml'
  rg -q 'Quota Testnet Signed Tx Broadcast' README.md \"\$signed_tx_workflow\"
  rg -q 'signed_raw_transaction' \"\$signed_tx_workflow\"
  ! rg -q 'node-version: 20' .github/workflows/quota-testnet-*.yml
  ! rg -q 'actions/(checkout|setup-node|upload-artifact)@v4' .github/workflows/quota-testnet-*.yml
  rg -q 'actions/checkout@v5' .github/workflows/quota-testnet-*.yml
  rg -q 'actions/setup-node@v5' .github/workflows/quota-testnet-*.yml
  rg -q 'actions/upload-artifact@v6' .github/workflows/quota-testnet-*.yml
  test \"\$(rg -c 'node-version: 18' .github/workflows/quota-testnet-*.yml | awk -F: '{sum += \$2} END {print sum + 0}')\" -eq 4
  malformed_output=\"\$(printf 'not-a-key' | npm run configure:testnet:quota-upgrade-secret -- --stdin --dry-run 2>&1 || true)\"
  rg -q 'Private key must be a 32-byte hex string' <<<\"\$malformed_output\"
  wrong_owner_output=\"\$(printf '0x0000000000000000000000000000000000000000000000000000000000000001' | npm run configure:testnet:quota-upgrade-secret -- --stdin --dry-run 2>&1 || true)\"
  rg -q 'expected ProxyAdmin owner' <<<\"\$wrong_owner_output\"
  wrong_signer_output=\"\$(printf '0x0000000000000000000000000000000000000000000000000000000000000001' | npm run sign:testnet:quota-upgrade-tx -- --stdin 2>&1 || true)\"
  rg -q 'expected ProxyAdmin owner' <<<\"\$wrong_signer_output\"
  signed_dispatch_output=\"\$(npm run dispatch:testnet:quota-signed-broadcast -- --no-dispatch 2>&1 || true)\"
  rg -q 'Pass --stdin' <<<\"\$signed_dispatch_output\"
  wrong_signed_tx=\"\$(node <<'NODE'
const { ethers } = require('ethers')
const wallet = new ethers.Wallet('0x0000000000000000000000000000000000000000000000000000000000000001')
const iface = new ethers.utils.Interface(['function upgrade(address proxy,address implementation)'])
const data = iface.encodeFunctionData('upgrade', [
  '0x824B93dE7221cf8a35FBd29d5202f6eFa3A29C5D',
  '0x80DA5f5e78c94EE5125Be515Ad4cd248469B57ba',
])
wallet.signTransaction({
  chainId: 206,
  nonce: 0,
  to: '0xcE154534e1E8F4Cc9Ab642Ad1816Ee1A237055F4',
  value: 0,
  gasLimit: 50000,
  gasPrice: 1,
  data,
}).then(console.log)
NODE
)\"
  broadcast_output=\"\$(printf '%s' \"\$wrong_signed_tx\" | npm run broadcast:testnet:quota-upgrade-tx -- --stdin --dry-run 2>&1 || true)\"
  rg -q 'Signed transaction from is' <<<\"\$broadcast_output\"
  prep_status=0
  prep_output=\"\$(npm run prepare:testnet:quota-upgrade-tx 2>&1)\" || prep_status=\$?
  if [ \"\$prep_status\" -eq 0 ]; then
    prep_json=\"\$(printf '%s\n' \"\$prep_output\" | sed -n '/^{/,\$p')\"
    jq -e '.chainId == 206 and .from == \"0x07B4eF04b62E69aE14A715cdcae692fa7033b9a5\" and .to == \"0xcE154534e1E8F4Cc9Ab642Ad1816Ee1A237055F4\" and .quotaProxy == \"0x824B93dE7221cf8a35FBd29d5202f6eFa3A29C5D\" and .previousImplementation == \"0x0c8735bD6b3E90eaD4cdAB917474Cc6e8E58ce82\" and .implementation == \"0x80DA5f5e78c94EE5125Be515Ad4cd248469B57ba\" and (.data | startswith(\"0x99a88ec4\")) and .simulationReturn == \"0x\" and .suggestedLegacyTransaction.type == 0 and .suggestedLegacyTransaction.chainId == .chainId and .suggestedLegacyTransaction.nonce == .ownerNonce and .suggestedLegacyTransaction.to == .to and .suggestedLegacyTransaction.value == .value and .suggestedLegacyTransaction.data == .data and .suggestedLegacyTransaction.gasPrice == .gasPriceWei and ((.suggestedLegacyTransaction.gasLimit | tonumber) >= (.gasEstimate | tonumber))' <<<\"\$prep_json\"
  else
    rg -q 'Quota proxy already points at 0x80DA5f5e78c94EE5125Be515Ad4cd248469B57ba' <<<\"\$prep_output\"
  fi"

run_shell_step \
  "Payback receiver finalizer wrapper validation" \
  "$VINUCHAIN_DIR" \
  "test -x scripts/finalize-payback-receiver-rollout.sh
  bash -n scripts/finalize-payback-receiver-rollout.sh
  rg -q 'audit-testnet-aws-opera.sh' scripts/finalize-payback-receiver-rollout.sh
  scripts/finalize-payback-receiver-rollout.sh --help | rg -q -- '--commit'
  scripts/finalize-payback-receiver-rollout.sh --help | rg -q -- '--push'
  commit_dry_run_output=\"\$(scripts/finalize-payback-receiver-rollout.sh --dry-run --commit 0x0000000000000000000000000000000000000000000000000000000000000000 2>&1 || true)\"
  rg -q -- '--dry-run cannot be combined with --commit or --push' <<<\"\$commit_dry_run_output\"
  push_without_commit_output=\"\$(scripts/finalize-payback-receiver-rollout.sh --push 0x0000000000000000000000000000000000000000000000000000000000000000 2>&1 || true)\"
  rg -q -- '--push requires --commit' <<<\"\$push_without_commit_output\""

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
  rg -q 'configure:testnet:quota-upgrade-secret' \"\$guide\"
  rg -q 'prepare:testnet:quota-upgrade-tx' \"\$guide\"
  rg -q 'dispatch:testnet:quota-prepare-upgrade-tx' \"\$guide\"
  rg -q 'Quota Testnet Prepare Upgrade Tx' \"\$guide\"
  rg -q 'quota-prepared-upgrade-testnet' \"\$guide\"
  rg -q 'download:testnet:quota-prepared-tx' \"\$guide\"
  rg -q 'GitHub artifact API helper' \"\$guide\"
  rg -q 'broadcast:testnet:quota-upgrade-tx' \"\$guide\"
  rg -q 'sign:testnet:quota-upgrade-tx' \"\$guide\"
  rg -q 'dispatch:testnet:quota-signed-broadcast' \"\$guide\"
  rg -q 'dispatch-dry-run' \"\$guide\"
  rg -q 'suggestedLegacyTransaction' \"\$guide\"
  rg -q 'Quota Testnet Signed Tx Broadcast' \"\$guide\"
  rg -q 'audit-testnet-aws-opera.sh' \"\$guide\"
  rg -q 'dispatch:testnet:quota-upgrade' \"\$guide\"
  rg -q 'dispatch:testnet:quota-upgrade:sequence' \"\$guide\"
  rg -q -- '--skip-secret-check' \"\$guide\"
  rg -q 'Check deployer secret' \"\$guide\"
  rg -q 'audit:testnet:quota-owner-action' \"\$guide\"
  rg -q 'live proxy implementation' \"\$guide\"
  rg -q 'latest prepared artifact download' \"\$guide\"
  rg -q 'prints private keys or signed raw transaction bytes' \"\$guide\"
  rg -q 'finalize-payback-receiver-rollout.sh' \"\$guide\"
  rg -q -- '--commit --push' \"\$guide\"
  rg -q 'finalize:vinuchain-quota' \"\$guide\"
  rg -q 'finalize:quota-testnet' \"\$guide\"
  rg -q 'finalize-payback-receiver-docs.sh' \"\$guide\"
  rg -q 'localDateString' scripts/finalize-payback-receiver-docs.sh
  ! rg -q 'toISOString\\(\\)\\.slice\\(0, 10\\)' scripts/finalize-payback-receiver-docs.sh
  rg -q 'is_fully_verified=true' \"\$guide\"
  rg -q 'is_partially_verified=false' \"\$guide\"
  rg -q 'local .* artifact' \"\$guide\"
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
