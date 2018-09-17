// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package node

import (
	"net"

	"github.com/BTWhite/go-btw-photon/chain"
	"github.com/BTWhite/go-btw-photon/config"
	"github.com/BTWhite/go-btw-photon/logger"
	"github.com/BTWhite/go-btw-photon/peer"
	"github.com/BTWhite/go-btw-photon/rpc"
	"github.com/BTWhite/go-btw-photon/rpc/net/http"
	"github.com/BTWhite/go-btw-photon/sync"
	"github.com/BTWhite/go-btw-photon/types"
)

type Params struct {
	Peers    []peer.Peer `json:"peers"`
	Ip       net.IP      `json:"ip"`
	Port     int         `json:"port"`
	Genesis  string      `json:"genesis"`
	Delegate string      `json:"delegate"`
	Magic    string      `json:"magic"`
	LogLevel string      `json:"logLevel"`
}

type node struct {
	cf *config.Config
	p  Params
}

func StartNode(cf *config.Config, params Params) {
	n := node{
		cf: cf,
		p:  params,
	}

	lp := peer.NewPeer(n.p.Ip, n.p.Port)
	peer.SetLocalPeer(&lp)

	logger.Init(n.p.LogLevel)
	go n.rpc()
	n.peers()
	n.snapshots()
	n.txs()

	cf.PeerManager().DisablerStart()
	err := chain.LoadGenesis(params.Genesis, cf.ChainHelper())
	if err != nil {
		logger.Err(err)
	}
}

func (n *node) snapshots() {
	if len(n.p.Delegate) > 0 {
		sf := n.cf.SnapShotFactory()
		sf.SetDelegate(types.NewKeyPair([]byte(n.p.Delegate)))
		sf.Start()
	}

	s := &sync.SnapShotSyncer{}
	s.SetConfig(n.cf)
	s.Start()
}

func (n *node) txs() {

	s := &sync.ChainSyncer{}
	s.SetConfig(n.cf)
	s.Start()
}

func (n *node) peers() {
	for _, peer := range n.p.Peers {
		err := n.cf.PeerManager().Save(peer)
		if err != nil {
			logger.Err(err.Error())
			continue
		}
	}

	s := &sync.PeerSyncer{}
	s.SetConfig(n.cf)
	s.Start()
}

func (n *node) rpc() {
	rpc.SetConfig(n.cf)
	http.Start(n.p.Port)
}
