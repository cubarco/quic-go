// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go (interfaces: QUICConn)

// Package quic is a generated GoMock package.
package quic

import (
	context "context"
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	congestion "github.com/quic-go/quic-go/congestion"
	protocol "github.com/quic-go/quic-go/internal/protocol"
	qerr "github.com/quic-go/quic-go/internal/qerr"
)

// MockQUICConn is a mock of QUICConn interface.
type MockQUICConn struct {
	ctrl     *gomock.Controller
	recorder *MockQUICConnMockRecorder
}

// MockQUICConnMockRecorder is the mock recorder for MockQUICConn.
type MockQUICConnMockRecorder struct {
	mock *MockQUICConn
}

// NewMockQUICConn creates a new mock instance.
func NewMockQUICConn(ctrl *gomock.Controller) *MockQUICConn {
	mock := &MockQUICConn{ctrl: ctrl}
	mock.recorder = &MockQUICConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQUICConn) EXPECT() *MockQUICConnMockRecorder {
	return m.recorder
}

// AcceptStream mocks base method.
func (m *MockQUICConn) AcceptStream(arg0 context.Context) (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptStream", arg0)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptStream indicates an expected call of AcceptStream.
func (mr *MockQUICConnMockRecorder) AcceptStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptStream", reflect.TypeOf((*MockQUICConn)(nil).AcceptStream), arg0)
}

// AcceptUniStream mocks base method.
func (m *MockQUICConn) AcceptUniStream(arg0 context.Context) (ReceiveStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptUniStream", arg0)
	ret0, _ := ret[0].(ReceiveStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptUniStream indicates an expected call of AcceptUniStream.
func (mr *MockQUICConnMockRecorder) AcceptUniStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptUniStream", reflect.TypeOf((*MockQUICConn)(nil).AcceptUniStream), arg0)
}

// CloseWithError mocks base method.
func (m *MockQUICConn) CloseWithError(arg0 qerr.ApplicationErrorCode, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseWithError", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseWithError indicates an expected call of CloseWithError.
func (mr *MockQUICConnMockRecorder) CloseWithError(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseWithError", reflect.TypeOf((*MockQUICConn)(nil).CloseWithError), arg0, arg1)
}

// ConnectionState mocks base method.
func (m *MockQUICConn) ConnectionState() ConnectionState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectionState")
	ret0, _ := ret[0].(ConnectionState)
	return ret0
}

// ConnectionState indicates an expected call of ConnectionState.
func (mr *MockQUICConnMockRecorder) ConnectionState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectionState", reflect.TypeOf((*MockQUICConn)(nil).ConnectionState))
}

// Context mocks base method.
func (m *MockQUICConn) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockQUICConnMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockQUICConn)(nil).Context))
}

// GetVersion mocks base method.
func (m *MockQUICConn) GetVersion() protocol.VersionNumber {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion")
	ret0, _ := ret[0].(protocol.VersionNumber)
	return ret0
}

// GetVersion indicates an expected call of GetVersion.
func (mr *MockQUICConnMockRecorder) GetVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockQUICConn)(nil).GetVersion))
}

// HandshakeComplete mocks base method.
func (m *MockQUICConn) HandshakeComplete() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandshakeComplete")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// HandshakeComplete indicates an expected call of HandshakeComplete.
func (mr *MockQUICConnMockRecorder) HandshakeComplete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandshakeComplete", reflect.TypeOf((*MockQUICConn)(nil).HandshakeComplete))
}

// LocalAddr mocks base method.
func (m *MockQUICConn) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr.
func (mr *MockQUICConnMockRecorder) LocalAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockQUICConn)(nil).LocalAddr))
}

// NextConnection mocks base method.
func (m *MockQUICConn) NextConnection() Connection {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextConnection")
	ret0, _ := ret[0].(Connection)
	return ret0
}

// NextConnection indicates an expected call of NextConnection.
func (mr *MockQUICConnMockRecorder) NextConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextConnection", reflect.TypeOf((*MockQUICConn)(nil).NextConnection))
}

// OpenStream mocks base method.
func (m *MockQUICConn) OpenStream() (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStream")
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStream indicates an expected call of OpenStream.
func (mr *MockQUICConnMockRecorder) OpenStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStream", reflect.TypeOf((*MockQUICConn)(nil).OpenStream))
}

// OpenStreamSync mocks base method.
func (m *MockQUICConn) OpenStreamSync(arg0 context.Context) (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStreamSync", arg0)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStreamSync indicates an expected call of OpenStreamSync.
func (mr *MockQUICConnMockRecorder) OpenStreamSync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStreamSync", reflect.TypeOf((*MockQUICConn)(nil).OpenStreamSync), arg0)
}

// OpenUniStream mocks base method.
func (m *MockQUICConn) OpenUniStream() (SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStream")
	ret0, _ := ret[0].(SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStream indicates an expected call of OpenUniStream.
func (mr *MockQUICConnMockRecorder) OpenUniStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStream", reflect.TypeOf((*MockQUICConn)(nil).OpenUniStream))
}

// OpenUniStreamSync mocks base method.
func (m *MockQUICConn) OpenUniStreamSync(arg0 context.Context) (SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStreamSync", arg0)
	ret0, _ := ret[0].(SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStreamSync indicates an expected call of OpenUniStreamSync.
func (mr *MockQUICConnMockRecorder) OpenUniStreamSync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStreamSync", reflect.TypeOf((*MockQUICConn)(nil).OpenUniStreamSync), arg0)
}

// ReceiveMessage mocks base method.
func (m *MockQUICConn) ReceiveMessage() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveMessage")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveMessage indicates an expected call of ReceiveMessage.
func (mr *MockQUICConnMockRecorder) ReceiveMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveMessage", reflect.TypeOf((*MockQUICConn)(nil).ReceiveMessage))
}

// RemoteAddr mocks base method.
func (m *MockQUICConn) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr.
func (mr *MockQUICConnMockRecorder) RemoteAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockQUICConn)(nil).RemoteAddr))
}

// SendMessage mocks base method.
func (m *MockQUICConn) SendMessage(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockQUICConnMockRecorder) SendMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockQUICConn)(nil).SendMessage), arg0)
}

// SetCongestionControl mocks base method.
func (m *MockQUICConn) SetCongestionControl(arg0 congestion.CongestionControl) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCongestionControl", arg0)
}

// SetCongestionControl indicates an expected call of SetCongestionControl.
func (mr *MockQUICConnMockRecorder) SetCongestionControl(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCongestionControl", reflect.TypeOf((*MockQUICConn)(nil).SetCongestionControl), arg0)
}

// destroy mocks base method.
func (m *MockQUICConn) destroy(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "destroy", arg0)
}

// destroy indicates an expected call of destroy.
func (mr *MockQUICConnMockRecorder) destroy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "destroy", reflect.TypeOf((*MockQUICConn)(nil).destroy), arg0)
}

// earlyConnReady mocks base method.
func (m *MockQUICConn) earlyConnReady() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "earlyConnReady")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// earlyConnReady indicates an expected call of earlyConnReady.
func (mr *MockQUICConnMockRecorder) earlyConnReady() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "earlyConnReady", reflect.TypeOf((*MockQUICConn)(nil).earlyConnReady))
}

// getPerspective mocks base method.
func (m *MockQUICConn) getPerspective() protocol.Perspective {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getPerspective")
	ret0, _ := ret[0].(protocol.Perspective)
	return ret0
}

// getPerspective indicates an expected call of getPerspective.
func (mr *MockQUICConnMockRecorder) getPerspective() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getPerspective", reflect.TypeOf((*MockQUICConn)(nil).getPerspective))
}

// handlePacket mocks base method.
func (m *MockQUICConn) handlePacket(arg0 receivedPacket) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "handlePacket", arg0)
}

// handlePacket indicates an expected call of handlePacket.
func (mr *MockQUICConnMockRecorder) handlePacket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handlePacket", reflect.TypeOf((*MockQUICConn)(nil).handlePacket), arg0)
}

// run mocks base method.
func (m *MockQUICConn) run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "run")
	ret0, _ := ret[0].(error)
	return ret0
}

// run indicates an expected call of run.
func (mr *MockQUICConnMockRecorder) run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "run", reflect.TypeOf((*MockQUICConn)(nil).run))
}

// shutdown mocks base method.
func (m *MockQUICConn) shutdown() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "shutdown")
}

// shutdown indicates an expected call of shutdown.
func (mr *MockQUICConnMockRecorder) shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "shutdown", reflect.TypeOf((*MockQUICConn)(nil).shutdown))
}
