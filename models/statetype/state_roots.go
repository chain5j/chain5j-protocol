// Package statetype
//
// @author: xwc1125
package statetype

import (
	"encoding/json"
	"github.com/chain5j/chain5j-pkg/codec/rlp"
	"github.com/chain5j/chain5j-pkg/types"
	"io"
	"reflect"
	"sort"
	"sync"
)

// StateRoots 并发map 多协程，使用安全
type StateRoots struct {
	m    map[types.TxType]types.Hash
	lock *sync.Mutex
}

func (r StateRoots) String() string {
	bytes, _ := json.Marshal(r.m)
	return string(bytes)
}

func NewRoots() *StateRoots {
	return &StateRoots{
		m:    map[types.TxType]types.Hash{},
		lock: new(sync.Mutex),
	}
}

// Ini 初始化
func (r *StateRoots) Ini() {
	r.m = map[types.TxType]types.Hash{}
	r.lock = new(sync.Mutex)
}

// Put 加入或修改
func (r *StateRoots) Put(k types.TxType, v types.Hash) {
	r.lock.Lock()
	r.m[k] = v
	r.lock.Unlock()
}

// Get 返回值，返回值类型，是否有返回
func (r *StateRoots) Get(k types.TxType) (types.Hash, bool) {
	r.lock.Lock()
	v, cb := r.m[k]
	var rv types.Hash
	var rs = false
	if cb {
		rv = v
		rs = true
	}
	r.lock.Unlock()
	return rv, rs
}
func (r *StateRoots) GetObj(k types.TxType) types.Hash {
	r.lock.Lock()
	v, cb := r.m[k]
	var rv types.Hash
	if cb {
		rv = v
	}
	r.lock.Unlock()
	return rv
}

// ContainsKey 判断是否包括key，如果包含key返回value的类型
func (r *StateRoots) ContainsKey(k types.TxType) (bool, string) {
	r.lock.Lock()
	v, cb := r.m[k]
	var rs = false
	var rt = ""
	if cb {
		rs = true
		rt = reflect.TypeOf(v).String()
	}
	r.lock.Unlock()
	return rs, rt
}

// Remove 移除一个对象
func (r *StateRoots) Remove(k types.TxType) (types.Hash, bool) {
	r.lock.Lock()
	v, cb := r.m[k]
	var rs = false
	var rv types.Hash
	if cb {
		rv = v
		rs = true
		delete(r.m, k)
	}
	r.lock.Unlock()
	return rv, rs
}

// ForEach 复制map用于外部遍历
func (r *StateRoots) ForEach() map[types.TxType]types.Hash {
	r.lock.Lock()
	mb := map[types.TxType]types.Hash{}
	for k, v := range r.m {
		mb[k] = v
	}
	r.lock.Unlock()
	return mb
}

// Size 返回个数
func (r *StateRoots) Size() int {
	r.lock.Lock()
	s := len(r.m)
	r.lock.Unlock()
	return s
}

type KV struct {
	Key   types.TxType
	Value types.Hash
}

// Sort 排序
func (r StateRoots) Sort() []KV {
	var newm []KV
	var keyArray []types.TxType
	for k, _ := range r.m {
		keyArray = append(keyArray, k)
	}
	sort.Sort(types.TxTypes(keyArray))
	for _, v := range keyArray {
		kv := &KV{
			Key:   v,
			Value: r.m[v],
		}
		newm = append(newm, *kv)
	}
	return newm
}

func (r *StateRoots) EncodeRLP(w io.Writer) error {
	kvs := r.Sort()
	if err := rlp.Encode(w, kvs); err != nil {
		return err
	}
	return nil
}
func (r *StateRoots) DecodeRLP(s *rlp.Stream) error {
	var data []KV
	if err := s.Decode(&data); err != nil {
		return err
	}
	for _, v := range data {
		value := v.Value
		r.Put(v.Key, value)
	}
	return nil
}

func (r *StateRoots) MarshalJSON() ([]byte, error) {
	kvs := r.Sort()
	return json.Marshal(kvs)
}
