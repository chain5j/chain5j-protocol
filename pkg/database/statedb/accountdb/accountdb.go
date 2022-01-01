// Package accountdb
//
// @author: xwc1125
package accountdb

import (
	"fmt"
	"github.com/chain5j/chain5j-pkg/codec/rlp"
	"github.com/chain5j/chain5j-pkg/collection/trees/tree"
	"github.com/chain5j/chain5j-pkg/crypto/hashalg/sha3"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-pkg/util/dateutil"
	"github.com/chain5j/chain5j-protocol/models/accounts"
	"github.com/chain5j/chain5j-protocol/pkg/database/basedb"
	"github.com/chain5j/logger"
	"math/big"
	"strings"
	"time"
)

type AccountDB struct {
	db   basedb.Database
	trie basedb.Tree

	// This map holds 'live' objects, which will get modified while processing a state transition.
	accountObjects map[string]*accountObject

	journal *basedb.Journal

	logger logger.Logger
}

// NewAccountDB Create a new state from a given trie.
func NewAccountDB(root types.Hash, db basedb.Database, journal *basedb.Journal) (*AccountDB, error) {
	tr, err := db.OpenTree(root)
	if err != nil {
		return nil, err
	}
	return &AccountDB{
		db:             db,
		trie:           tr,
		accountObjects: make(map[string]*accountObject),
		journal:        journal,
		logger:         logger.New("account_db"),
	}, nil
}

// Reset clears out all ephemeral state objects from the state db, but keeps
// the underlying state trie to avoid reloading data for the next operations.
func (a *AccountDB) Reset(root types.Hash) error {
	tr, err := a.db.OpenTree(root)
	if err != nil {
		return err
	}
	a.trie = tr
	a.accountObjects = make(map[string]*accountObject)
	return nil
}

// Exist reports whether the given account address exists in the state.
// Notably this also returns true for suicided accounts.
func (a *AccountDB) Exist(account string) bool {
	return a.getAccountObject(account) != nil
}

// Empty returns whether the state object is either non-existent
// or empty according to the EIP161 specification (balance = nonce = 0)
func (a *AccountDB) Empty(account string) bool {
	so := a.getAccountObject(account)
	return so == nil || so.empty()
}

// GetBalance Retrieve the balance from the given address or 0 if object not found
func (a *AccountDB) GetBalance(account string) *big.Int {
	stateObject := a.getAccountObject(account)
	if stateObject != nil {
		return stateObject.Balance()
	}
	return big.NewInt(0)
}

func (a *AccountDB) GetNonce(account string) uint64 {
	stateObject := a.getAccountObject(account)
	if stateObject != nil {
		return stateObject.Nonce()
	}

	return 0
}

// Database retrieves the low level database supporting the lower level trie ops.
func (a *AccountDB) Database() basedb.Database {
	return a.db
}

/*
 * SETTERS
 */

// AddBalance adds amount to the account associated with addr.
func (a *AccountDB) AddBalance(account string, amount *big.Int) {
	object := a.getAccountObject(account)
	if object != nil {
		object.AddBalance(amount)
	}
}

// SubBalance subtracts amount from the account associated with addr.
func (a *AccountDB) SubBalance(account string, amount *big.Int) {
	object := a.getAccountObject(account)
	if object != nil {
		object.SubBalance(amount)
	}
}

func (a *AccountDB) SetBalance(account string, amount *big.Int) {
	object := a.getAccountObject(account)
	if object != nil {
		object.SetBalance(amount)
	}
}

func (a *AccountDB) SetNonce(account string, nonce uint64) {
	object := a.getAccountObject(account)
	if object != nil {
		object.SetNonce(nonce)
	}
}

func (a *AccountDB) SetAddress(account string, address types.Address) {
	object := a.getAccountObject(account)
	if object != nil {
		object.SetAddress(address)
	}
}

func (a *AccountDB) FrozenAccount(account string) {
	obj := a.getAccountObject(account)
	if obj != nil {
		obj.FrozenAccount(true)
	}
}

func (a *AccountDB) UnFrozenAccount(account string) {
	obj := a.getAccountObject(account)
	if obj != nil {
		obj.FrozenAccount(false)
	}
}

func (a *AccountDB) UpdatePermission(account string, permission *accounts.Permissions) {
	obj := a.getAccountObject(account)
	if obj != nil {
		obj.UpdatePermission(permission)
	}
}

func (a *AccountDB) SetPartner(account string, data accounts.PartnerData) {
	obj := a.getAccountObject(account)
	if obj != nil {
		obj.SetPartner(data)
	}
}

func (a *AccountDB) SetLost(account string, data *accounts.LostStore) {
	obj := a.getAccountObject(account)
	if obj != nil {
		obj.SetLostStore(data)
	}
}

//
// Setting, updating & deleting state object methods.
//

// updateAccountObject writes the given object to the trie.
func (a *AccountDB) updateAccountObject(obj *accountObject) {
	// Encode the account and update the account trie
	account := obj.AccountName()

	data, err := rlp.EncodeToBytes(obj)
	if err != nil {
		panic(fmt.Errorf("can't encode object at %s: %v", account, err))
	}
	a.trie.TryUpdate([]byte(account), data)
}

// getAccountObject retrieves a state object given by the address, returning nil if
// the object is not found or was deleted in this execution context. If you need
// to differentiate between non-existent/just-deleted, use getDeletedAccountObject.
func (a *AccountDB) getAccountObject(account string) *accountObject {
	account = strings.ToLower(account)
	// Prefer live objects if any is available
	if obj := a.accountObjects[account]; obj != nil {
		return obj
	}
	if !strings.Contains(account, accounts.DomainLinkFlag) {
		if obj := a.accountObjects[account+accounts.DomainLinkFlag+accounts.ContractDomain]; obj != nil {
			return obj
		}
	}
	// Load the object from the database
	enc, _ := a.trie.TryGet([]byte(account))
	if len(enc) == 0 {
		if !strings.Contains(account, accounts.DomainLinkFlag) {
			enc, _ = a.trie.TryGet([]byte(account + accounts.DomainLinkFlag + accounts.ContractDomain))
		}
	}
	if len(enc) == 0 {
		return nil
	}

	var data accounts.AccountStore
	if err := rlp.DecodeBytes(enc, &data); err != nil {
		logger.Error("Failed to decode state object", "addr", account, "err", err)
		return nil
	}
	// Insert into the live set
	obj := newObject(a, account, &data)
	a.setAccountObject(obj)
	return obj
}

func (a *AccountDB) setAccountObject(object *accountObject) {
	a.accountObjects[object.AccountName()] = object
}

func (a *AccountDB) GetAccount(account string) *accounts.AccountStore {
	t := time.Now()
	object := a.getAccountObject(account)
	if object == nil {
		return nil
	}
	accountStore := object.AccountStore().Copy()
	a.logger.Debug("GetAccount Elapsed", "elapsed", dateutil.PrettyDuration(time.Since(t)))
	return accountStore
}

func (a *AccountDB) GetCode(account string) []byte {
	object := a.getAccountObject(account)
	if object == nil || !object.isContract {
		return nil
	}

	return object.Code(a.db)
}

func (a *AccountDB) SetCode(account string, code []byte) {
	object := a.getAccountObject(account)
	if object == nil {
		return
	}

	object.SetCode(types.BytesToHash(sha3.Keccak256(code)), code)
}

func (a *AccountDB) GetCodeSize(account string) int {
	object := a.getAccountObject(account)
	if object == nil {
		return 0
	}

	return len(object.Code(a.db))
}

func (a *AccountDB) GetCommittedState(account string, hash types.Hash) types.Hash {
	obj := a.getAccountObject(account)
	if obj != nil {
		return obj.GetCommittedState(a.db, hash)
	}
	return types.Hash{}
}

func (a *AccountDB) GetState(account string, hash types.Hash) types.Hash {
	obj := a.getAccountObject(account)
	if obj != nil {
		return obj.GetState(a.db, hash)
	}
	return types.Hash{}
}

func (a *AccountDB) SetState(account string, key, value types.Hash) {
	obj := a.getAccountObject(account)
	if obj != nil {
		obj.SetState(a.db, key, value)
	}
}

func (a *AccountDB) Suicide(account string) bool {
	obj := a.getAccountObject(account)
	if obj != nil {
		return false
	}

	a.journal.Append(suicideChange{
		db:          a,
		account:     obj.account,
		prev:        obj.contract.suicided,
		prevbalance: new(big.Int).Set(obj.Balance()),
	}, a)

	obj.markSuicided()
	obj.data.Balance = new(big.Int)

	return true
}

func (a *AccountDB) HasSuicided(account string) bool {
	obj := a.getAccountObject(account)
	if obj != nil {
		return obj.contract.suicided
	}
	return false
}

// createObject creates a new state object. If there is an existing account with
// the given address, it is overwritten and returned as the second return value.
func (a *AccountDB) createObject(account *accounts.AccountStore) (newobj, prev *accountObject) {
	cn := account.AccountName()
	prev = a.getAccountObject(account.AccountName()) // Note, prev might have been deleted, we need that!

	newobj = newObject(a, cn, account)
	newobj.setNonce(0) // sets the object to dirty
	if prev == nil {
		a.journal.Append(createObjectChange{a, cn}, a)
	} else {
		a.journal.Append(resetObjectChange{a, prev}, a)
	}
	a.setAccountObject(newobj)
	return newobj, prev
}

func (a *AccountDB) deleteObject(obj *accountObject) {
	a.trie.TryDelete([]byte(obj.account))
}

func (a *AccountDB) CreateAccount(account *accounts.AccountStore) {
	newObj, prev := a.createObject(account)
	if prev != nil {
		newObj.setBalance(prev.data.Balance)
	}
}

func (a *AccountDB) ForEachStorage(account string, cb func(key, value types.Hash) bool) error {
	so := a.getAccountObject(account)
	if so == nil {
		return nil
	}

	if !so.IsContract() {
		return nil
	}

	it := tree.NewIterator(so.getTree(a.db).NodeIterator(nil))

	for it.Next() {
		key := types.BytesToHash(a.trie.GetKey(it.Key))
		if value, dirty := so.contract.dirtyStorage[key]; dirty {
			if !cb(key, value) {
				return nil
			}
			continue
		}

		if len(it.Value) > 0 {
			_, content, _, err := rlp.Split(it.Value)
			if err != nil {
				return err
			}
			if !cb(key, types.BytesToHash(content)) {
				return nil
			}
		}
	}
	return nil
}

// Finalise finalises the state by removing the self destructed objects
// and clears the journal as well as the refunds.
func (a *AccountDB) FinaliseObject(account string, deleteEmptyObjects bool) {
	obj, exist := a.accountObjects[account]
	if !exist {
		return
	}

	if obj.IsContract() {
		if obj.contract.deleted {
			a.deleteObject(obj)
		} else {
			obj.updateRoot(a.db)
		}
	}

	a.updateAccountObject(obj)
}

func (a *AccountDB) CommitObject(account string, deleteEmptyObjects bool) {
	obj, exist := a.accountObjects[account]
	if !exist {
		return
	}

	if obj.IsContract() {
		if obj.contract.dirtyCode {
			logger.Trace("insert code", "hash", types.BytesToHash(obj.CodeHash()))

			a.db.TreeDB().InsertBlob(types.BytesToHash(obj.CodeHash()), obj.contract.code)
			obj.contract.dirtyCode = false
		}

		if !obj.contract.deleted {
			obj.CommitTree(a.db)
		}
	}

	a.updateAccountObject(obj)
}

func (a *AccountDB) IntermediateRoot() types.Hash {
	return a.trie.Hash()
}

func (a *AccountDB) CommitTree() (types.Hash, error) {
	root, err := a.trie.Commit(func(leaf []byte, parent types.Hash) error {
		var account accounts.AccountStore
		if err := rlp.DecodeBytes(leaf, &account); err != nil {
			return nil
		}
		if account.IsContract() {
			if account.StorageRoot() != types.EmptyRootHash {
				a.db.TreeDB().Reference(account.StorageRoot(), parent)
			}
			code := types.BytesToHash(account.CodeHash())
			if code != types.EmptyCodeHash {
				a.db.TreeDB().Reference(code, parent)
			}
		}

		return nil
	})

	if err != nil {
		return types.Hash{}, err
	}

	err = a.db.TreeDB().Commit(root, false)
	if err != nil {
		return types.Hash{}, err
	}

	return root, err
}
