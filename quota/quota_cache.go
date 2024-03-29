package quota

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/quota/contract/quotaProxy"
	"github.com/Fantom-foundation/go-opera/quota/contract/sfc"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/log"

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
	TxTypeNone    TxType = "none" // for all other transactions        = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"addressTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeRefundBlockCount\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
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

		if _, ok := qc.QuotaUsedMap[tx.Tx.From()]; ok && tx.Receipt.FeeRefund != nil {
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

		if receipt.FeeRefund != nil {
			if _, ok := qc.QuotaUsedMap[tx.From()]; !ok {
				qc.QuotaUsedMap[tx.From()] = big.NewInt(0)
			}
			qc.QuotaUsedMap[tx.From()].Add(qc.QuotaUsedMap[tx.From()], receipt.FeeRefund)
		}

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

func NewQuotaCache(store Store) *QuotaCache {
	var window uint64

	qc := QuotaCache{
		BlockBuffer:  NewCircularBuffer(1),
		TxCountMap:   make(map[common.Address]int64),
		QuotaUsedMap: make(map[common.Address]*big.Int),
		StakesMap:    make(map[common.Address]*big.Int),
		store:        store,
	}

	if store.GetRules().Upgrades.Podgorica {
		qc.contractAddress = store.GetRules().Economy.QuotaCacheAddress
		log.Info("NewQuotaCache: QuotaCacheAddress is set", "address", qc.contractAddress.String())
	} else {
		log.Info("HardFork Podgorica is not activated", "status", store.GetRules().Upgrades.Podgorica)
		log.Info("NewQuotaCache:", "Rules", store.GetRules())
	}

	abiQuotaProxy, err := abi.JSON(strings.NewReader(quotaProxy.QuotaProxyABI))
	if err != nil {
		panic(err)
	}
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
				epochIdx := store.FindBlockEpoch(idx.Block(qc.BlockBuffer.Buffer[k].BlockNumber))

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
		//for _, tx := range qc.BlockBuffer.Buffer[i].Txs {
		//	sb.WriteString(fmt.Sprintf("Txs: hash: %v, from: %v, to: %v, value: %v, type: %v\n", tx.Tx.Hash().Hex(), tx.Tx.From().Hex(), tx.Tx.To().Hex(), tx.Tx.Value().String(), tx.Type))
		//}
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
		if qc.store != nil {
			if qc.store.GetRules() != (opera.Rules{}) {
				log.Info("GetAvailableQuotaByAddress: HardFork Podgorica", "status", qc.store.GetRules().Upgrades.Podgorica)
			}
		}
	}

	if qc.contractAddress == (common.Address{}) || qc.BlockBuffer.Size == 1 {
		if qc.store != nil {
			if qc.store.GetRules() != (opera.Rules{}) {
				if qc.store.GetRules().Upgrades.Podgorica {
					qc.contractAddress = qc.store.GetRules().Economy.QuotaCacheAddress
					log.Info("GetAvailableQuotaByAddress: QuotaCacheAddress is set", "address", qc.contractAddress.String())

					window, err := qc.countBlocksInWindow(address)
					if err != nil {
						log.Warn("GetAvailableQuotaByAddress:", "error", err)
					}

					qc.InitializeBlockBuffer(window.Uint64())

					log.Info("GetAvailableQuotaByAddress: BlockBuffer initialized", "window", window.Uint64())
				} else {
					return quota
				}
			} else {
				return quota
			}
		} else {
			log.Warn("GetAvailableQuotaByAddress: store is nil")
			return quota
		}
	} else {
		if qc.store.GetRules().Economy.QuotaCacheAddress != qc.contractAddress {
			qc.contractAddress = qc.store.GetRules().Economy.QuotaCacheAddress
			log.Info("GetAvailableQuotaByAddress: QuotaCacheAddress has been changed", "address", qc.contractAddress.String())
		}
	}

	addressTotalStake, err := qc.getAddressTotalStake(address)
	if err != nil {
		log.Warn("GetAvailableQuotaByAddress:", "error", err)
	}

	minStake, err := qc.getMinStake(address)
	if err != nil {
		log.Warn("GetAvailableQuotaByAddress:", "error", err)
	}

	countBlocksInWindow, err := qc.countBlocksInWindow(address)
	if err != nil {
		log.Warn("GetAvailableQuotaByAddress:", "error", err)
	}

	// reinitialize quota cache if countBlocksInWindow is not equal to BlockBuffer.Size
	if countBlocksInWindow.Uint64() != qc.BlockBuffer.Size {
		log.Info("GetAvailableQuotaByAddress: reinitializing quota cache", "countBlocksInWindow", countBlocksInWindow.Uint64(), "BlockBuffer.Size", qc.BlockBuffer.Size)

		currentBlockInfo := qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex]
		qc.InitializeBlockBuffer(countBlocksInWindow.Uint64())
		if currentBlockInfo.BlockNumber < qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].BlockNumber { // should be never true
			log.Warn("Reinitializing quota cache: current block number is less than current block number in buffer", "current block number", currentBlockInfo.BlockNumber, "current index", qc.BlockBuffer.CurrentIndex, "current block number in buffer", qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].BlockNumber)
		} else if currentBlockInfo.BlockNumber == qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].BlockNumber {
			qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex] = currentBlockInfo
		} else {
			if currentBlockInfo.BlockNumber != qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].BlockNumber+1 {
				log.Warn("Reinitializing quota cache: current block number is not current or next", "current block number", currentBlockInfo.BlockNumber, "current index", qc.BlockBuffer.CurrentIndex, "current block number in buffer", qc.BlockBuffer.Buffer[qc.BlockBuffer.CurrentIndex].BlockNumber)
			}
		}
		err = qc.AddEmptyBlock(currentBlockInfo.BlockNumber)
		if err != nil {
			log.Warn("GetAvailableQuotaByAddress AddEmptyBlock:", "error", err)
		}
		for _, tx := range currentBlockInfo.Txs {
			err = qc.AddTransaction(tx.Tx, tx.Receipt)
			if err != nil {
				log.Warn("GetAvailableQuotaByAddress AddTransaction:", "error", err)
			}
		}
	}

	quotaFactor, err := qc.getQuotaFactor(address)
	if err != nil {
		log.Warn("GetAvailableQuotaByAddress:", "error", err)
	}

	if addressTotalStake.Cmp(minStake) < 0 {
		return big.NewInt(0)
	}

	quotaSum := big.NewInt(0)
	for j := 0; j < int(qc.BlockBuffer.Size); j++ {
		i := (qc.BlockBuffer.CurrentIndex - uint64(j) + qc.BlockBuffer.Size) % qc.BlockBuffer.Size
		if qc.BlockBuffer.Buffer[i].BlockNumber < 2 {
			break
		}

		quota = big.NewInt(0)

		if (qc.BlockBuffer.Buffer[i].BaseFeePerGas == nil || qc.BlockBuffer.Buffer[i].BaseFeePerGas.Cmp(big.NewInt(0)) == 0) && qc.BlockBuffer.Buffer[i].BlockNumber != 0 {
			blockIdx := idx.Block(qc.BlockBuffer.Buffer[i].BlockNumber)
			epochIdx := qc.store.FindBlockEpoch(blockIdx)

			qc.BlockBuffer.Buffer[i].BaseFeePerGas = qc.store.GetHistoryEpochState(epochIdx).Rules.Economy.MinGasPrice
		}

		// addressTotalStake * baseFeePerGas
		quota = quota.Mul(addressTotalStake, qc.BlockBuffer.Buffer[i].BaseFeePerGas)

		// addressTotalStake * baseFeePerGas * quotaFactor
		quota = quota.Mul(quota, quotaFactor)

		// addressTotalStake * baseFeePerGas * quotaFactor / countBlocksInWindow
		quota = quota.Div(quota, countBlocksInWindow)

		// addressTotalStake * baseFeePerGas * quotaFactor / countBlocksInWindow / minStake
		quota = quota.Div(quota, minStake)

		quotaSum = quotaSum.Add(quotaSum, quota)
	}

	quotaSum = quotaSum.Sub(quotaSum, qc.GetQuotaUsed(address))

	log.Info("GetAvailableQuotaByAddress:", "quotaSum after subtracting used quota", quotaSum)

	if quotaSum.Cmp(big.NewInt(0)) < 0 {
		log.Warn("GetAvailableQuotaByAddress: quotaSum is negative", "quotaSum", quotaSum)
		return big.NewInt(0)
	}

	return quotaSum
}

func (qc *QuotaCache) getAddressTotalStake(address common.Address) (*big.Int, error) {
	var (
		result []byte
		vmerr  error
	)

	sender := vm.AccountRef(address)
	packedData, err := qc.ContractABI.Pack("getStake", address)
	if err != nil {
		return big.NewInt(0), err
	}

	result, _, vmerr = qc.evm.StaticCall(sender, qc.contractAddress, packedData, 21000)
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
	functionSignature := []byte("minStake()")
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
	functionSignature := []byte("feeRefundBlockCount()")
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

func (qc *QuotaCache) getQuotaFactor(address common.Address) (*big.Int, error) {
	var (
		result []byte
		vmerr  error
	)

	sender := vm.AccountRef(address)
	functionSignature := []byte("quotaFactor()")
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

func (qc *QuotaCache) InitializeBlockBuffer(window uint64) {
	qc.BlockBuffer = NewCircularBuffer(window)

	abiSFC, err := abi.JSON(strings.NewReader(sfc.ContractABI))
	if err != nil {
		panic(err)
	}

	lastBlockIndex := qc.store.GetLatestBlockIndex()

	// if there are less blocks than window size
	if lastBlockIndex < window {
		window = lastBlockIndex
	}

	oldestBlockIndex := lastBlockIndex - window + 1
	for i := oldestBlockIndex; i <= lastBlockIndex; i++ {
		k := i - oldestBlockIndex

		txs, receipts := qc.store.GetBlockTransactionsAndReceipts(i)

		if len(txs) != len(receipts) {
			panic("txs and receipts length mismatch")
		}

		qc.BlockBuffer.Buffer[k].BlockNumber = uint64(i)

		if k > 0 {
			if qc.BlockBuffer.Buffer[k].BaseFeePerGas == nil {
				epochIdx := qc.store.FindBlockEpoch(idx.Block(qc.BlockBuffer.Buffer[k].BlockNumber))

				qc.BlockBuffer.Buffer[k].BaseFeePerGas = qc.store.GetHistoryEpochState(epochIdx).Rules.Economy.MinGasPrice
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

				if receipt.FeeRefund != nil {
					if _, ok := qc.QuotaUsedMap[tx.From()]; !ok {
						qc.QuotaUsedMap[tx.From()] = big.NewInt(0)
					}
					qc.QuotaUsedMap[tx.From()].Add(qc.QuotaUsedMap[tx.From()], receipt.FeeRefund)
				}
			}

		}

	}

	qc.BlockBuffer.CurrentIndex = window - 1

}
