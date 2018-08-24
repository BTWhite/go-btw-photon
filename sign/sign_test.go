// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package sign

import (
	"testing"

	"github.com/BTWhite/go-btw-photon/types"
)

func TestSign(t *testing.T) {
	tx := getSignedTx()
	if len(tx.Signature) == 0 {
		t.Error("Signing not work")
	}
}

func TestVerify(t *testing.T) {
	tx := getSignedTx()
	if !Verify(tx, tx.SenderPublicKey, tx.Signature, 0) {
		t.Error("Incorrect sign/verify")
	}
}

func getSignedTx() *types.Tx {
	tx := types.NewTx()
	tx.Amount = 100000
	tx.Fee = 10000
	tx.RecipientId = types.NewHash([]byte("51003903900029004900"))
	tx.SenderPublicKey = types.NewHash([]byte("f41dbb2329cdac8d28e3eac05fede2fe490d9221f83f21666bf5daa8eeb05fd8"))
	tx.GetBytes()

	kp := types.NewKeyPair([]byte("Hello World"))
	Sign(tx, kp, &tx.SenderPublicKey, &tx.Signature, 0)

	return tx
}
