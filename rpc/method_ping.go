// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package rpc

import (
	"fmt"
)

// PingRequest is a small way to check if the node is responding.
type PingRequest string

func init() {
	Register("ping", func() Executer { return new(PingRequest) })

}

func (preq *PingRequest) execute(r *Request) *Response {
	if *preq != "beep" {
		return response(nil, err(0, fmt.Sprintf("What is %s? :/", *preq)))
	}
	return response("boop", nil)
}
