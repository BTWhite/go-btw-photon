// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package rpc

import (
	"fmt"

	"github.com/BTWhite/go-btw-photon/chain"
	"github.com/BTWhite/go-btw-photon/types"
)

type LoadChainRequest struct {
	Chain string `json:"chain"`
	Start int    `json:"start"`
	Limit int    `json:"limit"`
}

type LoadChainResponse struct {
	Chain string      `json:"chain"`
	Start int         `json:"start"`
	Limit int         `json:"limit"`
	Txs   []*types.Tx `json:"txs"`
}

func init() {
	Register("chain.load", func() Executer { return new(LoadChainRequest) })
}

func (preq *LoadChainRequest) execute(r *Request) *Response {
	ch, e := cf.ChainHelper().GetChainById([]byte(preq.Chain))

	if e != nil {
		return response(nil, err(0, e.Error()))
	}

	if len(ch.Txs) <= preq.Start {
		return response(nil, err(0, fmt.Sprint("Start overflow, len:", len(ch.Txs))))
	}

	var result []*types.Tx
	for j := 0; j < preq.Limit; j++ {
		item := preq.Start + j
		if item > len(ch.Txs)-1 {
			break
		}

		tx, e := ch.GetTx(ch.Txs[item])
		if e != nil && e != chain.ErrTxNotFoundInChain {
			return response(nil, err(0, e.Error()))
		}

		result = append(result, tx)
	}
	resp := &LoadChainResponse{
		Chain: preq.Chain,
		Start: preq.Start,
		Limit: preq.Limit,
		Txs:   result,
	}

	return response(resp, nil)
}
