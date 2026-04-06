package launcher

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/kvdb"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"gopkg.in/urfave/cli.v1"

	"github.com/Fantom-foundation/go-opera/gossip"
	"github.com/Fantom-foundation/go-opera/inter"
)

// parseBlockNumber parses a block number from a decimal string and enforces the
// idx.Block range (uint32). strconv.ParseUint with bitSize=32 rejects values
// larger than math.MaxUint32, preventing silent truncation on cast to idx.Block.
func parseBlockNumber(s string) (uint32, error) {
	n, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(n), nil
}

// TracePayload is the RLP-encoded form for import/export of a single trace entry.
type TracePayload struct {
	Key    common.Hash
	Traces []byte
}

// importTxTraces imports transaction traces from a file into the trace store.
func importTxTraces(ctx *cli.Context) error {
	if len(ctx.Args()) < 1 {
		utils.Fatalf("This command requires an argument.")
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	cfg := makeAllConfigs(ctx)

	rawDbs := makeDirectDBsProducer(cfg)
	gdb, err := makeGossipStoreTrace(rawDbs, cfg)
	if err != nil {
		log.Crit("DB opening error", "datadir", cfg.Node.DataDir, "err", err)
	}
	defer gdb.Close()

	fn := ctx.Args().First()

	fh, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer fh.Close()

	var (
		reader  io.Reader = fh
		counter int
	)
	if strings.HasSuffix(fn, ".gz") {
		if reader, err = gzip.NewReader(reader); err != nil {
			return err
		}
		defer reader.(*gzip.Reader).Close()
	}

	log.Info("Importing transaction traces from file", "file", fn)
	start, reported := time.Now(), time.Now()

	stream := rlp.NewStream(reader, maxImportStreamSize)
	for {
		select {
		case <-interrupt:
			return fmt.Errorf("interrupted")
		default:
		}
		e := new(TracePayload)
		err = stream.Decode(e)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		gdb.TxTraceStore().SetTxTrace(e.Key, e.Traces)
		counter++
		if time.Since(reported) >= statsReportLimit {
			log.Info("Importing transaction traces", "imported", counter, "elapsed", common.PrettyDuration(time.Since(start)))
			reported = time.Now()
		}
	}
	log.Info("Imported transaction traces", "imported", counter, "elapsed", common.PrettyDuration(time.Since(start)))
	return nil
}

// deleteTxTraces removes transaction traces for the specified block range.
func deleteTxTraces(ctx *cli.Context) error {
	cfg := makeAllConfigs(ctx)

	rawDbs := makeDirectDBsProducer(cfg)
	gdb, err := makeGossipStoreTrace(rawDbs, cfg)
	if err != nil {
		log.Crit("DB opening error", "datadir", cfg.Node.DataDir, "err", err)
	}
	defer gdb.Close()

	from := idx.Block(1)
	if len(ctx.Args()) > 0 {
		n, err := parseBlockNumber(ctx.Args().Get(0))
		if err != nil {
			return fmt.Errorf("invalid from block: %w", err)
		}
		from = idx.Block(n)
	}
	to := gdb.GetLatestBlockIndex()
	if len(ctx.Args()) > 1 {
		n, err := parseBlockNumber(ctx.Args().Get(1))
		if err != nil {
			return fmt.Errorf("invalid to block: %w", err)
		}
		to = idx.Block(n)
	}

	log.Info("Deleting transaction traces", "from block", from, "to block", to)

	if err := deleteTraces(gdb, from, to); err != nil {
		utils.Fatalf("Deleting traces error: %v\n", err)
	}
	return nil
}

// exportTxTraces exports transaction traces for the specified block range to a file.
func exportTxTraces(ctx *cli.Context) error {
	if len(ctx.Args()) < 1 {
		utils.Fatalf("This command requires an argument.")
	}

	cfg := makeAllConfigs(ctx)

	rawDbs := makeDirectDBsProducer(cfg)
	gdb, err := makeGossipStoreTrace(rawDbs, cfg)
	if err != nil {
		log.Crit("DB opening error", "datadir", cfg.Node.DataDir, "err", err)
	}
	defer gdb.Close()

	fn := ctx.Args().First()

	fh, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer fh.Close()

	var writer io.Writer = fh
	if strings.HasSuffix(fn, ".gz") {
		writer = gzip.NewWriter(writer)
		defer writer.(*gzip.Writer).Close()
	}

	from := idx.Block(1)
	if len(ctx.Args()) > 1 {
		n, err := parseBlockNumber(ctx.Args().Get(1))
		if err != nil {
			return fmt.Errorf("invalid from block: %w", err)
		}
		from = idx.Block(n)
	}
	to := gdb.GetLatestBlockIndex()
	if len(ctx.Args()) > 2 {
		n, err := parseBlockNumber(ctx.Args().Get(2))
		if err != nil {
			return fmt.Errorf("invalid to block: %w", err)
		}
		to = idx.Block(n)
	}

	log.Info("Exporting transaction traces to file", "file", fn)

	if err := exportTraceTo(writer, gdb, from, to); err != nil {
		utils.Fatalf("Export error: %v\n", err)
	}
	return nil
}

// makeGossipStoreTrace opens a gossip store with trace transactions enabled.
func makeGossipStoreTrace(producer kvdb.FlushableDBProducer, cfg *config) (*gossip.Store, error) {
	cfg.OperaStore.TraceTransactions = true
	gdb := gossip.NewStore(producer, cfg.OperaStore)
	if gdb.TxTraceStore() == nil {
		return nil, errors.New("transaction traces db store is not initialized")
	}
	return gdb, nil
}

func exportTraceTo(w io.Writer, gdb *gossip.Store, from, to idx.Block) error {
	if from == 1 && to == gdb.GetLatestBlockIndex() {
		return exportAllTraceTo(w, gdb)
	}
	start, reported := time.Now(), time.Now()
	var (
		counter int
		block   *inter.Block
	)
	for i := from; i <= to; i++ {
		block = gdb.GetBlock(i)
		if block == nil {
			continue
		}
		for _, tx := range gdb.GetBlockTxs(i, block) {
			traces, traceErr := gdb.TxTraceStore().GetTx(tx.Hash())
			if traceErr != nil {
				log.Error("Failed to read tx trace", "tx", tx.Hash(), "err", traceErr)
				continue
			}
			if len(traces) > 0 {
				counter++
				if err := rlp.Encode(w, TracePayload{tx.Hash(), traces}); err != nil {
					return fmt.Errorf("failed to encode trace for tx %s: %w", tx.Hash(), err)
				}
			}
			if time.Since(reported) >= statsReportLimit {
				log.Info("Exporting transaction traces", "at block", i, "exported", counter, "elapsed", common.PrettyDuration(time.Since(start)))
				reported = time.Now()
			}
		}
	}
	log.Info("Exported transaction traces", "from block", from, "to block", to, "exported", counter, "elapsed", common.PrettyDuration(time.Since(start)))
	return nil
}

func exportAllTraceTo(w io.Writer, gdb *gossip.Store) error {
	start, reported := time.Now(), time.Now()
	var counter int
	var exportErr error

	gdb.TxTraceStore().ForEachTxtrace(func(key common.Hash, traces []byte) bool {
		counter++
		exportErr = rlp.Encode(w, TracePayload{key, traces})
		if exportErr != nil {
			return false
		}
		if time.Since(reported) >= statsReportLimit {
			log.Info("Exporting all transaction traces", "exported", counter, "elapsed", common.PrettyDuration(time.Since(start)))
			reported = time.Now()
		}
		return true
	})
	log.Info("Exported all transaction traces", "exported", counter, "elapsed", common.PrettyDuration(time.Since(start)))
	return exportErr
}

func deleteTraces(gdb *gossip.Store, from, to idx.Block) error {
	start, reported := time.Now(), time.Now()
	var counter int

	for i := from; i <= to; i++ {
		blk := gdb.GetBlock(i)
		if blk == nil {
			continue
		}
		for _, tx := range gdb.GetBlockTxs(i, blk) {
			if existing, _ := gdb.TxTraceStore().GetTx(tx.Hash()); existing != nil {
				counter++
				gdb.TxTraceStore().RemoveTxTrace(tx.Hash())
				if time.Since(reported) >= statsReportLimit {
					log.Info("Deleting traces", "deleted", counter, "elapsed", common.PrettyDuration(time.Since(start)))
					reported = time.Now()
				}
			}
		}
	}
	log.Info("Deleting transaction traces done", "deleted", counter, "from block", from, "to block", to, "elapsed", common.PrettyDuration(time.Since(start)))
	return nil
}
