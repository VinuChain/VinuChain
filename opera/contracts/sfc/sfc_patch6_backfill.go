package sfc

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	sfcGetStakeSlot      uint64 = 116
	sfcStakesSlot        uint64 = 125
	sfcStakePositionSlot uint64 = 126

	stakeWithoutAmountStorageSlots int64 = 2
)

type patch6DelegationBackfill struct {
	delegator   common.Address
	validatorID uint64
}

// Patch6BackfillStats reports deterministic storage repairs applied at the
// SfcV2Patch6 activation boundary.
type Patch6BackfillStats struct {
	Appended uint64
	Repaired uint64
}

// Changed reports whether the backfill wrote any SFC storage.
func (stats Patch6BackfillStats) Changed() bool {
	return stats.Appended != 0 || stats.Repaired != 0
}

var patch6TestnetDelegationBackfill = []patch6DelegationBackfill{
	patch6Backfill("0x2697b646c356f2cbd0ed1f24b9382683b7f437f9", 1),
	patch6Backfill("0x2dfe6aac8b6666424ef5e20a06861aac46d410de", 1),
	patch6Backfill("0x35bc53bfd049051724dc3199e09c7bd790dbc4e3", 1),
}

// Snapshot from live mainnet on 2026-05-17: non-zero getStake pairs present in
// staking GraphQL but missing from the SFC stakes[] enumeration. Refresh this
// list immediately before any mainnet SfcV2 activation release.
var mainnetSfcV2DelegationBackfill = []patch6DelegationBackfill{
	patch6Backfill("0x397a403C884CF71fCBd6a78e881cc00743d33472", 1),
	patch6Backfill("0xbF6042Fe73190b7A70c18ca589D040C8472C5d93", 1),
	patch6Backfill("0xD42c232373f306dD177354804A46Bf4e8A9FFcb5", 1),
	patch6Backfill("0xD9399d9466fcaFAD41E0E684185F0A4f17496410", 1),
	patch6Backfill("0x5D65C61900CEAb7Be7CEbb9040696A27522B3A4E", 6),
	patch6Backfill("0x93a590DbCFAc79950bF6422A5b70790C9729c9FE", 6),
	patch6Backfill("0xe817875283Efe9d4a9E6766C94579361EeBBea09", 6),
	patch6Backfill("0xF93cAf804d130ef6348138F10996bdcC55d394C2", 6),
	patch6Backfill("0x1d6b8A1303A7Ac1d163979486Dc6B399f1903Eed", 11),
	patch6Backfill("0x2DFE6AAC8b6666424EF5E20A06861AAc46d410dE", 11),
	patch6Backfill("0x96583A7D0A510d2595Cb86a7369AFe762A5778F4", 11),
	patch6Backfill("0xe89Ab3e678B34C7201551439c29E2eeFff1dfD72", 11),
	patch6Backfill("0x160462b1902A7B1454D19813f0017723BB23c048", 12),
	patch6Backfill("0xAF2B98736D027D0667f4269aF813d5F5431f92F2", 12),
	patch6Backfill("0xbAa41CC56Ce239eE616b5E5A6dAa8bE364fD1B95", 12),
	patch6Backfill("0xD294Cac5F20F3799f6583a22f82955631A6A9718", 12),
	patch6Backfill("0x2133a6261ff10322EE5ebDe0920afbe9EC68E2Be", 13),
	patch6Backfill("0x2DFE6AAC8b6666424EF5E20A06861AAc46d410dE", 13),
	patch6Backfill("0x8bceFCE8363a639a2a97132ac0920A2e7D12faFA", 13),
	patch6Backfill("0xAE1C67a6411912fFF867DcB0CeC275d82E656537", 13),
	patch6Backfill("0xf10f35Cc6c326F5d7C79Ecab22c2297ebCc87A0b", 13),
	patch6Backfill("0x12124D1f28ab4F68917e153A4Ba41Ae240a23efD", 14),
	patch6Backfill("0x19C2eC706e7245C32a6918406ca4Dcc6c8A8A6fC", 14),
	patch6Backfill("0xD9399d9466fcaFAD41E0E684185F0A4f17496410", 14),
	patch6Backfill("0xf10f35Cc6c326F5d7C79Ecab22c2297ebCc87A0b", 14),
	patch6Backfill("0xF93cAf804d130ef6348138F10996bdcC55d394C2", 14),
	patch6Backfill("0xfbd3AEb170f1cc319a50C47041062e09d5D71e1f", 14),
	patch6Backfill("0xe030322AC2526027e60e988906Ed64Df7E8C45f3", 15),
	patch6Backfill("0xFe1c93EB68e1D1Dbd4185f58CA0193b37B0eD689", 15),
	patch6Backfill("0x05A7C107812e4e0cb31CB159f4818A7aC20C7736", 16),
	patch6Backfill("0x0dee0083a3B170813dc536AE1fe4D85d3A375838", 16),
	patch6Backfill("0x3F43B8080a0f3Bf13A822C839f1d189072E11731", 16),
	patch6Backfill("0x42cbDeEb035655dEe0a71d50ed2580B725526674", 16),
	patch6Backfill("0x64758ef549B0e7714C2c69aE6097810D3c970d69", 16),
	patch6Backfill("0x660B13390ECC80736737acaFfe6325AA504d4EC2", 16),
	patch6Backfill("0x8C88273A678B5957dA7A4E004b98EFf16100E08b", 16),
	patch6Backfill("0x9d754Ffd84c5A8925FEad37bf7B1Fd4FbA40f48e", 16),
	patch6Backfill("0xb13053Afa9Ace2ba16B1B3A8ae4c69465CEe24b8", 16),
	patch6Backfill("0xB4F24E88e4b7863E43C8b4DaF475CA4160542144", 16),
	patch6Backfill("0xb70842eB8b2d820969D654ca4fc3B274D8953db4", 16),
	patch6Backfill("0xbEA395DD3AE7cE1Ec5fE2a909fe5DE46e63126Fa", 16),
	patch6Backfill("0xf5cb9BC697118600D7546Cb190e216159eb052C5", 16),
	patch6Backfill("0x3f2B98A836CC2199daa74DD1fd9726279E507FDc", 18),
	patch6Backfill("0x42cbDeEb035655dEe0a71d50ed2580B725526674", 18),
	patch6Backfill("0x50DD7B0F52434D58C793695505d3D722F0052Cd4", 18),
	patch6Backfill("0x7D757f62334d511103cC569c2BECEA6d18Ae3801", 18),
	patch6Backfill("0x8C88273A678B5957dA7A4E004b98EFf16100E08b", 18),
	patch6Backfill("0x93a590DbCFAc79950bF6422A5b70790C9729c9FE", 18),
	patch6Backfill("0xE4892C7156c9a8225679c0aa78BB8d791d6dC213", 18),
	patch6Backfill("0xF93cAf804d130ef6348138F10996bdcC55d394C2", 18),
	patch6Backfill("0x0dee0083a3B170813dc536AE1fe4D85d3A375838", 20),
	patch6Backfill("0x3888Ac6e79d261fe18ACF05677D17E644cae24ed", 20),
	patch6Backfill("0x4ddA7236b13Af67f5f41AC2C05dB21B6EdD8Bce6", 20),
	patch6Backfill("0x7f2F5245a8407A6568f64a6E1D70B5606cc253fe", 20),
	patch6Backfill("0xC14C601Cd7E9B72cBab29C41184900d8b592A481", 20),
	patch6Backfill("0xcB67fF716717FCe7aCE371804eae96C1c88A81dC", 20),
	patch6Backfill("0xDD63524c42732701b1C94ea8e9ECb021730Ac3c1", 20),
	patch6Backfill("0xEE611D8147fA75CB432b7672e5d6F349c8Cf12Ee", 20),
	patch6Backfill("0xF93cAf804d130ef6348138F10996bdcC55d394C2", 20),
	patch6Backfill("0x660B13390ECC80736737acaFfe6325AA504d4EC2", 21),
	patch6Backfill("0x252Ae5eE831763767962CECa5c67EadC9C05fE72", 23),
	patch6Backfill("0x0268C9524B6892aE40103E032ab996133Bc70AF7", 24),
	patch6Backfill("0x07722460E3d28E29432EDe52b1A3973E75Bc5438", 24),
	patch6Backfill("0x55F0ada5C312CD703b9555FF908E638D7f6E1225", 24),
	patch6Backfill("0x3f981CD5418A22CCb642185184E7247D6CDC3AAc", 25),
	patch6Backfill("0x11F1dDce1c588a0E8CFA4EF1c61e607B2eD192be", 26),
	patch6Backfill("0x2eA34893C3c7513A9457F782Ea35AC4390619992", 26),
	patch6Backfill("0x30d730300810CAAEE0F557288c49293d4EEB8fDA", 26),
	patch6Backfill("0x8C88273A678B5957dA7A4E004b98EFf16100E08b", 26),
	patch6Backfill("0xd253C9E4C56510E7f5321D2201539138495Ff2ED", 26),
	patch6Backfill("0xF93cAf804d130ef6348138F10996bdcC55d394C2", 26),
	patch6Backfill("0x0c5631F103e58F30fd49DcDb0160f19eDf80147C", 27),
	patch6Backfill("0x3c2e4BdAd51e47e19769e9530704d164B3953d35", 27),
	patch6Backfill("0xb13053Afa9Ace2ba16B1B3A8ae4c69465CEe24b8", 27),
	patch6Backfill("0xe817875283Efe9d4a9E6766C94579361EeBBea09", 27),
	patch6Backfill("0xf10f35Cc6c326F5d7C79Ecab22c2297ebCc87A0b", 27),
	patch6Backfill("0x5D65C61900CEAb7Be7CEbb9040696A27522B3A4E", 28),
	patch6Backfill("0xB3F913E5959e77070e44Ee2fEbff1174D9eE2060", 28),
	patch6Backfill("0x07722460E3d28E29432EDe52b1A3973E75Bc5438", 30),
	patch6Backfill("0x8C88273A678B5957dA7A4E004b98EFf16100E08b", 30),
	patch6Backfill("0xF93cAf804d130ef6348138F10996bdcC55d394C2", 30),
	patch6Backfill("0x93a590DbCFAc79950bF6422A5b70790C9729c9FE", 31),
}

func patch6Backfill(delegator string, validatorID uint64) patch6DelegationBackfill {
	return patch6DelegationBackfill{
		delegator:   common.HexToAddress(delegator),
		validatorID: validatorID,
	}
}

// BackfillPatch6TestnetDelegations repairs the live testnet delegation
// enumeration holes found before SfcV2Patch6 activation. The affected pairs
// have non-zero getStake but either no stakePosition or a stale stakePosition
// pointing at a different stakes[] row.
func BackfillPatch6TestnetDelegations(statedb *state.StateDB, blockTimestamp uint64) Patch6BackfillStats {
	return backfillDelegations(statedb, blockTimestamp, patch6TestnetDelegationBackfill)
}

// BackfillSfcV2MainnetDelegations repairs mainnet delegation enumeration holes
// when mainnet first activates SfcV2 and installs the latest SFC bytecode.
func BackfillSfcV2MainnetDelegations(statedb *state.StateDB, blockTimestamp uint64) Patch6BackfillStats {
	return backfillDelegations(statedb, blockTimestamp, mainnetSfcV2DelegationBackfill)
}

func backfillDelegations(statedb *state.StateDB, blockTimestamp uint64, pairs []patch6DelegationBackfill) Patch6BackfillStats {
	var stats Patch6BackfillStats
	if patch6StakesLength(statedb).Sign() == 0 {
		return stats
	}

	for _, pair := range pairs {
		if patch6GetStake(statedb, pair.delegator, pair.validatorID).Sign() == 0 {
			continue
		}
		if position, ok := patch6FindStakePosition(statedb, pair.delegator, pair.validatorID); ok {
			if stored := patch6StakePosition(statedb, pair.delegator, pair.validatorID); stored.Cmp(position) != 0 {
				patch6SetStakePosition(statedb, pair.delegator, pair.validatorID, position)
				stats.Repaired++
			}
			continue
		}
		patch6AppendStake(statedb, pair.delegator, pair.validatorID, blockTimestamp)
		stats.Appended++
	}

	return stats
}

func patch6GetStake(statedb *state.StateDB, delegator common.Address, validatorID uint64) *big.Int {
	return statedb.GetState(ContractAddress, patch6NestedAddressUintSlot(delegator, validatorID, sfcGetStakeSlot)).Big()
}

func patch6StakePosition(statedb *state.StateDB, delegator common.Address, validatorID uint64) *big.Int {
	return statedb.GetState(ContractAddress, patch6NestedAddressUintSlot(delegator, validatorID, sfcStakePositionSlot)).Big()
}

func patch6SetStakePosition(statedb *state.StateDB, delegator common.Address, validatorID uint64, position *big.Int) {
	statedb.SetState(ContractAddress, patch6NestedAddressUintSlot(delegator, validatorID, sfcStakePositionSlot), common.BigToHash(position))
}

func patch6StakesLength(statedb *state.StateDB) *big.Int {
	return statedb.GetState(ContractAddress, patch6Uint64Hash(sfcStakesSlot)).Big()
}

func patch6SetStakesLength(statedb *state.StateDB, length *big.Int) {
	statedb.SetState(ContractAddress, patch6Uint64Hash(sfcStakesSlot), common.BigToHash(length))
}

func patch6FindStakePosition(statedb *state.StateDB, delegator common.Address, validatorID uint64) (*big.Int, bool) {
	length := patch6StakesLength(statedb)
	if !length.IsUint64() {
		return nil, false
	}
	for i := uint64(1); i < length.Uint64(); i++ {
		position := new(big.Int).SetUint64(i)
		storedDelegator, storedValidatorID := patch6StakeAt(statedb, position)
		if storedDelegator == delegator && storedValidatorID == validatorID {
			return position, true
		}
	}
	return nil, false
}

func patch6AppendStake(statedb *state.StateDB, delegator common.Address, validatorID uint64, blockTimestamp uint64) {
	position := patch6StakesLength(statedb)
	entrySlot := patch6StakeEntrySlot(position)
	statedb.SetState(ContractAddress, entrySlot, patch6PackStakeHeader(delegator, blockTimestamp))
	statedb.SetState(ContractAddress, patch6AddSlot(entrySlot, 1), common.BigToHash(new(big.Int).SetUint64(validatorID)))
	patch6SetStakePosition(statedb, delegator, validatorID, position)
	patch6SetStakesLength(statedb, new(big.Int).Add(position, big.NewInt(1)))
}

func patch6StakeAt(statedb *state.StateDB, position *big.Int) (common.Address, uint64) {
	entrySlot := patch6StakeEntrySlot(position)
	packed := statedb.GetState(ContractAddress, entrySlot).Big()
	delegator := common.BigToAddress(new(big.Int).And(packed, patch6AddressMask()))
	validatorID := statedb.GetState(ContractAddress, patch6AddSlot(entrySlot, 1)).Big()
	if !validatorID.IsUint64() {
		return delegator, 0
	}
	return delegator, validatorID.Uint64()
}

func patch6PackStakeHeader(delegator common.Address, blockTimestamp uint64) common.Hash {
	packed := new(big.Int).SetBytes(delegator.Bytes())
	timestamp := new(big.Int).Lsh(new(big.Int).SetUint64(blockTimestamp), 160)
	return common.BigToHash(packed.Or(packed, timestamp))
}

func patch6StakeEntrySlot(position *big.Int) common.Hash {
	offset := new(big.Int).Mul(new(big.Int).Set(position), big.NewInt(stakeWithoutAmountStorageSlots))
	return patch6AddSlotBig(patch6DynamicArrayBaseSlot(sfcStakesSlot), offset)
}

func patch6AddSlot(slot common.Hash, offset int64) common.Hash {
	return common.BigToHash(new(big.Int).Add(slot.Big(), big.NewInt(offset)))
}

func patch6AddSlotBig(slot common.Hash, offset *big.Int) common.Hash {
	return common.BigToHash(new(big.Int).Add(slot.Big(), offset))
}

func patch6DynamicArrayBaseSlot(slot uint64) common.Hash {
	return crypto.Keccak256Hash(patch6Uint64Hash(slot).Bytes())
}

func patch6NestedAddressUintSlot(addr common.Address, key uint64, slot uint64) common.Hash {
	var enc [64]byte
	copy(enc[12:32], addr.Bytes())
	copy(enc[32:64], patch6Uint64Hash(slot).Bytes())
	inner := crypto.Keccak256Hash(enc[:])

	copy(enc[0:32], patch6Uint64Hash(key).Bytes())
	copy(enc[32:64], inner.Bytes())
	return crypto.Keccak256Hash(enc[:])
}

func patch6Uint64Hash(v uint64) common.Hash {
	return common.BigToHash(new(big.Int).SetUint64(v))
}

func patch6AddressMask() *big.Int {
	return new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 160), big.NewInt(1))
}
