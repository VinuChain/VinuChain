#!/usr/bin/env bash
#
# create-chaindata-snapshot.sh — Produce a VinuChain chaindata snapshot
# tarball with an embedded SNAPSHOT_INFO.txt metadata file at the root.
#
# After extracting this tarball into a datadir, operators can verify the
# snapshot was loaded correctly with a single `cat <datadir>/SNAPSHOT_INFO.txt`
# — no need to read opera logs to discover the resume block. If the file
# is missing from <datadir>/ after extraction, the tarball did NOT extract
# correctly (wrong path, wrong layout, or extraction failed silently).
#
# Caller precondition: opera is CLEANLY STOPPED on this host. This script
# does not verify — a running opera will produce a corrupt tarball because
# LevelDB files are being mutated mid-read.

set -euo pipefail

usage() {
  cat <<'EOF' >&2
Usage: create-chaindata-snapshot.sh [OPTIONS]

Required:
  --datadir DIR          Path to opera datadir (must contain chaindata/ and go-opera/)
  --output DIR           Output directory for tarball + .sha256 + .SNAPSHOT_INFO.txt
  --network NAME         mainnet | testnet | staging
  --network-id ID        e.g. 206 for testnet, 207 for mainnet
  --binary-version STR   e.g. v2.0.11-elemont
  --block N              Latest block height at snapshot time
  --epoch N              Latest epoch at snapshot time
  --source-host STR      Free-form source identifier (instance ID, hostname, etc.)

Optional:
  --genesis-file NAME    Distributed genesis filename (informational)
  --genesis-sha256 HEX   Distributed genesis file sha256
  --flags LIST           Comma-separated upgrade flags active in snapshot state
                         (default: Berlin,London,Llr,Podgorica,SfcV2,Elemont,
                          SfcV2Patch,SfcV2Patch2,SfcV2Patch3,SfcV2Patch4)
  --name BASE            Tarball base name (default: derived from network/version/UTC)
EOF
  exit 2
}

require() {
  if [[ -z "${!1:-}" ]]; then
    echo "create-chaindata-snapshot.sh: missing required --${1//_/-}" >&2
    usage
  fi
}

datadir="" output="" network="" network_id="" binary_version=""
block="" epoch="" source_host="" genesis_file="" genesis_sha256=""
flags="Berlin,London,Llr,Podgorica,SfcV2,Elemont,SfcV2Patch,SfcV2Patch2,SfcV2Patch3,SfcV2Patch4"
base_name=""

while [[ $# -gt 0 ]]; do
  case "$1" in
    --datadir)          datadir="$2"; shift 2;;
    --output)           output="$2"; shift 2;;
    --network)          network="$2"; shift 2;;
    --network-id)       network_id="$2"; shift 2;;
    --binary-version)   binary_version="$2"; shift 2;;
    --block)            block="$2"; shift 2;;
    --epoch)            epoch="$2"; shift 2;;
    --source-host)      source_host="$2"; shift 2;;
    --genesis-file)     genesis_file="$2"; shift 2;;
    --genesis-sha256)   genesis_sha256="$2"; shift 2;;
    --flags)            flags="$2"; shift 2;;
    --name)             base_name="$2"; shift 2;;
    -h|--help)          usage;;
    *) echo "unknown arg: $1" >&2; usage;;
  esac
done

for var in datadir output network network_id binary_version block epoch source_host; do
  require "$var"
done

if [[ ! -d "$datadir/chaindata" || ! -d "$datadir/go-opera" ]]; then
  echo "datadir missing chaindata/ or go-opera/: $datadir" >&2
  exit 1
fi
mkdir -p "$output"

ts_utc=$(date -u +%Y%m%dT%H%M%SZ)
ts_iso=$(date -u +%Y-%m-%dT%H:%M:%SZ)
if [[ -z "$base_name" ]]; then
  base_name="${network}-chaindata-${binary_version}-${ts_utc}-clean"
fi
info_path="$datadir/SNAPSHOT_INFO.txt"
tarball="$output/${base_name}.tar.gz"

# Emit structured-ish plain text: line-oriented, easy to grep, easy to diff.
{
  echo "# VinuChain chaindata snapshot metadata"
  echo "# Verify extraction: \`cat <datadir>/SNAPSHOT_INFO.txt\` after unpacking."
  echo "# If this file is missing from your datadir post-extract, the tarball"
  echo "# did NOT land correctly — do not start opera; re-extract into datadir root."
  echo
  echo "Network:                 $network"
  echo "NetworkID:               $network_id"
  echo "Source host:             $source_host"
  echo "Snapshot timestamp UTC:  $ts_iso"
  echo "Binary version:          $binary_version"
  echo
  echo "# Tip at snapshot time"
  echo "Latest block:            $block"
  echo "Latest epoch:            $epoch"
  echo
  echo "# Upgrade flags sealed in this snapshot state"
  IFS=',' read -ra flag_arr <<<"$flags"
  for f in "${flag_arr[@]}"; do
    printf '%-24s active\n' "${f}:"
  done
  if [[ -n "$genesis_file" ]]; then
    echo
    echo "# Genesis distribution (informational)"
    echo "Genesis filename:        $genesis_file"
    [[ -n "$genesis_sha256" ]] && echo "Genesis SHA256:          $genesis_sha256"
  fi
  echo
  echo "# Expected post-restart behavior"
  echo "# First 'New block' log line should appear at block >= $block."
  echo "# If opera resumes at a block below that, the chaindata in your datadir"
  echo "# is NOT from this snapshot — re-extract."
} > "$info_path"

# Tarball excludes: validator identity, in-process socket, shell history,
# pre-migration backups. These are either per-operator (nodekey / keystore)
# or irrelevant (shell history, chaindata.bak.*).
tar -C "$datadir" \
  --exclude='go-opera/nodekey' \
  --exclude='keystore' \
  --exclude='keystore/*' \
  --exclude='opera.ipc' \
  --exclude='go-opera/static-nodes.json' \
  --exclude='go-opera/trusted-nodes.json' \
  --exclude='chaindata.bak.*' \
  --exclude='.bash_history' \
  --exclude='history' \
  -czf "$tarball" \
  SNAPSHOT_INFO.txt chaindata go-opera

# Companion metadata for out-of-band consumers (curl before download, etc.)
cp "$info_path" "$output/${base_name}.SNAPSHOT_INFO.txt"

# SHA256 companion for integrity verification during curl-based restore.
( cd "$output" && sha256sum "${base_name}.tar.gz" > "${base_name}.tar.gz.sha256" )

echo "Wrote:"
echo "  $tarball"
echo "  $output/${base_name}.tar.gz.sha256"
echo "  $output/${base_name}.SNAPSHOT_INFO.txt"
sha256sum "$tarball"
