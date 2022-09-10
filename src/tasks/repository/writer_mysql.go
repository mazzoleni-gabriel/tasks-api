package repository

import (
	"context"
	"gorm.io/gorm"
	"tasks-api/src/tasks"
	"tasks-api/src/tasks/repository/models"
)

type WriterMySQL struct {
	db *gorm.DB
}

func NewWriterMySQL(db *gorm.DB) WriterMySQL {
	return WriterMySQL{db: db}
}

func (w WriterMySQL) Create(ctx context.Context, entity tasks.Task) (uint, error) {
	model := models.NewTaskFromEntity(entity)
	r := w.db.WithContext(ctx).Create(&model)
	return model.ID, r.Error
}
