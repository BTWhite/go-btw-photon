// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package account

import "github.com/BTWhite/go-btw-photon/types"

// Account is account information storage structure.
type Account struct {
	Address   types.Hash
	PublicKey types.PublicKey
	Balance   types.Coin
}

// NewAccount creates a blank account with address.
// If you want to receive valid information about the status of your account,
// go to the AccountManager.
func NewAccount(address types.Hash) *Account {
	return &Account{
		Address: address,
	}
}
