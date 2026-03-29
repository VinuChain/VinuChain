package valkeystore

import (
	"crypto/subtle"
	"errors"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Fantom-foundation/go-opera/inter/validatorpk"
	"github.com/Fantom-foundation/go-opera/valkeystore/encryption"
)

type MemKeystore struct {
	mem  map[string]*encryption.PrivateKey
	auth map[string]string
}

func NewMemKeystore() *MemKeystore {
	return &MemKeystore{
		mem:  make(map[string]*encryption.PrivateKey),
		auth: make(map[string]string),
	}
}

func (m *MemKeystore) Has(pubkey validatorpk.PubKey) bool {
	_, ok := m.mem[m.idxOf(pubkey)]
	return ok
}

func (m *MemKeystore) Add(pubkey validatorpk.PubKey, key []byte, auth string) error {
	if m.Has(pubkey) {
		return ErrAlreadyExists
	}
	if pubkey.Type != validatorpk.Types.Secp256k1 {
		return encryption.ErrNotSupportedType
	}
	decoded, err := crypto.ToECDSA(key)
	if err != nil {
		return err
	}
	keyCopy := make([]byte, len(key))
	copy(keyCopy, key)
	m.mem[m.idxOf(pubkey)] = &encryption.PrivateKey{
		Type:    pubkey.Type,
		Bytes:   keyCopy,
		Decoded: decoded,
	}
	m.auth[m.idxOf(pubkey)] = auth
	return nil
}

func (m *MemKeystore) Get(pubkey validatorpk.PubKey, auth string) (*encryption.PrivateKey, error) {
	idx := m.idxOf(pubkey)
	storedAuth, found := m.auth[idx]
	// Always perform the comparison to avoid leaking key existence
	// via timing. When the key doesn't exist, compare against an
	// empty string — the result is discarded but the timing is uniform.
	authMatch := subtle.ConstantTimeCompare([]byte(storedAuth), []byte(auth))
	if !found {
		return nil, ErrNotFound
	}
	if authMatch != 1 {
		return nil, errors.New("could not decrypt key with given password")
	}
	return m.mem[idx], nil
}

func (m *MemKeystore) idxOf(pubkey validatorpk.PubKey) string {
	return string(pubkey.Bytes())
}
