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
	"sync"

	"github.com/BTWhite/go-btw-photon/mine"
	"github.com/BTWhite/go-btw-photon/types"
)

// SnapShot is an alternative to blocks, stores only the latest balances after
// account changes and information on voting for delegates.
//
// It is the only synchronous chain that must be synchronized with all network members.
type SnapShot struct {
	Version            uint        `json:"version"`
	Id                 types.Hash  `json:"id"`
	Height             uint32      `json:"height"`
	PreviousSnapShot   types.Hash  `json:"previousSnapShot"`
	GeneratorPublicKey types.Hash  `json:"generatorPublicKey"`
	Votes              []Vote      `json:"votes"`
	Balances           []Balance   `json:"balances"`
	Timestamp          int64       `json:"timestamp"`
	Signatures         []Signature `json:"signaturess"`
	Signature          types.Hash  `json:"signature"`
	Nonce              uint32      `json:"nonce"`

	mu sync.Mutex
}

// AddVote supplements the unissued vote for further release.
func (s *SnapShot) AddVote(v Vote) {
	s.mu.Lock()
	s.Votes = append(s.Votes, v)
	s.mu.Unlock()
}

// AddVote supplements the unissued balance for further release.
func (s *SnapShot) AddBalance(b Balance) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for k, v := range s.Balances {
		if v.Address.Equals(b.Address) {
			s.Balances[k] = b
			return
		}
	}
	s.Balances = append(s.Balances, b)
}

// GetBytes is a implementation interface Byter.
func (s *SnapShot) GetBytes() []byte {
	buff := new(bytes.Buffer)

	binary.Write(buff, binary.LittleEndian, s.Height)
	binary.Write(buff, binary.LittleEndian, s.Timestamp)

	s.PreviousSnapShot.WriteToBuff(buff, 64)
	s.GeneratorPublicKey.WriteToBuff(buff, 64)

	for _, vote := range s.Votes {
		buff.Write(vote.GetBytes())
	}

	for _, balance := range s.Balances {
		buff.Write(balance.GetBytes())
	}

	return buff.Bytes()
}

func (s *SnapShot) Mine() {
	s.mu.Lock()
	data := s.GetBytes()
	cm := mine.StartMine(data, 6, 1)
	nonce := <-cm

	s.Nonce = nonce
	s.Id = s.GetId()
	s.mu.Unlock()
}

func (s *SnapShot) GetId() types.Hash {
	return mine.GetHashNonce(s.GetBytes(), s.Nonce)
}
