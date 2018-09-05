// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package peer

import (
	"bytes"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb/iterator"

	"github.com/BTWhite/go-btw-photon/db/leveldb"
)

var alpha = []byte("1234567890")

type PeerManager struct {
	db  *leveldb.Db
	tbl *leveldb.Tbl
	it  iterator.Iterator
}

func NewPeerManager(db *leveldb.Db) *PeerManager {
	return &PeerManager{
		db:  db,
		tbl: db.CreateTable([]byte("peer")),
		it:  db.NewIteratorPrefix([]byte("peer")),
	}
}

func (pm *PeerManager) Save(p Peer) error {
	buff := new(bytes.Buffer)
	buff.WriteString(fmt.Sprintf("%s:%d", p.Ip.String(), p.Port))
	return pm.tbl.PutObject(buff.Bytes(), p)
}

func (pm *PeerManager) Random(count int) []Peer {
	var peers []Peer

	for len(peers) < count {
		if !pm.it.Next() {
			pm.it.First()
		}
		p := &Peer{}
		leveldb.Decode(pm.it.Value(), p)
		peers = append(peers, *p)
	}

	return peers
}
