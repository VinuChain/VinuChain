package verwatcher

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Fantom-foundation/go-opera/opera/contracts/driver"
)

// TestOnNewLog_EmptyTopics verifies that OnNewLog does not panic when
// a log from the driver contract has zero topics.
func TestOnNewLog_EmptyTopics(t *testing.T) {
	w := &VerWarcher{
		store: NewStore(nil),
	}

	log := &types.Log{
		Address: driver.ContractAddress,
		Topics:  []common.Hash{}, // zero topics
	}

	// Must not panic.
	w.OnNewLog(log)
}
