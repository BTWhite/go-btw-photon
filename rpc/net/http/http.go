// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/BTWhite/go-btw-photon/json"
	"github.com/BTWhite/go-btw-photon/peer"
	"github.com/BTWhite/go-btw-photon/rpc"

	"github.com/BTWhite/go-btw-photon/logger"
)

var lp = "HTTP RPC:"

var parseError, _ = json.ToJson(rpc.Response{
	Id:     0,
	Result: nil,
	Error:  rpc.ErrParseError,
})

var internalError, _ = json.ToJson(rpc.Response{
	Id:     0,
	Result: nil,
	Error:  rpc.ErrInternalError,
})

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Err(lp, "Read Body:", err.Error())
		w.Write(internalError)
		return
	}

	args := &rpc.Args{}
	req := &rpc.Request{
		Params: args,
	}
	err = json.FromJson(data, req)
	if err != nil {
		logger.Err(lp, "Json Parse:", err.Error())
		w.Write(parseError)
		return
	}
	logger.Debug(lp, "method", req.Method, req.Id, "from", r.RemoteAddr)
	resp := rpc.ExecuteRequest(req, args)
	j, err := json.ToJson(resp)
	if err != nil {
		logger.Err(lp, "Internal Error:", err.Error())
		w.Write(internalError)
		return
	}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.Write(j)
}

// Start starts http server.
func Start(port int) error {
	http.HandleFunc("/jsonrpc/", handler)
	logger.Debug(lp, "Listen on", port, "port")
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		logger.Err(lp, err.Error())
		return err

	}

	return nil
}

// Send sends a request to the specified address.
func Send(addr string, request rpc.Request, respArgs interface{}) (*rpc.Response, error) {
	buff := new(bytes.Buffer)
	j, _ := json.ToJson(request)
	buff.Write(j)
	resp := &rpc.Response{
		Result: respArgs,
	}

	r, e := http.Post(addr, "javascript/json", buff)
	if e != nil {
		logger.Err(lp, e.Error())
		return resp, e
	}

	b, e := ioutil.ReadAll(r.Body)

	if e != nil {
		logger.Err(lp, e.Error())
		return resp, e
	}

	e = json.FromJson(b, resp)

	if e != nil {
		return resp, e
	}

	return resp, nil
}

// BroadCast sends requests to the `count` random peers.
func BroadCast(pm *peer.PeerManager, request rpc.Request, respArgs interface{}, count int) []*rpc.Response {
	if count <= 0 {
		count = 20
	}
	peers := pm.Random(count)
	results := make([]*rpc.Response, len(peers))
	//	for _, _ := range results {

	//	}

	wg := new(sync.WaitGroup)
	for i, peer := range peers {
		wg.Add(1)
		go func(i int, respArgs interface{}) {
			results[i], _ = Send(peer.HttpAddr(), request, &respArgs)

			wg.Done()
		}(i, respArgs)
	}

	wg.Wait()
	return results
}
