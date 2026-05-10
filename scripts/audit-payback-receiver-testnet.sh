#!/usr/bin/env bash
set -euo pipefail

RPC_URL="${RPC_URL:-https://vinufoundation-rpc.com}"
EXPECTED_CHAIN_ID="${EXPECTED_CHAIN_ID:-0xce}"
EXPECTED_CLIENT_VERSION="${EXPECTED_CLIENT_VERSION:-v2.0.17-elemont}"
EXPECTED_CLIENT_COMMIT="${EXPECTED_CLIENT_COMMIT:-bbc34e6}"
EXPECTED_QUOTA_PROXY="${EXPECTED_QUOTA_PROXY:-0x824B93dE7221cf8a35FBd29d5202f6eFa3A29C5D}"
EXPECTED_PROXY_ADMIN="${EXPECTED_PROXY_ADMIN:-0xcE154534e1E8F4Cc9Ab642Ad1816Ee1A237055F4}"
EXPECTED_PROXY_ADMIN_OWNER="${EXPECTED_PROXY_ADMIN_OWNER:-0x07B4eF04b62E69aE14A715cdcae692fa7033b9a5}"
KNOWN_PRE_RECEIVER_IMPLEMENTATION="${KNOWN_PRE_RECEIVER_IMPLEMENTATION:-0x0c8735bD6b3E90eaD4cdAB917474Cc6e8E58ce82}"
STAKE_FOR_SELECTOR="${STAKE_FOR_SELECTOR:-4bf69206}"
EIP1967_IMPLEMENTATION_SLOT="0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc"
OWNER_SELECTOR="0x8da5cb5b"

truthy() {
  case "${1:-}" in
    1 | true | TRUE | yes | YES | y | Y) return 0 ;;
    *) return 1 ;;
  esac
}

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

failures=()
ready_blockers=()

client_version="$(rpc web3_clientVersion | jq -r '.')"
chain_id="$(rpc eth_chainId | jq -r '.')"
syncing_json="$(rpc eth_syncing)"
block_number_hex="$(rpc eth_blockNumber | jq -r '.')"
rules_json="$(rpc eth_getRules '["latest"]')"

quota_proxy_from_rules="$(jq -r '.Economy.QuotaCacheAddress // ""' <<<"$rules_json")"
quota_proxy_lc="$(normalize_address "$EXPECTED_QUOTA_PROXY")"
rules_quota_lc="$(normalize_address "$quota_proxy_from_rules")"
proxy_admin_lc="$(normalize_address "$EXPECTED_PROXY_ADMIN")"
proxy_admin_owner_lc="$(normalize_address "$EXPECTED_PROXY_ADMIN_OWNER")"
known_pre_receiver_impl_lc="$(normalize_address "$KNOWN_PRE_RECEIVER_IMPLEMENTATION")"

implementation_slot="$(rpc eth_getStorageAt "[\"$EXPECTED_QUOTA_PROXY\",\"$EIP1967_IMPLEMENTATION_SLOT\",\"latest\"]" | jq -r '.')"
implementation_lc="$(normalize_address "$implementation_slot")"
admin_owner_call="$(rpc eth_call "[{\"to\":\"$EXPECTED_PROXY_ADMIN\",\"data\":\"$OWNER_SELECTOR\"},\"latest\"]" | jq -r '.')"
admin_owner_lc="$(normalize_address "$admin_owner_call")"
implementation_code="$(rpc eth_getCode "[\"$implementation_lc\",\"latest\"]" | jq -r '.')"
implementation_code_lc="$(printf '%s' "$implementation_code" | lower)"

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

if [[ "$client_version" != *"$EXPECTED_CLIENT_COMMIT"* ]]; then
  failures+=("client version does not include commit $EXPECTED_CLIENT_COMMIT")
fi

if [ "$syncing_json" != "false" ]; then
  failures+=("public RPC reports eth_syncing=$syncing_json")
fi

if [ "$rules_quota_lc" != "$quota_proxy_lc" ]; then
  failures+=("eth_getRules Economy.QuotaCacheAddress is $rules_quota_lc, expected $quota_proxy_lc")
fi

if [ "$admin_owner_lc" != "$proxy_admin_owner_lc" ]; then
  failures+=("ProxyAdmin owner is $admin_owner_lc, expected $proxy_admin_owner_lc")
fi

if [ "$implementation_lc" = "0x0000000000000000000000000000000000000000" ]; then
  failures+=("Quota proxy implementation slot is zero")
fi

if [ "$implementation_code" = "0x" ]; then
  failures+=("Quota implementation has no code at $implementation_lc")
  implementation_code_bytes=0
else
  implementation_code_bytes=$(((${#implementation_code} - 2) / 2))
fi

implementation_has_stake_for=false
if [[ "$implementation_code_lc" == *"$STAKE_FOR_SELECTOR"* ]]; then
  implementation_has_stake_for=true
else
  ready_blockers+=("Quota implementation bytecode does not contain stakeFor(address) selector 0x$STAKE_FOR_SELECTOR")
fi

implementation_is_known_pre_receiver=false
if [ "$implementation_lc" = "$known_pre_receiver_impl_lc" ]; then
  implementation_is_known_pre_receiver=true
  ready_blockers+=("Quota proxy still points at known pre-receiver implementation $known_pre_receiver_impl_lc")
fi

failures_json="$(json_array "${failures[@]}")"
ready_blockers_json="$(json_array "${ready_blockers[@]}")"
require_ready=false
if truthy "${REQUIRE_PAYBACK_RECEIVER_READY:-false}"; then
  require_ready=true
fi

ready_status=ready
if [ "${#ready_blockers[@]}" -gt 0 ]; then
  ready_status=not_ready
fi

status=passed
if [ "${#failures[@]}" -gt 0 ]; then
  status=failed
fi
if $require_ready && [ "${#ready_blockers[@]}" -gt 0 ]; then
  status=failed
fi

jq -n \
  --arg status "$status" \
  --arg receiverReadiness "$ready_status" \
  --argjson requirePaybackReceiverReady "$require_ready" \
  --arg rpcUrl "$RPC_URL" \
  --arg clientVersion "$client_version" \
  --arg expectedClientVersion "$EXPECTED_CLIENT_VERSION" \
  --arg expectedClientCommit "$EXPECTED_CLIENT_COMMIT" \
  --arg chainId "$chain_id" \
  --arg expectedChainId "$EXPECTED_CHAIN_ID" \
  --argjson syncing "$syncing_json" \
  --arg blockNumberHex "$block_number_hex" \
  --argjson blockNumber "$block_number" \
  --arg quotaProxy "$quota_proxy_lc" \
  --arg rulesQuotaProxy "$rules_quota_lc" \
  --arg proxyAdmin "$proxy_admin_lc" \
  --arg proxyAdminOwner "$admin_owner_lc" \
  --arg expectedProxyAdminOwner "$proxy_admin_owner_lc" \
  --arg implementation "$implementation_lc" \
  --arg knownPreReceiverImplementation "$known_pre_receiver_impl_lc" \
  --argjson implementationCodeBytes "$implementation_code_bytes" \
  --argjson implementationHasStakeFor "$implementation_has_stake_for" \
  --argjson implementationIsKnownPreReceiver "$implementation_is_known_pre_receiver" \
  --arg stakeForSelector "0x$STAKE_FOR_SELECTOR" \
  --argjson upgrades "$(jq -c '.Upgrades' <<<"$rules_json")" \
  --argjson failures "$failures_json" \
  --argjson readyBlockers "$ready_blockers_json" \
  '{
    status: $status,
    receiverReadiness: $receiverReadiness,
    requirePaybackReceiverReady: $requirePaybackReceiverReady,
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
      quotaProxy: $rulesQuotaProxy,
      expectedQuotaProxy: $quotaProxy,
      upgrades: $upgrades
    },
    quotaProxy: {
      proxy: $quotaProxy,
      proxyAdmin: $proxyAdmin,
      proxyAdminOwner: $proxyAdminOwner,
      expectedProxyAdminOwner: $expectedProxyAdminOwner,
      implementation: $implementation,
      knownPreReceiverImplementation: $knownPreReceiverImplementation,
      implementationCodeBytes: $implementationCodeBytes,
      implementationHasStakeFor: $implementationHasStakeFor,
      implementationIsKnownPreReceiver: $implementationIsKnownPreReceiver,
      stakeForSelector: $stakeForSelector
    },
    failures: $failures,
    readyBlockers: $readyBlockers
  }'

if [ "${#failures[@]}" -gt 0 ]; then
  exit 1
fi

if $require_ready && [ "${#ready_blockers[@]}" -gt 0 ]; then
  exit 1
fi
