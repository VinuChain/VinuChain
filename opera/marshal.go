package opera

import (
	"encoding/json"
	"errors"
	"math/big"
	"time"

	"github.com/Fantom-foundation/go-opera/inter"
)

// UpdateRules applies a JSON diff to the current rules, returning the updated
// ruleset. Certain fields are protected from governance changes and are always
// restored from the original rules after the diff is applied:
//   - NetworkID and Name (identity)
//   - Upgrades (hard fork flags — activation requires a new binary release)
//   - Economy.QuotaCacheAddress (payback contract binding — set at genesis or
//     via binary upgrade to prevent redirection to a rogue contract)
//
// After applying the diff, sanity bounds are checked on consensus-critical
// parameters to prevent a single malicious governance proposal from halting
// the chain or disabling Byzantine fault detection.
func UpdateRules(src Rules, diff []byte) (res Rules, err error) {
	changed := src.Copy()
	err = json.Unmarshal(diff, &changed)
	if err != nil {
		return src, err
	}
	// protect readonly fields
	res = changed
	res.NetworkID = src.NetworkID
	res.Name = src.Name
	res.Upgrades = src.Upgrades
	res.Economy.QuotaCacheAddress = src.Economy.QuotaCacheAddress

	// sanity bounds on consensus-critical parameters
	if err := validateRulesBounds(res); err != nil {
		return src, err
	}
	return
}

// validateRulesBounds rejects governance proposals that would set
// consensus-critical parameters to values that halt the chain or
// disable safety mechanisms.
func validateRulesBounds(r Rules) error {
	if r.Dag.MaxParents < 3 {
		return errors.New("Dag.MaxParents must be at least 3 for BFT connectivity")
	}
	if r.Economy.Gas.MaxEventGas == 0 {
		return errors.New("Economy.Gas.MaxEventGas cannot be zero")
	}
	if r.Economy.Gas.EventGas > r.Economy.Gas.MaxEventGas {
		return errors.New("Economy.Gas.EventGas exceeds MaxEventGas")
	}
	if r.Epochs.MaxEpochGas == 0 {
		return errors.New("Epochs.MaxEpochGas cannot be zero")
	}
	if r.Economy.MinGasPrice == nil {
		return errors.New("Economy.MinGasPrice cannot be nil")
	}
	if r.Economy.MinGasPrice.Sign() <= 0 {
		return errors.New("Economy.MinGasPrice must be positive")
	}
	if r.Economy.ShortGasPower.AllocPerSec == 0 {
		return errors.New("Economy.ShortGasPower.AllocPerSec cannot be zero")
	}
	if r.Economy.ShortGasPower.MaxAllocPeriod < inter.Timestamp(time.Second) {
		return errors.New("Economy.ShortGasPower.MaxAllocPeriod must be at least 1 second")
	}
	if r.Economy.LongGasPower.AllocPerSec == 0 {
		return errors.New("Economy.LongGasPower.AllocPerSec cannot be zero")
	}
	if r.Economy.LongGasPower.MaxAllocPeriod < inter.Timestamp(time.Second) {
		return errors.New("Economy.LongGasPower.MaxAllocPeriod must be at least 1 second")
	}
	// Upper bounds on gas power allocation — astronomical values distort the constructive gas price multiplier
	if r.Economy.ShortGasPower.AllocPerSec > 1e12 {
		return errors.New("Economy.ShortGasPower.AllocPerSec cannot exceed 1e12")
	}
	if r.Economy.LongGasPower.AllocPerSec > 1e12 {
		return errors.New("Economy.LongGasPower.AllocPerSec cannot exceed 1e12")
	}
	if r.Economy.ShortGasPower.MaxAllocPeriod > inter.Timestamp(7*24*time.Hour) {
		return errors.New("Economy.ShortGasPower.MaxAllocPeriod cannot exceed 1 week")
	}
	if r.Economy.LongGasPower.MaxAllocPeriod > inter.Timestamp(7*24*time.Hour) {
		return errors.New("Economy.LongGasPower.MaxAllocPeriod cannot exceed 1 week")
	}
	if r.Blocks.MaxBlockGas == 0 {
		return errors.New("Blocks.MaxBlockGas cannot be zero")
	}
	if r.Economy.Gas.MisbehaviourProofGas > r.Economy.Gas.MaxEventGas/2 {
		return errors.New("Economy.Gas.MisbehaviourProofGas must be less than half of MaxEventGas")
	}
	// Epoch duration bounds — prevent eternal epoch (MaxUint64) and per-block sealing (0)
	if r.Epochs.MaxEpochDuration == 0 || r.Epochs.MaxEpochDuration < inter.Timestamp(time.Minute) {
		return errors.New("Epochs.MaxEpochDuration must be at least 1 minute")
	}
	if r.Epochs.MaxEpochDuration > inter.Timestamp(7*24*time.Hour) {
		return errors.New("Epochs.MaxEpochDuration cannot exceed 1 week")
	}
	// MaxEpochGas upper bound — prevent eternal epoch via unreachable gas target
	if r.Epochs.MaxEpochGas > 1e15 {
		return errors.New("Epochs.MaxEpochGas cannot exceed 1e15")
	}
	// MaxFreeParents must not exceed MaxParents — underflow causes mempool freeze
	if r.Dag.MaxFreeParents > r.Dag.MaxParents {
		return errors.New("Dag.MaxFreeParents cannot exceed Dag.MaxParents")
	}
	// MinGasPrice upper bound — astronomical values censor all transactions
	maxMinGasPrice := new(big.Int).Exp(big.NewInt(10), big.NewInt(15), nil)
	if r.Economy.MinGasPrice.Cmp(maxMinGasPrice) > 0 {
		return errors.New("Economy.MinGasPrice cannot exceed 1e15 wei")
	}
	// BlockMissedSlack must not be zero — causes emitter flooding + mass deactivation
	if r.Economy.BlockMissedSlack == 0 {
		return errors.New("Economy.BlockMissedSlack cannot be zero")
	}
	// Vote gas params must not be zero — free votes enable unbounded storage inflation
	if r.Economy.Gas.BlockVoteGas == 0 {
		return errors.New("Economy.Gas.BlockVoteGas cannot be zero")
	}
	if r.Economy.Gas.EpochVoteGas == 0 {
		return errors.New("Economy.Gas.EpochVoteGas cannot be zero")
	}
	// Validate that maxEmptyEventGas (the formula used by MaxGasLimit) does not exceed
	// MaxEventGas. If it does, MaxGasLimit returns 0 and no EVM transactions can be
	// included in blocks. Overflow-safe: detect multiplication overflow before summing.
	maxParentDiff := uint64(r.Dag.MaxParents - r.Dag.MaxFreeParents)
	parentGasTerm := maxParentDiff * r.Economy.Gas.ParentGas
	if maxParentDiff > 0 && parentGasTerm/maxParentDiff != r.Economy.Gas.ParentGas {
		return errors.New("Economy.Gas.ParentGas overflows uint64 with current Dag.MaxParents")
	}
	extraGasTerm := uint64(r.Dag.MaxExtraData) * r.Economy.Gas.ExtraDataGas
	if r.Dag.MaxExtraData > 0 && extraGasTerm/uint64(r.Dag.MaxExtraData) != r.Economy.Gas.ExtraDataGas {
		return errors.New("Economy.Gas.ExtraDataGas overflows uint64 with current Dag.MaxExtraData")
	}
	minEventGas := r.Economy.Gas.EventGas + parentGasTerm + extraGasTerm
	if minEventGas < r.Economy.Gas.EventGas {
		return errors.New("gas parameters overflow uint64 in minimum event gas calculation")
	}
	if minEventGas > r.Economy.Gas.MaxEventGas {
		return errors.New("gas parameters make all events exceed MaxEventGas, blocking EVM transactions")
	}
	return nil
}
