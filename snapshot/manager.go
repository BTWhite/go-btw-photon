// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package snapshot

import (
	"time"

	"github.com/BTWhite/go-btw-photon/crypto/sign"
	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/types"
)

// SnapShotManager controls the work of snapshots, but does not release them
// into the network, see the SnapShotFactory.
type SnapShotManager struct {
	db  *leveldb.Db
	tbl *leveldb.Tbl
	cur *SnapShot
}

// NewSnapShotManager creates manager and empty SnapShot.
func NewSnapShotManager(db *leveldb.Db) *SnapShotManager {

	sm := &SnapShotManager{
		db:  db,
		tbl: db.CreateTable([]byte("ss")),
	}
	sm.Clear()

	return sm
}

// Clear clears SnapShot data for manager.
func (sm *SnapShotManager) Clear() *SnapShotManager {
	sm.cur = new(SnapShot)
	return sm
}

// AddVote creates and vote.
func (sm *SnapShotManager) CreateVote(kp *types.KeyPair,
	delegate types.Hash) types.Vote {

	vote := types.Vote{
		Delegate:  delegate,
		Timestamp: time.Now().Unix(),
	}

	sign.Sign(vote, kp, &vote.Voter, &vote.Signature, 0)
	return vote
}

func (sm *SnapShotManager) AddVote(v types.Vote) {
	sm.cur.AddVote(v)
}

//func (sm *SnapShotManager) Save(ss *SnapShotManager)
