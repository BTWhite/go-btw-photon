// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package mine

import "testing"

func TestMine(t *testing.T) {
	zeros := 6
	message := []byte("Hello World!")
	c := StartMine(message, zeros, 10)
	nonce := <-c

	h := GetHashNonce(message, nonce)

	for _, v := range h[:zeros] {
		if v != '0' {
			t.Error("Incorrect mine, got hash: ", string(h))
			break
		}
	}
}
