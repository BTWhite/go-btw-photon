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
	"io/ioutil"
	"time"

	"github.com/BTWhite/go-btw-photon/chain"
	"github.com/BTWhite/go-btw-photon/json"
	"github.com/BTWhite/go-btw-photon/logger"
	"github.com/BTWhite/go-btw-photon/types"
)

var (
	genesisChain = types.NewHash([]byte("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"))

	ErrGenesisLoaded = errors.New("The genesis chain has already been loaded")
)

// LoadGenesis loads the genesis chain from the file.
func (cb *ChainBook) LoadGenesis(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	ch := chain.NewChain(cb.txTbl, cb.chTbl)
	ch.CalcId()

	if ch.Id.Equals(genesisChain) {
		err, _ := cb.GetChain(genesisChain)
		if err == nil {
			return ErrGenesisLoaded
		}
	}
	logger.Debug("Loading genesis chain...")
	cb.AddChain(ch)

	var txs []*types.Tx
	json.FromJson(data, &txs)

	for _, tx := range txs {
		tx.Chain = ch.Id
		tx.Timestamp = time.Now().Unix()
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
