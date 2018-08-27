// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package types

import (
	"bytes"
	"encoding/hex"
)

// Hash a type for storing hashes.
type Hash []byte

// NewHash converts a byte array to Hash.
func NewHash(b []byte) Hash {
	h := make(Hash, len(b))
	copy(h, b)
	return h
}

// UnmarshalJSON required for deserialization and correction of standard byte processing in GO.
func (h *Hash) UnmarshalJSON(b []byte) error {
	if len(b) == 4 && b[0] == 'n' && b[1] == 'u' && b[2] == 'l' && b[3] == 'l' {
		*h = nil
		return nil
	}

	*h = make(Hash, len(b)-2)
	copy(*h, b[1:len(b)-1])

	return nil
}

// MarshalJSON required for serialization and correction of standard byte processing in GO.
func (h *Hash) MarshalJSON() ([]byte, error) {
	if h == nil || len(*h) == 0 {
		return []byte("null"), nil
	}
	b := make([]byte, len(*h)+2)
	b[0] = 34
	b[len(b)-1] = 34
	copy(b[1:len(b)-1], *h)
	return b, nil
}

// Equals checks the equivalence of two hashes.
func (h Hash) Equals(h2 Hash) bool {
	if len(h) != len(h2) {
		return false
	}

	for k := range h2 {
		if h[k] != h2[k] {
			return false
		}
	}
	return true
}

// String is the implementation of the Stringer interface.
func (h Hash) String() string {
	if len(h) == 0 {
		return "<nil hash>"
	}
	return string(h)
}

// ToHex gets an array of bytes or Hash and returns its text representation.
func (h Hash) ToHex() string {
	return hex.EncodeToString(h)
}

// ToBytes it simply helps to convert the Hash to bytes when needed.
func (h Hash) ToBytes() []byte {
	return h
}

// WriteToBuff writes the hash to the buffer.
// If the hash len equals 0 then defaultSize zeros will be written to the buffer.
// To avoid defaultSize, specify 0.
func (h Hash) WriteToBuff(buff *bytes.Buffer, defaultSize int) error {
	if len(h) == 0 {

		buff.Write(make([]byte, defaultSize))
	} else {
		tmp := make([]byte, defaultSize*2)
		_, err := hex.Decode(tmp, h)
		if err != nil {
			return err
		}
		sz := defaultSize
		if sz == 0 {
			sz = len(tmp)
		}
		buff.Write(tmp[:sz])
	}
	return nil
}
