// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package utils

// CopyOfRange takes an array of bytes and a line to be returned.
func CopyOfRange(src []byte, from, to int) []byte {
	return append([]byte(nil), src[from:to]...)
}

// FlipBytes flips an array with bytes
func FlipBytes(arr []byte) []byte {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
