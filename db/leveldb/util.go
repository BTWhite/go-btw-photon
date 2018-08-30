// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3

package leveldb

import (
	"bytes"
	"encoding/gob"
)

func decode(data []byte, obj interface{}) error {
	var b bytes.Buffer
	b.Write(data)
	d := gob.NewDecoder(&b)

	return d.Decode(obj)
}

func encode(obj interface{}) ([]byte, error) {
	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	if err := e.Encode(obj); err != nil {
		return make([]byte, 0), err
	}

	return b.Bytes(), nil
}
