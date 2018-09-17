// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package mine

import (
	"sync"

	"github.com/BTWhite/go-btw-photon/interfaces"
)

type Mine struct {
	data       []byte
	nonce      *uint32
	timestamp  *int64
	hash       *[]byte
	complexity int
	mu         sync.Mutex
	wg         sync.WaitGroup
	running    bool
}

type Miner interface {
	interfaces.Byter
	GetTimestamp() *int64
	GetNonce() *uint32
	GetId() *[]byte
	GetComplexity() int
}

func NewMine(obj Miner) *Mine {
	return &Mine{
		data:       obj.GetBytes(),
		nonce:      obj.GetNonce(),
		timestamp:  obj.GetTimestamp(),
		hash:       obj.GetId(),
		complexity: obj.GetComplexity(),
	}
}

func (m *Mine) Start(threads int) *sync.WaitGroup {
	m.running = true
	m.wg.Wait()
	m.wg.Add(1)
	for i := 0; i < threads; i++ {
		go func() {
			for m.Running() {
				m.mu.Lock()
				n := *m.nonce
				*m.nonce++
				m.mu.Unlock()

				if mine(m.data, n, m.complexity) {
					*m.hash = GetHashNonce(m.data, n)
					m.Stop()
					m.wg.Done()
				}

			}
		}()
	}
	return &m.wg
}

func (m *Mine) Running() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.running
}

func (m *Mine) Stop() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.running = false
}
