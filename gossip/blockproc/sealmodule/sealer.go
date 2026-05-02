package sealmodule

import (
	"math/big"
	"sort"

	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/inter/pos"
	"github.com/Fantom-foundation/lachesis-base/lachesis"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Fantom-foundation/go-opera/gossip/blockproc"
	"github.com/Fantom-foundation/go-opera/inter/iblockproc"
)

type OperaEpochsSealerModule struct{}

func New() *OperaEpochsSealerModule {
	return &OperaEpochsSealerModule{}
}

func (m *OperaEpochsSealerModule) Start(block iblockproc.BlockCtx, bs iblockproc.BlockState, es iblockproc.EpochState) blockproc.SealerProcessor {
	return &OperaEpochsSealer{
		block: block,
		es:    es,
		bs:    bs,
	}
}

type OperaEpochsSealer struct {
	block iblockproc.BlockCtx
	es    iblockproc.EpochState
	bs    iblockproc.BlockState
}

func (s *OperaEpochsSealer) EpochSealing() bool {
	sealEpoch := s.bs.EpochGas >= s.es.Rules.Epochs.MaxEpochGas
	sealEpoch = sealEpoch || (s.block.Time-s.es.EpochStart) >= s.es.Rules.Epochs.MaxEpochDuration
	sealEpoch = sealEpoch || s.bs.AdvanceEpochs > 0
	return sealEpoch || s.bs.EpochCheaters.Len() != 0
}

func (p *OperaEpochsSealer) Update(bs iblockproc.BlockState, es iblockproc.EpochState) {
	p.bs, p.es = bs, es
}

// SealEpoch is called after pre-internal transactions are executed
func (s *OperaEpochsSealer) SealEpoch() (iblockproc.BlockState, iblockproc.EpochState) {
	// Select new validators
	oldValidators := s.es.Validators
	builder := pos.NewBigBuilder()
	validatorIDs := make([]idx.ValidatorID, 0, len(s.bs.NextValidatorProfiles))
	for v := range s.bs.NextValidatorProfiles {
		validatorIDs = append(validatorIDs, v)
	}
	sort.Slice(validatorIDs, func(i, j int) bool { return validatorIDs[i] < validatorIDs[j] })
	for _, v := range validatorIDs {
		profile := s.bs.NextValidatorProfiles[v]
		if s.es.Rules.Upgrades.Elemont && len(profile.PubKey.Raw) == 0 {
			log.Warn("Skipping validator with empty pubkey at epoch seal", "id", v)
			continue
		}
		// ElemontPubkeyValidation rejects validators whose pubkey is non-empty
		// but malformed (wrong type byte or wrong length). Gated separately
		// from Elemont so existing chaindata replay — which already admitted
		// at least one such validator on testnet — stays bit-for-bit identical
		// up to flag activation.
		if s.es.Rules.Upgrades.ElemontPubkeyValidation {
			if err := profile.PubKey.Validate(); err != nil {
				log.Warn("Skipping validator with malformed pubkey at epoch seal", "id", v, "err", err)
				continue
			}
		}
		builder.Set(v, profile.Weight)
	}
	newValidators := builder.Build()
	s.es.Validators = newValidators
	s.es.ValidatorProfiles = s.bs.NextValidatorProfiles.Copy()

	// Build new []ValidatorEpochState and []ValidatorBlockState
	newValidatorEpochStates := make([]iblockproc.ValidatorEpochState, newValidators.Len())
	newValidatorBlockStates := make([]iblockproc.ValidatorBlockState, newValidators.Len())
	for newValIdx := idx.Validator(0); newValIdx < newValidators.Len(); newValIdx++ {
		// default values
		newValidatorBlockStates[newValIdx] = iblockproc.ValidatorBlockState{
			Originated: new(big.Int),
		}
		// inherit validator's state from previous epoch, if any
		valID := newValidators.GetID(newValIdx)
		if !oldValidators.Exists(valID) {
			// new validator
			newValidatorBlockStates[newValIdx].LastBlock = s.block.Idx
			newValidatorBlockStates[newValIdx].LastOnlineTime = s.block.Time
			continue
		}
		oldValIdx := oldValidators.GetIdx(valID)
		src := s.bs.ValidatorStates[oldValIdx]
		if src.Originated == nil {
			src.Originated = new(big.Int)
		}
		src.Originated = new(big.Int).Set(src.Originated)
		newValidatorBlockStates[newValIdx] = src
		newValidatorBlockStates[newValIdx].DirtyGasRefund = 0
		newValidatorBlockStates[newValIdx].Uptime = 0
		newValidatorEpochStates[newValIdx].GasRefund = s.bs.ValidatorStates[oldValIdx].DirtyGasRefund
		newValidatorEpochStates[newValIdx].PrevEpochEvent = s.bs.ValidatorStates[oldValIdx].LastEvent
	}
	s.es.ValidatorStates = newValidatorEpochStates
	s.bs.ValidatorStates = newValidatorBlockStates
	s.es.Validators = newValidators

	// dirty data becomes active
	s.es.PrevEpochStart = s.es.EpochStart
	s.es.EpochStart = s.block.Time
	if s.bs.DirtyRules != nil {
		s.es.Rules = s.bs.DirtyRules.Copy()
		s.bs.DirtyRules = nil
	}
	s.es.EpochStateRoot = s.bs.FinalizedStateRoot

	s.bs.EpochGas = 0
	s.bs.EpochCheaters = lachesis.Cheaters{}
	s.bs.CheatersWritten = 0
	newEpoch := s.es.Epoch + 1
	s.es.Epoch = newEpoch

	if s.bs.AdvanceEpochs > 0 {
		s.bs.AdvanceEpochs--
	}

	return s.bs, s.es
}
