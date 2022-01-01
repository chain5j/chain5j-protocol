// Package statedb
//
// @author: xwc1125
package statedb

import "github.com/chain5j/chain5j-pkg/types"

type (
	addLogChange struct {
		db     *StateDB
		txhash types.Hash
	}

	refundChange struct {
		db   *StateDB
		prev uint64
	}

	addPreimageChange struct {
		db   *StateDB
		hash types.Hash
	}
)

func (ch addLogChange) Revert() {
	s := ch.db
	logs := s.logs[ch.txhash]
	if len(logs) == 1 {
		delete(s.logs, ch.txhash)
	} else {
		s.logs[ch.txhash] = logs[:len(logs)-1]
	}
	s.logSize--
}

func (ch addLogChange) Dirtied() string {
	return ""
}

func (ch refundChange) Revert() {
	ch.db.refund = ch.prev
}

func (ch refundChange) Dirtied() string {
	return ""
}

func (ch addPreimageChange) Revert() {
	delete(ch.db.preimages, ch.hash)
}

func (ch addPreimageChange) Dirtied() string {
	return ""
}
