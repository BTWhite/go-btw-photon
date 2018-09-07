// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package rpc

type RequestEvent struct {
	r *Request
}

func (e *RequestEvent) GetObject(obj interface{}) error {
	re := obj.(*Request)
	*re = *e.r
	return nil
}

func (e *RequestEvent) SetRequest(r *Request) {
	e.r = r
}

func (e *RequestEvent) GetBytes() []byte {
	return nil
}

func (e *RequestEvent) SetBytes([]byte) int {
	return 0
}
