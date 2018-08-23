// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package sha256

import (
	"testing"
)

func TestSha256(t *testing.T) {
	hex := Sha256Hex([]byte("Hello World"))
	want := "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e"
	if hex != want {
		t.Error("Sha256 incorrect, got: " + hex + ", want: " + want)
	}
}
