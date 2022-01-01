// Package basedb
//
// @author: xwc1125
package basedb

type JournalEntry interface {
	Revert()
	Dirtied() string
}

type JournalCommitter interface {
	FinaliseObject(string, bool)
	CommitObject(string, bool)
}

type Journal struct {
	entries   []JournalEntry // Current changes tracked by the Journal
	Dirties   map[string]int // Dirty accounts and the number of changes
	Pending   map[string]struct{}
	Committer map[string]JournalCommitter
}

// NewJournal create a new initialized Journal.
func NewJournal() *Journal {
	return &Journal{
		Dirties:   make(map[string]int),
		Pending:   make(map[string]struct{}),
		Committer: make(map[string]JournalCommitter),
	}
}

// Append inserts a new modification entry to the end of the change journal.
func (j *Journal) Append(entry JournalEntry, commiter JournalCommitter) {
	j.entries = append(j.entries, entry)
	if key := entry.Dirtied(); key != "" {
		j.Dirties[key]++
		j.Pending[key] = struct{}{}
		j.Committer[key] = commiter
	}
}

// Revert undoes a batch of journalled modifications along with any reverted
// dirty handling too.
func (j *Journal) Revert(snapshot int) {
	for i := len(j.entries) - 1; i >= snapshot; i-- {
		// Undo the changes made by the operation
		j.entries[i].Revert()

		// Drop any dirty tracking induced by the change
		if key := j.entries[i].Dirtied(); key != "" {
			if j.Dirties[key]--; j.Dirties[key] == 0 {
				delete(j.Dirties, key)
				delete(j.Pending, key)
			}
		}
	}
	j.entries = j.entries[:snapshot]
}

// Dirty explicitly sets an address to dirty, even if the change entries would
// otherwise suggest it as clean. This method is an ugly hack to handle the RIPEMD
// precompile consensus exception.
func (j *Journal) Dirty(key string) {
	j.Dirties[key]++
}

// Length returns the current number of entries in the journal.
func (j *Journal) Length() int {
	return len(j.entries)
}

func (j *Journal) ClearDirties() {
	j.entries = []JournalEntry{}
	j.Dirties = make(map[string]int)
}
