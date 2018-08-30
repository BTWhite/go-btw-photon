// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package snapshot

import (
	"sync"
	"time"
)

var (
	tickInterval = time.Second * 10
)

// SnapShotFactory is releasing new snapshots for delegates.
type SnapShotFactory struct {
	sm       *SnapShotManager
	interval time.Duration
	mu       sync.Mutex
	running  bool
}

// NewSnapShotFactory creates new factory.
// Please note that the factory needs a manager.
func (sf *SnapShotFactory) NewSnapShotFactory(sm *SnapShotManager) *SnapShotFactory {
	return &SnapShotFactory{
		sm: sm,
	}
}

// Start starting the factory process.
func (sf *SnapShotFactory) Start() {
	sf.mu.Lock()
	defer sf.mu.Unlock()
	if sf.running {
		return
	}
	sf.running = true
	go sf.cycle()
}

// Start stopping the factory process.
func (sf *SnapShotFactory) Stop() {
	sf.mu.Lock()
	sf.running = false
	sf.mu.Unlock()
}

func (sf *SnapShotFactory) cycle() {
	for sf.running {
		sf.tick()
		time.Sleep(tickInterval)
	}
}

func (sf *SnapShotFactory) tick() {
	// todo
}
