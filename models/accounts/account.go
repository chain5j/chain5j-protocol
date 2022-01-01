// Package accounts
//
// @author: xwc1125
package accounts

import (
	"bytes"
	"fmt"
	"io"
	"math/big"
	"sort"
	"strings"

	"github.com/chain5j/chain5j-pkg/codec"
	"github.com/chain5j/chain5j-pkg/codec/json"
	"github.com/chain5j/chain5j-pkg/codec/rlp"
	"github.com/chain5j/chain5j-pkg/types"
)

type AccountOp uint

const (
	RegisterAccountOp      AccountOp = iota // 注册账户操作
	FrozenAccountOp                         // 冻结账户操作
	UpdateDataPermissionOp                  // 更新权限操作
	RegisterDomainOp                        // 注册域操作
	SetPartnerOp                            // 设置伙伴操作
	LostRequestOp                           // 丢失请求操作
	FoundRequestOp                          // 找回请求操作
	LostResetOp                             // 丢失重置操作
)

const (
	ContractDomain = "chain5j.contract" // 合约域
	DomainLinkFlag = "@"                // 域连接符
	CodeHashKey    = "code_hash"        // codeHash的键
	RootKey        = "root_hash"        // root的键
	PartnerKey     = "partner"          // 配偶的键
	LostKey        = "lost"             // 丢失的键
)

// Permissions 用户权限
type Permissions struct {
	EnableRegisterUser      bool `json:"enable_register_user,omitempty"`      // 是否允许注册用户
	EnableUpdateUser        bool `json:"enable_update_user,omitempty"`        // 是否允许更新用户权限
	EnableFrozenUser        bool `json:"enable_frozen_user,omitempty"`        // 是否允许冻结用户
	EnableRegisterDomain    bool `json:"enable_register_domain,omitempty"`    // 是否允许建立新的域名
	EnableRegisterSubdomain bool `json:"enable_register_subdomain,omitempty"` // 是否允许建立子域
}

// DomainStore 域名存储
type DomainStore struct {
	Admin  string `json:"admin"`  // 管理员名称、不包含域
	Number uint64 `json:"number"` // 域生成时所在区块
}

// AddressStore 账户下的地址属性
type AddressStore struct {
	KVS map[string][]byte `json:"kvs,omitempty" rlp:"nil"`
}

type kvStore struct {
	Key   string
	Value []byte
}

type kvStoreList []*kvStore

func (kv kvStoreList) Len() int           { return len(kv) }
func (kv kvStoreList) Less(i, j int) bool { return strings.Compare(kv[i].Key, kv[j].Key) < 0 }
func (kv kvStoreList) Swap(i, j int)      { kv[i], kv[j] = kv[j], kv[i] }
func (store *AddressStore) EncodeRLP(w io.Writer) error {
	var kvs kvStoreList
	for k, v := range store.KVS {
		kvs = append(kvs, &kvStore{Key: k, Value: v})
	}
	sort.Sort(kvs)
	return rlp.Encode(w, &kvs)
}
func (store *AddressStore) DecodeRLP(s *rlp.Stream) error {
	var kvs kvStoreList

	err := s.Decode(&kvs)
	if err != nil {
		return err
	}

	store.KVS = make(map[string][]byte)
	for _, kv := range kvs {
		if len(kv.Key) == 0 {
			continue
		}
		store.KVS[kv.Key] = kv.Value
	}

	return nil
}

// AccountStore 账户存储
type AccountStore struct {
	Nonce     uint64                          `json:"nonce,omitempty"`     // nonce，只要账户信息被修改，此值就得修改
	Balance   *big.Int                        `json:"balance,omitempty"`   // balance
	Addresses map[types.Address]*AddressStore `json:"addresses,omitempty"` // 地址集合

	CN     string `json:"cn"`     // 用户名称 common name，不带域的
	Domain string `json:"domain"` // 所在域

	IsAdmin              bool         `json:"is_admin,omitempty"`               // 是否为管理员
	IsFrozen             bool         `json:"is_frozen,omitempty"`              // 账户是否被冻结
	EnableDeployContract bool         `json:"enable_deploy_contract,omitempty"` // 是否允许部署合约
	Permissions          *Permissions `json:"permissions,omitempty" rlp:"nil"`  // 管理员权限

	XXX map[string][]byte `json:"xxx,omitempty" rlp:"nil"` // 扩展字段
}

func NewAccountStore(cn, domain string) *AccountStore {
	store := &AccountStore{
		Balance:   new(big.Int),
		CN:        cn,
		Domain:    domain,
		Addresses: make(map[types.Address]*AddressStore),
		XXX:       make(map[string][]byte),
	}
	// 将字符串全部转成小写
	store.Normalize()
	return store
}

func (store *AccountStore) Copy() *AccountStore {
	copy := &AccountStore{
		Nonce:                store.Nonce,
		Balance:              new(big.Int),
		Addresses:            make(map[types.Address]*AddressStore),
		CN:                   store.CN,
		Domain:               store.Domain,
		IsAdmin:              store.IsAdmin,
		IsFrozen:             store.IsFrozen,
		EnableDeployContract: store.EnableDeployContract,
		XXX:                  make(map[string][]byte),
	}

	if store.Balance != nil {
		copy.Balance.Set(store.Balance)
	}

	for k, v := range store.Addresses {
		copy.Addresses[k] = v
	}

	for k, v := range store.XXX {
		copy.XXX[k] = v
	}

	if store.IsAdmin && store.Permissions != nil {
		copy.Permissions = &Permissions{
			EnableRegisterUser:      store.Permissions.EnableRegisterUser,
			EnableUpdateUser:        store.Permissions.EnableUpdateUser,
			EnableFrozenUser:        store.Permissions.EnableFrozenUser,
			EnableRegisterDomain:    store.Permissions.EnableRegisterDomain,
			EnableRegisterSubdomain: store.Permissions.EnableRegisterSubdomain,
		}
	}

	return copy
}

// Normalize CN 和 domain 转化为小写
func (store *AccountStore) Normalize() {
	store.CN = strings.ToLower(store.CN)
	store.Domain = strings.ToLower(store.Domain)
}

// SetAddress 设置地址集
func (store *AccountStore) SetAddress(addr types.Address, addrStore *AddressStore) {
	store.Addresses[addr] = addrStore
}

// AccountName 账户名称
func (store *AccountStore) AccountName() string {
	if store.Domain != "" {
		// 如果域存在，那么cn需要添加域
		return store.CN + DomainLinkFlag + store.Domain
	}

	return store.CN
}

// ContainAddress 账户是否存在指定地址
func (store *AccountStore) ContainAddress(address types.Address) bool {
	_, ok := store.Addresses[address]
	return ok
}

// GetAddressStore 根据地址获取地址集
func (store *AccountStore) GetAddressStore(address types.Address) (*AddressStore, error) {
	addrStore, ok := store.Addresses[address]
	if !ok {
		return nil, fmt.Errorf(ErrFormatAddrNotExist, address)
	}

	return addrStore, nil
}

// AuthorizedRegisterUser 注册用户的权限是否已授予
func (store *AccountStore) AuthorizedRegisterUser() bool {
	if store.Permissions == nil {
		return false
	}
	return store.Permissions.EnableRegisterUser
}

// AuthorizedRegisterDomain 注册域的权限是否已授予
func (store *AccountStore) AuthorizedRegisterDomain() bool {
	if store.Permissions == nil {
		return false
	}
	return store.Permissions.EnableRegisterDomain
}

// IsContract 判断账户是否为合约地址
func (store *AccountStore) IsContract() bool {
	return store.Domain == ContractDomain
}

// SetCodeHash 设置codeHash
func (store *AccountStore) SetCodeHash(hash []byte) {
	store.XXX[CodeHashKey] = hash
}

// CodeHash 获取codeHash
func (store *AccountStore) CodeHash() []byte {
	if codeHash, ok := store.XXX[CodeHashKey]; ok {
		return codeHash
	} else {
		return types.EmptyCode
	}
}

// SetStorageRoot 设置存储的rootHash
func (store *AccountStore) SetStorageRoot(hash types.Hash) {
	store.XXX[RootKey] = hash[:]
}

// StorageRoot 获取存储的rootHash
func (store *AccountStore) StorageRoot() types.Hash {
	if root, ok := store.XXX[RootKey]; ok {
		return types.BytesToHash(root)
	} else {
		return types.EmptyRootHash
	}
}

// SetPartner 设置配偶[用于密码丢失找回]
func (store *AccountStore) SetPartner(cn, domain string) {
	pData := &PartnerData{
		CN:     cn,
		Domain: domain,
	}

	storeBytes, _ := codec.Coder().Encode(pData)
	store.XXX[PartnerKey] = storeBytes
}

// Partner 获取配偶地址
func (store *AccountStore) Partner() string {
	pBytes, ok := store.XXX[PartnerKey]
	if !ok {
		return ""
	}

	var data PartnerData
	if err := codec.Coder().Decode(pBytes, &data); err != nil {
		return ""
	}
	return data.CN + DomainLinkFlag + data.Domain
}

func (store *AccountStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(store)
}

type addrStore struct {
	Addr  types.Address
	Store *AddressStore
}

type addrStoreList []addrStore

func (a addrStoreList) Len() int           { return len(a) }
func (a addrStoreList) Less(i, j int) bool { return bytes.Compare(a[i].Addr[:], a[j].Addr[:]) < 0 }
func (a addrStoreList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type accountRlp struct {
	Nonce     uint64
	Balance   *big.Int
	Addresses addrStoreList
	CN        string
	Domain    string

	IsAdmin              bool
	EnableDeployContract bool
	Permissions          *Permissions `rlp:"nil"`

	IsFrozen bool

	XXX kvStoreList
}

func (store *AccountStore) EncodeRLP(w io.Writer) error {
	data := accountRlp{
		Nonce:                store.Nonce,
		Balance:              store.Balance,
		CN:                   store.CN,
		Domain:               store.Domain,
		IsAdmin:              store.IsAdmin,
		EnableDeployContract: store.EnableDeployContract,
		Permissions:          store.Permissions,
		IsFrozen:             store.IsFrozen,
	}

	for addr, store := range store.Addresses {
		data.Addresses = append(data.Addresses, struct {
			Addr  types.Address
			Store *AddressStore
		}{Addr: addr, Store: store})
	}
	sort.Sort(data.Addresses)

	for k, v := range store.XXX {
		data.XXX = append(data.XXX, &kvStore{Key: k, Value: v})
	}
	sort.Sort(data.XXX)

	return rlp.Encode(w, &data)
}
func (store *AccountStore) DecodeRLP(s *rlp.Stream) error {
	var data accountRlp
	err := s.Decode(&data)
	if err != nil {
		return err
	}

	store.Nonce = data.Nonce
	store.Balance = data.Balance
	store.CN = data.CN
	store.Domain = data.Domain
	store.IsAdmin = data.IsAdmin
	store.EnableDeployContract = data.EnableDeployContract
	store.Permissions = data.Permissions
	store.Addresses = make(map[types.Address]*AddressStore)
	store.IsFrozen = data.IsFrozen
	store.XXX = make(map[string][]byte)

	for _, addrStore := range data.Addresses {
		store.Addresses[addrStore.Addr] = addrStore.Store
	}

	for _, xStore := range data.XXX {
		store.XXX[xStore.Key] = xStore.Value
	}

	return nil
}

// AccountMap 地址到账户名称的映射
type AccountMap struct {
	objects map[types.Address]string
}

// AccountOpData 账户操作处理的对象
type AccountOpData struct {
	Operation AccountOp // 操作类型
	Data      []byte    // 具体操作内容
}

// DecodeAccountOpData 从交易input中解析出操作处理对象
func DecodeAccountOpData(input []byte, data *AccountOpData) error {
	if err := codec.Coder().Decode(input, &data); err != nil {
		return err
	}

	return nil
}

// FrozenAccountData 冻结/解冻账户的数据对象
type FrozenAccountData struct {
	CN     string
	Domain string
	Frozen bool
}

func (data *FrozenAccountData) Normalize() {
	data.CN = strings.ToLower(data.CN)
	data.Domain = strings.ToLower(data.Domain)
}

// UpdatePermissionData 更新账户权限的数据对象
type UpdatePermissionData struct {
	CN          string
	Domain      string
	Permissions Permissions
}

func (data *UpdatePermissionData) Normalize() {
	data.CN = strings.ToLower(data.CN)
	data.Domain = strings.ToLower(data.Domain)
}

// PartnerData 合作伙伴的数据对象
type PartnerData struct {
	CN     string
	Domain string
}

func (data *PartnerData) Normalize() {
	data.CN = strings.ToLower(data.CN)
	data.Domain = strings.ToLower(data.Domain)
}

// LostRequest 密钥丢失处理的数据对象
type LostRequest struct {
	CN          string
	Domain      string
	RecoverAddr types.Address
}

// LostStore 密钥丢失存储对象
type LostStore struct {
	*LostRequest
	TimeStamp uint64
}

func (req *LostRequest) Normalize() {
	req.CN = strings.ToLower(req.CN)
	req.Domain = strings.ToLower(req.Domain)
}
