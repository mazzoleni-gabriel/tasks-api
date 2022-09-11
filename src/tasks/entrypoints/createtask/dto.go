package createtask

import (
	"fmt"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	"tasks-api/internal/validator"
	"tasks-api/src/tasks"
	"time"
)

const userHeader = "X-User-ID"

// swagger:model CreateTaskRequest
type CreateTaskRequest struct {
	Summary     string    `json:"summary" validate:"required,max=2500"`
	PerformedAt time.Time `json:"performed_at" validate:"required"`
	CreatedBy   uint      `json:"-"`
}

// swagger:model CreateTaskResponse
type CreateTaskResponse struct {
	ID uint `json:"id"`
}

func (dto *CreateTaskRequest) Bind(_ *http.Request) error {
	return validator.Validate(dto)
}

func (dto *CreateTaskRequest) toEntity() tasks.Task {
	return tasks.Task{
		Summary:     dto.Summary,
		PerformedAt: dto.PerformedAt,
		CreatedBy:   dto.CreatedBy,
	}
}

func newTaskFromRequest(r *http.Request) (tasks.Task, error) {
	dto := CreateTaskRequest{}
	if err := render.Bind(r, &dto); err != nil {
		return tasks.Task{}, err
	}

	userID, err := getUserID(r)
	if err != nil {
		return tasks.Task{}, err
	}
	dto.CreatedBy = userID

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
