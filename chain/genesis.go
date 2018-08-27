// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package chain

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/BTWhite/go-btw-photon/json"
	"github.com/BTWhite/go-btw-photon/types"
)

type genesis struct {
	Chain        *Chain      `json:"chain"`
	Transactions []*types.Tx `json:"transactions"`
}

func LoadGenesis(filename string) (*Chain, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	g := &genesis{}

	json.FromJson(data, g)

	for _, tx := range g.Transactions {
		tx.Timestamp = time.Now().Unix()
		tx.Mine()

		err, hash := g.Chain.AddTx(tx)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(hash)
		}

	}
	_, _, tx := g.Chain.GetTx([]byte("0000accc2cecd6a183ae426739267b06c0e0db509d8fe1a38b6628807eb57f4f"))
	//	fmt.Println(err.Error())
	fmt.Println(string(json.ToJson(tx)))
	return nil, nil
}
