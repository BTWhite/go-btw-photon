// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package peer

import (
	"sync"

	"github.com/BTWhite/go-btw-photon/events"

	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/logger"
)

var alpha = []byte("1234567890")

// PeerManager controls the behavior of peers and establishes a relationship
// with the database.
type PeerManager struct {
	db       *leveldb.Db
	tbl      *leveldb.Tbl
	mu       sync.Mutex
	sLastKey []byte
	wg       sync.WaitGroup
}

// NewPeerManager creates new peer manager.
func NewPeerManager(db *leveldb.Db) *PeerManager {
	return &PeerManager{
		db:  db,
		tbl: db.CreateTable([]byte("peer")),
	}
}

// DisablerStart starts cycle for disable peers who stopped responding.
func (pm *PeerManager) DisablerStart() {
	go pm.disableCycle(events.Subscribe("peer-noconn"))
}

func (pm *PeerManager) Save(p Peer) error {
	pm.wg.Add(1)
	key := p.DBKey()
	err := pm.tbl.PutObject(key, &p)
	pm.wg.Done()
	return err
}

func (pm *PeerManager) Exist(p Peer) bool {
	pm.wg.Add(1)
	key := p.DBKey()
	exist, err := pm.tbl.Has(key)
	pm.wg.Done()
	if err != nil {
		logger.Err(err.Error())
		return false
	}
	return exist
}

// Disable sets peer inactive.
func (pm *PeerManager) Disable(p Peer) error {
	pm.wg.Add(1)
	bt := pm.db.NewBatch()
	key := p.DBKey()
	bt.Delete(append([]byte("peer"), key...))
	bt.PutObject(append([]byte("dsbl-peer"), key...), &p)
	err := bt.Write()
	pm.wg.Done()
	return err
}

// Random gets count random peers.
func (pm *PeerManager) Random(count int) []Peer {
	it := pm.db.NewIteratorPrefix([]byte("peer"))

	pm.wg.Wait()
	var peers []Peer
	tmp := make(map[string]bool)

	pm.mu.Lock()
	if len(pm.sLastKey) > 0 {
		it.Seek(pm.sLastKey)
	}

	for len(peers) < count {
		if !it.Next() {
			it.First()
		}
		if !it.Valid() {
			count--
			continue
		}
		p := &Peer{}
		err := leveldb.Decode(it.Value(), p)
		if err != nil {
			logger.Err("Peer Manager:", err.Error())
			count--
			continue
		}
		httpAddr := p.HttpAddr()
		if tmp[httpAddr] == true {
			count--
			continue
		}

		pm.sLastKey = it.Key()
		peers = append(peers, *p)
		tmp[httpAddr] = true
	}
	pm.mu.Unlock()

	return peers
}

func (pm *PeerManager) disableCycle(e chan events.Eventer) {
	for true {
		ev := <-e
		key := ev.GetBytes()
		p := Peer{}
		pm.tbl.GetObject(key, &p)

		pm.Disable(p)
	}
}
