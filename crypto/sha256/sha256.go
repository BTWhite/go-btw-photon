// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package sha256

import (
	"crypto/sha256"
	"encoding/hex"
)

// Sha256 hashes incoming data and returns hash bytes.
func Sha256(input []byte) []byte {
	hash := sha256.New()
	hash.Write(input)
	md := hash.Sum(nil)
	return md
}

// Sha256Hex works on Sha256, but returns a string representation of the hash.
func Sha256Hex(input []byte) string {
	return hex.EncodeToString(Sha256(input))
}

// EncodeX2 hash the value twice -> sha256(sha256(x)).
func Sha256X2(input []byte) []byte {
	return Sha256(Sha256(input))
}
