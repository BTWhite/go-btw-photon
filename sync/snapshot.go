// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package sync

import (
	"github.com/BTWhite/go-btw-photon/config"
	"github.com/BTWhite/go-btw-photon/logger"
	"github.com/BTWhite/go-btw-photon/peer"
	"github.com/BTWhite/go-btw-photon/snapshot"
)

var lps = "Snapshots sync:"

type SnapShotSyncer struct {
	sf *snapshot.SnapShotFactory
	sm *snapshot.SnapShotManager
	pm *peer.PeerManager
}

func (s *SnapShotSyncer) Start() {
	logger.Debug(lps, "Starting...")
}

func (s *SnapShotSyncer) Stop() {

}

func (s *SnapShotSyncer) SetConfig(cf *config.Config) {
	s.sf = cf.SnapShotFactory()
	s.sm = cf.SnapShotManager()
	s.pm = cf.PeerManager()
}
