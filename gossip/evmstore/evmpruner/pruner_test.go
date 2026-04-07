package evmpruner

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
)

// TestRecoverPruning_NoBloomFile verifies that RecoverPruning returns nil
// when no bloom filter file exists on disk (the common case on a clean node).
func TestRecoverPruning_NoBloomFile(t *testing.T) {
	dir := t.TempDir()
	db := rawdb.NewMemoryDatabase()
	defer db.Close()

	err := RecoverPruning(dir, db, common.Hash{})
	if err != nil {
		t.Fatalf("expected nil when no bloom file present, got: %v", err)
	}
}

// TestRecoverPruning_MissingSnapshotWithBloomFile verifies that RecoverPruning
// returns a non-nil error when a bloom filter file exists on disk but no valid
// snapshot is present in the database. The bloom filter's existence means
// pruning was interrupted and must be resumed — but without a snapshot the
// resume cannot succeed.
func TestRecoverPruning_MissingSnapshotWithBloomFile(t *testing.T) {
	dir := t.TempDir()
	db := rawdb.NewMemoryDatabase()
	defer db.Close()

	root := common.HexToHash("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef")
	bloomPath := bloomFilterName(dir, root)

	// Write an empty file so findBloomFilter detects it by filename, but the
	// in-memory DB has no snapshot data — either the bloom read or the snapshot
	// open will fail, producing the required non-nil error.
	if err := os.WriteFile(bloomPath, []byte{}, 0600); err != nil {
		t.Fatalf("failed to create bloom stub: %v", err)
	}

	err := RecoverPruning(dir, db, root)
	if err == nil {
		t.Fatal("expected an error when bloom file exists but snapshot is missing, got nil")
	}
}

// TestFindBloomFilter_WalkError verifies that findBloomFilter propagates
// filesystem walk errors instead of silently swallowing them.
func TestFindBloomFilter_WalkError(t *testing.T) {
	nonExistent := filepath.Join(t.TempDir(), "subdir_does_not_exist")
	_, _, err := findBloomFilter(nonExistent)
	if err == nil {
		t.Fatal("expected an error from findBloomFilter for non-existent path, got nil")
	}
}

// TestPrunerSnapshotsCount verifies that snapshotsCount is 128, matching
// the upstream go-ethereum diff layer window size.
func TestPrunerSnapshotsCount(t *testing.T) {
	const want = 128
	if snapshotsCount != want {
		t.Fatalf("snapshotsCount = %d, want %d", snapshotsCount, want)
	}
}
