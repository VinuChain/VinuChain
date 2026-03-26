package inter

import "fmt"

const SigSize = 64

// Signature is a secp256k1 signature in R|S format
type Signature [SigSize]byte

func (s Signature) Bytes() []byte {
	return s[:]
}

func BytesToSignature(b []byte) (sig Signature, err error) {
	if len(b) != SigSize {
		return sig, fmt.Errorf("invalid signature length: got %d, want %d", len(b), SigSize)
	}
	copy(sig[:], b)
	return sig, nil
}
