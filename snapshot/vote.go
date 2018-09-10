// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package snapshot

import (
	"bytes"
	"encoding/binary"

	"github.com/BTWhite/go-btw-photon/types"
)

type Vote struct {
	Voter     types.Hash `json:"voter"`
	Delegate  types.Hash `json:"delegate"`
	Timestamp int64      `json:"timestamp"`
	Signature types.Hash `json:"signature"`
}

func (v Vote) GetBytes() []byte {
	buff := new(bytes.Buffer)

	binary.Write(buff, binary.LittleEndian, v.Timestamp)
	v.Delegate.WriteToBuff(buff, 0)

	return buff.Bytes()
}
