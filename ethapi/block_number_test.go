package ethapi

import (
	"testing"
)

func TestBlockNumber_NilHeader_ReturnsZero(t *testing.T) {
	api := NewPublicBlockChainAPI(&stubBackend{})
	got := api.BlockNumber()
	if got != 0 {
		t.Fatalf("BlockNumber() = %d, want 0", got)
	}
}
