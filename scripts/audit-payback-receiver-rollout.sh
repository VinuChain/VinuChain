#!/usr/bin/env bash
set -u -o pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
VINUCHAIN_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
WORKSPACE_DIR="$(cd "$VINUCHAIN_DIR/.." && pwd)"

QUOTA_CONTRACT_DIR="${QUOTA_CONTRACT_DIR:-$WORKSPACE_DIR/vinu-quotacontract}"
VINUCHAIN_LISTS_DIR="${VINUCHAIN_LISTS_DIR:-$WORKSPACE_DIR/vinuchain-lists}"
VINUSCAN_FRONTEND_DIR="${VINUSCAN_FRONTEND_DIR:-$WORKSPACE_DIR/vinuscan-frontend}"
VINUCHAIN_DOCS_DIR="${VINUCHAIN_DOCS_DIR:-$WORKSPACE_DIR/VinuChain-Docs}"

if ! command -v go >/dev/null 2>&1 && [ -x /usr/local/go/bin/go ]; then
  export PATH="/usr/local/go/bin:$PATH"
fi

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
  test \"\$(git tag --contains \"\$live_commit\" --list 'v2.0.17-elemont')\" = 'v2.0.17-elemont'
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
  rg -q 'TestAddTransaction_StakeForReceiverRefundConsumesReceiverQuota' payback/payback_cache_test.go
  rg -q 'stakeFor must not record the payer as the payback stake owner' payback/payback_cache_test.go
  rg -q 'receiver payback stake sums' payback/payback_cache_test.go
  rg -q 'payer must not consume receiver payback refund quota' payback/payback_cache_test.go
  go test ./payback -run 'TestAddTransaction_(StakeRecordsSender|StakeForRecordsReceiver|StakeForReceiverRefundConsumesReceiverQuota)$' -count=1"

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
  "Quota contract receiver behavior tests" \
  "$QUOTA_CONTRACT_DIR" \
  "set -euo pipefail
  rg -q 'Should stake for another address successfully' test/Quota.test.ts
  rg -q 'Stake for another address should revert for invalid input' test/Quota.test.ts
  rg -q 'stakeFor\\(receiver.address' test/Quota.test.ts
  rg -q 'expect\\(await quota.getStake\\(payer.address\\)\\).eq\\(payerStakeBefore\\)' test/Quota.test.ts
  rg -q 'expect\\(await quota.getStake\\(receiver.address\\)\\).eq' test/Quota.test.ts
  rg -q 'quota.connect\\(payer\\).unstake\\(stake\\)' test/Quota.test.ts
  PRIVATE_TEST=0x0000000000000000000000000000000000000000000000000000000000000001 npm test -- --grep '[Ss]take for another address'"

run_shell_step \
  "Quota owner secret helper validation" \
  "$QUOTA_CONTRACT_DIR" \
  "set -euo pipefail
  helper=scripts/configure-quota-testnet-owner-secret.js
  owner_action_audit=scripts/audit-quota-testnet-owner-action.js
  aws_owner_route_helper=scripts/audit-quota-testnet-aws-owner-route.js
  tx_helper=scripts/prepare-quota-testnet-upgrade-tx.js
  sign_helper=scripts/sign-quota-testnet-upgrade-tx.js
  handoff_helper=scripts/print-quota-testnet-owner-handoff.js
  bundle_helper=scripts/package-quota-testnet-owner-handoff.js
  readme_helper=scripts/write-quota-testnet-owner-handoff-readme.js
  watch_helper=scripts/watch-quota-testnet-upgrade.js
  prepare_dispatch_helper=scripts/dispatch-quota-testnet-prepare-upgrade-tx.js
  download_prepare_helper=scripts/download-quota-testnet-prepared-tx.js
  validate_prepare_helper=scripts/validate-quota-testnet-prepared-tx.js
  export_wallet_helper=scripts/export-quota-testnet-wallet-tx.js
  wallet_page=scripts/quota-testnet-wallet-upgrade.html
  wallet_page_audit=scripts/audit-quota-testnet-wallet-upgrade-page.js
  broadcast_helper=scripts/broadcast-quota-testnet-upgrade-tx.js
  signed_dispatch_helper=scripts/dispatch-quota-testnet-signed-broadcast.js
  prepare_tx_workflow=.github/workflows/quota-testnet-prepare-upgrade-tx.yml
  signed_tx_workflow=.github/workflows/quota-testnet-broadcast-signed-tx.yml
  test -f \"\$helper\"
  test -f \"\$owner_action_audit\"
  test -f \"\$aws_owner_route_helper\"
  test -f \"\$tx_helper\"
  test -f \"\$sign_helper\"
  test -f \"\$handoff_helper\"
  test -f \"\$bundle_helper\"
  test -f \"\$readme_helper\"
  test -f \"\$watch_helper\"
  test -f \"\$prepare_dispatch_helper\"
  test -f \"\$download_prepare_helper\"
  test -f \"\$validate_prepare_helper\"
  test -f \"\$export_wallet_helper\"
  test -f \"\$wallet_page\"
  test -f \"\$wallet_page_audit\"
  test -f \"\$broadcast_helper\"
  test -f \"\$signed_dispatch_helper\"
  test -f \"\$prepare_tx_workflow\"
  test -f \"\$signed_tx_workflow\"
  node --check \"\$helper\"
  node --check \"\$owner_action_audit\"
  node --check \"\$aws_owner_route_helper\"
  node --check \"\$tx_helper\"
  node --check \"\$sign_helper\"
  node --check \"\$handoff_helper\"
  node --check \"\$bundle_helper\"
  node --check \"\$readme_helper\"
  node --check \"\$watch_helper\"
  node --check \"\$prepare_dispatch_helper\"
  node --check \"\$download_prepare_helper\"
  node --check \"\$validate_prepare_helper\"
  node --check \"\$export_wallet_helper\"
  node --check \"\$wallet_page_audit\"
  node --check \"\$broadcast_helper\"
  node --check \"\$signed_dispatch_helper\"
  rg -q 'configure:testnet:quota-upgrade-secret' package.json README.md
  rg -q 'audit:testnet:quota-aws-owner-route' package.json README.md
  rg -q 'audit:testnet:quota-owner-action' package.json README.md
  rg -q 'handoff:testnet:quota-owner' package.json README.md
  rg -q 'handoff:testnet:quota-owner-bundle' package.json README.md
  rg -q 'watch:testnet:quota-upgrade' package.json README.md
  rg -q 'README.md' README.md \"\$prepare_tx_workflow\"
  rg -q 'prepare:testnet:quota-upgrade-tx' package.json README.md
  rg -q 'sign:testnet:quota-upgrade-tx' package.json README.md
  rg -q 'dispatch:testnet:quota-prepare-upgrade-tx' package.json README.md
  rg -q 'download:testnet:quota-prepared-tx' package.json README.md
  rg -q 'audit:testnet:quota-prepared-tx' package.json README.md
  rg -q -- '--live' README.md \"\$validate_prepare_helper\"
  rg -q 'export:testnet:quota-wallet-tx' package.json README.md
  rg -q 'audit:testnet:quota-wallet-page' package.json
  rg -q 'broadcast:testnet:quota-upgrade-tx' package.json README.md
  rg -q 'dispatch:testnet:quota-signed-broadcast' package.json README.md
  rg -q -- '--skip-secret-check' README.md
  rg -q 'Check deployer secret' README.md
  rg -q 'live proxy implementation' README.md
  rg -q 'latest prepared artifact download' README.md
  rg -q 'source repository' README.md
  rg -q 'source provenance' README.md
  rg -q 'prints private keys' README.md
  rg -q -- '--ack-elevated-profile' README.md \"\$aws_owner_route_helper\"
  rg -q 'secretValuesPrinted' \"\$aws_owner_route_helper\"
  rg -q 'upgrade-tx auto' README.md
  rg -q 'Upgraded\(address\).*transaction hash' README.md
  rg -q 'PRIVATE_TEST empty in Actions' \"\$handoff_helper\"
  rg -q 'repositoryActionsSecrets' \"\$owner_action_audit\"
  rg -q 'repositoryEnvironments' \"\$owner_action_audit\"
  rg -q 'organizationActionsSecrets' \"\$owner_action_audit\"
  rg -q 'txpool_status' \"\$owner_action_audit\"
  rg -q 'ownerPendingTransactionCount' \"\$owner_action_audit\"
  rg -q 'GitHub secret surface' \"\$handoff_helper\"
  aws_owner_route_output=\"\$(npm run audit:testnet:quota-aws-owner-route -- --profile vinuchain-ops)\"
  rg -q '\"status\": \"checked_aws_owner_route\"' <<<\"\$aws_owner_route_output\"
  rg -q '\"profile\": \"vinuchain-ops\"' <<<\"\$aws_owner_route_output\"
  rg -q '\"secretValuesPrinted\": false' <<<\"\$aws_owner_route_output\"
  elevated_refusal=\"\$(npm run audit:testnet:quota-aws-owner-route -- --profile default-root 2>&1 || true)\"
  rg -q 'Refusing to use elevated-looking AWS profile \"default-root\" without --ack-elevated-profile' <<<\"\$elevated_refusal\"
  handoff_output=\"\$(npm run handoff:testnet:quota-owner)\"
  rg -q 'Local prepared files: .*quota' <<<\"\$handoff_output\"
  rg -q 'Wallet tx: .*quota-wallet-upgrade-testnet\\.json' <<<\"\$handoff_output\"
  rg -q 'Wallet sender: .*quota-testnet-wallet-upgrade\\.html' <<<\"\$handoff_output\"
  rg -q 'Live validate before signing: npm run audit:testnet:quota-prepared-tx -- --live' <<<\"\$handoff_output\"
  rg -q 'Export wallet tx: npm run export:testnet:quota-wallet-tx' <<<\"\$handoff_output\"
  rg -q 'GitHub secret surface' <<<\"\$handoff_output\"
  rg -q 'Repository PRIVATE_TEST secret present:' <<<\"\$handoff_output\"
  rg -q 'Repository environments:' <<<\"\$handoff_output\"
  rg -q 'Org Actions secrets inspectable:' <<<\"\$handoff_output\"
  rg -q 'Org secret note: Org Actions secrets require org admin or Actions-secrets fine-grained permission to inspect.' <<<\"\$handoff_output\"
  latest_prepared_run_id=\"\$(sed -n 's/^Run id: //p' <<<\"\$handoff_output\" | head -n1)\"
  prepared_source_commit=\"\$(sed -n 's/^Prepared source commit: //p' <<<\"\$handoff_output\" | head -n1)\"
  test -n \"\$latest_prepared_run_id\"
  test -n \"\$prepared_source_commit\"
  rg -q 'Transaction-affecting drift: no' <<<\"\$handoff_output\"
  rg -q 'Post-confirmation finalization' <<<\"\$handoff_output\"
  rg -q 'finalize-payback-receiver-rollout.sh --dry-run --upgrade-tx auto' <<<\"\$handoff_output\"
  rg -q 'finalize-payback-receiver-rollout.sh --commit --push --upgrade-tx auto' <<<\"\$handoff_output\"
  watch_output=\"\$(npm run watch:testnet:quota-upgrade -- --once 2>&1 || true)\"
  rg -q '\"expectedImplementation\":\"0x80DA5f5e78c94EE5125Be515Ad4cd248469B57ba\"' <<<\"\$watch_output\"
  rg -q '\"implementationHasStakeFor\":' <<<\"\$watch_output\"
  rg -q 'upgradeTransactionHash' \"\$watch_helper\"
  rg -q 'finalizerDryRunCommand' \"\$watch_helper\"
  npm run audit:testnet:quota-wallet-page
  rg -q 'eth_sendTransaction' \"\$wallet_page\"
  rg -q 'wallet_switchEthereumChain' \"\$wallet_page\"
  rg -q 'quota-wallet-upgrade-testnet.json' \"\$wallet_page\"
  bundle_output=\"\$(npm run handoff:testnet:quota-owner-bundle -- \"\$latest_prepared_run_id\" --dir /tmp/quota-owner-handoff-audit --no-zip)\"
  rg -q '\"status\": \"ready\"' <<<\"\$bundle_output\"
  test -f /tmp/quota-owner-handoff-audit/README.md
  test -f /tmp/quota-owner-handoff-audit/quota-prepared-upgrade-testnet.json
  test -f /tmp/quota-owner-handoff-audit/quota-wallet-upgrade-testnet.json
  rg -q 'Browser Wallet Path' /tmp/quota-owner-handoff-audit/README.md
  rg -q 'BUNDLE_DIR=\"\\$\\(pwd\\)\"' /tmp/quota-owner-handoff-audit/README.md
  rg -q 'cd /home/gypsey/vinu-quotacontract' /tmp/quota-owner-handoff-audit/README.md
  rg -q 'audit:testnet:quota-prepared-tx -- --live \"\\\$BUNDLE_DIR/quota-prepared-upgrade-testnet.json\"' /tmp/quota-owner-handoff-audit/README.md
  rg -q 'Post-Confirmation Finalization' /tmp/quota-owner-handoff-audit/README.md
  rg -q 'finalize-payback-receiver-rollout.sh --dry-run <upgrade-tx-hash>' /tmp/quota-owner-handoff-audit/README.md
  rg -q 'finalize-payback-receiver-rollout.sh --commit --push <upgrade-tx-hash>' /tmp/quota-owner-handoff-audit/README.md
  rg -q 'finalize-payback-receiver-rollout.sh --dry-run --upgrade-tx auto' /tmp/quota-owner-handoff-audit/README.md
  rg -q 'finalize-payback-receiver-rollout.sh --commit --push --upgrade-tx auto' /tmp/quota-owner-handoff-audit/README.md
  rg -q 'audit-payback-receiver-rollout.sh' /tmp/quota-owner-handoff-audit/README.md
  rg -q 'Prepared JSON SHA256' /tmp/quota-owner-handoff-audit/README.md
  rg -q 'Source commit' /tmp/quota-owner-handoff-audit/README.md
  rg -q 'prepared buffered gas price' /tmp/quota-owner-handoff-audit/README.md
  rg -q 'Observed gas price wei' /tmp/quota-owner-handoff-audit/README.md
  rg -q 'Gas price buffer bps' /tmp/quota-owner-handoff-audit/README.md
  rg -q 'Prepared gas price wei' /tmp/quota-owner-handoff-audit/README.md
  jq -e --arg commit \"\$prepared_source_commit\" --arg run \"\$latest_prepared_run_id\" '.sourceRepository == \"VinuChain/vinu-quotacontract\" and .sourceCommit == \$commit and .sourceRef == \"main\" and .sourceRunId == \$run and .sourceWorkflow == \"Quota Testnet Prepare Upgrade Tx\"' /tmp/quota-owner-handoff-audit/quota-prepared-upgrade-testnet.json
  jq -e '.observedGasPriceWei and .gasPriceBufferBps == 1000 and (.suggestedLegacyTransaction.gasPrice == .gasPriceWei) and ((.gasPriceWei | tonumber) > (.observedGasPriceWei | tonumber))' /tmp/quota-owner-handoff-audit/quota-prepared-upgrade-testnet.json
  jq -e --arg commit \"\$prepared_source_commit\" --arg run \"\$latest_prepared_run_id\" '.sourceRepository == \"VinuChain/vinu-quotacontract\" and .sourceCommit == \$commit and .sourceRef == \"main\" and .sourceRunId == \$run and .sourceWorkflow == \"Quota Testnet Prepare Upgrade Tx\"' /tmp/quota-owner-handoff-audit/quota-wallet-upgrade-testnet.json
  rg -q 'dispatch-dry-run' README.md \"\$signed_dispatch_helper\"
  rg -q 'suggestedLegacyTransaction' README.md
  rg -q 'buffered gas price' README.md
  rg -q 'observed RPC gas price' README.md
  rg -q 'is_fully_verified=true' README.md
  rg -q 'is_partially_verified=false' README.md
  rg -q 'local .* artifact bytecode' README.md
  rg -q 'Quota Testnet Prepare Upgrade Tx' README.md \"\$prepare_tx_workflow\"
  rg -q 'quota-prepared-upgrade-testnet' README.md \"\$prepare_tx_workflow\"
  rg -q 'quota-wallet-upgrade-testnet.json' README.md \"\$prepare_tx_workflow\"
  rg -q 'quota-testnet-wallet-upgrade.html' README.md \"\$prepare_tx_workflow\"
  rg -q 'write-quota-testnet-owner-handoff-readme.js' \"\$prepare_tx_workflow\"
  rg -q 'GitHub artifact API helper' README.md
  rg -q 'npm run download:testnet:quota-prepared-tx -- <run-id>' README.md
  rg -q 'actions/artifacts' \"\$download_prepare_helper\"
  rg -q 'suggestedLegacyTransaction' \"\$validate_prepare_helper\"
  download_help=\"\$(npm run download:testnet:quota-prepared-tx -- --help)\"
  rg -q -- '--dir' <<<\"\$download_help\"
  prepared_audit_help=\"\$(npm run audit:testnet:quota-prepared-tx -- --help)\"
  rg -q 'prepared transaction JSON file' <<<\"\$prepared_audit_help\"
  rg -q -- '--live' <<<\"\$prepared_audit_help\"
  wallet_export_help=\"\$(npm run export:testnet:quota-wallet-tx -- --help)\"
  rg -q 'wallet-friendly JSON shape' <<<\"\$wallet_export_help\"
  prepare_dispatch_dry_run=\"\$(npm run dispatch:testnet:quota-prepare-upgrade-tx -- --dry-run)\"
  rg -q 'quota-testnet-prepare-upgrade-tx.yml' <<<\"\$prepare_dispatch_dry_run\"
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
    jq -e '.observedGasPriceWei and .gasPriceBufferBps == 1000 and ((.gasPriceWei | tonumber) > (.observedGasPriceWei | tonumber))' <<<\"\$prep_json\"
    export_json=\"\$(tmpfile=\$(mktemp); printf '%s\n' \"\$prep_json\" > \"\$tmpfile\"; npm run export:testnet:quota-wallet-tx -- \"\$tmpfile\"; rm -f \"\$tmpfile\")\"
    printf '%s\n' \"\$export_json\" | sed -n '/^{/,\$p' | jq -e '.status == \"ready_for_owner_signature\" and .walletTransaction.chainId == \"0xce\" and .walletTransaction.to == \"0xcE154534e1E8F4Cc9Ab642Ad1816Ee1A237055F4\" and .quotaProxy == \"0x824B93dE7221cf8a35FBd29d5202f6eFa3A29C5D\" and .implementation == \"0x80DA5f5e78c94EE5125Be515Ad4cd248469B57ba\"'
  else
    rg -q 'Quota proxy already points at 0x80DA5f5e78c94EE5125Be515Ad4cd248469B57ba' <<<\"\$prep_output\"
  fi"

run_shell_step \
  "Payback receiver finalizer wrapper validation" \
  "$VINUCHAIN_DIR" \
  "test -x scripts/finalize-payback-receiver-rollout.sh
  test -x scripts/wait-finalize-payback-receiver-rollout.sh
  bash -n scripts/finalize-payback-receiver-rollout.sh
  bash -n scripts/wait-finalize-payback-receiver-rollout.sh
  rg -q 'audit-testnet-aws-opera.sh' scripts/finalize-payback-receiver-rollout.sh
  scripts/finalize-payback-receiver-rollout.sh --help | rg -q -- '--upgrade-tx <hash>'
  scripts/finalize-payback-receiver-rollout.sh --help | rg -q 'auto.*proxy Upgraded event'
  scripts/finalize-payback-receiver-rollout.sh --help | rg -q -- '--commit'
  scripts/finalize-payback-receiver-rollout.sh --help | rg -q -- '--push'
  scripts/wait-finalize-payback-receiver-rollout.sh --help | rg -q 'watch:testnet:quota-upgrade'
  scripts/wait-finalize-payback-receiver-rollout.sh --help | rg -q 'finalizer dry-run first'
  scripts/wait-finalize-payback-receiver-rollout.sh --help | rg -q -- '--timeout-seconds <n>'
  scripts/wait-finalize-payback-receiver-rollout.sh --help | rg -q -- '--commit'
  scripts/wait-finalize-payback-receiver-rollout.sh --help | rg -q -- '--push'
  commit_dry_run_output=\"\$(scripts/finalize-payback-receiver-rollout.sh --dry-run --commit 0x0000000000000000000000000000000000000000000000000000000000000000 2>&1 || true)\"
  rg -q -- '--dry-run cannot be combined with --commit or --push' <<<\"\$commit_dry_run_output\"
  push_without_commit_output=\"\$(scripts/finalize-payback-receiver-rollout.sh --push 0x0000000000000000000000000000000000000000000000000000000000000000 2>&1 || true)\"
  rg -q -- '--push requires --commit' <<<\"\$push_without_commit_output\"
  wait_push_without_commit_output=\"\$(scripts/wait-finalize-payback-receiver-rollout.sh --push 2>&1 || true)\"
  rg -q -- '--push requires --commit' <<<\"\$wait_push_without_commit_output\""

run_shell_step \
  "Quota upgrade dispatch secret gate" \
  "$QUOTA_CONTRACT_DIR" \
  "npm run dispatch:testnet:quota-upgrade:sequence -- --dry-run"

run_shell_step \
  "Quota owner handoff summary" \
  "$QUOTA_CONTRACT_DIR" \
  "npm run handoff:testnet:quota-owner"

run_shell_step \
  "Quota contract proxy upgrade and VinuExplorer verification" \
  "$QUOTA_CONTRACT_DIR" \
  "REQUIRE_QUOTA_UPGRADED=true REQUIRE_QUOTA_VERIFIED=true npm run audit:testnet:quota"

run_shell_step \
  "vinuchain-lists quota registry static tests" \
  "$VINUCHAIN_LISTS_DIR" \
  "npm test -- --grep 'VinuChain quota registry receiver metadata'"

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
  rg -q 'quota-wallet-upgrade-testnet.json' \"\$guide\"
  rg -q 'quota-testnet-wallet-upgrade.html' \"\$guide\"
  rg -q 'README.md' \"\$guide\"
  rg -q 'handoff:testnet:quota-owner-bundle' \"\$guide\"
  rg -q 'download:testnet:quota-prepared-tx' \"\$guide\"
  rg -q 'audit:testnet:quota-prepared-tx' \"\$guide\"
  rg -q -- '--live' \"\$guide\"
  rg -q 'export:testnet:quota-wallet-tx' \"\$guide\"
  rg -q 'GitHub artifact API helper' \"\$guide\"
  rg -q 'broadcast:testnet:quota-upgrade-tx' \"\$guide\"
  rg -q 'sign:testnet:quota-upgrade-tx' \"\$guide\"
  rg -q 'dispatch:testnet:quota-signed-broadcast' \"\$guide\"
  rg -q 'dispatch-dry-run' \"\$guide\"
  rg -q 'suggestedLegacyTransaction' \"\$guide\"
  rg -q 'buffered gas price' \"\$guide\"
  rg -q 'observed RPC gas' \"\$guide\"
  rg -q 'Quota Testnet Signed Tx Broadcast' \"\$guide\"
  rg -q 'audit-testnet-aws-opera.sh' \"\$guide\"
  rg -q 'dispatch:testnet:quota-upgrade' \"\$guide\"
  rg -q 'dispatch:testnet:quota-upgrade:sequence' \"\$guide\"
  rg -q -- '--skip-secret-check' \"\$guide\"
  rg -q 'Check deployer secret' \"\$guide\"
  rg -q 'audit:testnet:quota-owner-action' \"\$guide\"
  rg -q 'handoff:testnet:quota-owner' \"\$guide\"
  rg -q 'watch:testnet:quota-upgrade' \"\$guide\"
  rg -q 'Upgraded\(address\).*transaction hash' \"\$guide\"
  rg -q 'live proxy implementation' \"\$guide\"
  rg -q 'latest prepared artifact download' \"\$guide\"
  rg -q 'source repository' \"\$guide\"
  rg -q 'source provenance' \"\$guide\"
  rg -q 'prints private keys' \"\$guide\"
  rg -q 'finalize-payback-receiver-rollout.sh' \"\$guide\"
  rg -q 'wait-finalize-payback-receiver-rollout.sh' \"\$guide\"
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
  if [ \"\$readiness\" != ready ]; then
    quota_contract_dir=\"$QUOTA_CONTRACT_DIR\"
    handoff_output=\"\$(cd \"\$quota_contract_dir\" && npm run handoff:testnet:quota-owner)\"
    latest_prepared_run_id=\"\$(sed -n 's/^Run id: //p' <<<\"\$handoff_output\" | head -n1)\"
    latest_artifact_id=\"\$(sed -n 's/^Artifact id: //p' <<<\"\$handoff_output\" | head -n1)\"
    prepared_source_commit=\"\$(sed -n 's/^Prepared source commit: //p' <<<\"\$handoff_output\" | head -n1)\"
    test -n \"\$latest_prepared_run_id\"
    test -n \"\$latest_artifact_id\"
    test -n \"\$prepared_source_commit\"

    artifact_dir=\"\$(mktemp -d /tmp/quota-doc-artifact.XXXXXX)\"
    artifact_zip=\"\${artifact_dir}.zip\"
    trap 'rm -rf \"\$artifact_dir\" \"\$artifact_zip\"' EXIT
    (
      cd \"\$quota_contract_dir\" || exit 1
      npm run download:testnet:quota-prepared-tx -- \"\$latest_prepared_run_id\" --dir \"\$artifact_dir\"
    )
    prepared_sha=\"\$(sha256sum \"\$artifact_dir/quota-prepared-upgrade-testnet.json\" | cut -d ' ' -f1)\"
    wallet_sha=\"\$(sha256sum \"\$artifact_dir/quota-wallet-upgrade-testnet.json\" | cut -d ' ' -f1)\"
    sender_sha=\"\$(sha256sum \"\$artifact_dir/quota-testnet-wallet-upgrade.html\" | cut -d ' ' -f1)\"
    readme_sha=\"\$(sha256sum \"\$artifact_dir/README.md\" | cut -d ' ' -f1)\"
    zip_sha=\"\$(sha256sum \"\$artifact_zip\" | cut -d ' ' -f1)\"

    rg -q \"\$latest_prepared_run_id\" \"\$guide\"
    rg -q \"\$latest_artifact_id\" \"\$guide\"
    rg -q \"\$prepared_source_commit\" \"\$guide\"
    rg -q \"\$prepared_sha\" \"\$guide\"
    rg -q \"\$wallet_sha\" \"\$guide\"
    rg -q \"\$sender_sha\" \"\$guide\"
    rg -q \"\$readme_sha\" \"\$guide\"
    rg -q \"\$zip_sha\" \"\$guide\"
  fi
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
