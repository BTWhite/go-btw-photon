// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package chain

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/json"
	"github.com/BTWhite/go-btw-photon/types"
)

type genesis struct {
	Chain        *Chain      `json:"chain"`
	Transactions []*types.Tx `json:"transactions"`
}

// LoadGenesis loads the genesis chain from the file.
func LoadGenesis(filename string, chTbl *leveldb.Tbl, txTbl *leveldb.Tbl) (*Chain, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	g := &genesis{
		Chain: NewChain(txTbl, chTbl),
	}

	json.FromJson(data, &g.Transactions)

	for _, tx := range g.Transactions {
		tx.Timestamp = time.Now().Unix()
		tx.Mine()

		err := g.Chain.AddTx(tx)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	g.Chain.UpdatePayload()
	g.Chain.Id = g.Chain.CalcId()

	g.Chain.Save()
	return g.Chain, nil
}
