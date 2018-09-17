// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package snapshot

import (
	"sync"
	"time"

	"github.com/BTWhite/go-btw-photon/types"

	"github.com/BTWhite/go-btw-photon/db/leveldb"

	"github.com/BTWhite/go-btw-photon/account"
	"github.com/BTWhite/go-btw-photon/chain"
	"github.com/BTWhite/go-btw-photon/events"
	"github.com/BTWhite/go-btw-photon/logger"
)

var lpf = "SSFactory:"

// SnapShotFactory is releasing new snapshots for delegates.
type SnapShotFactory struct {
	sm       *SnapShotManager
	am       *account.AccountManager
	ch       *chain.ChainHelper
	interval time.Duration
	mu       sync.Mutex
	running  bool
	delegate *types.KeyPair
}

// NewSnapShotFactory creates new factory.
// Please note that the factory needs a manager.
func NewSnapShotFactory(sm *SnapShotManager,
	am *account.AccountManager, ch *chain.ChainHelper, db *leveldb.Db) *SnapShotFactory {

	return &SnapShotFactory{
		sm: sm,
		am: am,
		ch: ch,
	}
}

// SetDelegate sets delegate account for release new snapshots.
func (sf *SnapShotFactory) SetDelegate(kp *types.KeyPair) {
	sf.mu.Lock()
	sf.delegate = kp
	sf.mu.Unlock()
}

// Start starting the factory process.
func (sf *SnapShotFactory) Start() {
	sf.mu.Lock()
	defer sf.mu.Unlock()
	if sf.running {
		return
	}
	sf.running = true

	cTx := events.Subscribe("newtx")
	go sf.cycle(cTx)
	go sf.releaser()
}

// Start stopping the factory process.
func (sf *SnapShotFactory) Stop() {
	sf.mu.Lock()
	sf.running = false
	sf.mu.Unlock()
}

func (sf *SnapShotFactory) releaser() {
	for true {
		// TODO
		time.Sleep(time.Second * 10)
		sf.mu.Lock()
		delegate := sf.delegate
		sf.mu.Unlock()
		ss, err := sf.sm.Release(delegate)
		if err != nil {
			logger.Err(lpf, err.Error())
			continue
		}

		logger.Info(lpf, "Produced new snapshot:", ss.Id, "h:", ss.Height)
	}
}

func (sf *SnapShotFactory) cycle(cTx chan events.Eventer) {
	for sf.running {
		txE := <-cTx
		hash := txE.GetBytes()
		tx, err := sf.ch.GetTx(hash)
		if err != nil {
			logger.Err(err.Error())
			continue
		}
		accR := sf.am.Get(tx.RecipientId)
		accS := sf.am.Get(tx.SenderId)

		sf.sm.AddBalance(BalanceByAccount(accR))
		sf.sm.AddBalance(BalanceByAccount(accS))
		err = sf.sm.Commit()
		if err != nil {
			logger.Err(lpf, err.Error())
		}
		logger.Debug(lpf, "Add balances:", accR.Address, "->", accR.Balance, "and", accS.Address, "->", accS.Balance)
	}
}
