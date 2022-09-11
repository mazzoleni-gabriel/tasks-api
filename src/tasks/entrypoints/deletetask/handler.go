package deletetask

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"tasks-api/internal/render"
)

//go:generate mockery --name=UseCase --disable-version-string
type UseCase interface {
	Delete(ctx context.Context, id uint) (int64, error)
}

type Handler struct {
	useCase UseCase
}

func NewHandler(useCase UseCase) Handler {
	return Handler{
		useCase: useCase,
	}
}

func RegisterHandler(r *chi.Mux, l Handler) {
	r.Delete("/tasks/{task_id}", l.Handle)
}

// swagger:parameters delete-task
type _ struct {
	// in:path
	// required: true
	ID int `json:"task_id"`
}

// swagger:route DELETE /tasks/{task_id} Task delete-task
//
// responses:
//   200:
//   404: ErrorResponse
//   500: ErrorResponse
func (h Handler) Handle(w http.ResponseWriter, r *http.Request) {
	strTaskID := chi.URLParam(r, "task_id")
	taskID, err := strconv.ParseUint(strTaskID, 10, 64)
	if err != nil {
		err = fmt.Errorf("unnable to parse task_id: %w", err)
		render.NewError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	if taskID == 0 {
		render.NewError(w, r, "task_id is required", http.StatusBadRequest)
		return
	}

	affectedRows, err := h.useCase.Delete(r.Context(), uint(taskID))
	if err != nil {
		render.NewError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	if affectedRows == 0 {
		err = fmt.Errorf("the resource with id %d does not exists", taskID)
		render.NewError(w, r, err.Error(), http.StatusNotFound)
		return
	}

	render.JSON(w, r, map[string]string{
		"message": "resource deleted successfully",
	})
}
