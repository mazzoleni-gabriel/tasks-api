package topicpublisher

import (
	"context"
	"log"
)

type fakeTopicPublisher struct {
	TopicName string
}

func newFakeTopicPublisher(topic string) fakeTopicPublisher {
	return fakeTopicPublisher{
		TopicName: topic,
	}
}

func (t fakeTopicPublisher) Publish(_ context.Context, body interface{}) error {
	log.Printf("Fake publishing message %v to topic %s", body, t.TopicName)
	return nil
}
