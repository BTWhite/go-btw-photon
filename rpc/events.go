// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package rpc

import (
	"github.com/BTWhite/go-btw-photon/peer"
	"github.com/BTWhite/go-btw-photon/types"
)

type InsufficientDataEvent struct {
	Chain types.Hash
	To    types.Hash
	Peer  *peer.Peer
}

func (e *InsufficientDataEvent) GetObject(obj interface{}) error {
	ro := obj.(*InsufficientDataEvent)
	*ro = *e
	return nil
}

func (e *InsufficientDataEvent) SetData(chain types.Hash, to types.Hash, peer *peer.Peer) {
	e.Chain = chain
	e.To = to
	e.Peer = peer
}

func (e *InsufficientDataEvent) GetBytes() []byte {
	return nil
}

func (e *InsufficientDataEvent) SetBytes([]byte) int {
	return 0
}
