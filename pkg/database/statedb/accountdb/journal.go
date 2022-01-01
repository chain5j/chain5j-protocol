// Package accountdb
//
// @author: xwc1125
package accountdb

import (
	"math/big"

	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/models/accounts"
)

type (
	// Changes to the account trie.
	createObjectChange struct {
		db      *AccountDB
		account string
	}

	resetObjectChange struct {
		db   *AccountDB
		prev *accountObject
	}

	// Changes to individual accounts.
	balanceChange struct {
		db      *AccountDB
		account string
		prev    *big.Int
	}
	nonceChange struct {
		db      *AccountDB
		account string
		prev    uint64
	}
	frozenChange struct {
		db      *AccountDB
		account string
		prev    bool
	}

	permissionChange struct {
		db      *AccountDB
		account string
		prev    *accounts.Permissions
	}

	addressChange struct {
		db      *AccountDB
		account string
		set     types.Address
	}

	touchChange struct {
		account string
	}

	storageChange struct {
		db *AccountDB

		account       string
		key, prevalue types.Hash
	}

	codeChange struct {
		db *AccountDB

		account            string
		prevcode, prevhash []byte
	}

	suicideChange struct {
		db          *AccountDB
		account     string
		prev        bool // whether account had already suicided
		prevbalance *big.Int
	}

	xxxChange struct {
		db       *AccountDB
		account  string
		key      string
		preValue []byte
	}

	createMapChange struct {
		db      *AccountMap
		address string
	}
)

func (ch createObjectChange) Revert() {
	delete(ch.db.accountObjects, ch.account)
}

func (ch createObjectChange) Dirtied() string {
	return ch.account
}

func (ch resetObjectChange) Revert() {
	ch.db.setAccountObject(ch.prev)
}

func (ch resetObjectChange) Dirtied() string {
	return ""
}

func (ch touchChange) Revert() {
}

func (ch touchChange) Dirtied() string {
	return ch.account
}

func (ch balanceChange) Revert() {
	ch.db.getAccountObject(ch.account).setBalance(ch.prev)
}

func (ch balanceChange) Dirtied() string {
	return ch.account
}

func (ch nonceChange) Revert() {
	ch.db.getAccountObject(ch.account).setNonce(ch.prev)
}

func (ch nonceChange) Dirtied() string {
	return ch.account
}

func (ch frozenChange) Revert() {
	ch.db.getAccountObject(ch.account).frozenAccount(ch.prev)
}

func (ch frozenChange) Dirtied() string {
	return ch.account
}

func (ch permissionChange) Revert() {
	ch.db.getAccountObject(ch.account).updatePermission(ch.prev)
}

func (ch permissionChange) Dirtied() string {
	return ch.account
}

func (ch createMapChange) Revert() {
	delete(ch.db.objects, types.HexToAddress(ch.address))
}

func (ch createMapChange) Dirtied() string {
	return ch.address
}

func (ch storageChange) Revert() {
	ch.db.getAccountObject(ch.account).setState(ch.key, ch.prevalue)
}

func (ch storageChange) Dirtied() string {
	return ch.account
}

func (ch addressChange) Revert() {
	object := ch.db.getAccountObject(ch.account)
	delete(object.data.Addresses, ch.set)
}

func (ch addressChange) Dirtied() string {
	return ch.account
}

func (ch xxxChange) Revert() {
	ch.db.getAccountObject(ch.account).setXXX(ch.key, ch.preValue)
}

func (ch xxxChange) Dirtied() string {
	return ch.account
}

func (ch codeChange) Revert() {
	ch.db.getAccountObject(ch.account).setCode(types.BytesToHash(ch.prevhash), ch.prevcode)
}

func (ch codeChange) Dirtied() string {
	return ch.account
}

func (ch suicideChange) Revert() {
	obj := ch.db.getAccountObject(ch.account)
	if obj != nil {
		obj.contract.suicided = ch.prev
		obj.setBalance(ch.prevbalance)
	}
}

func (ch suicideChange) Dirtied() string {
	return ch.account
}
