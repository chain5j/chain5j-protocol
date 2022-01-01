// Package permission
//
// @author: xwc1125
package permission

import (
	"io"
	"sort"

	"github.com/chain5j/chain5j-pkg/codec/rlp"
)

// MemberInfo 成员信息
type MemberInfo struct {
	Name   string `json:"name"`
	Height uint64 `json:"height"` // 区块链高度【Admin&Peer和Height有关】
}

type MemberInfoMap map[string]MemberInfo

func NewMemberInfoMap() MemberInfoMap {
	return MemberInfoMap{}
}

func (m *MemberInfoMap) Put(key string, info MemberInfo) {
	infoMap := *m
	infoMap[key] = info
	*m = infoMap
}

type KMemberInfo struct {
	Key   string
	Value MemberInfo
}

// Sort 排序
func (m *MemberInfoMap) Sort() []KMemberInfo {
	var infos []KMemberInfo
	var keyArray []string
	infoMap := *m
	for k, _ := range infoMap {
		keyArray = append(keyArray, k)
	}
	sort.Strings(keyArray)
	for _, v := range keyArray {
		infos = append(infos, KMemberInfo{
			Key:   v,
			Value: infoMap[v],
		})
	}
	return infos
}

func (m *MemberInfoMap) EncodeRLP(w io.Writer) error {
	data := m.Sort()
	return rlp.Encode(w, data)
}

func (m *MemberInfoMap) DecodeRLP(s *rlp.Stream) error {
	var data []KMemberInfo
	if err := s.Decode(&data); err != nil {
		return err
	}
	for _, kv := range data {
		m.Put(kv.Key, kv.Value)
	}
	return nil
}

// PeerPermissionAlloc peer权限分配
type PeerPermissionAlloc struct {
	Colleague MemberInfoMap `json:"colleague"`
	Observer  MemberInfoMap `json:"observer"`
}

func (i *PeerPermissionAlloc) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(i)
}
func (i *PeerPermissionAlloc) Decode(bytes []byte) error {
	return rlp.DecodeBytes(bytes, i)
}

// GenesisPermission 创世权限
type GenesisPermission struct {
	Type                ChainAccessType                `json:"type"`                  // 权限类型
	SupervisorAlloc     MemberInfoMap                  `json:"supervisor_alloc"`      // 监管预设
	PeerPermissionAlloc map[string]PeerPermissionAlloc `json:"peer_permission_alloc"` // 节点预设
}
