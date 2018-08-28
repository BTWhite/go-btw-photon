// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package chainbook

import (
	"errors"

	"github.com/BTWhite/go-btw-photon/chain"
	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/interfaces"
	"github.com/BTWhite/go-btw-photon/types"
)

var (
	// ErrChainNotFound is returned is chain not found when writing a new transaction.
	ErrChainNotFound = errors.New("Chain not found")
)

// ChainBook is a `chain manager` that controls the invocation of the necessary
// transaction methods and chains.
type ChainBook struct {
	Chains    map[string]*chain.Chain
	processor interfaces.TxProcessor
	txTbl     *leveldb.Tbl
	chTbl     *leveldb.Tbl
}

// NewChainBook opens the chainbook.
// Waits for a table for transactions, a table for chains, and an implementing
// interfaces.TxProcessor object.
func NewChainBook(txTbl *leveldb.Tbl, chTbl *leveldb.Tbl,
	processor interfaces.TxProcessor) *ChainBook {

	cb := &ChainBook{
		chTbl:     chTbl,
		txTbl:     txTbl,
		processor: processor,
	}
	cb.Chains = make(map[string]*chain.Chain)
	return cb
}

// AddChain add chain to the chainbook list.
func (cb *ChainBook) AddChain(c *chain.Chain) {
	cb.Chains[c.Id.String()] = c
}

// GetChain gets chain from the chainbook list if chain exist.
func (cb *ChainBook) GetChain(hash types.Hash) (error, *chain.Chain) {
	var c *chain.Chain = cb.Chains[hash.String()]
	if c == nil {
		ch := chain.NewChain(cb.txTbl, cb.chTbl)
		err := cb.chTbl.GetObject(hash, ch)
		if err != nil {
			return ErrChainNotFound, nil
		}
		cb.Chains[ch.Id.String()] = ch
	}

	return nil, c
}

// AddTx is entry point for tx, the transaction will be transferred to the chain
// if it exists after the transaction is obtained (call `tx.Mine`).
// Before processing, transactions will also be changed `PreviousTx`.
// The processor's methods will also be called: `Validate` and `Process`.
func (cb *ChainBook) AddTx(tx *types.Tx) error {
	err, c := cb.GetChain(tx.Chain)
	if err != nil {
		return err
	}
	tx.PreviousTx = c.LastTx()
	tx.Chain = c.Id
	tx.Mine()

	err = cb.processor.Validate(tx, c)
	if err != nil {
		return err
	}
	err = c.AddTx(tx)
	if err != nil {
		return err
	}

	cb.processor.Process(tx, c)

	return nil
}
