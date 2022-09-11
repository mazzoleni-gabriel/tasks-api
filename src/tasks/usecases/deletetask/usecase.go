package deletetask

import (
	"context"
	"fmt"
	"tasks-api/src/tasks"
)

type TaskDeleter struct {
	writer    Writer
	publisher Publisher
}

//go:generate mockery --name=Writer --disable-version-string
type Writer interface {
	Delete(ctx context.Context, id uint) (int64, error)
}

//go:generate mockery --name=Publisher --disable-version-string
type Publisher interface {
	PublishDeleteMessage(ctx context.Context, task tasks.Task) error
}

func NewUseCase(writer Writer, publisher Publisher) TaskDeleter {
	return TaskDeleter{
		writer:    writer,
		publisher: publisher,
	}
}

func (u TaskDeleter) Delete(ctx context.Context, id uint) (int64, error) {
	affectedRows, err := u.writer.Delete(ctx, id)
	if err != nil {
		return 0, err
	}

	err = u.publisher.PublishDeleteMessage(ctx, tasks.Task{ID: id})
	if err != nil {
		fmt.Printf("fail to publish deletion message for task %d, this use case should trigger an alert", id)
	}
	return affectedRows, err
}
