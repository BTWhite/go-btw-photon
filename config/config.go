// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package config

import (
	"sync"

	"github.com/BTWhite/go-btw-photon/account"
	"github.com/BTWhite/go-btw-photon/chain"
	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/snapshot"
)

// Configer is a network configuration view.
type Configer interface {
	// SnapShotManager is getter for SnapShotManager.
	SnapShotManager() *snapshot.SnapShotManager

	// AccountManager is getter for AccountManager.
	AccountManager() *account.AccountManager

	// ChainHelper is getter for ChainHelper.
	ChainHelper() *chain.ChainHelper

	// SnapShotFactory is getter for SnapShotFactory.
	SnapShotFactory() *snapshot.SnapShotFactory

	//  Magic is getter for magic value of the network.
	Magic() []byte

	// Version is getter for Version instance.
	Version() Version
}

// Config combines the database, version and magic value of the network.
// It can also store the Accounts, Chains and Snapshots Managers.
type Config struct {
	db      *leveldb.Db
	mu      sync.Mutex
	am      *account.AccountManager
	ch      *chain.ChainHelper
	sm      *snapshot.SnapShotManager
	sf      *snapshot.SnapShotFactory
	magic   []byte
	version Version
}

// NewConfig creates new config, but not creates Managers.
func NewConfig(db *leveldb.Db, magic []byte, version [3]byte) *Config {

	return &Config{
		db:      db,
		magic:   magic,
		version: version,
	}
}

// SnapShotManager is getter for SnapShotManager, creates default if manager nil.
func (c *Config) SnapShotManager() *snapshot.SnapShotManager {
	if c.sm == nil {
		c.sm = snapshot.NewSnapShotManager(c.db)
	}

	return c.sm
}

// AccountManager is getter for AccountManager, creates default if manager nil.
func (c *Config) AccountManager() *account.AccountManager {
	if c.am == nil {
		c.am = account.NewAccountManager(c.db)
	}

	return c.am
}

// ChainHelper is getter for ChainHelper, creates default if helper nil.
func (c *Config) ChainHelper() *chain.ChainHelper {
	if c.ch == nil {
		c.ch = chain.NewChainHelper(c.db)
	}

	return c.ch
}

// SnapShotFactory is getter for SnapShotFactory, creates default if factory nil.
func (c *Config) SnapShotFactory() *snapshot.SnapShotFactory {
	if c.sf == nil {
		c.sf = snapshot.NewSnapShotFactory(c.SnapShotManager(), c.AccountManager(), c.ChainHelper(), c.db)
	}

	return c.sf
}

// DataBase is getter for Db instance.
func (c *Config) DataBase() *leveldb.Db {
	return c.db
}

// Magic is getter for magic value of the network.
func (c *Config) Magic() []byte {
	return c.magic
}

// Version is getter for Version instance.
func (c *Config) Version() Version {
	return c.version
}
