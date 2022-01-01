// Code generated by MockGen. DO NOT EDIT.
// Source: protocol/apis.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	types "github.com/chain5j/chain5j-pkg/types"
	hexutil "github.com/chain5j/chain5j-pkg/util/hexutil"
	models "github.com/chain5j/chain5j-protocol/models"
	statetype "github.com/chain5j/chain5j-protocol/models/statetype"
	protocol "github.com/chain5j/chain5j-protocol/protocol"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAPIs is a mock of APIs interface
type MockAPIs struct {
	ctrl     *gomock.Controller
	recorder *MockAPIsMockRecorder
}

// MockAPIsMockRecorder is the mock recorder for MockAPIs
type MockAPIsMockRecorder struct {
	mock *MockAPIs
}

// NewMockAPIs creates a new mock instance
func NewMockAPIs(ctrl *gomock.Controller) *MockAPIs {
	mock := &MockAPIs{ctrl: ctrl}
	mock.recorder = &MockAPIsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPIs) EXPECT() *MockAPIsMockRecorder {
	return m.recorder
}

// Syncing mocks base method
func (m *MockAPIs) Syncing(ctx context.Context) (*models.SyncingStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Syncing", ctx)
	ret0, _ := ret[0].(*models.SyncingStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Syncing indicates an expected call of Syncing
func (mr *MockAPIsMockRecorder) Syncing(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Syncing", reflect.TypeOf((*MockAPIs)(nil).Syncing), ctx)
}

// GetTransactionCount mocks base method
func (m *MockAPIs) GetTransactionCount(ctx context.Context, txType types.TxType, address string) (*hexutil.Uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionCount", ctx, txType, address)
	ret0, _ := ret[0].(*hexutil.Uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionCount indicates an expected call of GetTransactionCount
func (mr *MockAPIsMockRecorder) GetTransactionCount(ctx, txType, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionCount", reflect.TypeOf((*MockAPIs)(nil).GetTransactionCount), ctx, txType, address)
}

// SendRawTransaction mocks base method
func (m *MockAPIs) SendRawTransaction(ctx context.Context, txType types.TxType, rawTx hexutil.Bytes) (types.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRawTransaction", ctx, txType, rawTx)
	ret0, _ := ret[0].(types.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendRawTransaction indicates an expected call of SendRawTransaction
func (mr *MockAPIsMockRecorder) SendRawTransaction(ctx, txType, rawTx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRawTransaction", reflect.TypeOf((*MockAPIs)(nil).SendRawTransaction), ctx, txType, rawTx)
}

// GetTransaction mocks base method
func (m *MockAPIs) GetTransaction(ctx context.Context, hash types.Hash) models.Transaction {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransaction", ctx, hash)
	ret0, _ := ret[0].(models.Transaction)
	return ret0
}

// GetTransaction indicates an expected call of GetTransaction
func (mr *MockAPIsMockRecorder) GetTransaction(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockAPIs)(nil).GetTransaction), ctx, hash)
}

// GetTransactionReceipt mocks base method
func (m *MockAPIs) GetTransactionReceipt(ctx context.Context, hash types.Hash) (models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionReceipt", ctx, hash)
	ret0, _ := ret[0].(models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionReceipt indicates an expected call of GetTransactionReceipt
func (mr *MockAPIsMockRecorder) GetTransactionReceipt(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionReceipt", reflect.TypeOf((*MockAPIs)(nil).GetTransactionReceipt), ctx, hash)
}

// GetTransactionLogs mocks base method
func (m *MockAPIs) GetTransactionLogs(ctx context.Context, hash types.Hash) ([]*statetype.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionLogs", ctx, hash)
	ret0, _ := ret[0].([]*statetype.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionLogs indicates an expected call of GetTransactionLogs
func (mr *MockAPIsMockRecorder) GetTransactionLogs(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionLogs", reflect.TypeOf((*MockAPIs)(nil).GetTransactionLogs), ctx, hash)
}

// BlockHeight mocks base method
func (m *MockAPIs) BlockHeight(ctx context.Context) (*hexutil.Uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockHeight", ctx)
	ret0, _ := ret[0].(*hexutil.Uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockHeight indicates an expected call of BlockHeight
func (mr *MockAPIsMockRecorder) BlockHeight(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockHeight", reflect.TypeOf((*MockAPIs)(nil).BlockHeight), ctx)
}

// GetBlockByHash mocks base method
func (m *MockAPIs) GetBlockByHash(ctx context.Context, blockHash types.Hash) (*models.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockByHash", ctx, blockHash)
	ret0, _ := ret[0].(*models.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByHash indicates an expected call of GetBlockByHash
func (mr *MockAPIsMockRecorder) GetBlockByHash(ctx, blockHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByHash", reflect.TypeOf((*MockAPIs)(nil).GetBlockByHash), ctx, blockHash)
}

// GetBlockByHeight mocks base method
func (m *MockAPIs) GetBlockByHeight(ctx context.Context, blockHeight hexutil.Uint64) (*models.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockByHeight", ctx, blockHeight)
	ret0, _ := ret[0].(*models.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByHeight indicates an expected call of GetBlockByHeight
func (mr *MockAPIsMockRecorder) GetBlockByHeight(ctx, blockHeight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByHeight", reflect.TypeOf((*MockAPIs)(nil).GetBlockByHeight), ctx, blockHeight)
}

// GetBlockTransactionCountByHash mocks base method
func (m *MockAPIs) GetBlockTransactionCountByHash(ctx context.Context, blockHash types.Hash) (*hexutil.Uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockTransactionCountByHash", ctx, blockHash)
	ret0, _ := ret[0].(*hexutil.Uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockTransactionCountByHash indicates an expected call of GetBlockTransactionCountByHash
func (mr *MockAPIsMockRecorder) GetBlockTransactionCountByHash(ctx, blockHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockTransactionCountByHash", reflect.TypeOf((*MockAPIs)(nil).GetBlockTransactionCountByHash), ctx, blockHash)
}

// GetBlockTransactionCountByHeight mocks base method
func (m *MockAPIs) GetBlockTransactionCountByHeight(ctx context.Context, blockHeight hexutil.Uint64) (*hexutil.Uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockTransactionCountByHeight", ctx, blockHeight)
	ret0, _ := ret[0].(*hexutil.Uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockTransactionCountByHeight indicates an expected call of GetBlockTransactionCountByHeight
func (mr *MockAPIsMockRecorder) GetBlockTransactionCountByHeight(ctx, blockHeight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockTransactionCountByHeight", reflect.TypeOf((*MockAPIs)(nil).GetBlockTransactionCountByHeight), ctx, blockHeight)
}

// GetTransactionByBlockHashAndIndex mocks base method
func (m *MockAPIs) GetTransactionByBlockHashAndIndex(ctx context.Context, blockHash types.Hash, txIndex hexutil.Uint64) (models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByBlockHashAndIndex", ctx, blockHash, txIndex)
	ret0, _ := ret[0].(models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionByBlockHashAndIndex indicates an expected call of GetTransactionByBlockHashAndIndex
func (mr *MockAPIsMockRecorder) GetTransactionByBlockHashAndIndex(ctx, blockHash, txIndex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByBlockHashAndIndex", reflect.TypeOf((*MockAPIs)(nil).GetTransactionByBlockHashAndIndex), ctx, blockHash, txIndex)
}

// GetTransactionByBlockHeightAndIndex mocks base method
func (m *MockAPIs) GetTransactionByBlockHeightAndIndex(ctx context.Context, blockHeight, txIndex hexutil.Uint64) (models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByBlockHeightAndIndex", ctx, blockHeight, txIndex)
	ret0, _ := ret[0].(models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionByBlockHeightAndIndex indicates an expected call of GetTransactionByBlockHeightAndIndex
func (mr *MockAPIsMockRecorder) GetTransactionByBlockHeightAndIndex(ctx, blockHeight, txIndex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByBlockHeightAndIndex", reflect.TypeOf((*MockAPIs)(nil).GetTransactionByBlockHeightAndIndex), ctx, blockHeight, txIndex)
}

// GetCode mocks base method
func (m *MockAPIs) GetCode(ctx context.Context, contract types.Address, blockHeight hexutil.Uint64) (*hexutil.Bytes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCode", ctx, contract, blockHeight)
	ret0, _ := ret[0].(*hexutil.Bytes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCode indicates an expected call of GetCode
func (mr *MockAPIsMockRecorder) GetCode(ctx, contract, blockHeight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCode", reflect.TypeOf((*MockAPIs)(nil).GetCode), ctx, contract, blockHeight)
}

// Call mocks base method
func (m *MockAPIs) Call(ctx context.Context, hash models.VmMessage) (*hexutil.Bytes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", ctx, hash)
	ret0, _ := ret[0].(*hexutil.Bytes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call
func (mr *MockAPIsMockRecorder) Call(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockAPIs)(nil).Call), ctx, hash)
}

// EstimateGas mocks base method
func (m *MockAPIs) EstimateGas(ctx context.Context, transaction models.Transaction) (*hexutil.Uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateGas", ctx, transaction)
	ret0, _ := ret[0].(*hexutil.Uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateGas indicates an expected call of EstimateGas
func (mr *MockAPIsMockRecorder) EstimateGas(ctx, transaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateGas", reflect.TypeOf((*MockAPIs)(nil).EstimateGas), ctx, transaction)
}

// CompileContract mocks base method
func (m *MockAPIs) CompileContract(ctx context.Context, compileType protocol.CompileType, contract hexutil.Bytes) (*hexutil.Bytes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompileContract", ctx, compileType, contract)
	ret0, _ := ret[0].(*hexutil.Bytes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompileContract indicates an expected call of CompileContract
func (mr *MockAPIsMockRecorder) CompileContract(ctx, compileType, contract interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompileContract", reflect.TypeOf((*MockAPIs)(nil).CompileContract), ctx, compileType, contract)
}

// FilterSubscribeHeaders mocks base method
func (m *MockAPIs) FilterSubscribeHeaders(ctx context.Context) (*models.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterSubscribeHeaders", ctx)
	ret0, _ := ret[0].(*models.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterSubscribeHeaders indicates an expected call of FilterSubscribeHeaders
func (mr *MockAPIsMockRecorder) FilterSubscribeHeaders(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterSubscribeHeaders", reflect.TypeOf((*MockAPIs)(nil).FilterSubscribeHeaders), ctx)
}

// APIs mocks base method
func (m *MockAPIs) APIs() []protocol.API {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIs")
	ret0, _ := ret[0].([]protocol.API)
	return ret0
}

// APIs indicates an expected call of APIs
func (mr *MockAPIsMockRecorder) APIs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIs", reflect.TypeOf((*MockAPIs)(nil).APIs))
}

// RegisterAPI mocks base method
func (m *MockAPIs) RegisterAPI(apis []protocol.API) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterAPI", apis)
}

// RegisterAPI indicates an expected call of RegisterAPI
func (mr *MockAPIsMockRecorder) RegisterAPI(apis interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAPI", reflect.TypeOf((*MockAPIs)(nil).RegisterAPI), apis)
}
