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
	"errors"
	"fmt"

	"container/list"

	"github.com/BTWhite/go-btw-photon/crypto/sha256"
	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/db/sqlite3"
	"github.com/BTWhite/go-btw-photon/types"
)

var (
	tbl = leveldb.CreateTable([]byte("tx"))

	// ErrTxAlreadyExist is returned if tx already exist in tx list.
	ErrTxAlreadyExist = errors.New("Tx already exist")

	// ErrTxNotFound is returned if tx not found.
	ErrTxNotFound = errors.New("Tx not found")
)

// Chain is a branch in a network.
type Chain struct {
	Id      types.Hash `json:"id"`
	Asset   types.Hash `json:"asset"`
	Root    types.Hash `json:"root"`
	Payload types.Hash `json:"payload"`
	Txs     *list.List `json:"txs"`
}

// NewChain creates a new chain with hash name.
func NewChain(id types.Hash) *Chain {
	sqlite3.Init(fmt.Sprint("data/", id, ".chain"))
	chain := &Chain{}
	chain.Txs = list.New()
	return chain
}

// GetBytes returns chain bytes array.
func (c *Chain) GetBytes() []byte {
	buff := new(bytes.Buffer)

	c.Id.WriteToBuff(buff, 0)
	c.Asset.WriteToBuff(buff, 0)
	c.Root.WriteToBuff(buff, 0)

	return buff.Bytes()
}

// CalcId calculates a hash of a chain by byte representation.
func (c *Chain) CalcId() types.Hash {
	h := []byte(sha256.Sha256Hex(c.GetBytes()))
	return types.NewHash(h)
}

// AddTx adds a new transaction to the chain.
func (c *Chain) AddTx(tx *types.Tx) (error, types.Hash) {
	exist, err := tbl.Has(tx.Id)

	if err != nil {
		return err, nil
	}

	if exist {
		return ErrTxAlreadyExist, nil
	}

	tbl.PutObject(tx.Id, tx)
	return nil, tx.Id
}

// GetTx gets a transaction from the chain.
func (c *Chain) GetTx(hash types.Hash) (error, bool, *types.Tx) {
	exist, err := tbl.Has(hash)

	if err != nil {
		return err, false, nil
	}

	if !exist {
		return ErrTxNotFound, false, nil
	}

	tx := types.NewTx()
	err = tbl.GetObject(hash, tx)

	if err != nil {
		return err, false, nil
	}

	return nil, true, tx
}
