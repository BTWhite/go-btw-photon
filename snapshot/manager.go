// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package snapshot

import (
	"encoding/binary"
	"time"

	"github.com/BTWhite/go-btw-photon/crypto/sha256"

	"github.com/BTWhite/go-btw-photon/crypto/sign"
	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/logger"
	"github.com/BTWhite/go-btw-photon/types"
)

var lpm = "SSManager:"

// SnapShotManager controls the work of snapshots, but does not release them
// into the network, see the SnapShotFactory.
type SnapShotManager struct {
	db       *leveldb.Db
	tbl      *leveldb.Tbl
	tblBatch *leveldb.TblBatch
	cur      *SnapShot
}

type lastSnapShot struct {
	Id     types.Hash
	Height uint32
}

// NewSnapShotManager creates manager and empty SnapShot.
func NewSnapShotManager(db *leveldb.Db) *SnapShotManager {

	sm := &SnapShotManager{
		db:       db,
		tbl:      db.CreateTable([]byte("ss")),
		tblBatch: db.NewBatch().CreateTableBatch([]byte("ss")),
		cur:      new(SnapShot),
	}

	ss, err := sm.Get([]byte("unconfirmed"))
	if err != nil && err.Error() != "leveldb: not found" {
		logger.Err(lpm, err.Error())
	}

	if err == nil {
		sm.cur = ss
		logger.Debug(lpm, "Loaded unconfirmed SnapShot, balance changes:", len(ss.Balances))
	}

	return sm
}

// Clear clears SnapShot data for manager.
func (sm *SnapShotManager) Clear() *SnapShotManager {
	sm.cur = new(SnapShot)
	err := sm.Commit()
	if err != nil {
		logger.Err(lpm, "Clear:", err.Error())
	}
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

func (sm *SnapShotManager) AddBalance(b Balance) {
	sm.cur.AddBalance(b)
}

// Release confirms and releases a new snapshot.
func (sm *SnapShotManager) Release(pair *types.KeyPair) (lastSnapShot, error) {
	last, err := sm.Last()
	if err != nil {
		if err != ErrSSNotFound {
			return lastSnapShot{}, err
		}
		sm.cur.Height = 1
	} else {
		sm.cur.Height = last.Height + 1
		sm.cur.PreviousSnapShot = last.Id
	}

	sm.cur.Timestamp = time.Now().Unix()

	sign.Sign(sm.cur, pair, &sm.cur.GeneratorPublicKey, &sm.cur.Signature, 0)
	bts := sm.cur.GetBytes()
	hash := []byte(sha256.Sha256Hex(bts))
	sm.cur.Id = hash

	lss, err := sm.Write(sm.cur)
	if err == nil {
		sm.Clear()
	}
	return lss, err
}

// List gets snapshot list in range.
func (sm *SnapShotManager) List(offset int, limit int) []*SnapShot {
	it := sm.db.NewIteratorPrefix([]byte("ssh-"))
	var result = make([]*SnapShot, limit)

	if offset > 0 {
		var h uint32 = uint32(offset * limit)
		var buf = make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, h)
		it.Seek(append([]byte("ssh-"), buf...))
	}
	j := 0

	for i := 0; it.Next() && i < limit; i++ {
		key := it.Value()

		ss, err := sm.Get(key)
		if err != nil {
			logger.Err(lpm, err.Error(), string(key))
		}
		result[j] = ss
		j++
	}

	return result[:j]
}

// Write writes snapshot to the database.
func (sm *SnapShotManager) Write(ss *SnapShot) (lastSnapShot, error) {
	lss := lastSnapShot{
		Id:     sm.cur.Id,
		Height: sm.cur.Height,
	}

	var buf = make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, ss.Height)
	sm.tblBatch.Put(append([]byte("h-"), buf...), ss.Id)
	sm.tblBatch.PutObject(sm.cur.Id, ss)
	sm.tblBatch.PutObject([]byte("last"), lss)

	return lss, sm.tblBatch.Write()
}

// Last gets last snapshot
func (sm *SnapShotManager) Last() (*SnapShot, error) {
	lss := &lastSnapShot{}
	err := sm.tbl.GetObject([]byte("last"), lss)
	if err != nil {
		return nil, ErrSSNotFound
	}

	ss := new(SnapShot)
	err = sm.tbl.GetObject(lss.Id, ss)
	if err != nil {
		return nil, ErrSSNotFound
	}

	return ss, nil
}

// Commit saves any changes in unconfirmed snapshot.
func (sm *SnapShotManager) Commit() error {
	return sm.tbl.PutObject([]byte("unconfirmed"), sm.cur)
}

// Get finds snapshot in database.
func (sm *SnapShotManager) Get(id []byte) (*SnapShot, error) {
	ss := new(SnapShot)
	err := sm.tbl.GetObject(id, ss)
	if err != nil {
		return nil, err
	}
	return ss, nil
}
