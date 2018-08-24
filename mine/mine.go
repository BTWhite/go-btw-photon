// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package mine

import (
	"encoding/binary"
	"sync"

	"github.com/BTWhite/go-btw-photon/crypto/sha256"
)

func StartMine(message []byte, complexity int, threads int) *chan uint32 {
	var nonce uint32
	var proc = true
	var c = make(chan uint32)
	var mutex = &sync.Mutex{}

	for i := 0; i < threads; i++ {
		go func() {
			for proc {
				mutex.Lock()
				n := nonce
				nonce++
				mutex.Unlock()

				if mine(message, n, complexity) {
					proc = false
					c <- n
				}

			}
		}()
	}

	return &c
}

func GetHashNonce(data []byte, nonce uint32) []byte {
	nonceBuff := make([]byte, 4)
	binary.LittleEndian.PutUint32(nonceBuff, nonce)

	m := make([]byte, len(nonceBuff)+len(data))
	i := 0

	for _, t := range nonceBuff {
		m[i] = t
		i++
	}
	for _, t := range data {
		m[i] = t
		i++
	}

	return []byte(sha256.Sha256Hex(sha256.Sha256(m)))
}

func mine(message []byte, nonce uint32, complexity int) bool {
	hash := GetHashNonce(message, nonce)
	head := hash[:complexity]

	if hasValidHead(head) {
		return true
	}
	return false
}

func hasValidHead(head []byte) bool {
	for _, v := range head {
		if v != '0' {
			return false
		}
	}

	return true
}
