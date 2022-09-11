package publisher_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"tasks-api/src/tasks"
	"tasks-api/src/tasks/publisher"
	"tasks-api/src/tasks/publisher/mocks"
	"testing"
)

func Test_PublishCreateMessage(t *testing.T) {
	t.Run("Should return success", func(t *testing.T) {

		task := tasks.Task{ID: 1}
		expectedMsg := publisher.Message{
			Operation: publisher.CreateOperation,
			Entity: publisher.Entity{
				ID: task.ID,
			},
		}

		publisherMock := mocks.NewTopicPublisher(t)
		publisherMock.On("Publish", mock.Anything, expectedMsg).
			Return(nil)

		service := publisher.TaskFeedPublisher{TopicPublisher: publisherMock}
		err := service.PublishCreateMessage(context.TODO(), task)

		assert.NoError(t, err)
		publisherMock.AssertExpectations(t)
	})

	t.Run("Should return error when publisher fails", func(t *testing.T) {

		task := tasks.Task{ID: 1}
		expectedMsg := publisher.Message{
			Operation: publisher.CreateOperation,
			Entity: publisher.Entity{
				ID: task.ID,
			},
		}

		publisherMock := mocks.NewTopicPublisher(t)
		publisherMock.On("Publish", mock.Anything, expectedMsg).
			Return(assert.AnError)

		service := publisher.TaskFeedPublisher{TopicPublisher: publisherMock}
		err := service.PublishCreateMessage(context.TODO(), task)

		assert.Error(t, err)
		publisherMock.AssertExpectations(t)
	})
}

func Test_PublishUpdateMessage(t *testing.T) {
	t.Run("Should return success", func(t *testing.T) {

		task := tasks.Task{ID: 1}
		expectedMsg := publisher.Message{
			Operation: publisher.UpdateOperation,
			Entity: publisher.Entity{
				ID: task.ID,
			},
		}

		publisherMock := mocks.NewTopicPublisher(t)
		publisherMock.On("Publish", mock.Anything, expectedMsg).
			Return(nil)

		service := publisher.TaskFeedPublisher{TopicPublisher: publisherMock}
		err := service.PublishUpdateMessage(context.TODO(), task)

		assert.NoError(t, err)
		publisherMock.AssertExpectations(t)
	})

	t.Run("Should return error when publisher fails", func(t *testing.T) {

		task := tasks.Task{ID: 1}
		expectedMsg := publisher.Message{
			Operation: publisher.UpdateOperation,
			Entity: publisher.Entity{
				ID: task.ID,
			},
		}

		publisherMock := mocks.NewTopicPublisher(t)
		publisherMock.On("Publish", mock.Anything, expectedMsg).
			Return(assert.AnError)

		service := publisher.TaskFeedPublisher{TopicPublisher: publisherMock}
		err := service.PublishUpdateMessage(context.TODO(), task)

		assert.Error(t, err)
		publisherMock.AssertExpectations(t)
	})
}

func Test_PublishDeleteMessage(t *testing.T) {
	t.Run("Should return success", func(t *testing.T) {

		task := tasks.Task{ID: 1}
		expectedMsg := publisher.Message{
			Operation: publisher.DeleteOperation,
			Entity: publisher.Entity{
				ID: task.ID,
			},
		}

		publisherMock := mocks.NewTopicPublisher(t)
		publisherMock.On("Publish", mock.Anything, expectedMsg).
			Return(nil)

		service := publisher.TaskFeedPublisher{TopicPublisher: publisherMock}
		err := service.PublishDeleteMessage(context.TODO(), task)

		assert.NoError(t, err)
		publisherMock.AssertExpectations(t)
	})

	t.Run("Should return error when publisher fails", func(t *testing.T) {

		task := tasks.Task{ID: 1}
		expectedMsg := publisher.Message{
			Operation: publisher.DeleteOperation,
			Entity: publisher.Entity{
				ID: task.ID,
			},
		}

		publisherMock := mocks.NewTopicPublisher(t)
		publisherMock.On("Publish", mock.Anything, expectedMsg).
			Return(assert.AnError)

		service := publisher.TaskFeedPublisher{TopicPublisher: publisherMock}
		err := service.PublishDeleteMessage(context.TODO(), task)

		assert.Error(t, err)
		publisherMock.AssertExpectations(t)
	})
}
