package publisher

import (
	"context"
	"tasks-api/internal/config"
	"tasks-api/internal/topicpublisher"
	"tasks-api/src/tasks"
)

type TaskFeedPublisher struct {
	TopicPublisher TopicPublisher
}

//go:generate mockery --name=TopicPublisher --disable-version-string
type TopicPublisher interface {
	Publish(ctx context.Context, body interface{}) error
}

func NewTaskFeedPublisher(cfg config.Configuration) (TaskFeedPublisher, error) {
	publisher, err := topicpublisher.NewTopicPublisher(cfg.TopicTaskFeed, cfg.TopicTaskFeedScope)
	if err != nil {
		return TaskFeedPublisher{}, err
	}
	return TaskFeedPublisher{
		TopicPublisher: publisher,
	}, nil
}

func (p TaskFeedPublisher) PublishCreateMessage(ctx context.Context, task tasks.Task) error {
	return p.TopicPublisher.Publish(ctx, newCreateMessage(task))
}

func (p TaskFeedPublisher) PublishUpdateMessage(ctx context.Context, task tasks.Task) error {
	return p.TopicPublisher.Publish(ctx, newUpdateMessage(task))
}

func (p TaskFeedPublisher) PublishDeleteMessage(ctx context.Context, task tasks.Task) error {
	return p.TopicPublisher.Publish(ctx, newDeleteMessage(task))
}
