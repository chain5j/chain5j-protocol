// Package accountdb
//
// @author: xwc1125
package accountdb

import (
	"fmt"
	"github.com/chain5j/chain5j-pkg/codec/rlp"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/pkg/database/basedb"
	"github.com/chain5j/logger"
)

type mapObject struct {
	Name string
}

func newMapObject(name string) *mapObject {
	return &mapObject{
		Name: name,
	}
}

// AccountMap 地址到账户名称的映射
type AccountMap struct {
	db basedb.Database

	trie basedb.Tree

	objects map[types.Address]*mapObject

	journal *basedb.Journal

	dbErr error
}

func NewAccountMap(root types.Hash, db basedb.Database, journal *basedb.Journal) (*AccountMap, error) {
	tr, err := db.OpenTree(root)

	if err != nil {
		return nil, err
	}

	return &AccountMap{
		db:      db,
		trie:    tr,
		objects: make(map[types.Address]*mapObject),
		journal: journal,
	}, nil
}

// setError remembers the first non-nil error it is called with.
func (a *AccountMap) setError(err error) {
	if a.dbErr == nil {
		a.dbErr = err
	}
}

// Database retrieves the low level database supporting the lower level trie ops.
func (a *AccountMap) Database() basedb.Database {
	return a.db
}

func (a *AccountMap) AddressExist(addr types.Address) bool {
	return a.getOwner(addr) != ""
}

func (a *AccountMap) GetOwner(addr types.Address) string {
	return a.getOwner(addr)
}

func (a *AccountMap) getOwner(addr types.Address) string {
	if accountObj, ok := a.objects[addr]; ok {
		return accountObj.Name
	}

	// Load the object from the database
	enc, err := a.trie.TryGet(addr[:])
	if len(enc) == 0 {
		a.setError(err)
		return ""
	}

	var data mapObject
	if err := rlp.DecodeBytes(enc, &data); err != nil {
		logger.Error("Failed to decode state object", "addr", addr, "err", err)
		return ""
	}

	return data.Name
}

func (a *AccountMap) createMap(addr types.Address, name string) {
	if a.AddressExist(addr) {
		panic("address already exists")
	}

	a.journal.Append(createMapChange{
		db:      a,
		address: addr.Hex(),
	}, a)
	a.objects[addr] = newMapObject(name)
}

func (a *AccountMap) updateMapObject(addr types.Address, obj *mapObject) {
	data, err := rlp.EncodeToBytes(obj)
	if err != nil {
		panic(fmt.Errorf("can't encode object at %s: %v", addr, err))
	}

	a.setError(a.trie.TryUpdate(addr[:], data))
}

func (a *AccountMap) deleteMapObject(addr types.Address) {
	a.setError(a.trie.TryDelete(addr[:]))
}

func (a *AccountMap) CreateMap(addr types.Address, name string) {
	a.createMap(addr, name)
}

// FinaliseObject and clears the journal as well as the refunds.
func (a *AccountMap) FinaliseObject(account string, deleteEmptyObjects bool) {
	addr := types.HexToAddress(account)
	obj, exist := a.objects[addr]
	if !exist {
		return
	}

	a.updateMapObject(addr, obj)
}

func (a *AccountMap) CommitObject(account string, deleteEmptyObjects bool) {
	addr := types.HexToAddress(account)
	obj, exist := a.objects[addr]
	if !exist {
		return
	}

	a.updateMapObject(addr, obj)
}

func (a *AccountMap) IntermediateRoot() types.Hash {
	return a.trie.Hash()
}

func (a *AccountMap) CommitTree() (types.Hash, error) {
	root, err := a.trie.Commit(func(leaf []byte, parent types.Hash) error {
		return nil
	})

	if err != nil {
		return types.Hash{}, err
	}

	err = a.db.TreeDB().Commit(root, false)
	if err != nil {
		return types.Hash{}, err
	}

	return root, err
}
