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
	"sort"
	"sync"

	"github.com/BTWhite/go-btw-photon/crypto/sha256"
	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/logger"
	"github.com/BTWhite/go-btw-photon/types"
)

// Chain is a branch in a network.
type Chain struct {
	Id      types.Hash   `json:"id"`
	Height  uint32       `json:"height"`
	Payload types.Hash   `json:"payload"`
	Txs     []types.Hash `json:"txs"`

	txTbl   *leveldb.Tbl
	txBatch *leveldb.TblBatch
	chTbl   *leveldb.Tbl
	proc    TxProcessor
	mu      sync.Mutex
	muATX   sync.Mutex
}

// NewChain creates a new chain with hash name.
func NewChain(db *leveldb.Db, proc TxProcessor, id types.Hash,
	genesisTx *types.Tx) (*Chain, error) {

	chTbl := db.CreateTable([]byte("chn"))
	txTbl := db.CreateTable([]byte("tx"))
	txBatch := db.NewBatch().CreateTableBatch([]byte("tx"))

	chain := &Chain{
		Id:      id,
		txTbl:   txTbl,
		txBatch: txBatch,
		chTbl:   chTbl,
		proc:    proc,
	}

	chain.chTbl.GetObject(id, chain)

	if genesisTx != nil {
		err := chain.AddTx(genesisTx)

		if err != nil {
			return nil, err
		}
	}

	chain.UpdatePayload()

	return chain, nil
}

// Save writes chain to the database.
func (c *Chain) Save() error {
	return c.chTbl.PutObject(c.Id, c)
}

// UpdatePayload updates payload field responsible for the security
// of transactions inside.
func (c *Chain) UpdatePayload() types.Hash {
	//c.sortTx()
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
	if len(c.Txs) == 0 {
		return nil
	}
	return c.Txs[len(c.Txs)-1]
}

// AddTx adds a new transaction to the chain.
func (c *Chain) AddTx(tx *types.Tx) error {
	_, err := types.GetTx(tx.Id, c.txTbl)
	if err == nil {
		return ErrTxAlreadyExist
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	err = c.proc.Validate(tx, c)
	if err != nil {
		return err
	}

	err = c.proc.Process(tx, c)
	if err != nil {
		return err
	}

	hash, err := c.proc.Save(tx, c, c.txTbl, c.txBatch)
	if err != nil {
		return err
	}

	c.AddTxLink(hash)
	logger.Debug("Tx", tx.Id, "processed")

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

func CreateTx(kp *types.KeyPair, amount types.Coin,
	fee types.Coin, recipient types.Hash) (*types.Tx, error) {

	tx := types.NewTx()
	tx.Amount = amount
	tx.Fee = fee
	tx.RecipientId = recipient
	tx.SenderId = []byte(kp.Public().Address())
	tx.Chain = []byte(sha256.Sha256Hex(tx.SenderId))

	//	sign.Sign(tx, kp, &tx.SenderPublicKey, &tx.Signature, 0)

	return tx, nil
}

func (c *Chain) sortTx() {
	sort.Slice(c.Txs, func(a, b int) bool {
		return c.Txs[a].String() < c.Txs[b].String()
	})
}

func (c *Chain) AddTxLink(id types.Hash) {
	c.muATX.Lock()
	c.Txs = append(c.Txs, id)
	c.Height++
	c.muATX.Unlock()
}
