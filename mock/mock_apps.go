// Code generated by MockGen. DO NOT EDIT.
// Source: protocol/apps.go

// Package mock is a generated GoMock package.
package mock

import (
	types "github.com/chain5j/chain5j-pkg/types"
	models "github.com/chain5j/chain5j-protocol/models"
	protocol "github.com/chain5j/chain5j-protocol/protocol"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAppContext is a mock of AppContext interface
type MockAppContext struct {
	ctrl     *gomock.Controller
	recorder *MockAppContextMockRecorder
}

// MockAppContextMockRecorder is the mock recorder for MockAppContext
type MockAppContextMockRecorder struct {
	mock *MockAppContext
}

// NewMockAppContext creates a new mock instance
func NewMockAppContext(ctrl *gomock.Controller) *MockAppContext {
	mock := &MockAppContext{ctrl: ctrl}
	mock.recorder = &MockAppContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppContext) EXPECT() *MockAppContextMockRecorder {
	return m.recorder
}

// Caller mocks base method
func (m *MockAppContext) Caller() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Caller")
	ret0, _ := ret[0].(string)
	return ret0
}

// Caller indicates an expected call of Caller
func (mr *MockAppContextMockRecorder) Caller() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Caller", reflect.TypeOf((*MockAppContext)(nil).Caller))
}

// App mocks base method
func (m *MockAppContext) App() protocol.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "App")
	ret0, _ := ret[0].(protocol.Application)
	return ret0
}

// App indicates an expected call of App
func (mr *MockAppContextMockRecorder) App() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "App", reflect.TypeOf((*MockAppContext)(nil).App))
}

// MockAppContexts is a mock of AppContexts interface
type MockAppContexts struct {
	ctrl     *gomock.Controller
	recorder *MockAppContextsMockRecorder
}

// MockAppContextsMockRecorder is the mock recorder for MockAppContexts
type MockAppContextsMockRecorder struct {
	mock *MockAppContexts
}

// NewMockAppContexts creates a new mock instance
func NewMockAppContexts(ctrl *gomock.Controller) *MockAppContexts {
	mock := &MockAppContexts{ctrl: ctrl}
	mock.recorder = &MockAppContextsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppContexts) EXPECT() *MockAppContextsMockRecorder {
	return m.recorder
}

// Ctx mocks base method
func (m *MockAppContexts) Ctx(t types.TxType) protocol.AppContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ctx", t)
	ret0, _ := ret[0].(protocol.AppContext)
	return ret0
}

// Ctx indicates an expected call of Ctx
func (mr *MockAppContextsMockRecorder) Ctx(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ctx", reflect.TypeOf((*MockAppContexts)(nil).Ctx), t)
}

// MockTxValidator is a mock of TxValidator interface
type MockTxValidator struct {
	ctrl     *gomock.Controller
	recorder *MockTxValidatorMockRecorder
}

// MockTxValidatorMockRecorder is the mock recorder for MockTxValidator
type MockTxValidatorMockRecorder struct {
	mock *MockTxValidator
}

// NewMockTxValidator creates a new mock instance
func NewMockTxValidator(ctrl *gomock.Controller) *MockTxValidator {
	mock := &MockTxValidator{ctrl: ctrl}
	mock.recorder = &MockTxValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTxValidator) EXPECT() *MockTxValidatorMockRecorder {
	return m.recorder
}

// ValidateTx mocks base method
func (m *MockTxValidator) ValidateTx(ctx protocol.AppContext, txI models.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateTx", ctx, txI)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateTx indicates an expected call of ValidateTx
func (mr *MockTxValidatorMockRecorder) ValidateTx(ctx, txI interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateTx", reflect.TypeOf((*MockTxValidator)(nil).ValidateTx), ctx, txI)
}

// ValidateTxSafe mocks base method
func (m *MockTxValidator) ValidateTxSafe(ctx protocol.AppContext, txI models.Transaction, headerTimestamp uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateTxSafe", ctx, txI, headerTimestamp)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateTxSafe indicates an expected call of ValidateTxSafe
func (mr *MockTxValidatorMockRecorder) ValidateTxSafe(ctx, txI, headerTimestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateTxSafe", reflect.TypeOf((*MockTxValidator)(nil).ValidateTxSafe), ctx, txI, headerTimestamp)
}

// MockApplication is a mock of Application interface
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockApplication) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockApplicationMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockApplication)(nil).Start))
}

// Stop mocks base method
func (m *MockApplication) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockApplicationMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockApplication)(nil).Stop))
}

// NewAppContexts mocks base method
func (m *MockApplication) NewAppContexts(module string, args ...interface{}) (protocol.AppContext, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{module}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewAppContexts", varargs...)
	ret0, _ := ret[0].(protocol.AppContext)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewAppContexts indicates an expected call of NewAppContexts
func (mr *MockApplicationMockRecorder) NewAppContexts(module interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{module}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAppContexts", reflect.TypeOf((*MockApplication)(nil).NewAppContexts), varargs...)
}

// TxPool mocks base method
func (m *MockApplication) TxPool(config protocol.Config, apps protocol.Apps, blockReader protocol.BlockReader, broadcaster protocol.Broadcaster) (protocol.TxPool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxPool", config, apps, blockReader, broadcaster)
	ret0, _ := ret[0].(protocol.TxPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TxPool indicates an expected call of TxPool
func (mr *MockApplicationMockRecorder) TxPool(config, apps, blockReader, broadcaster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxPool", reflect.TypeOf((*MockApplication)(nil).TxPool), config, apps, blockReader, broadcaster)
}

// ValidateTx mocks base method
func (m *MockApplication) ValidateTx(ctx protocol.AppContext, txI models.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateTx", ctx, txI)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateTx indicates an expected call of ValidateTx
func (mr *MockApplicationMockRecorder) ValidateTx(ctx, txI interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateTx", reflect.TypeOf((*MockApplication)(nil).ValidateTx), ctx, txI)
}

// ValidateTxSafe mocks base method
func (m *MockApplication) ValidateTxSafe(ctx protocol.AppContext, txI models.Transaction, headerTimestamp uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateTxSafe", ctx, txI, headerTimestamp)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateTxSafe indicates an expected call of ValidateTxSafe
func (mr *MockApplicationMockRecorder) ValidateTxSafe(ctx, txI, headerTimestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateTxSafe", reflect.TypeOf((*MockApplication)(nil).ValidateTxSafe), ctx, txI, headerTimestamp)
}

// DeleteErrTx mocks base method
func (m *MockApplication) DeleteErrTx(txI models.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteErrTx", txI)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteErrTx indicates an expected call of DeleteErrTx
func (mr *MockApplicationMockRecorder) DeleteErrTx(txI interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteErrTx", reflect.TypeOf((*MockApplication)(nil).DeleteErrTx), txI)
}

// DeleteOkTx mocks base method
func (m *MockApplication) DeleteOkTx(txI models.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOkTx", txI)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOkTx indicates an expected call of DeleteOkTx
func (mr *MockApplicationMockRecorder) DeleteOkTx(txI interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOkTx", reflect.TypeOf((*MockApplication)(nil).DeleteOkTx), txI)
}

// GetCacheNonce mocks base method
func (m *MockApplication) GetCacheNonce(ctx protocol.AppContext, account string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCacheNonce", ctx, account)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCacheNonce indicates an expected call of GetCacheNonce
func (mr *MockApplicationMockRecorder) GetCacheNonce(ctx, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCacheNonce", reflect.TypeOf((*MockApplication)(nil).GetCacheNonce), ctx, account)
}

// Prepare mocks base method
func (m *MockApplication) Prepare(ctx protocol.AppContext, header *models.Header, txs models.TransactionSortedList, totalGas uint64) *models.TxsStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", ctx, header, txs, totalGas)
	ret0, _ := ret[0].(*models.TxsStatus)
	return ret0
}

// Prepare indicates an expected call of Prepare
func (mr *MockApplicationMockRecorder) Prepare(ctx, header, txs, totalGas interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockApplication)(nil).Prepare), ctx, header, txs, totalGas)
}

// Commit mocks base method
func (m *MockApplication) Commit(ctx protocol.AppContext, header *models.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", ctx, header)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockApplicationMockRecorder) Commit(ctx, header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockApplication)(nil).Commit), ctx, header)
}

// MockApps is a mock of Apps interface
type MockApps struct {
	ctrl     *gomock.Controller
	recorder *MockAppsMockRecorder
}

// MockAppsMockRecorder is the mock recorder for MockApps
type MockAppsMockRecorder struct {
	mock *MockApps
}

// NewMockApps creates a new mock instance
func NewMockApps(ctrl *gomock.Controller) *MockApps {
	mock := &MockApps{ctrl: ctrl}
	mock.recorder = &MockAppsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApps) EXPECT() *MockAppsMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockApps) Register(txType types.TxType, app protocol.Application) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Register", txType, app)
}

// Register indicates an expected call of Register
func (mr *MockAppsMockRecorder) Register(txType, app interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockApps)(nil).Register), txType, app)
}

// App mocks base method
func (m *MockApps) App(txType types.TxType) (protocol.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "App", txType)
	ret0, _ := ret[0].(protocol.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// App indicates an expected call of App
func (mr *MockAppsMockRecorder) App(txType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "App", reflect.TypeOf((*MockApps)(nil).App), txType)
}

// NewAppContexts mocks base method
func (m *MockApps) NewAppContexts(module string, preRoot []byte) (protocol.AppContexts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAppContexts", module, preRoot)
	ret0, _ := ret[0].(protocol.AppContexts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewAppContexts indicates an expected call of NewAppContexts
func (mr *MockAppsMockRecorder) NewAppContexts(module, preRoot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAppContexts", reflect.TypeOf((*MockApps)(nil).NewAppContexts), module, preRoot)
}

// Prepare mocks base method
func (m *MockApps) Prepare(ctx protocol.AppContexts, preRoot []byte, header *models.Header, txs models.Transactions, totalGas uint64) ([]byte, models.AppsStatus) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", ctx, preRoot, header, txs, totalGas)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(models.AppsStatus)
	return ret0, ret1
}

// Prepare indicates an expected call of Prepare
func (mr *MockAppsMockRecorder) Prepare(ctx, preRoot, header, txs, totalGas interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockApps)(nil).Prepare), ctx, preRoot, header, txs, totalGas)
}

// Commit mocks base method
func (m *MockApps) Commit(ctx protocol.AppContexts, header *models.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", ctx, header)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockAppsMockRecorder) Commit(ctx, header interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockApps)(nil).Commit), ctx, header)
}
