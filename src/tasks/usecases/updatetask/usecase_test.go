package updatetask_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"tasks-api/internal/apperror"
	"tasks-api/src/tasks"
	"tasks-api/src/tasks/usecases/updatetask"
	"tasks-api/src/tasks/usecases/updatetask/mocks"
	"testing"
)

func Test_UseCase(t *testing.T) {
	const userID = 1

	t.Run("Should return no error when success", func(t *testing.T) {
		entity := tasks.Task{
			ID: 2,
		}

		foundEntities := []tasks.Task{
			{
				CreatedBy: userID,
			},
		}

		filters := tasks.SearchFilters{
			ID: &entity.ID,
		}

		readerMock := mocks.NewReader(t)
		readerMock.On("Search", mock.Anything, filters).
			Return(foundEntities, nil)

		writerMock := mocks.NewWriter(t)
		writerMock.On("Update", mock.Anything, entity).
			Return(int64(1), nil)

		useCase := updatetask.NewUseCase(writerMock, readerMock)
		_, err := useCase.Update(context.TODO(), entity, userID)

		assert.NoError(t, err)
		writerMock.AssertExpectations(t)
	})

	t.Run("Should return error when update fails", func(t *testing.T) {
		entity := tasks.Task{
			ID: 2,
		}

		foundEntities := []tasks.Task{
			{
				CreatedBy: userID,
			},
		}

		filters := tasks.SearchFilters{
			ID: &entity.ID,
		}

		readerMock := mocks.NewReader(t)
		readerMock.On("Search", mock.Anything, filters).
			Return(foundEntities, nil)

		writerMock := mocks.NewWriter(t)
		writerMock.On("Update", mock.Anything, entity).
			Return(int64(0), assert.AnError)

		useCase := updatetask.NewUseCase(writerMock, readerMock)
		_, err := useCase.Update(context.TODO(), entity, userID)

		assert.Error(t, err)
		writerMock.AssertExpectations(t)
	})

	t.Run("Should return ErrOtherUserTask error when is other user task", func(t *testing.T) {
		entity := tasks.Task{
			ID: 2,
		}

		foundEntities := []tasks.Task{
			{
				CreatedBy: 123,
			},
		}

		filters := tasks.SearchFilters{
			ID: &entity.ID,
		}

		readerMock := mocks.NewReader(t)
		readerMock.On("Search", mock.Anything, filters).
			Return(foundEntities, nil)

		writerMock := mocks.NewWriter(t)

		useCase := updatetask.NewUseCase(writerMock, readerMock)
		_, err := useCase.Update(context.TODO(), entity, userID)

		assert.Error(t, err)
		assert.True(t, apperror.IsAppError(err, apperror.ErrOtherUserTask))
		writerMock.AssertExpectations(t)
	})

	t.Run("Should return ErrTaskNotFound error when not task was found", func(t *testing.T) {
		entity := tasks.Task{
			ID: 2,
		}

		var foundEntities []tasks.Task

		filters := tasks.SearchFilters{
			ID: &entity.ID,
		}

		readerMock := mocks.NewReader(t)
		readerMock.On("Search", mock.Anything, filters).
			Return(foundEntities, nil)

		writerMock := mocks.NewWriter(t)

		useCase := updatetask.NewUseCase(writerMock, readerMock)
		_, err := useCase.Update(context.TODO(), entity, userID)

		assert.Error(t, err)
		assert.True(t, apperror.IsAppError(err, apperror.ErrTaskNotFound))
		writerMock.AssertExpectations(t)
	})

	t.Run("Should return general error when search task fails", func(t *testing.T) {
		entity := tasks.Task{
			ID: 2,
		}

		filters := tasks.SearchFilters{
			ID: &entity.ID,
		}

		readerMock := mocks.NewReader(t)
		readerMock.On("Search", mock.Anything, filters).
			Return([]tasks.Task{}, assert.AnError)

		writerMock := mocks.NewWriter(t)

		useCase := updatetask.NewUseCase(writerMock, readerMock)
		_, err := useCase.Update(context.TODO(), entity, userID)

		assert.Error(t, err)
		writerMock.AssertExpectations(t)
	})
}
