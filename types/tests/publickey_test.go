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
	"github.com/BTWhite/go-btw-photon/utils"
)

func TestHex(t *testing.T) {
	pub := types.NewPublicKeyByHex("6343b517c40fa0c733599bb9291b8482b7ca9a16297446ab7ac9de0f148eaf4c")
	want := "6343b517c40fa0c733599bb9291b8482b7ca9a16297446ab7ac9de0f148eaf4c"

	if types.NewHash(pub).ToHex() != want {
		t.Error("PublicKeyByHex incorrect, got: " + types.NewHash(pub).ToHex() + ", want: " + want)
	}
}

func TestHexAddress(t *testing.T) {
	pub := types.NewPublicKeyByHex("6343b517c40fa0c733599bb9291b8482b7ca9a16297446ab7ac9de0f148eaf4c")
	want := utils.ADDR_PREFIX + "7MxUWmF6gJFcX1VJdXUMDvd9HxjKpxDfF"

	if pub.Address() != want {
		t.Error("Address incorrect, got: " + pub.Address() + ", want: " + want)
	}
}

func TestGetPublicHex(t *testing.T) {
	const testSecret string = "ecology cart dish athlete curious potato citizen more material spray coach age"

	kp := types.NewKeyPair([]byte(testSecret))
	if types.NewHash(*kp.Public()).ToHex() != "6343b517c40fa0c733599bb9291b8482b7ca9a16297446ab7ac9de0f148eaf4c" {
		t.Error("Getting hex of public incorrect. got: " + types.NewHash(*kp.Public()).ToHex() + ", want: " + "6343b517c40fa0c733599bb9291b8482b7ca9a16297446ab7ac9de0f148eaf4c")
	}
}
