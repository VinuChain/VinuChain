package gossip

import (
	"os"
	"testing"

	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/stretchr/testify/require"
)

// TestPaybackV2ActivationSwitchesQuotaCacheAddress is a source-structural
// pin for the PaybackV2 activation branches in sealEpochIfNeeded.
//
// Rationale: a live activation test would need a fully-wired BlockProcessor
// (Store, EVM state, sealer, validator set, etc.) which the existing
// sibling tests in this package likewise stub out. So we pin the contract:
//  1. The original activation site is gated on `Upgrades.PaybackV2 && !prevUpg.PaybackV2`
//     so it fires exactly once at the first PaybackV2 transition edge.
//  2. The patch activation site is gated on `Upgrades.PaybackV2Patch && !prevUpg.PaybackV2Patch`
//     so an already-active testnet can deterministically rebind the bad V2 address.
//  3. The site reads the new address from opera.PaybackV2ContractAddress so
//     a future address change cannot accidentally introduce a second source
//     of truth.
//  4. The site rejects the sentinel address via opera.PaybackV2AddressIsSentinel,
//     so a binary that lost the deployed address mid-build aborts at the
//     activation block rather than silently bricking payback.
//  5. The site overwrites bp.es.Rules.Economy.QuotaCacheAddress in place,
//     so the downstream payback cache resolver picks up the new address
//     in the same seal pass.
func TestPaybackV2ActivationSwitchesQuotaCacheAddress(t *testing.T) {
	src, err := os.ReadFile("block_processor.go")
	require.NoError(t, err)
	s := string(src)

	require.Contains(t, s, "Upgrades.PaybackV2 && !prevUpg.PaybackV2",
		"sealEpochIfNeeded must gate PaybackV2 activation on the !prevUpg edge so it fires exactly once at transition")
	require.Contains(t, s, "Upgrades.PaybackV2Patch && !prevUpg.PaybackV2Patch",
		"sealEpochIfNeeded must gate PaybackV2Patch on its own !prevUpg edge so already-active testnet can rebind once")
	require.Contains(t, s, "opera.PaybackV2ContractAddress(bp.es.Rules.NetworkID)",
		"activation must read the new address from opera.PaybackV2ContractAddress to keep a single source of truth")
	require.Contains(t, s, "opera.PaybackV2AddressIsSentinel(newAddr)",
		"activation must refuse to swap to the zero sentinel — this is the defence-in-depth guard for a binary that lost its deployed address")
	require.Contains(t, s, "bp.es.Rules.Economy.QuotaCacheAddress = newAddr",
		"activation must overwrite Economy.QuotaCacheAddress in place so the payback cache picks it up at the next resolve")
	require.Contains(t, s, "failed: V2 contract address is the zero sentinel",
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
	require.Contains(t, s, "hardcoded.Upgrades.PaybackV2Patch && !pending.Upgrades.PaybackV2Patch",
		"service.go must stage PaybackV2Patch from binary rules into pending DirtyRules so the next epoch seal rebinds active testnet")
	require.Contains(t, s, "Staged PaybackV2Patch upgrade from binary rules",
		"patch staging log line must follow the existing pattern so log-greppers can find the repair event")
}

// TestSfcV2Patch7StagingFromHardcodedRules pins the binary-rules-vs-stored-rules
// staging logic in service.go for SfcV2Patch7. On a live fleet where every
// prior flag is already sealed, the SfcV2Patch7 flag only enters DirtyRules if
// stageHardcodedUpgrades has a branch for it — without the branch the sealer
// never flips the flag, the block_processor activation guard
// (SfcV2Patch7 && !prevUpg.SfcV2Patch7) never fires, and the reward-cursor
// reflash + backfill never run. This is the BLOCKER-2 regression guard: it must
// follow the exact same staging shape as every previous SfcV2Patch and the
// PaybackV2 upgrade.
func TestSfcV2Patch7StagingFromHardcodedRules(t *testing.T) {
	src, err := os.ReadFile("service.go")
	require.NoError(t, err)
	s := string(src)

	require.Contains(t, s, "hardcoded.Upgrades.SfcV2Patch7 && !pending.Upgrades.SfcV2Patch7",
		"service.go must stage SfcV2Patch7 from binary rules into pending DirtyRules so the next epoch seal activates the reward-cursor reflash + backfill")
	require.Contains(t, s, "Staged SfcV2Patch7 upgrade from binary rules",
		"staging log line must follow the existing pattern so log-greppers can find the activation event")

	// The staging branch must sit AFTER the SfcV2Patch6 branch so the SFC
	// reflashes activate in cycle order.
	require.Less(t,
		indexOf(s, "hardcoded.Upgrades.SfcV2Patch6 && !pending.Upgrades.SfcV2Patch6"),
		indexOf(s, "hardcoded.Upgrades.SfcV2Patch7 && !pending.Upgrades.SfcV2Patch7"),
		"SfcV2Patch7 staging branch must be placed after the SfcV2Patch6 branch")
}

// TestSfcV2Patch7HardcodedRulesPerNetwork pins that stageHardcodedUpgrades will
// stage SfcV2Patch7 ONLY on testnet (NetworkID 206) and not on mainnet (207) or
// staging (205) — staging consults opera.MainNetRulesForNetwork, so this asserts
// the same source the staging branch reads.
func TestSfcV2Patch7HardcodedRulesPerNetwork(t *testing.T) {
	testnet := opera.MainNetRulesForNetwork(opera.VinuChainTestNetworkID)
	require.NotNil(t, testnet, "testnet rules must resolve")
	require.True(t, testnet.Upgrades.SfcV2Patch7,
		"testnet (206) hardcoded rules must carry SfcV2Patch7=true so stageHardcodedUpgrades stages it")

	mainnet := opera.MainNetRulesForNetwork(opera.VinuChainMainNetworkID)
	require.NotNil(t, mainnet, "mainnet rules must resolve")
	require.False(t, mainnet.Upgrades.SfcV2Patch7,
		"mainnet (207) hardcoded rules must NOT carry SfcV2Patch7 yet — it ships in a later mainnet release")

	staging := opera.MainNetRulesForNetwork(opera.VinuChainStagingNetworkID)
	require.NotNil(t, staging, "staging rules must resolve")
	require.False(t, staging.Upgrades.SfcV2Patch7,
		"staging (205) derives from mainnet rules, so SfcV2Patch7 must be false there too")
}

// indexOf is a tiny helper so the ordering assertion above reads clearly.
func indexOf(haystack, needle string) int {
	for i := 0; i+len(needle) <= len(haystack); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
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
	require.Contains(t, s, "bitmap.V |= paybackV2PatchBit", "EncodeRLP must OR in paybackV2PatchBit when PaybackV2Patch is true")
	require.Contains(t, s, "u.PaybackV2Patch = (bitmap.V & paybackV2PatchBit) != 0",
		"DecodeRLP must extract PaybackV2Patch from paybackV2PatchBit")
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
