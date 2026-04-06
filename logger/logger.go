package logger

import (
	"github.com/ethereum/go-ethereum/log"
)

func init() {
	log.Root().SetHandler(
		log.CallerStackHandler("%v", log.StdoutHandler))
}

// SetDSN appends a Sentry error-reporting handler to the log root handler.
// If value is empty, the call is a no-op and a warning is emitted.
func SetDSN(value string) {
	if value == "" {
		log.Warn("Sentry client DSN is empty")
		return
	}

	h, err := newSentryHandler(value)
	if err != nil {
		log.Warn("Probably Sentry host is not running", "err", err)
		return
	}

	log.Root().SetHandler(
		log.MultiHandler(
			log.Root().GetHandler(),
			h,
		))
}

// SetLevel sets level filter on log root handler.
// So it should be called last.
func SetLevel(l string) {
	lvl, err := log.LvlFromString(l)
	if err != nil {
		panic(err)
	}

	log.Root().SetHandler(
		log.LvlFilterHandler(
			lvl,
			log.Root().GetHandler()))
}
