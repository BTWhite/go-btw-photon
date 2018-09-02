// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package config

import "testing"

var inputs = []string{
	"0.0",
	"0.0.0",
	"1.4",
	"1.2.0",
	"1.2.3",
}

var outputs = []Version{
	{0, 0, 0},
	{0, 0, 0},
	{1, 4, 0},
	{1, 2, 0},
	{1, 2, 3},
}

func TestToString(t *testing.T) {
	for k, v := range outputs {
		if v.String() != inputs[k] {
			if len(inputs[k]) >= 4 && string(inputs[k][4]) == "0" && outputs[k][2] == 0 {
				continue
			}

			t.Fatal("String method incorrect,", "got:", v.String(), "want:", inputs[k])
		}
	}
}

func TestFromString(t *testing.T) {
	for k, v := range inputs {
		v := NewVersionByString(v)
		if !v.Equals(outputs[k]) {
			t.Fatal("FromString method incorrect", "got:", v, "want:", outputs[k])
		}
	}
}
