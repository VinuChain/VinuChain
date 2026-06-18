package metrics

import (
	"os"
	"path/filepath"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
)

var (
	// TODO: refactor it
	dbDir        atomic.Value
	//nolint:unused // registered as a side-effect metrics gauge; framework calls measureDbDir via callback
	dbSizeMetric = metrics.NewRegisteredFunctionalGauge("db_size", nil, measureDbDir)
)

func SetDataDir(datadir string) {
	dbDir.Store(datadir)
}

//nolint:unused // callback for dbSizeMetric gauge registered above
func measureDbDir() (size int64) {
	datadir, ok := dbDir.Load().(string)
	if !ok || datadir == "" || datadir == "inmemory" {
		return
	}
	return sizeOfDir(datadir)
}

//nolint:unused // helper for measureDbDir
func sizeOfDir(dir string) (size int64) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Debug("datadir walk", "path", path, "err", err)
			return filepath.SkipDir
		}

		if info.IsDir() {
			return nil
		}

		dst, err := filepath.EvalSymlinks(path)
		if err == nil && dst != path {
			size += sizeOfDir(dst)
		} else {
			size += info.Size()
		}

		return nil
	})

	if err != nil {
		log.Debug("datadir walk", "path", dir, "err", err)
	}

	return
}
