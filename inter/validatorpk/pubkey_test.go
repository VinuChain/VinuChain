package validatorpk

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestFromString(t *testing.T) {
	require := require.New(t)
	exp := PubKey{
		Type: Types.Secp256k1,
		Raw:  common.FromHex("45b86101f804f3f4f2012ef31fff807e87de579a3faa7947d1b487a810e35dc2c3b6071ac465046634b5f4a8e09bf8e1f2e7eccb699356b9e6fd496ca4b1677d1"),
	}
	{
		got, err := FromString("c0045b86101f804f3f4f2012ef31fff807e87de579a3faa7947d1b487a810e35dc2c3b6071ac465046634b5f4a8e09bf8e1f2e7eccb699356b9e6fd496ca4b1677d1")
		require.NoError(err)
		require.Equal(exp, got)
	}
	{
		got, err := FromString("0xc0045b86101f804f3f4f2012ef31fff807e87de579a3faa7947d1b487a810e35dc2c3b6071ac465046634b5f4a8e09bf8e1f2e7eccb699356b9e6fd496ca4b1677d1")
		require.NoError(err)
		require.Equal(exp, got)
	}
	{
		_, err := FromString("")
		require.Error(err)
	}
	{
		_, err := FromString("0x")
		require.Error(err)
	}
	{
		_, err := FromString("-")
		require.Error(err)
	}
}

func TestString(t *testing.T) {
	require := require.New(t)
	pk := PubKey{
		Type: Types.Secp256k1,
		Raw:  common.FromHex("45b86101f804f3f4f2012ef31fff807e87de579a3faa7947d1b487a810e35dc2c3b6071ac465046634b5f4a8e09bf8e1f2e7eccb699356b9e6fd496ca4b1677d1"),
	}
	require.Equal("0xc0045b86101f804f3f4f2012ef31fff807e87de579a3faa7947d1b487a810e35dc2c3b6071ac465046634b5f4a8e09bf8e1f2e7eccb699356b9e6fd496ca4b1677d1", pk.String())
}

// validBytes returns a canonical 66-byte secp256k1 pubkey serialization
// (1 type byte + 65 raw bytes). The 04-prefixed 64-byte body matches the
// SEC1 uncompressed point encoding used by Ethereum/secp256k1 keys.
func validBytes() []byte {
	return common.FromHex("c0045b86101f804f3f4f2012ef31fff807e87de579a3faa7947d1b487a810e35dc2c3b6071ac465046634b5f4a8e09bf8e1f2e7eccb699356b9e6fd496ca4b1677d1")
}

// malformedVal16Bytes mirrors testnet validator 16's pubkey shape: 65 bytes
// starting with 0x04 (the SEC1 uncompressed marker leaked into the type slot)
// instead of the 66-byte 0xc0-prefixed canonical form.
func malformedVal16Bytes() []byte {
	// Exactly 130 hex chars = 65 bytes, leading byte 0x04.
	return common.FromHex("04d9aaa24f5b86101f804f3f4f2012ef31fff807e87de579a3faa7947d1b487a810e35dc2c3b6071ac465046634b5f4a8e09bf8e1f2e7eccb699356b9e6fd496ca")
}

func TestFromBytesValidated_AcceptsCanonical(t *testing.T) {
	require := require.New(t)
	pk, err := FromBytesValidated(validBytes())
	require.NoError(err)
	require.Equal(Types.Secp256k1, pk.Type)
	require.Len(pk.Raw, 65)
	require.NoError(pk.Validate())
}

func TestFromBytesValidated_RejectsWrongType(t *testing.T) {
	require := require.New(t)
	// 66 bytes, but type byte is 0x04 not 0xc0.
	b := append([]byte{0x04}, validBytes()[1:]...)
	require.Len(b, 66)
	_, err := FromBytesValidated(b)
	require.ErrorIs(err, ErrMalformedPubkey)
}

func TestFromBytesValidated_RejectsWrongLength65WithCorrectType(t *testing.T) {
	require := require.New(t)
	// 65 bytes with type=0xc0 — short by one byte vs canonical 66.
	b := append([]byte{Types.Secp256k1}, validBytes()[2:]...)
	require.Len(b, 65)
	_, err := FromBytesValidated(b)
	require.ErrorIs(err, ErrMalformedPubkey)
}

func TestFromBytesValidated_RejectsVal16Shape(t *testing.T) {
	require := require.New(t)
	// The actual val 16 testnet shape: 65 bytes starting with 0x04.
	b := malformedVal16Bytes()
	require.Len(b, 65)
	require.Equal(byte(0x04), b[0])
	_, err := FromBytesValidated(b)
	require.ErrorIs(err, ErrMalformedPubkey)
}

func TestFromBytesValidated_RejectsEmpty(t *testing.T) {
	require := require.New(t)
	_, err := FromBytesValidated(nil)
	require.Error(err)
}

func TestFromBytes_AcceptsMalformedForBackwardCompat(t *testing.T) {
	require := require.New(t)
	// FromBytes must NOT validate — chain replay relies on its current behavior
	// to reproduce the existing on-chain admission of testnet val 16.
	pk, err := FromBytes(malformedVal16Bytes())
	require.NoError(err, "FromBytes must remain backward-compatible: no validation")
	require.Equal(byte(0x04), pk.Type)
	require.Len(pk.Raw, 64)
	// But Validate() must reject the same parsed PubKey.
	require.ErrorIs(pk.Validate(), ErrMalformedPubkey)
}

func TestValidate_AcceptsCanonical(t *testing.T) {
	require := require.New(t)
	pk, err := FromBytes(validBytes())
	require.NoError(err)
	require.NoError(pk.Validate())
}

func TestValidate_RejectsEmpty(t *testing.T) {
	require := require.New(t)
	pk := PubKey{}
	require.ErrorIs(pk.Validate(), ErrMalformedPubkey)
}

func TestValidate_RejectsWrongTypeOnly(t *testing.T) {
	require := require.New(t)
	pk := PubKey{Type: 0x04, Raw: make([]byte, 65)}
	require.ErrorIs(pk.Validate(), ErrMalformedPubkey)
}

func TestValidate_RejectsWrongRawLengthOnly(t *testing.T) {
	require := require.New(t)
	pk := PubKey{Type: Types.Secp256k1, Raw: make([]byte, 64)}
	require.ErrorIs(pk.Validate(), ErrMalformedPubkey)
}

func TestFromBytes_PreservesEmptyError(t *testing.T) {
	require := require.New(t)
	_, err := FromBytes(nil)
	require.Error(err, "FromBytes must keep returning the existing empty-pubkey error")
}
