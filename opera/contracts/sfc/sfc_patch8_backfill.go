package sfc

import (
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/log"
)

// SKELETON — no pre-upgrade stranded reactivation-gap pairs are known yet.
//
// The SfcV2Patch8 Solidity change makes reactivateValidator self-service AND
// heals a validator's offline reward gap forward via two appended mappings
// (reactivationHealFloor, reactivationHealFrom, captured at reactivation time).
// That heal only covers validators reactivated AFTER the Patch8 bytecode is
// live. A validator that was reactivated BEFORE the upgrade (owner-driven, under
// the old onlyOwner reactivateValidator) left no heal record, so its delegators
// may still be stranded behind an inverted reward index across that historical
// gap. This backfill is the deterministic, one-shot storage repair for those
// PRE-upgrade pairs — the Go analog of sfc_cursor_backfill.go.
//
// It is intentionally EMPTY for now: the corrupted pre-upgrade
// (validatorID, gap) set has not been enumerated off-chain, and we do not invent
// addresses. With an empty list this is a verified no-op, so wiring it at the
// Patch8 activation seal is safe and forward-compatible.
//
// TODO(reactivation-backfill): before the mainnet/testnet Patch8 rollout,
//   1. Enumerate validators reactivated before Patch8 with external delegators
//      whose stashedRewardsUntilEpoch cursor sits inside the offline gap
//      (inverted / frozen reward index). Solve each (validatorID -> R, fromEpoch)
//      off-chain: R = accumulatedRewardPerToken[validatorID] at the
//      deactivatedEpoch; fromEpoch = deactivatedEpoch + 1. Record as
//      include:true rows in a committed corrected_reactivations.json, mirroring
//      corrected_cursors.json.
//   2. Derive the SFC storage slots for reactivationHealFloor[validatorID] and
//      reactivationHealFrom[validatorID]. Both are single-level
//      mapping(uint256 => uint256) appended immediately AFTER
//      _reentrancyGuardCounter and BEFORE __gap (see the SFC.sol storage tail).
//      Confirm their base slots by counting the linearized layout the same way
//      sfc_cursor_backfill.go derived slot 114 for stashedRewardsUntilEpoch, and
//      verify against a live eth_getStorageAt before trusting them.
//   3. Populate reactivationBackfillPairs and set floor/from with a
//      write-once, no-overwrite guard (never clobber a heal record the Solidity
//      path already wrote post-upgrade).

// reactivationBackfillPair is a single pre-upgrade (validatorID) heal record to
// install: the pre-gap flat rate R (healFloor) and the first gap epoch
// (healFrom = deactivatedEpoch + 1). Delegators are healed per-validator, so no
// delegator address is needed here — the two SFC mappings are keyed by
// validatorID only.
type reactivationBackfillPair struct {
	validatorID uint64
	healFloor   uint64
	healFrom    uint64
}

// reactivationBackfillPairs is intentionally empty — see the SKELETON note and
// the TODO above. Do NOT invent entries; each must be an off-chain-solved,
// include:true row backed by live storage verification.
var reactivationBackfillPairs = []reactivationBackfillPair{}

// ReactivationBackfillStats reports the deterministic reactivation-heal records
// installed at the SfcV2Patch8 activation boundary.
type ReactivationBackfillStats struct {
	Installed uint64
}

// Changed reports whether the backfill wrote any SFC storage.
func (stats ReactivationBackfillStats) Changed() bool {
	return stats.Installed != 0
}

// BackfillReactivationHealRecords installs reactivationHealFloor/From records
// for the known PRE-upgrade stranded reactivation-gap pairs at the SfcV2Patch8
// activation seal. It is a ONE-SHOT, write-once storage repair. With the current
// empty pair list it is a verified no-op; the enumeration + storage-slot
// derivation are tracked in the TODO above and must land before any real entry
// is added.
func BackfillReactivationHealRecords(statedb *state.StateDB) ReactivationBackfillStats {
	var stats ReactivationBackfillStats
	if len(reactivationBackfillPairs) == 0 {
		return stats
	}

	// Deliberately unreachable until reactivationBackfillPairs is populated AND
	// the storage-slot derivation (step 2 of the TODO) is implemented and
	// verified against live state. Fail loud rather than silently write to
	// unverified slots.
	first := reactivationBackfillPairs[0]
	log.Crit("SfcV2Patch8 reactivation backfill has pairs but no verified storage-slot writer - refusing to write to unverified SFC slots",
		"pairs", len(reactivationBackfillPairs),
		"first_validator_id", first.validatorID,
		"first_heal_floor", first.healFloor,
		"first_heal_from", first.healFrom)
	return stats
}
