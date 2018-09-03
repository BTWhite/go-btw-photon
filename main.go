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
	"flag"

	"github.com/BTWhite/go-btw-photon/rpc"

	"github.com/BTWhite/go-btw-photon/chain"
	"github.com/BTWhite/go-btw-photon/config"
	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/logger"
	"github.com/BTWhite/go-btw-photon/rpc/net/http"
)

func main() {
	logger.Init("debug")

	port := flag.Int("http-port", 8886, "http json rpc port")
	magic := flag.String("magic", "b1m52ot80x", "magic value")
	delegate := flag.String("delegate", "none", "delegate secret")
	genesis := flag.String("genesis", "genesis.json", "genesis file")
	flag.Parse()

	db := leveldb.Open("data/")
	cf := config.NewConfig(db, []byte(*magic), [3]byte{0, 1, 0})

	if *delegate != "none" {
		cf.SnapShotFactory().Start()
	}

	chain.LoadGenesis(*genesis, cf.ChainHelper())
	rpc.SetConfig(cf)

	http.Start(*port)
}
