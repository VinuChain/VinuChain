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
			"enode://563b30428f48357f31c9d4906ca2f3d3815d663b151302c1ba9d58f3428265b554398c6fabf4b806a49525670cd9e031257c805375b9fdbcc015f60a7943e427@3.213.142.230:7946",
			"enode://8b53fe4410cde82d98d28697d56ccb793f9a67b1f8807c523eadafe96339d6e56bc82c0e702757ac5010972e966761b1abecb4935d9a86a9feed47e3e9ba27a6@3.227.34.226:7946",
			"enode://1703640d1239434dcaf010541cafeeb3c4c707be9098954c50aa705f6e97e2d0273671df13f6e447563e7d3a7c7ffc88de48318d8a3cc2cc59d196516054f17e@52.72.222.228:7946",
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
