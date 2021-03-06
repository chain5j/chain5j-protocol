// Code generated by MockGen. DO NOT EDIT.
// Source: protocol/consensus.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	models "github.com/chain5j/chain5j-protocol/models"
	protocol "github.com/chain5j/chain5j-protocol/protocol"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockConsensus is a mock of Consensus interface
type MockConsensus struct {
	ctrl     *gomock.Controller
	recorder *MockConsensusMockRecorder
}

// MockConsensusMockRecorder is the mock recorder for MockConsensus
type MockConsensusMockRecorder struct {
	mock *MockConsensus
}

// NewMockConsensus creates a new mock instance
func NewMockConsensus(ctrl *gomock.Controller) *MockConsensus {
	mock := &MockConsensus{ctrl: ctrl}
	mock.recorder = &MockConsensusMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConsensus) EXPECT() *MockConsensusMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockConsensus) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockConsensusMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockConsensus)(nil).Start))
}

// Stop mocks base method
func (m *MockConsensus) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockConsensusMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockConsensus)(nil).Stop))
}

// Begin mocks base method
func (m *MockConsensus) Begin() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(error)
	return ret0
}

// Begin indicates an expected call of Begin
func (mr *MockConsensusMockRecorder) Begin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockConsensus)(nil).Begin))
}

// VerifyHeader mocks base method
func (m *MockConsensus) VerifyHeader(blockReader protocol.BlockReader, header *models.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyHeader", blockReader, header)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyHeader indicates an expected call of VerifyHeader
func (mr *MockConsensusMockRecorder) VerifyHeader(blockReader, header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyHeader", reflect.TypeOf((*MockConsensus)(nil).VerifyHeader), blockReader, header)
}

// VerifyHeaders mocks base method
func (m *MockConsensus) VerifyHeaders(blockReader protocol.BlockReader, headers []*models.Header, seals []bool) (chan<- struct{}, <-chan error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyHeaders", blockReader, headers, seals)
	ret0, _ := ret[0].(chan<- struct{})
	ret1, _ := ret[1].(<-chan error)
	return ret0, ret1
}

// VerifyHeaders indicates an expected call of VerifyHeaders
func (mr *MockConsensusMockRecorder) VerifyHeaders(blockReader, headers, seals interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyHeaders", reflect.TypeOf((*MockConsensus)(nil).VerifyHeaders), blockReader, headers, seals)
}

// Prepare mocks base method
func (m *MockConsensus) Prepare(blockReader protocol.BlockReader, header *models.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", blockReader, header)
	ret0, _ := ret[0].(error)
	return ret0
}

// Prepare indicates an expected call of Prepare
func (mr *MockConsensusMockRecorder) Prepare(blockReader, header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockConsensus)(nil).Prepare), blockReader, header)
}

// Finalize mocks base method
func (m *MockConsensus) Finalize(blockReader protocol.BlockReader, header *models.Header, txs models.Transactions) (*models.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Finalize", blockReader, header, txs)
	ret0, _ := ret[0].(*models.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Finalize indicates an expected call of Finalize
func (mr *MockConsensusMockRecorder) Finalize(blockReader, header, txs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finalize", reflect.TypeOf((*MockConsensus)(nil).Finalize), blockReader, header, txs)
}

// Seal mocks base method
func (m *MockConsensus) Seal(ctx context.Context, blockReader protocol.BlockReader, block *models.Block, results chan<- *models.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seal", ctx, blockReader, block, results)
	ret0, _ := ret[0].(error)
	return ret0
}

// Seal indicates an expected call of Seal
func (mr *MockConsensusMockRecorder) Seal(ctx, blockReader, block, results interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seal", reflect.TypeOf((*MockConsensus)(nil).Seal), ctx, blockReader, block, results)
}
