package integration

import (
	"os"
	"path"
)

func isInterrupted(chaindataDir string) bool {
	_, err := os.Stat(path.Join(chaindataDir, "unfinished"))
	return err == nil
}

// FirstLaunchPending reports whether the next MakeEngine call will apply a
// genesis to chaindataDir: the directory is empty, or holds only the remains
// of an interrupted genesis processing (which get dropped and re-applied).
func FirstLaunchPending(chaindataDir string) bool {
	return isInterrupted(chaindataDir) || isEmpty(chaindataDir) || isEmptyDBSkeleton(chaindataDir)
}

// isEmptyDBSkeleton recognizes a MakeEngine crash between creating its database
// directories and writing the unfinished marker.
func isEmptyDBSkeleton(chaindataDir string) bool {
	entries, err := os.ReadDir(chaindataDir)
	if err != nil || len(entries) == 0 {
		return false
	}
	for _, entry := range entries {
		switch entry.Name() {
		case "leveldb-fsh", "leveldb-flg", "leveldb-drc", "pebble-fsh", "pebble-flg", "pebble-drc":
		default:
			return false
		}
		contents, err := os.ReadDir(path.Join(chaindataDir, entry.Name()))
		if err != nil || !entry.IsDir() || len(contents) != 0 {
			return false
		}
	}
	return true
}

func setGenesisProcessing(chaindataDir string) {
	f, _ := os.Create(path.Join(chaindataDir, "unfinished"))
	if f != nil {
		_ = f.Close()
	}
}

func setGenesisComplete(chaindataDir string) {
	_ = os.Remove(path.Join(chaindataDir, "unfinished"))
}
