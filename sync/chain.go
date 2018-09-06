// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package sync

import (
	"sync"

	"github.com/BTWhite/go-btw-photon/rpc"

	"github.com/BTWhite/go-btw-photon/events"
	"github.com/BTWhite/go-btw-photon/rpc/net/http"
	"github.com/BTWhite/go-btw-photon/types"

	"github.com/BTWhite/go-btw-photon/chain"
	"github.com/BTWhite/go-btw-photon/config"
	"github.com/BTWhite/go-btw-photon/logger"
	"github.com/BTWhite/go-btw-photon/peer"
)

var lpc = "Chain sync:"

type ChainSyncer struct {
	ch      *chain.ChainHelper
	pm      *peer.PeerManager
	mu      sync.Mutex
	running bool
	event   chan *events.Event
}

func (s *ChainSyncer) Start() {
	logger.Debug(lpc, "Starting...")
	s.event = events.Subscribe("newtx")
	s.running = true
	go s.cycle()
}

func (s *ChainSyncer) Stop() {
	s.running = false
}

func (s *ChainSyncer) SetConfig(cf *config.Config) {
	s.ch = cf.ChainHelper()
	s.pm = cf.PeerManager()
}

func (s *ChainSyncer) cycle() {
	for true {
		txh := <-s.event
		if !s.running {
			break
		}
		hash := txh.GetBytes()
		tx, err := s.ch.GetTx(hash)
		if err != nil {
			logger.Err(lpc, err.Error())
			continue
		}
		s.syncTx(tx)
	}
}

func (s *ChainSyncer) syncTx(tx *types.Tx) {

	r := rpc.Request{
		Id:     0,
		Method: "tx.post",
		Params: tx,
	}
	http.BroadCast(s.pm, r, nil, 0)
}
