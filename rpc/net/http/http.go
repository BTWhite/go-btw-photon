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

	"github.com/BTWhite/go-btw-photon/json"
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

func Send(addr string, request rpc.Request) *rpc.Response {
	buff := new(bytes.Buffer)
	j, _ := json.ToJson(request)
	buff.Write(j)

	response, err := http.Post(fmt.Sprintf("http://%s/jsonrpc/", addr), "javascript/json", buff)
	if err != nil {
		logger.Err(lp, err.Error())
		return nil
	}

	b, err := ioutil.ReadAll(response.Body)

	if err != nil {
		logger.Err(lp, err.Error())
		return nil
	}

	logger.Info(string(b))
	return nil
}
