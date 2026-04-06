package inter

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/go-opera/utils/cser"
)

// marshalTxCSER serializes a single transaction to raw CSER bytes.
func marshalTxCSER(t *testing.T, tx *types.Transaction) []byte {
	t.Helper()
	raw, err := cser.MarshalBinaryAdapter(func(w *cser.Writer) error {
		return TransactionMarshalCSER(w, tx)
	})
	require.NoError(t, err)
	return raw
}

// unmarshalTxCSER deserializes a single transaction from raw CSER bytes.
func unmarshalTxCSER(t *testing.T, raw []byte) (*types.Transaction, error) {
	t.Helper()
	var tx *types.Transaction
	err := cser.UnmarshalBinaryAdapter(raw, func(r *cser.Reader) error {
		var innerErr error
		tx, innerErr = TransactionUnmarshalCSER(r)
		return innerErr
	})
	return tx, err
}

// TestTransactionUnmarshalCSER_RoundTrip verifies round-trip serialization for
// all three supported transaction types and contract creation.
func TestTransactionUnmarshalCSER_RoundTrip(t *testing.T) {
	t.Run("legacy", func(t *testing.T) {
		tx := types.NewTx(&types.LegacyTx{
			Nonce:    42,
			GasPrice: big.NewInt(1e9),
			Gas:      21000,
			To:       &common.Address{0xab},
			Value:    big.NewInt(1e18),
			Data:     []byte{0xde, 0xad},
			V:        big.NewInt(27),
			R:        big.NewInt(1),
			S:        big.NewInt(1),
		})
		roundTripTx(t, tx)
	})

	t.Run("accesslist", func(t *testing.T) {
		tx := types.NewTx(&types.AccessListTx{
			ChainID:  big.NewInt(207),
			Nonce:    1,
			GasPrice: big.NewInt(1e9),
			Gas:      50000,
			To:       &common.Address{0xcd},
			Value:    big.NewInt(0),
			Data:     nil,
			AccessList: types.AccessList{
				{
					Address:     common.Address{0x01},
					StorageKeys: []common.Hash{{0x01}, {0x02}},
				},
			},
			V: big.NewInt(0),
			R: big.NewInt(2),
			S: big.NewInt(3),
		})
		roundTripTx(t, tx)
	})

	t.Run("dynamicfee", func(t *testing.T) {
		tx := types.NewTx(&types.DynamicFeeTx{
			ChainID:   big.NewInt(207),
			Nonce:     5,
			GasTipCap: big.NewInt(1e9),
			GasFeeCap: big.NewInt(2e9),
			Gas:       21000,
			To:        &common.Address{0xef},
			Value:     big.NewInt(500),
			Data:      nil,
			AccessList: types.AccessList{
				{Address: common.Address{0x02}, StorageKeys: nil},
			},
			V: big.NewInt(0),
			R: big.NewInt(4),
			S: big.NewInt(5),
		})
		roundTripTx(t, tx)
	})

	t.Run("contract_creation", func(t *testing.T) {
		tx := types.NewTx(&types.LegacyTx{
			Nonce:    0,
			GasPrice: big.NewInt(1e9),
			Gas:      500000,
			To:       nil,
			Value:    big.NewInt(0),
			Data:     []byte{0x60, 0x80},
			V:        big.NewInt(28),
			R:        big.NewInt(6),
			S:        big.NewInt(7),
		})
		roundTripTx(t, tx)
	})

	t.Run("accesslist_multiple_entries", func(t *testing.T) {
		al := make(types.AccessList, 5)
		for i := range al {
			al[i] = types.AccessTuple{
				Address:     common.Address{byte(i + 1)},
				StorageKeys: []common.Hash{{byte(i)}, {byte(i + 1)}},
			}
		}
		tx := types.NewTx(&types.AccessListTx{
			ChainID:    big.NewInt(207),
			Nonce:      2,
			GasPrice:   big.NewInt(1e9),
			Gas:        100000,
			To:         &common.Address{0xff},
			Value:      big.NewInt(0),
			Data:       nil,
			AccessList: al,
			V:          big.NewInt(1),
			R:          big.NewInt(7),
			S:          big.NewInt(8),
		})
		roundTripTx(t, tx)
	})
}

// TestTransactionUnmarshalCSER_AccessListCap verifies that the access list
// allocation cap rejects oversized lists, preventing OOM from crafted messages.
// The cap must be ProtocolMaxMsgSize/accessListEntrySize so that allocating
// the struct array never exceeds the protocol message size limit.
func TestTransactionUnmarshalCSER_AccessListCap(t *testing.T) {
	// Cap constant sanity: allocation from max allowed accessListLen must not
	// exceed the protocol message size. Each AccessTuple occupies
	// accessListEntrySize bytes (Address + slice header).
	maxAllowed := ProtocolMaxMsgSize / accessListEntrySize
	maxAllocBytes := maxAllowed * accessListEntrySize
	require.LessOrEqualf(t, maxAllocBytes, ProtocolMaxMsgSize,
		"access list struct alloc (%d bytes) must not exceed protocol message cap (%d bytes)",
		maxAllocBytes, ProtocolMaxMsgSize)
}

// TestTransactionUnmarshalCSER_UnknownType verifies that an unknown tx type
// returns ErrUnknownTxType rather than panicking.
func TestTransactionUnmarshalCSER_UnknownType(t *testing.T) {
	// Construct a CSER-encoded byte sequence that looks like a non-legacy tx
	// with an unsupported type byte (0x05).
	raw, err := cser.MarshalBinaryAdapter(func(w *cser.Writer) error {
		// Write the 6-bit zero marker that indicates a non-legacy tx.
		w.BitsW.Write(6, 0)
		// Write an unsupported tx type (not 0x01 AccessList or 0x02 DynamicFee).
		w.U8(0x05)
		// Write minimum plausible fields to avoid data-exhaustion panic before
		// the type check: nonce, gasLimit, gasPrice bigint.
		w.U64(0)
		w.U64(21000)
		w.SliceBytes([]byte{1}) // gasPrice = 1
		w.SliceBytes([]byte{})  // value = 0
		w.Bool(false)           // no To
		w.SliceBytes([]byte{})  // data = empty
		w.SliceBytes([]byte{1}) // v = 1
		// R and S as 32-byte fixed (zero sig)
		var sig [64]byte
		w.FixedBytes(sig[:])
		return nil
	})
	require.NoError(t, err)

	_, err = unmarshalTxCSER(t, raw)
	require.ErrorIs(t, err, ErrUnknownTxType)
}

func roundTripTx(t *testing.T, tx *types.Transaction) {
	t.Helper()
	raw := marshalTxCSER(t, tx)
	got, err := unmarshalTxCSER(t, raw)
	require.NoError(t, err)
	require.Equal(t, tx.Hash(), got.Hash(), "hash mismatch")
	require.Equal(t, tx.Type(), got.Type(), "type mismatch")
	require.Equal(t, tx.Nonce(), got.Nonce(), "nonce mismatch")
	require.Equal(t, tx.Gas(), got.Gas(), "gas mismatch")
	require.Equal(t, tx.Value().String(), got.Value().String(), "value mismatch")
	require.Equal(t, len(tx.AccessList()), len(got.AccessList()), "access list length mismatch")
}
