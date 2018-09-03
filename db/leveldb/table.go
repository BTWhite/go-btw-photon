// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3

package leveldb

// Tbl will in fact add a prefix to any records.
type Tbl struct {
	db     *Db
	prefix []byte
}

// CreateTable creates new table instance.
func (db *Db) CreateTable(prefix []byte) *Tbl {
	return &Tbl{
		db:     db,
		prefix: prefix,
	}
}

// Put puts bytes to the table.
func (t *Tbl) Put(key []byte, value []byte) error {
	return t.db.Put(prefix(t.prefix, key), value)
}

// PutObject puts object to the table.
func (t *Tbl) PutObject(key []byte, obj interface{}) error {

	return t.db.PutObject(prefix(t.prefix, key), obj)
}

// Get gets bytes from the table.
func (t *Tbl) Get(key []byte) ([]byte, error) {
	return t.db.Get(prefix(t.prefix, key))
}

// GetObject gets object from the table.
func (t *Tbl) GetObject(key []byte, obj interface{}) error {

	return t.db.GetObject(prefix(t.prefix, key), obj)
}

// Has checks the presence of an element in the table.
func (t *Tbl) Has(key []byte) (bool, error) {

	return t.db.Has(prefix(t.prefix, key))
}
