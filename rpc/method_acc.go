// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package rpc

import (
	"github.com/BTWhite/go-btw-photon/types"
)

func init() {
	Register("acc.open", new(OpenAccRequest))
	Register("acc.publicKey", new(GetAccPubKeyRequest))
}

type OpenAccRequest struct {
	PublicKey string `json:"publicKey"`
}

type GetAccPubKeyRequest struct {
	Secret string `json:"secret"`
}

func (preq *GetAccPubKeyRequest) execute(id int32) *Response {
	if len(preq.Secret) < 3 {
		return response(nil, err(0, "Please write correct `secret`"))
	}
	kp := types.NewKeyPair([]byte(preq.Secret))

	return response(types.NewHash(*kp.Public()).ToHex(), nil)
}

func (preq *OpenAccRequest) execute(id int32) *Response {
	if len(preq.PublicKey) < 3 {
		return response(nil, err(0, "Please write correct `publicKey`"))
	}

	acc := cf.AccountManager().GetByPublicKey(types.NewPublicKeyByHex(preq.PublicKey))
	return response(acc, nil)
}
