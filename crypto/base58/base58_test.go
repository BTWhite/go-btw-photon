// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package base58

import (
	"testing"
)

func TestEncode(t *testing.T) {
	encoded := Encode([]byte("Hello World"))
	want := "JxF12TrwUP45BMd"

	if string(encoded) != want {
		t.Error("Base58 encoding incorrect, got: " + string(encoded) + ", want: " + want)
	}
}

func TestDecoding(t *testing.T) {

	decoded := Decode([]byte("JxF12TrwUP45BMd"))
	want := "Hello World"

	if string(decoded) != want {
		t.Error("Base58 decoding incorrect, got:", decoded, ", want:", want)
	}
}

func TestCheck(t *testing.T) {
	encoded := Check([]byte("Hello World"))
	want := "32UWxgjUJDXeRwy6c6Fxf"

	if string(encoded) != want {
		t.Error("Base58Check incorrect, got: " + string(encoded) + ", want: " + want)
	}
}
