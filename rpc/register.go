// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package rpc

import (
	"github.com/BTWhite/go-btw-photon/json"
	"github.com/BTWhite/go-btw-photon/logger"
)

// Executer at its core the request object that consists of the necessary fields.
type Executer interface {
	execute(id int32) *Response
}

var data = make(map[string]Executer)

// Register a registers a new method.
func Register(name string, request Executer) {
	data[name] = request
	logger.Debug(lp, "Registered", name, "method")
}

// ExecuteRequest executes RPC request and returns a response if the method is
// not valid anyway, the response will return with the corresponding error.
func ExecuteRequest(request *Request, args *Args) *Response {
	method, exist := data[request.Method]

	if !exist {
		return request.Response(nil, ErrMethodNotFound)
	}

	if len(args.Bytes()) > 0 {
		err := json.FromJson(args.Bytes(), method)
		if err != nil {
			logger.Err(lp, "Execute Request:", err.Error(), "Json:", string(args.Bytes()))
			return request.Response(nil, ErrParseError)
		}
	}

	resp := method.execute(request.Id)
	resp.Id = request.Id
	return resp
}

// Response returns the answer and makes it easy to create answers.
func (req *Request) Response(result interface{}, err Error) *Response {
	return &Response{
		Id:     req.Id,
		Result: result,
		Error:  err,
	}
}

func response(result interface{}, err Error) *Response {
	return &Response{
		Result: result,
		Error:  err,
	}
}
