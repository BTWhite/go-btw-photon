// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package chain

import (
	"bytes"
	"encoding/binary"
	"errors"
	"sort"

	"github.com/BTWhite/go-btw-photon/crypto/sha256"
	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/types"
)

var (
	ErrTxNotFoundInChain = errors.New("Tx not found in chain")
)

// Chain is a branch in a network.
type Chain struct {
	Id      types.Hash   `json:"id"`
	Height  uint32       `json:"height"`
	RootCh  types.Hash   `json:"root_ch"`
	RootTx  types.Hash   `json:"root_tx"`
	Payload types.Hash   `json:"payload"`
	Txs     []types.Hash `json:"txs"`

	txTbl *leveldb.Tbl
	chTbl *leveldb.Tbl
}

// NewChain creates a new chain with hash name.
func NewChain(txTbl *leveldb.Tbl, chTbl *leveldb.Tbl) *Chain {

	chain := &Chain{
		txTbl: txTbl,
		chTbl: chTbl,
	}
	return chain
}

// CalcId calculates a hash of a chain by byte representation.
func (c *Chain) CalcId() types.Hash {
	buff := new(bytes.Buffer)

	c.RootCh.WriteToBuff(buff, 0)
	c.RootTx.WriteToBuff(buff, 0)

	h := []byte(sha256.Sha256Hex(buff.Bytes()))
	return types.NewHash(h)
}

// UpdatePayload updates payload field responsible for the security
// of transactions inside.
func (c *Chain) UpdatePayload() types.Hash {
	c.sortTx()
	buff := new(bytes.Buffer)
	binary.Write(buff, binary.LittleEndian, c.Height)

	for _, th := range c.Txs {

		th.WriteToBuff(buff, 64)
	}
	data := buff.Bytes()
	hash := sha256.Sha256Hex(data)

	c.Payload = []byte(hash)

	return c.Payload
}

// AddTx adds a new transaction to the chain.
func (c *Chain) AddTx(tx *types.Tx) error {
	_, hash := tx.Save(c.txTbl)

	c.Txs = append(c.Txs, hash)
	c.Height++
	return c.chTbl.PutObject(c.Id, c)
}

// GetTx gets a transaction from the chain.
func (c *Chain) GetTx(hash types.Hash) (error, *types.Tx) {
	err, tx := types.GetTx(hash, c.txTbl)
	if err != nil {
		return err, nil
	}

	if !tx.Chain.Equals(c.Id) {
		return ErrTxNotFoundInChain, nil
	}

	return nil, tx
}

func (c *Chain) sortTx() {
	sort.Slice(c.Txs, func(a, b int) bool {
		return c.Txs[a][0] < c.Txs[b][0]
	})
}
