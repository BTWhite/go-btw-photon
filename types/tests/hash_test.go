// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package tests

import (
	"testing"

	"github.com/BTWhite/go-btw-photon/types"
)

func TestBytesToHash(t *testing.T) {
	want := []byte("Hello World!")
	var hash types.Hash = types.NewHash(want)

	if !equals(want, hash) {
		t.Error("Error BytesToHash, got: \n", hash, "\nwant:\n", want)
	}
}

func TestToHex(t *testing.T) {
	h := types.NewHash([]byte("Hello World!"))
	want := "48656c6c6f20576f726c6421"
	got := h.ToHex()
	if want != got {
		t.Error("Invalid HashToHex, got:", got, ", want:", want)
	}
}

func equals(b []byte, hash types.Hash) bool {
	if len(b) != len(hash) {
		return false
	}
	for i := 0; i < len(b); i++ {
		if hash[i] != b[i] {
			return false
		}
	}
	return true
}
