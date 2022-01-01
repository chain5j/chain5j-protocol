// Package vmdb
//
// @author: xwc1125
package vmdb

import (
	"bytes"
	"fmt"
	"github.com/chain5j/chain5j-pkg/codec/rlp"
	"github.com/chain5j/chain5j-pkg/crypto/hashalg/sha3"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-pkg/util/hexutil"
	"github.com/chain5j/chain5j-protocol/models/statetype"
	"github.com/chain5j/chain5j-protocol/pkg/database/basedb"
	"github.com/chain5j/chain5j-protocol/pkg/database/statedb/vmdb/model"
	"io"
	"math/big"
)

type stateObject struct {
	db   *StateDB
	tree basedb.Tree

	address  types.DomainAddress // 地址
	addrHash types.Hash          // 地址对应的hash
	data     model.Account       // 账户信息

	dbErr error

	code         model.Code // 合约代码byteCodes，合约进行set,get时使用
	contractInfo []byte     // 合约信息

	originStorage  model.Storage // 缓存原始数据
	pendingStorage model.Storage // 即将需要写入磁盘的数据
	dirtyStorage   model.Storage // 当前交易需要进行更新的数据

	dirtyCode bool // 是否有数据需要进行更新
	suicided  bool //如果标记为true时，那么在进行update时，将会被剔除
	deleted   bool
}

func newObject(db *StateDB, address types.DomainAddress, data model.Account) *stateObject {
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
		addrHash:       types.BytesToHash(sha3.Keccak256(address.Bytes())),
		data:           data,
		originStorage:  make(model.Storage),
		pendingStorage: make(model.Storage),
		dirtyStorage:   make(model.Storage),
	}
}
func (s *stateObject) empty() bool {
	return s.data.IsEmpty()
}

// EncodeRLP implements rlp.Encoder.
func (s *stateObject) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &s.data)
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
	if s.address.Addr == ripemd {
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
func (s *stateObject) SubBalance(amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	s.SetBalance(new(big.Int).Sub(s.Balance(), amount))
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
	return stateObject
}

//
// Attribute accessors
//

// Returns the address of the contract/account
func (s *stateObject) Address() types.DomainAddress {
	return s.address
}

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

func (s *stateObject) CodeHash() []byte {
	return s.data.CodeHash
}

func (s *stateObject) Balance() *big.Int {
	return s.data.Balance
}

func (s *stateObject) Nonce() uint64 {
	return s.data.Nonce
}

// Never called, but must be present to allow stateObject to be used
// as a vm.Account interface that also satisfies the chain5j-vm.ContractRef
// interface. Interfaces are awesome.
func (s *stateObject) Value() *big.Int {
	panic("Value on stateObject should never be called")
}
func (c *stateObject) AddTokenBalance(token types.DomainAddress, amount *big.Int) {
	// EIP158: We must check emptiness for the objects such that the account
	// clearing (0,0,0 objects) can take effect.
	if amount.Sign() == 0 {
		if c.empty() {
			c.touch()
		}

		return
	}
	c.SetTokenBalance(token, new(big.Int).Add(c.TokenBalance(token), amount))
}

func (c *stateObject) SubTokenBalance(token types.DomainAddress, amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	c.SetTokenBalance(token, new(big.Int).Sub(c.TokenBalance(token), amount))
}

func (c *stateObject) SetTokenBalance(token types.DomainAddress, amount *big.Int) {
	if token == types.EmptyDomainAddress {
		c.SetBalance(amount)
		return
	}
	//if _, ok := c.data.Tokens[token]; !ok {
	//	c.data.Tokens[token] = big.NewInt(0)
	//}
	//c.db.journal.append(tokenBalanceChange{
	//	account: &c.address,
	//	token:   &token,
	//	prev:    new(big.Int).Set(c.data.Tokens[token]),
	//})
	c.setTokenBalance(token, amount)
}

func (c *stateObject) setTokenBalance(token types.DomainAddress, amount *big.Int) {
	//c.data.Tokens[token] = amount
}
func (c *stateObject) TokenBalance(token types.DomainAddress) *big.Int {
	//if token == types.EmptyDomainAddress {
	//	return c.data.Balance
	//}
	//if balance, ok := c.data.Tokens[token]; ok {
	//	return balance
	//}
	return big.NewInt(0)
}

func (c *stateObject) TokenBalances() statetype.TokenValues {
	//tv := make(statetype.TokenValues, 0, len(c.data.Tokens)+1)
	//if c.data.Balance.Sign() > 0 {
	//	tv = append(tv, statetype.TokenValue{
	//		TokenAddr: types.EmptyDomainAddress,
	//		Value:     big.NewInt(0).Set(c.data.Balance),
	//	})
	//}
	//for addr, val := range c.data.Tokens {
	//	if val.Sign() > 0 {
	//		tv = append(tv, statetype.TokenValue{
	//			TokenAddr: addr,
	//			Value:     big.NewInt(0).Set(val),
	//		})
	//	}
	//}
	//return tv
	return nil
}

// 合约信息
func (s *stateObject) ContractInfo(db basedb.Database) []byte {
	if s.contractInfo != nil {
		return s.contractInfo
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

func (s *stateObject) SetContractInfo(codeHash types.Hash, code []byte) {
	prevcode := s.Code(s.db.db)
	s.db.journal.append(codeChange{
		account:  &s.address,
		prevhash: s.CodeHash(),
		prevcode: prevcode,
	})
	s.setContractInfo(codeHash, code)
}

func (s *stateObject) setContractInfo(codeHash types.Hash, code []byte) {
	s.code = code
	s.data.CodeHash = codeHash[:]
	s.dirtyCode = true
}
