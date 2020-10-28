// Code generated by MockGen. DO NOT EDIT.
// Source: ./slack.go

// Package mock_slack is a generated GoMock package.
package mock_slack

import (
	bytes "bytes"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockSlack is a mock of Slack interface
type MockSlack struct {
	ctrl     *gomock.Controller
	recorder *MockSlackMockRecorder
}

// MockSlackMockRecorder is the mock recorder for MockSlack
type MockSlackMockRecorder struct {
	mock *MockSlack
}

// NewMockSlack creates a new mock instance
func NewMockSlack(ctrl *gomock.Controller) *MockSlack {
	mock := &MockSlack{ctrl: ctrl}
	mock.recorder = &MockSlackMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSlack) EXPECT() *MockSlackMockRecorder {
	return m.recorder
}

// Upload mocks base method
func (m *MockSlack) Upload(title, name, channel, ts string, r io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", title, name, channel, ts, r)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload
func (mr *MockSlackMockRecorder) Upload(title, name, channel, ts, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockSlack)(nil).Upload), title, name, channel, ts, r)
}

// Download mocks base method
func (m *MockSlack) Download(url string, b *bytes.Buffer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", url, b)
	ret0, _ := ret[0].(error)
	return ret0
}

// Download indicates an expected call of Download
func (mr *MockSlackMockRecorder) Download(url, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockSlack)(nil).Download), url, b)
}