// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/vitalik-ez/Chat-Golang/pkg/domain/entity"
)

// MockRoom is a mock of Room interface.
type MockRoom struct {
	ctrl     *gomock.Controller
	recorder *MockRoomMockRecorder
}

// MockRoomMockRecorder is the mock recorder for MockRoom.
type MockRoomMockRecorder struct {
	mock *MockRoom
}

// NewMockRoom creates a new mock instance.
func NewMockRoom(ctrl *gomock.Controller) *MockRoom {
	mock := &MockRoom{ctrl: ctrl}
	mock.recorder = &MockRoomMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoom) EXPECT() *MockRoomMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRoom) Create(room string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", room)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRoomMockRecorder) Create(room interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRoom)(nil).Create), room)
}

// MockMessage is a mock of Message interface.
type MockMessage struct {
	ctrl     *gomock.Controller
	recorder *MockMessageMockRecorder
}

// MockMessageMockRecorder is the mock recorder for MockMessage.
type MockMessageMockRecorder struct {
	mock *MockMessage
}

// NewMockMessage creates a new mock instance.
func NewMockMessage(ctrl *gomock.Controller) *MockMessage {
	mock := &MockMessage{ctrl: ctrl}
	mock.recorder = &MockMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessage) EXPECT() *MockMessageMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMessage) Create(message entity.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMessageMockRecorder) Create(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMessage)(nil).Create), message)
}
