// Package accountdb
//
// @author: xwc1125
package accountdb

import (
	"bytes"
	"fmt"
	"io"
	"math/big"

	"github.com/chain5j/chain5j-pkg/codec/rlp"
	"github.com/chain5j/chain5j-pkg/crypto/hashalg/sha3"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/models/accounts"
	"github.com/chain5j/chain5j-protocol/pkg/database/basedb"
	"github.com/chain5j/logger"
)

type Code []byte

func (c Code) String() string {
	return string(c) // strings.Join(Disassemble(self), " ")
}

type Storage map[types.Hash]types.Hash

func (s Storage) String() (str string) {
	for key, value := range s {
		str += fmt.Sprintf("%X : %X\n", key, value)
	}

	return
}

func (s Storage) Copy() Storage {
	cpy := make(Storage)
	for key, value := range s {
		cpy[key] = value
	}

	return cpy
}

type contractObject struct {
	tree basedb.Tree

	address  types.Address // 地址
	addrHash types.Hash    // 地址对应的hash

	code Code

	originStorage  Storage
	pendingStorage Storage // 即将需要写入磁盘的数据
	dirtyStorage   Storage // 当前交易需要进行更新的数据

	dirtyCode bool // 是否有数据需要进行更新
	suicided  bool // 如果标记为true时，那么在进行update时，将会被剔除
	deleted   bool
}

type accountObject struct {
	db      *AccountDB
	account string                 // 账户名称
	data    *accounts.AccountStore // 账户信息
	journal *basedb.Journal

	isContract bool
	contract   *contractObject
}

func newContractObject(account *accounts.AccountStore) *contractObject {
	addr := types.HexToAddress(account.CN)

	return &contractObject{
		address:        addr,
		addrHash:       types.BytesToHash(sha3.Keccak256(addr[:])),
		originStorage:  make(Storage),
		pendingStorage: make(Storage),
		dirtyStorage:   make(Storage),
	}
}

func newObject(db *AccountDB, account string, data *accounts.AccountStore) *accountObject {
	if data.Balance == nil {
		data.Balance = new(big.Int)
	}

	var (
		isContract bool
		contract   *contractObject
	)

	if data.IsContract() {
		isContract = true
		contract = newContractObject(data)
	}

	return &accountObject{
		db:         db,
		account:    account,
		data:       data,
		isContract: isContract,
		contract:   contract,
		journal:    db.journal,
	}
}

func (a *accountObject) IsContract() bool {
	return a.isContract
}

func (a *accountObject) empty() bool {
	return a.data.Nonce == 0 && a.data.Balance.Sign() == 0
}

// EncodeRLP implements rlp.Encoder.
func (a *accountObject) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &a.data)
}

func (a *accountObject) touch() {
	a.journal.Append(touchChange{
		account: a.account,
	}, a.db)
}

// AddBalance removes amount from c's balance.
// It is used to add funds to the destination account of a transfer.
func (a *accountObject) AddBalance(amount *big.Int) {
	// EIP158: We must check emptiness for the objects such that the account
	// clearing (0,0,0 objects) can take effect.
	if amount.Sign() == 0 {
		if a.empty() {
			a.touch()
		}

		return
	}
	a.SetBalance(new(big.Int).Add(a.Balance(), amount))
}

// SubBalance removes amount from c's balance.
// It is used to remove funds from the origin account of a transfer.
func (a *accountObject) SubBalance(amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	a.SetBalance(new(big.Int).Sub(a.Balance(), amount))
}

func (a *accountObject) SetAddress(address types.Address) {
	if _, ok := a.data.Addresses[address]; ok {
		return
	}
	a.db.journal.Append(addressChange{
		db:      a.db,
		account: a.account,
		set:     address,
	}, a.db)
	a.setAddress(address)
}

func (a *accountObject) setAddress(address types.Address) {
	a.data.Addresses[address] = &accounts.AddressStore{}
}

func (a *accountObject) SetBalance(amount *big.Int) {
	a.db.journal.Append(balanceChange{
		db:      a.db,
		account: a.account,
		prev:    new(big.Int).Set(a.data.Balance),
	}, a.db)
	a.setBalance(amount)
}

func (a *accountObject) setBalance(amount *big.Int) {
	a.data.Balance = amount
}

func (a *accountObject) deepCopy(db *AccountDB) *accountObject {
	stateObject := newObject(db, a.account, a.data)

	return stateObject
}

func (a *accountObject) AccountName() string {
	return a.account
}

func (a *accountObject) AccountStore() *accounts.AccountStore {
	return a.data
}

func (a *accountObject) SetNonce(nonce uint64) {
	a.db.journal.Append(nonceChange{
		db:      a.db,
		account: a.account,
		prev:    a.data.Nonce,
	}, a.db)
	a.setNonce(nonce)
}

func (a *accountObject) setNonce(nonce uint64) {
	a.data.Nonce = nonce
}

func (a *accountObject) FrozenAccount(frozen bool) {
	a.db.journal.Append(frozenChange{
		db:      a.db,
		account: a.account,
		prev:    a.data.IsFrozen,
	}, a.db)
	a.frozenAccount(frozen)
}

func (a *accountObject) frozenAccount(frozen bool) {
	a.data.IsFrozen = frozen
}

func (a *accountObject) UpdatePermission(permissions *accounts.Permissions) {
	a.journal.Append(permissionChange{
		db:      a.db,
		account: a.account,
		prev:    a.data.Permissions,
	}, a.db)
	a.updatePermission(permissions)
}

func (a *accountObject) updatePermission(permissions *accounts.Permissions) {
	a.data.Permissions = permissions
}

func (a *accountObject) SetPartner(data accounts.PartnerData) {
	a.journal.Append(xxxChange{
		db:       a.db,
		account:  a.account,
		key:      accounts.PartnerKey,
		preValue: a.data.XXX[accounts.PartnerKey],
	}, a.db)

	storeBytes, _ := rlp.EncodeToBytes(&data)
	a.setXXX(accounts.PartnerKey, storeBytes)
}

func (a *accountObject) SetLostStore(data *accounts.LostStore) {
	a.journal.Append(xxxChange{
		db:       a.db,
		account:  a.account,
		key:      accounts.LostKey,
		preValue: a.data.XXX[accounts.LostKey],
	}, a.db)

	var storeBytes []byte
	if data != nil {
		storeBytes, _ = rlp.EncodeToBytes(data)
	}

	a.setXXX(accounts.LostKey, storeBytes)
}

func (a *accountObject) setXXX(key string, value []byte) {
	if len(value) == 0 {
		delete(a.data.XXX, key)
	}

	a.data.XXX[key] = value
}

func (a *accountObject) Balance() *big.Int {
	return a.data.Balance
}

func (a *accountObject) Nonce() uint64 {
	return a.data.Nonce
}

// Never called, but must be present to allow stateObject to be used
// as a vm.AccountName interface that also satisfies the chain5j-vm.ContractRef
// interface. Interfaces are awesome.
func (a *accountObject) Value() *big.Int {
	panic("Value on stateObject should never be called")
}

func (a *accountObject) markSuicided() {
	a.contract.suicided = true
}

func (a *accountObject) getTree(db basedb.Database) basedb.Tree {
	if a.contract.tree == nil {
		var err error
		a.contract.tree, err = db.OpenStorageTree(a.contract.addrHash, a.data.StorageRoot())
		if err != nil {
			a.contract.tree, _ = db.OpenStorageTree(a.contract.addrHash, types.Hash{})
			// a.setError(fmt.Errorf("can't create storage trie: %v", err))
		}
	}
	return a.contract.tree
}

// GetState retrieves a value from the account storage trie.
func (a *accountObject) GetState(db basedb.Database, key types.Hash) types.Hash {
	// If we have a dirty value for this state entry, return it
	value, dirty := a.contract.dirtyStorage[key]
	if dirty {
		return value
	}
	// Otherwise return the entry'a original value
	return a.GetCommittedState(db, key)
}

// GetCommittedState retrieves a value from the committed account storage trie.
func (a *accountObject) GetCommittedState(db basedb.Database, key types.Hash) types.Hash {
	// If we have a pending write or clean cached, return that
	if value, pending := a.contract.pendingStorage[key]; pending {
		return value
	}
	if value, cached := a.contract.originStorage[key]; cached {
		return value
	}
	// Otherwise load the value from the database
	enc, err := a.getTree(db).TryGet(key[:])
	if err != nil {
		return types.Hash{}
	}
	var value types.Hash
	if len(enc) > 0 {
		_, content, _, err := rlp.Split(enc)
		if err != nil {
		}
		value.SetBytes(content)
	}
	a.contract.originStorage[key] = value
	return value
}

// SetState updates a value in account storage.
func (a *accountObject) SetState(db basedb.Database, key, value types.Hash) {
	// If the new value is the same as old, don't set
	prev := a.GetState(db, key)
	if prev == value {
		return
	}
	// New value is different, update and journal the change
	a.journal.Append(storageChange{
		db:       a.db,
		account:  a.account,
		key:      key,
		prevalue: prev,
	}, a.db)
	a.setState(key, value)
}

func (a *accountObject) setState(key, value types.Hash) {
	a.contract.dirtyStorage[key] = value
}

// Code returns the contract code associated with this object, if any.
func (a *accountObject) Code(db basedb.Database) []byte {
	if a.contract.code != nil {
		return a.contract.code
	}
	if bytes.Equal(a.CodeHash(), types.EmptyCode) {
		return nil
	}
	logger.Trace("Get code", "hash", types.BytesToHash(a.CodeHash()))
	code, err := db.ContractCode(a.contract.addrHash, types.BytesToHash(a.CodeHash()))
	if err != nil {
		// a.setError(fmt.Errorf("can't load code hash %x: %v", a.CodeHash(), err))
	}
	a.contract.code = code
	return code
}

func (a *accountObject) CodeHash() []byte {
	return a.data.CodeHash()
}

func (a *accountObject) SetCode(codeHash types.Hash, code []byte) {
	prevcode := a.Code(a.db.db)
	a.journal.Append(codeChange{
		db:       a.db,
		account:  a.account,
		prevhash: a.CodeHash(),
		prevcode: prevcode,
	}, a.db)
	a.setCode(codeHash, code)
}

func (a *accountObject) setCode(codeHash types.Hash, code []byte) {
	a.contract.code = code
	a.data.SetCodeHash(codeHash[:])
	a.contract.dirtyCode = true
}

func (a *accountObject) finalise() {
	for key, value := range a.contract.dirtyStorage {
		a.contract.pendingStorage[key] = value
	}
	if len(a.contract.dirtyStorage) > 0 {
		a.contract.dirtyStorage = make(Storage)
	}
}

// UpdateRoot sets the trie root to the current root hash of
func (a *accountObject) updateRoot(db basedb.Database) {
	a.updateTrie(db)
	a.data.SetStorageRoot(a.contract.tree.Hash())
}

// updateTrie writes cached storage modifications into the object's storage trie.
func (a *accountObject) updateTrie(db basedb.Database) basedb.Tree {
	// Make sure all dirty slots are finalized into the pending storage area
	a.finalise()

	tr := a.getTree(db)
	for key, value := range a.contract.pendingStorage {
		delete(a.contract.dirtyStorage, key)

		// Skip noop changes, persist actual changes
		if value == a.contract.originStorage[key] {
			continue
		}
		a.contract.originStorage[key] = value

		if (value == types.Hash{}) {
			tr.TryDelete(key[:])
			continue
		}
		// Encoding []byte cannot fail, ok to ignore the error.
		v, _ := rlp.EncodeToBytes(bytes.TrimLeft(value[:], "\x00"))
		tr.TryUpdate(key[:], v)
	}

	if len(a.contract.pendingStorage) > 0 {
		a.contract.pendingStorage = make(Storage)
	}
	return tr
}

func (a *accountObject) CommitTree(db basedb.Database) {
	a.updateTrie(db)

	root, err := a.contract.tree.Commit(nil)
	if err == nil {
		a.data.SetStorageRoot(root)
	}
}
