// Package statedb
//
// @author: xwc1125
package statedb

import (
	"math/big"

	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/models/accounts"
)

type EVMStateDB struct {
	*StateDB
}

func NewEvmStateDB(db *StateDB) *EVMStateDB {
	return &EVMStateDB{
		db,
	}
}

func (es *EVMStateDB) CreateAccount(address types.Address) {
	store := accounts.NewAccountStore(address.Hex(), accounts.ContractDomain)
	store.SetAddress(address, &accounts.AddressStore{})
	store.Normalize()

	es.StateDB.CreateAccount(store)
}

func (es *EVMStateDB) SubBalance(address types.Address, amount *big.Int) {
	account := es.StateDB.GetOwner(address)
	if account != "" {
		es.StateDB.SubBalance(account, amount)
	}
}

func (es *EVMStateDB) AddBalance(address types.Address, amount *big.Int) {
	account := es.StateDB.GetOwner(address)
	if account != "" {
		es.StateDB.AddBalance(account, amount)
	}
}

func (es *EVMStateDB) GetBalance(address types.Address) *big.Int {
	account := es.StateDB.GetOwner(address)
	if account != "" {
		return es.StateDB.GetBalance(account)
	}
	return big.NewInt(0)
}

func (es *EVMStateDB) GetNonce(address types.Address) uint64 {
	account := es.StateDB.GetOwner(address)
	if account != "" {
		return es.StateDB.GetNonce(account)
	}

	return 0
}

func (es *EVMStateDB) SetNonce(address types.Address, nonce uint64) {
	account := es.StateDB.GetOwner(address)
	if account != "" {
		es.StateDB.SetNonce(account, nonce)
	}
}

func (es *EVMStateDB) GetCodeHash(address types.Address) types.Hash {
	account := es.StateDB.GetOwner(address)
	if account != "" {
		return types.BytesToHash(es.StateDB.GetAccount(account).CodeHash())
	}

	return types.Hash{}
}

func (es *EVMStateDB) GetCode(address types.Address) []byte {
	account := es.StateDB.GetOwner(address)
	if account != "" {
		return es.StateDB.GetCode(account)
	}

	return nil
}

func (es *EVMStateDB) SetCode(address types.Address, code []byte) {
	account := es.StateDB.GetOwner(address)
	if account != "" {
		es.StateDB.SetCode(account, code)
	}
}

func (es *EVMStateDB) GetCodeSize(address types.Address) int {
	account := es.StateDB.GetOwner(address)

	if account != "" {
		return es.StateDB.GetCodeSize(account)
	}

	return 0
}

func (es *EVMStateDB) GetCommittedState(address types.Address, hash types.Hash) types.Hash {
	account := es.StateDB.GetOwner(address)

	if account != "" {
		return es.StateDB.GetCommittedState(account, hash)
	}

	return types.Hash{}
}

func (es *EVMStateDB) GetState(address types.Address, hash types.Hash) types.Hash {
	account := es.StateDB.GetOwner(address)

	if account != "" {
		return es.StateDB.GetState(account, hash)
	}

	return types.Hash{}
}

func (es *EVMStateDB) SetState(address types.Address, key, value types.Hash) {
	account := es.StateDB.GetOwner(address)
	if account != "" {
		es.StateDB.SetState(account, key, value)
	}
}

func (es *EVMStateDB) Suicide(address types.Address) bool {
	account := es.StateDB.GetOwner(address)
	if account != "" {
		return es.StateDB.Suicide(account)
	}

	return false
}

func (es *EVMStateDB) HasSuicided(address types.Address) bool {
	account := es.StateDB.GetOwner(address)

	if account != "" {
		return es.StateDB.HasSuicided(account)
	}

	return false
}

func (es *EVMStateDB) Exist(address types.Address) bool {
	return es.AddressExist(address)
}

func (es *EVMStateDB) Empty(address types.Address) bool {
	account := es.StateDB.GetOwner(address)

	if account != "" {
		return es.StateDB.Empty(account)
	}
	return false
}

func (es *EVMStateDB) ForEachStorage(address types.Address, cb func(types.Hash, types.Hash) bool) error {
	account := es.StateDB.GetOwner(address)

	if account != "" {
		return es.StateDB.ForEachStorage(account, cb)
	}

	return nil
}
