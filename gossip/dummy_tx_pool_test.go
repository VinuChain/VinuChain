package gossip

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestDummyTxPoolDeleteFromEmptyPool(t *testing.T) {
	pool := &dummyTxPool{}

	require.NotPanics(t, func() {
		pool.Delete(common.HexToHash("0x01"))
	})
	require.Zero(t, pool.Count())
}
