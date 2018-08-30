// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3

package leveldb

type Tbl struct {
	db     *Db
	prefix []byte
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
