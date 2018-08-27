// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package types

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/mine"
)

var (
	// ErrTxAlreadyExist is returned if tx already exist in tx list.
	ErrTxAlreadyExist = errors.New("Tx already exist")

	// ErrTxNotFound is returned if tx not found.
	ErrTxNotFound = errors.New("Tx not found")
)

const (
	complexity = 4
)

// Tx determines the structure of the transaction.
type Tx struct {
	Id              Hash   `json:"id"`
	SenderPublicKey Hash   `json:"senderPublicKey"`
	SenderId        Hash   `json:"senderId"`
	RecipientId     Hash   `json:"recipientId"`
	Amount          Coin   `json:"amount"`
	Fee             Coin   `json:"fee"`
	Signature       Hash   `json:"signature"`
	Timestamp       int64  `json:"timestamp"`
	Nonce           uint32 `json:"nonce"`
	Height          uint32 `json:"height"`
	Chain           Hash   `json:"chain"`
	PreviousTx      Hash   `json:"previousTx"`
}

// NewTx creates new empty transaction.
func NewTx() *Tx {
	return &Tx{}
}

// GetBytes gets byte array by tx object.
func (t *Tx) GetBytes() []byte {
	buff := new(bytes.Buffer)

	binary.Write(buff, binary.LittleEndian, t.Timestamp)
	binary.Write(buff, binary.LittleEndian, t.Amount)
	binary.Write(buff, binary.LittleEndian, t.Fee)
	binary.Write(buff, binary.LittleEndian, t.Height)

	t.SenderPublicKey.WriteToBuff(buff, 64)
	t.RecipientId.WriteToBuff(buff, 64)
	t.Chain.WriteToBuff(buff, 64)
	t.PreviousTx.WriteToBuff(buff, 64)
	return buff.Bytes()
}

// Mine generates a nonce field and automatically fills in the nonce and id fields.
// To make a transaction, you will need to spend a little bit of processing power,
// But it will be so fast that it will be almost unnoticeable,
// because quite low complexity is used.
// Look constant `complexity`.
func (t *Tx) Mine() {
	data := t.GetBytes()
	cm := mine.StartMine(data, complexity, 1)
	nonce := <-cm

	hash := mine.GetHashNonce(t.GetBytes(), nonce)
	t.Id = hash
	t.Nonce = nonce
}

// Save writes a tx to the database.
func (t *Tx) Save(tbl *leveldb.Tbl) (error, Hash) {
	exist, err := tbl.Has(t.Id)

	if err != nil {
		return err, nil
	}

	if exist {
		return ErrTxAlreadyExist, t.Id
	}

	tbl.PutObject(t.Id, t)
	return nil, t.Id
}

// GetTx tries to find a transaction in the entire network by its hash.
func GetTx(hash Hash, tbl *leveldb.Tbl) (error, *Tx) {
	exist, err := tbl.Has(hash)
	if err != nil {
		return err, nil
	}
	if !exist {
		return ErrTxNotFound, nil
	}

	tx := NewTx()
	err = tbl.GetObject(hash, tx)
	if err != nil {
		return err, nil
	}

	return nil, tx
}
