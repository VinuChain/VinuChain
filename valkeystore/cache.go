package valkeystore

import (
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/Fantom-foundation/go-opera/inter/validatorpk"
	"github.com/Fantom-foundation/go-opera/valkeystore/encryption"
)

var (
	ErrAlreadyUnlocked = errors.New("already unlocked")
	ErrLocked          = errors.New("key is locked")
)

type CachedKeystore struct {
	backend RawKeystoreI
	cache   map[string]*encryption.PrivateKey
}

func NewCachedKeystore(backend RawKeystoreI) *CachedKeystore {
	return &CachedKeystore{
		backend: backend,
		cache:   make(map[string]*encryption.PrivateKey),
	}
}

func (c *CachedKeystore) Unlocked(pubkey validatorpk.PubKey) bool {
	_, ok := c.cache[c.idxOf(pubkey)]
	return ok
}

func (c *CachedKeystore) Has(pubkey validatorpk.PubKey) bool {
	if c.Unlocked(pubkey) {
		return true
	}
	return c.backend.Has(pubkey)
}

func (c *CachedKeystore) Unlock(pubkey validatorpk.PubKey, auth string) error {
	if c.Unlocked(pubkey) {
		return ErrAlreadyUnlocked
	}
	key, err := c.backend.Get(pubkey, auth)
	if err != nil {
		return err
	}
	c.cache[c.idxOf(pubkey)] = key
	return nil
}

func (c *CachedKeystore) GetUnlocked(pubkey validatorpk.PubKey) (*encryption.PrivateKey, error) {
	if !c.Unlocked(pubkey) {
		return nil, ErrLocked
	}
	return copyPrivateKey(c.cache[c.idxOf(pubkey)]), nil
}

// copyPrivateKey returns a deep copy of the key so callers cannot mutate
// or retain a reference to the cached original. This also prevents a
// TOCTOU race where Lock() zeroes the cached key while Sign() uses it.
func copyPrivateKey(src *encryption.PrivateKey) *encryption.PrivateKey {
	if src == nil {
		return nil
	}
	dst := &encryption.PrivateKey{
		Type: src.Type,
	}
	if len(src.Bytes) > 0 {
		dst.Bytes = make([]byte, len(src.Bytes))
		copy(dst.Bytes, src.Bytes)
	}
	if ecKey, ok := src.Decoded.(*ecdsa.PrivateKey); ok && ecKey != nil {
		cpKey := new(ecdsa.PrivateKey)
		cpKey.PublicKey.Curve = ecKey.PublicKey.Curve
		cpKey.PublicKey.X = new(big.Int).Set(ecKey.PublicKey.X)
		cpKey.PublicKey.Y = new(big.Int).Set(ecKey.PublicKey.Y)
		cpKey.D = new(big.Int).Set(ecKey.D)
		dst.Decoded = cpKey
	}
	return dst
}

func (c *CachedKeystore) idxOf(pubkey validatorpk.PubKey) string {
	return string(pubkey.Bytes())
}

func (c *CachedKeystore) Add(pubkey validatorpk.PubKey, key []byte, auth string) error {
	return c.backend.Add(pubkey, key, auth)
}

func (c *CachedKeystore) Get(pubkey validatorpk.PubKey, auth string) (*encryption.PrivateKey, error) {
	return c.backend.Get(pubkey, auth)
}

func (c *CachedKeystore) Lock(pubkey validatorpk.PubKey) {
	idx := c.idxOf(pubkey)
	if key, ok := c.cache[idx]; ok {
		if key != nil {
			for i := range key.Bytes {
				key.Bytes[i] = 0
			}
			if ecKey, ok := key.Decoded.(*ecdsa.PrivateKey); ok && ecKey != nil {
				limbs := ecKey.D.Bits()
				for i := range limbs {
					limbs[i] = 0
				}
				ecKey.D.SetInt64(0)
			}
			key.Decoded = nil
		}
		delete(c.cache, idx)
	}
}
