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
	return isInterrupted(chaindataDir) || isEmpty(chaindataDir)
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
