package validatorpk

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type PubKey struct {
	Type uint8
	Raw  []byte
}

var Types = struct {
	Secp256k1 uint8
}{
	Secp256k1: 0xc0,
}

func (pk PubKey) Empty() bool {
	return len(pk.Raw) == 0 && pk.Type == 0
}

func (pk PubKey) String() string {
	return "0x" + common.Bytes2Hex(pk.Bytes())
}

func (pk PubKey) Bytes() []byte {
	return append([]byte{pk.Type}, pk.Raw...)
}

func (pk PubKey) Copy() PubKey {
	return PubKey{
		Type: pk.Type,
		Raw:  common.CopyBytes(pk.Raw),
	}
}

func FromString(str string) (PubKey, error) {
	return FromBytes(common.FromHex(str))
}

func FromBytes(b []byte) (PubKey, error) {
	if len(b) == 0 {
		return PubKey{}, errors.New("empty pubkey")
	}
	return PubKey{b[0], b[1:]}, nil
}

// ErrMalformedPubkey is returned by FromBytesValidated and Validate when a
// pubkey deviates from the canonical secp256k1 layout (1 type byte 0xc0 +
// 65 raw bytes). FromBytes intentionally does NOT use this — it must remain
// backward-compatible so chain replay reproduces existing on-chain state
// (including the malformed validator admitted on testnet at epoch 5682).
var ErrMalformedPubkey = errors.New("malformed pubkey")

// FromBytesValidated parses an on-wire pubkey serialization and rejects any
// shape that is not exactly 66 bytes prefixed with the Secp256k1 type byte.
// Use this at admission boundaries (RPC, parser of new createValidator txs)
// where rejecting a malformed input is correct. Do NOT use it on a chain
// replay path: existing chaindata may contain malformed pubkeys whose
// admission has already been finalized, and rejecting them retroactively
// would diverge consensus.
func FromBytesValidated(b []byte) (PubKey, error) {
	if len(b) != 66 || b[0] != Types.Secp256k1 {
		return PubKey{}, ErrMalformedPubkey
	}
	return PubKey{b[0], b[1:]}, nil
}

// Validate returns ErrMalformedPubkey if the parsed pubkey is not the
// canonical secp256k1 shape (Type==0xc0 and len(Raw)==65). Callers that
// already hold a parsed PubKey (e.g. read from validator profiles persisted
// in epoch state) can use this to selectively reject malformed entries
// behind an upgrade flag without changing FromBytes' permissive behavior.
func (pk PubKey) Validate() error {
	if pk.Type != Types.Secp256k1 || len(pk.Raw) != 65 {
		return ErrMalformedPubkey
	}
	return nil
}

// MarshalText returns the hex representation of a.
func (pk *PubKey) MarshalText() ([]byte, error) {
	return []byte(pk.String()), nil
}

// UnmarshalText parses a hash in hex syntax.
func (pk *PubKey) UnmarshalText(input []byte) error {
	res, err := FromString(string(input))
	if err != nil {
		return err
	}
	*pk = res
	return nil
}
