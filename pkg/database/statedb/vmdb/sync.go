// Package vmdb
//
// @author: xwc1125
package vmdb

import (
	"bytes"
	"github.com/chain5j/chain5j-pkg/codec/rlp"
	"github.com/chain5j/chain5j-pkg/collection/trees/tree"
	"github.com/chain5j/chain5j-pkg/database/kvstore"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/pkg/database/statedb/vmdb/model"
)

// NewStateSync create a new state trie download scheduler.
func NewStateSync(root types.Hash, database kvstore.KeyValueReader, bloom *tree.SyncBloom) *tree.Sync {
	var syncer *tree.Sync
	callback := func(leaf []byte, parent types.Hash) error {
		var obj model.Account
		if err := rlp.Decode(bytes.NewReader(leaf), &obj); err != nil {
			return err
		}
		syncer.AddSubTrie(obj.Root, 64, parent, nil)
		syncer.AddRawEntry(types.BytesToHash(obj.CodeHash), 64, parent)
		return nil
	}
	syncer = tree.NewSync(root, database, callback, bloom)
	return syncer
}
