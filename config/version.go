// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package config

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/BTWhite/go-btw-photon/logger"
)

// Version this is the type for storing the protocol version.
// Range versions: 0.0.0 - 255.255.255.
// First byte for major changes.
// Second byte for minor changes.
// Third byte is optional value for micro changes.
type Version [3]byte

// NewVersionByString converts version string to Version type.
func NewVersionByString(v string) Version {
	strs := strings.SplitN(v, ".", 3)
	ver := Version{}
	for i := 0; i < len(strs); i++ {
		v, err := strconv.ParseInt(strs[i], 10, 64)
		if err != nil {
			logger.Err(err.Error())
			continue
		}

		ver[i] = byte(v)
	}

	return ver
}

// String converts version to string representation.
func (v Version) String() string {
	if v[2] == 0 {
		return fmt.Sprintf("%d.%d", v[0], v[1])
	}

	return fmt.Sprintf("%d.%d.%d", v[0], v[1], v[2])
}

// Equals check if the two versions match
func (v Version) Equals(v2 Version) bool {
	if v[0] != v2[0] ||
		v[1] != v2[1] ||
		v[2] != v2[2] {
		return false
	}
	return true
}

// GetBytes converts version to the []byte array
func (v Version) GetBytes() []byte {
	buff := make([]byte, 3)
	for k, b := range v {
		buff[k] = b
	}

	return buff
}

// GetBytes3 converts version to the []byte array
func (v Version) GetBytes3() [3]byte {
	return v
}
