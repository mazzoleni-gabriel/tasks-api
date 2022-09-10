package deletetask

import (
	"context"
)

type TaskDeleter struct {
	writer Writer
}

//go:generate mockery --name=Writer --disable-version-string
type Writer interface {
	Delete(ctx context.Context, id uint) (int64, error)
}

func NewUseCase(writer Writer) TaskDeleter {
	return TaskDeleter{
		writer: writer,
	}
}

func (u TaskDeleter) Delete(ctx context.Context, id uint) (int64, error) {
	return u.writer.Delete(ctx, id)
}
