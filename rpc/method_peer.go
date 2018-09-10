// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package rpc

type GetPeersRequest struct {
	Limit int `json:"limit"`
}

func init() {
	Register("peers.get", func() Executer { return new(GetPeersRequest) })
}

func (req *GetPeersRequest) execute(r *Request) *Response {
	if req.Limit < 1 {
		req.Limit = 20
	}
	peers := cf.PeerManager().Random(req.Limit)

	if r.Peer != nil {
		if !cf.PeerManager().Exist(*r.Peer) {
			cf.PeerManager().Save(*r.Peer)
		}
	}
	return response(peers, nil)
}
