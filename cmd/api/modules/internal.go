package modules

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/httplog"
	"go.uber.org/fx"
	"net/http"
	"tasks-api/internal/db"
	"tasks-api/internal/server"
)

var internalModule = fx.Options(
	fx.Provide(
		newRouter,
		db.NewDatabase,
	),
	fx.Invoke(
		server.StartHTTPServer,
	),
)

func newRouter() *chi.Mux {
	logger := httplog.NewLogger("httplog-example", httplog.Options{
		JSON: true,
	})

	router := chi.NewRouter()
	router.Use(httplog.RequestLogger(logger))

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	return router
}
