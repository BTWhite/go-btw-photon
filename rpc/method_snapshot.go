// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package rpc

type GetSnapShotRequest struct {
	Id     string `json:"id"`
	Height int    `json:"height"`
}

type ListSnapShotRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type LastSnapShotRequest struct{}

func init() {
	Register("snapshot.get", func() Executer { return new(GetSnapShotRequest) })
	Register("snapshot.list", func() Executer { return new(ListSnapShotRequest) })
	Register("snapshot.last", func() Executer { return new(LastSnapShotRequest) })
}

func (req *GetSnapShotRequest) execute(r *Request) *Response {
	if len(req.Id) == 0 && req.Height <= 0 {
		return response(nil, err(0, "Please write correct `id` or `height`"))
	}

	ss, e := cf.SnapShotManager().Get([]byte(req.Id))
	if e != nil {
		response(nil, err(0, e.Error()))
	}
	return response(ss, nil)
}

func (req *ListSnapShotRequest) execute(r *Request) *Response {
	if req.Limit == 0 {
		req.Limit = 20
	}
	snapshots := cf.SnapShotManager().List(req.Offset, req.Limit)
	return response(snapshots, nil)
}

func (req *LastSnapShotRequest) execute(r *Request) *Response {

	snapshot, e := cf.SnapShotManager().Last()
	if e != nil {
		return response(nil, err(0, e.Error()))
	}
	return response(snapshot, nil)
}
