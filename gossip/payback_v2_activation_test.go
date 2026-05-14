package gossip

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestPaybackV2ActivationSwitchesQuotaCacheAddress is a source-structural
// pin for the PaybackV2 activation branch in sealEpochIfNeeded.
//
// Rationale: a live activation test would need a fully-wired BlockProcessor
// (Store, EVM state, sealer, validator set, etc.) which the existing
// sibling tests in this package likewise stub out. So we pin the contract:
//  1. The activation site exists and is gated on `Upgrades.PaybackV2 && !prevUpg.PaybackV2`
//     so it fires exactly once at the transition edge.
//  2. The site reads the new address from opera.PaybackV2ContractAddress so
//     a future address change cannot accidentally introduce a second source
//     of truth.
//  3. The site rejects the sentinel address via opera.PaybackV2AddressIsSentinel,
//     so a binary that lost the deployed address mid-build aborts at the
//     activation block rather than silently bricking payback.
//  4. The site overwrites bp.es.Rules.Economy.QuotaCacheAddress in place,
//     so the downstream payback cache resolver picks up the new address
//     in the same seal pass.
func TestPaybackV2ActivationSwitchesQuotaCacheAddress(t *testing.T) {
	src, err := os.ReadFile("block_processor.go")
	require.NoError(t, err)
	s := string(src)

	require.Contains(t, s, "Upgrades.PaybackV2 && !prevUpg.PaybackV2",
		"sealEpochIfNeeded must gate PaybackV2 activation on the !prevUpg edge so it fires exactly once at transition")
	require.Contains(t, s, "opera.PaybackV2ContractAddress(bp.es.Rules.NetworkID)",
		"activation must read the new address from opera.PaybackV2ContractAddress to keep a single source of truth")
	require.Contains(t, s, "opera.PaybackV2AddressIsSentinel(newAddr)",
		"activation must refuse to swap to the zero sentinel — this is the defence-in-depth guard for a binary that lost its deployed address")
	require.Contains(t, s, "bp.es.Rules.Economy.QuotaCacheAddress = newAddr",
		"activation must overwrite Economy.QuotaCacheAddress in place so the payback cache picks it up at the next resolve")
	require.Contains(t, s, "log.Crit(\"PaybackV2 activation failed: V2 contract address is the zero sentinel",
		"sentinel-detection log line must fingerprint the failure cause for operators")
}

// TestPaybackV2StagingFromHardcodedRules pins the binary-rules-vs-stored-rules
// staging logic in service.go. service.go's job at startup is to detect new
// flags introduced by the running binary that are not yet in the persisted
// pending rules, and stage them so the next epoch seal activates them.
// PaybackV2 must follow the same staging shape as every previous upgrade.
func TestPaybackV2StagingFromHardcodedRules(t *testing.T) {
	src, err := os.ReadFile("service.go")
	require.NoError(t, err)
	s := string(src)

	require.Contains(t, s, "hardcoded.Upgrades.PaybackV2 && !pending.Upgrades.PaybackV2",
		"service.go must stage PaybackV2 from binary rules into pending DirtyRules so the next epoch seal activates it")
	require.Contains(t, s, "Staged PaybackV2 upgrade from binary rules",
		"staging log line must follow the existing pattern so log-greppers can find the activation event")
}

// TestPaybackV2BitfieldEncodingPresent pins the RLP bitfield wiring in
// opera/legacy_serialization.go. Without this, a copy-paste regression
// could drop the flag from the on-wire encoding while keeping the Go
// struct field, silently de-activating PaybackV2 across the fleet.
func TestPaybackV2BitfieldEncodingPresent(t *testing.T) {
	src, err := os.ReadFile("../opera/legacy_serialization.go")
	require.NoError(t, err)
	s := string(src)
	require.Contains(t, s, "u.PaybackV2 {", "EncodeRLP must check u.PaybackV2 before setting the bit")
	require.Contains(t, s, "bitmap.V |= paybackV2Bit", "EncodeRLP must OR in paybackV2Bit when PaybackV2 is true")
	require.Contains(t, s, "u.PaybackV2 = (bitmap.V & paybackV2Bit) != 0",
		"DecodeRLP must extract PaybackV2 from paybackV2Bit")
}

// TestPaybackV2ActivationRefreshesEvmProcessorRules pins the H-01 fix from
// the 2026-05-14 security review. The activation branch in sealEpochIfNeeded
// MUST call bp.evmProcessor.SetRules(bp.es.Rules) immediately after mutating
// Economy.QuotaCacheAddress. Without this, evmProcessor.net stays at the
// pre-seal value copy and post-internal + user txs in the activation block
// see the OLD address:
//   - receipts encode FeeRefund against the OLD QuotaContract
//   - stakeFor txs targeting the NEW V2 contract are NOT recorded in
//     PaybackCache.StakesMap (txtype=TxTypeNone)
//   - EvmWriter.SetPaybackProxyAddr keeps the OLD address as the protected
//     system contract for the activation block, leaving the NEW address
//     unprotected against swapCode/setStorage for one block
//
// Source-structural pin rather than a runtime test because activating
// requires a fully-wired BlockProcessor + Store + sealer + EVM that the
// sibling tests in this package likewise stub out.
func TestPaybackV2ActivationRefreshesEvmProcessorRules(t *testing.T) {
	src, err := os.ReadFile("block_processor.go")
	require.NoError(t, err)
	s := string(src)
	require.Contains(t, s, "bp.evmProcessor.SetRules(bp.es.Rules)",
		"activation branch MUST call evmProcessor.SetRules(bp.es.Rules) right after the QuotaCacheAddress swap so the same-block evmProcessor sees the new address")
	require.Contains(t, s, "if bp.evmProcessor != nil {",
		"the SetRules call MUST be nil-guarded — endBlock-without-initProcessors is a valid intermediate state during recovery paths")
}
