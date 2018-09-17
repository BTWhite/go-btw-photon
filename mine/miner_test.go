// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package mine

import (
	"bytes"
	"encoding/binary"
	"testing"
	"time"
)

type Test struct {
	Id        []byte
	Nonce     uint32
	Timestamp int64
}

func (m *Test) GetTimestamp() *int64 {
	return &m.Timestamp
}

func (m *Test) GetNonce() *uint32 {
	return &m.Nonce
}

func (m *Test) GetId() *[]byte {
	return &m.Id
}

func (m *Test) GetBytes() []byte {
	buff := new(bytes.Buffer)
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, uint64(m.Timestamp))
	buff.Write(data)

	return buff.Bytes()
}

func (m *Test) GetComplexity() int {
	return 2
}

func TestMiner(t *testing.T) {
	tm := &Test{}
	tm.Timestamp = time.Now().Unix()

	m := NewMine(tm)
	wg := m.Start(1)
	t.Error("All ok")
	wg.Wait()

	t.Error("Err", string(tm.Id), len(tm.Id), tm.Nonce)

	//	for _, v := range tm.Id[:tm.GetComplexity()] {
	//		if v != '0' {
	//			t.Error("Incorrect mine, got hash: ", string(tm.Id))
	//			break
	//		}
	//	}
}
