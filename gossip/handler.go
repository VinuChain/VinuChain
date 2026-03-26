package gossip

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Fantom-foundation/lachesis-base/gossip/dagprocessor"
	"github.com/Fantom-foundation/lachesis-base/gossip/itemsfetcher"
	"github.com/Fantom-foundation/lachesis-base/inter/dag"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/utils/datasemaphore"
	"github.com/ethereum/go-ethereum/common"
	notify "github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover/discfilter"
	"github.com/ethereum/go-ethereum/trie"

	"github.com/Fantom-foundation/go-opera/eventcheck"
	"github.com/Fantom-foundation/go-opera/evmcore"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/blockrecords/brprocessor"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/blockrecords/brstream"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/blockrecords/brstream/brstreamleecher"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/blockrecords/brstream/brstreamseeder"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/blockvotes/bvprocessor"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/blockvotes/bvstream"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/blockvotes/bvstream/bvstreamleecher"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/blockvotes/bvstream/bvstreamseeder"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/dag/dagstream"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/dag/dagstream/dagstreamleecher"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/dag/dagstream/dagstreamseeder"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/epochpacks/epprocessor"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/epochpacks/epstream"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/epochpacks/epstream/epstreamleecher"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/epochpacks/epstream/epstreamseeder"
	"github.com/Fantom-foundation/go-opera/gossip/protocols/snap/snapstream/snapleecher"
	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/inter/ibr"
	"github.com/Fantom-foundation/go-opera/inter/ier"
	"github.com/Fantom-foundation/go-opera/logger"
)

const (
	softResponseLimitSize = 2 * 1024 * 1024    // Target maximum size of returned events, or other data.
	softLimitItems        = 250                // Target maximum number of events or transactions to request/response
	hardLimitItems        = softLimitItems * 4 // Maximum number of events or transactions to request/response

	// txChanSize is the size of channel listening to NewTxsNotify.
	// The number is referenced from the size of tx pool.
	txChanSize = 4096
)

func errResp(code errCode, format string, v ...interface{}) error {
	return fmt.Errorf("%v - %v", code, fmt.Sprintf(format, v...))
}

func checkLenLimits(size int, v interface{}) error {
	if size <= 0 {
		return errResp(ErrEmptyMessage, "%v", v)
	}
	if size > hardLimitItems {
		return errResp(ErrMsgTooLarge, "%v", v)
	}
	return nil
}

type dagNotifier interface {
	SubscribeNewEpoch(ch chan<- idx.Epoch) notify.Subscription
	SubscribeNewEmitted(ch chan<- *inter.EventPayload) notify.Subscription
}

type processCallback struct {
	Event            func(*inter.EventPayload) error
	SwitchEpochTo    func(idx.Epoch) error
	PauseEvmSnapshot func()
	BVs              func(inter.LlrSignedBlockVotes) error
	BR               func(ibr.LlrIdxFullBlockRecord) error
	EV               func(inter.LlrSignedEpochVote) error
	ER               func(ier.LlrIdxFullEpochRecord) error
}

// handlerConfig is the collection of initialization parameters to create a full
// node network handler.
type handlerConfig struct {
	config   Config
	notifier dagNotifier
	txpool   TxPool
	engineMu sync.Locker
	checkers *eventcheck.Checkers
	s        *Store
	process  processCallback
}

type snapsyncEpochUpd struct {
	epoch idx.Epoch
	root  common.Hash
}

type snapsyncCancelCmd struct {
	done chan struct{}
}

type snapsyncStateUpd struct {
	snapsyncEpochUpd  *snapsyncEpochUpd
	snapsyncCancelCmd *snapsyncCancelCmd
}

type snapsyncState struct {
	epoch     idx.Epoch
	cancel    func() error
	updatesCh chan snapsyncStateUpd
	quit      chan struct{}
}

type handler struct {
	NetworkID uint64
	config    Config

	syncStatus syncStatus

	txpool   TxPool
	maxPeers int

	peers *peerSet

	txsCh  chan evmcore.NewTxsNotify
	txsSub notify.Subscription

	dagLeecher   *dagstreamleecher.Leecher
	dagSeeder    *dagstreamseeder.Seeder
	dagProcessor *dagprocessor.Processor
	dagFetcher   *itemsfetcher.Fetcher

	bvLeecher   *bvstreamleecher.Leecher
	bvSeeder    *bvstreamseeder.Seeder
	bvProcessor *bvprocessor.Processor

	brLeecher   *brstreamleecher.Leecher
	brSeeder    *brstreamseeder.Seeder
	brProcessor *brprocessor.Processor

	epLeecher   *epstreamleecher.Leecher
	epSeeder    *epstreamseeder.Seeder
	epProcessor *epprocessor.Processor

	process processCallback

	txFetcher *itemsfetcher.Fetcher

	checkers *eventcheck.Checkers

	msgSemaphore *datasemaphore.DataSemaphore

	store    *Store
	engineMu sync.Locker

	notifier             dagNotifier
	emittedEventsCh      chan *inter.EventPayload
	emittedEventsSub     notify.Subscription
	newEpochsCh          chan idx.Epoch
	newEpochsSub         notify.Subscription
	quitProgressBroadcast chan struct{}

	// channels for syncer, txsyncLoop
	txsyncCh chan *txsync
	quitSync chan struct{}

	// snapsync fields
	chain       *ethBlockChain
	snapLeecher *snapleecher.Leecher
	snapState   snapsyncState

	// wait group is used for graceful shutdowns during downloading
	// and processing
	loopsWg sync.WaitGroup
	wg      sync.WaitGroup
	peerWG  sync.WaitGroup
	started sync.WaitGroup

	logger.Instance
}

// newHandler returns a new VinuChain sub protocol manager. The VinuChain sub protocol manages peers capable
// with the VinuChain network.
func newHandler(
	c handlerConfig,
) (
	*handler,
	error,
) {
	// Create the protocol manager with the base fields
	h := &handler{
		NetworkID:            c.s.GetRules().NetworkID,
		config:               c.config,
		notifier:             c.notifier,
		txpool:               c.txpool,
		msgSemaphore:         datasemaphore.New(c.config.Protocol.MsgsSemaphoreLimit, getSemaphoreWarningFn("P2P messages")),
		store:                c.s,
		process:              c.process,
		checkers:             c.checkers,
		peers:                newPeerSet(),
		engineMu:             c.engineMu,
		txsyncCh:             make(chan *txsync),
		quitSync:             make(chan struct{}),
		quitProgressBroadcast: make(chan struct{}),

		snapState: snapsyncState{
			updatesCh: make(chan snapsyncStateUpd, 128),
			quit:      make(chan struct{}),
		},

		Instance: logger.New("PM"),
	}
	h.started.Add(1)

	// TODO: configure it
	var (
		configBloomCache uint64 = 0 // Megabytes to alloc for fast sync bloom
	)

	var err error
	h.chain, err = newEthBlockChain(c.s)
	if err != nil {
		return nil, err
	}

	stateDb := h.store.EvmStore().EvmDb
	var stateBloom *trie.SyncBloom
	if false {
		// NOTE: Construct the downloader (long sync) and its backing state bloom if fast
		// sync is requested. The downloader is responsible for deallocating the state
		// bloom when it's done.
		// Note: we don't enable it if snap-sync is performed, since it's very heavy
		// and the heal-portion of the snap sync is much lighter than fast. What we particularly
		// want to avoid, is a 90%-finished (but restarted) snap-sync to begin
		// indexing the entire trie
		stateBloom = trie.NewSyncBloom(configBloomCache, stateDb)
	}
	h.snapLeecher = snapleecher.New(stateDb, stateBloom, h.removePeer)

	h.dagFetcher = itemsfetcher.New(h.config.Protocol.DagFetcher, itemsfetcher.Callback{
		OnlyInterested: func(ids []interface{}) []interface{} {
			return h.onlyInterestedEventsI(ids)
		},
		Suspend: func() bool {
			return false
		},
	})
	h.txFetcher = itemsfetcher.New(h.config.Protocol.TxFetcher, itemsfetcher.Callback{
		OnlyInterested: func(txids []interface{}) []interface{} {
			return txidsToInterfaces(h.txpool.OnlyNotExisting(interfacesToTxids(txids)))
		},
		Suspend: func() bool {
			return false
		},
	})

	h.dagProcessor = h.makeDagProcessor(c.checkers)
	h.dagLeecher = dagstreamleecher.New(h.store.GetEpoch(), h.store.GetHighestLamport() == 0, h.config.Protocol.DagStreamLeecher, dagstreamleecher.Callbacks{
		IsProcessed: h.store.HasEvent,
		RequestChunk: func(peer string, r dagstream.Request) error {
			p := h.peers.Peer(peer)
			if p == nil {
				return errNotRegistered
			}
			return p.RequestEventsStream(r)
		},
		Suspend: func(_ string) bool {
			return h.dagFetcher.Overloaded() || h.dagProcessor.Overloaded()
		},
		PeerEpoch: func(peer string) idx.Epoch {
			p := h.peers.Peer(peer)
			if p == nil || p.Useless() {
				return 0
			}
			return p.GetProgress().Epoch
		},
	})
	h.dagSeeder = dagstreamseeder.New(h.config.Protocol.DagStreamSeeder, dagstreamseeder.Callbacks{
		ForEachEvent: c.s.ForEachEventRLP,
	})

	h.bvProcessor = h.makeBvProcessor(c.checkers)
	h.bvLeecher = bvstreamleecher.New(h.config.Protocol.BvStreamLeecher, bvstreamleecher.Callbacks{
		LowestBlockToDecide: func() (idx.Epoch, idx.Block) {
			llrs := h.store.GetLlrState()
			epoch := h.store.FindBlockEpoch(llrs.LowestBlockToDecide)
			return epoch, llrs.LowestBlockToDecide
		},
		MaxEpochToDecide: func() idx.Epoch {
			if !h.syncStatus.RequestLLR() {
				return 0
			}
			return h.store.GetLlrState().LowestEpochToFill
		},
		IsProcessed: h.store.HasBlockVotes,
		RequestChunk: func(peer string, r bvstream.Request) error {
			p := h.peers.Peer(peer)
			if p == nil {
				return errNotRegistered
			}
			return p.RequestBVsStream(r)
		},
		Suspend: func(_ string) bool {
			return h.bvProcessor.Overloaded()
		},
		PeerBlock: func(peer string) idx.Block {
			p := h.peers.Peer(peer)
			if p == nil || p.Useless() {
				return 0
			}
			return p.GetProgress().LastBlockIdx
		},
	})
	h.bvSeeder = bvstreamseeder.New(h.config.Protocol.BvStreamSeeder, bvstreamseeder.Callbacks{
		Iterate: h.store.IterateOverlappingBlockVotesRLP,
	})

	h.brProcessor = h.makeBrProcessor()
	h.brLeecher = brstreamleecher.New(h.config.Protocol.BrStreamLeecher, brstreamleecher.Callbacks{
		LowestBlockToFill: func() idx.Block {
			return h.store.GetLlrState().LowestBlockToFill
		},
		MaxBlockToFill: func() idx.Block {
			if !h.syncStatus.RequestLLR() {
				return 0
			}
			// rough estimation for the max fill-able block
			llrs := h.store.GetLlrState()
			start := llrs.LowestBlockToFill
			end := llrs.LowestBlockToDecide
			if end > start+100 && h.store.HasBlock(start+100) {
				return start + 100
			}
			return end
		},
		IsProcessed: h.store.HasBlock,
		RequestChunk: func(peer string, r brstream.Request) error {
			p := h.peers.Peer(peer)
			if p == nil {
				return errNotRegistered
			}
			return p.RequestBRsStream(r)
		},
		Suspend: func(_ string) bool {
			return h.brProcessor.Overloaded()
		},
		PeerBlock: func(peer string) idx.Block {
			p := h.peers.Peer(peer)
			if p == nil || p.Useless() {
				return 0
			}
			return p.GetProgress().LastBlockIdx
		},
	})
	h.brSeeder = brstreamseeder.New(h.config.Protocol.BrStreamSeeder, brstreamseeder.Callbacks{
		Iterate: h.store.IterateFullBlockRecordsRLP,
	})

	h.epProcessor = h.makeEpProcessor(h.checkers)
	h.epLeecher = epstreamleecher.New(h.config.Protocol.EpStreamLeecher, epstreamleecher.Callbacks{
		LowestEpochToFetch: func() idx.Epoch {
			llrs := h.store.GetLlrState()
			if llrs.LowestEpochToFill < llrs.LowestEpochToDecide {
				return llrs.LowestEpochToFill
			}
			return llrs.LowestEpochToDecide
		},
		MaxEpochToFetch: func() idx.Epoch {
			if !h.syncStatus.RequestLLR() {
				return 0
			}
			return h.store.GetLlrState().LowestEpochToDecide + 10000
		},
		IsProcessed: h.store.HasHistoryBlockEpochState,
		RequestChunk: func(peer string, r epstream.Request) error {
			p := h.peers.Peer(peer)
			if p == nil {
				return errNotRegistered
			}
			return p.RequestEPsStream(r)
		},
		Suspend: func(_ string) bool {
			return h.epProcessor.Overloaded()
		},
		PeerEpoch: func(peer string) idx.Epoch {
			p := h.peers.Peer(peer)
			if p == nil || p.Useless() {
				return 0
			}
			return p.GetProgress().Epoch
		},
	})
	h.epSeeder = epstreamseeder.New(h.config.Protocol.EpStreamSeeder, epstreamseeder.Callbacks{
		Iterate: h.store.IterateEpochPacksRLP,
	})

	return h, nil
}

func (h *handler) peerMisbehaviour(peer string, err error) bool {
	if eventcheck.IsBan(err) {
		log.Warn("Dropping peer due to a misbehaviour", "peer", peer, "err", err)
		h.removePeer(peer)
		return true
	}
	return false
}

func (h *handler) removePeer(id string) {
	peer := h.peers.Peer(id)
	if peer != nil {
		peer.Peer.Disconnect(p2p.DiscUselessPeer)
	}
}

func (h *handler) unregisterPeer(id string) {
	// Short circuit if the peer was already removed
	peer := h.peers.Peer(id)
	if peer == nil {
		return
	}
	log.Debug("Removing peer", "peer", id)

	// Unregister the peer from the leecher's and seeder's and peer sets
	_ = h.epLeecher.UnregisterPeer(id)
	_ = h.epSeeder.UnregisterPeer(id)
	_ = h.dagLeecher.UnregisterPeer(id)
	_ = h.dagSeeder.UnregisterPeer(id)
	_ = h.brLeecher.UnregisterPeer(id)
	_ = h.brSeeder.UnregisterPeer(id)
	_ = h.bvLeecher.UnregisterPeer(id)
	_ = h.bvSeeder.UnregisterPeer(id)
	// Remove the `snap` extension if it exists
	if peer.snapExt != nil {
		_ = h.snapLeecher.SnapSyncer.Unregister(id)
	}
	if err := h.peers.UnregisterPeer(id); err != nil {
		log.Error("Peer removal failed", "peer", id, "err", err)
	}
}

func (h *handler) Start(maxPeers int) {
	h.snapsyncStageTick()

	h.maxPeers = maxPeers

	// broadcast transactions
	h.txsCh = make(chan evmcore.NewTxsNotify, txChanSize)
	h.txsSub = h.txpool.SubscribeNewTxsNotify(h.txsCh)

	h.loopsWg.Add(1)
	go h.txBroadcastLoop()

	if h.notifier != nil {
		// broadcast mined events
		h.emittedEventsCh = make(chan *inter.EventPayload, 4)
		h.emittedEventsSub = h.notifier.SubscribeNewEmitted(h.emittedEventsCh)
		// epoch changes
		h.newEpochsCh = make(chan idx.Epoch, 4)
		h.newEpochsSub = h.notifier.SubscribeNewEpoch(h.newEpochsCh)

		h.loopsWg.Add(3)
		go h.emittedBroadcastLoop()
		go h.progressBroadcastLoop()
		go h.onNewEpochLoop()
	}

	// start sync handlers
	go h.txsyncLoop()
	h.loopsWg.Add(2)
	go h.snapsyncStateLoop()
	go h.snapsyncStageLoop()
	h.dagFetcher.Start()
	h.txFetcher.Start()
	h.checkers.Heavycheck.Start()

	h.epProcessor.Start()
	h.epSeeder.Start()
	h.epLeecher.Start()

	h.dagProcessor.Start()
	h.dagSeeder.Start()
	h.dagLeecher.Start()

	h.bvProcessor.Start()
	h.bvSeeder.Start()
	h.bvLeecher.Start()

	h.brProcessor.Start()
	h.brSeeder.Start()
	h.brLeecher.Start()
	h.started.Done()
}

func (h *handler) Stop() {
	log.Info("Stopping VinuChain protocol")

	h.brLeecher.Stop()
	h.brSeeder.Stop()
	h.brProcessor.Stop()

	h.bvLeecher.Stop()
	h.bvSeeder.Stop()
	h.bvProcessor.Stop()

	h.dagLeecher.Stop()
	h.dagSeeder.Stop()
	h.dagProcessor.Stop()

	h.epLeecher.Stop()
	h.epSeeder.Stop()
	h.epProcessor.Stop()

	h.checkers.Heavycheck.Stop()
	h.txFetcher.Stop()
	h.dagFetcher.Stop()

	close(h.quitProgressBroadcast)
	close(h.snapState.quit)
	h.txsSub.Unsubscribe() // quits txBroadcastLoop
	if h.notifier != nil {
		h.emittedEventsSub.Unsubscribe() // quits eventBroadcastLoop
		h.newEpochsSub.Unsubscribe()     // quits onNewEpochLoop
	}

	// Wait for the subscription loops to come down.
	h.loopsWg.Wait()

	h.msgSemaphore.Terminate()
	// Quit the sync loop.
	// After this send has completed, no new peers will be accepted.
	close(h.quitSync)

	// Disconnect existing sessions.
	// This also closes the gate for any new registrations on the peer set.
	// sessions which are already established but not added to h.peers yet
	// will exit when they try to register.
	h.peers.Close()

	// Wait for all peer handler goroutines to come down.
	h.wg.Wait()
	h.peerWG.Wait()

	log.Info("VinuChain protocol stopped")
}

func (h *handler) myProgress() PeerProgress {
	bs := h.store.GetBlockState()
	epoch := h.store.GetEpoch()
	return PeerProgress{
		Epoch:            epoch,
		LastBlockIdx:     bs.LastBlock.Idx,
		LastBlockAtropos: bs.LastBlock.Atropos,
	}
}

// handle is the callback invoked to manage the life cycle of a peer. When
// this function terminates, the peer is disconnected.
func (h *handler) handle(p *peer) error {
	// If the peer has a `snap` extension, wait for it to connect so we can have
	// a uniform initialization/teardown mechanism
	snap, err := h.peers.WaitSnapExtension(p)
	if err != nil {
		p.Log().Error("Snapshot extension barrier failed", "err", err)
		return err
	}
	useless := discfilter.Banned(p.Node().ID(), p.Node().Record())
	if !useless && (!eligibleForSnap(p.Peer) || !strings.Contains(strings.ToLower(p.Name()), "opera")) {
		useless = true
		discfilter.Ban(p.ID())
	}
	if !p.Peer.Info().Network.Trusted && useless {
		if h.peers.UselessNum() >= h.maxPeers/10 {
			// don't allow more than 10% of useless peers
			return p2p.DiscTooManyPeers
		}
		p.SetUseless()
	}

	h.peerWG.Add(1)
	defer h.peerWG.Done()

	// Execute the handshake
	var (
		genesis    = *h.store.GetGenesisID()
		myProgress = h.myProgress()
	)
	if err := p.Handshake(h.NetworkID, myProgress, common.Hash(genesis)); err != nil {
		p.Log().Debug("Handshake failed", "err", err)
		if !useless {
			discfilter.Ban(p.ID())
		}
		return err
	}

	// Ignore maxPeers if this is a trusted peer
	if h.peers.Len() >= h.maxPeers && !p.Peer.Info().Network.Trusted {
		return p2p.DiscTooManyPeers
	}
	p.Log().Debug("Peer connected", "name", p.Name())

	// Register the peer locally
	if err := h.peers.RegisterPeer(p, snap); err != nil {
		p.Log().Warn("Peer registration failed", "err", err)
		return err
	}
	if err := h.dagLeecher.RegisterPeer(p.id); err != nil {
		p.Log().Warn("Leecher peer registration failed", "err", err)
		return err
	}
	if p.RunningCap(ProtocolName, []uint{FTM63}) {
		if err := h.epLeecher.RegisterPeer(p.id); err != nil {
			p.Log().Warn("Leecher peer registration failed", "err", err)
			return err
		}
		if err := h.bvLeecher.RegisterPeer(p.id); err != nil {
			p.Log().Warn("Leecher peer registration failed", "err", err)
			return err
		}
		if err := h.brLeecher.RegisterPeer(p.id); err != nil {
			p.Log().Warn("Leecher peer registration failed", "err", err)
			return err
		}
	}
	if snap != nil {
		if err := h.snapLeecher.SnapSyncer.Register(snap); err != nil {
			p.Log().Error("Failed to register peer in snap syncer", "err", err)
			return err
		}
	}
	defer h.unregisterPeer(p.id)

	// Propagate existing transactions. new transactions appearing
	// after this will be sent via broadcasts.
	h.syncTransactions(p, h.txpool.SampleHashes(h.config.Protocol.MaxInitialTxHashesSend))

	// Handle incoming messages until the connection is torn down
	for {
		if err := h.handleMsg(p); err != nil {
			p.Log().Debug("Message handling failed", "err", err)
			return err
		}
	}
}

// NodeInfo represents a short summary of the sub-protocol metadata
// known about the host peer.
type NodeInfo struct {
	Network     uint64      `json:"network"` // network ID
	Genesis     common.Hash `json:"genesis"` // SHA3 hash of the host's genesis object
	Epoch       idx.Epoch   `json:"epoch"`
	NumOfBlocks idx.Block   `json:"blocks"`
	//Config  *params.ChainConfig `json:"config"`  // Chain configuration for the fork rules
}

// NodeInfo retrieves some protocol metadata about the running host node.
func (h *handler) NodeInfo() *NodeInfo {
	numOfBlocks := h.store.GetLatestBlockIndex()
	return &NodeInfo{
		Network:     h.NetworkID,
		Genesis:     common.Hash(*h.store.GetGenesisID()),
		Epoch:       h.store.GetEpoch(),
		NumOfBlocks: numOfBlocks,
	}
}

func getSemaphoreWarningFn(name string) func(dag.Metric, dag.Metric, dag.Metric) {
	return func(received dag.Metric, processing dag.Metric, releasing dag.Metric) {
		log.Warn(fmt.Sprintf("%s semaphore inconsistency", name),
			"receivedNum", received.Num, "receivedSize", received.Size,
			"processingNum", processing.Num, "processingSize", processing.Size,
			"releasingNum", releasing.Num, "releasingSize", releasing.Size)
	}
}
