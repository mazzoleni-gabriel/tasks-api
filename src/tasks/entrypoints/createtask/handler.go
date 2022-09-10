package createtask

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"tasks-api/internal/render"
	"tasks-api/src/tasks"
)

//go:generate mockery --name=UseCase --disable-version-string
type UseCase interface {
	Create(ctx context.Context, task tasks.Task) (uint, error)
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
	r.Post("/tasks", l.Handle)
}

func (h Handler) Handle(w http.ResponseWriter, r *http.Request) {
	entity, err := newTaskFromRequest(r)
	if err != nil {
		render.NewError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	createdID, err := h.useCase.Create(r.Context(), entity)
	if err != nil {
		render.NewError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, CreateTaskResponse{ID: createdID})
}
