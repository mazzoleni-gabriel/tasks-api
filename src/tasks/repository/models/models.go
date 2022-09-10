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
	return Task{
		Summary:     entity.Summary,
		CreatedBy:   entity.CreatedBy,
		PerformedAt: entity.PerformedAt,
	}
}
