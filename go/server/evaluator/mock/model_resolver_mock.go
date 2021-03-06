// Code generated by MockGen. DO NOT EDIT.
// Source: evaluator/model_resolver.go

// Package mock_evaluator is a generated GoMock package.
package mock_evaluator

import (
	repositories "github.com/marekgalovic/photon/go/core/repositories"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockModelResolver is a mock of ModelResolver interface
type MockModelResolver struct {
	ctrl     *gomock.Controller
	recorder *MockModelResolverMockRecorder
}

// MockModelResolverMockRecorder is the mock recorder for MockModelResolver
type MockModelResolverMockRecorder struct {
	mock *MockModelResolver
}

// NewMockModelResolver creates a new mock instance
func NewMockModelResolver(ctrl *gomock.Controller) *MockModelResolver {
	mock := &MockModelResolver{ctrl: ctrl}
	mock.recorder = &MockModelResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModelResolver) EXPECT() *MockModelResolverMockRecorder {
	return m.recorder
}

// GetModel mocks base method
func (m *MockModelResolver) GetModel(arg0 string) (*repositories.Model, *repositories.ModelVersion, error) {
	ret := m.ctrl.Call(m, "GetModel", arg0)
	ret0, _ := ret[0].(*repositories.Model)
	ret1, _ := ret[1].(*repositories.ModelVersion)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetModel indicates an expected call of GetModel
func (mr *MockModelResolverMockRecorder) GetModel(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModel", reflect.TypeOf((*MockModelResolver)(nil).GetModel), arg0)
}
