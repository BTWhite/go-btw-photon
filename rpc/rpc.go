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

// Request is the basic structure for JSON RPC requests.
type Request struct {
	Id     int32       `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params,omitempty"`
}

// Response is the basic structure for JSON RPC responses.
type Response struct {
	Id     int32       `json:"id"`
	Error  Error       `json:"error,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

// Args is necessary at the stage of receiving json, because to determine the
// necessary parameters, you first need to know what method should be called.
type Args []byte

// UnmarshalJSON allows you to put json "as is" in text form, then deserialize
// as soon as the method is known.
func (a *Args) UnmarshalJSON(b []byte) error {
	*a = b
	return nil
}

// Bytes return bytes array.
func (a *Args) Bytes() []byte {
	return *a
}

// SetConfig must be called before any methods are executed.
func SetConfig(config *config.Config) {
	cf = config
}
