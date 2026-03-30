# VinuChain

EVM-compatible blockchain secured by the Lachesis asynchronous DAG-based BFT consensus algorithm. Forked from [Fantom's go-opera](https://github.com/Fantom-foundation/go-opera) with VinuChain-specific features including a fee refund (payback) system, SFC V2 staking contract, and extensive security hardening.

## Networks

| Network | ID | Hex |
|---------|----|-----|
| Mainnet | 207 | `0xcf` |
| Testnet | 206 | `0xce` |
| Fakenet | 27 | `0x1b` |

## Architecture

VinuChain is built on three repositories:

| Repository | Purpose |
|------------|---------|
| [VinuChain](https://github.com/VinuChain/VinuChain) | Node binary: consensus integration, block processing, gossip, RPC, payback system |
| [go-vinu](https://github.com/VinuChain/go-vinu) | Forked go-ethereum: EVM execution, P2P, accounts, FeeRefund receipt field |
| [lachesis-base](https://github.com/VinuChain/lachesis-base) | Forked Lachesis consensus: DAG-BFT, event processing, stream hardening |

Key packages in this repository:

| Package | Purpose |
|---------|---------|
| `gossip/` | Main node service: P2P, event/block processing, tx pool, emitter |
| `evmcore/` | EVM state transitions and block execution |
| `opera/` | Chain rules, network configs, upgrade flags, on-chain contract bindings |
| `payback/` | Fee refund cache: per-address refunds based on staking activity |
| `inter/` | Core data types: events, blocks, timestamps |
| `vecmt/` | Vector clocks for DAG consensus (median time, cheater detection) |
| `integration/` | Genesis creation and store wiring |

### Upgrade system

Hard forks are gated by boolean flags in `opera.Upgrades`. Each flag enables consensus-critical behavior changes and is activated by shipping a new binary with the flag set in the network's hardcoded rule constructor.

| Flag | What it enables |
|------|----------------|
| `Berlin` | EIP-2565, EIP-2929, EIP-2718, EIP-2930 |
| `London` | EIP-1559 base fee, EIP-3529 |
| `Llr` | Lightweight Lachesis revision |
| `Podgorica` | Fee refund / payback system |
| `SfcV2` | SFC V2 contract upgrade, 30% base fee burn |
| `Elemont` | Consensus fixes: NoCheaters merged view, epoch boundary hardening |

## Building the source

Requires Go 1.22+ and a C compiler.

```shell
make opera
```

The build output is `build/opera`.

## Running `opera`

### Launching a network

You will need a genesis file to join a network, which may be found at https://drive.google.com/drive/folders/1_LKq9ljXYwH4LkxO-6e8UnVgiWncGCBH?usp=sharing

Launching a readonly (non-validator) node:

```shell
$ opera --genesis file.g
```

### Configuration

```shell
$ opera --config /path/to/your_config.toml
```

Export your current configuration:

```shell
$ opera --your-favourite-flags dumpconfig
```

### Validator

Create a new validator private key:

```shell
$ opera validator new
```

Launch a validator:

```shell
$ opera --nousb --validator.id YOUR_ID --validator.pubkey 0xYOUR_PUBKEY
```

`opera` will prompt for a password to decrypt the validator private key. Optionally, specify a password file with `--validator.password`.

### Participation in discovery

Specify your public IP for better connectivity. Ensure TCP/UDP p2p port (5050 by default) is open.

```shell
$ opera --nat extip:1.2.3.4
```

## Dev

### Running testnet

The network is specified by its genesis file:

```shell
$ opera --genesis /path/to/testnet.g
```

Use a separate datadir for testnet to avoid collisions:

```shell
$ opera --genesis /path/to/testnet.g --datadir /path/to/datadir
$ opera --datadir /path/to/datadir account new
$ opera --datadir /path/to/datadir attach
```

### Testing

```shell
make test               # run all tests
go test ./...           # equivalent
go test -race ./...     # with race detector
make coverage           # coverage report
```

### Fakenet (private testing)

Fakenet generates a genesis with N validators with equal stakes. Validator private keys are deterministic — use only for testing.

Single validator (PoA-like):

```shell
$ opera --fakenet 1/1
```

Five validators (run per validator):

```shell
$ opera --fakenet 1/5   # first validator
$ opera --fakenet 2/5   # second validator
```

Non-validator node:

```shell
$ opera --fakenet 0/5
```

Connect nodes via bootnode:

```shell
$ opera --fakenet 1/5 --bootnodes "enode://verylonghex@1.2.3.4:5050"
```

### Demo

```shell
cd demo/
./start.sh    # start VinuChain processes
./stop.sh     # stop
./clean.sh    # erase chain data
```

## License

GNU Lesser General Public License v3.0 — see [COPYING.LESSER](COPYING.LESSER) and [COPYING](COPYING).
