// Package statedb
//
// @author: xwc1125
package statedb

import (
	"fmt"
	"sort"

	"github.com/chain5j/chain5j-pkg/codec/rlp"
	"github.com/chain5j/chain5j-pkg/crypto/hashalg/sha3"
	"github.com/chain5j/chain5j-pkg/database/kvstore"
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/models/accounts"
	"github.com/chain5j/chain5j-protocol/models/statetype"
	"github.com/chain5j/chain5j-protocol/pkg/database/basedb"
	"github.com/chain5j/chain5j-protocol/pkg/database/statedb/accountdb"
	"github.com/chain5j/logger"
)

var (
	rootPrefix   = []byte("state-root-")
	domainPrefix = "domains-"
)

type stateRoots struct {
	AccountRoot types.Hash
	MapRoot     types.Hash
	KVSRoot     types.Hash
	XRoot       types.Hash
}

func (roots *stateRoots) Hash() types.Hash {
	var rootBytes []byte
	rootBytes = append(roots.AccountRoot.Bytes(), roots.MapRoot.Bytes()...)
	rootBytes = append(rootBytes, roots.KVSRoot.Bytes()...)
	rootBytes = append(rootBytes, roots.XRoot.Bytes()...)
	return types.BytesToHash(sha3.Keccak256(rootBytes))
}

type revision struct {
	id           int
	journalIndex int
}

type StateDB struct {
	log    logger.Logger
	config *Config
	db     kvstore.Database

	*accountdb.AccountDB
	*accountdb.AccountMap
	*accountdb.AccountKVS

	xstore *xStore

	thash, bhash types.Hash
	txIndex      int
	logs         map[types.Hash][]*statetype.Log
	logSize      uint

	// Journal of state modifications. This is the backbone of
	// Snapshot and RevertToSnapshot.
	journal        *basedb.Journal
	validRevisions []revision
	nextRevisionId int

	// The refund counter, also used by state transitioning.
	refund uint64

	preimages map[types.Hash][]byte

	err error
}

// New Create a new state from a given trie.
func New(root types.Hash, db kvstore.Database, opts ...option) (*StateDB, error) {
	s := &StateDB{
		log: logger.New("state_db"),
		config: &Config{
			Metrics:      false,
			MetricsLevel: 0,
		},
		db:        db,
		preimages: make(map[types.Hash][]byte),
		logs:      make(map[types.Hash][]*statetype.Log),
	}
	if err := apply(s, opts...); err != nil {
		return nil, err
	}
	var (
		roots stateRoots
		err   error
	)
	if root != (types.Hash{}) && root != types.EmptyRootHash {
		rootBytes, err := db.Get(rootKey(root))
		if err != nil {
			return nil, err
		}

		if err = rlp.DecodeBytes(rootBytes, &roots); err != nil {
			return nil, err
		}
	}

	if s.config.IsMetrics(1) {
		s.log.Trace("new state db", "account_root", roots.AccountRoot, "map_root", roots.MapRoot, "kvs_root", roots.KVSRoot, "x_root", roots.XRoot, "state_roots", root)
	}

	s.journal = basedb.NewJournal()

	s.AccountDB, err = accountdb.NewAccountDB(roots.AccountRoot, basedb.NewDatabase(db), s.journal)
	if err != nil {
		s.log.Error("NewAccountDB Err", "err", err)
		return nil, err
	}

	s.AccountMap, err = accountdb.NewAccountMap(roots.MapRoot, basedb.NewDatabase(db), s.journal)
	if err != nil {
		s.log.Error("NewAccountMap Err", "err", err)
		return nil, err
	}

	s.AccountKVS, err = accountdb.NewAccountKVS(roots.KVSRoot, basedb.NewDatabase(db), s.journal)
	if err != nil {
		s.log.Error("NewAccountMap Err", "err", err)
		return nil, err
	}

	s.xstore, err = NewXStore(roots.XRoot, basedb.NewDatabase(db))
	if err != nil {
		s.log.Error("New XStore Err", "err", err)
		return nil, err
	}
	return s, nil
}

func (s *StateDB) SetConfig(config *Config) {
	if s.config != nil {
		s.config = config
	}
}

// setError remembers the first non-nil error it is called with.
func (s *StateDB) setError(err error) {
	if s.err == nil {
		s.err = err
	}
}

func (s *StateDB) Error() error {
	return s.err
}

func (s *StateDB) Prepare(thash, bhash types.Hash, ti int) {
	s.thash = thash
	s.bhash = bhash
	s.txIndex = ti
}

func (s *StateDB) CreateAccount(account *accounts.AccountStore) {
	// 避免外部修改account的值，影响db.
	store := account.Copy()

	s.AccountDB.CreateAccount(store)
	for address := range store.Addresses {
		s.CreateMap(address, store.AccountName())
	}
}

func (s *StateDB) SetAddress(account string, address types.Address) {
	s.AccountDB.SetAddress(account, address)

	s.CreateMap(address, account)
}

func (s *StateDB) Finalise(deleteEmptyObjects bool) {
	for key := range s.journal.Dirties {
		committer := s.journal.Committer[key]
		if committer != nil {
			committer.FinaliseObject(key, deleteEmptyObjects)
		}
	}

	s.clearJournal()
}

func (s *StateDB) IntermediateRoot(deleteEmptyObjects bool) types.Hash {
	s.Finalise(deleteEmptyObjects)

	accountRoot := s.AccountDB.IntermediateRoot()
	mapRoot := s.AccountMap.IntermediateRoot()
	kvsRoot := s.AccountKVS.IntermediateRoot()
	xRoot := s.xstore.IntermediateRoot()
	roots := stateRoots{
		AccountRoot: accountRoot,
		MapRoot:     mapRoot,
		KVSRoot:     kvsRoot,
		XRoot:       xRoot,
	}
	if s.config.IsMetrics(1) {
		logger.Debug("intermediate root", "account_root", accountRoot, "map_root", mapRoot, "kvs_root", kvsRoot, "x_root", xRoot, "state_roots", roots.Hash())
	}
	return roots.Hash()
}

func (s *StateDB) Commit(deleteEmptyObjects bool) (types.Hash, error) {
	var err error
	defer s.setError(err)

	for key := range s.journal.Pending {
		committer := s.journal.Committer[key]
		if committer != nil {
			committer.CommitObject(key, deleteEmptyObjects)
		}
	}

	accountRoot, err := s.AccountDB.CommitTree()
	if err != nil {
		return types.Hash{}, err
	}

	mapRoot, err := s.AccountMap.CommitTree()
	if err != nil {
		return types.Hash{}, err
	}

	kvsRoot, err := s.AccountKVS.CommitTree()
	if err != nil {
		return types.Hash{}, err
	}

	xRoot, err := s.xstore.CommitTree()
	if err != nil {
		return types.Hash{}, err
	}

	roots := &stateRoots{
		AccountRoot: accountRoot,
		MapRoot:     mapRoot,
		KVSRoot:     kvsRoot,
		XRoot:       xRoot,
	}

	rlpBytes, _ := rlp.EncodeToBytes(roots)
	root := roots.Hash()

	return root, s.db.Put(rootKey(root), rlpBytes)
}

func rootKey(hash types.Hash) []byte {
	return append(rootPrefix, hash.Bytes()...)
}

// Snapshot returns an identifier for the current revision of the state.
func (s *StateDB) Snapshot() int {
	id := s.nextRevisionId
	s.nextRevisionId++
	s.validRevisions = append(s.validRevisions, revision{id, s.journal.Length()})
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
	s.journal.Revert(snapshot)
	s.validRevisions = s.validRevisions[:idx]
}

func (s *StateDB) clearJournal() {
	s.journal.ClearDirties()
}

func (s *StateDB) AddLog(log *statetype.Log) {
	s.journal.Append(addLogChange{db: s, txhash: s.thash}, nil)

	log.TransactionHash = s.thash
	log.BlockHash = s.bhash
	log.TxIndex = uint(s.txIndex)
	log.Index = s.logSize
	s.logs[s.thash] = append(s.logs[s.thash], log)
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

// AddRefund adds gas to the refund counter
func (s *StateDB) AddRefund(gas uint64) {
	s.journal.Append(refundChange{db: s, prev: s.refund}, nil)
	s.refund += gas
}

// SubRefund removes gas from the refund counter.
// This method will panic if the refund counter goes below zero
func (s *StateDB) SubRefund(gas uint64) {
	s.journal.Append(refundChange{db: s, prev: s.refund}, nil)
	if gas > s.refund {
		panic("Refund counter below zero")
	}
	s.refund -= gas
}

func (s *StateDB) GetRefund() uint64 {
	return s.refund
}

func (s *StateDB) AddPreimage(hash types.Hash, preimage []byte) {
	if _, ok := s.preimages[hash]; !ok {
		s.journal.Append(addPreimageChange{db: s, hash: hash}, nil)
		pi := make([]byte, len(preimage))
		copy(pi, preimage)
		s.preimages[hash] = pi
	}
}

func (s *StateDB) GetDomain(domain string) *accounts.DomainStore {
	domainBytes, err := s.xstore.Get(domainKey(domain))
	if err != nil {
		return nil
	}

	var store accounts.DomainStore
	if err = rlp.DecodeBytes(domainBytes, &store); err != nil {
		return nil
	}

	return &store
}

func (s *StateDB) AddDomain(domain string, store accounts.DomainStore) {
	storeBytes, _ := rlp.EncodeToBytes(&store)

	s.xstore.Store(domainKey(domain), storeBytes)
}

func domainKey(domain string) string {
	return domainPrefix + domain
}

func (s *StateDB) KVExist(account string, namespace string, key string) bool {
	return s.AccountKVS.KeyExist(account, namespace, key)
}
