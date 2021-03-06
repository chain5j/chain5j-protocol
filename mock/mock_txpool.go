// Code generated by MockGen. DO NOT EDIT.
// Source: protocol/txpool.go

// Package mock is a generated GoMock package.
package mock

import (
	event "github.com/chain5j/chain5j-pkg/event"
	types "github.com/chain5j/chain5j-pkg/types"
	models "github.com/chain5j/chain5j-protocol/models"
	protocol "github.com/chain5j/chain5j-protocol/protocol"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTxPools is a mock of TxPools interface
type MockTxPools struct {
	ctrl     *gomock.Controller
	recorder *MockTxPoolsMockRecorder
}

// MockTxPoolsMockRecorder is the mock recorder for MockTxPools
type MockTxPoolsMockRecorder struct {
	mock *MockTxPools
}

// NewMockTxPools creates a new mock instance
func NewMockTxPools(ctrl *gomock.Controller) *MockTxPools {
	mock := &MockTxPools{ctrl: ctrl}
	mock.recorder = &MockTxPoolsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTxPools) EXPECT() *MockTxPoolsMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockTxPools) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockTxPoolsMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockTxPools)(nil).Start))
}

// Stop mocks base method
func (m *MockTxPools) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockTxPoolsMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockTxPools)(nil).Stop))
}

// Register mocks base method
func (m *MockTxPools) Register(txType types.TxType, txPool protocol.TxPool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Register", txType, txPool)
}

// Register indicates an expected call of Register
func (mr *MockTxPoolsMockRecorder) Register(txType, txPool interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockTxPools)(nil).Register), txType, txPool)
}

// TxPool mocks base method
func (m *MockTxPools) TxPool(txType types.TxType) (protocol.TxPool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxPool", txType)
	ret0, _ := ret[0].(protocol.TxPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TxPool indicates an expected call of TxPool
func (mr *MockTxPoolsMockRecorder) TxPool(txType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxPool", reflect.TypeOf((*MockTxPools)(nil).TxPool), txType)
}

// Add mocks base method
func (m *MockTxPools) Add(peerId *models.P2PID, tx models.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", peerId, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockTxPoolsMockRecorder) Add(peerId, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTxPools)(nil).Add), peerId, tx)
}

// Get mocks base method
func (m *MockTxPools) Get(txType types.TxType, hash types.Hash) (models.Transaction, models.TxStatus) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", txType, hash)
	ret0, _ := ret[0].(models.Transaction)
	ret1, _ := ret[1].(models.TxStatus)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTxPoolsMockRecorder) Get(txType, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTxPools)(nil).Get), txType, hash)
}

// GetTxs mocks base method
func (m *MockTxPools) GetTxs(txsLimit uint64) map[types.TxType][]models.Transaction {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxs", txsLimit)
	ret0, _ := ret[0].(map[types.TxType][]models.Transaction)
	return ret0
}

// GetTxs indicates an expected call of GetTxs
func (mr *MockTxPoolsMockRecorder) GetTxs(txsLimit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxs", reflect.TypeOf((*MockTxPools)(nil).GetTxs), txsLimit)
}

// FetchTxs mocks base method
func (m *MockTxPools) FetchTxs(txsLimit, headerTimestamp uint64) models.Transactions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchTxs", txsLimit, headerTimestamp)
	ret0, _ := ret[0].(models.Transactions)
	return ret0
}

// FetchTxs indicates an expected call of FetchTxs
func (mr *MockTxPoolsMockRecorder) FetchTxs(txsLimit, headerTimestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchTxs", reflect.TypeOf((*MockTxPools)(nil).FetchTxs), txsLimit, headerTimestamp)
}

// Fallback mocks base method
func (m *MockTxPools) Fallback(txType types.TxType, txs []models.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fallback", txType, txs)
	ret0, _ := ret[0].(error)
	return ret0
}

// Fallback indicates an expected call of Fallback
func (mr *MockTxPoolsMockRecorder) Fallback(txType, txs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fallback", reflect.TypeOf((*MockTxPools)(nil).Fallback), txType, txs)
}

// Delete mocks base method
func (m *MockTxPools) Delete(txType types.TxType, txs []models.Transaction, noErr bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", txType, txs, noErr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockTxPoolsMockRecorder) Delete(txType, txs, noErr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTxPools)(nil).Delete), txType, txs, noErr)
}

// Len mocks base method
func (m *MockTxPools) Len() map[types.TxType]uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Len")
	ret0, _ := ret[0].(map[types.TxType]uint64)
	return ret0
}

// Len indicates an expected call of Len
func (mr *MockTxPoolsMockRecorder) Len() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Len", reflect.TypeOf((*MockTxPools)(nil).Len))
}

// Subscribe mocks base method
func (m *MockTxPools) Subscribe(ch chan []models.Transaction) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ch)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockTxPoolsMockRecorder) Subscribe(ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockTxPools)(nil).Subscribe), ch)
}

// MockTxPool is a mock of TxPool interface
type MockTxPool struct {
	ctrl     *gomock.Controller
	recorder *MockTxPoolMockRecorder
}

// MockTxPoolMockRecorder is the mock recorder for MockTxPool
type MockTxPoolMockRecorder struct {
	mock *MockTxPool
}

// NewMockTxPool creates a new mock instance
func NewMockTxPool(ctrl *gomock.Controller) *MockTxPool {
	mock := &MockTxPool{ctrl: ctrl}
	mock.recorder = &MockTxPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTxPool) EXPECT() *MockTxPoolMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockTxPool) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockTxPoolMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockTxPool)(nil).Start))
}

// Stop mocks base method
func (m *MockTxPool) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockTxPoolMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockTxPool)(nil).Stop))
}

// Add mocks base method
func (m *MockTxPool) Add(peerId *models.P2PID, tx models.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", peerId, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockTxPoolMockRecorder) Add(peerId, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTxPool)(nil).Add), peerId, tx)
}

// Exist mocks base method
func (m *MockTxPool) Exist(hash types.Hash) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exist", hash)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exist indicates an expected call of Exist
func (mr *MockTxPoolMockRecorder) Exist(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exist", reflect.TypeOf((*MockTxPool)(nil).Exist), hash)
}

// Get mocks base method
func (m *MockTxPool) Get(hash types.Hash) (models.Transaction, models.TxStatus) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", hash)
	ret0, _ := ret[0].(models.Transaction)
	ret1, _ := ret[1].(models.TxStatus)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTxPoolMockRecorder) Get(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTxPool)(nil).Get), hash)
}

// GetTxs mocks base method
func (m *MockTxPool) GetTxs(txsLimit uint64) []models.Transaction {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxs", txsLimit)
	ret0, _ := ret[0].([]models.Transaction)
	return ret0
}

// GetTxs indicates an expected call of GetTxs
func (mr *MockTxPoolMockRecorder) GetTxs(txsLimit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxs", reflect.TypeOf((*MockTxPool)(nil).GetTxs), txsLimit)
}

// FetchTxs mocks base method
func (m *MockTxPool) FetchTxs(txsLimit, headerTimestamp uint64) []models.Transaction {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchTxs", txsLimit, headerTimestamp)
	ret0, _ := ret[0].([]models.Transaction)
	return ret0
}

// FetchTxs indicates an expected call of FetchTxs
func (mr *MockTxPoolMockRecorder) FetchTxs(txsLimit, headerTimestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchTxs", reflect.TypeOf((*MockTxPool)(nil).FetchTxs), txsLimit, headerTimestamp)
}

// Fallback mocks base method
func (m *MockTxPool) Fallback(txs []models.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fallback", txs)
	ret0, _ := ret[0].(error)
	return ret0
}

// Fallback indicates an expected call of Fallback
func (mr *MockTxPoolMockRecorder) Fallback(txs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fallback", reflect.TypeOf((*MockTxPool)(nil).Fallback), txs)
}

// Delete mocks base method
func (m *MockTxPool) Delete(tx models.Transaction, noErr bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", tx, noErr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockTxPoolMockRecorder) Delete(tx, noErr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTxPool)(nil).Delete), tx, noErr)
}

// Len mocks base method
func (m *MockTxPool) Len() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Len")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Len indicates an expected call of Len
func (mr *MockTxPoolMockRecorder) Len() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Len", reflect.TypeOf((*MockTxPool)(nil).Len))
}
