#!/usr/bin/env bash
set -euo pipefail

cat >&2 <<'EOF'
wait-finalize-payback-receiver-rollout.sh is retired.

The old watcher waited for a proxy Upgraded event to the receiver-capable
implementation. Testnet now uses the PaybackV2Patch binary rule to rebind
Economy.QuotaCacheAddress to corrected non-proxy QuotaContractV2.

Use scripts/audit-payback-receiver-testnet.sh after rollout instead.
EOF

exit 1
