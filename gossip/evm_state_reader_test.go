package gossip

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-opera/opera"
)

func TestMaxGasLimitForVinuChainMainnetCapsSingleTransactionGas(t *testing.T) {
	cap := maxGasLimitForRules(opera.VinuChainMainNetRules())

	require.Equal(t, uint64(9_980_000), cap)
	require.Less(t, cap, uint64(1<<24), "VinuChain's per-transaction gas cap should remain below Ethereum's Fusaka cap")
}

func TestMaxGasLimitForRulesReturnsZeroWhenEmptyEventExceedsMaxEventGas(t *testing.T) {
	rules := opera.VinuChainMainNetRules()
	rules.Economy.Gas.MaxEventGas = rules.Economy.Gas.EventGas

	require.Zero(t, maxGasLimitForRules(rules))
}
