// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package ethStatedb

import (
	"bytes"
	"github.com/chain5j/chain5j-pkg/codec/rlp"
	"github.com/chain5j/chain5j-pkg/collection/trees/tree"
	"github.com/chain5j/chain5j-pkg/database/kvstore"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/pkg/database/ethStatedb/model"
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
