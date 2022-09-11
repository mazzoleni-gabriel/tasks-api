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

func (w WriterMySQL) Update(ctx context.Context, entity tasks.Task) (int64, error) {
	model := models.NewTaskFromEntity(entity)
	r := w.db.WithContext(ctx).
		Model(&model).
		Updates(model)
	return r.RowsAffected, r.Error
}

func (w WriterMySQL) Delete(ctx context.Context, id uint) (int64, error) {
	r := w.db.WithContext(ctx).Delete(&models.Task{}, id)
	return r.RowsAffected, r.Error
}
