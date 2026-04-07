package drivermodule

import (
	"io"
	"math"
	"math/big"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"
	ethmath "github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Fantom-foundation/go-opera/gossip/blockproc"
	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/inter/drivertype"
	"github.com/Fantom-foundation/go-opera/inter/iblockproc"
	"github.com/Fantom-foundation/go-opera/inter/validatorpk"
	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/contracts/driver"
	"github.com/Fantom-foundation/go-opera/opera/contracts/driver/drivercall"
	"github.com/Fantom-foundation/go-opera/opera/contracts/driver/driverpos"
	"github.com/Fantom-foundation/go-opera/tracing"
)

const (
	maxAdvanceEpochs       = 1 << 16
	maxValidators          = 500
	baseFeeBurnPercentage  = 30
	baseFeeBurnDenominator = 100
)

// burnNumeratorVal and burnDenominatorVal return fresh big.Int values
// to prevent accidental mutation of consensus-critical constants.
func burnNumeratorVal() *big.Int   { return big.NewInt(baseFeeBurnPercentage) }
func burnDenominatorVal() *big.Int { return big.NewInt(baseFeeBurnDenominator) }

type DriverTxListenerModule struct{}

func NewDriverTxListenerModule() *DriverTxListenerModule {
	return &DriverTxListenerModule{}
}

func (m *DriverTxListenerModule) Start(block iblockproc.BlockCtx, bs iblockproc.BlockState, es iblockproc.EpochState, statedb *state.StateDB) blockproc.TxListener {
	// Snapshot MinGasPrice at block start so fee calculations are
	// consistent for the entire block, even if governance updates
	// DirtyRules mid-block (which wouldn't affect es.Rules today,
	// but this makes the invariant explicit and refactor-proof).
	var minGasPrice *big.Int
	if es.Rules.Economy.MinGasPrice != nil {
		minGasPrice = new(big.Int).Set(es.Rules.Economy.MinGasPrice)
	}
	return &DriverTxListener{
		block:            block,
		es:               es,
		bs:               bs,
		statedb:          statedb,
		blockMinGasPrice: minGasPrice,
	}
}

type DriverTxListener struct {
	block          iblockproc.BlockCtx
	es             iblockproc.EpochState
	bs             iblockproc.BlockState
	statedb        *state.StateDB
	blockMinGasPrice *big.Int // snapshotted at block start; governance DirtyRules cannot change this mid-block
}

type DriverTxTransactor struct{}

type DriverTxPreTransactor struct{}

func NewDriverTxTransactor() *DriverTxTransactor {
	return &DriverTxTransactor{}
}

func NewDriverTxPreTransactor() *DriverTxPreTransactor {
	return &DriverTxPreTransactor{}
}

func InternalTxBuilder(statedb *state.StateDB) func(calldata []byte, addr common.Address) *types.Transaction {
	var nonce uint64
	nonceInitialized := false
	return func(calldata []byte, addr common.Address) *types.Transaction {
		if !nonceInitialized {
			nonce = statedb.GetNonce(common.Address{})
			nonceInitialized = true
		}
		tx := types.NewTransaction(nonce, addr, common.Big0, 1e10, common.Big0, calldata)
		nonce++
		return tx
	}
}

func maxBlockIdx(a, b idx.Block) idx.Block {
	if a > b {
		return a
	}
	return b
}

func (p *DriverTxPreTransactor) PopInternalTxs(block iblockproc.BlockCtx, bs iblockproc.BlockState, es iblockproc.EpochState, sealing bool, statedb *state.StateDB) types.Transactions {
	buildTx := InternalTxBuilder(statedb)
	internalTxs := make(types.Transactions, 0, 8)

	// write cheaters
	for _, validatorID := range bs.EpochCheaters[bs.CheatersWritten:] {
		calldata := drivercall.DeactivateValidator(validatorID, drivertype.DoublesignBit)
		internalTxs = append(internalTxs, buildTx(calldata, driver.ContractAddress))
	}

	// push data into Driver before epoch sealing
	if sealing {
		metrics := make([]drivercall.ValidatorEpochMetric, es.Validators.Len())
		for oldValIdx := idx.Validator(0); oldValIdx < es.Validators.Len(); oldValIdx++ {
			info := bs.ValidatorStates[oldValIdx]
			// forgive downtime if below BlockMissedSlack
			missed := opera.BlocksMissed{
				BlocksNum: maxBlockIdx(block.Idx, info.LastBlock) - info.LastBlock,
				Period:    inter.MaxTimestamp(block.Time, info.LastOnlineTime) - info.LastOnlineTime,
			}
			uptime := info.Uptime
			if missed.BlocksNum <= es.Rules.Economy.BlockMissedSlack {
				missed = opera.BlocksMissed{}
				prevOnlineTime := inter.MaxTimestamp(info.LastOnlineTime, es.EpochStart)
				uptime += inter.MaxTimestamp(block.Time, prevOnlineTime) - prevOnlineTime
			}
			originatedFee := info.Originated
			if es.Rules.Upgrades.Elemont {
				valID := es.Validators.GetID(oldValIdx)
				for _, cheater := range bs.EpochCheaters {
					if cheater == valID {
						originatedFee = new(big.Int)
						break
					}
				}
			}
			metrics[oldValIdx] = drivercall.ValidatorEpochMetric{
				Missed:          missed,
				Uptime:          uptime,
				OriginatedTxFee: originatedFee,
			}
		}
		calldata := drivercall.SealEpoch(metrics)
		internalTxs = append(internalTxs, buildTx(calldata, driver.ContractAddress))
	}
	return internalTxs
}

func (p *DriverTxTransactor) PopInternalTxs(_ iblockproc.BlockCtx, _ iblockproc.BlockState, es iblockproc.EpochState, sealing bool, statedb *state.StateDB) types.Transactions {
	buildTx := InternalTxBuilder(statedb)
	internalTxs := make(types.Transactions, 0, 1)
	// push data into Driver after epoch sealing
	if sealing {
		calldata := drivercall.SealEpochValidators(es.Validators.SortedIDs())
		internalTxs = append(internalTxs, buildTx(calldata, driver.ContractAddress))
	}
	return internalTxs
}

func (p *DriverTxListener) OnNewReceipt(tx *types.Transaction, r *types.Receipt, originator idx.ValidatorID) {
	tracing.FinishTx(tx.Hash(), "DriverTxListener.OnNewReceipt()")
	if originator == 0 {
		return
	}
	originatorIdx := p.es.Validators.GetIdx(originator)

	gasUsed := new(big.Int).SetUint64(r.GasUsed)

	// Compute effective gas price using the block-start MinGasPrice snapshot
	// to ensure txFee matches what refundGas computed via the EVM baseFee.
	effectiveGasPrice := tx.GasPrice()
	if p.es.Rules.Upgrades.London && p.blockMinGasPrice != nil {
		effectiveGasPrice = ethmath.BigMin(
			new(big.Int).Add(tx.GasTipCap(), p.blockMinGasPrice),
			tx.GasFeeCap(),
		)
	}
	txFee := new(big.Int).Mul(gasUsed, effectiveGasPrice)

	// Guard against nil FeeRefund
	feeRefund := r.FeeRefund
	if feeRefund == nil {
		feeRefund = new(big.Int)
	}

	// Calculate validator's share: total fee minus payback refund
	validatorFee := new(big.Int).Sub(txFee, feeRefund)
	if validatorFee.Sign() < 0 {
		validatorFee.SetUint64(0)
	}

	// Burn 30% of base fee portion from the validator's share when SfcV2 is active
	burnAmount := new(big.Int)
	if p.es.Rules.Upgrades.SfcV2 && p.es.Rules.Upgrades.London && p.blockMinGasPrice != nil && p.blockMinGasPrice.Sign() > 0 {
		baseFeeUsed := new(big.Int).Mul(p.blockMinGasPrice, gasUsed)
		burnAmount.Mul(baseFeeUsed, burnNumeratorVal())
		burnAmount.Div(burnAmount, burnDenominatorVal())
		// burn cannot exceed what validators would receive
		if burnAmount.Cmp(validatorFee) > 0 {
			burnAmount.Set(validatorFee)
		}
	}

	originated := p.bs.ValidatorStates[originatorIdx].Originated
	delta := new(big.Int).Sub(validatorFee, burnAmount)
	p.bs.ValidatorStates[originatorIdx].Originated = new(big.Int).Add(originated, delta)

	if feeRefund.Sign() > 0 {
		log.Debug("Payback fee refund", "tx", tx.Hash().Hex(), "fee", txFee, "refund", feeRefund)
	}

	// Send burned amount to 0x0 address for on-chain accounting
	if burnAmount.Sign() > 0 {
		p.statedb.AddBalance(common.Address{}, burnAmount)
		log.Debug("Base fee burned", "tx", tx.Hash().Hex(), "burn", burnAmount, "baseFee", p.blockMinGasPrice, "fee", txFee)
	}

	// track gas power refunds
	notUsedGas := tx.Gas() - r.GasUsed
	if notUsedGas != 0 {
		cur := p.bs.ValidatorStates[originatorIdx].DirtyGasRefund
		if cur+notUsedGas >= cur {
			p.bs.ValidatorStates[originatorIdx].DirtyGasRefund = cur + notUsedGas
		} else {
			p.bs.ValidatorStates[originatorIdx].DirtyGasRefund = math.MaxUint64
		}
	}
}

func decodeDataBytes(l *types.Log) ([]byte, error) {
	if len(l.Data) < 32 {
		return nil, io.ErrUnexpectedEOF
	}
	start := new(big.Int).SetBytes(l.Data[24:32]).Uint64()
	if start > math.MaxUint64-32 {
		return nil, io.ErrUnexpectedEOF
	}
	if start+32 > uint64(len(l.Data)) {
		return nil, io.ErrUnexpectedEOF
	}
	size := new(big.Int).SetBytes(l.Data[start+24 : start+32]).Uint64()
	if size > math.MaxUint64-start-32 {
		return nil, io.ErrUnexpectedEOF
	}
	if start+32+size > uint64(len(l.Data)) {
		return nil, io.ErrUnexpectedEOF
	}
	return l.Data[start+32 : start+32+size], nil
}

func (p *DriverTxListener) OnNewLog(l *types.Log) {
	if l.Address != driver.ContractAddress {
		return
	}
	if len(l.Topics) == 0 {
		return
	}
	// Track validator weight changes
	if l.Topics[0] == driverpos.Topics.UpdateValidatorWeight && len(l.Topics) > 1 && len(l.Data) >= 32 {
		validatorID := idx.ValidatorID(new(big.Int).SetBytes(l.Topics[1][:]).Uint64())
		weight := new(big.Int).SetBytes(l.Data[0:32])

		if weight.Sign() == 0 {
			delete(p.bs.NextValidatorProfiles, validatorID)
		} else {
			profile, ok := p.bs.NextValidatorProfiles[validatorID]
			if !ok {
				if len(p.bs.NextValidatorProfiles) >= maxValidators {
					log.Warn("Validator count at cap, ignoring new validator", "id", validatorID, "cap", maxValidators)
				} else {
					profile.PubKey = validatorpk.PubKey{
						Type: 0,
						Raw:  []byte{},
					}
					profile.Weight = weight
					p.bs.NextValidatorProfiles[validatorID] = profile
				}
			} else {
				profile.Weight = weight
				p.bs.NextValidatorProfiles[validatorID] = profile
			}
		}
	}
	// Track validator pubkey changes
	if l.Topics[0] == driverpos.Topics.UpdateValidatorPubkey && len(l.Topics) > 1 {
		validatorID := idx.ValidatorID(new(big.Int).SetBytes(l.Topics[1][:]).Uint64())
		pubkey, err := decodeDataBytes(l)
		if err != nil {
			log.Warn("Malformed UpdatedValidatorPubkey Driver event")
			return
		}

		profile, ok := p.bs.NextValidatorProfiles[validatorID]
		if !ok {
			log.Warn("Unexpected UpdatedValidatorPubkey Driver event")
			return
		}
		parsedKey, err := validatorpk.FromBytes(pubkey)
		if err != nil {
			log.Warn("Malformed validator pubkey in UpdateValidatorPubkey event", "validator", validatorID, "err", err)
			return
		}
		profile.PubKey = parsedKey
		p.bs.NextValidatorProfiles[validatorID] = profile
	}
	// Update rules
	if l.Topics[0] == driverpos.Topics.UpdateNetworkRules && len(l.Data) >= 64 {
		diff, err := decodeDataBytes(l)
		if err != nil {
			log.Warn("Malformed UpdateNetworkRules Driver event")
			return
		}

		last := &p.es.Rules
		if p.bs.DirtyRules != nil {
			last = p.bs.DirtyRules
		}
		updated, err := opera.UpdateRules(*last, diff)
		if err != nil {
			log.Warn("Network rules update error", "err", err)
			return
		}
		p.bs.DirtyRules = &updated
	}
	// Advance epochs
	if l.Topics[0] == driverpos.Topics.AdvanceEpochs && len(l.Data) >= 32 {
		var epochsNum uint64
		if p.es.Rules.Upgrades.Elemont {
			epochsNum = new(big.Int).SetBytes(l.Data[0:32]).Uint64()
		} else {
			epochsNum = new(big.Int).SetBytes(l.Data[29:32]).Uint64()
		}

		if epochsNum > uint64(maxAdvanceEpochs) {
			epochsNum = maxAdvanceEpochs
		}
		p.bs.AdvanceEpochs += idx.Epoch(epochsNum)
		if p.bs.AdvanceEpochs > maxAdvanceEpochs {
			p.bs.AdvanceEpochs = maxAdvanceEpochs
		}
	}
}

func (p *DriverTxListener) Update(bs iblockproc.BlockState, es iblockproc.EpochState) {
	p.bs, p.es = bs, es
}

func (p *DriverTxListener) Finalize() iblockproc.BlockState {
	return p.bs
}
