## Fuzzing

Fuzzing hardens the code that parses untrusted input. Here it targets the P2P
message-handling path (`gossip/handler.handleMsg`), which decodes messages from
network peers. It uses Go's built-in fuzzing (`go test -fuzz`), not the
abandoned `dvyukov/go-fuzz` toolchain.

### Run locally

```
make fuzz-native
```

or directly:

```
go test -run='^$' -fuzz='^FuzzHandleMsg$' -fuzztime=60s ./gossip/
```

### Targets (`gossip/handler_fuzz_test.go`)

- `FuzzHandleMsg` — fast per-PR target: a synced handler exercising the
  non-stream message decoders. Run by the `fuzz` job in
  `.github/workflows/ci.yml` on every push / PR.
- `FuzzHandleMsgDeep` — also starts the dag/bv/br/ep seeders so the
  `Request*Stream` codes are exercised (event-stage decoders).
- `FuzzHandleMsgDeepLLR` — runs in the block-records sync stage so the BV/BR/EP
  stream-response decoders are reached.

The two deep targets run nightly via `.github/workflows/fuzz-nightly.yml`, which
persists the generated corpus across runs as a workflow artifact.

### Crashers

If the fuzzer finds a crashing input, Go writes the reproducer under
`gossip/testdata/fuzz/<Target>/` and fails the run. Commit that file as a
regression seed once the underlying bug is fixed, then re-run the target to
confirm it no longer reproduces.
