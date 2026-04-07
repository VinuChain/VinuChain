package evmstore

import (
	"testing"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-opera/logger"
)

func TestPruneReceiptsUpTo_DeletesOldReceipts(t *testing.T) {
	logger.SetTestMode(t)
	store := cachedStore()

	receipt := []*types.ReceiptForStorage{
		{Status: 1, CumulativeGasUsed: 21000, Logs: []*types.Log{}},
	}
	for _, n := range []idx.Block{1, 2, 3, 4, 5} {
		store.SetRawReceipts(n, receipt)
	}

	count, err := store.PruneReceiptsUpTo(3)
	require.NoError(t, err)
	assert.Equal(t, 3, count)

	assert.Nil(t, store.GetRawReceiptsRLP(1))
	assert.Nil(t, store.GetRawReceiptsRLP(2))
	assert.Nil(t, store.GetRawReceiptsRLP(3))
	assert.NotNil(t, store.GetRawReceiptsRLP(4))
	assert.NotNil(t, store.GetRawReceiptsRLP(5))
}

func TestPruneReceiptsUpTo_DeletesTxPositions(t *testing.T) {
	logger.SetTestMode(t)
	store := cachedStore()

	txhash := common.HexToHash("0xdeadbeef")
	pos := TxPosition{Block: 2}
	store.SetTxPosition(txhash, pos)

	count, err := store.PruneReceiptsUpTo(2)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, count, 1)

	assert.Nil(t, store.GetTxPosition(txhash))
}

func TestPruneReceiptsUpTo_KeepsTxPositionsAboveThreshold(t *testing.T) {
	logger.SetTestMode(t)
	store := cachedStore()

	kept := common.HexToHash("0xbeeffeed")
	store.SetTxPosition(kept, TxPosition{Block: 10})

	_, err := store.PruneReceiptsUpTo(5)
	require.NoError(t, err)

	assert.NotNil(t, store.GetTxPosition(kept))
}
