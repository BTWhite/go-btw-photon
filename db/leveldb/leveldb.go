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
)

type Tbl struct {
	prefix []byte
}

var inst *leveldb.DB

func Open(filepath string) {
	if inst != nil {
		Close()
	}

	db, err := leveldb.OpenFile(filepath, nil)

	if err != nil {
		panic(err.Error())
	}

	inst = db
}

func Close() {
	inst.Close()
}

func Put(key []byte, value []byte) error {
	return inst.Put(key, value, nil)
}

func PutObject(key []byte, obj interface{}) error {
	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	if err := e.Encode(obj); err != nil {
		panic(err)
	}

	return Put(key, b.Bytes())
}

func Get(key []byte) ([]byte, error) {
	return inst.Get(key, nil)
}

func GetObject(key []byte, obj interface{}) error {
	tmp, err := Get(key)
	if err != nil {
		obj = nil
		return err
	}

	var b bytes.Buffer
	b.Write(tmp)
	d := gob.NewDecoder(&b)
	return d.Decode(obj)
}

func Delete(key []byte) error {
	return inst.Delete(key, nil)
}

func Has(key []byte) (bool, error) {
	return inst.Has(key, nil)
}

func CreateTable(prefix []byte) *Tbl {
	return &Tbl{prefix}
}

func (t *Tbl) Put(key []byte, value []byte) error {
	return Put(append(t.prefix, key...), value)
}

func (t *Tbl) PutObject(key []byte, obj interface{}) error {

	return PutObject(append(t.prefix, key...), obj)
}

func (t *Tbl) Get(key []byte) ([]byte, error) {
	return Get(append(t.prefix, key...))
}

func (t *Tbl) GetObject(key []byte, obj interface{}) error {

	return GetObject(append(t.prefix, key...), obj)
}

func (t *Tbl) Has(key []byte) (bool, error) {

	return Has(append(t.prefix, key...))
}
