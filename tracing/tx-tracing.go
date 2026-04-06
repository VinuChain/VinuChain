package tracing

import (
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/opentracing/opentracing-go"
)

var (
	enabled   atomic.Bool
	txSpans   = make(map[common.Hash]opentracing.Span)
	txSpansMu sync.RWMutex

	noopSpan = opentracing.NoopTracer{}.StartSpan("")
)

func SetEnabled(val bool) {
	enabled.Store(val)
}

func Enabled() bool {
	return enabled.Load()
}

func StartTx(tx common.Hash, operation string) {
	if !enabled.Load() {
		return
	}

	txSpansMu.Lock()
	defer txSpansMu.Unlock()

	if _, ok := txSpans[tx]; ok {
		return
	}

	span := opentracing.StartSpan("lifecycle")
	span.SetTag("txhash", tx.String())
	span.SetTag("enter", operation)
	txSpans[tx] = span
}

func FinishTx(tx common.Hash, operation string) {
	if !enabled.Load() {
		return
	}

	txSpansMu.Lock()
	defer txSpansMu.Unlock()

	span, ok := txSpans[tx]
	if !ok {
		return
	}

	span.SetTag("exit", operation)
	span.Finish()
	delete(txSpans, tx)
}

func CheckTx(tx common.Hash, operation string) opentracing.Span {
	if !enabled.Load() {
		return noopSpan
	}

	txSpansMu.RLock()
	defer txSpansMu.RUnlock()

	span, ok := txSpans[tx]

	if !ok {
		return noopSpan
	}

	return opentracing.GlobalTracer().StartSpan(
		operation,
		opentracing.ChildOf(span.Context()),
	)
}
