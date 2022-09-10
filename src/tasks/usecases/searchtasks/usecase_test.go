package searchtasks_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"tasks-api/src/tasks"
	"tasks-api/src/tasks/usecases/searchtasks"
	"tasks-api/src/tasks/usecases/searchtasks/mocks"
	"testing"
)

func Test_UseCase(t *testing.T) {
	t.Run("Should return tasks", func(t *testing.T) {
		var entities []tasks.Task
		var filters tasks.SearchFilters

		readerMock := mocks.NewReader(t)
		readerMock.On("Search", mock.Anything, filters).
			Return(entities, nil)

		useCase := searchtasks.NewUseCase(readerMock)
		result, err := useCase.Search(context.TODO(), filters)

		assert.NoError(t, err)
		assert.Equal(t, entities, result)
		readerMock.AssertExpectations(t)
	})

	t.Run("Should return error when reader fails", func(t *testing.T) {
		var filters tasks.SearchFilters

		readerMock := mocks.NewReader(t)
		readerMock.On("Search", mock.Anything, filters).
			Return([]tasks.Task{}, assert.AnError)

		useCase := searchtasks.NewUseCase(readerMock)
		_, err := useCase.Search(context.TODO(), filters)

		assert.Error(t, err)
		readerMock.AssertExpectations(t)
	})
}
