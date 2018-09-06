// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package sync

import (
	"github.com/BTWhite/go-btw-photon/config"
)

type Syncer interface {
	// Start starts sync.
	Start()

	// Stop stops sync.
	Stop()

	// SetConfig put config in syncer.
	SetConfig(cf *config.Config)
}
