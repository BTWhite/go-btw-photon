// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package sync

import (
	"errors"
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
	ch        *chain.ChainHelper
	pm        *peer.PeerManager
	mu        sync.Mutex
	running   bool
	eventTx   chan events.Eventer
	eventData chan events.Eventer
}

func (s *ChainSyncer) Start() {
	logger.Debug(lpc, "Starting...")
	s.eventTx = events.Subscribe("newtx")
	s.eventData = events.Subscribe("insufficent_data_tx")
	s.running = true

	go s.cycleTx()
	go s.cycleData()
}

func (s *ChainSyncer) Stop() {
	s.running = false
}

func (s *ChainSyncer) SetConfig(cf *config.Config) {
	s.ch = cf.ChainHelper()
	s.pm = cf.PeerManager()
}

func (s *ChainSyncer) cycleTx() {
	for true {
		txh := <-s.eventTx
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

func (s *ChainSyncer) cycleData() {
	for true {
		e := <-s.eventData
		if !s.running {
			break
		}
		re := &rpc.InsufficientDataEvent{}
		e.GetObject(re)

		if re.Peer != nil {

			for i := 0; true; i++ {
				end, err := s.loadChain(re.Peer.HttpAddr(), 20*i, 20, re.Chain, re.To)
				if err != nil {
					logger.Err(lpc, err.Error())
					break
				}

				if end {
					break
				}
			}

		}

	}
}

func (s *ChainSyncer) loadChain(addr string, start int, limit int,
	ch types.Hash, to types.Hash) (bool, error) {
	re := rpc.Request{}
	re.Method = "chain.load"
	re.Params = rpc.LoadChainRequest{
		Chain: ch.String(),
		Start: start,
		Limit: limit,
	}

	respArgs := rpc.LoadChainResponse{}

	r, err := http.Send(addr, re, &respArgs)

	if err != nil {
		return false, err
	}

	if r.Error != nil {
		return false, errors.New(r.Error.Message())
	}

	for _, tx := range respArgs.Txs {
		s.mu.Lock()
		err := s.ch.ProcessTx(tx)

		if err != nil && err != chain.ErrTxAlreadyExist {
			s.mu.Unlock()
			return false, err
		}

		if tx.Id.Equals(to) {
			s.mu.Unlock()
			return true, nil
		}
		s.mu.Unlock()
	}

	return false, nil
}

func (s *ChainSyncer) syncTx(tx *types.Tx) {
	s.mu.Lock()
	r := rpc.Request{
		Id:     0,
		Method: "tx.post",
		Params: tx,
	}
	http.BroadCast(s.pm, r, nil, 0)
	s.mu.Unlock()
}
