package quota

import (
	"github.com/ethereum/go-ethereum/core/types"
)

// This interface is used by the QuotaService to get information about the chain
// It can be implemented by wraped gossip.Store with additional methods
type Store interface {
	// GetLatestBlockIndex() returns the latest block index
	GetLatestBlockIndex() uint64
	// GetBlockTransactionsAndReceipts returns transactions and receipts for the block with index i
	GetBlockTransactionsAndReceipts(n uint64) (types.Transactions, types.Receipts)
}
