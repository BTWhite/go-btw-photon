// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package events

import "sync"

type Eventer interface {
	GetBytes() []byte
	SetBytes([]byte) int
}

type Event struct {
	mu sync.Mutex
	b  []byte
}

func (e *Event) SetBytes(bytes []byte) int {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.b = bytes

	return len(e.b)
}

func (e *Event) GetBytes() []byte {
	return e.b
}
