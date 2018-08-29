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
	"sync"

	"github.com/BTWhite/go-btw-photon/crypto/sha256"
	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/types"
)

var (
	// ErrTxNotFoundInChain is returned if the transaction is not found
	// or is not in a particular chain.
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
	Last    types.Hash   `json:"lastTx"`

	txTbl *leveldb.Tbl
	chTbl *leveldb.Tbl
	mu    sync.Mutex
}

// NewChain creates a new chain with hash name.
func NewChain(txTbl *leveldb.Tbl, chTbl *leveldb.Tbl) *Chain {

	chain := &Chain{
		txTbl: txTbl,
		chTbl: chTbl,
	}
	return chain
}

// Save writes chain to the database.
func (c *Chain) Save() error {
	return c.chTbl.PutObject(c.Id, c)
}

// CalcId calculates a hash of a chain by byte representation.
func (c *Chain) CalcId() types.Hash {
	buff := new(bytes.Buffer)

	c.RootCh.WriteToBuff(buff, 0)
	c.RootTx.WriteToBuff(buff, 0)

	h := []byte(sha256.Sha256Hex(buff.Bytes()))
	hash := types.NewHash(h)
	c.Id = hash
	return hash
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

// LastTx returns last tx hash in this chain.
func (c *Chain) LastTx() types.Hash {
	if len(c.Last) != 0 {
		return c.Last
	}
	return c.RootTx
}

// AddTx adds a new transaction to the chain.
func (c *Chain) AddTx(tx *types.Tx) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	hash, err := tx.Save(c.txTbl)

	if err != nil {
		return err
	}

	c.Txs = append(c.Txs, hash)
	c.Height++
	c.Last = hash

	return c.chTbl.PutObject(c.Id, c)
}

// GetTx gets a transaction from the chain.
func (c *Chain) GetTx(hash types.Hash) (*types.Tx, error) {
	tx, err := types.GetTx(hash, c.txTbl)
	if err != nil {
		return nil, err
	}

	if !tx.Chain.Equals(c.Id) {
		return nil, ErrTxNotFoundInChain
	}

	return tx, nil
}

func (c *Chain) sortTx() {
	sort.Slice(c.Txs, func(a, b int) bool {
		return c.Txs[a].String() < c.Txs[b].String()
	})
}
