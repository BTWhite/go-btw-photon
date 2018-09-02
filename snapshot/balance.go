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

	"github.com/BTWhite/go-btw-photon/account"
	"github.com/BTWhite/go-btw-photon/types"
)

type Balance struct {
	Address types.Hash `json:"addr"`
	Balance types.Coin `json:"blnc"`
}

func BalanceByAccount(acc *account.Account) Balance {
	return Balance{
		Address: acc.Address,
		Balance: acc.Balance,
	}
}

func (b Balance) GetBytes() []byte {
	buff := new(bytes.Buffer)

	binary.Write(buff, binary.LittleEndian, b.Balance)
	if b.Address != nil {
		b.Address.WriteToBuff(buff, 0)
	}

	return buff.Bytes()
}
