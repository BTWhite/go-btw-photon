// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package chainbook

import (
	"io/ioutil"

	"github.com/BTWhite/go-btw-photon/chain"
	"github.com/BTWhite/go-btw-photon/json"
	"github.com/BTWhite/go-btw-photon/types"
)

// LoadGenesis loads the genesis chain from the file.
func (cb *ChainBook) LoadGenesis(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	ch := chain.NewChain(cb.txTbl, cb.chTbl)
	ch.CalcId()
	cb.AddChain(ch)

	var txs []*types.Tx
	json.FromJson(data, &txs)

	for _, tx := range txs {
		tx.Chain = ch.Id
		err := cb.AddTx(tx)

		if err != nil {
			return err
		}
	}

	ch.UpdatePayload()
	ch.CalcId()
	ch.Save()

	return nil
}
