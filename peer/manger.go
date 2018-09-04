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
	"encoding/binary"

	"github.com/BTWhite/go-btw-photon/db/leveldb"
)

type PeerManager struct {
	db  *leveldb.Db
	tbl *leveldb.Tbl
}

func NewPeerManager(db *leveldb.Db) *PeerManager {
	return &PeerManager{
		db:  db,
		tbl: db.CreateTable([]byte("peer")),
	}
}

func (pm PeerManager) Save(p Peer) error {
	buff := new(bytes.Buffer)
	buff.WriteString(p.Ip.String())
	binary.Write(buff, binary.LittleEndian, p.Port)
	return pm.tbl.PutObject(buff.Bytes(), p)
}
