package drivercall

import (
	"math/big"
	"testing"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/inter/validatorpk"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/genesis/gpos"
)

func TestSealEpochValidators_ProducesValidCalldata(t *testing.T) {
	validators := []idx.ValidatorID{1, 2, 3}
	data := SealEpochValidators(validators)
	require.NotNil(t, data)
	require.True(t, len(data) > 4, "calldata should contain method ID + encoded args")
}

func TestSealEpochValidators_EmptyList(t *testing.T) {
	data := SealEpochValidators([]idx.ValidatorID{})
	require.NotNil(t, data)
	require.True(t, len(data) > 4)
}

func TestSealEpoch_ProducesValidCalldata(t *testing.T) {
	metrics := []ValidatorEpochMetric{
		{
			Missed: opera.BlocksMissed{
				BlocksNum: 10,
				Period:    inter.Timestamp(1000),
			},
			Uptime:          inter.Timestamp(5000),
			OriginatedTxFee: big.NewInt(1e18),
		},
	}
	data := SealEpoch(metrics)
	require.NotNil(t, data)
	require.True(t, len(data) > 4)
}

func TestSealEpoch_EmptyMetrics(t *testing.T) {
	data := SealEpoch([]ValidatorEpochMetric{})
	require.NotNil(t, data)
	require.True(t, len(data) > 4)
}

func TestSetGenesisValidator_ProducesValidCalldata(t *testing.T) {
	v := gpos.Validator{
		ID:      1,
		Address: common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
		PubKey: validatorpk.PubKey{
			Raw:  make([]byte, 65),
			Type: validatorpk.Types.Secp256k1,
		},
		CreationTime:     inter.Timestamp(1608600000),
		CreationEpoch:    1,
		DeactivatedTime:  0,
		DeactivatedEpoch: 0,
		Status:           0,
	}
	data := SetGenesisValidator(v)
	require.NotNil(t, data)
	require.True(t, len(data) > 4)
}

func TestSetGenesisDelegation_ProducesValidCalldata(t *testing.T) {
	d := Delegation{
		Address:            common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
		ValidatorID:        1,
		Stake:              big.NewInt(1e18),
		LockedStake:        new(big.Int),
		LockupFromEpoch:    0,
		LockupEndTime:      0,
		LockupDuration:     0,
		EarlyUnlockPenalty: new(big.Int),
		Rewards:            new(big.Int),
	}
	data := SetGenesisDelegation(d)
	require.NotNil(t, data)
	require.True(t, len(data) > 4)
}

func TestDeactivateValidator_ProducesValidCalldata(t *testing.T) {
	data := DeactivateValidator(1, 1)
	require.NotNil(t, data)
	require.True(t, len(data) > 4)
}

func TestAllFunctions_CalldataStartsWithMethodSelector(t *testing.T) {
	t.Run("SealEpochValidators", func(t *testing.T) {
		data := SealEpochValidators([]idx.ValidatorID{1})
		method, ok := sAbi.Methods["sealEpochValidators"]
		require.True(t, ok)
		require.Equal(t, method.ID, data[:4])
	})

	t.Run("SealEpoch", func(t *testing.T) {
		metrics := []ValidatorEpochMetric{{
			Missed:          opera.BlocksMissed{},
			Uptime:          0,
			OriginatedTxFee: new(big.Int),
		}}
		data := SealEpoch(metrics)
		method, ok := sAbi.Methods["sealEpoch"]
		require.True(t, ok)
		require.Equal(t, method.ID, data[:4])
	})

	t.Run("SetGenesisValidator", func(t *testing.T) {
		v := gpos.Validator{
			ID:      1,
			Address: common.HexToAddress("0x01"),
			PubKey: validatorpk.PubKey{
				Raw:  make([]byte, 65),
				Type: validatorpk.Types.Secp256k1,
			},
		}
		data := SetGenesisValidator(v)
		method, ok := sAbi.Methods["setGenesisValidator"]
		require.True(t, ok)
		require.Equal(t, method.ID, data[:4])
	})

	t.Run("SetGenesisDelegation", func(t *testing.T) {
		d := Delegation{
			Address:            common.HexToAddress("0x01"),
			ValidatorID:        1,
			Stake:              new(big.Int),
			LockedStake:        new(big.Int),
			EarlyUnlockPenalty: new(big.Int),
			Rewards:            new(big.Int),
		}
		data := SetGenesisDelegation(d)
		method, ok := sAbi.Methods["setGenesisDelegation"]
		require.True(t, ok)
		require.Equal(t, method.ID, data[:4])
	})

	t.Run("DeactivateValidator", func(t *testing.T) {
		data := DeactivateValidator(1, 0)
		method, ok := sAbi.Methods["deactivateValidator"]
		require.True(t, ok)
		require.Equal(t, method.ID, data[:4])
	})
}
