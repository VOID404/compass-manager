// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Registrator is an autogenerated mock type for the Registrator type
type Registrator struct {
	mock.Mock
}

// ConfigureRuntimeAgent provides a mock function with given fields: kubeconfig, runtimeID
func (_m *Registrator) ConfigureRuntimeAgent(kubeconfig string, runtimeID string) error {
	ret := _m.Called(kubeconfig, runtimeID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(kubeconfig, runtimeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Register provides a mock function with given fields: name
func (_m *Registrator) Register(name string) (string, error) {
	ret := _m.Called(name)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRegistrator interface {
	mock.TestingT
	Cleanup(func())
}

// NewRegistrator creates a new instance of Registrator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRegistrator(t mockConstructorTestingTNewRegistrator) *Registrator {
	mock := &Registrator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
