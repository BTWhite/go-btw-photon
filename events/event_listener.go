// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package events

import "sync"

var gEventListener = NewEventListener()

// EventListener sends out certain events to their subscribers.
type EventListener struct {
	mu     sync.Mutex
	events map[string][]chan *Event
}

// NewEventListener creates event listener.
func NewEventListener() *EventListener {
	return &EventListener{
		events: make(map[string][]chan *Event),
	}
}

// Subscribe creates a new chan and signs it for news, then returns.
func (el *EventListener) Subscribe(title string) chan *Event {
	c := make(chan *Event)

	el.mu.Lock()
	el.events[title] = append(el.events[title], c)
	el.mu.Unlock()

	return c
}

// Push sends a new event to its subscribers.
func (el *EventListener) Push(title string, e *Event) {
	el.mu.Lock()
	for _, event := range el.events[title] {
		event <- e
	}
	el.mu.Unlock()
}

// PushBytes converts bytes to an event and executes Push
func (el *EventListener) PushBytes(title string, bytes []byte) {
	e := new(Event)
	e.SetBytes(bytes)
	el.Push(title, e)
}

// SetEventListener sets default EventListener (without instance).
func SetEventListener(el *EventListener) {
	gEventListener = el
}

// Subscribe subscribes to default EventListener.
func Subscribe(title string) chan *Event {
	return gEventListener.Subscribe(title)
}

// Subscribe push to default EventListener.
func Push(title string, e *Event) {
	gEventListener.Push(title, e)
}

// Subscribe push bytes to default EventListener.
func PushBytes(title string, bytes []byte) {
	gEventListener.PushBytes(title, bytes)
}
