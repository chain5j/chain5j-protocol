// Package vmdb
//
// @author: xwc1125
package vmdb

import (
	"errors"
	"fmt"
	"github.com/chain5j/chain5j-pkg/codec/rlp"
	"github.com/chain5j/chain5j-pkg/collection/trees/tree"
	"github.com/chain5j/chain5j-pkg/crypto/hashalg/sha3"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/models/accounts"
	"github.com/chain5j/chain5j-protocol/models/statetype"
	"github.com/chain5j/chain5j-protocol/pkg/database/basedb"
	"github.com/chain5j/chain5j-protocol/pkg/database/statedb/accountdb"
	"github.com/chain5j/chain5j-protocol/pkg/database/statedb/vmdb/model"
	"github.com/chain5j/logger"
	"math/big"
	"sort"
	"time"
)

type revision struct {
	id           int
	journalIndex int
}

type proofList [][]byte

func (n *proofList) Put(key []byte, value []byte) error {
	*n = append(*n, value)
	return nil
}

func (n *proofList) Delete(key []byte) error {
	panic("not supported")
}

// StateDB within the ethereum protocol are used to store anything
// within the merkle trie. StateDBs take care of caching and storing
// nested states. It's the general query interface to retrieve:
// * Contracts
// * Accounts
type StateDB struct {
	accountDB *accountdb.AccountDB

	db   basedb.Database
	trie basedb.Tree

	// This map holds 'live' objects, which will get modified while processing a state transition.
	stateObjects        map[types.DomainAddress]*stateObject
	stateObjectsPending map[types.DomainAddress]struct{} // State objects finalized but not yet written to the trie
	stateObjectsDirty   map[types.DomainAddress]struct{} // State objects modified in the current execution

	// DB error.
	// State objects are used by the consensus core and VM which are
	// unable to deal with database-level errors. Any error that occurs
	// during a database read is memoized here and will eventually be returned
	// by StateDB.Commit.
	dbErr error

	// The refund counter, also used by state transitioning.
	refund uint64

	thash, bhash types.Hash
	txIndex      int
	logs         map[types.Hash][]*statetype.Log
	logSize      uint

	preimages map[types.Hash][]byte

	// Journal of state modifications. This is the backbone of
	// Snapshot and RevertToSnapshot.
	journal        *journal
	validRevisions []revision
	nextRevisionId int

	// Measurements gathered during execution for debugging purposes
	AccountReads   time.Duration
	AccountHashes  time.Duration
	AccountUpdates time.Duration
	AccountCommits time.Duration
	StorageReads   time.Duration
	StorageHashes  time.Duration
	StorageUpdates time.Duration
	StorageCommits time.Duration

	contractInfos map[types.DomainAddress][]byte
}

// Create a new state from a given trie.
func New(root types.Hash, db basedb.Database, accountDB *accountdb.AccountDB) (*StateDB, error) {
	tr, err := db.OpenTree(root)
	if err != nil {
		return nil, err
	}
	return &StateDB{
		accountDB: accountDB,

		db:                  db,
		trie:                tr,
		stateObjects:        make(map[types.DomainAddress]*stateObject),
		stateObjectsPending: make(map[types.DomainAddress]struct{}),
		stateObjectsDirty:   make(map[types.DomainAddress]struct{}),
		logs:                make(map[types.Hash][]*statetype.Log),
		preimages:           make(map[types.Hash][]byte),
		journal:             newJournal(),
		contractInfos:       make(map[types.DomainAddress][]byte),
	}, nil
}

// setError remembers the first non-nil error it is called with.
func (s *StateDB) setError(err error) {
	if s.dbErr == nil {
		s.dbErr = err
	}
}

func (s *StateDB) Error() error {
	return s.dbErr
}

// Reset clears out all ephemeral state objects from the state db, but keeps
// the underlying state trie to avoid reloading data for the next operations.
func (s *StateDB) Reset(root types.Hash) error {
	tr, err := s.db.OpenTree(root)
	if err != nil {
		return err
	}
	s.trie = tr
	s.stateObjects = make(map[types.DomainAddress]*stateObject)
	s.stateObjectsPending = make(map[types.DomainAddress]struct{})
	s.stateObjectsDirty = make(map[types.DomainAddress]struct{})
	s.thash = types.Hash{}
	s.bhash = types.Hash{}
	s.txIndex = 0
	s.logs = make(map[types.Hash][]*statetype.Log)
	s.logSize = 0
	s.preimages = make(map[types.Hash][]byte)

	s.contractInfos = make(map[types.DomainAddress][]byte)
	s.clearJournalAndRefund()
	return nil
}

func (s *StateDB) AddLog(logger *statetype.Log) {
	s.journal.append(addLogChange{txhash: s.thash})

	logger.TransactionHash = s.thash
	logger.BlockHash = s.bhash
	logger.TxIndex = uint(s.txIndex)
	logger.Index = s.logSize
	s.logs[s.thash] = append(s.logs[s.thash], logger)
	s.logSize++
}

func (s *StateDB) GetLogs(hash types.Hash) []*statetype.Log {
	return s.logs[hash]
}

func (s *StateDB) Logs() []*statetype.Log {
	var logs []*statetype.Log
	for _, lgs := range s.logs {
		logs = append(logs, lgs...)
	}
	return logs
}

// AddPreimage records a SHA3 preimage seen by the VM.
func (s *StateDB) AddPreimage(hash types.Hash, preimage []byte) {
	if _, ok := s.preimages[hash]; !ok {
		s.journal.append(addPreimageChange{hash: hash})
		pi := make([]byte, len(preimage))
		copy(pi, preimage)
		s.preimages[hash] = pi
	}
}

// Preimages returns a list of SHA3 preimages that have been submitted.
func (s *StateDB) Preimages() map[types.Hash][]byte {
	return s.preimages
}

// AddRefund adds gas to the refund counter
func (s *StateDB) AddRefund(gas uint64) {
	s.journal.append(refundChange{prev: s.refund})
	s.refund += gas
}

// SubRefund removes gas from the refund counter.
// This method will panic if the refund counter goes below zero
func (s *StateDB) SubRefund(gas uint64) {
	s.journal.append(refundChange{prev: s.refund})
	if gas > s.refund {
		panic("Refund counter below zero")
	}
	s.refund -= gas
}

// Exist reports whether the given account address exists in the state.
// Notably this also returns true for suicided accounts.
func (s *StateDB) Exist(addr types.DomainAddress) bool {
	return s.getStateObject(addr) != nil
}

// Empty returns whether the state object is either non-existent
// or empty according to the EIP161 specification (balance = nonce = code = 0)
func (s *StateDB) Empty(addr types.DomainAddress) bool {
	so := s.getStateObject(addr)
	return so == nil || so.empty()
}

// Retrieve the balance from the given address or 0 if object not found
func (s *StateDB) GetBalance(addr types.DomainAddress) *big.Int {
	return s.accountDB.GetBalance(addr.DomainAddr)
}

func (s *StateDB) GetNonce(addr types.DomainAddress) uint64 {
	return s.accountDB.GetNonce(addr.DomainAddr)
}

// TxIndex returns the current transaction index set by Prepare.
func (s *StateDB) TxIndex() int {
	return s.txIndex
}

// BlockHash returns the current block hash set by Prepare.
func (s *StateDB) BlockHash() types.Hash {
	return s.bhash
}

func (s *StateDB) GetCode(addr types.DomainAddress) []byte {
	stateObject := s.getStateObject(addr)
	if stateObject != nil {
		return stateObject.Code(s.db)
	}
	return nil
}

func (s *StateDB) GetCodeSize(addr types.DomainAddress) int {
	stateObject := s.getStateObject(addr)
	if stateObject == nil {
		return 0
	}
	if stateObject.code != nil {
		return len(stateObject.code)
	}
	size, err := s.db.ContractCodeSize(stateObject.addrHash, types.BytesToHash(stateObject.CodeHash()))
	if err != nil {
		s.setError(err)
	}
	return size
}

func (s *StateDB) GetCodeHash(addr types.DomainAddress) types.Hash {
	stateObject := s.getStateObject(addr)
	if stateObject == nil {
		return types.Hash{}
	}
	return types.BytesToHash(stateObject.CodeHash())
}

// GetState retrieves a value from the given account's storage trie.
func (s *StateDB) GetState(addr types.DomainAddress, hash types.Hash) types.Hash {
	stateObject := s.getStateObject(addr)
	if stateObject != nil {
		return stateObject.GetState(s.db, hash)
	}
	return types.Hash{}
}

// GetProof returns the MerkleProof for a given Account
func (s *StateDB) GetProof(a types.DomainAddress) ([][]byte, error) {
	var proof proofList
	err := s.trie.Prove(sha3.Keccak256(a.Bytes()), 0, &proof)
	return [][]byte(proof), err
}

// GetProof returns the StorageProof for given key
func (s *StateDB) GetStorageProof(a types.DomainAddress, key types.Hash) ([][]byte, error) {
	var proof proofList
	trie := s.StorageTree(a)
	if trie == nil {
		return proof, errors.New("storage trie for requested address does not exist")
	}
	err := trie.Prove(sha3.Keccak256(key.Bytes()), 0, &proof)
	return [][]byte(proof), err
}

// GetCommittedState retrieves a value from the given account's committed storage trie.
func (s *StateDB) GetCommittedState(addr types.DomainAddress, hash types.Hash) types.Hash {
	stateObject := s.getStateObject(addr)
	if stateObject != nil {
		return stateObject.GetCommittedState(s.db, hash)
	}
	return types.Hash{}
}

// Database retrieves the low level database supporting the lower level trie ops.
func (s *StateDB) Database() basedb.Database {
	return s.db
}

// StorageTrie returns the storage trie of an account.
// The return value is a copy and is nil for non-existent accounts.
func (s *StateDB) StorageTree(addr types.DomainAddress) tree.Tree {
	stateObject := s.getStateObject(addr)
	if stateObject == nil {
		return nil
	}
	cpy := stateObject.deepCopy(s)
	return cpy.updateTree(s.db)
}

func (s *StateDB) HasSuicided(addr types.DomainAddress) bool {
	stateObject := s.getStateObject(addr)
	if stateObject != nil {
		return stateObject.suicided
	}
	return false
}

/*
 * SETTERS
 */

// AddBalance adds amount to the account associated with addr.
func (s *StateDB) AddBalance(addr types.DomainAddress, amount *big.Int) {
	s.accountDB.AddBalance(addr.DomainAddr, amount)
}

// SubBalance subtracts amount from the account associated with addr.
func (s *StateDB) SubBalance(addr types.DomainAddress, amount *big.Int) {
	s.accountDB.SubBalance(addr.DomainAddr, amount)
}

func (s *StateDB) SetBalance(addr types.DomainAddress, amount *big.Int) {
	s.accountDB.SetBalance(addr.DomainAddr, amount)
}

func (s *StateDB) SetNonce(addr types.DomainAddress, nonce uint64) {
	s.accountDB.SetNonce(addr.DomainAddr, nonce)
}

func (s *StateDB) SetCode(addr types.DomainAddress, code []byte) {
	stateObject := s.GetOrNewStateObject(addr)
	if stateObject != nil {
		stateObject.SetCode(types.BytesToHash(sha3.Keccak256(code)), code)
	}
}

func (s *StateDB) SetState(addr types.DomainAddress, key, value types.Hash) {
	stateObject := s.GetOrNewStateObject(addr)
	if stateObject != nil {
		stateObject.SetState(s.db, key, value)
	}
}

// SetStorage replaces the entire storage for the specified account with given
// storage. This function should only be used for debugging.
//func (s *StateDB) SetStorage(addr types.Address, storage map[types.Hash]types.Hash) {
//	stateObject := s.GetOrNewAccountObject(addr)
//	if stateObject != nil {
//		stateObject.SetStorage(storage)
//	}
//}

// Suicide marks the given account as suicided.
// This clears the account balance.
//
// The account's state object is still available until the state is committed,
// getStateObject will return a non-nil account after Suicide.
func (s *StateDB) Suicide(addr types.DomainAddress) bool {
	stateObject := s.getStateObject(addr)
	if stateObject == nil {
		return false
	}
	s.journal.append(suicideChange{
		account:     &addr,
		prev:        stateObject.suicided,
		prevbalance: new(big.Int).Set(stateObject.Balance()),
	})
	stateObject.markSuicided()
	stateObject.data.Balance = new(big.Int)

	return true
}

//
// Setting, updating & deleting state object methods.
//

// updateStateObject writes the given object to the trie.
func (s *StateDB) updateStateObject(obj *stateObject) {
	// Encode the account and update the account trie
	addr := obj.Address()

	data, err := rlp.EncodeToBytes(obj)
	if err != nil {
		panic(fmt.Errorf("can't encode object at %x: %v", addr.DomainAddr, err))
	}
	s.setError(s.trie.TryUpdate(addr.Bytes(), data))
}

// deleteStateObject removes the given object from the state trie.
func (s *StateDB) deleteStateObject(obj *stateObject) {
	// Delete the account from the trie
	addr := obj.Address()
	s.setError(s.trie.TryDelete(addr.Bytes()))
}

// getStateObject retrieves a state object given by the address, returning nil if
// the object is not found or was deleted in this execution context. If you need
// to differentiate between non-existent/just-deleted, use getDeletedStateObject.
func (s *StateDB) getStateObject(addr types.DomainAddress) *stateObject {
	if obj := s.getDeletedStateObject(addr); obj != nil && !obj.deleted {
		return obj
	}
	return nil
}

// getDeletedStateObject is similar to getStateObject, but instead of returning
// nil for a deleted state object, it returns the actual object with the deleted
// flag set. This is needed by the state journal to revert to the correct s-
// destructed object instead of wiping all knowledge about the state object.
func (s *StateDB) getDeletedStateObject(addr types.DomainAddress) *stateObject {
	// Prefer live objects if any is available
	if obj := s.stateObjects[addr]; obj != nil {
		return obj
	}
	// Load the object from the database
	enc, err := s.trie.TryGet(addr.Bytes())
	if len(enc) == 0 {
		s.setError(err)
		return nil
	}
	//var data model.Account
	data := model.NewAccount(types.EmptyDomainAddress)
	if err := rlp.DecodeBytes(enc, &data); err != nil {
		logger.Error("Failed to decode state object", "addr", addr, "err", err)
		return nil
	}
	// Insert into the live set
	obj := newObject(s, addr, *data)
	s.setStateObject(obj)
	return obj
}

func (s *StateDB) setStateObject(object *stateObject) {
	s.stateObjects[object.Address()] = object
}

// Retrieve a state object or create a new state object if nil.
func (s *StateDB) GetOrNewStateObject(addr types.DomainAddress) *stateObject {
	stateObject := s.getStateObject(addr)
	if stateObject == nil {
		stateObject, _ = s.createObject(addr)
	}
	return stateObject
}

// createObject creates a new state object. If there is an existing account with
// the given address, it is overwritten and returned as the second return value.
func (s *StateDB) createObject(addr types.DomainAddress) (newobj, prev *stateObject) {
	prev = s.getDeletedStateObject(addr) // Note, prev might have been deleted, we need that!

	account := model.NewAccount(addr)
	newobj = newObject(s, addr, *account)
	newobj.setNonce(0) // sets the object to dirty
	if prev == nil {
		s.journal.append(createObjectChange{account: &addr})
	} else {
		s.journal.append(resetObjectChange{prev: prev})
	}
	s.setStateObject(newobj)
	return newobj, prev
}

// CreateAccount explicitly creates a state object. If a state object with the address
// already exists the balance is carried over to the new account.
//
// CreateAccount is called during the EVM CREATE operation. The situation might arise that
// a contract does the following:
//
//   1. sends funds to sha(account ++ (nonce + 1))
//   2. tx_create(sha(account ++ nonce)) (note that this gets the address of 1)
//
// Carrying over the balance ensures that Ether doesn't disappear.
func (s *StateDB) CreateAccount(addr types.DomainAddress) {
	s.accountDB.CreateAccount(&accounts.AccountStore{
		CN:     addr.Addr.Hex(),
		Domain: addr.DomainAddr,
	})
	//newObj, prev := s.createObject(addr)
	//if prev != nil {
	//	newObj.setBalance(prev.data.Balance)
	//}
}

func (db *StateDB) ForEachStorage(addr types.DomainAddress, cb func(key, value types.Hash) bool) error {
	so := db.getStateObject(addr)
	if so == nil {
		return nil
	}
	it := tree.NewIterator(so.getTree(db.db).NodeIterator(nil))

	for it.Next() {
		key := types.BytesToHash(db.trie.GetKey(it.Key))
		if value, dirty := so.dirtyStorage[key]; dirty {
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

// Copy creates a deep, independent copy of the state.
// Snapshots of the copied state cannot be applied to the copy.
func (s *StateDB) Copy() *StateDB {
	// Copy all the basic fields, initialize the memory ones
	state := &StateDB{
		db:                  s.db,
		trie:                s.db.CopyTree(s.trie),
		stateObjects:        make(map[types.DomainAddress]*stateObject),
		stateObjectsPending: make(map[types.DomainAddress]struct{}, len(s.stateObjectsPending)),
		stateObjectsDirty:   make(map[types.DomainAddress]struct{}, len(s.journal.dirties)),
		refund:              s.refund,
		logs:                make(map[types.Hash][]*statetype.Log, len(s.logs)),
		logSize:             s.logSize,
		preimages:           make(map[types.Hash][]byte, len(s.preimages)),
		journal:             newJournal(),
	}
	// Copy the dirty states, logs, and preimages
	for addr := range s.journal.dirties {
		// and in the Finalise-method, there is a case where an object is in the journal but not
		// in the stateObjects: OOG after touch on ripeMD prior to Byzantium. Thus, we need to check for
		// nil
		if object, exist := s.stateObjects[addr]; exist {
			// Even though the original object is dirty, we are not copying the journal,
			// so we need to make sure that anyside effect the journal would have caused
			// during a commit (or similar op) is already applied to the copy.
			state.stateObjects[addr] = object.deepCopy(state)

			state.stateObjectsDirty[addr] = struct{}{}   // Mark the copy dirty to force internal (code/state) commits
			state.stateObjectsPending[addr] = struct{}{} // Mark the copy pending to force external (account) commits
		}
	}
	// Above, we don't copy the actual journal. This means that if the copy is copied, the
	// loop above will be a no-op, since the copy's journal is empty.
	// Thus, here we iterate over stateObjects, to enable copies of copies
	for addr := range s.stateObjectsPending {
		if _, exist := state.stateObjects[addr]; !exist {
			state.stateObjects[addr] = s.stateObjects[addr].deepCopy(state)
		}
		state.stateObjectsPending[addr] = struct{}{}
	}
	for addr := range s.stateObjectsDirty {
		if _, exist := state.stateObjects[addr]; !exist {
			state.stateObjects[addr] = s.stateObjects[addr].deepCopy(state)
		}
		state.stateObjectsDirty[addr] = struct{}{}
	}
	for hash, logs := range s.logs {
		cpy := make([]*statetype.Log, len(logs))
		for i, l := range logs {
			cpy[i] = new(statetype.Log)
			*cpy[i] = *l
		}
		state.logs[hash] = cpy
	}
	for hash, preimage := range s.preimages {
		state.preimages[hash] = preimage
	}
	return state
}

// Snapshot returns an identifier for the current revision of the state.
func (s *StateDB) Snapshot() int {
	id := s.nextRevisionId
	s.nextRevisionId++
	s.validRevisions = append(s.validRevisions, revision{id, s.journal.length()})
	return id
}

// RevertToSnapshot reverts all state changes made since the given revision.
func (s *StateDB) RevertToSnapshot(revid int) {
	// Find the snapshot in the stack of valid snapshots.
	idx := sort.Search(len(s.validRevisions), func(i int) bool {
		return s.validRevisions[i].id >= revid
	})
	if idx == len(s.validRevisions) || s.validRevisions[idx].id != revid {
		panic(fmt.Errorf("revision id %v cannot be reverted", revid))
	}
	snapshot := s.validRevisions[idx].journalIndex

	// Replay the journal to undo changes and remove invalidated snapshots
	s.journal.revert(s, snapshot)
	s.validRevisions = s.validRevisions[:idx]
}

// GetRefund returns the current value of the refund counter.
func (s *StateDB) GetRefund() uint64 {
	return s.refund
}

// Finalise finalises the state by removing the s destructed objects and clears
// the journal as well as the refunds. Finalise, however, will not push any updates
// into the tries just yet. Only IntermediateRoot or Commit will do that.
func (s *StateDB) Finalise(deleteEmptyObjects bool) {
	for addr := range s.journal.dirties {
		obj, exist := s.stateObjects[addr]
		if !exist {
			// ripeMD is 'touched' at block 1714175, in tx 0x1237f737031e40bcde4a8b7e717b2d15e3ecadfe49bb1bbc71ee9deb09c6fcf2
			// That tx goes out of gas, and although the notion of 'touched' does not exist there, the
			// touch-event will still be recorded in the journal. Since ripeMD is a special snowflake,
			// it will persist in the journal even though the journal is reverted. In this special circumstance,
			// it may exist in `s.journal.dirties` but not in `s.stateObjects`.
			// Thus, we can safely ignore it here
			continue
		}
		if obj.suicided || (deleteEmptyObjects && obj.empty()) {
			obj.deleted = true
		} else {
			obj.finalise()
		}
		s.stateObjectsPending[addr] = struct{}{}
		s.stateObjectsDirty[addr] = struct{}{}
	}
	// Invalidate journal because reverting across transactions is not allowed.
	s.clearJournalAndRefund()
}

// IntermediateRoot computes the current root hash of the state trie.
// It is called in between transactions to get the root hash that
// goes into transaction receipts.
func (s *StateDB) IntermediateRoot(deleteEmptyObjects bool) types.Hash {
	// Finalise all the dirty storage states and write them into the tries
	s.Finalise(deleteEmptyObjects)

	for addr := range s.stateObjectsPending {
		obj := s.stateObjects[addr]

		if obj.deleted {
			s.deleteStateObject(obj)
		} else {
			obj.updateRoot(s.db)
			s.updateStateObject(obj)
		}
	}
	if len(s.stateObjectsPending) > 0 {
		s.stateObjectsPending = make(map[types.DomainAddress]struct{})
	}
	// Track the amount of time wasted on hashing the account trie
	return s.trie.Hash()
}

// Prepare sets the current transaction hash and index and block hash which is
// used when the EVM emits new state logs.
func (s *StateDB) Prepare(thash, bhash types.Hash, ti int) {
	s.thash = thash
	s.bhash = bhash
	s.txIndex = ti
}

func (s *StateDB) clearJournalAndRefund() {
	if len(s.journal.entries) > 0 {
		s.journal = newJournal()
		s.refund = 0
	}
	s.validRevisions = s.validRevisions[:0] // Snapshots can be created without journal entires
}

// Commit writes the state to the underlying in-memory trie database.
func (s *StateDB) Commit(deleteEmptyObjects bool) (types.Hash, error) {
	// Finalize any pending changes and merge everything into the tries
	s.IntermediateRoot(deleteEmptyObjects)

	// Commit objects to the trie, measuring the elapsed time
	for addr := range s.stateObjectsDirty {
		if obj := s.stateObjects[addr]; !obj.deleted {
			// Write any contract code associated with the state object
			if obj.code != nil && obj.dirtyCode {
				s.db.TreeDB().InsertBlob(types.BytesToHash(obj.CodeHash()), obj.code)
				obj.dirtyCode = false
			}
			// Write any storage changes in the state object to its storage trie
			if err := obj.CommitTree(s.db); err != nil {
				return types.Hash{}, err
			}
		}
	}
	if len(s.stateObjectsDirty) > 0 {
		s.stateObjectsDirty = make(map[types.DomainAddress]struct{})
	}
	// Write the account trie changes, measuing the amount of wasted time
	return s.trie.Commit(func(leaf []byte, parent types.Hash) error {
		//var account model.Account
		account := model.NewAccount(types.EmptyDomainAddress)
		if err := rlp.DecodeBytes(leaf, account); err != nil {
			return nil
		}
		if account.Root != types.EmptyRootHash {
			s.db.TreeDB().Reference(account.Root, parent)
		}
		code := types.BytesToHash(account.CodeHash)
		if code != types.EmptyCodeHash {
			s.db.TreeDB().Reference(code, parent)
		}
		return nil
	})
}

func (s *StateDB) SubTokenBalance(addr types.DomainAddress, token types.DomainAddress, amount *big.Int) {
	stateObject := s.GetOrNewStateObject(addr)
	if stateObject != nil {
		stateObject.SubTokenBalance(token, amount)
	}
}

func (s *StateDB) SetTokenBalance(addr types.DomainAddress, token types.DomainAddress, amount *big.Int) {
	stateObject := s.GetOrNewStateObject(addr)
	if stateObject != nil {
		stateObject.SetTokenBalance(token, amount)
	}
}

func (s *StateDB) AddTokenBalance(addr types.DomainAddress, token types.DomainAddress, amount *big.Int) {
	stateObject := s.GetOrNewStateObject(addr)
	if stateObject != nil {
		stateObject.AddTokenBalance(token, amount)
	}
}
func (s *StateDB) GetTokenBalance(addr types.DomainAddress, token types.DomainAddress) *big.Int {
	stateObject := s.getStateObject(addr)
	if stateObject != nil {
		return stateObject.TokenBalance(token)
	}
	return big.NewInt(0)
}
func (s *StateDB) GetTokenBalances(addr types.DomainAddress) statetype.TokenValues {
	stateObject := s.getStateObject(addr)
	if stateObject != nil {
		return stateObject.TokenBalances()
	}
	return nil
}

func (s *StateDB) GetContractInfo(addr types.DomainAddress) []byte {
	return s.contractInfos[addr]
}
func (s *StateDB) SetContractInfo(addr types.DomainAddress, info []byte) {
	stateObject := s.GetOrNewStateObject(addr)
	if stateObject != nil {
		stateObject.SetContractInfo(types.BytesToHash(sha3.Keccak256(info)), info)
	}
	s.contractInfos[addr] = info
}
