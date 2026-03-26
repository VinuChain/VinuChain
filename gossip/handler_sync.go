package gossip

import (
	"errors"
	"fmt"
	"time"

	"github.com/Fantom-foundation/lachesis-base/gossip/dagprocessor"
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/dag"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/utils/datasemaphore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/Fantom-foundation/go-opera/eventcheck"
	"github.com/Fantom-foundation/go-opera/eventcheck/bvallcheck"
	"github.com/Fantom-foundation/go-opera/eventcheck/epochcheck"
	"github.com/Fantom-foundation/go-opera/eventcheck/evallcheck"
	"github.com/Fantom-foundation/go-opera/eventcheck/heavycheck"
	"github.com/Fantom-foundation/go-opera/eventcheck/parentlesscheck"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/blockrecords/brprocessor"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/blockrecords/brstream"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/blockrecords/brstream/brstreamseeder"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/blockvotes/bvprocessor"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/blockvotes/bvstream"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/blockvotes/bvstream/bvstreamleecher"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/blockvotes/bvstream/bvstreamseeder"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/dag/dagstream"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/dag/dagstream/dagstreamseeder"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/epochpacks/epprocessor"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/epochpacks/epstream"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/epochpacks/epstream/epstreamseeder"
	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/inter/ibr"
	"github.com/Fantom-foundation/go-opera/inter/ier"
)

const (
	// maxPeerEpochDrift is the maximum number of epochs a peer may claim
	// ahead of our local epoch before we consider it invalid.
	maxPeerEpochDrift = idx.Epoch(1000)
	// maxPeerBlockDrift is the maximum number of blocks a peer may claim
	// ahead of our local head before we consider it invalid.
	maxPeerBlockDrift = idx.Block(5000)
)

// validatePeerProgress checks that a peer-reported progress is within
// plausible bounds relative to our local state. Returns a non-nil error
// (suitable for errResp) if the progress should be rejected.
func (h *handler) validatePeerProgress(progress PeerProgress) error {
	if progress.Epoch == 0 {
		return errResp(ErrDecode, "peer progress with zero epoch")
	}
	localEpoch := h.store.GetEpoch()
	if progress.Epoch > localEpoch+maxPeerEpochDrift {
		return errResp(ErrDecode, "peer epoch %d too far ahead of local %d", progress.Epoch, localEpoch)
	}
	localBlock := h.store.GetLatestBlockIndex()
	if progress.LastBlockIdx > localBlock+maxPeerBlockDrift {
		return errResp(ErrDecode, "peer block %d too far ahead of local %d", progress.LastBlockIdx, localBlock)
	}
	return nil
}

func (h *handler) highestPeerProgress() PeerProgress {
	peers := h.peers.List()
	max := h.myProgress()
	for _, peer := range peers {
		peerProgress := peer.GetProgress()
		if max.LastBlockIdx < peerProgress.LastBlockIdx {
			max = peerProgress
		}
	}
	return max
}

func (h *handler) makeDagProcessor(checkers *eventcheck.Checkers) *dagprocessor.Processor {
	// checkers
	lightCheck := func(e dag.Event) (retErr error) {
		defer func() {
			if r := recover(); r != nil {
				retErr = fmt.Errorf("type assertion failed in lightCheck: %v", r)
			}
		}()
		if h.store.GetEpoch() != e.ID().Epoch() {
			return epochcheck.ErrNotRelevant
		}
		if h.dagProcessor.IsBuffered(e.ID()) {
			return eventcheck.ErrDuplicateEvent
		}
		if h.store.HasEvent(e.ID()) {
			return eventcheck.ErrAlreadyConnectedEvent
		}
		if err := checkers.Basiccheck.Validate(e.(inter.EventPayloadI)); err != nil {
			return err
		}
		if err := checkers.Epochcheck.Validate(e.(inter.EventPayloadI)); err != nil {
			return err
		}
		return nil
	}
	bufferedCheck := func(_e dag.Event, _parents dag.Events) (retErr error) {
		defer func() {
			if r := recover(); r != nil {
				retErr = fmt.Errorf("type assertion failed in bufferedCheck: %v", r)
			}
		}()
		e := _e.(inter.EventI)
		parents := make(inter.EventIs, len(_parents))
		for i := range _parents {
			parents[i] = _parents[i].(inter.EventI)
		}
		var selfParent inter.EventI
		if e.SelfParent() != nil {
			selfParent = parents[0].(inter.EventI)
		}
		if err := checkers.Parentscheck.Validate(e, parents); err != nil {
			return err
		}
		if err := checkers.Gaspowercheck.Validate(e, selfParent); err != nil {
			return err
		}
		return nil
	}
	parentlessChecker := parentlesscheck.Checker{
		HeavyCheck: &heavycheck.EventsOnly{Checker: checkers.Heavycheck},
		LightCheck: lightCheck,
	}
	newProcessor := dagprocessor.New(datasemaphore.New(h.config.Protocol.EventsSemaphoreLimit, getSemaphoreWarningFn("DAG events")), h.config.Protocol.DagProcessor, dagprocessor.Callback{
		// DAG callbacks
		Event: dagprocessor.EventCallback{
			Process: func(_e dag.Event) error {
				e := _e.(*inter.EventPayload)
				preStart := time.Now()
				h.engineMu.Lock()
				defer h.engineMu.Unlock()

				err := h.process.Event(e)
				if err != nil {
					return err
				}

				// event is connected, announce it
				passedSinceEvent := preStart.Sub(e.CreationTime().Time())
				h.BroadcastEvent(e, passedSinceEvent)

				return nil
			},
			Released: func(e dag.Event, peer string, err error) {
				if eventcheck.IsBan(err) {
					log.Warn("Incoming event rejected", "event", e.ID().String(), "creator", e.Creator(), "err", err)
					h.removePeer(peer)
				}
			},

			Exists: func(id hash.Event) bool {
				return h.store.HasEvent(id)
			},

			Get: func(id hash.Event) dag.Event {
				e := h.store.GetEventPayload(id)
				if e == nil {
					return nil
				}
				return e
			},

			CheckParents:    bufferedCheck,
			CheckParentless: parentlessChecker.Enqueue,
		},
		HighestLamport: h.store.GetHighestLamport,
	})

	return newProcessor
}

func (h *handler) makeBvProcessor(checkers *eventcheck.Checkers) *bvprocessor.Processor {
	// checkers
	lightCheck := func(bvs inter.LlrSignedBlockVotes) error {
		if h.store.HasBlockVotes(bvs.Val.Epoch, bvs.Val.LastBlock(), bvs.Signed.Locator.ID()) {
			return eventcheck.ErrAlreadyProcessedBVs
		}
		return checkers.Basiccheck.ValidateBVs(bvs)
	}
	allChecker := bvallcheck.Checker{
		HeavyCheck: &heavycheck.BVsOnly{Checker: checkers.Heavycheck},
		LightCheck: lightCheck,
	}
	return bvprocessor.New(datasemaphore.New(h.config.Protocol.BVsSemaphoreLimit, getSemaphoreWarningFn("BVs")), h.config.Protocol.BvProcessor, bvprocessor.Callback{
		// DAG callbacks
		Item: bvprocessor.ItemCallback{
			Process: h.process.BVs,
			Released: func(bvs inter.LlrSignedBlockVotes, peer string, err error) {
				if eventcheck.IsBan(err) {
					log.Warn("Incoming BVs rejected", "BVs", bvs.Signed.Locator.ID(), "creator", bvs.Signed.Locator.Creator, "err", err)
					h.removePeer(peer)
				}
			},
			Check: allChecker.Enqueue,
		},
	})
}

func (h *handler) makeBrProcessor() *brprocessor.Processor {
	// checkers
	return brprocessor.New(datasemaphore.New(h.config.Protocol.BVsSemaphoreLimit, getSemaphoreWarningFn("BR")), h.config.Protocol.BrProcessor, brprocessor.Callback{
		// DAG callbacks
		Item: brprocessor.ItemCallback{
			Process: h.process.BR,
			Released: func(br ibr.LlrIdxFullBlockRecord, peer string, err error) {
				if eventcheck.IsBan(err) {
					log.Warn("Incoming BR rejected", "block", br.Idx, "err", err)
					h.removePeer(peer)
				}
			},
		},
	})
}

func (h *handler) makeEpProcessor(checkers *eventcheck.Checkers) *epprocessor.Processor {
	// checkers
	lightCheck := func(ev inter.LlrSignedEpochVote) error {
		if h.store.HasEpochVote(ev.Val.Epoch, ev.Signed.Locator.ID()) {
			return eventcheck.ErrAlreadyProcessedEV
		}
		return checkers.Basiccheck.ValidateEV(ev)
	}
	allChecker := evallcheck.Checker{
		HeavyCheck: &heavycheck.EVOnly{Checker: checkers.Heavycheck},
		LightCheck: lightCheck,
	}
	// checkers
	return epprocessor.New(datasemaphore.New(h.config.Protocol.BVsSemaphoreLimit, getSemaphoreWarningFn("BR")), h.config.Protocol.EpProcessor, epprocessor.Callback{
		// DAG callbacks
		Item: epprocessor.ItemCallback{
			ProcessEV: h.process.EV,
			ProcessER: h.process.ER,
			ReleasedEV: func(ev inter.LlrSignedEpochVote, peer string, err error) {
				if eventcheck.IsBan(err) {
					log.Warn("Incoming EV rejected", "event", ev.Signed.Locator.ID(), "creator", ev.Signed.Locator.Creator, "err", err)
					h.removePeer(peer)
				}
			},
			ReleasedER: func(er ier.LlrIdxFullEpochRecord, peer string, err error) {
				if eventcheck.IsBan(err) {
					log.Warn("Incoming ER rejected", "epoch", er.Idx, "err", err)
					h.removePeer(peer)
				}
			},
			CheckEV: allChecker.Enqueue,
		},
	})
}

func (h *handler) isEventInterested(id hash.Event, epoch idx.Epoch) bool {
	if id.Epoch() != epoch {
		return false
	}

	if h.dagProcessor.IsBuffered(id) || h.store.HasEvent(id) {
		return false
	}
	return true
}

func (h *handler) onlyInterestedEventsI(ids []interface{}) []interface{} {
	if len(ids) == 0 {
		return ids
	}
	epoch := h.store.GetEpoch()
	interested := make([]interface{}, 0, len(ids))
	for _, id := range ids {
		eid, ok := id.(hash.Event)
		if !ok {
			continue
		}
		if h.isEventInterested(eid, epoch) {
			interested = append(interested, id)
		}
	}
	return interested
}

func interfacesToEventIDs(ids []interface{}) hash.Events {
	res := make(hash.Events, 0, len(ids))
	for _, id := range ids {
		eid, ok := id.(hash.Event)
		if !ok {
			continue
		}
		res = append(res, eid)
	}
	return res
}

func eventIDsToInterfaces(ids hash.Events) []interface{} {
	res := make([]interface{}, len(ids))
	for i, id := range ids {
		res[i] = id
	}
	return res
}

func interfacesToTxids(ids []interface{}) []common.Hash {
	res := make([]common.Hash, 0, len(ids))
	for _, id := range ids {
		txid, ok := id.(common.Hash)
		if !ok {
			continue
		}
		res = append(res, txid)
	}
	return res
}

func txidsToInterfaces(ids []common.Hash) []interface{} {
	res := make([]interface{}, len(ids))
	for i, id := range ids {
		res[i] = id
	}
	return res
}

func (h *handler) handleTxHashes(p *peer, announces []common.Hash) {
	// Mark the hashes as present at the remote node
	for _, id := range announces {
		p.MarkTransaction(id)
	}
	// Schedule all the unknown hashes for retrieval
	requestTransactions := func(ids []interface{}) error {
		return p.RequestTransactions(interfacesToTxids(ids))
	}
	_ = h.txFetcher.NotifyAnnounces(p.id, txidsToInterfaces(announces), time.Now(), requestTransactions)
}

func (h *handler) handleTxs(p *peer, txs types.Transactions) {
	// Mark the hashes as present at the remote node
	for _, tx := range txs {
		p.MarkTransaction(tx.Hash())
	}
	h.txpool.AddRemotes(txs)
}

func (h *handler) handleEventHashes(p *peer, announces hash.Events) {
	// Mark the hashes as present at the remote node
	for _, id := range announces {
		p.MarkEvent(id)
	}
	// filter too high IDs
	notTooHigh := make(hash.Events, 0, len(announces))
	sessionCfg := h.config.Protocol.DagStreamLeecher.Session
	for _, id := range announces {
		maxLamport := h.store.GetHighestLamport() + idx.Lamport(sessionCfg.DefaultChunkItemsNum+1)*idx.Lamport(sessionCfg.ParallelChunksDownload)
		if id.Lamport() <= maxLamport {
			notTooHigh = append(notTooHigh, id)
		}
	}
	if len(announces) != len(notTooHigh) {
		h.dagLeecher.ForceSyncing()
	}
	if len(notTooHigh) == 0 {
		return
	}
	// Schedule all the unknown hashes for retrieval
	requestEvents := func(ids []interface{}) error {
		return p.RequestEvents(interfacesToEventIDs(ids))
	}
	_ = h.dagFetcher.NotifyAnnounces(p.id, eventIDsToInterfaces(notTooHigh), time.Now(), requestEvents)
}

func (h *handler) handleEvents(p *peer, events dag.Events, ordered bool) {
	// Mark the hashes as present at the remote node
	for _, e := range events {
		p.MarkEvent(e.ID())
	}
	// filter too high events
	notTooHigh := make(dag.Events, 0, len(events))
	sessionCfg := h.config.Protocol.DagStreamLeecher.Session
	now := time.Now()
	for _, e := range events {
		maxLamport := h.store.GetHighestLamport() + idx.Lamport(sessionCfg.DefaultChunkItemsNum+1)*idx.Lamport(sessionCfg.ParallelChunksDownload)
		if e.Lamport() <= maxLamport {
			notTooHigh = append(notTooHigh, e)
		}
		if ei, ok := e.(inter.EventI); ok {
			if now.Sub(ei.CreationTime().Time()) < 10*time.Minute {
				h.syncStatus.MarkMaybeSynced()
			}
		}
	}
	if len(events) != len(notTooHigh) {
		h.dagLeecher.ForceSyncing()
	}
	if len(notTooHigh) == 0 {
		return
	}
	// Schedule all the events for connection.
	// We capture p (pointer) rather than copying *p because peer contains a sync.RWMutex.
	// If the peer disconnects before the callback fires, p2p.Send returns an error gracefully.
	peerID := p.id
	requestEvents := func(ids []interface{}) error {
		return p.RequestEvents(interfacesToEventIDs(ids))
	}
	notifyAnnounces := func(ids hash.Events) {
		_ = h.dagFetcher.NotifyAnnounces(peerID, eventIDsToInterfaces(ids), now, requestEvents)
	}
	_ = h.dagProcessor.Enqueue(peerID, notTooHigh, ordered, notifyAnnounces, nil)
}

// handleMsg is invoked whenever an inbound message is received from a remote
// peer. The remote connection is torn down upon returning any error.
func (h *handler) handleMsg(p *peer) error {
	// Read the next message from the remote peer, and ensure it's fully consumed
	msg, err := p.rw.ReadMsg()
	if err != nil {
		return err
	}
	if msg.Size > protocolMaxMsgSize {
		return errResp(ErrMsgTooLarge, "%v > %v", msg.Size, protocolMaxMsgSize)
	}
	defer msg.Discard()
	// Acquire semaphore for serialized messages
	eventsSizeEst := dag.Metric{
		Num:  1,
		Size: uint64(msg.Size),
	}
	if !h.msgSemaphore.Acquire(eventsSizeEst, h.config.Protocol.MsgsSemaphoreTimeout) {
		h.Log.Warn("Failed to acquire semaphore for p2p message", "size", msg.Size, "peer", p.id)
		return nil
	}
	defer h.msgSemaphore.Release(eventsSizeEst)

	// Handle the message depending on its contents
	switch {
	case msg.Code == HandshakeMsg:
		// Status messages should never arrive after the handshake
		return errResp(ErrExtraStatusMsg, "uncontrolled status message")

	case msg.Code == ProgressMsg:
		var progress PeerProgress
		if err := msg.Decode(&progress); err != nil {
			return errResp(ErrDecode, "%v: %v", msg, err)
		}
		if err := h.validatePeerProgress(progress); err != nil {
			return err
		}
		p.SetProgress(progress)

	case msg.Code == EvmTxsMsg:
		// Transactions arrived, make sure we have a valid and fresh graph to handle them
		if !h.syncStatus.AcceptTxs() {
			break
		}
		// Transactions can be processed, parse all of them and deliver to the pool
		var txs types.Transactions
		if err := msg.Decode(&txs); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		if err := checkLenLimits(len(txs), txs); err != nil {
			return err
		}
		txids := make([]interface{}, txs.Len())
		for i, tx := range txs {
			txids[i] = tx.Hash()
		}
		_ = h.txFetcher.NotifyReceived(txids)
		h.handleTxs(p, txs)

	case msg.Code == NewEvmTxHashesMsg:
		// Transactions arrived, make sure we have a valid and fresh graph to handle them
		if !h.syncStatus.AcceptTxs() {
			break
		}
		// Transactions can be processed, parse all of them and deliver to the pool
		var txHashes []common.Hash
		if err := msg.Decode(&txHashes); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		if err := checkLenLimits(len(txHashes), txHashes); err != nil {
			return err
		}
		h.handleTxHashes(p, txHashes)

	case msg.Code == GetEvmTxsMsg:
		var requests []common.Hash
		if err := msg.Decode(&requests); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		if err := checkLenLimits(len(requests), requests); err != nil {
			return err
		}

		txs := make(types.Transactions, 0, len(requests))
		for _, txid := range requests {
			tx := h.txpool.Get(txid)
			if tx == nil {
				continue
			}
			txs = append(txs, tx)
		}
		SplitTransactions(txs, func(batch types.Transactions) {
			p.EnqueueSendTransactions(batch, p.queue)
		})

	case msg.Code == EventsMsg:
		if !h.syncStatus.AcceptEvents() {
			break
		}

		var events inter.EventPayloads
		if err := msg.Decode(&events); err != nil {
			return errResp(ErrDecode, "%v: %v", msg, err)
		}
		if err := checkLenLimits(len(events), events); err != nil {
			return err
		}
		_ = h.dagFetcher.NotifyReceived(eventIDsToInterfaces(events.IDs()))
		h.handleEvents(p, events.Bases(), events.Len() > 1)

	case msg.Code == NewEventIDsMsg:
		// Fresh events arrived, make sure we have a valid and fresh graph to handle them
		if !h.syncStatus.AcceptEvents() {
			break
		}
		var announces hash.Events
		if err := msg.Decode(&announces); err != nil {
			return errResp(ErrDecode, "%v: %v", msg, err)
		}
		if err := checkLenLimits(len(announces), announces); err != nil {
			return err
		}
		h.handleEventHashes(p, announces)

	case msg.Code == GetEventsMsg:
		var requests hash.Events
		if err := msg.Decode(&requests); err != nil {
			return errResp(ErrDecode, "%v: %v", msg, err)
		}
		if err := checkLenLimits(len(requests), requests); err != nil {
			return err
		}

		rawEvents := make([]rlp.RawValue, 0, len(requests))
		ids := make(hash.Events, 0, len(requests))
		size := 0
		for _, id := range requests {
			if raw := h.store.GetEventPayloadRLP(id); raw != nil {
				rawEvents = append(rawEvents, raw)
				ids = append(ids, id)
				size += len(raw)
			} else {
				h.Log.Debug("requested event not found", "hash", id)
			}
			if size >= softResponseLimitSize {
				break
			}
		}
		if len(rawEvents) != 0 {
			p.EnqueueSendEventsRLP(rawEvents, ids, p.queue)
		}

	case msg.Code == RequestEventsStream:
		var request dagstream.Request
		if err := msg.Decode(&request); err != nil {
			return errResp(ErrDecode, "%v: %v", msg, err)
		}
		if request.Limit.Num > hardLimitItems-1 {
			return errResp(ErrMsgTooLarge, "%v", msg)
		}
		if request.Limit.Size > protocolMaxMsgSize*2/3 {
			return errResp(ErrMsgTooLarge, "%v", msg)
		}

		pid := p.id
		_, peerErr := h.dagSeeder.NotifyRequestReceived(dagstreamseeder.Peer{
			ID:        pid,
			SendChunk: p.SendEventsStream,
			Misbehaviour: func(err error) {
				h.peerMisbehaviour(pid, err)
			},
		}, request)
		if peerErr != nil {
			return peerErr
		}

	case msg.Code == EventsStreamResponse:
		if !h.syncStatus.AcceptEvents() {
			break
		}

		var chunk dagChunk
		if err := msg.Decode(&chunk); err != nil {
			return errResp(ErrDecode, "%v: %v", msg, err)
		}
		if err := checkLenLimits(len(chunk.Events)+len(chunk.IDs)+1, chunk); err != nil {
			return err
		}

		if (len(chunk.Events) != 0) && (len(chunk.IDs) != 0) {
			return errors.New("expected either events or event hashes")
		}
		var last hash.Event
		if len(chunk.IDs) != 0 {
			h.handleEventHashes(p, chunk.IDs)
			last = chunk.IDs[len(chunk.IDs)-1]
		}
		if len(chunk.Events) != 0 {
			h.handleEvents(p, chunk.Events.Bases(), true)
			last = chunk.Events[len(chunk.Events)-1].ID()
		}

		_ = h.dagLeecher.NotifyChunkReceived(chunk.SessionID, last, chunk.Done)

	case msg.Code == RequestBVsStream:
		var request bvstream.Request
		if err := msg.Decode(&request); err != nil {
			return errResp(ErrDecode, "%v: %v", msg, err)
		}
		if request.Limit.Num > hardLimitItems-1 {
			return errResp(ErrMsgTooLarge, "%v", msg)
		}
		if request.Limit.Size > protocolMaxMsgSize*2/3 {
			return errResp(ErrMsgTooLarge, "%v", msg)
		}

		pid := p.id
		_, peerErr := h.bvSeeder.NotifyRequestReceived(bvstreamseeder.Peer{
			ID:        pid,
			SendChunk: p.SendBVsStream,
			Misbehaviour: func(err error) {
				h.peerMisbehaviour(pid, err)
			},
		}, request)
		if peerErr != nil {
			return peerErr
		}

	case msg.Code == BVsStreamResponse:
		var chunk bvsChunk
		if err := msg.Decode(&chunk); err != nil {
			return errResp(ErrDecode, "%v: %v", msg, err)
		}
		if err := checkLenLimits(len(chunk.BVs)+1, chunk); err != nil {
			return err
		}

		var last bvstreamleecher.BVsID
		if len(chunk.BVs) != 0 {
			_ = h.bvProcessor.Enqueue(p.id, chunk.BVs, nil)
			last = bvstreamleecher.BVsID{
				Epoch:     chunk.BVs[len(chunk.BVs)-1].Val.Epoch,
				LastBlock: chunk.BVs[len(chunk.BVs)-1].Val.LastBlock(),
				ID:        chunk.BVs[len(chunk.BVs)-1].Signed.Locator.ID(),
			}
		}

		_ = h.bvLeecher.NotifyChunkReceived(chunk.SessionID, last, chunk.Done)

	case msg.Code == RequestBRsStream:
		var request brstream.Request
		if err := msg.Decode(&request); err != nil {
			return errResp(ErrDecode, "%v: %v", msg, err)
		}
		if request.Limit.Num > hardLimitItems-1 {
			return errResp(ErrMsgTooLarge, "%v", msg)
		}
		if request.Limit.Size > protocolMaxMsgSize*2/3 {
			return errResp(ErrMsgTooLarge, "%v", msg)
		}

		pid := p.id
		_, peerErr := h.brSeeder.NotifyRequestReceived(brstreamseeder.Peer{
			ID:        pid,
			SendChunk: p.SendBRsStream,
			Misbehaviour: func(err error) {
				h.peerMisbehaviour(pid, err)
			},
		}, request)
		if peerErr != nil {
			return peerErr
		}

	case msg.Code == BRsStreamResponse:
		if !h.syncStatus.AcceptBlockRecords() {
			break
		}

		msgSize := uint64(msg.Size)
		var chunk brsChunk
		if err := msg.Decode(&chunk); err != nil {
			return errResp(ErrDecode, "%v: %v", msg, err)
		}
		if err := checkLenLimits(len(chunk.BRs)+1, chunk); err != nil {
			return err
		}

		var last idx.Block
		if len(chunk.BRs) != 0 {
			_ = h.brProcessor.Enqueue(p.id, chunk.BRs, msgSize, nil)
			last = chunk.BRs[len(chunk.BRs)-1].Idx
		}

		_ = h.brLeecher.NotifyChunkReceived(chunk.SessionID, last, chunk.Done)

	case msg.Code == RequestEPsStream:
		var request epstream.Request
		if err := msg.Decode(&request); err != nil {
			return errResp(ErrDecode, "%v: %v", msg, err)
		}
		if request.Limit.Num > hardLimitItems-1 {
			return errResp(ErrMsgTooLarge, "%v", msg)
		}
		if request.Limit.Size > protocolMaxMsgSize*2/3 {
			return errResp(ErrMsgTooLarge, "%v", msg)
		}

		pid := p.id
		_, peerErr := h.epSeeder.NotifyRequestReceived(epstreamseeder.Peer{
			ID:        pid,
			SendChunk: p.SendEPsStream,
			Misbehaviour: func(err error) {
				h.peerMisbehaviour(pid, err)
			},
		}, request)
		if peerErr != nil {
			return peerErr
		}

	case msg.Code == EPsStreamResponse:
		msgSize := uint64(msg.Size)
		var chunk epsChunk
		if err := msg.Decode(&chunk); err != nil {
			return errResp(ErrDecode, "%v: %v", msg, err)
		}
		if err := checkLenLimits(len(chunk.EPs)+1, chunk); err != nil {
			return err
		}

		var last idx.Epoch
		if len(chunk.EPs) != 0 {
			_ = h.epProcessor.Enqueue(p.id, chunk.EPs, msgSize, nil)
			last = chunk.EPs[len(chunk.EPs)-1].Record.Idx
		}

		_ = h.epLeecher.NotifyChunkReceived(chunk.SessionID, last, chunk.Done)

	default:
		return errResp(ErrInvalidMsgCode, "%v", msg.Code)
	}
	return nil
}
