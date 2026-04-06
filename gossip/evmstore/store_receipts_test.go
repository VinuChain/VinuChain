package evmstore

import (
	"math/big"
	"testing"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"

	"github.com/Fantom-foundation/go-opera/logger"
)

func equalStorageReceipts(t *testing.T, expect, got []*types.ReceiptForStorage) {
	assert.EqualValues(t, len(expect), len(got))
	for i := range expect {
		assert.EqualValues(t, expect[i].CumulativeGasUsed, got[i].CumulativeGasUsed)
		assert.EqualValues(t, expect[i].Logs, got[i].Logs)
		assert.EqualValues(t, expect[i].Status, got[i].Status)
		assert.Equal(t, expect[i].FeeRefund, got[i].FeeRefund)
	}
}

func TestStoreGetCachedReceipts(t *testing.T) {
	logger.SetTestMode(t)

	block, expect := fakeReceipts()
	store := cachedStore()
	store.SetRawReceipts(block, expect)

	got, _ := store.GetRawReceipts(block)
	assert.EqualValues(t, expect, got)
}

func TestStoreGetNonCachedReceipts(t *testing.T) {
	logger.SetTestMode(t)

	block, expect := fakeReceipts()
	store := nonCachedStore()
	store.SetRawReceipts(block, expect)

	got, _ := store.GetRawReceipts(block)
	equalStorageReceipts(t, expect, got)
}

func BenchmarkStoreGetRawReceipts(b *testing.B) {
	logger.SetTestMode(b)

	b.Run("cache on", func(b *testing.B) {
		benchStoreGetRawReceipts(b, cachedStore())
	})
	b.Run("cache off", func(b *testing.B) {
		benchStoreGetRawReceipts(b, nonCachedStore())
	})
}

func benchStoreGetRawReceipts(b *testing.B, store *Store) {
	block, receipt := fakeReceipts()

	store.SetRawReceipts(block, receipt)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if v, _ := store.GetRawReceipts(block); v == nil {
			b.Fatal("invalid result")
		}
	}
}

func BenchmarkStoreSetRawReceipts(b *testing.B) {
	logger.SetTestMode(b)

	b.Run("cache on", func(b *testing.B) {
		benchStoreSetRawReceipts(b, cachedStore())
	})
	b.Run("cache off", func(b *testing.B) {
		benchStoreSetRawReceipts(b, nonCachedStore())
	})
}

func benchStoreSetRawReceipts(b *testing.B, store *Store) {
	block, receipt := fakeReceipts()

	for i := 0; i < b.N; i++ {
		store.SetRawReceipts(block, receipt)
	}
}

func fakeReceipts() (idx.Block, []*types.ReceiptForStorage) {
	return idx.Block(1),
		[]*types.ReceiptForStorage{
			{
				PostState:         nil,
				Status:            0,
				CumulativeGasUsed: 0,
				Bloom:             types.Bloom{},
				Logs:              []*types.Log{},
				TxHash:            common.Hash{},
				ContractAddress:   common.Address{},
				GasUsed:           0,
				BlockHash:         common.Hash{},
				BlockNumber:       nil,
				TransactionIndex:  0,
				FeeRefund:         big.NewInt(12345),
			},
		}
}

func TestStoreGetCachedReceiptsWithFeeRefund(t *testing.T) {
	logger.SetTestMode(t)

	block := idx.Block(2)
	feeRefund := big.NewInt(9876543210)
	receipts := []*types.ReceiptForStorage{
		{
			Status:            1,
			CumulativeGasUsed: 21000,
			Logs:              []*types.Log{},
			FeeRefund:         feeRefund,
		},
	}

	cached := cachedStore()
	cached.SetRawReceipts(block, receipts)
	gotCached, _ := cached.GetRawReceipts(block)
	assert.EqualValues(t, len(receipts), len(gotCached))
	assert.Equal(t, feeRefund, gotCached[0].FeeRefund)

	nonCached := nonCachedStore()
	nonCached.SetRawReceipts(block, receipts)
	gotNonCached, _ := nonCached.GetRawReceipts(block)
	assert.EqualValues(t, len(receipts), len(gotNonCached))
	assert.Equal(t, feeRefund, gotNonCached[0].FeeRefund)
}
