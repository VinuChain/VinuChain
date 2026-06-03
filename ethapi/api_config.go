package ethapi

import (
	"context"
	"fmt"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/forkid"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

type ethConfigResponse struct {
	Current *ethForkConfig `json:"current"`
	Next    *ethForkConfig `json:"next"`
	Last    *ethForkConfig `json:"last"`
}

type ethForkConfig struct {
	ActivationTime  uint64                    `json:"activationTime"`
	ActivationBlock uint64                    `json:"activationBlock"`
	ChainID         string                    `json:"chainId"`
	ForkID          string                    `json:"forkId"`
	Precompiles     map[string]common.Address `json:"precompiles"`
}

// Config implements eth_config. VinuChain is block-height forked, so the
// response includes activationBlock as a Vinu extension; activationTime is only
// knowable after the activation block exists.
func (s *PublicBlockChainAPI) Config(ctx context.Context) (*ethConfigResponse, error) {
	cfg := s.b.ChainConfig()
	if cfg == nil {
		return nil, fmt.Errorf("chain config unavailable")
	}
	head := s.b.CurrentBlock()
	if head == nil {
		return nil, fmt.Errorf("current block unavailable")
	}
	genesis, err := s.b.HeaderByNumber(ctx, rpc.BlockNumber(0))
	if err != nil {
		return nil, err
	}
	if genesis == nil {
		return nil, fmt.Errorf("genesis header unavailable")
	}

	headNum := head.NumberU64()
	forks := ethConfigForkBlocks(cfg)
	currentBlock := forks[0]
	var future []uint64
	for _, block := range forks {
		if block <= headNum {
			currentBlock = block
			continue
		}
		future = append(future, block)
	}

	resp := &ethConfigResponse{}
	resp.Current, err = s.ethForkConfigAt(ctx, cfg, genesis.Hash, currentBlock, headNum)
	if err != nil {
		return nil, err
	}
	if len(future) > 0 {
		resp.Next, err = s.ethForkConfigAt(ctx, cfg, genesis.Hash, future[0], headNum)
		if err != nil {
			return nil, err
		}
		resp.Last, err = s.ethForkConfigAt(ctx, cfg, genesis.Hash, future[len(future)-1], headNum)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (s *PublicBlockChainAPI) ethForkConfigAt(ctx context.Context, cfg *params.ChainConfig, genesis common.Hash, block uint64, head uint64) (*ethForkConfig, error) {
	var activationTime uint64
	if block <= head {
		header, err := s.b.HeaderByNumber(ctx, rpc.BlockNumber(block))
		if err != nil {
			return nil, err
		}
		if header != nil {
			activationTime = uint64(header.Time.Unix())
		}
	}
	id := forkid.NewID(cfg, genesis, block)
	chainID := "0x0"
	if cfg.ChainID != nil {
		chainID = "0x" + cfg.ChainID.Text(16)
	}
	return &ethForkConfig{
		ActivationTime:  activationTime,
		ActivationBlock: block,
		ChainID:         chainID,
		ForkID:          hexutil.Encode(id.Hash[:]),
		Precompiles:     ethConfigPrecompiles(cfg.Rules(new(big.Int).SetUint64(block))),
	}, nil
}

func ethConfigForkBlocks(cfg *params.ChainConfig) []uint64 {
	seen := map[uint64]struct{}{0: {}}
	for _, block := range []*big.Int{
		cfg.HomesteadBlock,
		cfg.EIP150Block,
		cfg.EIP155Block,
		cfg.EIP158Block,
		cfg.ByzantiumBlock,
		cfg.ConstantinopleBlock,
		cfg.PetersburgBlock,
		cfg.IstanbulBlock,
		cfg.BerlinBlock,
		cfg.LondonBlock,
		cfg.ShanghaiBlock,
		cfg.CancunBlock,
		cfg.PragueBlock,
		cfg.VinuBLSBlock,
		cfg.VinuLatestEVMBlock,
	} {
		if block == nil || !block.IsUint64() {
			continue
		}
		seen[block.Uint64()] = struct{}{}
	}
	forks := make([]uint64, 0, len(seen))
	for block := range seen {
		forks = append(forks, block)
	}
	sort.Slice(forks, func(i, j int) bool { return forks[i] < forks[j] })
	return forks
}

func ethConfigPrecompiles(rules params.Rules) map[string]common.Address {
	precompiles := map[string]common.Address{
		"ECREC":     common.BytesToAddress([]byte{0x01}),
		"ID":        common.BytesToAddress([]byte{0x04}),
		"RIPEMD160": common.BytesToAddress([]byte{0x03}),
		"SHA256":    common.BytesToAddress([]byte{0x02}),
	}
	if rules.IsByzantium {
		precompiles["BN254_ADD"] = common.BytesToAddress([]byte{0x06})
		precompiles["BN254_MUL"] = common.BytesToAddress([]byte{0x07})
		precompiles["BN254_PAIRING"] = common.BytesToAddress([]byte{0x08})
		precompiles["MODEXP"] = common.BytesToAddress([]byte{0x05})
	}
	if rules.IsIstanbul || rules.IsBerlin {
		precompiles["BLAKE2F"] = common.BytesToAddress([]byte{0x09})
	}
	if rules.IsVinuBLS || rules.IsVinuLatestEVM {
		precompiles["BLS12_G1ADD"] = common.BytesToAddress([]byte{0x0b})
		precompiles["BLS12_G1MSM"] = common.BytesToAddress([]byte{0x0c})
		precompiles["BLS12_G2ADD"] = common.BytesToAddress([]byte{0x0d})
		precompiles["BLS12_G2MSM"] = common.BytesToAddress([]byte{0x0e})
		precompiles["BLS12_MAP_FP2_TO_G2"] = common.BytesToAddress([]byte{0x11})
		precompiles["BLS12_MAP_FP_TO_G1"] = common.BytesToAddress([]byte{0x10})
		precompiles["BLS12_PAIRING_CHECK"] = common.BytesToAddress([]byte{0x0f})
	}
	if rules.IsVinuLatestEVM {
		precompiles["P256VERIFY"] = common.BytesToAddress([]byte{0x01, 0x00})
	}
	return precompiles
}
