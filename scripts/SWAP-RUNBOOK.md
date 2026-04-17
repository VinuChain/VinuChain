# Opera Binary Swap Runbook — Testnet RPC

This runbook covers swapping the `opera` binary on the testnet RPC instance
(`i-0317eabd98fdf3951`). The same flow applies to mainnet RPC and validator
hosts with different target paths.

It pairs with `scripts/swap-testnet-opera.sh`, which enforces every check
below programmatically. Prefer the script. This document exists for when
you need to reason about what the script is doing or drop back to manual
steps during an incident.

## Why this runbook exists

On **2026-04-17 ~15:41 UTC** an SSM swap command against
`i-0317eabd98fdf3951` targeted `/home/ubuntu/VinuChain/build/` instead of
`/home/ubuntu/vinu-txtrace/build/`. The public trace RPC lost `debug_*` /
`trace_*` support until the correct checkout was identified and rebuilt.

Root cause: the swap logic lived only in ad-hoc SSM command bodies composed
in a Claude session. There was no persisted script, no checked-in runbook,
and no precondition guard against "wrong checkout". This runbook plus
`swap-testnet-opera.sh` closes that gap.

## Target layout on the testnet RPC instance

| Path                             | Role                                           | Notes                                                                    |
| -------------------------------- | ---------------------------------------------- | ------------------------------------------------------------------------ |
| `/home/ubuntu/vinu-txtrace/`     | **Trace RPC node** — serves public RPC traffic | `run_node.sh` contains `--tracenode` and `--db.preset pbl-1`. **Always the swap target on this host.** |
| `/home/ubuntu/VinuChain/`        | Secondary checkout — non-trace                 | Does not serve public traffic. Do not swap here to upgrade the RPC.      |

Verify before doing anything else:

```bash
grep -l -- --tracenode /home/ubuntu/*/build/run_node.sh
```

The path that grep prints is the trace node's checkout. That is the value
you pass as `--target-dir`.

## Pre-flight checklist

Run through this before invoking any swap command.

- [ ] You are logged in (or SSM-exec'd in) as `ubuntu`. Never as `root`.
- [ ] The instance ID matches the host you intend to operate on. Cross-check via `aws ssm describe-instance-information`.
- [ ] The new binary was **built on the instance** (never shipped from a
      workstation — glibc mismatches will silently produce a binary that
      fails to start or runs against a different libc version). See the
      "Build on instance" section below.
- [ ] The new binary's expected version string (e.g. `2.0.5-elemont`) is
      known up front. `swap-testnet-opera.sh` will refuse to proceed if
      `opera version` does not match.
- [ ] The target's `run_node.sh` still contains `--tracenode` (for the
      trace node) and `--db.preset pbl-1`.
- [ ] You know the rollback path: the previous binary stays at
      `build/opera.bak.<version>-<timestamp>` — rollback is
      `mv build/opera.bak.* build/opera && bash build/run_node.sh`.

If any box is unchecked, stop and fix that first.

## Build on instance

The task agent that kicks off a release prebuilds the new binary on the
instance. The output path convention is `<target-dir>/build/opera.new`:

```bash
# On the instance, as ubuntu
cd /home/ubuntu/vinu-txtrace
git fetch --tags origin
git checkout v2.0.5-elemont
make opera                 # produces build/opera (still the old name)
mv build/opera build/opera.new
```

Do NOT overwrite `build/opera` here — the node is still serving traffic
from that file. The swap script moves the binary into place only after the
node has stopped cleanly.

## Performing the swap (preferred path)

From an SSM send-command body on the instance:

```bash
sudo -iu ubuntu bash -lc '
cd /home/ubuntu/VinuChain-ops &&  # wherever this repo is cloned on-box
./scripts/swap-testnet-opera.sh \
  --target-dir /home/ubuntu/vinu-txtrace \
  --expected-version 2.0.5-elemont
'
```

The script will:

1. Refuse to run as root.
2. Require the target path to exist and contain `build/run_node.sh`.
3. Refuse to proceed if `run_node.sh` lacks `--tracenode` (pass
   `--skip-trace-check` only when swapping a validator or non-trace RPC).
4. Refuse if an `opera` process is running from a path that doesn't match
   the target. This is the 2026-04-17 failure mode.
5. Verify `opera.new version` exactly matches `--expected-version`.
6. Graceful SIGTERM with a 60-second budget; no force-kill.
7. Back up the current binary to `build/opera.bak.<version>-<timestamp>`.
8. Install the new binary and re-verify its reported version. If the
   installed binary somehow reports the wrong version, roll back to the
   backup and exit non-zero.
9. Restart the node via `bash build/run_node.sh`.

## Manual swap (if you cannot use the script)

Only if the script itself is broken or unavailable.

```bash
# Under user ubuntu
TARGET=/home/ubuntu/vinu-txtrace
EXPECTED=2.0.5-elemont

# 1. Sanity: is this the trace node?
grep -- --tracenode "$TARGET/build/run_node.sh" >/dev/null \
  || { echo "FATAL: target is NOT a trace checkout"; exit 1; }

# 2. Version check on the candidate binary before stopping anything.
"$TARGET/build/opera.new" version | grep -F "Version: $EXPECTED" \
  || { echo "FATAL: new binary version != $EXPECTED"; exit 1; }

# 3. Stop gracefully.
pkill -TERM -u ubuntu -f '^'"$TARGET/build/opera"'( |$)' || true
for i in $(seq 1 60); do pgrep -u ubuntu -f opera >/dev/null || break; sleep 1; done
pgrep -u ubuntu -f opera >/dev/null && { echo "FATAL: opera still running"; exit 1; }

# 4. Back up, install, verify.
cp -p "$TARGET/build/opera" "$TARGET/build/opera.bak.manual-$(date -u +%Y%m%dT%H%M%SZ)"
install -m 0755 "$TARGET/build/opera.new" "$TARGET/build/opera"
"$TARGET/build/opera" version | grep -F "Version: $EXPECTED" \
  || { echo "FATAL: installed binary version != $EXPECTED — roll back manually"; exit 1; }

# 5. Restart.
( cd "$TARGET/build" && bash run_node.sh )
```

## Post-swap verification

1. `tail -F /home/ubuntu/vinu-txtrace/build/logs/*.log` — confirm block
   production resumes within ~30 seconds.
2. `curl -s -X POST http://localhost:18545 -H 'content-type: application/json' \
   -d '{"jsonrpc":"2.0","id":1,"method":"web3_clientVersion","params":[]}'`
   — response should include the new version string.
3. Trace probe:
   `curl -s -X POST http://localhost:18545 -H 'content-type: application/json' \
   -d '{"jsonrpc":"2.0","id":1,"method":"debug_traceBlockByNumber","params":["latest",{"tracer":"callTracer"}]}'`
   — must return a result, not `method not found`. If this fails you are
   on a non-trace checkout.
4. ALB target group (`ec2-vinu-new-alb-ec2`) goes back to healthy within a
   few health check intervals.

## Rollback

Fast path (binary backup is on-disk):

```bash
cd /home/ubuntu/vinu-txtrace/build
pkill -TERM -u ubuntu -f opera && sleep 10
install -m 0755 opera.bak.<version>-<timestamp> opera
bash run_node.sh
```

If the backup file is missing, rebuild the previous tag on the instance
(same `git fetch --tags && git checkout <old-tag> && make opera` flow),
rename `build/opera` to `build/opera.new`, and re-run the swap script.

## Why each check exists

| Check                          | Incident it prevents                                                                                          |
| ------------------------------ | ------------------------------------------------------------------------------------------------------------- |
| `--target-dir` required, no default | 2026-04-17 — swap landed in the wrong checkout because the path was implicit                           |
| `--tracenode` presence guard   | Same incident — trace RPC silently dropped `debug_*`/`trace_*` support                                        |
| Running-process path match     | Operator rebuilds in checkout A while node runs from checkout B; swap finishes, node never reloaded            |
| Must run as `ubuntu`, not root | File ownership / systemd permission drift                                                                      |
| Graceful SIGTERM, no force kill | LevelDB corruption from hard-killed opera mid-write                                                           |
| Version string match before swap | Caller asked for `2.0.5-elemont` but built something else (wrong tag, uncommitted changes)                  |
| Backup before install          | One-command rollback                                                                                           |
| Version re-check after install | Catches partial `cp` / filesystem issues before the node restarts into an unknown state                       |

## Known follow-ups

- A sibling `scripts/swap-mainnet-opera.sh` (or a single script that takes
  the role as a flag) is worth adding for `i-083fffeeb03583a18`. Mainnet
  RPC has different path conventions and a larger chaindata footprint —
  it warrants its own pre-flight (disk headroom, for one).
- The same class of "ad-hoc SSM with empty `--comment`" will keep producing
  this failure mode until the tooling that composes SSM commands is
  required to include a comment and a target-path audit line.
