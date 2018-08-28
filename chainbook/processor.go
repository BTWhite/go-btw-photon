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

	"github.com/BTWhite/go-btw-photon/account"
	"github.com/BTWhite/go-btw-photon/chain"
	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/sign"
	"github.com/BTWhite/go-btw-photon/types"
)

var (
	// ErrTxInvalidPrevTx is returned if the transaction specified an invalid
	// previous transaction hash.
	ErrTxInvalidPrevTx = errors.New("Invalid previous tx")

	// ErrTxInvalidSignature is returned if signature incorrect.
	ErrTxInvalidSignature = errors.New("Invalid tx signature")

	// ErrTxInsufficientBalance is returned is sender does not have enough coins.
	ErrTxInsufficientBalance = errors.New("Insufficient balance")
)

// DefaultProcessor is the base processor for blocks.
type DefaultProcessor struct {
	db *leveldb.Db
	am *account.AccountManager
}

// NewProcessor creates a new DefaultProcessor.
func NewProcessor(db *leveldb.Db) *DefaultProcessor {
	dp := &DefaultProcessor{}
	dp.db = db
	dp.am = account.NewAccountManager(db)
	return dp
}

// Process called directly for transaction processing.
// Do not use this method to write to the chain, here only the results are processed.
func (p *DefaultProcessor) Process(tx *types.Tx, ch *chain.Chain) error {
	recipient := p.am.Get(tx.RecipientId)
	recipient.Balance = types.NewCoin(recipient.Balance.Uint64() + tx.Amount.Uint64())

	// TODO fee to delegates.
	sender := p.am.GetByPublicKey(types.NewPublicKeyByHex(tx.SenderPublicKey.String()))
	balance := sender.Balance.Uint64() - (tx.Amount.Uint64() + tx.Fee.Uint64())
	sender.Balance = types.NewCoin(balance)

	err := p.am.Save(recipient)
	if err != nil {
		return err
	}

	err = p.am.Save(sender)
	if err != nil {
		return err
	}

	return nil
}

// Validate called before the transaction is written, if nil is returned,
// it is considered that the transaction is valid.
func (p *DefaultProcessor) Validate(tx *types.Tx, ch *chain.Chain) error {
	if tx.Chain.Equals(genesisChain) && tx.Height == 1 {
		return nil
	}

	if !tx.PreviousTx.Equals(ch.LastTx()) {
		return ErrTxInvalidPrevTx
	}

	if !sign.Verify(tx, tx.SenderPublicKey, tx.Signature, 0) {
		return ErrTxInvalidSignature
	}

	sender := p.am.GetByPublicKey(types.NewPublicKeyByHex(tx.SenderPublicKey.String()))
	if sender.Balance.Uint64() < (tx.Amount.Uint64() + tx.Fee.Uint64()) {
		return ErrTxInsufficientBalance
	}

	return nil
}
