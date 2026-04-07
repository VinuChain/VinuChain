package ethapi

import (
	"strings"
	"testing"
)

func newExtRPCGuardedAPI() *PrivateAccountAPI {
	return NewPrivateAccountAPI(&stubBackend{extRPC: true}, new(AddrLocker))
}

func TestImportRawKey_BlockedOnExtRPC(t *testing.T) {
	api := newExtRPCGuardedAPI()
	_, err := api.ImportRawKey("deadbeef", "password")
	if err == nil {
		t.Fatal("expected error when ExtRPCEnabled is true")
	}
	if !strings.Contains(err.Error(), "not available over external RPC") {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestOpenWallet_BlockedOnExtRPC(t *testing.T) {
	api := newExtRPCGuardedAPI()
	err := api.OpenWallet("http://example.com", nil)
	if err == nil {
		t.Fatal("expected error when ExtRPCEnabled is true")
	}
	if !strings.Contains(err.Error(), "not available over external RPC") {
		t.Fatalf("unexpected error message: %s", err.Error())
	}
}

func TestListWallets_EmptyOnExtRPC(t *testing.T) {
	api := newExtRPCGuardedAPI()
	wallets := api.ListWallets()
	if wallets == nil {
		t.Fatal("expected non-nil empty slice, got nil")
	}
	if len(wallets) != 0 {
		t.Fatalf("expected empty slice, got %d wallets", len(wallets))
	}
}
