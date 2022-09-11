package updatetask

import (
	"context"
	"github.com/go-chi/chi"
	"net/http"
	"tasks-api/internal/apperror"
	"tasks-api/internal/render"
	"tasks-api/src/tasks"
)

//go:generate mockery --name=UseCase --disable-version-string
type UseCase interface {
	Update(ctx context.Context, task tasks.Task, userID uint) (int64, error)
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
	r.Put("/tasks/{task_id}", l.Handle)
}

var mapErrors = map[apperror.ErrCode]int{
	apperror.ErrTaskNotFound:  http.StatusNotFound,
	apperror.ErrOtherUserTask: http.StatusForbidden,
}

// swagger:parameters create-task
type _ struct {
	// in:body
	// required: true
	Body UpdateTaskRequest
	// in:header
	// required: true
	UserID string `json:"X-User-ID"`
}

// swagger:route PUT /tasks Task update-task
//
// responses:
//   200:
//   404: ErrorResponse
//   500: ErrorResponse
func (h Handler) Handle(w http.ResponseWriter, r *http.Request) {
	taskID, err := getTaskID(r)
	if err != nil {
		render.NewError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	entity, err := newTaskFromRequest(r)
	if err != nil {
		render.NewError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	entity.ID = taskID

	userID, err := getUserID(r)
	if err != nil {
		render.NewError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.useCase.Update(r.Context(), entity, userID)
	if err != nil {
		render.NewAppError(w, r, err, mapErrors)
		return
	}

	render.JSON(w, r, map[string]string{
		"message": "resource updated successfully",
	})
}
