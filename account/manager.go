// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package account

import (
	"sync"

	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/types"
)

// AccountManager searches and writes account information.
type AccountManager struct {
	db *leveldb.Tbl
	mu sync.Mutex
	bt *leveldb.TblBatch
}

// NewAccountManager creates a new AccountManager instance.
// The manager works with the database, the connection of which you will give him.
func NewAccountManager(db *leveldb.Db) *AccountManager {
	return &AccountManager{
		db: db.CreateTable([]byte("usr")),
		bt: db.NewBatch().CreateTableBatch([]byte("usr")),
	}
}

// Save overwrites an account in the database
func (am *AccountManager) Save(a *Account) error {
	am.mu.Lock()
	err := am.db.PutObject(a.Address, a)
	am.mu.Unlock()
	return err
}

// Commit commiting all changes from batch.
func (am *AccountManager) Commit() error {
	err := am.bt.Write()
	am.bt.Reset()
	return err
}

// Get finds an account in the database or returns a base account if it was not found.
func (am *AccountManager) Get(address types.Hash) *Account {
	acc := NewAccount(address)
	am.mu.Lock()
	defer am.mu.Unlock()
	ok, _ := am.db.Has(address)
	if !ok {
		return acc
	}

	am.db.GetObject(acc.Address, acc)
	return acc
}

// GetByPublicKey converts a public key to an address and delegates authority
// to the Get method.
func (am *AccountManager) GetByPublicKey(pub types.PublicKey) *Account {
	return am.Get([]byte(pub.Address()))
}
