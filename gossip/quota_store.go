package gossip

import (
	"math/big"

	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// QuotaStore is a wrapper around gossip.Store
type QuotaStore struct {
	store *Store
}

// NewQuotaStore creates a new QuotaStore
func NewQuotaStore(s *Store) *QuotaStore {
	return &QuotaStore{s}
}

// GetLatestBlockIndex returns the latest block index
func (qs *QuotaStore) GetLatestBlockIndex() uint64 {
	return uint64(qs.store.GetLatestBlockIndex())
}

// GetBlockTransactionsAndReceipts returns transactions and receipts for the block with index i
func (qs *QuotaStore) GetBlockTransactionsAndReceipts(n uint64) (txs types.Transactions, receipts types.Receipts) {
	idxBlock := idx.Block(n)
	block := qs.getBlock(idxBlock)
	txs = qs.store.GetBlockTxs(idxBlock, block)
	receipts = qs.store.EvmStore().GetReceipts(
		idxBlock,
		types.LatestSignerForChainID(qs.getChainId()),
		common.Hash{},
		txs,
	)
	return
}

// GetBlock returns the block with index i
func (qs *QuotaStore) getBlock(i idx.Block) *inter.Block {
	return qs.store.GetBlock(i)
}

// GetChainId returns the chain id
func (qs *QuotaStore) getChainId() *big.Int {
	return qs.store.GetEvmChainConfig().ChainID
}
