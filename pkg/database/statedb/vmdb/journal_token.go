// Package vmdb
// 
// @author: xwc1125
package vmdb

import (
	"github.com/chain5j/chain5j-pkg/types"
	"math/big"
)

type
(
	tokenBalanceChange struct {
		account *types.DomainAddress
		token   *types.DomainAddress
		prev    *big.Int
	}
)

func (ch tokenBalanceChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).setTokenBalance(*ch.token, ch.prev)
}

func (ch tokenBalanceChange) dirtied() *types.DomainAddress {
	return ch.account
}
