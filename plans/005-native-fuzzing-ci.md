# Plan 005: Wire continuous fuzzing into CI (native Go fuzzing)

> **Executor instructions**: This is a migration + integration task with real
> unknowns — treat the early steps as investigation and STOP-and-report rather
> than forcing a result. Run every verification command. When done (or blocked),
> update the status row for this plan in `plans/README.md`.
>
> **Drift check (run first)**: `git diff --stat af41ca7..HEAD -- gossip/handler_fuzz.go gossip/handler.go gossip/service.go .github/workflows/ci.yml`
> If `gossip/handler_fuzz.go` or the handler wiring changed, re-read them and
> compare against "Current state" before proceeding.

## Status

- **Priority**: P2
- **Effort**: L
- **Risk**: LOW (new test-only code + a CI job; touches no production logic)
- **Depends on**: none
- **Category**: tests / dx / security
- **Planned at**: commit `af41ca7`, 2026-06-16

## Why this matters

VinuChain is a consensus node that decodes messages from untrusted network
peers — exactly the surface `FUZZING.md` says fuzzing exists to harden. A fuzz
harness already exists (`gossip/handler_fuzz.go`, target `FuzzHandler`), but:

1. It uses the **abandoned** `github.com/dvyukov/go-fuzz` toolchain (the
   `//go:build gofuzz` file + `make fuzz`), which predates Go's native fuzzing
   (Go 1.18+). Nobody runs it.
2. It is **never executed in CI**, so regressions in the P2P decode path are not
   caught continuously.
3. The harness itself is **incomplete** — `makeFuzzedHandler` has a literal
   `// TODO: init` where `heavyCheckReader` / `gasPowerCheckReader` are left as
   zero-value, so the deep validation path may not be exercised (or may panic).

This plan migrates the harness to **native Go fuzzing** (`go test -fuzz`) and
adds a **time-boxed** CI job that fuzzes the P2P message-handling path on every
push/PR. The minimum-viable outcome is a native fuzz target that compiles, runs
for the CI budget without spurious failures, and is wired into CI. Fully wiring
the heavy-check readers is a stretch goal with an explicit fallback.

## Current state

- `gossip/handler_fuzz.go` — behind `//go:build gofuzz`, so it is **not** compiled
  during a normal `go test`. It defines (study these; you will adapt them):
  - `FuzzHandler(data []byte) int` — feeds `data` into a constructed `*handler`
    via `fuzzedHandler.handleMsg(other)`.
  - `makeFuzzedHandler() (*handler, error)` — builds a handler from a fakenet
    genesis (`makefakegenesis.FakeGenesisStore(3, ...)`), `NewMemStore()`,
    `DefaultConfig(cachescale.Identity)`, `makeCheckers(...)`,
    `evmcore.NewTxPool(...)`, `newHandler(handlerConfig{...})`, then `h.Start(3)`.
    **Contains the gap:**
    ```go
    var (
        network             = opera.FakeNetRules()
        heavyCheckReader    HeavyCheckReader
        gasPowerCheckReader GasPowerCheckReader
        // TODO: init
    )
    ```
  - `newFuzzMsg(data []byte) (*p2p.Msg, error)` — maps `data[0]` to one of 8
    protocol message codes (`HandshakeMsg`, `EvmTxsMsg`, `ProgressMsg`,
    `NewEventIDsMsg`, `GetEventsMsg`, `EventsMsg`, `RequestEventsStream`,
    `EventsStreamResponse`) and wraps the rest as the payload.
  - `fuzzMsgReadWriter`, `randomID()` helpers.
- These symbols are compiled **only** with `-tags gofuzz`, so a new normal
  `_test.go` file can define its own (use distinct names to avoid confusion).
- `Makefile` has a `fuzz` target invoking `go-fuzz-build` + `go-fuzz` (legacy).
- `.github/workflows/ci.yml` — no fuzz job.
- **The `// TODO: init` gap is a real trap — do NOT just port the legacy harness.**
  Leaving `heavyCheckReader`/`gasPowerCheckReader` zero-value is not harmless:
  `makeCheckers` (`gossip/service.go:497`) always builds the gaspower checker, and
  the zero-value `gasPowerCheckReader.Ctx` is an empty `atomic.Value`, so
  `GetValidationContext()`'s `Ctx.Load().(*gaspowercheck.ValidationContext)`
  type-asserts a nil → **panics** in a background goroutine once `h.Start(...)`
  runs and any event is enqueued. A harness built that way reports panics that are
  setup bugs, not attacker-input findings.
- **Use the repo's already-wired path instead.** `gossip/common_test.go:133`
  defines `newTestEnv(firstEpoch idx.Epoch, validatorsNum idx.Validator) *testEnv`,
  which builds a full `Service` via `newService(...)`. `newService` initializes
  BOTH readers at `gossip/service.go:411-413`
  (`svc.heavyCheckReader.Store = store`; `svc.gasPowerCheckReader.Ctx.Store(...)`).
  `testEnv` embeds `*Service`, and `Service` has an unexported `handler *handler`
  field (`gossip/service.go:146`, set at `:453`) — reachable from test code in
  `package gossip` as `env.handler`. `gossip/heavycheck_test.go:34` uses
  `newTestEnv(startEpoch, validatorsNum)` as the working exemplar. This is the
  construction the fuzz target should use.
- For reference, `Config.HeavyCheck` is typed `heavycheck.Config` (a struct
  `{MaxQueuedTasks, Threads}` — `gossip/config.go:89`), **not a bool** — there is
  no "disable heavy checks" flag to set. Do not try to turn checks off; wire the
  readers correctly via `newTestEnv` instead.

## Commands you will need

| Purpose | Command | Expected |
|---------|---------|----------|
| Build (sanity) | `go build ./...` | exit 0 |
| Compile + run normal tests | `go test -run=^$ -count=1 ./gossip/` | exit 0 (new file compiles; existing tests run) |
| Run fuzz briefly | `go test -run=^$ -fuzz=FuzzHandleMsg -fuzztime=30s ./gossip/` | runs ~30s, exits 0, no crash file written |
| Inspect wired construction | `grep -rn "func newTestEnv\|svc.handler\|gasPowerCheckReader" gossip/common_test.go gossip/service.go` | the wired path |
| YAML sanity | `python3 -c "import yaml; yaml.safe_load(open('.github/workflows/ci.yml'))"` | exit 0 (skip if PyYAML absent) |

## Scope

**In scope** (create/modify):
- `gossip/handler_fuzz_test.go` (create — the native fuzz target + adapted helpers)
- `.github/workflows/ci.yml` (add one job)
- Optionally `Makefile` (add a `fuzz-native` target; you MAY leave the legacy
  `fuzz` target in place or note it for retirement — do not delete it without
  saying so)

**Out of scope** (do NOT touch):
- Any production `.go` file in `gossip/` (handler logic, checkers, txpool). The
  fuzz target must adapt to the code as-is; if it can only run by changing
  production code, STOP and report.
- `gossip/handler_fuzz.go` — leave the legacy harness as-is (a later cleanup can
  retire it once the native target is proven).

## Git workflow

- Branch: `advisor/005-native-fuzzing`.
- Commits: Conventional Commits, e.g. `test(gossip): add native fuzz target for P2P handleMsg` and `ci: run native fuzzing (time-boxed)`.
- Do NOT push or open a PR unless instructed.

## Steps

### Step 1: Confirm the wired construction path

Read `gossip/handler_fuzz.go` (the legacy harness), `gossip/common_test.go`
around `newTestEnv` (line 133), `gossip/heavycheck_test.go` around line 34 (the
exemplar caller), and `gossip/service.go:359` (`newService`) + `:411-413` (where
the readers are initialized) + `:146`/`:453` (the `handler` field). Run:

```sh
grep -rn "func newTestEnv\|env.handler\|\.handler\b" gossip/common_test.go gossip/*_test.go
grep -n "heavyCheckReader\|gasPowerCheckReader\|svc.handler" gossip/service.go
```

**Verify**: you can state, in your report, (a) that `newTestEnv` → `newService`
initializes both readers, and (b) how to reach the `*handler` from the returned
`*testEnv` (it is `env.handler` via the embedded `*Service`). If `newTestEnv` or
the embedded-`Service` handler field no longer exists, STOP — the construction
basis for this plan has drifted.

### Step 2: Create a native fuzz target that compiles

Create `gossip/handler_fuzz_test.go` (a **normal** test file — no build tag).
**Do NOT port the legacy `makeFuzzedHandler`** — it has the `// TODO: init` gap
and panics (see Current state). Instead construct the handler from a fully-wired
`testEnv`. Port only the message helpers `newFuzzMsg`, `fuzzMsgReadWriter`,
`randomID` under **distinct names** (`decodeFuzzMsg`, `fuzzRW`, `randFuzzID`) so
they do not overlap with the gofuzz-tagged file. The required imports (`testing`,
`bytes`, `errors`, `github.com/ethereum/go-ethereum/p2p`,
`.../p2p/enode`) can be read from `gossip/handler_fuzz.go`.

Target shape (reach the handler via the wired env):

```go
func FuzzHandleMsg(f *testing.F) {
    // Build a fully-initialized service+handler (readers wired by newService).
    env := newTestEnv(1, 3)        // firstEpoch=1, 3 validators — mirror heavycheck_test.go
    defer env.Close()
    h := env.handler               // *handler from the embedded *Service (service.go:146)

    // Seed corpus: byte[0] selects a message code, the rest is the payload.
    f.Add([]byte{0x00})
    f.Add([]byte{0x04, 0x01, 0x02, 0x03})

    f.Fuzz(func(t *testing.T, data []byte) {
        msg, err := decodeFuzzMsg(data)
        if err != nil {
            return // not interesting
        }
        peer := &peer{
            version: ProtocolVersion,
            Peer:    p2p.NewPeer(randFuzzID(), "fuzz-peer", []p2p.Cap{}),
            rw:      &fuzzRW{msg},
        }
        // handleMsg returning an error on malformed input is EXPECTED and fine;
        // a panic/hang is the bug the fuzzer hunts.
        _ = h.handleMsg(peer)
    })
}
```

> If `handleMsg` requires the handler to have been started (the legacy harness
> called `h.Start(3)`), call the equivalent start on `env.handler` once before
> `f.Fuzz` — determine from `gossip/handler.go`/`service.go` whether `newTestEnv`
> already starts it (it starts `blockProcTasks`/`verWatcher` but may not start the
> handler). Whatever you do, the handler MUST be backed by `newService`'s wired
> readers — never by a zero-value reader.

**Verify**: `go test -run=^$ -count=1 ./gossip/` → compiles and runs the normal
tests (0 fuzz iterations because no `-fuzz` flag). It must exit 0. (`-count=1`,
not `-count=0`; `-count=0` is an unusual no-op — `-count=1` compiles AND confirms
the package's existing tests still pass with your new file present.)

### Step 3: Run the fuzzer briefly and confirm no immediate crash

```sh
go test -run=^$ -fuzz=FuzzHandleMsg -fuzztime=30s ./gossip/
```

**Verify**: it runs for ~30s and exits 0 with no `--- FAIL` and no
`testdata/fuzz/FuzzHandleMsg/<hash>` crash corpus file written.

If it **does** find a crasher: that is a genuine finding. Capture the
`testdata/fuzz/...` reproducer, STOP, and report it — do not "fix" production
code under this plan. A reproducible panic in `handleMsg` on peer input is
exactly what this harness is meant to surface and would be its own remediation.

### Step 4: Distinguish a real crasher from a harness-setup panic

Because the handler is built via `newTestEnv`/`newService`, the readers are
already wired, so the deep validation path is reachable — you do NOT need to
disable anything. But before treating any panic as a finding, confirm it
originates from **fed input**, not harness setup:

- A panic whose stack runs through `handleMsg` → message decode/validation on
  the fuzz `data` is a genuine finding → follow Step 3's STOP (capture the
  reproducer, report).
- A panic during `newTestEnv`/`env.handler` setup, or a `nil`-context panic in
  `gasPowerCheckReader.GetValidationContext` that fires regardless of input
  (i.e. it also fires on the `f.Add` seeds), is a **harness bug**, not a finding.
  If you see this, your handler is not actually backed by `newService`'s wired
  readers — re-check Step 2. Do NOT report it as a vulnerability, and do NOT
  "fix" production code; fix the harness construction.

**Verify**: the 30s run from Step 3 completes with no crash file, OR any crash
you do surface has a `handleMsg`-rooted stack tied to the fuzz input (record
which in your report).

### Step 5: Add a time-boxed fuzz job to CI

```yaml
  fuzz:
    name: fuzz (time-boxed)
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version-file: go.mod
          cache: false

      # Native Go fuzzing, bounded so it fits a normal CI run. Any crash the
      # fuzzer finds fails the job and is saved under testdata/fuzz/.
      - name: Fuzz P2P handleMsg
        run: go test -run=^$ -fuzz=FuzzHandleMsg -fuzztime=120s ./gossip/
```

**Verify**: `python3 -c "import yaml; yaml.safe_load(open('.github/workflows/ci.yml'))"` → exit 0.

### Step 6 (optional): add a `make fuzz-native` convenience target

```makefile
.PHONY: fuzz-native
fuzz-native:
	go test -run=^$$ -fuzz=FuzzHandleMsg -fuzztime=60s ./gossip/
```

**Verify**: `make fuzz-native` runs ~60s and exits 0.

## Test plan

- The fuzz target compiles and the package's tests still pass:
  `go test -run=^$ -count=1 ./gossip/` exits 0.
- A 30s local fuzz run finds no crasher (Step 3), or any crasher is `handleMsg`-rooted (Step 4).
- The existing suite is unaffected: `go test ./gossip/` exits 0 (the new
  `_test.go` adds a fuzz func but no failing unit test).
- CI YAML parses (skip if PyYAML absent).

## Done criteria

ALL must hold:

- [ ] `gossip/handler_fuzz_test.go` exists and `go test -run=^$ -count=1 ./gossip/` exits 0.
- [ ] `go test -run=^$ -fuzz=FuzzHandleMsg -fuzztime=30s ./gossip/` exits 0 with no
      crasher (or any crasher is a `handleMsg`-rooted finding per Step 4, reported).
- [ ] `go test ./gossip/` exits 0 (no regression to the normal suite).
- [ ] `.github/workflows/ci.yml` has a `fuzz` job invoking `-fuzz=FuzzHandleMsg`
      (`grep -n "FuzzHandleMsg" .github/workflows/ci.yml` → ≥1).
- [ ] `python3 -c "import yaml; yaml.safe_load(open('.github/workflows/ci.yml'))"` exits 0.
- [ ] No production `.go` file modified (`git diff --name-only af41ca7 -- 'gossip/*.go' | grep -v _test.go` is empty).
- [ ] `plans/README.md` status row for 005 updated, with the Step 4 decision and
      any Step 3 crasher recorded.

## STOP conditions

Stop and report back (do not improvise) if:

- The fuzzer finds a reproducible crash in `handleMsg` (capture the
  `testdata/fuzz/...` reproducer — this is a real finding for separate remediation).
- The only way to make the target compile/run is to modify production `gossip/`
  code.
- `newTestEnv` no longer exists, no longer wires the readers, or the `*handler`
  is no longer reachable from the returned env — the construction basis has
  drifted; report what changed (do NOT fall back to the panicking legacy harness).
- The live `gossip/handler_fuzz.go`, `gossip/common_test.go` (`newTestEnv`), or
  `gossip/service.go` reader-init (`:411-413`) differs structurally from "Current state".

## Maintenance notes

- Once the native target is proven, the legacy `gossip/handler_fuzz.go`
  (`//go:build gofuzz`) and the `Makefile` `fuzz` target can be retired in a
  follow-up — call that out for the reviewer but don't bundle it here.
- A longer scheduled fuzz (e.g. a nightly `schedule:` workflow with a multi-minute
  `-fuzztime` and a persisted corpus artifact) is a high-value follow-up; the
  per-PR 120s job is the floor, not the ceiling.
- If `f.Fuzz` seed corpus is committed under `gossip/testdata/fuzz/`, keep it —
  it accelerates convergence and pins past crashers as regression tests.
- Reviewer should confirm the job's `-fuzztime` keeps total CI duration acceptable
  and that a found crasher actually fails the job (fail-closed).
