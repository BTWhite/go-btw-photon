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
	"fmt"

	"github.com/BTWhite/go-btw-photon/chain"
	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/interfaces"
	"github.com/BTWhite/go-btw-photon/types"
)

var (
	ErrChainNotFound = errors.New("Chain not found")
)

type ChainBook struct {
	Chains    map[string]*chain.Chain
	processor interfaces.TxProcessor
	txTbl     *leveldb.Tbl
	chTbl     *leveldb.Tbl
}

func NewChainBook(txTbl *leveldb.Tbl, chTbl *leveldb.Tbl, processor interfaces.TxProcessor) *ChainBook {
	cb := &ChainBook{
		chTbl:     chTbl,
		txTbl:     txTbl,
		processor: processor,
	}
	cb.Chains = make(map[string]*chain.Chain)
	return cb
}

func (cb *ChainBook) AddChain(c *chain.Chain) {
	cb.Chains[c.Id.String()] = c
}

func (cb *ChainBook) AddTx(tx *types.Tx) error {
	var c *chain.Chain = cb.Chains[tx.Chain.String()]
	if c == nil {
		ch := chain.NewChain(cb.txTbl, cb.chTbl)
		err := cb.chTbl.GetObject(tx.Chain, ch)
		if err != nil {
			fmt.Println(tx.Id, err.Error())
			return err
		}
		cb.Chains[ch.Id.String()] = ch
		c = ch
	}

	tx.PreviousTx = c.LastTx()
	tx.Chain = c.Id
	tx.Mine()

	err := cb.processor.Validate(tx, c)
	if err != nil {
		return err
	}
	err = c.AddTx(tx)
	if err != nil {
		return err
	}

	return cb.processor.Process(tx, c)
}
