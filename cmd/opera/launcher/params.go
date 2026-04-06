package launcher

import (
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/ethereum/go-ethereum/params"

	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/genesis"
	"github.com/Fantom-foundation/go-opera/opera/genesisstore"
)

var (
	Bootnodes = map[string][]string{
		"main": {
			"enode://0281626c7d7fc8696300688cbb19f3781aabd981d74cd16f3f5cd7885a32da4d1d9d64afbb2416b93654935a3088afbe1a4a05d823ff2146e5d1d0c2cbdeca46@188.165.195.122:3000",
		},
		"test": {
			"enode://e2a95c1b8d85b018b8e88133bec342801b42e19b59a52e030462d04a5549f02fc57215b4ca97771ec6b3a0d30a78603fdccd2b5091c44f6ac439d6c8be8bc539@44.239.129.39:3000",
			"enode://7a45d086b9c82bd3677a76d36e003b9490066d56b612f33d05cb4d242212acd4e5cab4abbcb15a0df9aa499e41b4b4e868d82ba1c509c1990c9217dfe4607775@44.239.129.39:3001",
			"enode://d8e37eeba79b2c52dcba6e396ff907f27a6a8f7db34528cb8636bc3271291657a01c5649bff53429cea8a23b03fac13a178813c34c6d17d14f7b810a988393b5@44.239.129.39:3002",
			"enode://3f15b5ac22dea3e37a90cd9378cf0cd4ed9ea122851846c8108fcc7d2c7e709ea4a089cf3da93c0d3d3053250417cf0ea9ad9eff0aa77ff07d76b6cf267a2937@44.239.129.39:3003",
		},
	}

	vinuChainTestnetHeader = genesis.Header{
		GenesisID:   hash.HexToHash("0xbf7a3d7f49cd99745acd2aa1c828c81576c41a84fddc9c6ffb9857bab02fe260"),
		NetworkID:   opera.VinuChainTestNetworkID,
		NetworkName: "VinuChain Testnet",
	}

	vinuChainTestMainnetHeader = genesis.Header{
		GenesisID:   hash.HexToHash("0xb1b0e08cb0d53d0fb1067658c5af0b3d3ff334d574679f5f74eee2b3448394ce"),
		NetworkID:   opera.VinuChainStagingNetworkID,
		NetworkName: "VinuChain Staging Mainnet",
	}

	vinuChainMainnetHeader = genesis.Header{
		GenesisID:   hash.HexToHash("0xca7941e04fc93391af59a3a87e2ad386312d6b74922deeaa05068b1c08d9caa4"),
		NetworkID:   opera.VinuChainMainNetworkID,
		NetworkName: "VinuChain Mainnet",
	}

	AllowedOperaGenesis = []GenesisTemplate{

		// Mainnet
		{
			Name:   "VinuChain mainnet without history",
			Header: vinuChainMainnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0xc6640c1b62156a63c3121fba3aaf755ee88b84935d2ebf1497611e4ee7f09144"),
				genesisstore.BlocksSection(0): hash.HexToHash("0x438be95bb65eee5e23d7f78d39773646f2f21a6b18266b4d73d1e723c55fb94e"),
				genesisstore.EvmSection(0):    hash.HexToHash("0x765b90e4674d426b37d05a0e4a35addb3deec44a0cc948391b738e0e815682be"),
			},
		},

		// Mainnet with deployed contracts
		{
			Name:   "VinuChain mainnet with deployed contracts",
			Header: vinuChainMainnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0x482f104dc843b2f86265a3494b1047c65a8568b0578ef1c43ea9aa8c961e6a6f"),
				genesisstore.BlocksSection(0): hash.HexToHash("0x9aab452d91d99fe26457feac40c2be7f2b31facf8edf66d815e2b0a184b871de"),
				genesisstore.EvmSection(0):    hash.HexToHash("0x53f30bbcc37b7ba4d705aad4e79b1e1007673d64b6a1ab703e2319776a62bb3d"),
			},
		},

		// VinuChain testnet
		{
			Name:   "VinuChain testnet without history",
			Header: vinuChainTestnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0x7a74f234769f2285be94ac48c4a97abf98de32b93a05ddfd6cc934027b6d2d4f"),
				genesisstore.BlocksSection(0): hash.HexToHash("0xbfe43b2d77e7d672c4b0130d0a43f0710704f53ebb2c39a379c93076a43bddce"),
				genesisstore.EvmSection(0):    hash.HexToHash("0x7c3476d667f7912172df77a5e5804428380541bf98282442689e6b442d16da34"),
			},
		},

		// VinuChain test mainnet
		{
			Name:   "VinuChain test mainnet without history",
			Header: vinuChainTestMainnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0x4d9b7946e4c2afba6d015e4a4282dd8d2299c1494f13ad6908846e4f09ed11be"),
				genesisstore.BlocksSection(0): hash.HexToHash("0xf51e8759171b4109bbd5d592d62a93d9d882cc0d2846323d20df0ef7b7cd27db"),
				genesisstore.EvmSection(0):    hash.HexToHash("0xe5f319e1c7c064c3f8f2a8226e1481b0102447478089e0602d40eda08055f893"),
			},
		},
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
