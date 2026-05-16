#!/usr/bin/env bash
set -euo pipefail

cat >&2 <<'EOF'
finalize-payback-receiver-rollout.sh is retired.

The proxy receiver rollout was superseded by the PaybackV2 binary-level
contract replacement and the 2026-05-16 PaybackV2Patch corrected-contract
rebind. Do not run the old proxy finalizer; it targets the stale
QuotaContractReceiverImplementation path.

Use scripts/audit-payback-receiver-testnet.sh after the PaybackV2Patch rollout
to verify that vc_getRules points at the corrected QuotaContractV2 address.
EOF

exit 1
