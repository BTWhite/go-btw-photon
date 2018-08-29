// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.
//
// This package is required for development and will be deleted after the expiration of the time.
//
// Genesis account:
// AoqtQfSZyfCex8q4fGwMAFEQbcDt4mZfP | HelloSubject

package main

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/BTWhite/go-btw-photon/account"
	"github.com/BTWhite/go-btw-photon/chain"
	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/json"
	"github.com/BTWhite/go-btw-photon/logger"
	"github.com/BTWhite/go-btw-photon/types"
)

func main() {
	logger.Init("debug")
	logger.Debug("BitWhite Node starting...")
	db := leveldb.Open("data/")
	txTbl := db.CreateTable([]byte("tx"))

	cb := chain.NewChainBook(db, chain.NewProcessor(db))
	err := cb.LoadGenesis("genesis.json")
	if err != nil && err != chain.ErrGenesisLoaded {
		panic(err.Error())
	}
	cb.GetChain([]byte("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"))
	am := account.NewAccountManager(db)

	for true {
		var tmp string
		fmt.Print("> ")
		fmt.Scanf("%s", &tmp)

		switch tmp {
		case "tx":
			createTx(cb, am)
		case "gettx":
			getTx(txTbl)
		case "balance":
			getBalance(am)
		case "chain":
			getChain(cb)
		}
	}
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func getChain(cb *chain.ChainBook) {
	var tmp string
	fmt.Print("Chain: ")
	fmt.Scanf("%s", &tmp)

	ch, err := cb.GetChain([]byte(tmp))
	if err != nil {
		logger.Err(err.Error())
		return
	}

	logger.Info(string(json.ToJson(ch)))
}

func getTx(txTbl *leveldb.Tbl) {
	var tmp string
	fmt.Print("Tx: ")
	fmt.Scanf("%s", &tmp)

	tx, err := types.GetTx(types.NewHash([]byte(tmp)), txTbl)

	if err != nil {
		logger.Err(err.Error())
		return
	}

	logger.Info(string(json.ToJson(tx)))
}

func createTx(cb *chain.ChainBook, am *account.AccountManager) {
	var tmp string
	var tmpI uint64
	fmt.Print("Secret: ")
	fmt.Scanf("%s", &tmp)

	kp := types.NewKeyPair([]byte(tmp))
	fmt.Println("Address:", kp.Public().Address(), "| Balance:",
		am.Get([]byte(kp.Public().Address())).Balance)

	fmt.Print("To: ")
	fmt.Scanf("%s", &tmp)

	fmt.Println("To address:", tmp, "| Balance:", am.Get([]byte(tmp)).Balance)

	fmt.Print("Amount: ")
	fmt.Scanf("%d", &tmpI)

	tx, err := cb.CreateTx(kp, types.NewCoin(tmpI), types.NewCoin(10000000),
		types.NewHash([]byte(tmp)), []byte("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"))
	if err != nil {
		logger.Err(err.Error())
	}

	err = cb.AddTx(tx)
	if err != nil {
		logger.Err(err.Error())
	}
}

func getBalance(am *account.AccountManager) {
	var tmp string
	fmt.Print("Address: ")
	fmt.Scanf("%s", &tmp)

	acc := am.Get([]byte(tmp))

	logger.Info(acc.Address, acc.Balance)
}
