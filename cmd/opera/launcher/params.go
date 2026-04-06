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
			"enode://03c70d4597d731ef182678b7664f2a4a3add07056f23d4e01aba86f066080d18fa13abbd2e13e9d4ea762a2715a983b5ac6151162d05ee0434f1847da1a626e9@34.242.220.16:5050",
			"enode://01c64d1a9dd8a65c56f2d4e373795eb6efd27b714b2b5999363a42a0edc39d7417a431416ceb5c67b1a170983af109e8a15d0c2d44a2ac41ecfb5c23c1a1a48a@3.35.200.210:5050",
			"enode://7044c88daa5df059e2f7a2667471a8149a5cf66e68643dcb86f399d48c4ff6475b73ee91486ea830d225f7f78a2fdf955208673da51c6852230c3a90a3701c06@3.1.103.70:5050",
			"enode://594d26c2338566daca9391d73f1b1821bb0b454e6f3d48715116bf42f320924d569534c143b640feec8a8eaa137a0b822426fb62b52a90162270ea5868bdc37c@18.138.254.181:5050",
			"enode://339e331912e5239a9e13eb82b47be58ea4d3946e91caa2992103a8d4f0226c1e86f9134822d5b238f25c9cbdd473f806caa8e4f8ef1748a6c66395f4bf0dd569@54.66.206.151:5050",
		},
		"test": {
			"enode://e2a95c1b8d85b018b8e88133bec342801b42e19b59a52e030462d04a5549f02fc57215b4ca97771ec6b3a0d30a78603fdccd2b5091c44f6ac439d6c8be8bc539@44.239.129.39:3000",
			"enode://7a45d086b9c82bd3677a76d36e003b9490066d56b612f33d05cb4d242212acd4e5cab4abbcb15a0df9aa499e41b4b4e868d82ba1c509c1990c9217dfe4607775@44.239.129.39:3001",
			"enode://d8e37eeba79b2c52dcba6e396ff907f27a6a8f7db34528cb8636bc3271291657a01c5649bff53429cea8a23b03fac13a178813c34c6d17d14f7b810a988393b5@44.239.129.39:3002",
			"enode://3f15b5ac22dea3e37a90cd9378cf0cd4ed9ea122851846c8108fcc7d2c7e709ea4a089cf3da93c0d3d3053250417cf0ea9ad9eff0aa77ff07d76b6cf267a2937@44.239.129.39:3003",
		},
	}

	vinuTestnetHeader = genesis.Header{
		GenesisID:   hash.HexToHash("0xbf7a3d7f49cd99745acd2aa1c828c81576c41a84fddc9c6ffb9857bab02fe260"),
		NetworkID:   opera.VinuTestNetworkID,
		NetworkName: "VinuChain Testnet",
	}

	vinuTestMainnetHeader = genesis.Header{
		GenesisID:   hash.HexToHash("0xb1b0e08cb0d53d0fb1067658c5af0b3d3ff334d574679f5f74eee2b3448394ce"),
		NetworkID:   205,
		NetworkName: "VinuChain Mainnet",
	}

	vinuMainnetHeader = genesis.Header{
		GenesisID:   hash.HexToHash("0xca7941e04fc93391af59a3a87e2ad386312d6b74922deeaa05068b1c08d9caa4"),
		NetworkID:   opera.VinuMainNetworkID,
		NetworkName: "VinuChain Mainnet",
	}

	AllowedOperaGenesis = []GenesisTemplate{

		// Mainnet
		{
			Name:   "VitaInu mainnet without history",
			Header: vinuMainnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0xc6640c1b62156a63c3121fba3aaf755ee88b84935d2ebf1497611e4ee7f09144"),
				genesisstore.BlocksSection(0): hash.HexToHash("0x438be95bb65eee5e23d7f78d39773646f2f21a6b18266b4d73d1e723c55fb94e"),
				genesisstore.EvmSection(0):    hash.HexToHash("0x765b90e4674d426b37d05a0e4a35addb3deec44a0cc948391b738e0e815682be"),
			},
		},

		// Mainnet with deployed contracts
		{
			Name:   "VitaInu mainnet with deployed contracts",
			Header: vinuMainnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0x482f104dc843b2f86265a3494b1047c65a8568b0578ef1c43ea9aa8c961e6a6f"),
				genesisstore.BlocksSection(0): hash.HexToHash("0x9aab452d91d99fe26457feac40c2be7f2b31facf8edf66d815e2b0a184b871de"),
				genesisstore.EvmSection(0):    hash.HexToHash("0x53f30bbcc37b7ba4d705aad4e79b1e1007673d64b6a1ab703e2319776a62bb3d"),
			},
		},

		// Vita Inu testnet
		{
			Name:   "VitaInu testnet without history",
			Header: vinuTestnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0x7a74f234769f2285be94ac48c4a97abf98de32b93a05ddfd6cc934027b6d2d4f"),
				genesisstore.BlocksSection(0): hash.HexToHash("0xbfe43b2d77e7d672c4b0130d0a43f0710704f53ebb2c39a379c93076a43bddce"),
				genesisstore.EvmSection(0):    hash.HexToHash("0x7c3476d667f7912172df77a5e5804428380541bf98282442689e6b442d16da34"),
			},
		},

		// Vita Inu test mainnet
		{
			Name:   "VitaInu test mainnet without history",
			Header: vinuTestMainnetHeader,
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
