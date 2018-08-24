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
)

// Tx determines the structure of the transaction.
type Tx struct {
	Id              Hash   `json:"id"`
	SenderPublicKey Hash   `json:"senderId"`
	RecipientId     Hash   `json:"recipientId"`
	Amount          uint64 `json:"amount"`
	Fee             uint64 `json:"fee"`
	Signature       Hash   `json:"signature"`
	Timestamp       uint32 `json:"timestamp"`
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

	t.SenderPublicKey.WriteToBuff(buff, 32)
	t.RecipientId.WriteToBuff(buff, 32)
	return buff.Bytes()
}
