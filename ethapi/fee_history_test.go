package ethapi

import (
	"math"
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

func TestGasUsedRatio(t *testing.T) {
	cases := []struct {
		name     string
		gasUsed  uint64
		gasLimit uint64
		want     float64
	}{
		{
			name:     "typical block at MaxBlockGas",
			gasUsed:  5_000_000,
			gasLimit: 20_500_000,
			want:     5_000_000.0 / 20_500_000.0,
		},
		{
			name:     "empty block",
			gasUsed:  0,
			gasLimit: 20_500_000,
			want:     0.0,
		},
		{
			name:     "fully used block",
			gasUsed:  20_500_000,
			gasLimit: 20_500_000,
			want:     1.0,
		},
		{
			name:     "zero gas limit must not panic or produce NaN",
			gasUsed:  0,
			gasLimit: 0,
			want:     0.0,
		},
		{
			name:     "zero gas limit with nonzero used must not divide-by-zero",
			gasUsed:  1_000_000,
			gasLimit: 0,
			want:     0.0,
		},
		{
			name:     "malformed header where used exceeds limit is clamped to 1.0",
			gasUsed:  30_000_000,
			gasLimit: 20_500_000,
			want:     1.0,
		},
	}

	const epsilon = 1e-12
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := gasUsedRatio(tc.gasUsed, tc.gasLimit)
			if math.IsNaN(got) {
				t.Fatalf("gasUsedRatio(%d, %d) returned NaN", tc.gasUsed, tc.gasLimit)
			}
			if math.IsInf(got, 0) {
				t.Fatalf("gasUsedRatio(%d, %d) returned Inf", tc.gasUsed, tc.gasLimit)
			}
			if math.Abs(got-tc.want) > epsilon {
				t.Fatalf("gasUsedRatio(%d, %d) = %v, want %v", tc.gasUsed, tc.gasLimit, got, tc.want)
			}
		})
	}
}
