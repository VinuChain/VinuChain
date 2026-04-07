package ethapi

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestFeeHistoryRewardSliceIndependence(t *testing.T) {
	tips := []*hexutil.Big{
		(*hexutil.Big)(big.NewInt(100)),
		(*hexutil.Big)(big.NewInt(200)),
	}

	const entries = 3
	reward := make([][]*hexutil.Big, 0, entries)
	for i := 0; i < entries; i++ {
		tip := make([]*hexutil.Big, len(tips))
		copy(tip, tips)
		reward = append(reward, tip)
	}

	// Mutate the first entry's first element.
	reward[0][0] = (*hexutil.Big)(big.NewInt(999))

	// All other entries must be unaffected.
	for i := 1; i < entries; i++ {
		if reward[i][0].ToInt().Cmp(big.NewInt(100)) != 0 {
			t.Fatalf("reward[%d][0] = %s, want 100; slice was shared", i, reward[i][0].ToInt())
		}
	}
}
