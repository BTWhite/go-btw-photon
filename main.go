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
	"io/ioutil"

	"github.com/BTWhite/go-btw-photon/config"
	"github.com/BTWhite/go-btw-photon/db/leveldb"
	"github.com/BTWhite/go-btw-photon/json"
	"github.com/BTWhite/go-btw-photon/logger"
	"github.com/BTWhite/go-btw-photon/node"
)

func main() {

	c := initParams()

	port := flag.Int("http-port", c.Port, "http json rpc port")
	magic := flag.String("magic", c.Magic, "magic value")
	delegate := flag.String("delegate", c.Delegate, "delegate secret")
	genesis := flag.String("genesis", c.Genesis, "genesis file")
	logLevel := flag.String("log", c.LogLevel, "log level (debug|error|info)")

	flag.Parse()

	c.Port = *port
	c.Magic = *magic
	c.Delegate = *delegate
	c.Genesis = *genesis
	c.LogLevel = *logLevel

	db := leveldb.Open("data/")
	cf := config.NewConfig(db, []byte(c.Magic), [3]byte{0, 1, 0})

	node.StartNode(cf, c)
}

func initParams() node.Params {
	c := node.Params{}

	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = json.FromJson(b, &c)
	if err != nil {
		logger.Fatal(err.Error())
	}

	return c
}
