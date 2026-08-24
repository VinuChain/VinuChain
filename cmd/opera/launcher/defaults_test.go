package launcher

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// mkChainState makes dir look like a data directory that has been run: a
// non-empty chaindata/.
func mkChainState(t *testing.T, dir string) string {
	t.Helper()
	cd := filepath.Join(dir, "chaindata")
	if err := os.MkdirAll(cd, 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(cd, "CURRENT"), []byte("x"), 0600); err != nil {
		t.Fatal(err)
	}
	return dir
}

// mkShell reproduces what a read-only subcommand (`opera account list`) leaves
// behind: go-opera/nodekey, go-opera/LOCK and keystore/, but no chaindata.
func mkShell(t *testing.T, dir string) string {
	t.Helper()
	for _, sub := range []string{"go-opera", "keystore"} {
		if err := os.MkdirAll(filepath.Join(dir, sub), 0700); err != nil {
			t.Fatal(err)
		}
	}
	for _, f := range []string{"nodekey", "LOCK"} {
		if err := os.WriteFile(filepath.Join(dir, "go-opera", f), []byte("x"), 0600); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

func mkEmpty(t *testing.T, dir string) string {
	t.Helper()
	if err := os.MkdirAll(dir, 0700); err != nil {
		t.Fatal(err)
	}
	return dir
}

// mkInterrupted makes dir hold a chaindata containing only the "unfinished"
// marker of a genesis that was killed part-way.
func mkInterrupted(t *testing.T, dir string) string {
	t.Helper()
	cd := filepath.Join(dir, "chaindata")
	if err := os.MkdirAll(cd, 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(cd, "unfinished"), nil, 0600); err != nil {
		t.Fatal(err)
	}
	return dir
}

func TestResolveDataDirLinux(t *testing.T) {
	const (
		wantLegacy = ".opera"
		wantNew    = ".vinuchain"
	)
	cases := []struct {
		name     string
		setup    func(t *testing.T, home string)
		want     string
		wantNote string // substring; "" means no note at all
	}{
		{
			name:  "fresh home uses the new directory",
			setup: func(t *testing.T, home string) {},
			want:  wantNew,
		},
		{
			name:  "populated legacy with no new directory stays on legacy",
			setup: func(t *testing.T, home string) { mkChainState(t, filepath.Join(home, wantLegacy)) },
			want:  wantLegacy, wantNote: "Using legacy data directory",
		},
		{
			// Regression: a flagless start that died on a missing --genesis
			// leaves an empty .vinuchain behind. It must not strand .opera.
			name: "populated legacy is not stranded by an empty new directory",
			setup: func(t *testing.T, home string) {
				mkChainState(t, filepath.Join(home, wantLegacy))
				mkEmpty(t, filepath.Join(home, wantNew))
			},
			want: wantLegacy, wantNote: "Using legacy data directory",
		},
		{
			// Regression: `opera account list` creates a go-opera/+keystore/
			// shell. Non-empty, but carries no chain state.
			name: "populated legacy is not stranded by a read-only-command shell",
			setup: func(t *testing.T, home string) {
				mkChainState(t, filepath.Join(home, wantLegacy))
				mkShell(t, filepath.Join(home, wantNew))
			},
			want: wantLegacy, wantNote: "Using legacy data directory",
		},
		{
			name: "new directory wins once it carries chain state",
			setup: func(t *testing.T, home string) {
				mkChainState(t, filepath.Join(home, wantLegacy))
				mkChainState(t, filepath.Join(home, wantNew))
			},
			want: wantNew, wantNote: "carry chain state",
		},
		{
			// Regression: MakeEngine drops an interrupted chaindata and then
			// requires --genesis, so the marker is not state anyone can start
			// from. Preferring it would turn a flagless start that would have
			// opened the populated .opera into a missing-genesis fatal.
			name: "interrupted genesis in the new directory does not strand legacy",
			setup: func(t *testing.T, home string) {
				mkChainState(t, filepath.Join(home, wantLegacy))
				mkInterrupted(t, filepath.Join(home, wantNew))
			},
			want: wantLegacy, wantNote: "Using legacy data directory",
		},
		{
			// Pre-existing behaviour, deliberately preserved: an empty legacy
			// directory still wins. Changing it would move hosts toward
			// .vinuchain, which is the unsafe direction.
			name:  "empty legacy directory still wins",
			setup: func(t *testing.T, home string) { mkEmpty(t, filepath.Join(home, wantLegacy)) },
			want:  wantLegacy, wantNote: "Using legacy data directory",
		},
		{
			name:  "no legacy directory uses the new one silently",
			setup: func(t *testing.T, home string) { mkChainState(t, filepath.Join(home, wantNew)) },
			want:  wantNew,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			home := t.TempDir()
			tc.setup(t, home)

			got, note := resolveDataDir(home, "linux", "")

			if want := filepath.Join(home, tc.want); got != want {
				t.Errorf("resolveDataDir() = %q, want %q", got, want)
			}
			switch {
			case tc.wantNote == "" && note != "":
				t.Errorf("unexpected note %q", note)
			case tc.wantNote != "" && !strings.Contains(note, tc.wantNote):
				t.Errorf("note = %q, want it to contain %q", note, tc.wantNote)
			}
		})
	}
}

// TestResolveDataDirOtherPlatforms locks in the current darwin and windows
// behaviour. These branches are deliberately untouched by the linux migration
// fix; the test exists so a future change to them is a visible decision.
func TestResolveDataDirOtherPlatforms(t *testing.T) {
	t.Run("darwin prefers a populated Lachesis directory", func(t *testing.T) {
		home := t.TempDir()
		legacy := filepath.Join(home, "Library", "Lachesis")
		mkEmpty(t, legacy)
		if err := os.WriteFile(filepath.Join(legacy, "x"), []byte("x"), 0600); err != nil {
			t.Fatal(err)
		}
		if got, _ := resolveDataDir(home, "darwin", ""); got != legacy {
			t.Errorf("got %q, want %q", got, legacy)
		}
	})

	t.Run("darwin falls through to VinuChain", func(t *testing.T) {
		home := t.TempDir()
		want := filepath.Join(home, "Library", "VinuChain")
		if got, _ := resolveDataDir(home, "darwin", ""); got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("windows without appdata uses the roaming fallback", func(t *testing.T) {
		home := t.TempDir()
		want := filepath.Join(home, "AppData", "Roaming", "Lachesis")
		if got, _ := resolveDataDir(home, "windows", ""); got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("windows falls through to VinuChain under appdata", func(t *testing.T) {
		home := t.TempDir()
		appdata := t.TempDir()
		want := filepath.Join(appdata, "VinuChain")
		if got, _ := resolveDataDir(home, "windows", appdata); got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("no home yields no directory", func(t *testing.T) {
		if got, note := resolveDataDir("", "linux", ""); got != "" || note != "" {
			t.Errorf("got (%q, %q), want empty", got, note)
		}
	})
}

// TestDefaultDataDirIsMemoised pins that repeated calls do not re-probe, which
// is what stopped the legacy notice being emitted once per package-level
// initialiser.
func TestDefaultDataDirIsMemoised(t *testing.T) {
	first := DefaultDataDir()
	for i := 0; i < 3; i++ {
		if got := DefaultDataDir(); got != first {
			t.Fatalf("DefaultDataDir() not stable: %q then %q", first, got)
		}
	}
	firstNote := DefaultDataDirNote()
	if again := DefaultDataDirNote(); again != firstNote {
		t.Errorf("DefaultDataDirNote() not stable: %q then %q", firstNote, again)
	}
}

// mkFakenet reproduces a --fakenet run, which lives in a fakenet-N
// subdirectory of the data directory rather than in chaindata/ directly.
func mkFakenet(t *testing.T, dir string, n int) string {
	t.Helper()
	mkChainState(t, filepath.Join(dir, fmt.Sprintf("fakenet-%d", n)))
	return dir
}

func TestIsUnusedDataDir(t *testing.T) {
	if dir := mkChainState(t, t.TempDir()); isUnusedDataDir(dir) {
		t.Error("populated chaindata is in use")
	}
	if dir := mkShell(t, t.TempDir()); !isUnusedDataDir(dir) {
		t.Error("a read-only-command shell holds no chain state")
	}
	if dir := mkEmpty(t, t.TempDir()); !isUnusedDataDir(dir) {
		t.Error("an empty directory holds no chain state")
	}
	if dir := mkInterrupted(t, t.TempDir()); !isUnusedDataDir(dir) {
		t.Error("an interrupted genesis is dropped by MakeEngine, so it is not openable state")
	}
	if dir := mkFakenet(t, t.TempDir(), 3); !isUnusedDataDir(dir) {
		t.Error("a fakenet subdirectory is not top-level chain state")
	}
}

// TestIsUnusedDataDirDoesNotGuessOnUnreadable pins that inability to inspect is
// not read as absence. integration.FirstLaunchPending reports an unreadable
// directory as empty; concluding "unused" from that would silently hand a node
// its older .opera state instead of failing on the datadir it was meant to open.
func TestIsUnusedDataDirDoesNotGuessOnUnreadable(t *testing.T) {
	if os.Geteuid() == 0 {
		t.Skip("root bypasses directory permissions")
	}
	dir := mkChainState(t, t.TempDir())
	chaindata := filepath.Join(dir, "chaindata")
	if err := os.Chmod(chaindata, 0000); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chmod(chaindata, 0700) })

	if isUnusedDataDir(dir) {
		t.Error("an unreadable chaindata must not be treated as unused")
	}
}

// TestResolveDataDirKeepsUnreadableNewDir is the same guarantee at the
// resolution level: a .vinuchain that cannot be inspected keeps priority.
func TestResolveDataDirKeepsUnreadableNewDir(t *testing.T) {
	if os.Geteuid() == 0 {
		t.Skip("root bypasses directory permissions")
	}
	home := t.TempDir()
	mkChainState(t, filepath.Join(home, ".opera"))
	newDir := mkChainState(t, filepath.Join(home, ".vinuchain"))
	chaindata := filepath.Join(newDir, "chaindata")
	if err := os.Chmod(chaindata, 0000); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chmod(chaindata, 0700) })

	got, _ := resolveDataDir(home, "linux", "")
	if want := filepath.Join(home, ".vinuchain"); got != want {
		t.Errorf("resolveDataDir() = %q, want %q", got, want)
	}
}

// TestResolveDataDirFakenetDoesNotStrandLegacy pins the priority between the
// two failure modes here: a dev fakenet sitting under .vinuchain must not
// redirect a production start away from real state in .opera. The accepted cost
// is that a --fakenet run then builds its network under .opera.
func TestResolveDataDirFakenetDoesNotStrandLegacy(t *testing.T) {
	home := t.TempDir()
	mkChainState(t, filepath.Join(home, ".opera"))
	mkFakenet(t, filepath.Join(home, ".vinuchain"), 3)

	got, _ := resolveDataDir(home, "linux", "")
	if want := filepath.Join(home, ".opera"); got != want {
		t.Errorf("resolveDataDir() = %q, want %q", got, want)
	}
}

func TestUsesDefaultDataDir(t *testing.T) {
	def := filepath.Join("/home/u", ".vinuchain")
	cases := []struct {
		dataDir string
		want    bool
	}{
		{def, true},
		{filepath.Join(def, "fakenet-3"), true}, // --fakenet lives beneath the default
		{filepath.Join("/home/u", ".opera"), false},
		{filepath.Join("/srv", "chain"), false},
		{def + "-backup", false}, // prefix string match must not count
		{"", false},
	}
	for _, tc := range cases {
		if got := usesDefaultDataDir(tc.dataDir, def); got != tc.want {
			t.Errorf("usesDefaultDataDir(%q, %q) = %v, want %v", tc.dataDir, def, got, tc.want)
		}
	}
}
