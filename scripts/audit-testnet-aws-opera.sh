#!/usr/bin/env bash
set -euo pipefail

AWS_PROFILE="${AWS_PROFILE:-vinuchain-ops}"
AWS_REGION="${AWS_REGION:-us-west-2}"
TRACE_INSTANCE_ID="${TRACE_INSTANCE_ID:-i-0317eabd98fdf3951}"
VALIDATOR_INSTANCE_ID="${VALIDATOR_INSTANCE_ID:-i-029476269e84beb4a}"
EXPECTED_VERSION="${EXPECTED_VERSION:-2.0.19-elemont}"
EXPECTED_COMMIT="${EXPECTED_COMMIT:-$(git rev-parse --short HEAD 2>/dev/null || printf unknown)}"
TRACE_OPERA_PATH="${TRACE_OPERA_PATH:-/home/ubuntu/vinu-txtrace/build/opera}"
VALIDATOR_OPERA_PATH="${VALIDATOR_OPERA_PATH:-/home/ubuntu/vinu-testnet-prod/vinu-blockchain/build/opera}"

aws_cli() {
  AWS_PROFILE="$AWS_PROFILE" AWS_REGION="$AWS_REGION" aws "$@"
}

send_readonly_command() {
  local label="$1"
  local instance_id="$2"
  local script_body="$3"
  local tmp_script
  local encoded
  local command_id

  tmp_script="$(mktemp)"
  printf '%s\n' "$script_body" >"$tmp_script"
  encoded="$(base64 -w0 "$tmp_script")"
  rm -f "$tmp_script"

  command_id="$(
    aws_cli ssm send-command \
      --document-name VinuChain-RunAsUbuntu \
      --instance-ids "$instance_id" \
      --parameters "commands=[\"echo '$encoded' | base64 -d | bash\"]" \
      --timeout-seconds 30 \
      --comment "read-only audit: $label opera version and service state" \
      --query 'Command.CommandId' \
      --output text
  )"

  aws_cli ssm wait command-executed \
    --command-id "$command_id" \
    --instance-id "$instance_id"

  aws_cli ssm get-command-invocation \
    --command-id "$command_id" \
    --instance-id "$instance_id" \
    --query 'StandardOutputContent' \
    --output text
}

bool() {
  if "$@"; then
    printf 'true'
  else
    printf 'false'
  fi
}

require_line() {
  local output="$1"
  local expected="$2"
  if ! rg -q "^${expected}$" <<<"$output"; then
    printf 'missing expected AWS audit line: %s\n' "$expected" >&2
    return 1
  fi
}

identity_json="$(aws_cli sts get-caller-identity --output json)"
instances_json="$(aws_cli ssm describe-instance-information --output json)"

if ! jq -e --arg id "$TRACE_INSTANCE_ID" '.InstanceInformationList[] | select(.InstanceId == $id and .PingStatus == "Online")' <<<"$instances_json" >/dev/null; then
  printf 'Trace RPC SSM instance is not online: %s\n' "$TRACE_INSTANCE_ID" >&2
  exit 1
fi

if ! jq -e --arg id "$VALIDATOR_INSTANCE_ID" '.InstanceInformationList[] | select(.InstanceId == $id and .PingStatus == "Online")' <<<"$instances_json" >/dev/null; then
  printf 'Validator SSM instance is not online: %s\n' "$VALIDATOR_INSTANCE_ID" >&2
  exit 1
fi

trace_script="$(cat <<EOF
set -euo pipefail
unit='vinu-opera.service'
expected_path='$TRACE_OPERA_PATH'
expected_version='$EXPECTED_VERSION'
expected_commit='$EXPECTED_COMMIT'
active=\$(systemctl is-active "\$unit" 2>/dev/null || true)
pid=\$(systemctl show -p MainPID --value "\$unit" 2>/dev/null || true)
exe=''
has_tracenode=false
if [ -n "\$pid" ] && [ "\$pid" != 0 ]; then
  exe=\$(readlink -f "/proc/\$pid/exe" 2>/dev/null || true)
  mapfile -d '' args <"/proc/\$pid/cmdline"
  for arg in "\${args[@]}"; do
    if [ "\$arg" = '--tracenode' ]; then
      has_tracenode=true
    fi
  done
fi
version=\$("\$expected_path" version 2>&1 | tr '\\n' '|' || true)
version_ok=false
commit_ok=false
case "\$version" in
  *"Version: \$expected_version"*) version_ok=true ;;
esac
case "\$version" in
  *"Git Commit: \$expected_commit"*) commit_ok=true ;;
esac
printf 'TRACE_ACTIVE=%s\\n' "\$active"
printf 'TRACE_EXE_OK=%s\\n' "\$([ "\$exe" = "\$expected_path" ] && printf true || printf false)"
printf 'TRACE_TRACENODE_OK=%s\\n' "\$has_tracenode"
printf 'TRACE_VERSION_OK=%s\\n' "\$version_ok"
printf 'TRACE_COMMIT_OK=%s\\n' "\$commit_ok"
EOF
)"

validator_script="$(cat <<EOF
set -euo pipefail
expected_path='$VALIDATOR_OPERA_PATH'
expected_version='$EXPECTED_VERSION'
expected_commit='$EXPECTED_COMMIT'
version=\$("\$expected_path" version 2>&1 | tr '\\n' '|' || true)
version_ok=false
commit_ok=false
case "\$version" in
  *"Version: \$expected_version"*) version_ok=true ;;
esac
case "\$version" in
  *"Git Commit: \$expected_commit"*) commit_ok=true ;;
esac
for index in 1 2 3 4; do
  unit="vinu-validator-v\${index}.service"
  active=\$(systemctl is-active "\$unit" 2>/dev/null || true)
  pid=\$(systemctl show -p MainPID --value "\$unit" 2>/dev/null || true)
  exe=''
  validator_id=''
  datadir=''
  if [ -n "\$pid" ] && [ "\$pid" != 0 ]; then
    exe=\$(readlink -f "/proc/\$pid/exe" 2>/dev/null || true)
    mapfile -d '' args <"/proc/\$pid/cmdline"
    for ((i = 0; i < \${#args[@]}; i++)); do
      case "\${args[\$i]}" in
        --validator.id)
          validator_id="\${args[\$((i + 1))]:-}"
          ;;
        --datadir)
          datadir="\${args[\$((i + 1))]:-}"
          ;;
      esac
    done
  fi
  printf 'VALIDATOR_%s_ACTIVE=%s\\n' "\$index" "\$active"
  printf 'VALIDATOR_%s_EXE_OK=%s\\n' "\$index" "\$([ "\$exe" = "\$expected_path" ] && printf true || printf false)"
  printf 'VALIDATOR_%s_ID_OK=%s\\n' "\$index" "\$([ "\$validator_id" = "\$index" ] && printf true || printf false)"
  printf 'VALIDATOR_%s_DATADIR_OK=%s\\n' "\$index" "\$([[ "\$datadir" == *"datadir_opera\${index}" ]] && printf true || printf false)"
  printf 'VALIDATOR_%s_VERSION_OK=%s\\n' "\$index" "\$version_ok"
  printf 'VALIDATOR_%s_COMMIT_OK=%s\\n' "\$index" "\$commit_ok"
done
EOF
)"

trace_output="$(send_readonly_command "testnet trace RPC" "$TRACE_INSTANCE_ID" "$trace_script")"
validator_output="$(send_readonly_command "testnet validators" "$VALIDATOR_INSTANCE_ID" "$validator_script")"

require_line "$trace_output" 'TRACE_ACTIVE=active'
require_line "$trace_output" 'TRACE_EXE_OK=true'
require_line "$trace_output" 'TRACE_TRACENODE_OK=true'
require_line "$trace_output" 'TRACE_VERSION_OK=true'
require_line "$trace_output" 'TRACE_COMMIT_OK=true'

for index in 1 2 3 4; do
  require_line "$validator_output" "VALIDATOR_${index}_ACTIVE=active"
  require_line "$validator_output" "VALIDATOR_${index}_EXE_OK=true"
  require_line "$validator_output" "VALIDATOR_${index}_ID_OK=true"
  require_line "$validator_output" "VALIDATOR_${index}_DATADIR_OK=true"
  require_line "$validator_output" "VALIDATOR_${index}_VERSION_OK=true"
  require_line "$validator_output" "VALIDATOR_${index}_COMMIT_OK=true"
done

jq -n \
  --arg profile "$AWS_PROFILE" \
  --arg region "$AWS_REGION" \
  --arg identity "$(jq -r '.Arn' <<<"$identity_json")" \
  --arg traceInstance "$TRACE_INSTANCE_ID" \
  --arg validatorInstance "$VALIDATOR_INSTANCE_ID" \
  --arg expectedVersion "$EXPECTED_VERSION" \
  --arg expectedCommit "$EXPECTED_COMMIT" \
  '{
    status: "passed",
    awsProfile: $profile,
    awsRegion: $region,
    callerArn: $identity,
    traceInstance: $traceInstance,
    validatorInstance: $validatorInstance,
    expectedVersion: $expectedVersion,
    expectedCommit: $expectedCommit,
    checks: {
      traceRpc: "active expected binary with tracenode",
      validators: "v1-v4 active expected binary and datadir"
    }
  }'
