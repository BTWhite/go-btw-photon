// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3

package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
)

// NewIterator creates new iterator instance.
func (db *Db) NewIterator() iterator.Iterator {
	return db.core.NewIterator(nil, nil)
}

// NewIteratorPrefix creates new iterator instance with prefix.
func (db *Db) NewIteratorPrefix(prefix []byte) iterator.Iterator {
	return db.core.NewIterator(util.BytesPrefix(prefix), nil)
}
