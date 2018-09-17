// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3

package leveldb

// Batcher is a general representation of the batch, which can both tablebatch and batch.
type Batcher interface {
	// Put puts bytes to the batch.
	Put(key, value []byte) error

	// PutObject puts object to the batch.
	PutObject(key []byte, obj interface{}) error

	// Delete deletes value by key from the batch.
	Delete(key []byte) error

	// Write writes data from batch to the database
	Write() error

	// Reset resets batch.
	Reset()

	// CreateTableBatch creates new tablebatch instance.
	CreateTableBatch(prefix []byte) Batcher
}

// Writer is a database object with write access.
type Writer interface {
	// Put puts bytes to the database.
	Put(key, value []byte) error

	// PutObject puts object to the database.
	PutObject(key []byte, obj interface{}) error

	// Delete deletes value by key from the database.
	Delete(key []byte) error

	// Has checks the presence of an element in the database.
	Has(key []byte) (bool, error)
}

// Reader is a database object with read access.
type Reader interface {
	// Get gets bytes from the database.
	Get(key []byte) ([]byte, error)

	// GetObject gets object from the database.
	GetObject(key []byte, obj interface{}) error

	// Has checks the presence of an element in the database.
	Has(key []byte) (bool, error)
}

// Corrector is a combination of Reader and Writer.
type Corrector interface {
	// Put puts bytes to the database.
	Put(key, value []byte) error

	// PutObject puts object to the database.
	PutObject(key []byte, obj interface{}) error

	// Delete deletes value by key from the database.
	Delete(key []byte) error

	// Get gets bytes from the database.
	Get(key []byte) ([]byte, error)

	// GetObject gets object from the database.
	GetObject(key []byte, obj interface{}) error

	// Has checks the presence of an element in the database.
	Has(key []byte) (bool, error)
}
