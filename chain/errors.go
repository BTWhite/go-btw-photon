// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package chain

import "errors"

var (
	// ErrGenesisLoaded returned if genesis chain already been loaded.
	ErrGenesisLoaded = errors.New("The genesis chain has already been loaded")

	// ErrTxInvalidPrevTx is returned if the transaction specified an invalid
	// previous transaction hash.
	ErrTxInvalidPrevTx = errors.New("Invalid previous tx")

	// ErrTxInvalidSignature is returned if signature incorrect.
	ErrTxInvalidSignature = errors.New("Invalid tx signature")

	// ErrTxInsufficientBalance is returned is sender does not have enough coins.
	ErrTxInsufficientBalance = errors.New("Insufficient balance")

	// ErrTxNotFoundInChain is returned if the transaction is not found
	// or is not in a particular chain.
	ErrTxNotFoundInChain = errors.New("Tx not found in chain")

	// ErrTxAlreadyExist is returned if tx already exist in tx list.
	ErrTxAlreadyExist = errors.New("Tx already exist")

	// ErrTxNil if a null pointer was passed when creating a chain
	ErrChainTxNil = errors.New("Genesis tx is nil")

	// ErrChainNotFound is returned is chain not found when writing a new transaction.
	ErrChainNotFound = errors.New("Chain not found")

	// ErrChainEmpty is returned is chain empty.
	ErrChainEmpty = errors.New("Chain empty")

	// ErrInsufficientData is returned if not found previous tx and sync needed.
	ErrInsufficientData = errors.New("Insufficient data")
)
