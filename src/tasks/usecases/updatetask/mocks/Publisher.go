// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"
	tasks "tasks-api/src/tasks"

	mock "github.com/stretchr/testify/mock"
)

// Publisher is an autogenerated mock type for the Publisher type
type Publisher struct {
	mock.Mock
}

// PublishUpdateMessage provides a mock function with given fields: ctx, task
func (_m *Publisher) PublishUpdateMessage(ctx context.Context, task tasks.Task) error {
	ret := _m.Called(ctx, task)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, tasks.Task) error); ok {
		r0 = rf(ctx, task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPublisher interface {
	mock.TestingT
	Cleanup(func())
}

// NewPublisher creates a new instance of Publisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPublisher(t mockConstructorTestingTNewPublisher) *Publisher {
	mock := &Publisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
