// Code generated by MockGen. DO NOT EDIT.
// Source: protocol/broadcaster.go

// Package mock is a generated GoMock package.
package mock

import (
	event "github.com/chain5j/chain5j-pkg/event"
	models "github.com/chain5j/chain5j-protocol/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBroadcaster is a mock of Broadcaster interface
type MockBroadcaster struct {
	ctrl     *gomock.Controller
	recorder *MockBroadcasterMockRecorder
}

// MockBroadcasterMockRecorder is the mock recorder for MockBroadcaster
type MockBroadcasterMockRecorder struct {
	mock *MockBroadcaster
}

// NewMockBroadcaster creates a new mock instance
func NewMockBroadcaster(ctrl *gomock.Controller) *MockBroadcaster {
	mock := &MockBroadcaster{ctrl: ctrl}
	mock.recorder = &MockBroadcasterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBroadcaster) EXPECT() *MockBroadcasterMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockBroadcaster) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockBroadcasterMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockBroadcaster)(nil).Start))
}

// Stop mocks base method
func (m *MockBroadcaster) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockBroadcasterMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockBroadcaster)(nil).Stop))
}

// SubscribeMsg mocks base method
func (m *MockBroadcaster) SubscribeMsg(msgType uint, ch chan<- *models.P2PMessage) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeMsg", msgType, ch)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeMsg indicates an expected call of SubscribeMsg
func (mr *MockBroadcasterMockRecorder) SubscribeMsg(msgType, ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeMsg", reflect.TypeOf((*MockBroadcaster)(nil).SubscribeMsg), msgType, ch)
}

// SubscribeNewPeer mocks base method
func (m *MockBroadcaster) SubscribeNewPeer(newPeerCh chan models.P2PID) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeNewPeer", newPeerCh)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeNewPeer indicates an expected call of SubscribeNewPeer
func (mr *MockBroadcasterMockRecorder) SubscribeNewPeer(newPeerCh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeNewPeer", reflect.TypeOf((*MockBroadcaster)(nil).SubscribeNewPeer), newPeerCh)
}

// SubscribeDropPeer mocks base method
func (m *MockBroadcaster) SubscribeDropPeer(dropPeerCh chan models.P2PID) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeDropPeer", dropPeerCh)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeDropPeer indicates an expected call of SubscribeDropPeer
func (mr *MockBroadcasterMockRecorder) SubscribeDropPeer(dropPeerCh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeDropPeer", reflect.TypeOf((*MockBroadcaster)(nil).SubscribeDropPeer), dropPeerCh)
}

// Broadcast mocks base method
func (m *MockBroadcaster) Broadcast(peers []models.P2PID, mType uint, payload []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Broadcast", peers, mType, payload)
	ret0, _ := ret[0].(error)
	return ret0
}

// Broadcast indicates an expected call of Broadcast
func (mr *MockBroadcasterMockRecorder) Broadcast(peers, mType, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockBroadcaster)(nil).Broadcast), peers, mType, payload)
}

// BroadcastTxs mocks base method
func (m *MockBroadcaster) BroadcastTxs(peerId *models.P2PID, txs []models.Transaction, isForce bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BroadcastTxs", peerId, txs, isForce)
}

// BroadcastTxs indicates an expected call of BroadcastTxs
func (mr *MockBroadcasterMockRecorder) BroadcastTxs(peerId, txs, isForce interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastTxs", reflect.TypeOf((*MockBroadcaster)(nil).BroadcastTxs), peerId, txs, isForce)
}

// RegisterTrustPeer mocks base method
func (m *MockBroadcaster) RegisterTrustPeer(peerID models.P2PID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterTrustPeer", peerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterTrustPeer indicates an expected call of RegisterTrustPeer
func (mr *MockBroadcasterMockRecorder) RegisterTrustPeer(peerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterTrustPeer", reflect.TypeOf((*MockBroadcaster)(nil).RegisterTrustPeer), peerID)
}

// DeregisterTrustPeer mocks base method
func (m *MockBroadcaster) DeregisterTrustPeer(peerID models.P2PID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterTrustPeer", peerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeregisterTrustPeer indicates an expected call of DeregisterTrustPeer
func (mr *MockBroadcasterMockRecorder) DeregisterTrustPeer(peerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterTrustPeer", reflect.TypeOf((*MockBroadcaster)(nil).DeregisterTrustPeer), peerID)
}
