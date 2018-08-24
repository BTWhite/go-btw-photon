package mine

import "testing"

func TestMine(t *testing.T) {
	zeros := 2
	message := []byte("Hello World!")
	c := StartMine(message, zeros, 10)
	nonce := <-c

	h := GetHashNonce(message, nonce)

	for _, v := range h[:zeros] {
		if v != '0' {
			t.Error("Incorrect mine, got hash: ", string(h))
			break
		}
	}
}
