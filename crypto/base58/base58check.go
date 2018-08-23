// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package base58

import (
	"bytes"

	"github.com/BTWhite/go-btw-photon/crypto/sha256"
)

// Check produces a hash in sha256 and maps it into a special base58 format.
// Used in address generation.
func Check(input []byte) []byte {
	checksum := sha256.Sha256X2(input)

	/*b := make([]byte, len(input)+4)
	buffer := bytes.NewBuffer(b)*/

	buffer := bytes.NewBuffer(make([]byte, 0))

	buffer.Write(input)
	buffer.Write(checksum[:4])

	return Encode(buffer.Bytes())
}
