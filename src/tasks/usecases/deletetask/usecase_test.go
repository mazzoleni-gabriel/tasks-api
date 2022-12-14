package deletetask_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"tasks-api/src/tasks"
	"tasks-api/src/tasks/usecases/deletetask"
	"tasks-api/src/tasks/usecases/deletetask/mocks"
	"testing"
)

func Test_UseCase(t *testing.T) {
	t.Run("Should return success", func(t *testing.T) {
		taskID := uint(1)

		writerMock := mocks.NewWriter(t)
		writerMock.On("Delete", mock.Anything, taskID).
			Return(int64(1), nil)

		publishedTask := tasks.Task{ID: taskID}
		publisherMock := mocks.NewPublisher(t)
		publisherMock.On("PublishDeleteMessage", mock.Anything, publishedTask).
			Return(nil)

		useCase := deletetask.NewUseCase(writerMock, publisherMock)
		result, err := useCase.Delete(context.TODO(), taskID)

		assert.NoError(t, err)
		assert.Equal(t, int64(1), result)
		writerMock.AssertExpectations(t)
		publisherMock.AssertExpectations(t)
	})

	t.Run("Should return error when writer fails", func(t *testing.T) {
		taskID := uint(1)

		writerMock := mocks.NewWriter(t)
		writerMock.On("Delete", mock.Anything, taskID).
			Return(int64(0), assert.AnError)

		useCase := deletetask.NewUseCase(writerMock, nil)
		_, err := useCase.Delete(context.TODO(), taskID)

		assert.Error(t, err)
		writerMock.AssertExpectations(t)
	})
}
