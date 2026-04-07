package valkeystore

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Fantom-foundation/go-opera/inter/validatorpk"
	"github.com/Fantom-foundation/go-opera/valkeystore/encryption"
)

type SignerI interface {
	Sign(pubkey validatorpk.PubKey, digest []byte) ([]byte, error)
}

type Signer struct {
	backend KeystoreI
}

func NewSigner(backend KeystoreI) *Signer {
	return &Signer{
		backend: backend,
	}
}

func (s *Signer) Sign(pubkey validatorpk.PubKey, digest []byte) ([]byte, error) {
	if pubkey.Type != validatorpk.Types.Secp256k1 {
		return nil, encryption.ErrNotSupportedType
	}
	key, err := s.backend.GetUnlocked(pubkey)
	if err != nil {
		return nil, err
	}
	defer zeroPrivateKey(key)

	secp256k1Key, ok := key.Decoded.(*ecdsa.PrivateKey)
	if !ok {
		return nil, errors.New("decoded key is not *ecdsa.PrivateKey")
	}

	sigRSV, err := crypto.Sign(digest, secp256k1Key)
	if err != nil {
		return nil, err
	}
	sigRS := make([]byte, 64)
	copy(sigRS, sigRSV[:64])
	return sigRS, nil
}

// zeroPrivateKey overwrites the key material in a PrivateKey copy so it
// doesn't persist on the Go heap waiting for GC.
func zeroPrivateKey(key *encryption.PrivateKey) {
	for i := range key.Bytes {
		key.Bytes[i] = 0
	}
	if ecKey, ok := key.Decoded.(*ecdsa.PrivateKey); ok && ecKey != nil && ecKey.D != nil {
		ecKey.D.SetUint64(0)
	}
	key.Decoded = nil
}
