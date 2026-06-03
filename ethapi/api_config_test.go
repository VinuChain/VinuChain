package ethapi

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-opera/evmcore"
	"github.com/Fantom-foundation/go-opera/inter"
)

type configStubBackend struct {
	stubBackend
	config  *params.ChainConfig
	current *evmcore.EvmBlock
	headers map[uint64]*evmcore.EvmHeader
}

func (b *configStubBackend) ChainConfig() *params.ChainConfig {
	return b.config
}

func (b *configStubBackend) CurrentBlock() *evmcore.EvmBlock {
	return b.current
}

func (b *configStubBackend) HeaderByNumber(_ context.Context, number rpc.BlockNumber) (*evmcore.EvmHeader, error) {
	if number == rpc.LatestBlockNumber {
		return &b.current.EvmHeader, nil
	}
	return b.headers[uint64(number)], nil
}

func TestEthConfigReportsCurrentNextAndLastForks(t *testing.T) {
	cfg := *params.TestChainConfig
	cfg.ChainID = big.NewInt(206)
	cfg.HomesteadBlock = common.Big0
	cfg.EIP150Block = common.Big0
	cfg.EIP155Block = common.Big0
	cfg.EIP158Block = common.Big0
	cfg.ByzantiumBlock = common.Big0
	cfg.ConstantinopleBlock = common.Big0
	cfg.PetersburgBlock = common.Big0
	cfg.IstanbulBlock = common.Big0
	cfg.BerlinBlock = common.Big0
	cfg.LondonBlock = common.Big0
	cfg.ShanghaiBlock = common.Big0
	cfg.CancunBlock = common.Big0
	cfg.PragueBlock = common.Big0
	cfg.VinuBLSBlock = big.NewInt(10)
	cfg.VinuLatestEVMBlock = big.NewInt(20)

	backend := &configStubBackend{
		config: &cfg,
		current: evmcore.NewEvmBlock(&evmcore.EvmHeader{
			Number: big.NewInt(12),
			Hash:   common.HexToHash("0x12"),
			Time:   inter.FromUnix(120),
		}, nil),
		headers: map[uint64]*evmcore.EvmHeader{
			0: {
				Number: big.NewInt(0),
				Hash:   common.HexToHash("0x01"),
				Time:   inter.FromUnix(1),
			},
			10: {
				Number: big.NewInt(10),
				Hash:   common.HexToHash("0x10"),
				Time:   inter.FromUnix(100),
			},
			12: {
				Number: big.NewInt(12),
				Hash:   common.HexToHash("0x12"),
				Time:   inter.FromUnix(120),
			},
		},
	}
	api := NewPublicBlockChainAPI(backend)

	resp, err := api.Config(context.Background())
	require.NoError(t, err)
	require.NotNil(t, resp.Current)
	require.NotNil(t, resp.Next)
	require.NotNil(t, resp.Last)

	require.Equal(t, "0xce", resp.Current.ChainID)
	require.Equal(t, uint64(10), resp.Current.ActivationBlock)
	require.Equal(t, uint64(100), resp.Current.ActivationTime)
	require.Contains(t, resp.Current.Precompiles, "BLS12_G1ADD")
	require.NotContains(t, resp.Current.Precompiles, "P256VERIFY")
	require.Regexp(t, "^0x[0-9a-f]{8}$", resp.Current.ForkID)

	require.Equal(t, uint64(20), resp.Next.ActivationBlock)
	require.Zero(t, resp.Next.ActivationTime)
	require.Contains(t, resp.Next.Precompiles, "BLS12_G1ADD")
	require.Equal(t, common.BytesToAddress([]byte{0x01, 0x00}), resp.Next.Precompiles["P256VERIFY"])
	require.Equal(t, resp.Next, resp.Last)
}

func TestEthConfigReportsDistinctNextAndLastFutureForks(t *testing.T) {
	cfg := *params.TestChainConfig
	cfg.ChainID = big.NewInt(206)
	cfg.HomesteadBlock = common.Big0
	cfg.EIP150Block = common.Big0
	cfg.EIP155Block = common.Big0
	cfg.EIP158Block = common.Big0
	cfg.ByzantiumBlock = common.Big0
	cfg.ConstantinopleBlock = common.Big0
	cfg.PetersburgBlock = common.Big0
	cfg.IstanbulBlock = common.Big0
	cfg.BerlinBlock = common.Big0
	cfg.LondonBlock = common.Big0
	cfg.ShanghaiBlock = common.Big0
	cfg.CancunBlock = common.Big0
	cfg.PragueBlock = common.Big0
	cfg.VinuBLSBlock = big.NewInt(10)
	cfg.VinuLatestEVMBlock = big.NewInt(20)

	backend := &configStubBackend{
		config: &cfg,
		current: evmcore.NewEvmBlock(&evmcore.EvmHeader{
			Number: big.NewInt(5),
			Hash:   common.HexToHash("0x05"),
			Time:   inter.FromUnix(50),
		}, nil),
		headers: map[uint64]*evmcore.EvmHeader{
			0: {Number: big.NewInt(0), Hash: common.HexToHash("0x01"), Time: inter.FromUnix(1)},
			5: {Number: big.NewInt(5), Hash: common.HexToHash("0x05"), Time: inter.FromUnix(50)},
		},
	}
	api := NewPublicBlockChainAPI(backend)

	resp, err := api.Config(context.Background())
	require.NoError(t, err)
	require.NotNil(t, resp.Current)
	require.NotNil(t, resp.Next)
	require.NotNil(t, resp.Last)
	require.Equal(t, uint64(0), resp.Current.ActivationBlock)
	require.Equal(t, uint64(10), resp.Next.ActivationBlock)
	require.Equal(t, uint64(20), resp.Last.ActivationBlock)
	require.Contains(t, resp.Next.Precompiles, "BLS12_G1ADD")
	require.NotContains(t, resp.Next.Precompiles, "P256VERIFY")
	require.Equal(t, common.BytesToAddress([]byte{0x01, 0x00}), resp.Last.Precompiles["P256VERIFY"])
	require.NotEqual(t, resp.Next.ForkID, resp.Last.ForkID)
}

func TestEthConfigHasNoNextAfterLastConfiguredFork(t *testing.T) {
	cfg := *params.TestChainConfig
	cfg.ChainID = big.NewInt(206)
	cfg.HomesteadBlock = common.Big0
	cfg.EIP150Block = common.Big0
	cfg.EIP155Block = common.Big0
	cfg.EIP158Block = common.Big0
	cfg.ByzantiumBlock = common.Big0
	cfg.ConstantinopleBlock = common.Big0
	cfg.PetersburgBlock = common.Big0
	cfg.IstanbulBlock = common.Big0
	cfg.BerlinBlock = common.Big0
	cfg.LondonBlock = common.Big0
	cfg.ShanghaiBlock = common.Big0
	cfg.CancunBlock = common.Big0
	cfg.PragueBlock = common.Big0
	cfg.VinuBLSBlock = common.Big0
	cfg.VinuLatestEVMBlock = common.Big0

	header := &evmcore.EvmHeader{
		Number: big.NewInt(25),
		Hash:   common.HexToHash("0x25"),
		Time:   inter.FromUnix(250),
	}
	backend := &configStubBackend{
		config:  &cfg,
		current: evmcore.NewEvmBlock(header, nil),
		headers: map[uint64]*evmcore.EvmHeader{
			0:  {Number: big.NewInt(0), Hash: common.HexToHash("0x01"), Time: inter.FromUnix(1)},
			25: header,
		},
	}
	api := NewPublicBlockChainAPI(backend)

	resp, err := api.Config(context.Background())
	require.NoError(t, err)
	require.NotNil(t, resp.Current)
	require.Nil(t, resp.Next)
	require.Nil(t, resp.Last)
	require.Contains(t, resp.Current.Precompiles, "P256VERIFY")
}
