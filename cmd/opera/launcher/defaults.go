package launcher

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/Fantom-foundation/go-opera/integration"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/nat"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	DefaultP2PPort  = 5050  // Default p2p port for listening
	DefaultHTTPPort = 18545 // Default TCP port for the HTTP RPC server
	DefaultWSPort   = 18546 // Default TCP port for the websocket RPC server
)

func overrideFlags() {
	utils.ListenPortFlag.Value = DefaultP2PPort
	utils.HTTPPortFlag.Value = DefaultHTTPPort
	utils.LegacyRPCPortFlag.Value = DefaultHTTPPort
	utils.WSPortFlag.Value = DefaultWSPort
}

// NodeDefaultConfig contains reasonable default settings.
var NodeDefaultConfig = node.Config{
	DataDir:             DefaultDataDir(),
	HTTPPort:            DefaultHTTPPort,
	HTTPModules:         []string{},
	HTTPVirtualHosts:    []string{"localhost"},
	HTTPTimeouts:        rpc.DefaultHTTPTimeouts,
	WSPort:              DefaultWSPort,
	WSModules:           []string{},
	GraphQLVirtualHosts: []string{"localhost"},
	P2P: p2p.Config{
		NoDiscovery: false, // enable discovery v4 by default
		DiscoveryV5: true,  // enable discovery v5 by default
		ListenAddr:  fmt.Sprintf(":%d", DefaultP2PPort),
		MaxPeers:    50,
		NAT:         nat.Any(),
	},
}

// DefaultDataDir is the default data directory to use for the databases and other
// persistence requirements. The resolution probes the filesystem and is reached
// from package-level initialisers, so it is memoised: it must not re-run, and
// must not re-emit its diagnostic, once per call site.
func DefaultDataDir() string {
	dir, _ := defaultDataDir()
	return dir
}

// DefaultDataDirNote returns a human-readable diagnostic about how the default
// data directory was resolved, or "" when there is nothing to report. It is
// deliberately not logged at resolution time: the default is frequently
// computed on runs that then override it with --datadir, and logging there
// reports a directory the node never opens. It is surfaced only once the node
// is about to open that directory — see LogDefaultDataDirNote.
func DefaultDataDirNote() string {
	_, note := defaultDataDir()
	return note
}

var (
	defaultDataDirOnce sync.Once
	defaultDataDirVal  string
	defaultDataDirNote string
)

func defaultDataDir() (string, string) {
	defaultDataDirOnce.Do(func() {
		defaultDataDirVal, defaultDataDirNote = resolveDataDir(homeDir(), runtime.GOOS, windowsAppData())
	})
	return defaultDataDirVal, defaultDataDirNote
}

// resolveDataDir is the pure, testable core of DefaultDataDir. It takes the
// environment it depends on as arguments so every platform branch can be
// exercised from a test on any host.
func resolveDataDir(home, goos, winAppData string) (dir string, note string) {
	if home == "" {
		// As we cannot guess a stable location, return empty and handle later.
		return "", ""
	}
	switch goos {
	case "darwin":
		legacy := filepath.Join(home, "Library", "Lachesis")
		if isNonEmptyDir(legacy) {
			return legacy, ""
		}
		return filepath.Join(home, "Library", "VinuChain"), ""
	case "windows":
		legacyFallback := filepath.Join(home, "AppData", "Roaming", "Lachesis")
		if winAppData == "" || isNonEmptyDir(legacyFallback) {
			return legacyFallback, ""
		}
		legacy := filepath.Join(winAppData, "Lachesis")
		if isNonEmptyDir(legacy) {
			return legacy, ""
		}
		return filepath.Join(winAppData, "VinuChain"), ""
	default:
		legacyDir := filepath.Join(home, ".opera")
		newDir := filepath.Join(home, ".vinuchain")
		if _, err := os.Stat(legacyDir); err != nil {
			return newDir, ""
		}
		// The legacy directory keeps winning until .vinuchain independently
		// holds openable chain state. Mere existence of .vinuchain is not
		// enough: read-only subcommands (account/console/attach) create it as a
		// side effect, a flagless start that dies on a missing --genesis leaves
		// it behind, and an interrupted genesis leaves only a marker that gets
		// dropped on the next run. Treating any of those as a completed
		// migration would strand a populated .opera and resync from genesis.
		//
		// Consequence, deliberately: the migration no longer completes on its
		// own. Moving to .vinuchain is an explicit operator action — copy the
		// directory and restart with --datadir.
		if isUnusedDataDir(newDir) {
			return legacyDir, fmt.Sprintf(
				"Using legacy data directory %s. To migrate: stop the node, copy it to %s, and restart with an explicit --datadir",
				legacyDir, newDir)
		}
		return newDir, fmt.Sprintf(
			"Both %s and %s carry chain state; using %s. Remove the unused directory and pass --datadir explicitly to avoid ambiguity",
			legacyDir, newDir, newDir)
	}
}

// isUnusedDataDir reports whether dir is positively known to hold no chain
// state a node could open — an empty directory, the go-opera/+keystore/ shell
// that read-only subcommands leave behind, or the remains of an interrupted
// genesis that MakeEngine drops and re-initialises.
//
// Inability to inspect answers false, deliberately.
// integration.FirstLaunchPending treats an unreadable directory as empty, which
// is correct for "should I apply a genesis here" but wrong for this decision: a
// .vinuchain whose chaindata cannot be read must keep priority so the node
// fails loudly on it, rather than silently opening older state from .opera.
//
// Only the top-level chaindata is consulted. --fakenet state lives in fakenet-N
// subdirectories, and a stray dev fakenet must not decide the default for a
// production start; letting it would strand exactly the real .opera state this
// fallback exists to protect. The accepted consequence is dev-only: a --fakenet
// run on a host that still has a legacy .opera builds its network under .opera
// rather than reusing one under .vinuchain. Pass --datadir to pin it.
func isUnusedDataDir(dir string) bool {
	chaindata := filepath.Join(dir, "chaindata")
	f, err := os.Open(chaindata)
	if err != nil {
		return os.IsNotExist(err)
	}
	_ = f.Close()
	return integration.FirstLaunchPending(chaindata)
}

// LogDefaultDataDirNote surfaces how the default data directory resolved, for
// the node that is about to open dataDir.
//
// It is deliberately not called while configuration is merely being assembled.
// dumpconfig and checkconfig build a full config without opening anything, and
// an explicit --datadir or a config-file [Node] DataDir means the note
// describes a directory nothing will open — reporting it there is what made the
// legacy notice look like the node was using .opera when it was not.
func LogDefaultDataDirNote(dataDir string) {
	def, note := defaultDataDir()
	if note == "" || def == "" {
		return
	}
	if !usesDefaultDataDir(dataDir, def) {
		return
	}
	log.Warn(note)
}

// usesDefaultDataDir reports whether dataDir is the resolved default or lives
// beneath it, which is what --fakenet does with its fakenet-N directories.
func usesDefaultDataDir(dataDir, def string) bool {
	return dataDir == def || strings.HasPrefix(dataDir, def+string(filepath.Separator))
}

func windowsAppData() string {
	v := os.Getenv("LOCALAPPDATA")
	if v == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return ""
		}
		return filepath.Join(home, "AppData", "Local")
	}
	return v
}

func isNonEmptyDir(dir string) bool {
	f, err := os.Open(dir)
	if err != nil {
		return false
	}
	names, _ := f.Readdir(1)
	_ = f.Close()
	return len(names) > 0
}

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}
