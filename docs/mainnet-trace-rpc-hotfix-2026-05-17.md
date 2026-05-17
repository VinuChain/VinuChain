# Mainnet Trace RPC Hotfix - 2026-05-17

This note records the mainnet RPC trace hotfix deployed for VinuExplorer
internal native-transfer indexing.

## Scope

- Source baseline: `v2.0.0-rc.1`
- Hotfix source commit: `60e43d6fa0f2a6d181fc6a0ef3fb897c709740df`
- Provenance tag: `v2.0.0-rc.1-trace.1`
- Review PR: `https://github.com/VinuChain/VinuChain/pull/14`
- Mainnet RPC host: `i-083fffeeb03583a18` (`ip-10-0-142-137`)
- Runtime binary: `/home/ubuntu/VinuChain/build/opera`
- Rollback binary: `/home/ubuntu/VinuChain/build/opera.pre-trace-20260517T100059Z`
- Built candidate SHA256:
  `1ce6cd848046974958cc9ef9b23447f0b95d81d0ed9189d0e758062e5556ecda`

The annotated tag intentionally points at the deployed source commit. Do not
move `v2.0.0-rc.1-trace.1`. If code changes are needed after this note, build a
new binary and tag a new source commit such as `v2.0.0-rc.1-trace.2`.

## Behavior

The hotfix registers the public `trace` RPC namespace and adds replay-backed
support for:

- `trace_block`
- `trace_transaction`
- `trace_get`

VinuExplorer uses `trace_block` to index native internal transfers emitted by
contracts. This fixed transaction
`0xdd2007fae2f91fd0db5321a06097706c738ddc79d147e3f6eb4acb8153977b56`, where
the splitter contract sent four native VC transfers that were not previously
visible in explorer internal transactions.

## Verification

Live verification after deployment:

- `web3_clientVersion`:
  `go-opera/v2.0.0-rc.1-60e43d6f-1779011622/linux-amd64/go1.19.13`
- `rpc_modules` includes `"trace": "1.0"`
- `trace_block 0xcf1ae9` returns the target transaction and four child native
  transfers from `0xda90cd8fb13370f5e720b2d6a5cd4b8459817f1f`
- VinuExplorer mainnet backend imported block `13572841` and now serves those
  internal transfers through `/api/v2/transactions/:hash/internal-transactions`

Local build/test checks used for the source branch:

```bash
/usr/local/go/bin/go test ./ethapi ./txtrace ./gossip ./cmd/opera -run '^$'
/usr/local/go/bin/go build -o build/opera-trace ./cmd/opera
git diff --check
```

## Operational Constraints

Trace replay is more expensive than ordinary read-only JSON-RPC calls. Keep
public RPC access behind the existing proxy/rate-limit layer and monitor RPC
latency, CPU, memory, and request volume after enabling any explorer backfill.

Do not enable VinuExplorer's global internal-transaction fetcher while
`pending_block_operations` contains a large historical backlog. Use the backend
range tooling to prune or process bounded ranges first.

## Rollback

Rollback is a binary swap back to:

```text
/home/ubuntu/VinuChain/build/opera.pre-trace-20260517T100059Z
```

Before rollback, record current `web3_clientVersion`, service status, and the
binary SHA256. After rollback, verify:

- `vinu-opera.service` is running
- public RPC responds to `eth_blockNumber`
- `rpc_modules` no longer exposes unexpected namespaces
- VinuExplorer backend health remains `200`

Mainnet RPC swaps require explicit operator confirmation of target host, binary
path, rollback path, and post-swap probes before execution.
