// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package peer

import (
	"fmt"
	"sync"

	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/logger"
)

var alpha = []byte("1234567890")

type PeerManager struct {
	db       *leveldb.Db
	tbl      *leveldb.Tbl
	mu       sync.Mutex
	sLastKey []byte
	wg       sync.WaitGroup
}

func NewPeerManager(db *leveldb.Db) *PeerManager {
	return &PeerManager{
		db:  db,
		tbl: db.CreateTable([]byte("peer")),
	}
}

func (pm *PeerManager) Save(p Peer) error {
	pm.wg.Add(1)
	key := []byte(fmt.Sprintf("%s:%d", p.Ip.String(), p.Port))
	err := pm.tbl.PutObject(key, &p)
	pm.wg.Done()
	return err
}

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

		p := &Peer{}
		err := leveldb.Decode(it.Value(), p)
		if err != nil {
			logger.Err("Peer Manager:", string(it.Key()), err.Error())
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
