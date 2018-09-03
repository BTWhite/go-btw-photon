// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package rpc

import "github.com/BTWhite/go-btw-photon/config"

var lp = "RPC:"
var cf *config.Config

type Request struct {
	Id     int32       `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params,omitempty"`
}

type Response struct {
	Id     int32       `json:"id"`
	Error  Error       `json:"error,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

type Args []byte

// UnmarshalJSON required for deserialization and correction of standard byte processing in GO.
func (a *Args) UnmarshalJSON(b []byte) error {
	*a = b
	return nil
}

func (a *Args) Bytes() []byte {
	return *a
}

func SetConfig(config *config.Config) {
	cf = config
}
