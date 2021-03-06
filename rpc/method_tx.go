// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package rpc

import (
	"github.com/BTWhite/go-btw-photon/chain"
	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/events"
	"github.com/BTWhite/go-btw-photon/logger"
	"github.com/BTWhite/go-btw-photon/rawdb"
	"github.com/BTWhite/go-btw-photon/types"
)

var (
	ErrTxNotFound = err(0, rawdb.ErrTxNotFound.Error())
)

func init() {
	Register("tx.get", func() Executer { return new(GetTxRequest) })
	Register("tx.list", func() Executer { return new(GetTxListRequest) })
	Register("tx.create", func() Executer { return new(CreateTxRequest) })
	Register("tx.post", func() Executer { return new(PostTxRequest) })
}

type GetTxRequest struct {
	Id types.Hash `json:"id"`
}

func (preq *GetTxRequest) execute(r *Request) *Response {
	tx, err := cf.ChainHelper().GetTx(preq.Id)
	if err != nil {
		if err == rawdb.ErrTxNotFound {
			return response(nil, ErrTxNotFound)
		}
		logger.Err(lp, err.Error())
		return response(nil, ErrInternalError)
	}
	return response(tx, nil)
}

type GetTxListRequest struct {
	Limit int `json:"limit"`
}

func (preq *GetTxListRequest) execute(r *Request) *Response {
	it := cf.DataBase().NewIteratorPrefix([]byte("tx"))
	var txs []*types.Tx

	for i := 0; i < preq.Limit && it.Next(); i++ {
		bytes := it.Value()
		tx := types.NewTx()
		err := leveldb.Decode(bytes, tx)
		if err != nil {
			logger.Err(lp, err.Error(), bytes)
			continue
		}

		txs = append(txs, tx)

	}

	return response(txs, nil)
}

type CreateTxRequest struct {
	Secret  string `json:"secret"`
	Address string `json:"address"`
	Amount  uint64 `json:"amount"`
}

func (preq *CreateTxRequest) execute(r *Request) *Response {
	if len(preq.Secret) < 3 {
		return response(nil, err(0, "Please write correct secret"))
	}

	if !types.HasAddr([]byte(preq.Address)) {
		return response(nil, err(0, "Please write correct address"))
	}

	kp := types.NewKeyPair([]byte(preq.Secret))
	tx, e := cf.ChainHelper().NewTx(kp, types.Coin(preq.Amount),
		types.Coin(10000000), []byte(preq.Address))

	if e != nil {
		logger.Err(lp, e.Error())
		return response(nil, ErrInternalError)
	}

	e = cf.ChainHelper().ProcessTx(tx)

	if e != nil {
		logger.Err(lp, e.Error())
		return response(nil, err(0, e.Error()))
	}

	return response(tx.Id.String(), nil)
}

type PostTxRequest struct {
	types.Tx
}

func (preq *PostTxRequest) execute(r *Request) *Response {
	e := cf.ChainHelper().ProcessTx(&preq.Tx)

	if e != nil {
		if e == chain.ErrInsufficientData && r.Peer != nil {

			e := &InsufficientDataEvent{}
			e.SetData(preq.Chain, preq.Id, r.Peer)
			go events.Push("insufficent_data_tx", e)
		}

		return response(nil, err(0, e.Error()))
	}

	return response(preq.Id.String(), nil)
}
