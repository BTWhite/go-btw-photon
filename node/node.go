// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package node

import (
	"github.com/BTWhite/go-btw-photon/chain"
	"github.com/BTWhite/go-btw-photon/config"
	"github.com/BTWhite/go-btw-photon/logger"
	"github.com/BTWhite/go-btw-photon/peer"
	"github.com/BTWhite/go-btw-photon/rpc"
	"github.com/BTWhite/go-btw-photon/rpc/net/http"
)

type Params struct {
	Peers    []peer.Peer `json:"peers"`
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

	logger.Init(n.p.LogLevel)
	chain.LoadGenesis(params.Genesis, cf.ChainHelper())

	n.peers()
	n.delegate()
	n.rpc()
}

func (n *node) delegate() {
	if len(n.p.Delegate) > 0 {
		n.cf.SnapShotFactory().Start()
	}
}

func (n *node) peers() {
	for _, peer := range n.p.Peers {
		err := n.cf.PeerManager().Save(peer)
		if err != nil {
			logger.Err(err.Error())
			continue
		}
	}
}

func (n *node) rpc() {
	rpc.SetConfig(n.cf)
	http.Start(n.p.Port)
}
