// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package snapshot

import "github.com/BTWhite/go-btw-photon/types"

// Signature is type for compare signature and publickey for multisigns.
type Signature struct {
	PublicKey types.Hash
	Signature types.Hash
}
