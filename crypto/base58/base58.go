// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package base58

const ALPHABET string = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var INDEXES []int = make([]int, 128)

// Init is necessary for correct operation of decoding.
// You must call this function once before using the package.
func Init() {
	var i int = 0
	for i < len(INDEXES) {
		INDEXES[i] = -1
		i++
	}

	i = 0
	for i < len(ALPHABET) {
		INDEXES[int(ALPHABET[i])] = i
		i++
	}
}

// Encode encodes bytes in base58.
func Encode(input []byte) []byte {
	if len(input) == 0 {
		return []byte{}
	}

	var zeroCount int = 0
	for zeroCount < len(input) && input[zeroCount] == 0 {
		zeroCount++
	}

	var temp []byte = make([]byte, len(input)*2)
	var j int = len(temp)

	var startAt int = zeroCount

	for startAt < len(input) {
		var mod byte = divmod58(input, startAt)
		if input[startAt] == 0 {
			startAt++
		}
		j--
		temp[j] = byte(ALPHABET[mod])
	}

	for j < len(temp) && temp[j] == ALPHABET[0] {
		j++
	}

	for true {
		zeroCount--
		if zeroCount < 0 {
			break
		} else {
			j--
			temp[j] = byte(ALPHABET[0])
		}
	}

	var output []byte = temp[j:len(temp)]
	return output
}

// Decode decodes bytes from base58.
func Decode(input []byte) []byte {
	if len(input) == 0 {
		return []byte{}
	}

	var i int = 0
	for i < len(input) {
		var c uint8 = input[i]
		var digit58 int = -1
		if int(c) >= 0 && int(c) < 128 {
			digit58 = INDEXES[c]
		}
		if digit58 < 0 {
			return []byte{}
		}

		input[i] = byte(digit58)
		i++
	}

	var zeroCount int = 0
	for zeroCount < len(input) && input[zeroCount] == 0 {
		zeroCount++
	}

	var temp []byte = make([]byte, len(input))
	var j = len(temp)
	var startAt = zeroCount

	for startAt < len(input) {
		var mod byte = divmod256(input, startAt)
		if input[startAt] == 0 {
			startAt++
		}
		j--
		temp[j] = mod
	}

	for j < len(temp) && temp[j] == 0 {
		j++
	}

	return temp[j-zeroCount : len(temp)]
}

func divmod58(number []byte, startAt int) byte {
	var remainder int = 0
	var i int = startAt
	for i < len(number) {
		var digit256 int = int(number[i] & 0xFF)
		var temp int = remainder*256 + digit256
		number[i] = byte(temp / 58)
		remainder = temp % 58
		i++
	}

	return byte(remainder)
}

func divmod256(number58 []byte, startAt int) byte {
	var remainder int = 0
	var i = startAt
	for i < len(number58) {
		var digit58 int = int(number58[i] & 0xFF)
		var temp int = remainder*58 + digit58
		number58[i] = byte(temp / 256)
		remainder = temp % 256
		i++
	}
	return byte(remainder)
}
