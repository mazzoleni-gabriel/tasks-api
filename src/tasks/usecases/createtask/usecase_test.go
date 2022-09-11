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
		createdID := uint(1)

		writerMock := mocks.NewWriter(t)
		writerMock.On("Create", mock.Anything, entity).
			Return(createdID, nil)

		publishedTask := tasks.Task{ID: createdID}
		publisherMock := mocks.NewPublisher(t)
		publisherMock.On("PublishCreateMessage", mock.Anything, publishedTask).
			Return(nil)

		useCase := createtask.NewUseCase(writerMock, publisherMock)
		createdID, err := useCase.Create(context.TODO(), entity)

		assert.NoError(t, err)
		assert.Equal(t, createdID, createdID)
		writerMock.AssertExpectations(t)
		publisherMock.AssertExpectations(t)
	})

	t.Run("Should return error when writer fails", func(t *testing.T) {
		entity := tasks.Task{}

		writerMock := mocks.NewWriter(t)
		writerMock.On("Create", mock.Anything, entity).
			Return(uint(0), assert.AnError)

		useCase := createtask.NewUseCase(writerMock, nil)
		_, err := useCase.Create(context.TODO(), entity)

		assert.Error(t, err)
		writerMock.AssertExpectations(t)
	})
}
