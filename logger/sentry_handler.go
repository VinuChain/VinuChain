package logger

import (
	"fmt"

	"github.com/ethereum/go-ethereum/log"
	raven "github.com/getsentry/raven-go"
)

// newSentryHandler returns a log.Handler that forwards LvlError and LvlCrit
// records to the Sentry project identified by dsn. Records below LvlError are
// dropped. Uses raven-go (Sentry SDK v1) which ships fire-and-forget delivery.
func newSentryHandler(dsn string) (log.Handler, error) {
	client, err := raven.NewClient(dsn, nil)
	if err != nil {
		return nil, err
	}
	return log.FuncHandler(func(r *log.Record) error {
		if r.Lvl > log.LvlError {
			return nil
		}
		packet := raven.NewPacketWithExtra(r.Msg, ctxToExtra(r.Ctx))
		packet.Level = ravenLevel(r.Lvl)
		client.Capture(packet, nil)
		return nil
	}), nil
}

func ctxToExtra(ctx []interface{}) raven.Extra {
	extra := make(raven.Extra, len(ctx)/2)
	for i := 1; i < len(ctx); i += 2 {
		k, ok := ctx[i-1].(string)
		if !ok {
			k = fmt.Sprintf("key%d", i-1)
		}
		extra[k] = ctx[i]
	}
	return extra
}

func ravenLevel(lvl log.Lvl) raven.Severity {
	switch lvl {
	case log.LvlCrit:
		return raven.FATAL
	case log.LvlError:
		return raven.ERROR
	default:
		return raven.WARNING
	}
}
