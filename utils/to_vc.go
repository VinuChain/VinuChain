package utils

import "math/big"

// ToVC converts VC to Wei
func ToVC(vc uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(vc), big.NewInt(1e18))
}
