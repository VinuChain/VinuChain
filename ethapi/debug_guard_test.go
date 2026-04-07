package ethapi

import (
	"context"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

func TestGetBlockRlp_RejectsExternalRPC(t *testing.T) {
	api := NewPublicDebugAPI(&stubBackend{extRPC: true})
	_, err := api.GetBlockRlp(context.Background(), 0)
	if err == nil {
		t.Fatal("expected error when ExtRPCEnabled is true")
	}
	if !strings.Contains(err.Error(), "debug method not available over external RPC") {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestPrintBlock_RejectsExternalRPC(t *testing.T) {
	api := NewPublicDebugAPI(&stubBackend{extRPC: true})
	_, err := api.PrintBlock(context.Background(), 0)
	if err == nil {
		t.Fatal("expected error when ExtRPCEnabled is true")
	}
	if !strings.Contains(err.Error(), "debug method not available over external RPC") {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestBlocksTransactionTimes_RejectsExternalRPC(t *testing.T) {
	api := NewPublicDebugAPI(&stubBackend{extRPC: true})
	_, err := api.BlocksTransactionTimes(context.Background(), rpc.BlockNumber(0), hexutil.Uint64(10))
	if err == nil {
		t.Fatal("expected error when ExtRPCEnabled is true")
	}
	if !strings.Contains(err.Error(), "debug method not available over external RPC") {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}
