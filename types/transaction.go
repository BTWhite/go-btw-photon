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

	"github.com/BTWhite/go-btw-photon/crypto/sha256"
	"github.com/BTWhite/go-btw-photon/mine"
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
//
// Deprecated: Use GenerateId.
func (t *Tx) Mine() {
	data := t.GetBytes()
	cm := mine.StartMine(data, complexity, 1)
	nonce := <-cm

	hash := mine.GetHashNonce(t.GetBytes(), nonce)
	t.Id = hash
	t.Nonce = nonce
}

// GenerateId generatates id and fills in Id field.
func (t *Tx) GenerateId() {
	data := t.GetBytes()
	hash := []byte(sha256.Sha256Hex(data))
	t.Id = hash
}
