// Code generated by MockGen. DO NOT EDIT.
// Source: protocol/vm.go

// Package mock is a generated GoMock package.
package mock

import (
	types "github.com/chain5j/chain5j-pkg/types"
	models "github.com/chain5j/chain5j-protocol/models"
	statetype "github.com/chain5j/chain5j-protocol/models/statetype"
	protocol "github.com/chain5j/chain5j-protocol/protocol"
	gomock "github.com/golang/mock/gomock"
	big "math/big"
	reflect "reflect"
	time "time"
)

// MockContractRef is a mock of ContractRef interface
type MockContractRef struct {
	ctrl     *gomock.Controller
	recorder *MockContractRefMockRecorder
}

// MockContractRefMockRecorder is the mock recorder for MockContractRef
type MockContractRefMockRecorder struct {
	mock *MockContractRef
}

// NewMockContractRef creates a new mock instance
func NewMockContractRef(ctrl *gomock.Controller) *MockContractRef {
	mock := &MockContractRef{ctrl: ctrl}
	mock.recorder = &MockContractRefMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContractRef) EXPECT() *MockContractRefMockRecorder {
	return m.recorder
}

// Address mocks base method
func (m *MockContractRef) Address() types.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(types.Address)
	return ret0
}

// Address indicates an expected call of Address
func (mr *MockContractRefMockRecorder) Address() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockContractRef)(nil).Address))
}

// MockVM is a mock of VM interface
type MockVM struct {
	ctrl     *gomock.Controller
	recorder *MockVMMockRecorder
}

// MockVMMockRecorder is the mock recorder for MockVM
type MockVMMockRecorder struct {
	mock *MockVM
}

// NewMockVM creates a new mock instance
func NewMockVM(ctrl *gomock.Controller) *MockVM {
	mock := &MockVM{ctrl: ctrl}
	mock.recorder = &MockVMMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVM) EXPECT() *MockVMMockRecorder {
	return m.recorder
}

// VmName mocks base method
func (m *MockVM) VmName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VmName")
	ret0, _ := ret[0].(string)
	return ret0
}

// VmName indicates an expected call of VmName
func (mr *MockVMMockRecorder) VmName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VmName", reflect.TypeOf((*MockVM)(nil).VmName))
}

// Cancel mocks base method
func (m *MockVM) Cancel() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Cancel")
}

// Cancel indicates an expected call of Cancel
func (mr *MockVMMockRecorder) Cancel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockVM)(nil).Cancel))
}

// Cancelled mocks base method
func (m *MockVM) Cancelled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancelled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Cancelled indicates an expected call of Cancelled
func (mr *MockVMMockRecorder) Cancelled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancelled", reflect.TypeOf((*MockVM)(nil).Cancelled))
}

// Call mocks base method
func (m *MockVM) Call(caller protocol.ContractRef, addr types.Address, input []byte, gas uint64, value *big.Int) ([]byte, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", caller, addr, input, gas, value)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Call indicates an expected call of Call
func (mr *MockVMMockRecorder) Call(caller, addr, input, gas, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockVM)(nil).Call), caller, addr, input, gas, value)
}

// CallCode mocks base method
func (m *MockVM) CallCode(caller protocol.ContractRef, addr types.Address, input []byte, gas uint64, value *big.Int) ([]byte, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallCode", caller, addr, input, gas, value)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CallCode indicates an expected call of CallCode
func (mr *MockVMMockRecorder) CallCode(caller, addr, input, gas, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallCode", reflect.TypeOf((*MockVM)(nil).CallCode), caller, addr, input, gas, value)
}

// DelegateCall mocks base method
func (m *MockVM) DelegateCall(caller protocol.ContractRef, addr types.Address, input []byte, gas uint64) ([]byte, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelegateCall", caller, addr, input, gas)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DelegateCall indicates an expected call of DelegateCall
func (mr *MockVMMockRecorder) DelegateCall(caller, addr, input, gas interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelegateCall", reflect.TypeOf((*MockVM)(nil).DelegateCall), caller, addr, input, gas)
}

// StaticCall mocks base method
func (m *MockVM) StaticCall(caller protocol.ContractRef, addr types.Address, input []byte, gas uint64) ([]byte, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StaticCall", caller, addr, input, gas)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// StaticCall indicates an expected call of StaticCall
func (mr *MockVMMockRecorder) StaticCall(caller, addr, input, gas interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StaticCall", reflect.TypeOf((*MockVM)(nil).StaticCall), caller, addr, input, gas)
}

// Create mocks base method
func (m *MockVM) Create(caller protocol.ContractRef, code []byte, gas uint64, value *big.Int) ([]byte, types.Address, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", caller, code, gas, value)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(types.Address)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Create indicates an expected call of Create
func (mr *MockVMMockRecorder) Create(caller, code, gas, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVM)(nil).Create), caller, code, gas, value)
}

// Create2 mocks base method
func (m *MockVM) Create2(caller protocol.ContractRef, code []byte, gas uint64, endowment, salt *big.Int) ([]byte, types.Address, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create2", caller, code, gas, endowment, salt)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(types.Address)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Create2 indicates an expected call of Create2
func (mr *MockVMMockRecorder) Create2(caller, code, gas, endowment, salt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create2", reflect.TypeOf((*MockVM)(nil).Create2), caller, code, gas, endowment, salt)
}

// DB mocks base method
func (m *MockVM) DB() protocol.StateDB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DB")
	ret0, _ := ret[0].(protocol.StateDB)
	return ret0
}

// DB indicates an expected call of DB
func (mr *MockVMMockRecorder) DB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockVM)(nil).DB))
}

// Coinbase mocks base method
func (m *MockVM) Coinbase() types.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Coinbase")
	ret0, _ := ret[0].(types.Address)
	return ret0
}

// Coinbase indicates an expected call of Coinbase
func (mr *MockVMMockRecorder) Coinbase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Coinbase", reflect.TypeOf((*MockVM)(nil).Coinbase))
}

// MockContract is a mock of Contract interface
type MockContract struct {
	ctrl     *gomock.Controller
	recorder *MockContractMockRecorder
}

// MockContractMockRecorder is the mock recorder for MockContract
type MockContractMockRecorder struct {
	mock *MockContract
}

// NewMockContract creates a new mock instance
func NewMockContract(ctrl *gomock.Controller) *MockContract {
	mock := &MockContract{ctrl: ctrl}
	mock.recorder = &MockContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContract) EXPECT() *MockContractMockRecorder {
	return m.recorder
}

// AsDelegate mocks base method
func (m *MockContract) AsDelegate() protocol.Contract {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AsDelegate")
	ret0, _ := ret[0].(protocol.Contract)
	return ret0
}

// AsDelegate indicates an expected call of AsDelegate
func (mr *MockContractMockRecorder) AsDelegate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsDelegate", reflect.TypeOf((*MockContract)(nil).AsDelegate))
}

// Caller mocks base method
func (m *MockContract) Caller() types.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Caller")
	ret0, _ := ret[0].(types.Address)
	return ret0
}

// Caller indicates an expected call of Caller
func (mr *MockContractMockRecorder) Caller() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Caller", reflect.TypeOf((*MockContract)(nil).Caller))
}

// UseGas mocks base method
func (m *MockContract) UseGas(gas uint64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UseGas", gas)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UseGas indicates an expected call of UseGas
func (mr *MockContractMockRecorder) UseGas(gas interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseGas", reflect.TypeOf((*MockContract)(nil).UseGas), gas)
}

// Address mocks base method
func (m *MockContract) Address() types.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(types.Address)
	return ret0
}

// Address indicates an expected call of Address
func (mr *MockContractMockRecorder) Address() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockContract)(nil).Address))
}

// Value mocks base method
func (m *MockContract) Value() *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// Value indicates an expected call of Value
func (mr *MockContractMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockContract)(nil).Value))
}

// SetCallCode mocks base method
func (m *MockContract) SetCallCode(addr *types.Address, hash types.Hash, code []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCallCode", addr, hash, code)
}

// SetCallCode indicates an expected call of SetCallCode
func (mr *MockContractMockRecorder) SetCallCode(addr, hash, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCallCode", reflect.TypeOf((*MockContract)(nil).SetCallCode), addr, hash, code)
}

// MockChainContext2 is a mock of ChainContext2 interface
type MockChainContext2 struct {
	ctrl     *gomock.Controller
	recorder *MockChainContext2MockRecorder
}

// MockChainContext2MockRecorder is the mock recorder for MockChainContext2
type MockChainContext2MockRecorder struct {
	mock *MockChainContext2
}

// NewMockChainContext2 creates a new mock instance
func NewMockChainContext2(ctrl *gomock.Controller) *MockChainContext2 {
	mock := &MockChainContext2{ctrl: ctrl}
	mock.recorder = &MockChainContext2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChainContext2) EXPECT() *MockChainContext2MockRecorder {
	return m.recorder
}

// GetStateDb mocks base method
func (m *MockChainContext2) GetStateDb() protocol.StateDB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateDb")
	ret0, _ := ret[0].(protocol.StateDB)
	return ret0
}

// GetStateDb indicates an expected call of GetStateDb
func (mr *MockChainContext2MockRecorder) GetStateDb() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateDb", reflect.TypeOf((*MockChainContext2)(nil).GetStateDb))
}

// GetOrigin mocks base method
func (m *MockChainContext2) GetOrigin() types.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrigin")
	ret0, _ := ret[0].(types.Address)
	return ret0
}

// GetOrigin indicates an expected call of GetOrigin
func (mr *MockChainContext2MockRecorder) GetOrigin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrigin", reflect.TypeOf((*MockChainContext2)(nil).GetOrigin))
}

// GetTime mocks base method
func (m *MockChainContext2) GetTime() *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTime")
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// GetTime indicates an expected call of GetTime
func (mr *MockChainContext2MockRecorder) GetTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTime", reflect.TypeOf((*MockChainContext2)(nil).GetTime))
}

// GetBlockNum mocks base method
func (m *MockChainContext2) GetBlockNum() *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockNum")
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// GetBlockNum indicates an expected call of GetBlockNum
func (mr *MockChainContext2MockRecorder) GetBlockNum() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockNum", reflect.TypeOf((*MockChainContext2)(nil).GetBlockNum))
}

// MockChainContext is a mock of ChainContext interface
type MockChainContext struct {
	ctrl     *gomock.Controller
	recorder *MockChainContextMockRecorder
}

// MockChainContextMockRecorder is the mock recorder for MockChainContext
type MockChainContextMockRecorder struct {
	mock *MockChainContext
}

// NewMockChainContext creates a new mock instance
func NewMockChainContext(ctrl *gomock.Controller) *MockChainContext {
	mock := &MockChainContext{ctrl: ctrl}
	mock.recorder = &MockChainContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChainContext) EXPECT() *MockChainContextMockRecorder {
	return m.recorder
}

// GetHeader mocks base method
func (m *MockChainContext) GetHeader(bHash types.Hash, height uint64) *models.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeader", bHash, height)
	ret0, _ := ret[0].(*models.Header)
	return ret0
}

// GetHeader indicates an expected call of GetHeader
func (mr *MockChainContextMockRecorder) GetHeader(bHash, height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeader", reflect.TypeOf((*MockChainContext)(nil).GetHeader), bHash, height)
}

// MockStateDB is a mock of StateDB interface
type MockStateDB struct {
	ctrl     *gomock.Controller
	recorder *MockStateDBMockRecorder
}

// MockStateDBMockRecorder is the mock recorder for MockStateDB
type MockStateDBMockRecorder struct {
	mock *MockStateDB
}

// NewMockStateDB creates a new mock instance
func NewMockStateDB(ctrl *gomock.Controller) *MockStateDB {
	mock := &MockStateDB{ctrl: ctrl}
	mock.recorder = &MockStateDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStateDB) EXPECT() *MockStateDBMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method
func (m *MockStateDB) CreateAccount(addr types.Address) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateAccount", addr)
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockStateDBMockRecorder) CreateAccount(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockStateDB)(nil).CreateAccount), addr)
}

// SubBalance mocks base method
func (m *MockStateDB) SubBalance(addr types.Address, amount *big.Int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SubBalance", addr, amount)
}

// SubBalance indicates an expected call of SubBalance
func (mr *MockStateDBMockRecorder) SubBalance(addr, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubBalance", reflect.TypeOf((*MockStateDB)(nil).SubBalance), addr, amount)
}

// AddBalance mocks base method
func (m *MockStateDB) AddBalance(addr types.Address, amount *big.Int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddBalance", addr, amount)
}

// AddBalance indicates an expected call of AddBalance
func (mr *MockStateDBMockRecorder) AddBalance(addr, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBalance", reflect.TypeOf((*MockStateDB)(nil).AddBalance), addr, amount)
}

// GetBalance mocks base method
func (m *MockStateDB) GetBalance(addr types.Address) *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", addr)
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// GetBalance indicates an expected call of GetBalance
func (mr *MockStateDBMockRecorder) GetBalance(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockStateDB)(nil).GetBalance), addr)
}

// GetNonce mocks base method
func (m *MockStateDB) GetNonce(addr types.Address) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNonce", addr)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetNonce indicates an expected call of GetNonce
func (mr *MockStateDBMockRecorder) GetNonce(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNonce", reflect.TypeOf((*MockStateDB)(nil).GetNonce), addr)
}

// SetNonce mocks base method
func (m *MockStateDB) SetNonce(addr types.Address, nonce uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNonce", addr, nonce)
}

// SetNonce indicates an expected call of SetNonce
func (mr *MockStateDBMockRecorder) SetNonce(addr, nonce interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNonce", reflect.TypeOf((*MockStateDB)(nil).SetNonce), addr, nonce)
}

// GetCodeHash mocks base method
func (m *MockStateDB) GetCodeHash(addr types.Address) types.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCodeHash", addr)
	ret0, _ := ret[0].(types.Hash)
	return ret0
}

// GetCodeHash indicates an expected call of GetCodeHash
func (mr *MockStateDBMockRecorder) GetCodeHash(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCodeHash", reflect.TypeOf((*MockStateDB)(nil).GetCodeHash), addr)
}

// GetCode mocks base method
func (m *MockStateDB) GetCode(addr types.Address) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCode", addr)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetCode indicates an expected call of GetCode
func (mr *MockStateDBMockRecorder) GetCode(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCode", reflect.TypeOf((*MockStateDB)(nil).GetCode), addr)
}

// SetCode mocks base method
func (m *MockStateDB) SetCode(addr types.Address, code []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCode", addr, code)
}

// SetCode indicates an expected call of SetCode
func (mr *MockStateDBMockRecorder) SetCode(addr, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCode", reflect.TypeOf((*MockStateDB)(nil).SetCode), addr, code)
}

// GetCodeSize mocks base method
func (m *MockStateDB) GetCodeSize(addr types.Address) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCodeSize", addr)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetCodeSize indicates an expected call of GetCodeSize
func (mr *MockStateDBMockRecorder) GetCodeSize(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCodeSize", reflect.TypeOf((*MockStateDB)(nil).GetCodeSize), addr)
}

// AddRefund mocks base method
func (m *MockStateDB) AddRefund(gas uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddRefund", gas)
}

// AddRefund indicates an expected call of AddRefund
func (mr *MockStateDBMockRecorder) AddRefund(gas interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRefund", reflect.TypeOf((*MockStateDB)(nil).AddRefund), gas)
}

// SubRefund mocks base method
func (m *MockStateDB) SubRefund(gas uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SubRefund", gas)
}

// SubRefund indicates an expected call of SubRefund
func (mr *MockStateDBMockRecorder) SubRefund(gas interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubRefund", reflect.TypeOf((*MockStateDB)(nil).SubRefund), gas)
}

// GetRefund mocks base method
func (m *MockStateDB) GetRefund() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRefund")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetRefund indicates an expected call of GetRefund
func (mr *MockStateDBMockRecorder) GetRefund() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRefund", reflect.TypeOf((*MockStateDB)(nil).GetRefund))
}

// GetCommittedState mocks base method
func (m *MockStateDB) GetCommittedState(addr types.Address, hash types.Hash) types.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommittedState", addr, hash)
	ret0, _ := ret[0].(types.Hash)
	return ret0
}

// GetCommittedState indicates an expected call of GetCommittedState
func (mr *MockStateDBMockRecorder) GetCommittedState(addr, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommittedState", reflect.TypeOf((*MockStateDB)(nil).GetCommittedState), addr, hash)
}

// GetState mocks base method
func (m *MockStateDB) GetState(addr types.Address, hash types.Hash) types.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetState", addr, hash)
	ret0, _ := ret[0].(types.Hash)
	return ret0
}

// GetState indicates an expected call of GetState
func (mr *MockStateDBMockRecorder) GetState(addr, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockStateDB)(nil).GetState), addr, hash)
}

// SetState mocks base method
func (m *MockStateDB) SetState(addr types.Address, key, value types.Hash) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetState", addr, key, value)
}

// SetState indicates an expected call of SetState
func (mr *MockStateDBMockRecorder) SetState(addr, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetState", reflect.TypeOf((*MockStateDB)(nil).SetState), addr, key, value)
}

// Suicide mocks base method
func (m *MockStateDB) Suicide(addr types.Address) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Suicide", addr)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Suicide indicates an expected call of Suicide
func (mr *MockStateDBMockRecorder) Suicide(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Suicide", reflect.TypeOf((*MockStateDB)(nil).Suicide), addr)
}

// HasSuicided mocks base method
func (m *MockStateDB) HasSuicided(addr types.Address) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasSuicided", addr)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasSuicided indicates an expected call of HasSuicided
func (mr *MockStateDBMockRecorder) HasSuicided(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasSuicided", reflect.TypeOf((*MockStateDB)(nil).HasSuicided), addr)
}

// Exist mocks base method
func (m *MockStateDB) Exist(addr types.Address) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exist", addr)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exist indicates an expected call of Exist
func (mr *MockStateDBMockRecorder) Exist(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exist", reflect.TypeOf((*MockStateDB)(nil).Exist), addr)
}

// Empty mocks base method
func (m *MockStateDB) Empty(addr types.Address) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Empty", addr)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Empty indicates an expected call of Empty
func (mr *MockStateDBMockRecorder) Empty(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Empty", reflect.TypeOf((*MockStateDB)(nil).Empty), addr)
}

// RevertToSnapshot mocks base method
func (m *MockStateDB) RevertToSnapshot(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RevertToSnapshot", arg0)
}

// RevertToSnapshot indicates an expected call of RevertToSnapshot
func (mr *MockStateDBMockRecorder) RevertToSnapshot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevertToSnapshot", reflect.TypeOf((*MockStateDB)(nil).RevertToSnapshot), arg0)
}

// Snapshot mocks base method
func (m *MockStateDB) Snapshot() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshot")
	ret0, _ := ret[0].(int)
	return ret0
}

// Snapshot indicates an expected call of Snapshot
func (mr *MockStateDBMockRecorder) Snapshot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockStateDB)(nil).Snapshot))
}

// AddLog mocks base method
func (m *MockStateDB) AddLog(log *statetype.Log) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddLog", log)
}

// AddLog indicates an expected call of AddLog
func (mr *MockStateDBMockRecorder) AddLog(log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLog", reflect.TypeOf((*MockStateDB)(nil).AddLog), log)
}

// AddPreimage mocks base method
func (m *MockStateDB) AddPreimage(hash types.Hash, preimage []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddPreimage", hash, preimage)
}

// AddPreimage indicates an expected call of AddPreimage
func (mr *MockStateDBMockRecorder) AddPreimage(hash, preimage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPreimage", reflect.TypeOf((*MockStateDB)(nil).AddPreimage), hash, preimage)
}

// ForEachStorage mocks base method
func (m *MockStateDB) ForEachStorage(addr types.Address, cb func(types.Hash, types.Hash) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForEachStorage", addr, cb)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForEachStorage indicates an expected call of ForEachStorage
func (mr *MockStateDBMockRecorder) ForEachStorage(addr, cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForEachStorage", reflect.TypeOf((*MockStateDB)(nil).ForEachStorage), addr, cb)
}

// MockPrecompiledContract is a mock of PrecompiledContract interface
type MockPrecompiledContract struct {
	ctrl     *gomock.Controller
	recorder *MockPrecompiledContractMockRecorder
}

// MockPrecompiledContractMockRecorder is the mock recorder for MockPrecompiledContract
type MockPrecompiledContractMockRecorder struct {
	mock *MockPrecompiledContract
}

// NewMockPrecompiledContract creates a new mock instance
func NewMockPrecompiledContract(ctrl *gomock.Controller) *MockPrecompiledContract {
	mock := &MockPrecompiledContract{ctrl: ctrl}
	mock.recorder = &MockPrecompiledContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPrecompiledContract) EXPECT() *MockPrecompiledContractMockRecorder {
	return m.recorder
}

// RequiredGas mocks base method
func (m *MockPrecompiledContract) RequiredGas(input []byte) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequiredGas", input)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// RequiredGas indicates an expected call of RequiredGas
func (mr *MockPrecompiledContractMockRecorder) RequiredGas(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequiredGas", reflect.TypeOf((*MockPrecompiledContract)(nil).RequiredGas), input)
}

// Run mocks base method
func (m *MockPrecompiledContract) Run(input []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", input)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run
func (mr *MockPrecompiledContractMockRecorder) Run(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockPrecompiledContract)(nil).Run), input)
}

// MockTracer is a mock of Tracer interface
type MockTracer struct {
	ctrl     *gomock.Controller
	recorder *MockTracerMockRecorder
}

// MockTracerMockRecorder is the mock recorder for MockTracer
type MockTracerMockRecorder struct {
	mock *MockTracer
}

// NewMockTracer creates a new mock instance
func NewMockTracer(ctrl *gomock.Controller) *MockTracer {
	mock := &MockTracer{ctrl: ctrl}
	mock.recorder = &MockTracerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTracer) EXPECT() *MockTracerMockRecorder {
	return m.recorder
}

// CaptureStart mocks base method
func (m *MockTracer) CaptureStart(from, to types.Address, call bool, input []byte, gas uint64, value *big.Int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CaptureStart", from, to, call, input, gas, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// CaptureStart indicates an expected call of CaptureStart
func (mr *MockTracerMockRecorder) CaptureStart(from, to, call, input, gas, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CaptureStart", reflect.TypeOf((*MockTracer)(nil).CaptureStart), from, to, call, input, gas, value)
}

// CaptureState mocks base method
func (m *MockTracer) CaptureState(env protocol.VM, pc uint64, op protocol.OPCode, gas, cost uint64, contract protocol.Contract, depth int, err error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CaptureState", env, pc, op, gas, cost, contract, depth, err)
	ret0, _ := ret[0].(error)
	return ret0
}

// CaptureState indicates an expected call of CaptureState
func (mr *MockTracerMockRecorder) CaptureState(env, pc, op, gas, cost, contract, depth, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CaptureState", reflect.TypeOf((*MockTracer)(nil).CaptureState), env, pc, op, gas, cost, contract, depth, err)
}

// CaptureLog mocks base method
func (m *MockTracer) CaptureLog(env protocol.VM, msg string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CaptureLog", env, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// CaptureLog indicates an expected call of CaptureLog
func (mr *MockTracerMockRecorder) CaptureLog(env, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CaptureLog", reflect.TypeOf((*MockTracer)(nil).CaptureLog), env, msg)
}

// CaptureFault mocks base method
func (m *MockTracer) CaptureFault(env protocol.VM, pc uint64, op protocol.OPCode, gas, cost uint64, contract protocol.Contract, depth int, err error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CaptureFault", env, pc, op, gas, cost, contract, depth, err)
	ret0, _ := ret[0].(error)
	return ret0
}

// CaptureFault indicates an expected call of CaptureFault
func (mr *MockTracerMockRecorder) CaptureFault(env, pc, op, gas, cost, contract, depth, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CaptureFault", reflect.TypeOf((*MockTracer)(nil).CaptureFault), env, pc, op, gas, cost, contract, depth, err)
}

// CaptureEnd mocks base method
func (m *MockTracer) CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CaptureEnd", output, gasUsed, t, err)
	ret0, _ := ret[0].(error)
	return ret0
}

// CaptureEnd indicates an expected call of CaptureEnd
func (mr *MockTracerMockRecorder) CaptureEnd(output, gasUsed, t, err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CaptureEnd", reflect.TypeOf((*MockTracer)(nil).CaptureEnd), output, gasUsed, t, err)
}

// MockOPCode is a mock of OPCode interface
type MockOPCode struct {
	ctrl     *gomock.Controller
	recorder *MockOPCodeMockRecorder
}

// MockOPCodeMockRecorder is the mock recorder for MockOPCode
type MockOPCodeMockRecorder struct {
	mock *MockOPCode
}

// NewMockOPCode creates a new mock instance
func NewMockOPCode(ctrl *gomock.Controller) *MockOPCode {
	mock := &MockOPCode{ctrl: ctrl}
	mock.recorder = &MockOPCodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOPCode) EXPECT() *MockOPCodeMockRecorder {
	return m.recorder
}

// IsPush mocks base method
func (m *MockOPCode) IsPush() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPush")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPush indicates an expected call of IsPush
func (mr *MockOPCodeMockRecorder) IsPush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPush", reflect.TypeOf((*MockOPCode)(nil).IsPush))
}

// String mocks base method
func (m *MockOPCode) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockOPCodeMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockOPCode)(nil).String))
}

// Byte mocks base method
func (m *MockOPCode) Byte() byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Byte")
	ret0, _ := ret[0].(byte)
	return ret0
}

// Byte indicates an expected call of Byte
func (mr *MockOPCodeMockRecorder) Byte() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Byte", reflect.TypeOf((*MockOPCode)(nil).Byte))
}
