// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package types

import (
	"math/big"
)

const (
	unitCoin uint64 = 100000000
)

// Coin is the type for safe interaction with the values of the number of coins.
type Coin uint64

// NewCoin creates a type of coins, as an argument expects to get
// a cipher representation of the sum.
// For example 100 000 000 is 1 coin, 10 000 000 is 0.1 coin.
func NewCoin(u uint64) Coin {
	return Coin(u)
}

// String returns a string with the number of coins understandable to the user.
func (c Coin) String() string {
	cf := big.NewFloat(float64(c))
	cu := big.NewFloat(float64(unitCoin))

	return big.NewFloat(0).Quo(cf, cu).String()
}

// Uint64 return the root view of the coin in uint64.
func (c Coin) Uint64() uint64 {
	return uint64(c)
}
