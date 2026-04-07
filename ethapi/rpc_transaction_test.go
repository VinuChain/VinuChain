package ethapi

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

// TestNewRPCTransaction_NilFeeRefund_OmitsField verifies that when feeRefund
// is nil, the JSON output omits the feeRefund field entirely (omitempty).
func TestNewRPCTransaction_NilFeeRefund_OmitsField(t *testing.T) {
	tx := types.NewTransaction(0, common.HexToAddress("0x1234"), big.NewInt(0), 21000, big.NewInt(1e9), nil)
	rpcTx := newRPCTransaction(tx, common.Hash{}, 0, 0, nil, nil)

	data, err := json.Marshal(rpcTx)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	if _, ok := m["feeRefund"]; ok {
		t.Errorf("feeRefund should be omitted when nil, got %v", m["feeRefund"])
	}
}

// TestNewRPCTransaction_WithFeeRefund_IncludesField verifies that when feeRefund
// is provided, the JSON output includes the feeRefund field.
func TestNewRPCTransaction_WithFeeRefund_IncludesField(t *testing.T) {
	tx := types.NewTransaction(0, common.HexToAddress("0x1234"), big.NewInt(0), 21000, big.NewInt(1e9), nil)
	refund := (*hexutil.Big)(big.NewInt(500))
	rpcTx := newRPCTransaction(tx, common.Hash{}, 0, 0, nil, refund)

	data, err := json.Marshal(rpcTx)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	val, ok := m["feeRefund"]
	if !ok {
		t.Fatal("feeRefund should be present when provided")
	}
	if val != "0x1f4" {
		t.Errorf("expected feeRefund=0x1f4, got %v", val)
	}
}
