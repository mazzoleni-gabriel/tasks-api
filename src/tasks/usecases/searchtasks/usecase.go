package searchtasks

import (
	"context"
	"tasks-api/src/tasks"
)

type TaskSearcher struct {
	reader Reader
}

//go:generate mockery --name=Reader --disable-version-string
type Reader interface {
	Search(ctx context.Context, filters tasks.SearchFilters) ([]tasks.Task, error)
}

func NewUseCase(reader Reader) TaskSearcher {
	return TaskSearcher{
		reader: reader,
	}
}

func (u TaskSearcher) Search(ctx context.Context, filters tasks.SearchFilters) ([]tasks.Task, error) {
	return u.reader.Search(ctx, filters)
}
