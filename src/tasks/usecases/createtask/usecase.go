package createtask

import (
	"context"
	"tasks-api/src/tasks"
)

type TaskCreator struct {
	writer Writer
}

//go:generate mockery --name=Writer --disable-version-string
type Writer interface {
	Create(ctx context.Context, task tasks.Task) (uint, error)
}

func NewUseCase(writer Writer) TaskCreator {
	return TaskCreator{
		writer: writer,
	}
}

func (u TaskCreator) Create(ctx context.Context, task tasks.Task) (uint, error) {
	return u.writer.Create(ctx, task)
}
