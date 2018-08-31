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

type EventListener struct {
	mu     sync.Mutex
	events map[string][]chan *Event
}

func NewEventListener() *EventListener {
	return &EventListener{
		events: make(map[string][]chan *Event),
	}
}

func (el *EventListener) Subscribe(title string) chan *Event {
	c := make(chan *Event)

	el.mu.Lock()
	el.events[title] = append(el.events[title], c)
	el.mu.Unlock()

	return c
}

func (el *EventListener) Push(title string, e *Event) {
	el.mu.Lock()
	for _, event := range el.events[title] {
		event <- e
	}
	el.mu.Unlock()
}

func SetEventListener(el *EventListener) {
	gEventListener = el
}

func Subscribe(title string) chan *Event {
	return gEventListener.Subscribe(title)
}

func Push(title string, e *Event) {
	gEventListener.Push(title, e)
}
