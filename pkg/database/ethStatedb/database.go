// Package ethStatedb
//
// @author: xwc1125
package ethStatedb

import (
	"github.com/chain5j/chain5j-pkg/collection/trees/tree"
	"github.com/chain5j/chain5j-pkg/database/kvstore"
	"github.com/chain5j/chain5j-protocol/pkg/database/basedb"
	lru "github.com/hashicorp/golang-lru"
)

const (
	// Number of codehash->size associations to keep.
	codeSizeCacheSize = 100000
)

func NewDatabase(db kvstore.Database) basedb.Database {
	return NewDatabaseWithCache(db, 0)
}

// NewDatabaseWithCache creates a backing store for state. The returned database
// is safe for concurrent use and retains a lot of collapsed RLP trie nodes in a
// large memory cache.
func NewDatabaseWithCache(db kvstore.Database, cache int) basedb.Database {
	csc, _ := lru.New(codeSizeCacheSize)
	return &cachingDB{
		db:            tree.NewDatabaseWithCache(db, cache),
		codeSizeCache: csc,
	}
}
