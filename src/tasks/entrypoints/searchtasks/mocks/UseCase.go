// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	tasks "tasks-api/src/tasks"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// Search provides a mock function with given fields: ctx, filters
func (_m *UseCase) Search(ctx context.Context, filters tasks.SearchFilters) ([]tasks.Task, error) {
	ret := _m.Called(ctx, filters)

	var r0 []tasks.Task
	if rf, ok := ret.Get(0).(func(context.Context, tasks.SearchFilters) []tasks.Task); ok {
		r0 = rf(ctx, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]tasks.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, tasks.SearchFilters) error); ok {
		r1 = rf(ctx, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCase(t mockConstructorTestingTNewUseCase) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
