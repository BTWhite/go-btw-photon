// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package chain

import (
	"io/ioutil"

	"github.com/BTWhite/go-btw-photon/json"
	"github.com/BTWhite/go-btw-photon/logger"
	"github.com/BTWhite/go-btw-photon/types"
)

var (
	genesisChain = types.NewHash([]byte("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"))
)

// LoadGenesis loads the genesis chain from the file.
func LoadGenesis(filename string, h *ChainHelper) error {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var txs []*types.Tx
	json.FromJson(data, &txs)
	if len(txs) == 0 {
		return nil
	}

	ch, err := h.GetChainById(genesisChain)
	if err != nil {
		return err
	}

	if len(ch.Txs) > 0 {
		return ErrGenesisLoaded
	}

	logger.Debug("Chain:", "Loading genesis chain...")

	for _, tx := range txs {

		tx.Chain = ch.Id
		tx.GenerateId()
		ch.AddTxLink(tx.Id)
		err := h.ProcessTx(tx)

		if err != nil {
			return err
		}
	}
	ch.UpdatePayload()
	ch.Save()
	return nil
}
