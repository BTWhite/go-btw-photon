// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3

package utils

import (
	"encoding/json"
)

// ToJson attempts to convert transferred structure to the json
func ToJson(o interface{}) []byte {
	b, err := json.Marshal(o)

	if err != nil {
		return nil
	}
	return b
}

// FromJson attempts to convert json to the transferred structure
func FromJson(data []byte, o interface{}) error {
	err := json.Unmarshal(data, o)
	if err != nil {
		return err
	}
	return nil
}
