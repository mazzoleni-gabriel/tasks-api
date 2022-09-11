package topicpublisher

import (
	"context"
	"fmt"
)

type TopicPublisher interface {
	Publish(ctx context.Context, body interface{}) error
}

func NewTopicPublisher(topic string, scope string) (TopicPublisher, error) {
	if scope == "fake" {
		return newFakeTopicPublisher(topic), nil
	}
	return nil, fmt.Errorf("topic scope not supported")
}
