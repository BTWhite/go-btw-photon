// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package chain

import (
	"github.com/BTWhite/go-btw-photon/crypto/sign"
	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/logger"
	"github.com/BTWhite/go-btw-photon/types"
)

// ChainBook is a `chain manager` that controls the invocation of the necessary
// transaction methods and chains.
type ChainBook struct {
	Chains    map[string]*Chain
	processor TxProcessor
	txTbl     *leveldb.Tbl
	chTbl     *leveldb.Tbl
	db        *leveldb.Db
}

// NewChainBook opens the chainbook.
// Waits for a table for transactions, a table for chains, and an implementing
// interfaces.TxProcessor object.
func NewChainBook(db *leveldb.Db,
	processor TxProcessor) *ChainBook {
	chTbl := db.CreateTable([]byte("chn"))
	txTbl := db.CreateTable([]byte("tx"))

	cb := &ChainBook{
		chTbl:     chTbl,
		txTbl:     txTbl,
		db:        db,
		processor: processor,
	}
	cb.Chains = make(map[string]*Chain)
	return cb
}

// AddChain add chain to the chainbook list.
func (cb *ChainBook) AddChain(c *Chain) {
	cb.Chains[c.Id.String()] = c
}

// GetChain gets chain from the chainbook list if chain exist.
func (cb *ChainBook) GetChain(hash types.Hash) (*Chain, error) {
	var c *Chain = cb.Chains[hash.String()]
	if c == nil {
		ch := NewChain(cb.db)
		err := cb.chTbl.GetObject(hash, ch)
		if err != nil {
			return nil, ErrChainNotFound
		}
		cb.Chains[ch.Id.String()] = ch
	}

	return c, nil
}

// AddTx is entry point for tx, the transaction will be transferred to the chain
// if it exists after the transaction is obtained (call `tx.GenerateId`).
// Before processing, transactions will also be changed `PreviousTx`.
// The processor's methods will also be called: `Validate` and `Process`.
func (cb *ChainBook) AddTx(tx *types.Tx) error {
	c, err := cb.GetChain(tx.Chain)
	if err != nil {
		return err
	}
	tx.PreviousTx = c.LastTx()
	tx.Chain = c.Id
	tx.GenerateId()

	err = cb.processor.Validate(tx, c)
	if err != nil {
		return err
	}
	err = c.AddTx(tx)
	if err != nil {
		return err
	}
	err = cb.processor.Process(tx, c)
	if err != nil {
		return err
	}

	logger.Debug("Tx", tx.Id, "processed")
	return nil
}

// CreateTx creates and safe signing transaction
func (cb *ChainBook) CreateTx(kp *types.KeyPair, amount types.Coin,
	fee types.Coin, recipient types.Hash, chain types.Hash) *types.Tx {

	ch, _ := cb.GetChain(chain)

	tx := types.NewTx()
	tx.Amount = amount
	tx.Fee = fee
	tx.RecipientId = recipient
	tx.SenderId = []byte(kp.Public().Address())
	tx.Chain = chain
	tx.PreviousTx = ch.LastTx()
	sign.Sign(tx, kp, &tx.SenderPublicKey, &tx.Signature, 0)

	return tx
}
