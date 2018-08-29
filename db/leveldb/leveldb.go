// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3

package leveldb

import (
	"bytes"
	"encoding/gob"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type Tbl struct {
	db     *Db
	prefix []byte
}

type Db struct {
	core *leveldb.DB
}

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

func (db *Db) Close() {
	db.core.Close()
}

func (db *Db) Put(key []byte, value []byte) error {
	return db.core.Put(key, value, nil)
}

func (db *Db) PutObject(key []byte, obj interface{}) error {
	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	if err := e.Encode(obj); err != nil {
		panic(err)
	}

	return db.Put(key, b.Bytes())
}

func (db *Db) Get(key []byte) ([]byte, error) {
	return db.core.Get(key, nil)
}

func (db *Db) GetObject(key []byte, obj interface{}) error {
	tmp, err := db.Get(key)
	if err != nil {
		obj = nil
		return err
	}

	var b bytes.Buffer
	b.Write(tmp)
	d := gob.NewDecoder(&b)
	return d.Decode(obj)
}

func (db *Db) Delete(key []byte) error {
	return db.core.Delete(key, nil)
}

func (db *Db) Has(key []byte) (bool, error) {
	return db.core.Has(key, nil)
}

func (db *Db) CreateTable(prefix []byte) *Tbl {
	return &Tbl{
		db:     db,
		prefix: prefix,
	}
}

func (t *Tbl) Put(key []byte, value []byte) error {
	return t.db.Put(append(t.prefix, key...), value)
}

func (t *Tbl) PutObject(key []byte, obj interface{}) error {

	return t.db.PutObject(append(t.prefix, key...), obj)
}

func (t *Tbl) Get(key []byte) ([]byte, error) {
	return t.db.Get(append(t.prefix, key...))
}

func (t *Tbl) GetObject(key []byte, obj interface{}) error {

	return t.db.GetObject(append(t.prefix, key...), obj)
}

func (t *Tbl) Has(key []byte) (bool, error) {

	return t.db.Has(append(t.prefix, key...))
}
