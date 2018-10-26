// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3

package rawdb

import (
	"errors"

	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/types"
	"github.com/BTWhite/go-btw-photon/utils"
)

var (
	// ErrTxAlreadyExist is returned if tx already exist in tx list.
	ErrTxAlreadyExist = errors.New("Tx already exist")

	// ErrTxNotFound is returned if tx not found.
	ErrTxNotFound = errors.New("Tx not found")
)

// WriteTx writes tx to the database with indexes
func WriteTx(w leveldb.Writer, tx *types.Tx) error {

	err := w.PutObject(tx.Id, tx)
	if err != nil {
		return err
	}

	bNum := utils.Uint32ToBytes(tx.Height)
	err = w.Put(bNum, tx.Id)
	if err != nil {
		return err
	}

	return nil
}

// GetTx tries to find a transaction in the entire network by its hash.
func GetTxByHash(r leveldb.Reader, hash types.Hash) (*types.Tx, error) {

	tx := types.NewTx()
	err := r.GetObject(hash, tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

//func GetTxByHeight(r leveldb.Reader, height uint32,
//	chain types.Hash) (*types.Tx, error) {

//	key := utils.Uint32ToBytes(height)
//	exist, err := r.Has(key)
//	if err != nil {
//		return nil, err
//	}
//	if !exist {
//		return nil, ErrTxNotFound
//	}

//	hash, err := r.Get(key)
//	return
//}

func lastTxNum(r leveldb.Reader) uint32 {
	b, err := r.Get([]byte("lastNum"))
	if err != nil {
		return 0
	}

	return utils.BytesToUint32(b)
}
