// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package events

import "sync"

// Eventer is interface that defines an object that can be perceived as an event.
type Eventer interface {
	// GetBytes gets bytes from the event.
	GetBytes() []byte
	// SetBytes sets bytes to the event.
	SetBytes([]byte) int
}

// Event is default event implementation.
type Event struct {
	mu sync.Mutex
	b  []byte
}

// SetBytes sets bytes to the event.
func (e *Event) SetBytes(bytes []byte) int {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.b = bytes

	return len(e.b)
}

// GetBytes gets bytes from the event.
func (e *Event) GetBytes() []byte {
	return e.b
}
