package valkeystore

import (
	"crypto/ecdsa"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetUnlockedReturnsDefensiveCopy(t *testing.T) {
	require := require.New(t)
	ks := NewDefaultMemKeystore()

	err := ks.backend.Add(pubkey1, key1, "auth1")
	require.NoError(err)
	err = ks.backend.Unlock(pubkey1, "auth1")
	require.NoError(err)

	// Get two copies — they should be independent
	k1, err := ks.GetUnlocked(pubkey1)
	require.NoError(err)
	k2, err := ks.GetUnlocked(pubkey1)
	require.NoError(err)

	// Verify both have correct bytes
	require.Equal(key1, k1.Bytes)
	require.Equal(key1, k2.Bytes)

	// Mutate k1's bytes — k2 should be unaffected
	k1.Bytes[0] ^= 0xff
	require.NotEqual(k1.Bytes[0], k2.Bytes[0], "mutation of one copy must not affect the other")

	// Mutate k1's decoded key D — k2 should be unaffected
	ecK1 := k1.Decoded.(*ecdsa.PrivateKey)
	ecK2 := k2.Decoded.(*ecdsa.PrivateKey)
	originalD := ecK2.D.Int64()
	ecK1.D.SetInt64(0)
	require.Equal(originalD, ecK2.D.Int64(), "mutation of decoded key must not affect the other copy")
}

func TestLockDoesNotCorruptConcurrentSign(t *testing.T) {
	require := require.New(t)
	ks := NewDefaultMemKeystore()

	err := ks.backend.Add(pubkey1, key1, "auth1")
	require.NoError(err)
	err = ks.backend.Unlock(pubkey1, "auth1")
	require.NoError(err)

	signer := NewSigner(ks)
	digest := make([]byte, 32)
	for i := range digest {
		digest[i] = byte(i)
	}

	// Concurrent signing and locking — sign should succeed or return ErrLocked,
	// never panic or use a zeroed key
	var wg sync.WaitGroup
	errs := make(chan error, 100)

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := signer.Sign(pubkey1, digest)
			if err != nil {
				errs <- err
			}
		}()
	}

	// Lock midway through the signing storm
	wg.Add(1)
	go func() {
		defer wg.Done()
		ks.Lock(pubkey1)
	}()

	wg.Wait()
	close(errs)

	// Every error should be ErrLocked — never a panic or zeroed-key signature
	for err := range errs {
		require.ErrorIs(err, ErrLocked, "expected ErrLocked, got: %v", err)
	}
}
