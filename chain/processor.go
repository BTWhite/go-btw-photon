// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package chain

import (
	"encoding/binary"
	"sync"

	"github.com/BTWhite/go-btw-photon/account"
	"github.com/BTWhite/go-btw-photon/crypto/sign"
	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/types"
)

// TxProcessor responsible for the validation of the transaction and its processing.
type TxProcessor interface {
	// Process processes the transaction, but does not write to the chain.
	Process(tx *types.Tx, ch *Chain) error

	// Validate checks the transaction for validity.
	Validate(tx *types.Tx, ch *Chain) error

	// Save saves new tx to the database.
	Save(tx *types.Tx, ch *Chain, tbl *leveldb.Tbl, batch *leveldb.TblBatch) (types.Hash, error)
}

// DefaultProcessor is the base processor for blocks.
type DefaultProcessor struct {
	db *leveldb.Db
	am *account.AccountManager

	mup sync.Mutex
	mus sync.Mutex
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
func (p *DefaultProcessor) Process(tx *types.Tx, ch *Chain) error {
	p.mup.Lock()
	defer p.mup.Unlock()
	recipient := p.am.Get(tx.RecipientId)
	recipient.Balance = types.NewCoin(recipient.Balance.Uint64() + tx.Amount.Uint64())

	// TODO fee to delegates.
	senderPub := types.NewPublicKeyByHex(tx.SenderPublicKey.String())
	sender := p.am.GetByPublicKey(senderPub)
	if senderPub.Address() == tx.SenderId.String() {
		sender.PublicKey = senderPub
	}
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

	return p.am.Commit()
}

// Validate called before the transaction is written, if nil is returned,
// it is considered that the transaction is valid.
func (p *DefaultProcessor) Validate(tx *types.Tx, ch *Chain) error {
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

// Save saves new tx to the database.
func (p *DefaultProcessor) Save(tx *types.Tx, ch *Chain, tbl *leveldb.Tbl,
	batch *leveldb.TblBatch) (types.Hash, error) {

	p.mus.Lock()
	defer p.mus.Unlock()

	exist, err := tbl.Has(tx.Id)
	if err != nil {
		return nil, err
	}

	if exist {
		return tx.Id, ErrTxAlreadyExist
	}

	err = batch.PutObject(tx.Id, tx)
	if err != nil {
		return nil, err
	}

	bNum := uint32ToBytes(lastTxNum(tbl) + 1)
	err = batch.Put(bNum, tx.Id)
	if err != nil {
		return nil, err
	}

	return tx.Id, batch.Write()
}

func uint32ToBytes(u uint32) []byte {
	buff := make([]byte, 4)
	binary.LittleEndian.PutUint32(buff, u)
	return buff
}

func bytesToUint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

func lastTxNum(tbl *leveldb.Tbl) uint32 {
	b, err := tbl.Get([]byte("lastNum"))
	if err != nil {
		return 0
	}

	return bytesToUint32(b)
}
