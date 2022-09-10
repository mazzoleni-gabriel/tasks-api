package createtask_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"tasks-api/src/tasks"
	"tasks-api/src/tasks/usecases/createtask"
	"tasks-api/src/tasks/usecases/createtask/mocks"
	"testing"
)

func Test_UseCase(t *testing.T) {
	t.Run("Should return created id", func(t *testing.T) {
		entity := tasks.Task{}

		writerMock := mocks.NewWriter(t)
		writerMock.On("Create", mock.Anything, entity).
			Return(uint(1), nil)

		useCase := createtask.NewUseCase(writerMock)
		createdID, err := useCase.Create(context.TODO(), entity)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), createdID)
		writerMock.AssertExpectations(t)
	})

	t.Run("Should return error when writer fails", func(t *testing.T) {
		entity := tasks.Task{}

		writerMock := mocks.NewWriter(t)
		writerMock.On("Create", mock.Anything, entity).
			Return(uint(0), assert.AnError)

		useCase := createtask.NewUseCase(writerMock)
		_, err := useCase.Create(context.TODO(), entity)

		assert.Error(t, err)
		writerMock.AssertExpectations(t)
	})
}
