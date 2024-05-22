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

// PaybackStore is a wrapper around gossip.Store
type PaybackStore struct {
	store *Store
}

// NewPaybackStore creates a new PaybackStore
func NewPaybackStore(s *Store) *PaybackStore {
	return &PaybackStore{s}
}

// GetLatestBlockIndex returns the latest block index
func (ps *PaybackStore) GetLatestBlockIndex() uint64 {
	return uint64(ps.store.GetLatestBlockIndex())
}

// GetBlockTransactionsAndReceipts returns transactions and receipts for the block with index i
func (ps *PaybackStore) GetBlockTransactionsAndReceipts(n uint64) (txs types.Transactions, receipts types.Receipts) {
	idxBlock := idx.Block(n)
	block := ps.getBlock(idxBlock)
	txs = ps.store.GetBlockTxs(idxBlock, block)
	receipts = ps.store.EvmStore().GetReceipts(
		idxBlock,
		types.LatestSignerForChainID(ps.getChainId()),
		common.Hash{},
		txs,
	)
	return
}

// GetBlock returns the block with index i
func (ps *PaybackStore) getBlock(i idx.Block) *inter.Block {
	return ps.store.GetBlock(i)
}

// GetChainId returns the chain id
func (ps *PaybackStore) getChainId() *big.Int {
	return ps.store.GetEvmChainConfig().ChainID
}

// FindBlockEpoch returns the epoch of the block with index i
func (ps *PaybackStore) FindBlockEpoch(i idx.Block) idx.Epoch {
	return ps.store.FindBlockEpoch(i)
}

// GetHistoryEpochState returns the state of the epoch with index i
func (ps *PaybackStore) GetHistoryEpochState(i idx.Epoch) *iblockproc.EpochState {
	return ps.store.GetHistoryEpochState(i)
}

// GetRules returns the rules
func (ps *PaybackStore) GetRules() opera.Rules {
	return ps.store.GetRules()
}

// GetCurrentEpoch returns the current epoch
func (ps *PaybackStore) GetCurrentEpoch() idx.Epoch {
	return ps.store.GetEpoch()
}

// GetBlock returns the block
func (ps *PaybackStore) GetBlock(i idx.Block) *inter.Block {
	return ps.store.GetBlock(i)
}
