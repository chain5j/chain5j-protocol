// Package vmdb
//
// @author: xwc1125
package vmdb

import (
	"github.com/chain5j/chain5j-pkg/collection/trees/tree"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/models/statetype"
	"math/big"
)

type StateDBFace interface {
	Error() error
	Reset(root types.Hash) error
	Exist(addr types.Address) bool
	Empty(addr types.Address) bool
	Copy() *StateDB

	GetState(addr types.Address, hash types.Hash) types.Hash
	GetCommittedState(addr types.Address, hash types.Hash) types.Hash
	SetState(addr types.Address, key, value types.Hash)

	GetProof(a types.Address) ([][]byte, error)
	GetStorageProof(a types.Address, key types.Hash) ([][]byte, error)

	TxIndex() int
	BlockHash() types.Hash

	AddPreimage(hash types.Hash, preimage []byte)
	Preimages() map[types.Hash][]byte

	StorageTree(addr types.Address) tree.Tree

	HasSuicided(addr types.Address) bool
	Suicide(addr types.Address) bool

	Snapshot() int
	RevertToSnapshot(revid int)

	IntermediateRoot(deleteEmptyObjects bool) types.Hash
	Finalise(deleteEmptyObjects bool)
	Prepare(thash, bhash types.Hash, ti int)
	Commit(deleteEmptyObjects bool) (types.Hash, error)
}

type BzStateDb interface {
	AddBalance(addr types.Address, amount *big.Int)
	SubBalance(addr types.Address, amount *big.Int)
	GetBalance(addr types.Address) *big.Int
	SetBalance(addr types.Address, amount *big.Int)

	GetNonce(addr types.Address) *big.Int
	SetNonce(addr types.Address, nonce *big.Int)

	AddLog(log *statetype.Log)
	GetLogs(hash types.Hash) []*statetype.Log
	Logs() []*statetype.Log

	AddRefund(gas uint64)
	SubRefund(gas uint64)
	GetRefund() uint64

	GetCode(addr types.Address) []byte
	GetCodeSize(addr types.Address) int
	GetCodeHash(addr types.Address) types.Hash
	SetCode(addr types.Address, code []byte)
}
