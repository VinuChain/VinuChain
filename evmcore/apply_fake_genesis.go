// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package evmcore

import (
	"crypto/ecdsa"
	"math"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Fantom-foundation/go-opera/inter"
)

var FakeGenesisTime = inter.Timestamp(1608600000 * time.Second)

// ApplyFakeGenesis writes or updates the genesis block in db.
func ApplyFakeGenesis(statedb *state.StateDB, time inter.Timestamp, balances map[common.Address]*big.Int) (*EvmBlock, error) {
	for acc, balance := range balances {
		statedb.SetBalance(acc, balance)
	}

	// initial block
	root, err := flush(statedb, true)
	if err != nil {
		return nil, err
	}
	block := genesisBlock(time, root)

	return block, nil
}

func flush(statedb *state.StateDB, clean bool) (root common.Hash, err error) {
	root, err = statedb.Commit(clean)
	if err != nil {
		return
	}
	err = statedb.Database().TrieDB().Commit(root, false, nil)
	if err != nil {
		return
	}

	if !clean {
		err = statedb.Database().TrieDB().Cap(0)
	}

	return
}

// genesisBlock makes genesis block with pretty hash.
func genesisBlock(time inter.Timestamp, root common.Hash) *EvmBlock {
	block := &EvmBlock{
		EvmHeader: EvmHeader{
			Number:   big.NewInt(0),
			Time:     time,
			GasLimit: math.MaxUint64,
			Root:     root,
			TxHash:   types.EmptyRootHash,
		},
	}

	return block
}

// MustApplyFakeGenesis writes the genesis block and state to db, panicking on error.
func MustApplyFakeGenesis(statedb *state.StateDB, time inter.Timestamp, balances map[common.Address]*big.Int) *EvmBlock {
	block, err := ApplyFakeGenesis(statedb, time, balances)
	if err != nil {
		log.Crit("ApplyFakeGenesis", "err", err)
	}
	return block
}

// FakeKey gets n-th fake private key.
func origFakeKey(n int) *ecdsa.PrivateKey {
	reader := rand.New(rand.NewSource(int64(n)))

	key, err := ecdsa.GenerateKey(crypto.S256(), reader)
	if err != nil {
		panic(err)
	}

	return key
}

func FakeKey(n int) *ecdsa.PrivateKey {
	var key *ecdsa.PrivateKey

	if n == 0 {
		key, _ = crypto.HexToECDSA("d881dc14cdf9621a36ae939138ed90ea787a1adc87b0f320bd67febbd56e2e35") // 0x9b761aEf8C0B5c751D20217634f8188dCE866EbA
	}
	if n == 1 {
		key, _ = crypto.HexToECDSA("7f947dbc55f9c12cf976ed5adaa62d1935a6659d659035efad789a98b19f94ee") // 0x5Bb0880ab787b348C48ee51D00e41e3cBc134272
	}
	if n == 2 {
		key, _ = crypto.HexToECDSA("e1dc824daaf5757c6e8d1a65de22a88c63d6dfa02df0a09f98074a4ff23b08b5") // 0x626e9F7f9C101Fb0dBf14E0f175BbFFd757F5e2C
	}
	if n == 3 {
		key, _ = crypto.HexToECDSA("82410b11d55de7d52e916888742365ea74489b6f7f52032162fc1423f60e26c8") // 0x7421Af93A2B5E8b63C7Cb2267903967784Bd9fd4
	}

	return key
}
