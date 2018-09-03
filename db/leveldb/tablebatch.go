// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3

package leveldb

// TblBatch it's the same as batch, but with a prefix.
type TblBatch struct {
	b      *Batch
	prefix []byte
}

// CreateTable creates new tablebatch instance.
func (b *Batch) CreateTableBatch(prefix []byte) *TblBatch {
	return &TblBatch{
		b:      b,
		prefix: prefix,
	}
}

// Put puts bytes to the tablebatch.
func (t *TblBatch) Put(key []byte, value []byte) error {
	return t.b.Put(prefix(t.prefix, key), value)
}

// PutObject puts object to the tablebatch.
func (t *TblBatch) PutObject(key []byte, obj interface{}) error {
	return t.b.PutObject(prefix(t.prefix, key), obj)
}

// Delete deletes value by key from the tablebatch.
func (t *TblBatch) Delete(key []byte) error {
	return t.b.Delete(prefix(t.prefix, key))
}

// Write writes data from batch to the database.
func (t *TblBatch) Write() error {
	return t.b.Write()
}

// Reset resets tablebatch.
func (t *TblBatch) Reset() {
	t.b.Reset()
}
