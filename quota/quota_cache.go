package quota

import (
	"fmt"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/quota/contract/quotaProxy"
	"github.com/Fantom-foundation/go-opera/quota/contract/sfc"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
)

type TxType string

const (
	TxTypeStake   TxType = "stake"
	TxTypeUnstake TxType = "unstake"
	TxTypeNone    TxType = "none" // for all other transactions
)

type TxInfo struct {
	Tx      *types.Transaction
	Receipt *types.Receipt
	Type    TxType
}

type BlockInfo struct {
	BlockNumber   uint64
	Txs           []TxInfo
	BaseFeePerGas *big.Int
}

type CircularBuffer struct {
	Size         uint64
	CurrentIndex uint64
	Buffer       []BlockInfo
}

func NewCircularBuffer(size uint64) *CircularBuffer {
	return &CircularBuffer{
		Size:         size,
		CurrentIndex: 0,
		Buffer:       make([]BlockInfo, size),
	}
}

type QuotaCache struct {
	// BlockBuffer contains last window blocks
	// Buffer[CurrentIndex] is the latest block
	// Buffer[CurrentIndex-1] is the previous block
	// Buffer[CurrentIndex+1] is the oldest block
	BlockBuffer *CircularBuffer

	// TxCountMap contains last tx count
	TxCountMap map[common.Address]int64

	// QuotaUsedMap contains last quota used
	QuotaUsedMap map[common.Address]*big.Int

	// StakesMap contains last stakes and unstakes
	// TODO: add stake and unstake events to cache (not one number but []TxInfo)
	StakesMap map[common.Address]*big.Int

	// ABI to get nemes of called methods
	ContractABI *abi.ABI

	store Store

	evm *vm.EVM

	contractAddress common.Address
}

func (qc *QuotaCache) deleteCurrentBlock() error {
	for _, tx := range qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].Txs {
		isEmpty := false
		if _, ok := qc.TxCountMap[tx.Tx.From()]; ok {
			qc.TxCountMap[tx.Tx.From()]--
		} else {
			return fmt.Errorf("consistency error: tx count map does not contain tx")
		}
		if qc.TxCountMap[tx.Tx.From()] == 0 {
			isEmpty = true
			delete(qc.TxCountMap, tx.Tx.From())
		}

		if tx.Type == TxTypeStake {
			if _, ok := qc.StakesMap[tx.Tx.From()]; ok {
				qc.StakesMap[tx.Tx.From()].Sub(qc.StakesMap[tx.Tx.From()], tx.Tx.Value())
			} else {
				return fmt.Errorf("consistency error: stakes map does not contain tx")
			}
			if qc.StakesMap[tx.Tx.From()].Cmp(big.NewInt(0)) == 0 {
				delete(qc.StakesMap, tx.Tx.From())
			} else {
				if isEmpty {
					return fmt.Errorf("consistency error: stakes map is not empty")
				}
			}
		}

		if tx.Type == TxTypeUnstake {
			if _, ok := qc.StakesMap[tx.Tx.From()]; ok {
				qc.StakesMap[tx.Tx.From()].Add(qc.StakesMap[tx.Tx.From()], tx.Tx.Value())
			} else {
				return fmt.Errorf("consistency error: stakes map does not contain tx")
			}
			if qc.StakesMap[tx.Tx.From()].Cmp(big.NewInt(0)) == 0 {
				delete(qc.StakesMap, tx.Tx.From())
			} else {
				if isEmpty {
					return fmt.Errorf("consistency error: stakes map is not empty")
				}
			}
		}

		if _, ok := qc.QuotaUsedMap[tx.Tx.From()]; ok {
			qc.QuotaUsedMap[tx.Tx.From()].Sub(qc.QuotaUsedMap[tx.Tx.From()], tx.Receipt.FeeRefund)
		}

		//if qc.QuotaUsedMap[tx.Tx.From()].Cmp(big.NewInt(0)) == 0 {
		//	delete(qc.QuotaUsedMap, tx.Tx.From())
		//} else {
		//	if isEmpty {
		//		return fmt.Errorf("consistency error: quota used map is not empty")
		//	}
		//}

		//if qc.QuotaUsedMap[tx.Tx.From()].Cmp(big.NewInt(0)) < 0 {
		//	return fmt.Errorf("consistency error: quota used map is negative")
		//}
	}
	return nil
}

func (qc *QuotaCache) AddTransaction(tx *types.Transaction, receipt *types.Receipt) error {
	if receipt.Status == types.ReceiptStatusSuccessful {
		// add tx to block buffer
		// check if this transaction for current block or next block
		if receipt.BlockNumber.Uint64() == qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].BlockNumber {
		} else if receipt.BlockNumber.Uint64() == qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].BlockNumber+1 {
			qc.BlockBuffer.CurrentIndex = (qc.BlockBuffer.CurrentIndex + 1) % qc.BlockBuffer.Size

			// change maps related to oldest block deliting from buffer
			if err := qc.deleteCurrentBlock(); err != nil {
				return err
			}
			qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].BlockNumber = receipt.BlockNumber.Uint64()
			qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].Txs = make([]TxInfo, 0, 1)
		} else if receipt.BlockNumber.Uint64() > qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].BlockNumber {
			return fmt.Errorf(
				"consistency error: receipt block number is greater than current block number, "+
					"receipt block number: %v, "+
					"current block number: %v",
				receipt.BlockNumber,
				qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].BlockNumber,
			)
		} else {
			return fmt.Errorf("consistency error: receipt block number is not current or next, receipt block number: %v, current block number: %v", receipt.BlockNumber, qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].BlockNumber)
		}

		qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].Txs = append(qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].Txs, TxInfo{tx, receipt, TxTypeNone})

		if _, ok := qc.TxCountMap[tx.From()]; !ok {
			qc.TxCountMap[tx.From()] = 0
		}
		qc.TxCountMap[tx.From()]++

		txtype := getTxType(tx, *qc.ContractABI)
		if txtype == TxTypeStake || txtype == TxTypeUnstake {
			if _, ok := qc.StakesMap[tx.From()]; !ok {
				qc.StakesMap[tx.From()] = big.NewInt(0)
			}
			switch txtype {
			case TxTypeStake:
				qc.StakesMap[tx.From()].Add(qc.StakesMap[tx.From()], tx.Value())
			case TxTypeUnstake:
				qc.StakesMap[tx.From()].Sub(qc.StakesMap[tx.From()], tx.Value())
			}
		}

		// add quota used to quota used map
		// check if tx is delegate or undelegate
		if tx.Data() != nil && len(tx.Data()) >= 4 {
			if method, err := qc.ContractABI.MethodById(tx.Data()[:4]); err == nil {
				if method.Name == "delegate" {
					if _, ok := qc.StakesMap[tx.From()]; !ok {
						qc.StakesMap[tx.From()] = big.NewInt(0)
					}
					qc.StakesMap[tx.From()].Add(qc.StakesMap[tx.From()], tx.Value())
					qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].Txs[len(qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].Txs)-1].Type = TxTypeStake
				}

				if method.Name == "undelegate" {
					if _, ok := qc.StakesMap[tx.From()]; !ok {
						qc.StakesMap[tx.From()] = big.NewInt(0)
					}
					qc.StakesMap[tx.From()].Sub(qc.StakesMap[tx.From()], tx.Value())
					qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].Txs[len(qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].Txs)-1].Type = TxTypeUnstake

					if qc.StakesMap[tx.From()].Cmp(big.NewInt(0)) == 0 {
						delete(qc.StakesMap, tx.From())
					}
				}

			}
		}

		if _, ok := qc.QuotaUsedMap[tx.From()]; !ok {
			qc.QuotaUsedMap[tx.From()] = big.NewInt(0)
		}
		qc.QuotaUsedMap[tx.From()].Add(qc.QuotaUsedMap[tx.From()], receipt.FeeRefund)

	}
	return nil
}

func (qc *QuotaCache) GetQuotaUsed(address common.Address) *big.Int {
	if _, ok := qc.QuotaUsedMap[address]; ok {
		return qc.QuotaUsedMap[address]
	}
	return big.NewInt(0)
}

func (qc *QuotaCache) GetLastStakes(address common.Address) *big.Int {
	if _, ok := qc.StakesMap[address]; ok {
		return qc.StakesMap[address]
	}
	return big.NewInt(0)
}

func (qc *QuotaCache) GetTxCount(address common.Address) int64 {
	if _, ok := qc.TxCountMap[address]; ok {
		return qc.TxCountMap[address]
	}
	return 0
}

func (qc *QuotaCache) AddEmptyBlock(blockNumber uint64) error {
	if qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].BlockNumber == blockNumber {
		return nil
	}

	qc.BlockBuffer.CurrentIndex = (qc.BlockBuffer.CurrentIndex + 1) % qc.BlockBuffer.Size

	// change maps related to oldest block deliting from buffer
	if err := qc.deleteCurrentBlock(); err != nil {
		return err
	}

	qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].BlockNumber = blockNumber
	qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].Txs = []TxInfo{}
	return nil
}

func NewQuotaCache(store Store, window uint64) *QuotaCache {
	qc := QuotaCache{
		BlockBuffer:  NewCircularBuffer(window),
		TxCountMap:   make(map[common.Address]int64),
		QuotaUsedMap: make(map[common.Address]*big.Int),
		StakesMap:    make(map[common.Address]*big.Int),
		store:        store,
	}

	qc.contractAddress = store.GetRules().Economy.QuotaCacheAddress
	log.Info("NewQuotaCache: QuotaCacheAddress is set", "address", qc.contractAddress.String())

	abiQuotaProxy, _ := abi.JSON(strings.NewReader(quotaProxy.QuotaProxyMetaData.ABI))
	qc.ContractABI = &abiQuotaProxy

	abiSFC, err := abi.JSON(strings.NewReader(sfc.ContractABI))
	if err != nil {
		panic(err)
	}

	lastBlockIndex := store.GetLatestBlockIndex()

	// if there are less blocks than window size
	if lastBlockIndex < window {
		window = lastBlockIndex
	}

	oldestBlockIndex := lastBlockIndex - window + 1
	for i := oldestBlockIndex; i <= lastBlockIndex; i++ {
		k := i - oldestBlockIndex

		txs, receipts := store.GetBlockTransactionsAndReceipts(i)

		if len(txs) != len(receipts) {
			panic("txs and receipts length mismatch")
		}

		qc.BlockBuffer.Buffer[k].BlockNumber = uint64(i)

		if k > 0 {
			if qc.BlockBuffer.Buffer[k].BaseFeePerGas == nil {
				blockIdx := idx.Block(k)
				epochIdx := store.FindBlockEpoch(blockIdx)

				qc.BlockBuffer.Buffer[k].BaseFeePerGas = store.GetHistoryEpochState(epochIdx).Rules.Economy.MinGasPrice
			}
		}

		for j := 0; j < len(txs); j++ {
			tx := txs[j]
			receipt := receipts[j]
			if receipt.Status == types.ReceiptStatusSuccessful {
				txtype := getTxType(tx, abiSFC)
				qc.BlockBuffer.Buffer[k].Txs = append(qc.BlockBuffer.Buffer[k].Txs, TxInfo{tx, receipt, txtype})
				if _, ok := qc.TxCountMap[tx.From()]; !ok {
					qc.TxCountMap[tx.From()] = 0
				}
				qc.TxCountMap[tx.From()]++

				if _, ok := qc.QuotaUsedMap[tx.From()]; !ok {
					qc.QuotaUsedMap[tx.From()] = big.NewInt(0)
				}
				qc.QuotaUsedMap[tx.From()].Add(qc.QuotaUsedMap[tx.From()], receipt.FeeRefund)

				if txtype == TxTypeStake || txtype == TxTypeUnstake {
					if _, ok := qc.StakesMap[tx.From()]; !ok {
						qc.StakesMap[tx.From()] = big.NewInt(0)
					}
					switch txtype {
					case TxTypeStake:
						qc.StakesMap[tx.From()].Add(qc.StakesMap[tx.From()], tx.Value())
					case TxTypeUnstake:
						qc.StakesMap[tx.From()].Sub(qc.StakesMap[tx.From()], tx.Value())
					}
				}
			}

		}

	}

	return &qc
}

// String() prints all data in cache
func (qc *QuotaCache) String() string {
	var sb strings.Builder
	for i := (qc.BlockBuffer.CurrentIndex + 1) % qc.BlockBuffer.Size; i != qc.BlockBuffer.CurrentIndex; i = (i + 1) % qc.BlockBuffer.Size {
		sb.WriteString(fmt.Sprintf("Block: %d\n", qc.BlockBuffer.Buffer[i].BlockNumber))
		for _, tx := range qc.BlockBuffer.Buffer[i].Txs {
			sb.WriteString(fmt.Sprintf("Txs: hash: %v, from: %v, to: %v, value: %v, type: %v\n", tx.Tx.Hash().Hex(), tx.Tx.From().Hex(), tx.Tx.To().Hex(), tx.Tx.Value().String(), tx.Type))
		}
		sb.WriteString(fmt.Sprintf("BaseFeePerGas: %v\n", qc.BlockBuffer.Buffer[i].BaseFeePerGas))
	}
	sb.WriteString(fmt.Sprintf("TxCountMap: %v\n", qc.TxCountMap))
	sb.WriteString(fmt.Sprintf("QuotaUsedMap: %v\n", qc.QuotaUsedMap))
	sb.WriteString(fmt.Sprintf("StakesMap: %v\n", qc.StakesMap))
	return sb.String()
}

func getTxType(tx *types.Transaction, abi abi.ABI) TxType {
	if tx.Data() != nil && len(tx.Data()) >= 4 {
		if method, err := abi.MethodById(tx.Data()[:4]); err == nil {
			if method.Name == "delegate" {
				return TxTypeStake
			}

			if method.Name == "undelegate" {
				return TxTypeUnstake
			}
		}
	}
	return TxTypeNone
}

func (qc *QuotaCache) GetAvailableQuotaByAddress(address common.Address) *big.Int {
	quota := big.NewInt(0)

	if qc.contractAddress == (common.Address{}) {
		log.Info("GetAvailableQuotaByAddress: QuotaCacheAddress is not set")
		if qc.store != nil {
			if qc.store.GetRules() != (opera.Rules{}) {
				log.Info("GetAvailableQuotaByAddress: QuotaCacheAddress is not set, but rules are set", "rules", qc.store.GetRules())
				if qc.store.GetRules().Economy.QuotaCacheAddress != (common.Address{}) {
					qc.contractAddress = qc.store.GetRules().Economy.QuotaCacheAddress
					log.Info("GetAvailableQuotaByAddress: QuotaCacheAddress is set", "address", qc.contractAddress.String())
				}
			} else {
				return quota
			}
		} else {
			log.Info("GetAvailableQuotaByAddress: QuotaCacheAddress is not set and store is not set")
			return quota
		}
	}

	addressTotalStake, err := qc.getAddressTotalStake(address)
	if err != nil {
		panic(err)
	}

	log.Info("address total stake", "address", address, "total stake", addressTotalStake)

	minStake, err := qc.getMinStake(address)
	if err != nil {
		panic(err)
	}

	log.Info("min stake", "min stake", minStake)

	countBlocksInWindow, err := qc.countBlocksInWindow(address)
	if err != nil {
		panic(err)
	}

	log.Info("count blocks in window", "count blocks in window", countBlocksInWindow)

	if addressTotalStake.Cmp(minStake) < 0 {
		return big.NewInt(0)
	}

	// addressTotalStake * baseFeePerGas
	quota = quota.Mul(addressTotalStake, qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].BaseFeePerGas)

	// addressTotalStake * baseFeePerGas * 21000
	quota = quota.Mul(quota, big.NewInt(21000))

	// addressTotalStake * baseFeePerGas * 21000 / countBlocksInWindow
	quota = quota.Div(quota, countBlocksInWindow)

	// addressTotalStake * baseFeePerGas * 21000 / countBlocksInWindow / minStake
	quota = quota.Div(quota, minStake)

	// subtract already used quota
	quota = quota.Sub(quota, qc.GetQuotaUsed(address))

	if quota.Cmp(big.NewInt(0)) < 0 {
		return big.NewInt(0)
	}

	return quota
}

func (qc *QuotaCache) getAddressTotalStake(address common.Address) (*big.Int, error) {
	var (
		result []byte
		vmerr  error
	)

	sender := vm.AccountRef(address)
	functionSignature := []byte("addressTotalStake(sender)")
	hash := crypto.Keccak256Hash(functionSignature)
	methodID := hash[:4]

	//addressParam := common.LeftPadBytes(address.Bytes(), 32)
	packedData, err := qc.ContractABI.Pack("addressTotalStake", address)
	if err != nil {
		return big.NewInt(0), err
	}

	data := append(methodID, packedData...)

	//gas, err := evmcore.IntrinsicGas(methodID, nil, false)
	//if err != nil {
	//	return big.NewInt(0), err
	//}

	result, _, vmerr = qc.evm.StaticCall(sender, qc.contractAddress, data, 21000)
	if vmerr != nil {
		return big.NewInt(0), vmerr
	}

	resultValue := big.NewInt(0)
	resultValue.SetBytes(result)

	return resultValue, nil
}

func (qc *QuotaCache) getMinStake(address common.Address) (*big.Int, error) {
	var (
		result []byte
		vmerr  error
	)

	sender := vm.AccountRef(address)
	functionSignature := []byte("getMinStake()")
	hash := crypto.Keccak256Hash(functionSignature)
	methodID := hash[:4]

	result, _, vmerr = qc.evm.StaticCall(sender, qc.contractAddress, methodID, 21000)
	if vmerr != nil {
		return big.NewInt(0), vmerr
	}

	resultValue := big.NewInt(0)
	resultValue.SetBytes(result)

	return resultValue, nil
}

func (qc *QuotaCache) GetStore() Store {
	return qc.store
}

func (qc *QuotaCache) AddBaseFeePerGas(blockNumber uint64, baseFeePerGas *big.Int) {
	if baseFeePerGas == nil {
		blockIdx := idx.Block(blockNumber)
		epochIdx := qc.store.FindBlockEpoch(blockIdx)

		qc.BlockBuffer.Buffer[blockNumber].BaseFeePerGas = qc.store.GetHistoryEpochState(epochIdx).Rules.Economy.MinGasPrice
	} else {
		qc.BlockBuffer.Buffer[blockNumber].BaseFeePerGas = baseFeePerGas
	}
}

func (qc *QuotaCache) countBlocksInWindow(address common.Address) (*big.Int, error) {
	var (
		result []byte
		vmerr  error
	)

	sender := vm.AccountRef(address)
	functionSignature := []byte("getFeeRefundBlockCount()")
	hash := crypto.Keccak256Hash(functionSignature)
	methodID := hash[:4]

	result, _, vmerr = qc.evm.StaticCall(sender, qc.contractAddress, methodID, 21000)
	if vmerr != nil {
		return big.NewInt(0), vmerr
	}

	resultValue := big.NewInt(0)
	resultValue.SetBytes(result)

	return resultValue, nil
}

func (qc *QuotaCache) SetEVM(evm *vm.EVM) {
	qc.evm = evm
}
