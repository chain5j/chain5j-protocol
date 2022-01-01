// Code generated by MockGen. DO NOT EDIT.
// Source: protocol/nodekey.go

// Package mock is a generated GoMock package.
package mock

import (
	crypto "crypto"
	signature "github.com/chain5j/chain5j-pkg/crypto/signature"
	models "github.com/chain5j/chain5j-protocol/models"
	protocol "github.com/chain5j/chain5j-protocol/protocol"
	gomock "github.com/golang/mock/gomock"
	hash "hash"
	reflect "reflect"
)

// MockNodeKey is a mock of NodeKey interface
type MockNodeKey struct {
	ctrl     *gomock.Controller
	recorder *MockNodeKeyMockRecorder
}

// MockNodeKeyMockRecorder is the mock recorder for MockNodeKey
type MockNodeKeyMockRecorder struct {
	mock *MockNodeKey
}

// NewMockNodeKey creates a new mock instance
func NewMockNodeKey(ctrl *gomock.Controller) *MockNodeKey {
	mock := &MockNodeKey{ctrl: ctrl}
	mock.recorder = &MockNodeKeyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeKey) EXPECT() *MockNodeKeyMockRecorder {
	return m.recorder
}

// ID mocks base method
func (m *MockNodeKey) ID() (models.NodeID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(models.NodeID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ID indicates an expected call of ID
func (mr *MockNodeKeyMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockNodeKey)(nil).ID))
}

// IdFromPub mocks base method
func (m *MockNodeKey) IdFromPub(pub crypto.PublicKey) (models.NodeID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IdFromPub", pub)
	ret0, _ := ret[0].(models.NodeID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IdFromPub indicates an expected call of IdFromPub
func (mr *MockNodeKeyMockRecorder) IdFromPub(pub interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IdFromPub", reflect.TypeOf((*MockNodeKey)(nil).IdFromPub), pub)
}

// PubKey mocks base method
func (m *MockNodeKey) PubKey(pubKey crypto.PublicKey) (protocol.PubKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PubKey", pubKey)
	ret0, _ := ret[0].(protocol.PubKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PubKey indicates an expected call of PubKey
func (mr *MockNodeKeyMockRecorder) PubKey(pubKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubKey", reflect.TypeOf((*MockNodeKey)(nil).PubKey), pubKey)
}

// Sign mocks base method
func (m *MockNodeKey) Sign(data []byte) (*signature.SignResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", data)
	ret0, _ := ret[0].(*signature.SignResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign
func (mr *MockNodeKeyMockRecorder) Sign(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockNodeKey)(nil).Sign), data)
}

// Verify mocks base method
func (m *MockNodeKey) Verify(data []byte, signResult *signature.SignResult) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", data, signResult)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Verify indicates an expected call of Verify
func (mr *MockNodeKeyMockRecorder) Verify(data, signResult interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockNodeKey)(nil).Verify), data, signResult)
}

// RecoverId mocks base method
func (m *MockNodeKey) RecoverId(data []byte, signResult *signature.SignResult) (models.NodeID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecoverId", data, signResult)
	ret0, _ := ret[0].(models.NodeID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecoverId indicates an expected call of RecoverId
func (mr *MockNodeKeyMockRecorder) RecoverId(data, signResult interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecoverId", reflect.TypeOf((*MockNodeKey)(nil).RecoverId), data, signResult)
}

// RecoverPub mocks base method
func (m *MockNodeKey) RecoverPub(data []byte, signResult *signature.SignResult) (protocol.PubKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecoverPub", data, signResult)
	ret0, _ := ret[0].(protocol.PubKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecoverPub indicates an expected call of RecoverPub
func (mr *MockNodeKeyMockRecorder) RecoverPub(data, signResult interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecoverPub", reflect.TypeOf((*MockNodeKey)(nil).RecoverPub), data, signResult)
}

// MockKey is a mock of Key interface
type MockKey struct {
	ctrl     *gomock.Controller
	recorder *MockKeyMockRecorder
}

// MockKeyMockRecorder is the mock recorder for MockKey
type MockKeyMockRecorder struct {
	mock *MockKey
}

// NewMockKey creates a new mock instance
func NewMockKey(ctrl *gomock.Controller) *MockKey {
	mock := &MockKey{ctrl: ctrl}
	mock.recorder = &MockKeyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKey) EXPECT() *MockKeyMockRecorder {
	return m.recorder
}

// Marshal mocks base method
func (m *MockKey) Marshal() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Marshal")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Marshal indicates an expected call of Marshal
func (mr *MockKeyMockRecorder) Marshal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marshal", reflect.TypeOf((*MockKey)(nil).Marshal))
}

// Unmarshal mocks base method
func (m *MockKey) Unmarshal(input []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmarshal", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmarshal indicates an expected call of Unmarshal
func (mr *MockKeyMockRecorder) Unmarshal(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmarshal", reflect.TypeOf((*MockKey)(nil).Unmarshal), input)
}

// Equals mocks base method
func (m *MockKey) Equals(arg0 protocol.Key) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equals", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equals indicates an expected call of Equals
func (mr *MockKeyMockRecorder) Equals(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equals", reflect.TypeOf((*MockKey)(nil).Equals), arg0)
}

// Raw mocks base method
func (m *MockKey) Raw() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Raw")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Raw indicates an expected call of Raw
func (mr *MockKeyMockRecorder) Raw() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*MockKey)(nil).Raw))
}

// Type mocks base method
func (m *MockKey) Type() protocol.KeyType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(protocol.KeyType)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockKeyMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockKey)(nil).Type))
}

// Hash mocks base method
func (m *MockKey) Hash() func() hash.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash")
	ret0, _ := ret[0].(func() hash.Hash)
	return ret0
}

// Hash indicates an expected call of Hash
func (mr *MockKeyMockRecorder) Hash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockKey)(nil).Hash))
}

// MockPrivKey is a mock of PrivKey interface
type MockPrivKey struct {
	ctrl     *gomock.Controller
	recorder *MockPrivKeyMockRecorder
}

// MockPrivKeyMockRecorder is the mock recorder for MockPrivKey
type MockPrivKeyMockRecorder struct {
	mock *MockPrivKey
}

// NewMockPrivKey creates a new mock instance
func NewMockPrivKey(ctrl *gomock.Controller) *MockPrivKey {
	mock := &MockPrivKey{ctrl: ctrl}
	mock.recorder = &MockPrivKeyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPrivKey) EXPECT() *MockPrivKeyMockRecorder {
	return m.recorder
}

// Marshal mocks base method
func (m *MockPrivKey) Marshal() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Marshal")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Marshal indicates an expected call of Marshal
func (mr *MockPrivKeyMockRecorder) Marshal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marshal", reflect.TypeOf((*MockPrivKey)(nil).Marshal))
}

// Unmarshal mocks base method
func (m *MockPrivKey) Unmarshal(input []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmarshal", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmarshal indicates an expected call of Unmarshal
func (mr *MockPrivKeyMockRecorder) Unmarshal(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmarshal", reflect.TypeOf((*MockPrivKey)(nil).Unmarshal), input)
}

// Equals mocks base method
func (m *MockPrivKey) Equals(arg0 protocol.Key) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equals", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equals indicates an expected call of Equals
func (mr *MockPrivKeyMockRecorder) Equals(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equals", reflect.TypeOf((*MockPrivKey)(nil).Equals), arg0)
}

// Raw mocks base method
func (m *MockPrivKey) Raw() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Raw")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Raw indicates an expected call of Raw
func (mr *MockPrivKeyMockRecorder) Raw() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*MockPrivKey)(nil).Raw))
}

// Type mocks base method
func (m *MockPrivKey) Type() protocol.KeyType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(protocol.KeyType)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockPrivKeyMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockPrivKey)(nil).Type))
}

// Hash mocks base method
func (m *MockPrivKey) Hash() func() hash.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash")
	ret0, _ := ret[0].(func() hash.Hash)
	return ret0
}

// Hash indicates an expected call of Hash
func (mr *MockPrivKeyMockRecorder) Hash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockPrivKey)(nil).Hash))
}

// Sign mocks base method
func (m *MockPrivKey) Sign(data []byte) (*signature.SignResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", data)
	ret0, _ := ret[0].(*signature.SignResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign
func (mr *MockPrivKeyMockRecorder) Sign(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockPrivKey)(nil).Sign), data)
}

// GetPublic mocks base method
func (m *MockPrivKey) GetPublic() protocol.PubKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublic")
	ret0, _ := ret[0].(protocol.PubKey)
	return ret0
}

// GetPublic indicates an expected call of GetPublic
func (mr *MockPrivKeyMockRecorder) GetPublic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublic", reflect.TypeOf((*MockPrivKey)(nil).GetPublic))
}

// MockPubKey is a mock of PubKey interface
type MockPubKey struct {
	ctrl     *gomock.Controller
	recorder *MockPubKeyMockRecorder
}

// MockPubKeyMockRecorder is the mock recorder for MockPubKey
type MockPubKeyMockRecorder struct {
	mock *MockPubKey
}

// NewMockPubKey creates a new mock instance
func NewMockPubKey(ctrl *gomock.Controller) *MockPubKey {
	mock := &MockPubKey{ctrl: ctrl}
	mock.recorder = &MockPubKeyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPubKey) EXPECT() *MockPubKeyMockRecorder {
	return m.recorder
}

// Marshal mocks base method
func (m *MockPubKey) Marshal() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Marshal")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Marshal indicates an expected call of Marshal
func (mr *MockPubKeyMockRecorder) Marshal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marshal", reflect.TypeOf((*MockPubKey)(nil).Marshal))
}

// Unmarshal mocks base method
func (m *MockPubKey) Unmarshal(input []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmarshal", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmarshal indicates an expected call of Unmarshal
func (mr *MockPubKeyMockRecorder) Unmarshal(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmarshal", reflect.TypeOf((*MockPubKey)(nil).Unmarshal), input)
}

// Equals mocks base method
func (m *MockPubKey) Equals(arg0 protocol.Key) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equals", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equals indicates an expected call of Equals
func (mr *MockPubKeyMockRecorder) Equals(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equals", reflect.TypeOf((*MockPubKey)(nil).Equals), arg0)
}

// Raw mocks base method
func (m *MockPubKey) Raw() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Raw")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Raw indicates an expected call of Raw
func (mr *MockPubKeyMockRecorder) Raw() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*MockPubKey)(nil).Raw))
}

// Type mocks base method
func (m *MockPubKey) Type() protocol.KeyType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(protocol.KeyType)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockPubKeyMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockPubKey)(nil).Type))
}

// Hash mocks base method
func (m *MockPubKey) Hash() func() hash.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash")
	ret0, _ := ret[0].(func() hash.Hash)
	return ret0
}

// Hash indicates an expected call of Hash
func (mr *MockPubKeyMockRecorder) Hash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockPubKey)(nil).Hash))
}

// Verify mocks base method
func (m *MockPubKey) Verify(data []byte, signResult *signature.SignResult) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", data, signResult)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Verify indicates an expected call of Verify
func (mr *MockPubKeyMockRecorder) Verify(data, signResult interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockPubKey)(nil).Verify), data, signResult)
}