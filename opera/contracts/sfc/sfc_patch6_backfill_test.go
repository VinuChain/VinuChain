package sfc

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
)

func TestBackfillPatch6TestnetDelegations_AppendsMissingStalePositionPair(t *testing.T) {
	statedb := newPatch6BackfillState(t)
	target := patch6TestnetDelegationBackfill[0]
	otherDelegator := common.HexToAddress("0x31882eb0fd71bfb5d88d6559c84d1317fd9d4234")
	const (
		otherValidatorID = uint64(6)
		oldPosition      = uint64(1)
		newPosition      = uint64(2)
		blockTimestamp   = uint64(1_706_000_001)
	)

	patch6SeedStakeEntry(statedb, oldPosition, otherDelegator, otherValidatorID, 1_700_000_000)
	patch6SetStakesLength(statedb, new(big.Int).SetUint64(newPosition))
	patch6SetStake(statedb, target.delegator, target.validatorID, patch6MustBigInt(t, "40000000000000000000"))
	patch6SetStakePosition(statedb, target.delegator, target.validatorID, new(big.Int).SetUint64(oldPosition))

	stats := BackfillPatch6TestnetDelegations(statedb, blockTimestamp)
	if stats.Appended != 1 || stats.Repaired != 0 {
		t.Fatalf("unexpected stats after backfill: %+v", stats)
	}
	if got := patch6StakesLength(statedb).Uint64(); got != newPosition+1 {
		t.Fatalf("unexpected stakes length: got %d, want %d", got, newPosition+1)
	}
	if got := patch6StakePosition(statedb, target.delegator, target.validatorID).Uint64(); got != newPosition {
		t.Fatalf("unexpected target position: got %d, want %d", got, newPosition)
	}
	gotDelegator, gotValidatorID := patch6StakeAt(statedb, new(big.Int).SetUint64(newPosition))
	if gotDelegator != target.delegator || gotValidatorID != target.validatorID {
		t.Fatalf("unexpected appended stake: got (%s,%d), want (%s,%d)",
			gotDelegator, gotValidatorID, target.delegator, target.validatorID)
	}
	if gotTimestamp := patch6StakeTimestamp(statedb, new(big.Int).SetUint64(newPosition)); gotTimestamp != blockTimestamp {
		t.Fatalf("unexpected appended timestamp: got %d, want %d", gotTimestamp, blockTimestamp)
	}

	stats = BackfillPatch6TestnetDelegations(statedb, blockTimestamp+1)
	if stats.Changed() {
		t.Fatalf("second backfill was not idempotent: %+v", stats)
	}
	if got := patch6StakesLength(statedb).Uint64(); got != newPosition+1 {
		t.Fatalf("second backfill changed stakes length: got %d, want %d", got, newPosition+1)
	}
}

func TestBackfillPatch6TestnetDelegations_RepairsStalePositionWithoutDuplicate(t *testing.T) {
	statedb := newPatch6BackfillState(t)
	target := patch6TestnetDelegationBackfill[1]
	staleDelegator := common.HexToAddress("0x5f9ce86cfe3dcd47f2921dd46c15a2a2f7bf156f")
	const (
		stalePosition  = uint64(1)
		actualPosition = uint64(2)
	)

	patch6SeedStakeEntry(statedb, stalePosition, staleDelegator, 12, 1_715_000_000)
	patch6SeedStakeEntry(statedb, actualPosition, target.delegator, target.validatorID, 1_706_000_001)
	patch6SetStakesLength(statedb, big.NewInt(3))
	patch6SetStake(statedb, target.delegator, target.validatorID, big.NewInt(101_000_000_000_000_000))
	patch6SetStakePosition(statedb, target.delegator, target.validatorID, new(big.Int).SetUint64(stalePosition))

	stats := BackfillPatch6TestnetDelegations(statedb, 1_706_000_100)
	if stats.Appended != 0 || stats.Repaired != 1 {
		t.Fatalf("unexpected stats after repair: %+v", stats)
	}
	if got := patch6StakePosition(statedb, target.delegator, target.validatorID).Uint64(); got != actualPosition {
		t.Fatalf("unexpected repaired position: got %d, want %d", got, actualPosition)
	}
	if got := patch6StakesLength(statedb).Uint64(); got != 3 {
		t.Fatalf("repair appended a duplicate stake: length got %d, want 3", got)
	}
}

func TestBackfillPatch6TestnetDelegations_SkipsZeroStake(t *testing.T) {
	statedb := newPatch6BackfillState(t)
	patch6SetStakesLength(statedb, big.NewInt(1))

	stats := BackfillPatch6TestnetDelegations(statedb, 1_706_000_001)
	if stats.Changed() {
		t.Fatalf("zero-stake backfill changed state: %+v", stats)
	}
}

func TestPatch6TestnetDelegationBackfillListPinned(t *testing.T) {
	assertPatch6BackfillList(t, "testnet", patch6TestnetDelegationBackfill, 3)
}

func TestMainnetSfcV2DelegationBackfillListPinned(t *testing.T) {
	assertPatch6BackfillList(t, "mainnet", mainnetSfcV2DelegationBackfill, 82)
}

func assertPatch6BackfillList(t *testing.T, name string, pairs []patch6DelegationBackfill, want int) {
	t.Helper()

	if got := len(pairs); got != want {
		t.Fatalf("unexpected %s backfill list length: got %d, want %d", name, got, want)
	}

	seen := make(map[string]struct{}, len(pairs))
	for _, pair := range pairs {
		if pair.delegator == (common.Address{}) {
			t.Fatalf("%s backfill list contains zero delegator for validator %d", name, pair.validatorID)
		}
		if pair.validatorID == 0 {
			t.Fatalf("%s backfill list contains zero validator for delegator %s", name, pair.delegator)
		}
		key := pair.delegator.Hex() + ":" + new(big.Int).SetUint64(pair.validatorID).String()
		if _, ok := seen[key]; ok {
			t.Fatalf("duplicate %s backfill pair %s", name, key)
		}
		seen[key] = struct{}{}
	}
}

func newPatch6BackfillState(t *testing.T) *state.StateDB {
	t.Helper()

	statedb, err := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	if err != nil {
		t.Fatalf("failed to create state: %v", err)
	}
	return statedb
}

func patch6SetStake(statedb *state.StateDB, delegator common.Address, validatorID uint64, amount *big.Int) {
	statedb.SetState(ContractAddress, patch6NestedAddressUintSlot(delegator, validatorID, sfcGetStakeSlot), common.BigToHash(amount))
}

func patch6MustBigInt(t *testing.T, value string) *big.Int {
	t.Helper()

	amount, ok := new(big.Int).SetString(value, 10)
	if !ok {
		t.Fatalf("invalid big integer %q", value)
	}
	return amount
}

func patch6SeedStakeEntry(statedb *state.StateDB, position uint64, delegator common.Address, validatorID uint64, timestamp uint64) {
	entrySlot := patch6StakeEntrySlot(new(big.Int).SetUint64(position))
	statedb.SetState(ContractAddress, entrySlot, patch6PackStakeHeader(delegator, timestamp))
	statedb.SetState(ContractAddress, patch6AddSlot(entrySlot, 1), common.BigToHash(new(big.Int).SetUint64(validatorID)))
}

func patch6StakeTimestamp(statedb *state.StateDB, position *big.Int) uint64 {
	entrySlot := patch6StakeEntrySlot(position)
	return new(big.Int).Rsh(statedb.GetState(ContractAddress, entrySlot).Big(), 160).Uint64()
}
