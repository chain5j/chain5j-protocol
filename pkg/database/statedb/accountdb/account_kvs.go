// Package accountdb
//
// @author: xwc1125
package accountdb

import (
	"errors"
	"fmt"
	"github.com/chain5j/chain5j-pkg/codec/rlp"
	linkedHashMap "github.com/chain5j/chain5j-pkg/collection/maps/linked_hashmap"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/pkg/database/basedb"
)

// AccountKVS 账户kvs
type AccountKVS struct {
	db   basedb.Database
	trie basedb.Tree

	objects map[string]*linkedHashMap.LinkedHashMap // 额外存储的相关数据（namespace==>key==>value）
	journal *basedb.Journal
	dbErr   error
}

func NewAccountKVS(root types.Hash, db basedb.Database, journal *basedb.Journal) (*AccountKVS, error) {
	tr, err := db.OpenTree(root)

	if err != nil {
		return nil, err
	}

	return &AccountKVS{
		db:      db,
		trie:    tr,
		objects: make(map[string]*linkedHashMap.LinkedHashMap),
		journal: journal,
	}, nil
}

func (a *AccountKVS) setError(err error) {
	if a.dbErr == nil {
		a.dbErr = err
	}
}

func (a *AccountKVS) Database() basedb.Database {
	return a.db
}

func (a *AccountKVS) NamespaceExist(account string, namespace string) bool {
	obj, ok := a.objects[account]
	if !ok {
		return false
	}
	_, ok = obj.Get(namespace)
	return ok
}

func (a *AccountKVS) KeyExist(account string, namespace string, key string) bool {
	obj, ok := a.objects[account]
	if !ok {
		return false
	}
	kvnp, _ := obj.Get(namespace)
	if kvnp == nil {
		return false
	}
	hashMap := kvnp.(*linkedHashMap.LinkedHashMap)
	keyVal, _ := hashMap.Get(key)
	if keyVal == nil {
		return false
	}
	return true
}

func (a *AccountKVS) AddKV(account string, namespace string, key string, val interface{}) {
	obj, ok := a.objects[account]
	if !ok {
		obj = linkedHashMap.NewLinkedHashMap()
		a.objects[account] = obj
	}
	kvnp, _ := obj.Get(namespace)
	if kvnp == nil {
		kvnp = linkedHashMap.NewLinkedHashMap()
		obj.Add(namespace, kvnp)
	}
	hashMap := kvnp.(*linkedHashMap.LinkedHashMap)
	hashMap.Add(key, val)

	//s.dirtyKvs = true
}

func (a *AccountKVS) setKVS(account string, kvs *linkedHashMap.LinkedHashMap) {
	a.objects[account] = kvs
	//s.kvsHash = kvsHash[:]
	//s.dirtyKvs = true
}

func (a *AccountKVS) setKVNP(account string, namespace string, _kvnp *linkedHashMap.LinkedHashMap) {
	obj, ok := a.objects[account]
	if !ok {
		obj = linkedHashMap.NewLinkedHashMap()
		a.objects[account] = obj
	}
	obj.Add(namespace, _kvnp)
	//a.kvsHash = kvsHash[:]
	//a.dirtyKvs = true
}

func (a *AccountKVS) KVS(account string) *linkedHashMap.LinkedHashMap {
	return a.objects[account]
}

func (a *AccountKVS) KVNP(account string, namespace string) (*linkedHashMap.LinkedHashMap, error) {
	obj, ok := a.objects[account]
	if !ok {
		obj = linkedHashMap.NewLinkedHashMap()
		a.objects[account] = obj
		return nil, errors.New("object is empty")
	}
	kvnp, _ := obj.Get(namespace)
	if kvnp == nil {
		return nil, errors.New("namespace is nil")
	}
	hashMap := kvnp.(*linkedHashMap.LinkedHashMap)
	return hashMap, nil
}

func (a *AccountKVS) KV(account string, namespace string, key string) (interface{}, error) {
	obj, ok := a.objects[account]
	if !ok {
		obj = linkedHashMap.NewLinkedHashMap()
		a.objects[account] = obj
		return nil, errors.New("object is empty")
	}
	kvnp, _ := obj.Get(namespace)
	if kvnp == nil {
		return nil, errors.New("namespace is nil")
	}
	hashMap := kvnp.(*linkedHashMap.LinkedHashMap)
	val, _ := hashMap.Get(key)
	return val, nil
}

// and clears the journal as well as the refunds.
func (a *AccountKVS) FinaliseObject(account string, deleteEmptyObjects bool) {
	obj, exist := a.objects[account]
	if !exist {
		return
	}

	a.updateObject(account, obj)
}

func (a *AccountKVS) updateObject(account string, obj *linkedHashMap.LinkedHashMap) {
	data, err := rlp.EncodeToBytes(obj)
	if err != nil {
		panic(fmt.Errorf("can't encode object at %s: %v", account, err))
	}

	a.setError(a.trie.TryUpdate([]byte(account), data))
}

func (a *AccountKVS) CommitObject(account string, deleteEmptyObjects bool) {
	obj, exist := a.objects[account]
	if !exist {
		return
	}

	a.updateObject(account, obj)
}

func (a *AccountKVS) IntermediateRoot() types.Hash {
	return a.trie.Hash()
}

func (a *AccountKVS) CommitTree() (types.Hash, error) {
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
