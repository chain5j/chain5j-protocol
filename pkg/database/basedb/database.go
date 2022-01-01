// Package basedb
//
// @author: xwc1125
package basedb

import (
	"github.com/chain5j/chain5j-pkg/collection/trees/tree"
	"github.com/chain5j/chain5j-pkg/database/kvstore"
	"github.com/chain5j/chain5j-pkg/types"
	lru "github.com/hashicorp/golang-lru"
)

const (
	// Number of codehash->size associations to keep.
	codeSizeCacheSize = 100000
)

type Database interface {
	// OpenTree 打开主账户树
	OpenTree(root types.Hash) (Tree, error)
	// OpenStorageTree 打开账户树
	OpenStorageTree(addrHash types.Hash, root types.Hash) (Tree, error)
	// CopyTree 复制树
	CopyTree(Tree) Tree

	// ContractCode 获取合约的bytes
	ContractCode(addrHash, codeHash types.Hash) ([]byte, error)

	// ContractCodeSize 获取合约的bytes长度
	ContractCodeSize(addrHash, codeHash types.Hash) (int, error)

	// TreeDB 获取tree database
	TreeDB() *tree.Database
}

func NewDatabase(db kvstore.Database) Database {
	return NewDatabaseWithCache(db, 0)
}

// NewDatabaseWithCache creates a backing store for state. The returned database
// is safe for concurrent use and retains a lot of collapsed RLP trie nodes in a
// large memory cache.
func NewDatabaseWithCache(db kvstore.Database, cache int) Database {
	csc, _ := lru.New(codeSizeCacheSize)
	return &cachingDB{
		db:            tree.NewDatabaseWithCache(db, cache),
		codeSizeCache: csc,
	}
}
