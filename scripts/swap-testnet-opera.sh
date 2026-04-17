#!/usr/bin/env bash
#
# swap-testnet-opera.sh — safely swap a freshly-built opera binary into a
# running testnet node's checkout and restart it.
#
# Background: On 2026-04-17 a swap run against the testnet RPC instance
# (i-0317eabd98fdf3951) targeted /home/ubuntu/VinuChain/build/ instead of
# /home/ubuntu/vinu-txtrace/build/. The trace node needs --tracenode and
# --db.preset pbl-1 to serve debug_* / trace_* RPCs; swapping the non-trace
# checkout silently removed trace capability from the public RPC for hours.
#
# This script exists so that the swap logic lives somewhere persistent and
# reviewable instead of in an ad-hoc SSM command body. Run it on the target
# instance (via SSM send-command or an interactive ubuntu shell) — not from
# the operator's workstation.
#
# Usage:
#   swap-testnet-opera.sh --target-dir <path> --expected-version <semver> [options]
#
# Required flags:
#   --target-dir <path>         Checkout whose build/opera is being swapped.
#                               Example: /home/ubuntu/vinu-txtrace
#   --expected-version <string> Version string the new binary must print, e.g.
#                               "2.0.5-elemont". Compared against `opera version`.
#
# Optional flags:
#   --source-bin <path>         Path to the newly-built binary. Defaults to
#                               <target-dir>/build/opera.new (atomic staging).
#   --skip-trace-check          Swap a non-trace checkout (mainnet RPC, validator).
#                               MUST be supplied explicitly; the default is to
#                               refuse if run_node.sh lacks --tracenode.
#   --no-restart                Stop + swap but do not restart.
#   -h, --help                  Print this help.
#
# Exit codes:
#   0   success
#   1   argument / pre-flight validation error
#   2   target sanity check failed (wrong checkout, missing --tracenode, etc.)
#   3   version mismatch
#   4   runtime / swap failure
#

set -euo pipefail

log() {
	printf '[%s] %s\n' "$(date -u +%FT%TZ)" "$*" >&2
}

die() {
	local code=$1
	shift
	log "FATAL: $*"
	exit "$code"
}

usage() {
	sed -n '3,38p' "$0" | sed 's/^# \{0,1\}//'
}

# --- Argument parsing -------------------------------------------------------

TARGET_DIR=""
EXPECTED_VERSION=""
SOURCE_BIN=""
SKIP_TRACE_CHECK=0
RESTART=1

while [[ $# -gt 0 ]]; do
	case "$1" in
		--target-dir)
			TARGET_DIR="${2:-}"
			shift 2
			;;
		--expected-version)
			EXPECTED_VERSION="${2:-}"
			shift 2
			;;
		--source-bin)
			SOURCE_BIN="${2:-}"
			shift 2
			;;
		--skip-trace-check)
			SKIP_TRACE_CHECK=1
			shift
			;;
		--no-restart)
			RESTART=0
			shift
			;;
		-h|--help)
			usage
			exit 0
			;;
		*)
			die 1 "unknown argument: $1 (see --help)"
			;;
	esac
done

[[ -n "$TARGET_DIR" ]] || die 1 "--target-dir is required (no default; state the target checkout explicitly)"
[[ -n "$EXPECTED_VERSION" ]] || die 1 "--expected-version is required (e.g. 2.0.5-elemont)"

[[ -d "$TARGET_DIR" ]] || die 1 "--target-dir does not exist: $TARGET_DIR"
TARGET_DIR="$(cd "$TARGET_DIR" && pwd)"

BUILD_DIR="$TARGET_DIR/build"
[[ -d "$BUILD_DIR" ]] || die 1 "no build/ under target-dir: $BUILD_DIR"

RUN_NODE="$BUILD_DIR/run_node.sh"
[[ -f "$RUN_NODE" ]] || die 1 "no run_node.sh under $BUILD_DIR"

CURRENT_BIN="$BUILD_DIR/opera"
[[ -x "$CURRENT_BIN" ]] || die 1 "no existing opera binary at $CURRENT_BIN (is this really a node checkout?)"

if [[ -z "$SOURCE_BIN" ]]; then
	SOURCE_BIN="$BUILD_DIR/opera.new"
fi
[[ -x "$SOURCE_BIN" ]] || die 1 "source binary missing or not executable: $SOURCE_BIN (build it first)"

# --- Identity / environment pre-flight --------------------------------------

# Refuse root. The node runs as ubuntu; swapping as root leaves files with
# wrong ownership and can confuse systemd later.
if [[ "$(id -u)" == "0" ]]; then
	die 1 "do not run as root; re-invoke as user 'ubuntu'"
fi

if [[ "$(id -un)" != "ubuntu" ]]; then
	die 1 "must run as user 'ubuntu' (got $(id -un))"
fi

log "target-dir       = $TARGET_DIR"
log "build-dir        = $BUILD_DIR"
log "run_node.sh      = $RUN_NODE"
log "current binary   = $CURRENT_BIN"
log "source binary    = $SOURCE_BIN"
log "expected version = $EXPECTED_VERSION"

# --- Checkout sanity: is this a trace node? ---------------------------------

if grep -q -- '--tracenode' "$RUN_NODE"; then
	HAS_TRACENODE=1
else
	HAS_TRACENODE=0
fi

if [[ "$HAS_TRACENODE" == "0" && "$SKIP_TRACE_CHECK" == "0" ]]; then
	die 2 "target's run_node.sh does NOT contain '--tracenode'. This is the exact failure mode of the 2026-04-17 outage. If you are intentionally swapping a non-trace checkout (mainnet RPC, validator), pass --skip-trace-check. Target: $RUN_NODE"
fi

if [[ "$HAS_TRACENODE" == "1" ]]; then
	log "OK: run_node.sh contains --tracenode"
	# Warn if the db preset expected by trace nodes is missing.
	if ! grep -q -- '--db.preset pbl-1' "$RUN_NODE"; then
		log "WARNING: --tracenode present but '--db.preset pbl-1' not found in run_node.sh; double-check this is the right checkout"
	fi
else
	log "non-trace checkout (confirmed via --skip-trace-check)"
fi

# --- Running-process sanity -------------------------------------------------

# If opera is currently running, it must be the binary under this checkout.
# Swapping while another opera instance runs from a different path is the
# exact class of error we are guarding against.
RUNNING_PIDS="$(pgrep -u ubuntu -f '[o]pera' 2>/dev/null || true)"
if [[ -n "$RUNNING_PIDS" ]]; then
	for pid in $RUNNING_PIDS; do
		exe="$(readlink -f "/proc/$pid/exe" 2>/dev/null || true)"
		if [[ -z "$exe" ]]; then
			continue
		fi
		if [[ "$exe" != "$CURRENT_BIN" ]]; then
			die 2 "an opera process (pid $pid) is running from $exe, which does not match $CURRENT_BIN. Refusing to swap — you are probably targeting the wrong checkout."
		fi
		log "running opera pid $pid → $exe (matches target) — OK"
	done
else
	log "no opera process currently running under user ubuntu"
fi

# --- Version checks ---------------------------------------------------------

# Current running binary's version — informational only.
if current_version_output="$("$CURRENT_BIN" version 2>&1)"; then
	current_version="$(printf '%s\n' "$current_version_output" | awk -F': ' '/^Version:/ {print $2; exit}')"
	log "current binary reports version: ${current_version:-<unparseable>}"
else
	log "WARNING: current binary failed to report version (may be corrupt): $current_version_output"
fi

# New binary's version — must match --expected-version exactly.
new_version_output="$("$SOURCE_BIN" version 2>&1)" || die 4 "new binary failed to run: $new_version_output"
new_version="$(printf '%s\n' "$new_version_output" | awk -F': ' '/^Version:/ {print $2; exit}')"
if [[ -z "$new_version" ]]; then
	die 3 "could not parse Version: line from new binary output: $new_version_output"
fi

if [[ "$new_version" != "$EXPECTED_VERSION" ]]; then
	die 3 "new binary version mismatch: expected '$EXPECTED_VERSION', got '$new_version'"
fi
log "OK: new binary reports version: $new_version"

# --- Stop the node ----------------------------------------------------------

if [[ -n "$RUNNING_PIDS" ]]; then
	log "sending SIGTERM to opera pids: $RUNNING_PIDS"
	kill -TERM $RUNNING_PIDS || true

	# Wait up to 60s for graceful exit.
	for _ in $(seq 1 60); do
		if ! pgrep -u ubuntu -f '[o]pera' >/dev/null 2>&1; then
			break
		fi
		sleep 1
	done

	if pgrep -u ubuntu -f '[o]pera' >/dev/null 2>&1; then
		die 4 "opera did not exit within 60s of SIGTERM; refusing to force-kill (LevelDB corruption risk). Investigate manually."
	fi
	log "opera stopped cleanly"
fi

# --- Back up and swap -------------------------------------------------------

backup_suffix="${current_version:-unknown}-$(date -u +%Y%m%dT%H%M%SZ)"
backup_path="$BUILD_DIR/opera.bak.$backup_suffix"
log "backing up current binary to $backup_path"
cp -p "$CURRENT_BIN" "$backup_path"

log "installing new binary at $CURRENT_BIN"
install -m 0755 "$SOURCE_BIN" "$CURRENT_BIN"

# Sanity: the installed binary must also report the expected version.
installed_version_output="$("$CURRENT_BIN" version 2>&1)"
installed_version="$(printf '%s\n' "$installed_version_output" | awk -F': ' '/^Version:/ {print $2; exit}')"
if [[ "$installed_version" != "$EXPECTED_VERSION" ]]; then
	log "installed binary version mismatch ($installed_version != $EXPECTED_VERSION) — rolling back"
	install -m 0755 "$backup_path" "$CURRENT_BIN"
	die 4 "rolled back to backup after version verification failure"
fi
log "OK: installed binary reports version: $installed_version"

# --- Restart ----------------------------------------------------------------

if [[ "$RESTART" == "0" ]]; then
	log "--no-restart given; leaving node stopped"
	log "swap complete; backup at $backup_path"
	exit 0
fi

log "starting node via $RUN_NODE"
cd "$BUILD_DIR"
# run_node.sh is expected to background the node (nohup / &) itself. If it
# blocks, the caller (SSM) will hang until the command times out — which is
# the caller's problem, not this script's.
bash "$RUN_NODE"

log "swap complete; backup at $backup_path"
