// Code generated by MockGen. DO NOT EDIT.
// Source: db/types.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	db "github.com/weijun-sh/cosmos-sdk/db"
	gomock "github.com/golang/mock/gomock"
)

// MockDBConnection is a mock of DBConnection interface.
type MockDBConnection struct {
	ctrl     *gomock.Controller
	recorder *MockDBConnectionMockRecorder
}

// MockDBConnectionMockRecorder is the mock recorder for MockDBConnection.
type MockDBConnectionMockRecorder struct {
	mock *MockDBConnection
}

// NewMockDBConnection creates a new mock instance.
func NewMockDBConnection(ctrl *gomock.Controller) *MockDBConnection {
	mock := &MockDBConnection{ctrl: ctrl}
	mock.recorder = &MockDBConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBConnection) EXPECT() *MockDBConnectionMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockDBConnection) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDBConnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDBConnection)(nil).Close))
}

// DeleteVersion mocks base method.
func (m *MockDBConnection) DeleteVersion(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVersion", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVersion indicates an expected call of DeleteVersion.
func (mr *MockDBConnectionMockRecorder) DeleteVersion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVersion", reflect.TypeOf((*MockDBConnection)(nil).DeleteVersion), arg0)
}

// ReadWriter mocks base method.
func (m *MockDBConnection) ReadWriter() db.DBReadWriter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadWriter")
	ret0, _ := ret[0].(db.DBReadWriter)
	return ret0
}

// ReadWriter indicates an expected call of ReadWriter.
func (mr *MockDBConnectionMockRecorder) ReadWriter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadWriter", reflect.TypeOf((*MockDBConnection)(nil).ReadWriter))
}

// Reader mocks base method.
func (m *MockDBConnection) Reader() db.DBReader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reader")
	ret0, _ := ret[0].(db.DBReader)
	return ret0
}

// Reader indicates an expected call of Reader.
func (mr *MockDBConnectionMockRecorder) Reader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reader", reflect.TypeOf((*MockDBConnection)(nil).Reader))
}

// ReaderAt mocks base method.
func (m *MockDBConnection) ReaderAt(arg0 uint64) (db.DBReader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReaderAt", arg0)
	ret0, _ := ret[0].(db.DBReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReaderAt indicates an expected call of ReaderAt.
func (mr *MockDBConnectionMockRecorder) ReaderAt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReaderAt", reflect.TypeOf((*MockDBConnection)(nil).ReaderAt), arg0)
}

// Revert mocks base method.
func (m *MockDBConnection) Revert() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revert")
	ret0, _ := ret[0].(error)
	return ret0
}

// Revert indicates an expected call of Revert.
func (mr *MockDBConnectionMockRecorder) Revert() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revert", reflect.TypeOf((*MockDBConnection)(nil).Revert))
}

// SaveNextVersion mocks base method.
func (m *MockDBConnection) SaveNextVersion() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveNextVersion")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveNextVersion indicates an expected call of SaveNextVersion.
func (mr *MockDBConnectionMockRecorder) SaveNextVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveNextVersion", reflect.TypeOf((*MockDBConnection)(nil).SaveNextVersion))
}

// SaveVersion mocks base method.
func (m *MockDBConnection) SaveVersion(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveVersion", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveVersion indicates an expected call of SaveVersion.
func (mr *MockDBConnectionMockRecorder) SaveVersion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveVersion", reflect.TypeOf((*MockDBConnection)(nil).SaveVersion), arg0)
}

// Versions mocks base method.
func (m *MockDBConnection) Versions() (db.VersionSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Versions")
	ret0, _ := ret[0].(db.VersionSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Versions indicates an expected call of Versions.
func (mr *MockDBConnectionMockRecorder) Versions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Versions", reflect.TypeOf((*MockDBConnection)(nil).Versions))
}

// Writer mocks base method.
func (m *MockDBConnection) Writer() db.DBWriter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Writer")
	ret0, _ := ret[0].(db.DBWriter)
	return ret0
}

// Writer indicates an expected call of Writer.
func (mr *MockDBConnectionMockRecorder) Writer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Writer", reflect.TypeOf((*MockDBConnection)(nil).Writer))
}

// MockDBReader is a mock of DBReader interface.
type MockDBReader struct {
	ctrl     *gomock.Controller
	recorder *MockDBReaderMockRecorder
}

// MockDBReaderMockRecorder is the mock recorder for MockDBReader.
type MockDBReaderMockRecorder struct {
	mock *MockDBReader
}

// NewMockDBReader creates a new mock instance.
func NewMockDBReader(ctrl *gomock.Controller) *MockDBReader {
	mock := &MockDBReader{ctrl: ctrl}
	mock.recorder = &MockDBReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBReader) EXPECT() *MockDBReaderMockRecorder {
	return m.recorder
}

// Discard mocks base method.
func (m *MockDBReader) Discard() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Discard")
	ret0, _ := ret[0].(error)
	return ret0
}

// Discard indicates an expected call of Discard.
func (mr *MockDBReaderMockRecorder) Discard() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discard", reflect.TypeOf((*MockDBReader)(nil).Discard))
}

// Get mocks base method.
func (m *MockDBReader) Get(arg0 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDBReaderMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDBReader)(nil).Get), arg0)
}

// Has mocks base method.
func (m *MockDBReader) Has(key []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", key)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Has indicates an expected call of Has.
func (mr *MockDBReaderMockRecorder) Has(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockDBReader)(nil).Has), key)
}

// Iterator mocks base method.
func (m *MockDBReader) Iterator(start, end []byte) (db.Iterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iterator", start, end)
	ret0, _ := ret[0].(db.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Iterator indicates an expected call of Iterator.
func (mr *MockDBReaderMockRecorder) Iterator(start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterator", reflect.TypeOf((*MockDBReader)(nil).Iterator), start, end)
}

// ReverseIterator mocks base method.
func (m *MockDBReader) ReverseIterator(start, end []byte) (db.Iterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReverseIterator", start, end)
	ret0, _ := ret[0].(db.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReverseIterator indicates an expected call of ReverseIterator.
func (mr *MockDBReaderMockRecorder) ReverseIterator(start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReverseIterator", reflect.TypeOf((*MockDBReader)(nil).ReverseIterator), start, end)
}

// MockDBWriter is a mock of DBWriter interface.
type MockDBWriter struct {
	ctrl     *gomock.Controller
	recorder *MockDBWriterMockRecorder
}

// MockDBWriterMockRecorder is the mock recorder for MockDBWriter.
type MockDBWriterMockRecorder struct {
	mock *MockDBWriter
}

// NewMockDBWriter creates a new mock instance.
func NewMockDBWriter(ctrl *gomock.Controller) *MockDBWriter {
	mock := &MockDBWriter{ctrl: ctrl}
	mock.recorder = &MockDBWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBWriter) EXPECT() *MockDBWriterMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockDBWriter) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockDBWriterMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockDBWriter)(nil).Commit))
}

// Delete mocks base method.
func (m *MockDBWriter) Delete(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDBWriterMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDBWriter)(nil).Delete), arg0)
}

// Discard mocks base method.
func (m *MockDBWriter) Discard() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Discard")
	ret0, _ := ret[0].(error)
	return ret0
}

// Discard indicates an expected call of Discard.
func (mr *MockDBWriterMockRecorder) Discard() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discard", reflect.TypeOf((*MockDBWriter)(nil).Discard))
}

// Set mocks base method.
func (m *MockDBWriter) Set(arg0, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockDBWriterMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockDBWriter)(nil).Set), arg0, arg1)
}

// MockDBReadWriter is a mock of DBReadWriter interface.
type MockDBReadWriter struct {
	ctrl     *gomock.Controller
	recorder *MockDBReadWriterMockRecorder
}

// MockDBReadWriterMockRecorder is the mock recorder for MockDBReadWriter.
type MockDBReadWriterMockRecorder struct {
	mock *MockDBReadWriter
}

// NewMockDBReadWriter creates a new mock instance.
func NewMockDBReadWriter(ctrl *gomock.Controller) *MockDBReadWriter {
	mock := &MockDBReadWriter{ctrl: ctrl}
	mock.recorder = &MockDBReadWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBReadWriter) EXPECT() *MockDBReadWriterMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockDBReadWriter) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockDBReadWriterMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockDBReadWriter)(nil).Commit))
}

// Delete mocks base method.
func (m *MockDBReadWriter) Delete(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDBReadWriterMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDBReadWriter)(nil).Delete), arg0)
}

// Discard mocks base method.
func (m *MockDBReadWriter) Discard() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Discard")
	ret0, _ := ret[0].(error)
	return ret0
}

// Discard indicates an expected call of Discard.
func (mr *MockDBReadWriterMockRecorder) Discard() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discard", reflect.TypeOf((*MockDBReadWriter)(nil).Discard))
}

// Get mocks base method.
func (m *MockDBReadWriter) Get(arg0 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDBReadWriterMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDBReadWriter)(nil).Get), arg0)
}

// Has mocks base method.
func (m *MockDBReadWriter) Has(key []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", key)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Has indicates an expected call of Has.
func (mr *MockDBReadWriterMockRecorder) Has(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockDBReadWriter)(nil).Has), key)
}

// Iterator mocks base method.
func (m *MockDBReadWriter) Iterator(start, end []byte) (db.Iterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iterator", start, end)
	ret0, _ := ret[0].(db.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Iterator indicates an expected call of Iterator.
func (mr *MockDBReadWriterMockRecorder) Iterator(start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterator", reflect.TypeOf((*MockDBReadWriter)(nil).Iterator), start, end)
}

// ReverseIterator mocks base method.
func (m *MockDBReadWriter) ReverseIterator(start, end []byte) (db.Iterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReverseIterator", start, end)
	ret0, _ := ret[0].(db.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReverseIterator indicates an expected call of ReverseIterator.
func (mr *MockDBReadWriterMockRecorder) ReverseIterator(start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReverseIterator", reflect.TypeOf((*MockDBReadWriter)(nil).ReverseIterator), start, end)
}

// Set mocks base method.
func (m *MockDBReadWriter) Set(arg0, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockDBReadWriterMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockDBReadWriter)(nil).Set), arg0, arg1)
}

// MockIterator is a mock of Iterator interface.
type MockIterator struct {
	ctrl     *gomock.Controller
	recorder *MockIteratorMockRecorder
}

// MockIteratorMockRecorder is the mock recorder for MockIterator.
type MockIteratorMockRecorder struct {
	mock *MockIterator
}

// NewMockIterator creates a new mock instance.
func NewMockIterator(ctrl *gomock.Controller) *MockIterator {
	mock := &MockIterator{ctrl: ctrl}
	mock.recorder = &MockIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIterator) EXPECT() *MockIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIterator) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIterator)(nil).Close))
}

// Domain mocks base method.
func (m *MockIterator) Domain() ([]byte, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Domain")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// Domain indicates an expected call of Domain.
func (mr *MockIteratorMockRecorder) Domain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Domain", reflect.TypeOf((*MockIterator)(nil).Domain))
}

// Error mocks base method.
func (m *MockIterator) Error() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockIteratorMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockIterator)(nil).Error))
}

// Key mocks base method.
func (m *MockIterator) Key() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Key")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Key indicates an expected call of Key.
func (mr *MockIteratorMockRecorder) Key() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Key", reflect.TypeOf((*MockIterator)(nil).Key))
}

// Next mocks base method.
func (m *MockIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockIterator)(nil).Next))
}

// Value mocks base method.
func (m *MockIterator) Value() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Value indicates an expected call of Value.
func (mr *MockIteratorMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockIterator)(nil).Value))
}

// MockVersionSet is a mock of VersionSet interface.
type MockVersionSet struct {
	ctrl     *gomock.Controller
	recorder *MockVersionSetMockRecorder
}

// MockVersionSetMockRecorder is the mock recorder for MockVersionSet.
type MockVersionSetMockRecorder struct {
	mock *MockVersionSet
}

// NewMockVersionSet creates a new mock instance.
func NewMockVersionSet(ctrl *gomock.Controller) *MockVersionSet {
	mock := &MockVersionSet{ctrl: ctrl}
	mock.recorder = &MockVersionSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVersionSet) EXPECT() *MockVersionSetMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockVersionSet) Count() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(int)
	return ret0
}

// Count indicates an expected call of Count.
func (mr *MockVersionSetMockRecorder) Count() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockVersionSet)(nil).Count))
}

// Equal mocks base method.
func (m *MockVersionSet) Equal(arg0 db.VersionSet) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockVersionSetMockRecorder) Equal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockVersionSet)(nil).Equal), arg0)
}

// Exists mocks base method.
func (m *MockVersionSet) Exists(arg0 uint64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockVersionSetMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockVersionSet)(nil).Exists), arg0)
}

// Iterator mocks base method.
func (m *MockVersionSet) Iterator() db.VersionIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iterator")
	ret0, _ := ret[0].(db.VersionIterator)
	return ret0
}

// Iterator indicates an expected call of Iterator.
func (mr *MockVersionSetMockRecorder) Iterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterator", reflect.TypeOf((*MockVersionSet)(nil).Iterator))
}

// Last mocks base method.
func (m *MockVersionSet) Last() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Last")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Last indicates an expected call of Last.
func (mr *MockVersionSetMockRecorder) Last() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Last", reflect.TypeOf((*MockVersionSet)(nil).Last))
}

// MockVersionIterator is a mock of VersionIterator interface.
type MockVersionIterator struct {
	ctrl     *gomock.Controller
	recorder *MockVersionIteratorMockRecorder
}

// MockVersionIteratorMockRecorder is the mock recorder for MockVersionIterator.
type MockVersionIteratorMockRecorder struct {
	mock *MockVersionIterator
}

// NewMockVersionIterator creates a new mock instance.
func NewMockVersionIterator(ctrl *gomock.Controller) *MockVersionIterator {
	mock := &MockVersionIterator{ctrl: ctrl}
	mock.recorder = &MockVersionIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVersionIterator) EXPECT() *MockVersionIteratorMockRecorder {
	return m.recorder
}

// Next mocks base method.
func (m *MockVersionIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockVersionIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockVersionIterator)(nil).Next))
}

// Value mocks base method.
func (m *MockVersionIterator) Value() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Value indicates an expected call of Value.
func (mr *MockVersionIteratorMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockVersionIterator)(nil).Value))
}
