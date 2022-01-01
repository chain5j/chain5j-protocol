// Package basedb
//
// @author: xwc1125
package basedb

import (
	"fmt"
	"github.com/chain5j/chain5j-pkg/collection/trees/tree"
	"github.com/chain5j/chain5j-pkg/types"
	lru "github.com/hashicorp/golang-lru"
)

type cachingDB struct {
	db            *tree.Database
	codeSizeCache *lru.Cache
}

// OpenTree opens the main account trie at a specific root hash.
func (db *cachingDB) OpenTree(root types.Hash) (Tree, error) {
	return tree.NewSecure(root, db.db)
}

// OpenStorageTree opens the storage trie of an account.
func (db *cachingDB) OpenStorageTree(addrHash, root types.Hash) (Tree, error) {
	return tree.NewSecure(root, db.db)
}

// CopyTree returns an independent copy of the given trie.
func (db *cachingDB) CopyTree(t Tree) Tree {
	switch t := t.(type) {
	case *tree.SecureTrie:
		return t.Copy()
	default:
		panic(fmt.Errorf("unknown trie type %T", t))
	}
}

// ContractCode retrieves a particular contract's code.
func (db *cachingDB) ContractCode(addrHash, codeHash types.Hash) ([]byte, error) {
	code, err := db.db.Node(codeHash)
	if err == nil {
		db.codeSizeCache.Add(codeHash, len(code))
	}
	return code, err
}

// ContractCodeSize retrieves a particular contracts code's size.
func (db *cachingDB) ContractCodeSize(addrHash, codeHash types.Hash) (int, error) {
	if cached, ok := db.codeSizeCache.Get(codeHash); ok {
		return cached.(int), nil
	}
	code, err := db.ContractCode(addrHash, codeHash)
	return len(code), err
}

// TreeDB retrieves any intermediate trie-node caching layer.
func (db *cachingDB) TreeDB() *tree.Database {
	return db.db
}
