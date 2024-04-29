package gossip

import (
	"github.com/Fantom-foundation/go-opera/inter/iblockproc"
	"github.com/Fantom-foundation/go-opera/opera"
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

// FindBlockEpoch returns the epoch of the block with index i
func (qs *QuotaStore) FindBlockEpoch(i idx.Block) idx.Epoch {
	return qs.store.FindBlockEpoch(i)
}

// GetHistoryEpochState returns the state of the epoch with index i
func (qs *QuotaStore) GetHistoryEpochState(i idx.Epoch) *iblockproc.EpochState {
	return qs.store.GetHistoryEpochState(i)
}

// GetRules returns the rules
func (qs *QuotaStore) GetRules() opera.Rules {
	return qs.store.GetRules()
}

// GetCurrentEpoch returns the current epoch
func (qs *QuotaStore) GetCurrentEpoch() idx.Epoch {
	return qs.store.GetEpoch()
}

// GetBlock returns the block
func (qc *QuotaStore) GetBlock(i idx.Block) *inter.Block {
	return qc.store.GetBlock(i)
}
