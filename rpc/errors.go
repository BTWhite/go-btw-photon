// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package rpc

type Error interface {
	Code() int32
	Message() string
}

type defaultError struct {
	C int32  `json:"code"`
	M string `json:"message"`
}

var (
	// ErrParseError returned if invalid JSON was received by the server.
	// An error occurred on the server while parsing the JSON text.
	ErrParseError = err(-32700, "Parse Error")

	/// ErrInvalidRequest returned if JSON sent is not a valid Request object.
	ErrInvalidRequest = err(-32600, "Invalid Request")

	// ErrMethodNotFound returned if method does not exist / is not available.
	ErrMethodNotFound = err(-32601, "Method not found")

	// ErrInvalidParams returned if received invalid method parameter(s).
	ErrInvalidParams = err(-32602, "Invalid method parameter(s)")

	// ErrInternalError returned if occurred internal JSON-RPC error.
	ErrInternalError = err(-32603, "Internal error")

	// New errors code:
	// -32000 to -32099
)

func err(code int32, message string) Error {
	return &defaultError{code, message}
}

func (e *defaultError) Code() int32 {
	return e.C
}

func (e *defaultError) Message() string {
	return e.M
}
