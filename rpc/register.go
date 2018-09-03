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

type Executer interface {
	Execute(id int32) *Response
}

var data = make(map[string]Executer)

func Register(name string, request Executer) {
	data[name] = request
	logger.Debug(lp, "Registered", name, "method")
}

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

	resp := method.Execute(request.Id)
	resp.Id = request.Id
	return resp
}

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
