// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package snapshot

import (
	"github.com/BTWhite/go-btw-photon/account"
)

//
type VoteProcessor interface {
	Validate(v Vote) error
	Process(v Vote, am *account.AccountManager) error
	Commit() error
}

type BalanceProcessor interface {
	Validate(b Balance) error
}

type SnapShotProcessor interface {
	Vote() VoteProcessor
	Balance() BalanceProcessor
}

type DefaultSnapShotProcessor struct {
	v VoteProcessor
	b BalanceProcessor
}

func NewSnapShotProcessor(v VoteProcessor, b BalanceProcessor) *DefaultSnapShotProcessor {
	return &DefaultSnapShotProcessor{
		v: v,
		b: b,
	}
}

func (sp *DefaultSnapShotProcessor) Vote() VoteProcessor {
	return sp.v
}

func (sp *DefaultSnapShotProcessor) Balance() BalanceProcessor {
	return sp.b
}
