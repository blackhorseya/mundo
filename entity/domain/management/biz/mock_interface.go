// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package biz is a generated GoMock package.
package biz

import (
	reflect "reflect"

	agg "github.com/blackhorseya/mundo/entity/domain/identity/agg"
	agg0 "github.com/blackhorseya/mundo/entity/domain/management/agg"
	contextx "github.com/blackhorseya/mundo/pkg/contextx"
	gomock "go.uber.org/mock/gomock"
)

// MockIManagementBiz is a mock of IManagementBiz interface.
type MockIManagementBiz struct {
	ctrl     *gomock.Controller
	recorder *MockIManagementBizMockRecorder
}

// MockIManagementBizMockRecorder is the mock recorder for MockIManagementBiz.
type MockIManagementBizMockRecorder struct {
	mock *MockIManagementBiz
}

// NewMockIManagementBiz creates a new mock instance.
func NewMockIManagementBiz(ctrl *gomock.Controller) *MockIManagementBiz {
	mock := &MockIManagementBiz{ctrl: ctrl}
	mock.recorder = &MockIManagementBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIManagementBiz) EXPECT() *MockIManagementBizMockRecorder {
	return m.recorder
}

// CreateWordBook mocks base method.
func (m *MockIManagementBiz) CreateWordBook(ctx contextx.Contextx, by *agg.Member, name string) (*agg0.Wordbook, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWordBook", ctx, by, name)
	ret0, _ := ret[0].(*agg0.Wordbook)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWordBook indicates an expected call of CreateWordBook.
func (mr *MockIManagementBizMockRecorder) CreateWordBook(ctx, by, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWordBook", reflect.TypeOf((*MockIManagementBiz)(nil).CreateWordBook), ctx, by, name)
}

// GetWordBookByID mocks base method.
func (m *MockIManagementBiz) GetWordBookByID(ctx contextx.Contextx, id string) (*agg0.Wordbook, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWordBookByID", ctx, id)
	ret0, _ := ret[0].(*agg0.Wordbook)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWordBookByID indicates an expected call of GetWordBookByID.
func (mr *MockIManagementBizMockRecorder) GetWordBookByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWordBookByID", reflect.TypeOf((*MockIManagementBiz)(nil).GetWordBookByID), ctx, id)
}