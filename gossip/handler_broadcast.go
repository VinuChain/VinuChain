package gossip

import (
	"math"
	"time"

	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Fantom-foundation/go-opera/inter"
)

func (h *handler) decideBroadcastAggressiveness(size int, passed time.Duration, peersNum int) int {
	percents := 100
	maxPercents := 1000000 * percents
	latencyVsThroughputTradeoff := maxPercents
	cfg := h.config.Protocol
	if cfg.ThroughputImportance != 0 {
		latencyVsThroughputTradeoff = (cfg.LatencyImportance * percents) / cfg.ThroughputImportance
	}

	broadcastCost := passed * time.Duration(128+size) / 128
	broadcastAllCostTarget := time.Duration(latencyVsThroughputTradeoff) * (700 * time.Millisecond) / time.Duration(percents)
	broadcastSqrtCostTarget := broadcastAllCostTarget * 10

	fullRecipients := 0
	if latencyVsThroughputTradeoff >= maxPercents {
		// edge case
		fullRecipients = peersNum
	} else if latencyVsThroughputTradeoff <= 0 {
		// edge case
		fullRecipients = 0
	} else if broadcastCost <= broadcastAllCostTarget {
		// if event is small or was created recently, always send to everyone full event
		fullRecipients = peersNum
	} else if broadcastCost <= broadcastSqrtCostTarget || passed == 0 {
		// if event is big but was created recently, send full event to subset of peers
		fullRecipients = int(math.Sqrt(float64(peersNum)))
		if fullRecipients < 4 {
			fullRecipients = 4
		}
	}
	if fullRecipients > peersNum {
		fullRecipients = peersNum
	}
	return fullRecipients
}

// BroadcastEvent will either propagate a event to a subset of it's peers, or
// will only announce it's availability (depending what's requested).
func (h *handler) BroadcastEvent(event *inter.EventPayload, passed time.Duration) int {
	if passed < 0 {
		passed = 0
	}
	id := event.ID()
	peers := h.peers.PeersWithoutEvent(id)
	if len(peers) == 0 {
		log.Trace("Event is already known to all peers", "hash", id)
		return 0
	}

	fullRecipients := h.decideBroadcastAggressiveness(event.Size(), passed, len(peers))

	// Exclude low quality peers from fullBroadcast
	var fullBroadcast = make([]*peer, 0, fullRecipients)
	var hashBroadcast = make([]*peer, 0, len(peers))
	for _, p := range peers {
		if !p.Useless() && len(fullBroadcast) < fullRecipients {
			fullBroadcast = append(fullBroadcast, p)
		} else {
			hashBroadcast = append(hashBroadcast, p)
		}
	}
	for _, peer := range fullBroadcast {
		peer.AsyncSendEvents(inter.EventPayloads{event}, peer.queue)
	}
	// Broadcast of event hash to the rest peers
	for _, peer := range hashBroadcast {
		peer.AsyncSendEventIDs(hash.Events{event.ID()}, peer.queue)
	}
	log.Trace("Broadcast event", "hash", id, "fullRecipients", len(fullBroadcast), "hashRecipients", len(hashBroadcast))
	return len(peers)
}

// BroadcastTxs will propagate a batch of transactions to all peers which are not known to
// already have the given transaction.
func (h *handler) BroadcastTxs(txs types.Transactions) {
	totalSize := common.StorageSize(0)
	for _, tx := range txs {
		totalSize += tx.Size()
	}

	peers := h.peers.List()
	fullRecipients := h.decideBroadcastAggressiveness(int(totalSize), time.Second, len(peers))
	fullSent := 0
	for _, p := range peers {
		var peerTxs types.Transactions
		for _, tx := range txs {
			if !p.knownTxs.Contains(tx.Hash()) {
				peerTxs = append(peerTxs, tx)
			}
		}
		if len(peerTxs) == 0 {
			continue
		}
		sendFull := fullSent < fullRecipients
		SplitTransactions(peerTxs, func(batch types.Transactions) {
			if sendFull {
				p.AsyncSendTransactions(batch, p.queue)
			} else {
				txids := make([]common.Hash, batch.Len())
				for j, tx := range batch {
					txids[j] = tx.Hash()
				}
				p.AsyncSendTransactionHashes(txids, p.queue)
			}
		})
		if sendFull {
			fullSent++
		}
	}
}

// Mined broadcast loop
func (h *handler) emittedBroadcastLoop() {
	defer h.loopsWg.Done()
	for {
		select {
		case emitted := <-h.emittedEventsCh:
			h.BroadcastEvent(emitted, 0)
		// Err() channel will be closed when unsubscribing.
		case <-h.emittedEventsSub.Err():
			return
		}
	}
}

func (h *handler) broadcastProgress() {
	progress := h.myProgress()
	for _, peer := range h.peers.List() {
		peer.AsyncSendProgress(progress, peer.queue)
	}
}

// Progress broadcast loop
func (h *handler) progressBroadcastLoop() {
	ticker := time.NewTicker(h.config.Protocol.ProgressBroadcastPeriod)
	defer ticker.Stop()
	defer h.loopsWg.Done()
	// automatically stops if unsubscribe
	for {
		select {
		case <-ticker.C:
			h.broadcastProgress()
		case <-h.quitProgressBroadcast:
			return
		}
	}
}

func (h *handler) onNewEpochLoop() {
	defer h.loopsWg.Done()
	for {
		select {
		case myEpoch := <-h.newEpochsCh:
			h.dagProcessor.Clear()
			h.dagLeecher.OnNewEpoch(myEpoch)
		// Err() channel will be closed when unsubscribing.
		case <-h.newEpochsSub.Err():
			return
		}
	}
}

func (h *handler) txBroadcastLoop() {
	ticker := time.NewTicker(h.config.Protocol.RandomTxHashesSendPeriod)
	defer ticker.Stop()
	defer h.loopsWg.Done()
	for {
		select {
		case notify := <-h.txsCh:
			h.BroadcastTxs(notify.Txs)

		// Err() channel will be closed when unsubscribing.
		case <-h.txsSub.Err():
			return

		case <-ticker.C:
			if !h.syncStatus.AcceptTxs() {
				break
			}
			peers := h.peers.List()
			if len(peers) == 0 {
				continue
			}
			randPeer := peers[cryptoRandIntn(len(peers))]
			h.syncTransactions(randPeer, h.txpool.SampleHashes(h.config.Protocol.MaxRandomTxHashesSend))
		}
	}
}
