#!/usr/bin/env bash
set -euo pipefail

RPC_URL="${RPC_URL:-https://vinufoundation-rpc.com}"
EXPECTED_CHAIN_ID="${EXPECTED_CHAIN_ID:-0xce}"
EXPECTED_CLIENT_VERSION="${EXPECTED_CLIENT_VERSION:-v2.0.19-elemont}"
EXPECTED_CLIENT_COMMIT="${EXPECTED_CLIENT_COMMIT:-}"
EXPECTED_PAYBACK_V2="${EXPECTED_PAYBACK_V2:-0x89D1cBD9DEAaB4dFf6f800a336FBDd9A5c6829e4}"
EXPECTED_OWNER="${EXPECTED_OWNER:-0xf9c82B1117e8BeA97843042521B8FBC93044f347}"

OWNER_SELECTOR="0x8da5cb5b"
FEE_REFUND_BLOCK_COUNT_SELECTOR="0x0fe34e68"
MIN_STAKE_SELECTOR="0x375b3c0a"
QUOTA_FACTOR_SELECTOR="0x976dd021"
HOLD_TIME_SELECTOR="0x097d5155"
STAKE_FOR_SELECTOR="4bf69206"
UNSTAKE_FOR_SELECTOR="36ef088c"

lower() {
  tr '[:upper:]' '[:lower:]'
}

normalize_address() {
  local value="${1#0x}"
  value="${value: -40}"
  printf '0x%s' "$value" | lower
}

rpc() {
  local method="$1"
  local params="${2:-[]}"
  local payload response

  payload="$(jq -nc --arg method "$method" --argjson params "$params" \
    '{jsonrpc:"2.0",id:1,method:$method,params:$params}')"

  if ! response="$(curl -sS --fail -H 'content-type: application/json' --data "$payload" "$RPC_URL")"; then
    printf 'RPC request failed: %s\n' "$method" >&2
    exit 1
  fi

  if jq -e '.error' >/dev/null <<<"$response"; then
    printf 'RPC returned error for %s: %s\n' "$method" "$(jq -c '.error' <<<"$response")" >&2
    exit 1
  fi

  jq -c '.result' <<<"$response"
}

json_array() {
  if [ "$#" -eq 0 ]; then
    printf '[]'
  else
    printf '%s\n' "$@" | jq -R . | jq -s .
  fi
}

eth_call() {
  local to="$1"
  local data="$2"
  rpc eth_call "$(jq -nc --arg to "$to" --arg data "$data" '[{to:$to,data:$data},"latest"]')" | jq -r '.'
}

failures=()

client_version="$(rpc web3_clientVersion | jq -r '.')"
chain_id="$(rpc eth_chainId | jq -r '.')"
syncing_json="$(rpc eth_syncing)"
block_number_hex="$(rpc eth_blockNumber | jq -r '.')"
rules_json="$(rpc vc_getRules '["latest"]')"

expected_payback_v2_lc="$(normalize_address "$EXPECTED_PAYBACK_V2")"
rules_payback_v2_lc="$(normalize_address "$(jq -r '.Economy.QuotaCacheAddress // ""' <<<"$rules_json")")"
expected_owner_lc="$(normalize_address "$EXPECTED_OWNER")"
payback_v2_code="$(rpc eth_getCode "[\"$EXPECTED_PAYBACK_V2\",\"latest\"]" | jq -r '.')"
payback_v2_code_lc="$(printf '%s' "$payback_v2_code" | lower)"

owner_call="$(eth_call "$EXPECTED_PAYBACK_V2" "$OWNER_SELECTOR")"
owner_lc="$(normalize_address "$owner_call")"
fee_refund_block_count="$(eth_call "$EXPECTED_PAYBACK_V2" "$FEE_REFUND_BLOCK_COUNT_SELECTOR")"
min_stake="$(eth_call "$EXPECTED_PAYBACK_V2" "$MIN_STAKE_SELECTOR")"
quota_factor="$(eth_call "$EXPECTED_PAYBACK_V2" "$QUOTA_FACTOR_SELECTOR")"
hold_time="$(eth_call "$EXPECTED_PAYBACK_V2" "$HOLD_TIME_SELECTOR")"

if [ "$block_number_hex" = "0x" ]; then
  block_number=0
else
  block_number=$((16#${block_number_hex#0x}))
fi

if [ "$chain_id" != "$EXPECTED_CHAIN_ID" ]; then
  failures+=("chain id is $chain_id, expected $EXPECTED_CHAIN_ID")
fi

if [[ "$client_version" != *"$EXPECTED_CLIENT_VERSION"* ]]; then
  failures+=("client version does not include $EXPECTED_CLIENT_VERSION")
fi

if [ -n "$EXPECTED_CLIENT_COMMIT" ] && [[ "$client_version" != *"$EXPECTED_CLIENT_COMMIT"* ]]; then
  failures+=("client version does not include commit $EXPECTED_CLIENT_COMMIT")
fi

if [ "$syncing_json" != "false" ]; then
  failures+=("public RPC reports eth_syncing=$syncing_json")
fi

if [ "$rules_payback_v2_lc" != "$expected_payback_v2_lc" ]; then
  failures+=("vc_getRules Economy.QuotaCacheAddress is $rules_payback_v2_lc, expected $expected_payback_v2_lc")
fi

if [ "$(jq -r '.Upgrades.PaybackV2 // false' <<<"$rules_json")" != "true" ]; then
  failures+=("vc_getRules Upgrades.PaybackV2 is not true")
fi

if [ "$(jq -r '.Upgrades.PaybackV2Patch // false' <<<"$rules_json")" != "true" ]; then
  failures+=("vc_getRules Upgrades.PaybackV2Patch is not true")
fi

if [ "$payback_v2_code" = "0x" ]; then
  failures+=("corrected PaybackV2 has no code at $expected_payback_v2_lc")
  payback_v2_code_bytes=0
else
  payback_v2_code_bytes=$(((${#payback_v2_code} - 2) / 2))
fi

if [[ "$payback_v2_code_lc" != *"$STAKE_FOR_SELECTOR"* ]]; then
  failures+=("corrected PaybackV2 bytecode does not contain stakeFor(address) selector 0x$STAKE_FOR_SELECTOR")
fi

if [[ "$payback_v2_code_lc" != *"$UNSTAKE_FOR_SELECTOR"* ]]; then
  failures+=("corrected PaybackV2 bytecode does not contain unstakeFor(address,uint256) selector 0x$UNSTAKE_FOR_SELECTOR")
fi

if [ "$owner_lc" != "$expected_owner_lc" ]; then
  failures+=("corrected PaybackV2 owner is $owner_lc, expected $expected_owner_lc")
fi

failures_json="$(json_array "${failures[@]}")"
status=passed
if [ "${#failures[@]}" -gt 0 ]; then
  status=failed
fi

jq -n \
  --arg status "$status" \
  --arg rpcUrl "$RPC_URL" \
  --arg clientVersion "$client_version" \
  --arg expectedClientVersion "$EXPECTED_CLIENT_VERSION" \
  --arg expectedClientCommit "$EXPECTED_CLIENT_COMMIT" \
  --arg chainId "$chain_id" \
  --arg expectedChainId "$EXPECTED_CHAIN_ID" \
  --argjson syncing "$syncing_json" \
  --arg blockNumberHex "$block_number_hex" \
  --argjson blockNumber "$block_number" \
  --arg paybackV2 "$rules_payback_v2_lc" \
  --arg expectedPaybackV2 "$expected_payback_v2_lc" \
  --arg owner "$owner_lc" \
  --arg expectedOwner "$expected_owner_lc" \
  --argjson codeBytes "$payback_v2_code_bytes" \
  --arg feeRefundBlockCount "$fee_refund_block_count" \
  --arg minStake "$min_stake" \
  --arg quotaFactor "$quota_factor" \
  --arg holdTime "$hold_time" \
  --argjson upgrades "$(jq -c '.Upgrades' <<<"$rules_json")" \
  --argjson failures "$failures_json" \
  '{
    status: $status,
    rpc: {
      url: $rpcUrl,
      clientVersion: $clientVersion,
      expectedClientVersion: $expectedClientVersion,
      expectedClientCommit: $expectedClientCommit,
      chainId: $chainId,
      expectedChainId: $expectedChainId,
      syncing: $syncing,
      blockNumberHex: $blockNumberHex,
      blockNumber: $blockNumber
    },
    rules: {
      paybackV2: $paybackV2,
      expectedPaybackV2: $expectedPaybackV2,
      upgrades: $upgrades
    },
    contract: {
      address: $expectedPaybackV2,
      owner: $owner,
      expectedOwner: $expectedOwner,
      codeBytes: $codeBytes,
      feeRefundBlockCount: $feeRefundBlockCount,
      minStake: $minStake,
      quotaFactor: $quotaFactor,
      holdTime: $holdTime
    },
    failures: $failures
  }'

if [ "${#failures[@]}" -gt 0 ]; then
  exit 1
fi
