// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package account

import (
	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/types"
)

type AccountManager struct {
	db *leveldb.Tbl
}

func NewAccountManager(db *leveldb.Db) *AccountManager {
	return &AccountManager{
		db: db.CreateTable([]byte("usr")),
	}
}

func (am *AccountManager) Save(a *Account) error {
	return am.db.PutObject(a.Address, a)
}

func (am *AccountManager) Get(address types.Hash) *Account {
	acc := NewAccount(address)
	ok, _ := am.db.Has(address)
	if !ok {
		return acc
	}

	am.db.GetObject(acc.Address, acc)
	return acc
}

func (am *AccountManager) GetByPublicKey(pub types.PublicKey) *Account {
	return am.Get([]byte(pub.Address()))
}
