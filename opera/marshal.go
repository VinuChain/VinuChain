package opera

import "encoding/json"

// UpdateRules applies a JSON diff to the current rules, returning the updated
// ruleset. Certain fields are protected from governance changes and are always
// restored from the original rules after the diff is applied:
//   - NetworkID and Name (identity)
//   - Upgrades (hard fork flags — activation requires a new binary release)
//   - Economy.QuotaCacheAddress (payback contract binding — set at genesis or
//     via binary upgrade to prevent redirection to a rogue contract)
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
	return
}
