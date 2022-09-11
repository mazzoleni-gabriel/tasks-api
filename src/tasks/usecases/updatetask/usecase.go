package updatetask

import (
	"context"
	"fmt"
	"tasks-api/internal/apperror"
	"tasks-api/src/tasks"
)

type TaskUpdater struct {
	writer Writer
	reader Reader
}

//go:generate mockery --name=Writer --disable-version-string
type Writer interface {
	Update(ctx context.Context, entity tasks.Task) (int64, error)
}

//go:generate mockery --name=Reader --disable-version-string
type Reader interface {
	Search(ctx context.Context, filters tasks.SearchFilters) (tasks []tasks.Task, err error)
}

func NewUseCase(writer Writer, reader Reader) TaskUpdater {
	return TaskUpdater{
		writer: writer,
		reader: reader,
	}
}

func (u TaskUpdater) Update(ctx context.Context, task tasks.Task, userID uint) (int64, error) {
	if err := u.validateTask(ctx, task.ID, userID); err != nil {
		return 0, err
	}

	return u.writer.Update(ctx, task)
}

func (u TaskUpdater) validateTask(ctx context.Context, id uint, userID uint) error {
	foundTasks, err := u.reader.Search(ctx, tasks.SearchFilters{ID: &id})
	if err != nil {
		return err
	}
	if len(foundTasks) == 0 {
		return apperror.New(apperror.ErrTaskNotFound, fmt.Sprintf("the resource with id %d does not exists", id))
	}

	if foundTasks[0].CreatedBy != userID {
		return apperror.New(apperror.ErrOtherUserTask, fmt.Sprintf("the task %d belong to another user", id))
	}

	return nil
}
