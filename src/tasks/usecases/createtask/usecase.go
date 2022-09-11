package createtask

import (
	"context"
	"fmt"
	"tasks-api/src/tasks"
)

type TaskCreator struct {
	writer    Writer
	publisher Publisher
}

//go:generate mockery --name=Writer --disable-version-string
type Writer interface {
	Create(ctx context.Context, task tasks.Task) (uint, error)
}

//go:generate mockery --name=Publisher --disable-version-string
type Publisher interface {
	PublishCreateMessage(ctx context.Context, task tasks.Task) error
}

func NewUseCase(writer Writer, publisher Publisher) TaskCreator {
	return TaskCreator{
		writer:    writer,
		publisher: publisher,
	}
}

func (u TaskCreator) Create(ctx context.Context, task tasks.Task) (uint, error) {
	createdID, err := u.writer.Create(ctx, task)
	if err != nil {
		return 0, err
	}

	task.ID = createdID
	err = u.publisher.PublishCreateMessage(ctx, task)
	if err != nil {
		fmt.Printf("fail to publish creation message for task %d, this use case should trigger an alert", createdID)
	}

	return createdID, nil
}
