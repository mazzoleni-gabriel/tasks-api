package updatetask

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	"tasks-api/internal/validator"
	"tasks-api/src/tasks"
	"time"
)

const userHeader = "X-User-ID"

type UpdateTaskRequest struct {
	Summary     string    `json:"summary" validate:"max=2500"`
	PerformedAt time.Time `json:"performed_at"`
}

func (dto *UpdateTaskRequest) Bind(_ *http.Request) error {
	return validator.Validate(dto)
}

func (dto *UpdateTaskRequest) toEntity() tasks.Task {
	return tasks.Task{
		Summary:     dto.Summary,
		PerformedAt: dto.PerformedAt,
	}
}

func newTaskFromRequest(r *http.Request) (tasks.Task, error) {
	dto := UpdateTaskRequest{}
	if err := render.Bind(r, &dto); err != nil {
		return tasks.Task{}, err
	}
	return dto.toEntity(), nil
}

func getUserID(r *http.Request) (uint, error) {
	strUserID := r.Header.Get(userHeader)
	if strUserID == "" {
		return 0, fmt.Errorf("header %s is required", userHeader)
	}

	userID, err := strconv.ParseUint(strUserID, 10, 64)
	return uint(userID), err
}

func getTaskID(r *http.Request) (uint, error) {
	strTaskID := chi.URLParam(r, "task_id")
	taskID, err := strconv.ParseUint(strTaskID, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("unnable to parse task_id: %w", err)
	}

	if taskID == 0 {
		return 0, fmt.Errorf("task_id is required")
	}

	return uint(taskID), nil
}
