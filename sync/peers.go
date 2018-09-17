// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package sync

import (
	"time"

	"github.com/BTWhite/go-btw-photon/rpc"

	"github.com/BTWhite/go-btw-photon/config"
	"github.com/BTWhite/go-btw-photon/logger"
	"github.com/BTWhite/go-btw-photon/peer"
	"github.com/BTWhite/go-btw-photon/rpc/net/http"
)

var lpp = "Peers sync:"

type PeerSyncer struct {
	pm      *peer.PeerManager
	running bool
}

func (s *PeerSyncer) Start() {
	logger.Debug(lpp, "Starting...")
	s.running = true
	go s.cycle()
}

func (s *PeerSyncer) cycle() {
	for s.running {
		time.Sleep(time.Second * 5)
		peers := s.pm.Random(1)
		if len(peers) == 0 {
			continue
		}
		req := rpc.Request{
			Method: "peers.get",
			Params: rpc.GetPeersRequest{
				Limit: 1,
			},
		}
		var resp []peer.Peer
		http.Send(peers[0], req, &resp)
		for _, p := range resp {
			if !s.pm.Exist(p) {
				s.pm.Save(p)
			}
		}
	}
}

func (s *PeerSyncer) Stop() {

}

func (s *PeerSyncer) SetConfig(cf *config.Config) {
	s.pm = cf.PeerManager()
}
