package searchtasks

import (
	"context"
	"github.com/go-chi/chi"
	"net/http"
	"tasks-api/internal/render"
	"tasks-api/src/tasks"
)

//go:generate mockery --name=UseCase --disable-version-string
type UseCase interface {
	Search(ctx context.Context, filters tasks.SearchFilters) ([]tasks.Task, error)
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
	r.Get("/tasks/search", l.Handle)
}

func (h Handler) Handle(w http.ResponseWriter, r *http.Request) {
	filters, err := newFiltersFromRequest(r)
	if err != nil {
		render.NewError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	tasks, err := h.useCase.Search(r.Context(), filters)
	if err != nil {
		render.NewError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, newResponse(tasks))
}
