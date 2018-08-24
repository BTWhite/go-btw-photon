package tests

import (
	"testing"

	"github.com/BTWhite/go-btw-photon/types"
	"github.com/BTWhite/go-btw-photon/utils"
)

func TestSign(t *testing.T) {
	tx := getSignedTx()
	if len(tx.Signature) == 0 {
		t.Error("Signing not work")
	}
}

func TestVerify(t *testing.T) {
	tx := getSignedTx()
	if !utils.Verify(tx, tx.SenderPublicKey, tx.Signature, 0) {
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
	utils.Sign(tx, kp, &tx.SenderPublicKey, &tx.Signature, 0)

	return tx
}
