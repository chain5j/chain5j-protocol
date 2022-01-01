// Package statedb
//
// @author: xwc1125
package statedb

import (
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/pkg/database/basedb"
)

// xStore用于不经常使用的数据存储，不设置回滚。
// 存储前应有足够的条件确保数据没问题
type xStore struct {
	db      basedb.Database
	trie    basedb.Tree
	preRoot types.Hash

	cache map[string][]byte

	store   map[string][]byte
	unStore map[string]struct{}
}

func NewXStore(root types.Hash, db basedb.Database) (*xStore, error) {
	tr, err := db.OpenTree(root)
	if err != nil {
		return nil, err
	}

	return &xStore{
		db:      db,
		trie:    tr,
		preRoot: root,
		cache:   make(map[string][]byte),
		store:   make(map[string][]byte),
		unStore: make(map[string]struct{}),
	}, nil
}

func (x *xStore) Store(key string, value []byte) {
	x.store[key] = value
}

func (x *xStore) UnStore(key string) {
	x.unStore[key] = struct{}{}
}

func (x *xStore) Get(key string) ([]byte, error) {
	if v, ok := x.cache[key]; ok {
		return v, nil
	}

	v, err := x.trie.TryGet([]byte(key))
	if err != nil {
		return nil, err
	}

	x.cache[key] = v

	return v, nil
}

func (x *xStore) IntermediateRoot() types.Hash {
	if len(x.store) == 0 && len(x.unStore) == 0 {
		return x.preRoot
	}

	for k, v := range x.store {
		x.trie.TryUpdate([]byte(k), v)
	}

	for k := range x.unStore {
		x.trie.TryDelete([]byte(k))
	}

	return x.trie.Hash()
}

func (x *xStore) CommitTree() (types.Hash, error) {
	if len(x.store) == 0 && len(x.unStore) == 0 {
		return x.preRoot, nil
	}

	root, err := x.trie.Commit(func(leaf []byte, parent types.Hash) error {
		return nil
	})

	if err != nil {
		return types.Hash{}, err
	}

	err = x.db.TreeDB().Commit(root, false)
	if err != nil {
		return types.Hash{}, err
	}

	return root, nil
}
