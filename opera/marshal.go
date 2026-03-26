package opera

import (
	"encoding/json"
	"errors"
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
	if r.Economy.MinGasPrice != nil && r.Economy.MinGasPrice.Sign() < 0 {
		return errors.New("Economy.MinGasPrice cannot be negative")
	}
	if r.Economy.ShortGasPower.AllocPerSec == 0 {
		return errors.New("Economy.ShortGasPower.AllocPerSec cannot be zero")
	}
	if r.Economy.LongGasPower.AllocPerSec == 0 {
		return errors.New("Economy.LongGasPower.AllocPerSec cannot be zero")
	}
	if r.Blocks.MaxBlockGas == 0 {
		return errors.New("Blocks.MaxBlockGas cannot be zero")
	}
	if r.Economy.Gas.MisbehaviourProofGas > r.Economy.Gas.MaxEventGas/2 {
		return errors.New("Economy.Gas.MisbehaviourProofGas must be less than half of MaxEventGas")
	}
	return nil
}
