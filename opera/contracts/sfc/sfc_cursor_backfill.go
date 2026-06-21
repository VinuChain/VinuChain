package sfc

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

// sfcStashedRewardsUntilEpochSlot is the SFC storage base slot of the
// `stashedRewardsUntilEpoch[address][uint256]` nested mapping. It was derived
// from the SFC storage layout (contracts/vinuchain/SFC.sol) by counting state
// variables in linearized C3 inheritance order:
//
//	Initializable: slots 0-50 (initialized+initializing packed in slot 0, __gap[50])
//	Ownable:       slots 51-101 (_owner, _pendingOwner, _pendingOwnerDeadline, __gap[48])
//	StakersConstants + Version: 0 slots (constants / pure functions only)
//	SFC own storage starts at slot 102:
//	  102 node, 103 currentSealedEpoch, 104 _legacyGenesisValidator,
//	  105 getValidator, 106 getValidatorID, 107 getValidatorPubkey,
//	  108 lastValidatorID, 109 totalStake, 110 totalActiveStake,
//	  111 totalSlashedStake, 112 totalPenalty, 113 _rewardsStash,
//	  114 stashedRewardsUntilEpoch  <-- this slot
//	  115 getWithdrawalRequest, 116 getStake, ... 125 stakes, 126 stakePosition
//
// The 116/125/126 anchors are exactly the slots sfc_patch6_backfill.go derived
// for getStake/stakes/stakePosition, which validates the count. The base slot
// was ALSO verified against live testnet (NetworkID 206) before being trusted:
// for the known pairs (0x8406413be0…,17)=300, (0xfe62bd156e…,17)=1900, and
// (0x031a75844f…,15)=100, eth_getStorageAt at
// keccak(validatorID || keccak(delegator || 114)) returned the same value as
// eth_call stashedRewardsUntilEpoch(delegator, validatorID).
const sfcStashedRewardsUntilEpochSlot uint64 = 114

// cursorBackfillPair is a single (delegator, validatorID) cursor correction.
// correctedEpoch is the smallest epoch whose single-cursor ARPT integral lands
// at or just under the genuine owed reward (computed off-chain in the prior
// forensic stage); raising the stored cursor to this value lets a subsequent
// claim mint only the genuine owed reward instead of reverting "zero rewards".
type cursorBackfillPair struct {
	delegator      common.Address
	validatorID    uint64
	correctedEpoch uint64
}

func cursorPair(delegator string, validatorID uint64, correctedEpoch uint64) cursorBackfillPair {
	return cursorBackfillPair{
		delegator:      common.HexToAddress(delegator),
		validatorID:    validatorID,
		correctedEpoch: correctedEpoch,
	}
}

// testnetCursorBackfill holds the 12 already-corrupted testnet (NetworkID 206)
// (delegator, validatorID) pairs whose stashedRewardsUntilEpoch cursor is stuck
// far below the validator's createdEpoch in the all-zero-ARPT dead zone (the
// "cursor_overmint" class). correctedEpoch is the off-chain-solved under-pay
// epoch from corrected_cursors.json (include:true rows only).
//
// Deliberately EXCLUDED (see corrected_cursors.json pairsExcluded):
//   - the 4 near-miss pairs (validators 6, 10, 11, 14) whose cursor is exactly
//     createdEpoch-1: the dead zone is a single pre-creation ARPT==0 epoch so
//     the stash booked at createdEpoch is already valid and claim succeeds.
//   - the 3 validator-16 pairs: validator 16 is the malformed-pubkey validator
//     that earns zero rewards at every epoch (ARPT==0 everywhere), so genuine
//     owed is truly zero and correcting the cursor would mint phantom pay.
//     Those delegators recover via the zero-penalty unlockStake -> undelegate ->
//     withdraw path, not a cursor backfill.
var testnetCursorBackfill = []cursorBackfillPair{
	cursorPair("0x031a75844f399de13f70b9be4828cb040614d5be", 15, 5643),
	cursorPair("0x0b5323811594d0c570ea6d2214a71c9fbd81db43", 15, 5754),
	cursorPair("0x8406413be017b3b7c72416a39cb8b3eeb7b9c2ef", 15, 5754),
	cursorPair("0x84db64db1bad3cc8dcedb67d9611c268e597dde7", 15, 5684),
	cursorPair("0xf10f35cc6c326f5d7c79ecab22c2297ebcc87a0b", 15, 5682),
	cursorPair("0x0b5323811594d0c570ea6d2214a71c9fbd81db43", 17, 6010),
	cursorPair("0x2ea34893c3c7513a9457f782ea35ac4390619992", 17, 6012),
	cursorPair("0x8406413be017b3b7c72416a39cb8b3eeb7b9c2ef", 17, 5824),
	cursorPair("0xfe62bd156e62214fbdd18233d9a101526b9f4dde", 17, 5753),
	cursorPair("0x2ea34893c3c7513a9457f782ea35ac4390619992", 18, 6012),
	cursorPair("0x8406413be017b3b7c72416a39cb8b3eeb7b9c2ef", 18, 5800),
	cursorPair("0xfe62bd156e62214fbdd18233d9a101526b9f4dde", 18, 5933),
}

// mainnetCursorBackfill is intentionally empty: mainnet has not activated SfcV2
// and ships the permanent reward-cursor fix via GetLatestContractBin() on its
// first activation, so there are no already-corrupted mainnet pairs to correct.
var mainnetCursorBackfill = []cursorBackfillPair{}

// CursorBackfillStats reports the deterministic cursor corrections applied at
// the SfcV2Patch7 activation boundary.
type CursorBackfillStats struct {
	Raised uint64
}

// Changed reports whether the backfill wrote any SFC storage.
func (stats CursorBackfillStats) Changed() bool {
	return stats.Raised != 0
}

// BackfillTestnetRewardCursors raises the stuck stashedRewardsUntilEpoch cursor
// for the known testnet (NetworkID 206) reward-cursor-corrupted pairs at the
// SfcV2Patch7 activation seal. currentSealedEpoch is read from SFC storage and
// used as a hard upper bound so no cursor can ever be raised past the chain's
// sealed-epoch frontier. This is a ONE-SHOT, raise-only, capped correction.
func BackfillTestnetRewardCursors(statedb *state.StateDB) CursorBackfillStats {
	return backfillRewardCursors(statedb, testnetCursorBackfill)
}

func backfillRewardCursors(statedb *state.StateDB, pairs []cursorBackfillPair) CursorBackfillStats {
	var stats CursorBackfillStats
	if len(pairs) == 0 {
		return stats
	}

	currentSealedEpoch := cursorCurrentSealedEpoch(statedb)

	for _, pair := range pairs {
		corrected := new(big.Int).SetUint64(pair.correctedEpoch)

		// Defensive cap: never raise a cursor past the chain's sealed-epoch
		// frontier. _newRewardsOf integrates ARPT up to currentSealedEpoch, so
		// a cursor above it can never under-integrate into a positive reward
		// but could leave the cursor in an inconsistent forward position.
		if currentSealedEpoch.Sign() != 0 && corrected.Cmp(currentSealedEpoch) > 0 {
			corrected = new(big.Int).Set(currentSealedEpoch)
		}

		current := cursorGet(statedb, pair.delegator, pair.validatorID)

		// Raise-only: require correctedEpoch strictly greater than the stored
		// value. Lowering a cursor could expand the integrated reward window
		// and over-mint, so a corrected value that is not an increase is a
		// no-op (defensive against a stale list or a re-run after partial
		// application).
		if corrected.Cmp(current) <= 0 {
			continue
		}

		cursorSet(statedb, pair.delegator, pair.validatorID, corrected)
		stats.Raised++
		log.Info("Raised SFC reward cursor (SfcV2Patch7 backfill)",
			"delegator", pair.delegator.Hex(),
			"validatorID", pair.validatorID,
			"from", current.String(),
			"to", corrected.String())
	}

	return stats
}

func cursorCurrentSealedEpoch(statedb *state.StateDB) *big.Int {
	return statedb.GetState(ContractAddress, cursorUint64Hash(sfcCurrentSealedEpochSlot)).Big()
}

// sfcCurrentSealedEpochSlot is the SFC storage slot of currentSealedEpoch
// (slot 103, one above stashedRewardsUntilEpoch's base slot 114's predecessors
// — see the layout comment on sfcStashedRewardsUntilEpochSlot).
const sfcCurrentSealedEpochSlot uint64 = 103

func cursorGet(statedb *state.StateDB, delegator common.Address, validatorID uint64) *big.Int {
	return statedb.GetState(ContractAddress, cursorSlot(delegator, validatorID)).Big()
}

func cursorSet(statedb *state.StateDB, delegator common.Address, validatorID uint64, value *big.Int) {
	statedb.SetState(ContractAddress, cursorSlot(delegator, validatorID), common.BigToHash(value))
}

// cursorSlot computes the storage key for
// stashedRewardsUntilEpoch[delegator][validatorID] using Solidity's nested
// mapping layout: key = keccak(abi.encode(validatorID, keccak(abi.encode(
// delegator, baseSlot)))). Mirrors patch6NestedAddressUintSlot but pinned to
// the stashedRewardsUntilEpoch base slot.
func cursorSlot(delegator common.Address, validatorID uint64) common.Hash {
	var enc [64]byte
	copy(enc[12:32], delegator.Bytes())
	copy(enc[32:64], cursorUint64Hash(sfcStashedRewardsUntilEpochSlot).Bytes())
	inner := crypto.Keccak256Hash(enc[:])

	copy(enc[0:32], cursorUint64Hash(validatorID).Bytes())
	copy(enc[32:64], inner.Bytes())
	return crypto.Keccak256Hash(enc[:])
}

func cursorUint64Hash(v uint64) common.Hash {
	return common.BigToHash(new(big.Int).SetUint64(v))
}
