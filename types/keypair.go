// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package types

import (
	"bytes"
	"encoding/hex"

	"github.com/BTWhite/go-btw-photon/crypto/base58"
	"github.com/BTWhite/go-btw-photon/crypto/sha256"
	"github.com/BTWhite/go-btw-photon/utils"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ripemd160"
)

// SecretKey this is a private key.
type SecretKey Hash

// PublicKey this is a public key.
type PublicKey Hash

// KeyPair this is a pair of keys: private and public.
type KeyPair struct {
	secret SecretKey
	public PublicKey
}

// Public is getter for public key.
func (k *KeyPair) Public() *PublicKey {
	return &k.public
}

// NewKeyPair creates and returns a pointer to a key pair (public, private).
func NewKeyPair(seed []byte) *KeyPair {
	hash := sha256.Sha256(seed)
	key := ed25519.NewKeyFromSeed(hash)

	publicKey := make([]byte, 32)
	copy(publicKey, key[32:])

	secretKey := make([]byte, 64)
	copy(secretKey[:32], hash)
	copy(secretKey[32:], publicKey)

	keypair := KeyPair{
		secretKey,
		publicKey,
	}
	return &keypair
}

// Secret is getter for private key.
func (k KeyPair) Secret() *SecretKey {
	return &k.secret
}

// Address creates address by public key.
func (k *PublicKey) Address() string {
	hash := sha256.Sha256(*k)
	ripemd := ripemd160.New()
	ripemd.Write(hash)
	sum := ripemd.Sum([]byte{})

	buffer := bytes.NewBuffer(nil)
	buffer.WriteString(utils.ADDR_PREFIX)
	buffer.Write(base58.Check(sum))

	return buffer.String()
}

// NewPublicKeyByHex creates a public key for the string representation.
func NewPublicKeyByHex(s string) PublicKey {
	b, _ := hex.DecodeString(s)
	return PublicKey(b)
}

func HasAddr(addr []byte) bool {
	return len(addr) > 2 && addr[0] == 'B' && addr[1] == '0'
}
