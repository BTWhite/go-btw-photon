// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package chain

import (
	"sync"
	"time"

	"github.com/BTWhite/go-btw-photon/account"
	"github.com/BTWhite/go-btw-photon/crypto/sha256"
	"github.com/BTWhite/go-btw-photon/crypto/sign"
	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/types"
)

type ChainHelper struct {
	db    *leveldb.Db
	tblTx *leveldb.Tbl
	tblCh *leveldb.Tbl
	am    *account.AccountManager
	proc  TxProcessor
	mu    sync.Mutex
}

func NewChainHelper(db *leveldb.Db) *ChainHelper {
	h := &ChainHelper{
		db:    db,
		tblCh: db.CreateTable([]byte("tbl")),
		tblTx: db.CreateTable([]byte("tx")),
		proc:  NewProcessor(db),
		am:    account.NewAccountManager(db),
	}

	return h
}

func (h *ChainHelper) GetChainByAddress(address []byte) (*Chain, error) {
	return h.GetChainById([]byte(sha256.Sha256Hex(address)))
}

func (h *ChainHelper) GetChainById(id []byte) (*Chain, error) {
	return NewChain(h.db, h.proc, id, nil)
}

func (h *ChainHelper) NewTx(kp *types.KeyPair, amount types.Coin, fee types.Coin,
	recipient types.Hash) (*types.Tx, error) {

	addr := []byte(kp.Public().Address())
	pub := types.NewHash(*kp.Public())

	ch, err := h.GetChainByAddress(addr)
	if err != nil {
		return nil, err
	}

	lastTx := ch.LastTx()
	if lastTx == nil {
		return nil, ErrChainEmpty
	}

	tx := types.NewTx()
	tx.Amount = amount
	tx.Fee = amount
	tx.Timestamp = time.Now().Unix()
	tx.SenderId = addr
	tx.SenderPublicKey = pub
	tx.RecipientId = recipient
	tx.PreviousTx = lastTx
	tx.Chain = ch.Id
	tx.GenerateId()
	sign.Sign(tx, kp, &tx.SenderPublicKey, &tx.Signature, 0)

	return tx, nil
}

func (h *ChainHelper) ProcessTx(tx *types.Tx) error {
	sch, err := h.GetChainByAddress(tx.SenderId)
	if err != nil {
		return err
	}

	rch, err := h.GetChainByAddress(tx.RecipientId)
	if err != nil {
		return err
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	err = sch.AddTx(tx)
	if err != nil {
		return err
	}
	rch.AddTxLink(tx.Id)

	return rch.Save()
}

func (h *ChainHelper) AccountManager() *account.AccountManager {
	return h.am
}

func (h *ChainHelper) GetTx(hash types.Hash) (*types.Tx, error) {
	return types.GetTx(hash, h.tblTx)
}
