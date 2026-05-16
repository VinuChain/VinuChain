#!/usr/bin/env bash
set -euo pipefail

cat >&2 <<'EOF'
scripts/audit-payback-receiver-rollout.sh has been retired.

The legacy receiver rollout audit described the superseded proxy rollout where
stakeFor(address) made the receiver the stake owner. The corrected PaybackV2
testnet upgrade uses staker-owned withdrawals and is audited with:

  scripts/audit-payback-receiver-testnet.sh
EOF

exit 1
