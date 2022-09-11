package models

import (
	"github.com/jinzhu/gorm"
	"tasks-api/src/tasks"
	"time"
)

type Task struct {
	gorm.Model
	Summary     string
	CreatedBy   uint
	PerformedAt time.Time
}

func NewTaskFromEntity(entity tasks.Task) Task {
	task := Task{
		Summary:     entity.Summary,
		CreatedBy:   entity.CreatedBy,
		PerformedAt: entity.PerformedAt,
	}
	if entity.ID != 0 {
		task.ID = entity.ID
	}
	return task
}

func (t Task) ToEntity() tasks.Task {
	return tasks.Task{
		ID:          t.ID,
		Summary:     t.Summary,
		CreatedBy:   t.CreatedBy,
		PerformedAt: t.PerformedAt,
	}
}
