// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3

package sign

import (
	"encoding/hex"

	"github.com/BTWhite/go-btw-photon/crypto/sha256"
	"github.com/BTWhite/go-btw-photon/interfaces"
	"github.com/BTWhite/go-btw-photon/types"
	"golang.org/x/crypto/ed25519"
)

// Sign creates a signature for any element that Byteble implements.
// Requires that you also specify generator and signature, where the public key
// of the signer and his signature will be placed.
// dataOffset will also help you delete the last n bytes of the byte array.
func Sign(object interfaces.Byter, pair *types.KeyPair, generator *types.Hash,
	signature *types.Hash, dataOffset int) {

	generatorHex := make([]byte, 64)
	pub := *pair.Public()
	hex.Encode(generatorHex, pub)
	*generator = generatorHex

	bytes := object.GetBytes()
	data := bytes[:len(bytes)-dataOffset]
	hashBytes := sha256.Sha256(data)

	secret := *pair.Secret()
	privateBytes := types.NewHash(secret).ToBytes()
	sign := ed25519.Sign(privateBytes, hashBytes)

	signHex := make([]byte, 128)
	hex.Encode(signHex, sign)
	*signature = signHex
}

// Verify verifies the validity of the generator signature by Byter object.
// Unlike Sign does not require pointers.
// For the validity of events, do not forget to specify the same dataOffset
// when signing and verifying.
func Verify(object interfaces.Byter, generator types.Hash,
	signature types.Hash, dataOffset int) bool {

	bytes := object.GetBytes()
	data := bytes[:len(bytes)-dataOffset]
	hashBytes := sha256.Sha256(data)

	sig := make([]byte, 64)
	gen := make([]byte, 32)
	hex.Decode(sig, signature.ToBytes())
	hex.Decode(gen, generator.ToBytes())

	return ed25519.Verify(gen, hashBytes, sig)
}
