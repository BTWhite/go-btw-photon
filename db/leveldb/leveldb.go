// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3

package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

// Db is wrapper over the standard leveldb.DB.
type Db struct {
	core *leveldb.DB
}

// Open opens a new connection to the database.
func Open(filepath string) *Db {
	var (
		cache   = 512
		handles = 512
	)
	opts := &opt.Options{
		OpenFilesCacheCapacity: handles,
		BlockCacheCapacity:     cache / 2 * opt.MiB,
		WriteBuffer:            cache / 4 * opt.MiB, // Two of these are used internally
		Filter:                 filter.NewBloomFilter(10),
	}
	db, err := leveldb.OpenFile(filepath, opts)

	if err != nil {
		panic(err.Error())
	}

	return &Db{
		core: db,
	}
}

// Close closes connection to the database.
func (db *Db) Close() {
	db.core.Close()
}

// Put puts bytes to the database.
func (db *Db) Put(key []byte, value []byte) error {
	return db.core.Put(key, value, nil)
}

// PutObject puts object to the database.
func (db *Db) PutObject(key []byte, obj interface{}) error {
	b, err := encode(obj)
	if err != nil {
		return err
	}
	return db.Put(key, b)
}

// Get gets bytes from the database.
func (db *Db) Get(key []byte) ([]byte, error) {
	return db.core.Get(key, nil)
}

// GetObject gets object from the database.
func (db *Db) GetObject(key []byte, obj interface{}) error {
	tmp, err := db.Get(key)
	if err != nil {
		return err
	}

	return Decode(tmp, obj)
}

// Delete deletes value by key from the database.
func (db *Db) Delete(key []byte) error {
	return db.core.Delete(key, nil)
}

// Hash checks the presence of an element in the database.
func (db *Db) Has(key []byte) (bool, error) {
	return db.core.Has(key, nil)
}
