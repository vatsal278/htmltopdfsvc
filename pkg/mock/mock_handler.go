// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vatsal278/htmltopdfsvc/internal/handler (interfaces: HtmltopdfsvcHandler)

// Package mock is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHtmltopdfsvcHandler is a mock of HtmltopdfsvcHandler interface.
type MockHtmltopdfsvcHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHtmltopdfsvcHandlerMockRecorder
}

// MockHtmltopdfsvcHandlerMockRecorder is the mock recorder for MockHtmltopdfsvcHandler.
type MockHtmltopdfsvcHandlerMockRecorder struct {
	mock *MockHtmltopdfsvcHandler
}

// NewMockHtmltopdfsvcHandler creates a new mock instance.
func NewMockHtmltopdfsvcHandler(ctrl *gomock.Controller) *MockHtmltopdfsvcHandler {
	mock := &MockHtmltopdfsvcHandler{ctrl: ctrl}
	mock.recorder = &MockHtmltopdfsvcHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHtmltopdfsvcHandler) EXPECT() *MockHtmltopdfsvcHandlerMockRecorder {
	return m.recorder
}

// HealthCheck mocks base method.
func (m *MockHtmltopdfsvcHandler) HealthCheck() (string, string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockHtmltopdfsvcHandlerMockRecorder) HealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockHtmltopdfsvcHandler)(nil).HealthCheck))
}

// Ping mocks base method.
func (m *MockHtmltopdfsvcHandler) Ping(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Ping", arg0, arg1)
}

// Ping indicates an expected call of Ping.
func (mr *MockHtmltopdfsvcHandlerMockRecorder) Ping(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockHtmltopdfsvcHandler)(nil).Ping), arg0, arg1)
}
