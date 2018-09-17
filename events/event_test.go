// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package events

import (
	"fmt"
	"testing"
	"time"
)

func TestSub(t *testing.T) {
	c := Subscribe("test")
	succ := false
	go func(c chan Eventer, suss *bool) {
		e := <-c
		bytes := e.GetBytes()
		want := []byte("Hello World")

		if len(bytes) != len(want) {
			t.Fatal("Chain returned len:", len(bytes), "want:", len(want))
		}

		for i, b := range bytes {
			if b != want[i] {
				t.Fatal("Invalid event", "got:", string(bytes), "want:", string(want))
			}
		}

		succ = true
	}(c, &succ)

	e := new(Event)
	e.SetBytes([]byte("Hello World"))
	Push("test", e)

	time.Sleep(time.Millisecond * 10)

	if !succ {
		t.Fatal("Event undelivered")
	}
}

func TestMultiPush(t *testing.T) {
	c := Subscribe("multipush")
	run := true

	go func(c chan Eventer) {
		i := 0
		for run {
			<-c
			i++
			if i == 5 {
				run = false
			}
		}

	}(c)

	for i := 0; i < 5; i++ {
		go func() {
			e := new(Event)
			e.SetBytes([]byte(fmt.Sprint("Hello from gorutine", i+1)))
			Push("multipush", e)
		}()
	}

	for j := 0; ; j++ {
		if run && j >= 1000 {
			t.Fatal("Events undelivered")
		} else if !run {
			break
		}
		time.Sleep(time.Millisecond)
	}

}
