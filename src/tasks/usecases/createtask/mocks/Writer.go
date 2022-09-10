// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	tasks "tasks-api/src/tasks"
)

// Writer is an autogenerated mock type for the Writer type
type Writer struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, task
func (_m *Writer) Create(ctx context.Context, task tasks.Task) (uint, error) {
	ret := _m.Called(ctx, task)

	var r0 uint
	if rf, ok := ret.Get(0).(func(context.Context, tasks.Task) uint); ok {
		r0 = rf(ctx, task)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, tasks.Task) error); ok {
		r1 = rf(ctx, task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewWriter interface {
	mock.TestingT
	Cleanup(func())
}

// NewWriter creates a new instance of Writer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWriter(t mockConstructorTestingTNewWriter) *Writer {
	mock := &Writer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}