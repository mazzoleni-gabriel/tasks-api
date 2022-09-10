package repository

import (
	"context"
	"gorm.io/gorm"
	"tasks-api/src/tasks"
	"tasks-api/src/tasks/repository/models"
)

type ReaderMySQL struct {
	db *gorm.DB
}

func NewReaderMySQL(db *gorm.DB) ReaderMySQL {
	return ReaderMySQL{db: db}
}

func (r ReaderMySQL) Search(ctx context.Context, filters tasks.SearchFilters) (tasks []tasks.Task, err error) {
	var tasksModel []models.Task
	query := r.db.WithContext(ctx)
	if filters.CreatedBy != nil {
		query = query.Where("created_by = ?", filters.CreatedBy)
	}
	query = query.Find(&tasksModel)

	for _, t := range tasksModel {
		tasks = append(tasks, t.ToEntity())
	}

	return tasks, query.Error
}
