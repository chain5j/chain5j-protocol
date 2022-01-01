// Package ethStatedb
//
// @author: xwc1125
package ethStatedb

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/chain5j/chain5j-pkg/codec/rlp"
	linkedHashMap "github.com/chain5j/chain5j-pkg/collection/maps/linked_hashmap"
	"github.com/chain5j/chain5j-pkg/crypto/hashalg/sha3"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-pkg/util/hexutil"
	"github.com/chain5j/chain5j-protocol/pkg/database/basedb"
	"github.com/chain5j/chain5j-protocol/pkg/database/ethStatedb/model"
	"io"
	"math/big"
)

type stateObject struct {
	db    *StateDB
	tree  basedb.Tree
	dbErr error

	address  types.Address // 地址
	addrHash types.Hash    // 地址对应的hash
	data     model.Account // 账户信息

	code    model.Code                   // 合约代码byteCodes，合约进行set,get时使用
	kvsHash []byte                       // kvs hash
	kvs     *linkedHashMap.LinkedHashMap // 额外存储的相关数据（namespace==>key==>value）

	originStorage  model.Storage // 缓存原始数据
	pendingStorage model.Storage // 即将需要写入磁盘的数据
	dirtyStorage   model.Storage // 当前交易需要进行更新的数据

	dirtyKvs  bool // kvs是否有更新
	dirtyCode bool // 是否有数据需要进行更新
	suicided  bool //如果标记为true时，那么在进行update时，将会被剔除
	deleted   bool
}

func newObject(db *StateDB, address types.Address, data model.Account) *stateObject {
	if data.Balance == nil {
		data.Balance = new(big.Int)
	}
	if data.CodeHash == nil {
		data.CodeHash = types.EmptyCode
	}
	if data.Root.Nil() {
		data.Root = types.EmptyRootHash
	}

	return &stateObject{
		db:             db,
		address:        address,
		addrHash:       types.BytesToHash(sha3.Keccak256(address[:])),
		data:           data,
		kvsHash:        types.EmptyCode,
		kvs:            linkedHashMap.NewLinkedHashMap(),
		originStorage:  make(model.Storage),
		pendingStorage: make(model.Storage),
		dirtyStorage:   make(model.Storage),
	}
}

func (s *stateObject) empty() bool {
	return s.data.Nonce == 0 && s.data.Balance.Sign() == 0 && bytes.Equal(s.data.CodeHash, types.EmptyCode) && bytes.Equal(s.kvsHash, types.EmptyCode)
}

// EncodeRLP implements rlp.Encoder.
func (s *stateObject) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{
		s.address,
		s.addrHash,
		s.data,
		s.code,
		s.kvsHash,
		s.kvs,
	})
}

func (s *stateObject) DecodeRLP(st *rlp.Stream) error {
	var obj struct {
		Address  types.Address
		AddrHash types.Hash
		Data     model.Account
		Code     model.Code
		KvsHash  []byte
		Kvs      *linkedHashMap.LinkedHashMap
	}

	if err := st.Decode(&obj); err != nil {
		return err
	}
	s.address = obj.Address
	s.addrHash = obj.AddrHash
	s.data = obj.Data
	s.code = obj.Code
	s.kvsHash = obj.KvsHash
	s.kvs = obj.Kvs
	return nil
}

// setError remembers the first non-nil error it is called with.
func (s *stateObject) setError(err error) {
	if s.dbErr == nil {
		s.dbErr = err
	}
}

func (s *stateObject) markSuicided() {
	s.suicided = true
}

func (s *stateObject) touch() {
	s.db.journal.append(touchChange{
		account: &s.address,
	})
	if s.address == ripemd {
		// Explicitly put it in the dirty-cache, which is otherwise generated from
		// flattened journals.
		s.db.journal.dirty(s.address)
	}
}

func (s *stateObject) getTree(db basedb.Database) basedb.Tree {
	if s.tree == nil {
		var err error
		s.tree, err = db.OpenStorageTree(s.addrHash, s.data.Root)
		if err != nil {
			s.tree, _ = db.OpenStorageTree(s.addrHash, types.Hash{})
			s.setError(fmt.Errorf("can't create storage trie: %v", err))
		}
	}
	return s.tree
}

// ===========State===========
// GetState retrieves a value from the account storage trie.
func (s *stateObject) GetState(db basedb.Database, key types.Hash) types.Hash {
	// If we have a dirty value for this state entry, return it
	value, dirty := s.dirtyStorage[key]
	if dirty {
		return value
	}
	// Otherwise return the entry's original value
	return s.GetCommittedState(db, key)
}

// GetCommittedState retrieves a value from the committed account storage trie.
func (s *stateObject) GetCommittedState(db basedb.Database, key types.Hash) types.Hash {
	// If we have a pending write or clean cached, return that
	if value, pending := s.pendingStorage[key]; pending {
		return value
	}
	if value, cached := s.originStorage[key]; cached {
		return value
	}
	// Otherwise load the value from the database
	enc, err := s.getTree(db).TryGet(key[:])
	if err != nil {
		s.setError(err)
		return types.Hash{}
	}
	var value types.Hash
	if len(enc) > 0 {
		_, content, _, err := rlp.Split(enc)
		if err != nil {
			s.setError(err)
		}
		value.SetBytes(content)
	}
	s.originStorage[key] = value
	return value
}

// SetState updates a value in account storage.
func (s *stateObject) SetState(db basedb.Database, key, value types.Hash) {
	// If the new value is the same as old, don't set
	prev := s.GetState(db, key)
	if prev == value {
		return
	}
	// New value is different, update and journal the change
	s.db.journal.append(storageChange{
		account:  &s.address,
		key:      key,
		prevalue: prev,
	})
	s.setState(key, value)
}

func (s *stateObject) setState(key, value types.Hash) {
	s.dirtyStorage[key] = value
}

// finalise moves all dirty storage slots into the pending area to be hashed or
// committed later. It is invoked at the end of every transaction.
func (s *stateObject) finalise() {
	for key, value := range s.dirtyStorage {
		s.pendingStorage[key] = value
	}
	if len(s.dirtyStorage) > 0 {
		s.dirtyStorage = make(model.Storage)
	}
}

// updateTrie writes cached storage modifications into the object's storage trie.
func (s *stateObject) updateTree(db basedb.Database) basedb.Tree {
	// Make sure all dirty slots are finalized into the pending storage area
	s.finalise()

	// Insert all the pending updates into the Tree
	tr := s.getTree(db)
	for key, value := range s.pendingStorage {
		// Skip noop changes, persist actual changes
		if value == s.originStorage[key] {
			continue
		}
		s.originStorage[key] = value

		if (value == types.Hash{}) {
			s.setError(tr.TryDelete(key[:]))
			continue
		}
		// Encoding []byte cannot fail, ok to ignore the error.
		v, _ := rlp.EncodeToBytes(hexutil.TrimLeftZeroes(value[:]))
		s.setError(tr.TryUpdate(key[:], v))
	}
	if len(s.pendingStorage) > 0 {
		s.pendingStorage = make(model.Storage)
	}
	return tr
}

// UpdateRoot sets the Tree root to the current root hash of
func (s *stateObject) updateRoot(db basedb.Database) {
	s.updateTree(db)

	s.data.Root = s.tree.Hash()
}

// CommitTree the storage Tree of the object to db.
// This updates the Tree root.
func (s *stateObject) CommitTree(db basedb.Database) error {
	s.updateTree(db)
	if s.dbErr != nil {
		return s.dbErr
	}
	root, err := s.tree.Commit(nil)
	if err == nil {
		s.data.Root = root
	}
	return err
}

// ===========Balance===========
// AddBalance removes amount from c's balance.
// It is used to add funds to the destination account of a transfer.
func (s *stateObject) AddBalance(amount *big.Int) {
	// EIP158: We must check emptiness for the objects such that the account
	// clearing (0,0,0 objects) can take effect.
	if amount.Sign() == 0 {
		if s.empty() {
			s.touch()
		}

		return
	}
	s.SetBalance(new(big.Int).Add(s.Balance(), amount))
}

// SubBalance removes amount from c's balance.
// It is used to remove funds from the origin account of a transfer.
func (s *stateObject) Balance() *big.Int {
	return s.data.Balance
}

func (s *stateObject) SetBalance(amount *big.Int) {
	s.db.journal.append(balanceChange{
		account: &s.address,
		prev:    new(big.Int).Set(s.data.Balance),
	})
	s.setBalance(amount)
}

func (s *stateObject) setBalance(amount *big.Int) {
	s.data.Balance = amount
}

func (s *stateObject) SubBalance(amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	s.SetBalance(new(big.Int).Sub(s.Balance(), amount))
}

// ===========Gas===========
// Return the gas back to the origin. Used by the Virtual machine or Closures
func (s *stateObject) ReturnGas(gas *big.Int) {}

func (s *stateObject) deepCopy(db *StateDB) *stateObject {
	stateObject := newObject(db, s.address, s.data)
	if s.tree != nil {
		stateObject.tree = db.db.CopyTree(s.tree)
	}
	stateObject.code = s.code
	stateObject.dirtyStorage = s.dirtyStorage.Copy()
	stateObject.originStorage = s.originStorage.Copy()
	stateObject.pendingStorage = s.pendingStorage.Copy()
	stateObject.suicided = s.suicided
	stateObject.dirtyCode = s.dirtyCode
	stateObject.deleted = s.deleted
	stateObject.kvs = s.kvs
	stateObject.dirtyKvs = s.dirtyKvs
	return stateObject
}

//
// Attribute accessors
//

// Returns the address of the contract/account
func (s *stateObject) Address() types.Address {
	return s.address
}

// ===========code===========

// Code returns the contract code associated with this object, if any.
func (s *stateObject) Code(db basedb.Database) []byte {
	if s.code != nil {
		return s.code
	}
	if bytes.Equal(s.CodeHash(), types.EmptyCode) {
		return nil
	}
	code, err := db.ContractCode(s.addrHash, types.BytesToHash(s.CodeHash()))
	if err != nil {
		s.setError(fmt.Errorf("can't load code hash %x: %v", s.CodeHash(), err))
	}
	s.code = code
	return code
}

func (s *stateObject) SetCode(codeHash types.Hash, code []byte) {
	prevcode := s.Code(s.db.db)
	s.db.journal.append(codeChange{
		account:  &s.address,
		prevhash: s.CodeHash(),
		prevcode: prevcode,
	})
	s.setCode(codeHash, code)
}

func (s *stateObject) setCode(codeHash types.Hash, code []byte) {
	s.code = code
	s.data.CodeHash = codeHash[:]
	s.dirtyCode = true
}

func (s *stateObject) CodeHash() []byte {
	return s.data.CodeHash
}

//===========Nonce===========
func (s *stateObject) Nonce() uint64 {
	return s.data.Nonce
}

func (s *stateObject) SetNonce(nonce uint64) {
	s.db.journal.append(nonceChange{
		account: &s.address,
		prev:    s.data.Nonce,
	})
	s.setNonce(nonce)
}

func (s *stateObject) setNonce(nonce uint64) {
	s.data.Nonce = nonce
}

// kvs
func (s *stateObject) KVHash() []byte {
	return s.kvsHash
}

func (s *stateObject) AddKV(namespace string, key string, val interface{}) {
	nameMap, _ := s.kvs.Get(namespace)
	if nameMap == nil {
		nameMap = linkedHashMap.NewLinkedHashMap()
		s.kvs.Add(namespace, nameMap)
	}
	hashMap := nameMap.(*linkedHashMap.LinkedHashMap)
	hashMap.Add(key, val)
	s.dirtyKvs = true
}

func (s *stateObject) setKVS(kvsHash types.Hash, kvs *linkedHashMap.LinkedHashMap) {
	s.kvs = kvs
	s.kvsHash = kvsHash[:]
	s.dirtyKvs = true
}

func (s *stateObject) setKVNP(kvsHash types.Hash, namespace string, kvnp *linkedHashMap.LinkedHashMap) {
	s.kvs.Add(namespace, kvnp)
	s.kvsHash = kvsHash[:]
	s.dirtyKvs = true
}

func (s *stateObject) KVS() *linkedHashMap.LinkedHashMap {
	return s.kvs
}

func (s *stateObject) KVNP(namespace string) (*linkedHashMap.LinkedHashMap, error) {
	nameMap, _ := s.kvs.Get(namespace)
	if nameMap == nil {
		return nil, errors.New("namespace is nil")
	}
	hashMap := nameMap.(*linkedHashMap.LinkedHashMap)
	return hashMap, nil
}

func (s *stateObject) KV(namespace string, key string) (interface{}, error) {
	nameMap, _ := s.kvs.Get(namespace)
	if nameMap == nil {
		return nil, errors.New("namespace is nil")
	}
	hashMap := nameMap.(*linkedHashMap.LinkedHashMap)
	val, _ := hashMap.Get(key)
	return val, nil
}

func (s *stateObject) KVExist(namespace string, key string) bool {
	nameMap, _ := s.kvs.Get(namespace)
	if nameMap == nil {
		return false
	}
	hashMap := nameMap.(*linkedHashMap.LinkedHashMap)
	_, b := hashMap.Get(key)
	return b
}

// Never called, but must be present to allow stateObject to be used
// as a vm.AccountName interface that also satisfies the chain5j-vm.ContractRef
// interface. Interfaces are awesome.
func (s *stateObject) Value() *big.Int {
	panic("Value on stateObject should never be called")
}
