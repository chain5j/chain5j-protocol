// Package protocol
//
// @author: xwc1125
package protocol

import (
	"math/big"
	"time"

	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/models"
	"github.com/chain5j/chain5j-protocol/models/statetype"
)

// CompileType 合约编译器类型
type CompileType int

const (
	EVM  CompileType = iota // evm
	WASM                    // wasm
)

// ContractRef 合约接口
type ContractRef interface {
	Address() types.Address
}

// VM vm调用合约的基础接口
type VM interface {
	VmName() string // vm名称

	Cancel()         // 取消合约处理
	Cancelled() bool // 合约是否已经取消

	Call(caller ContractRef, addr types.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error)
	CallCode(caller ContractRef, addr types.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error)
	DelegateCall(caller ContractRef, addr types.Address, input []byte, gas uint64) (ret []byte, leftOverGas uint64, err error)
	StaticCall(caller ContractRef, addr types.Address, input []byte, gas uint64) (ret []byte, leftOverGas uint64, err error)

	Create(caller ContractRef, code []byte, gas uint64, value *big.Int) (ret []byte, contractAddr types.Address, leftOverGas uint64, err error)
	Create2(caller ContractRef, code []byte, gas uint64, endowment *big.Int, salt *big.Int) (ret []byte, contractAddr types.Address, leftOverGas uint64, err error)

	DB() StateDB
	Coinbase() types.Address // 当前节点的地址

	// ChainConfig() *models.ChainConfig
	// GetStateDb() StateDB
	// GetContext() Context
}

// Contract 合约接口
type Contract interface {
	AsDelegate() Contract                                          // 将contract设置为委托代理调用并返回当前contract
	Caller() types.Address                                         // 当contract是委托代理调用时，将递归调用，包括调用方的调用方的委托调用。
	UseGas(gas uint64) (ok bool)                                   // 尝试使用gas并将其减去，成功后返回true
	Address() types.Address                                        // 返回合约地址
	Value() *big.Int                                               // 返回合约的值（从其调用者发送给它）
	SetCallCode(addr *types.Address, hash types.Hash, code []byte) // 设置合约地址对应的code和hash
}

type ChainContext2 interface {
	GetStateDb() StateDB
	GetOrigin() types.Address
	GetTime() *big.Int
	GetBlockNum() *big.Int
}

type ChainContext interface {
	GetHeader(bHash types.Hash, height uint64) *models.Header // 获取header
}

// StateDB is an EVM database for full state querying.
// type StateDB interface {
//	//Reset(root types.Hash) error// evm no use
//	//StorageTree(addr types.DomainAddress) tree.Tree// evm no use
//
//	Suicide(address types.DomainAddress) bool
//	//HasSuicided(addr types.DomainAddress) bool// evm no use
//
//	// Exist reports whether the given account exists in state.
//	// Notably this should also return true for suicided accounts.
//	Exist(types.DomainAddress) bool
//	// Empty returns whether the given account is empty. Empty
//	// is defined according to EIP161 (balance = nonce = code = 0).
//	Empty(types.DomainAddress) bool
//
//	RevertToSnapshot(int)
//	Snapshot() int
//
//	CreateAccount(types.DomainAddress)
//
//	SubBalance(types.DomainAddress, *big.Int)
//	AddBalance(types.DomainAddress, *big.Int)
//	GetBalance(types.DomainAddress) *big.Int
//
//	GetNonce(types.DomainAddress) uint64
//	SetNonce(types.DomainAddress, uint64)
//
//	GetCodeHash(types.DomainAddress) types.Hash
//	GetCode(types.DomainAddress) []byte
//	SetCode(types.DomainAddress, []byte)
//	GetCodeSize(types.DomainAddress) int
//
//	GetState(types.DomainAddress, types.Hash) types.Hash
//	SetState(types.DomainAddress, types.Hash, types.Hash)
//	//GetCommittedState(addr types.DomainAddress, hash types.Hash) types.Hash
//
//	//Commit(deleteEmptyObjects bool) (root types.Hash, err error)
//	//Finalise(deleteEmptyObjects bool)
//	//IntermediateRoot(deleteEmptyObjects bool) types.Hash
//
//	//ChargeGas(types.DomainAddress, uint64)
//
//	AddLog(*statetype.Log)
//	//GetLogs(hash types.Hash) []*statetype.Log
//	//Logs() []*statetype.Log
//
//	//GetContractRef(address types.DomainAddress) ContractRef
//
//	AddRefund(uint64)
//	//SubRefund(uint64)
//	GetRefund() uint64 // 在state_transiction中用到
//
//	//SubTokenBalance(addr types.DomainAddress, token types.DomainAddress, amount *big.Int)
//	//AddTokenBalance(addr types.DomainAddress, token types.DomainAddress, amount *big.Int)
//	//GetTokenBalance(addr types.DomainAddress, token types.DomainAddress) *big.Int
//	//GetTokenBalances(addr types.DomainAddress) statetype.TokenValues
//	//
//	//GetContractInfo(addr types.DomainAddress) []byte
//	//SetContractInfo(addr types.DomainAddress, info []byte)
// }

// StateDB 状态DB
type StateDB interface {
	CreateAccount(addr types.Address) // 创建账户

	SubBalance(addr types.Address, amount *big.Int) // 减少余额
	AddBalance(addr types.Address, amount *big.Int) // 增加余额
	GetBalance(addr types.Address) *big.Int         // 获取余额

	GetNonce(addr types.Address) uint64        // 获取nonce
	SetNonce(addr types.Address, nonce uint64) // 设置nonce

	GetCodeHash(addr types.Address) types.Hash // 根据contract获取对应的codeHash
	GetCode(addr types.Address) []byte         // 根据contract获取code
	SetCode(addr types.Address, code []byte)   // 设置contract对应的code
	GetCodeSize(addr types.Address) int        // 获取code对应的size

	AddRefund(gas uint64) // 添加退款
	SubRefund(gas uint64) // 减少退款
	GetRefund() uint64    // 获取退款

	GetCommittedState(addr types.Address, hash types.Hash) types.Hash
	GetState(addr types.Address, hash types.Hash) types.Hash
	SetState(addr types.Address, key types.Hash, value types.Hash)

	Suicide(addr types.Address) bool
	HasSuicided(addr types.Address) bool

	// Exist reports whether the given account exists in state.
	// Notably this should also return true for suicided accounts.
	Exist(addr types.Address) bool
	// Empty returns whether the given account is empty. Empty
	// is defined according to EIP161 (balance = nonce = code = 0).
	Empty(addr types.Address) bool

	RevertToSnapshot(int)
	Snapshot() int

	AddLog(log *statetype.Log)
	AddPreimage(hash types.Hash, preimage []byte)

	ForEachStorage(addr types.Address, cb func(key, value types.Hash) bool) error
}

// contract_native.go
// PrecompiledContract is the basic interface for native Go contracts.
// The implementation requires a deterministic gas count based on
// the input size of the Run method of the contract.
type PrecompiledContract interface {
	RequiredGas(input []byte) uint64  // RequiredPrice calculates the contract gas use
	Run(input []byte) ([]byte, error) // Run runs the precompiled contract
}

// Tracer is used to collect execution traces from an VM transaction
// execution. CaptureState is called for each step of the VM with the
// current VM state.
// Note that reference types are actual VM data structures; make copies
// if you need to retain them beyond the current call.
type Tracer interface {
	CaptureStart(from types.Address, to types.Address, call bool, input []byte, gas uint64, value *big.Int) error
	CaptureState(env VM, pc uint64, op OPCode, gas, cost uint64, contract Contract, depth int, err error) error
	CaptureLog(env VM, msg string) error
	CaptureFault(env VM, pc uint64, op OPCode, gas, cost uint64, contract Contract, depth int, err error) error
	CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error) error
}

type OPCode interface {
	IsPush() bool
	String() string
	Byte() byte
}
