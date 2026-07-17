package sfc

import (
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/log"
)

// SKELETON — no pre-upgrade stranded reward-cursor/reactivation-gap pairs are
// known yet.
//
// The SfcV2Patch9 Solidity change ships two reward fixes on top of Patch8's
// self-service reactivation: (1) _rawDelegate seeds a delegator's reward
// cursor to currentSealedEpoch+1 on first delegation, avoiding a one-epoch
// over-mint for new delegators; (2) reactivateValidator physically backfills
// a prior offline gap on repeated reactivation so passive delegators are not
// re-stranded. Both fixes are forward-acting in the bytecode itself and only
// cover delegations/reactivations that occur AFTER the Patch9 bytecode is
// live. A pair that was affected BEFORE the upgrade leaves no corrected
// record, so its delegators may still be sitting behind the pre-upgrade
// defect. This backfill is the deterministic, one-shot storage repair for
// those PRE-upgrade pairs — the Go analog of sfc_cursor_backfill.go /
// sfc_patch8_backfill.go.
//
// It is intentionally EMPTY for now: the corrupted pre-upgrade set has not
// been enumerated off-chain, and we do not invent addresses. With an empty
// list this is a verified no-op, so wiring it at the Patch9 activation seal
// is safe and forward-compatible.
//
// TODO(patch9-backfill): before the mainnet/testnet Patch9 rollout,
//  1. Enumerate delegators/validators affected before Patch9 whose reward
//     cursor or reactivation-heal state still reflects the pre-fix defect.
//     Solve each pair off-chain and record as include:true rows in a
//     committed corrected file, mirroring corrected_cursors.json /
//     corrected_reactivations.json.
//  2. Derive the SFC storage slots touched by the fix and confirm against a
//     live eth_getStorageAt before trusting them.
//  3. Populate patch9BackfillPairs and write with a write-once, no-overwrite
//     guard (never clobber a record the Solidity path already wrote
//     post-upgrade).

// patch9BackfillPair is a single pre-upgrade heal record to install: the
// pre-gap flat rate R (healFloor) and the first gap epoch (healFrom).
type patch9BackfillPair struct {
	validatorID uint64
	healFloor   uint64
	healFrom    uint64
}

// patch9BackfillPairs is intentionally empty — see the SKELETON note and the
// TODO above. Do NOT invent entries; each must be an off-chain-solved,
// include:true row backed by live storage verification.
var patch9BackfillPairs = []patch9BackfillPair{}

// Patch9BackfillStats reports the deterministic heal records installed at the
// SfcV2Patch9 activation boundary.
type Patch9BackfillStats struct {
	Installed uint64
}

// Changed reports whether the backfill wrote any SFC storage.
func (stats Patch9BackfillStats) Changed() bool {
	return stats.Installed != 0
}

// BackfillPatch9ReactivationHealRecords installs heal records for the known
// PRE-upgrade stranded pairs at the SfcV2Patch9 activation seal. It is a
// ONE-SHOT, write-once storage repair. With the current empty pair list it is
// a verified no-op; the enumeration + storage-slot derivation are tracked in
// the TODO above and must land before any real entry is added.
func BackfillPatch9ReactivationHealRecords(statedb *state.StateDB) Patch9BackfillStats {
	var stats Patch9BackfillStats
	if len(patch9BackfillPairs) == 0 {
		return stats
	}

	// Deliberately unreachable until patch9BackfillPairs is populated AND the
	// storage-slot derivation (step 2 of the TODO) is implemented and
	// verified against live state. Fail loud rather than silently write to
	// unverified slots.
	log.Crit("SfcV2Patch9 backfill has pairs but no verified storage-slot writer - refusing to write to unverified SFC slots",
		"pairs", len(patch9BackfillPairs))
	return stats
}
