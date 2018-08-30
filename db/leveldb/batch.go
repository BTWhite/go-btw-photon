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
)

type Batch struct {
	core *leveldb.DB
	b    *leveldb.Batch
	size int
}

func (db *Db) NewBatch() *Batch {
	return &Batch{core: db.core, b: new(leveldb.Batch)}
}

func (b *Batch) Put(key, value []byte) error {
	b.b.Put(key, value)
	b.size += len(value)
	return nil
}

func (b *Batch) Delete(key []byte) error {
	b.b.Delete(key)
	b.size += 1
	return nil
}

func (b *Batch) Write() error {
	return b.core.Write(b.b, nil)
}

func (b *Batch) Reset() {
	b.b.Reset()
	b.size = 0
}

func (b *Batch) PutObject(key []byte, obj interface{}) error {
	bt, err := encode(obj)
	if err != nil {
		return err
	}

	return b.Put(key, bt)
}
