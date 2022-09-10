package server

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"go.uber.org/fx"
	"net/http"
	"tasks-api/internal/config"
)

const defaultAddr = ":8080"

func StartHTTPServer(l fx.Lifecycle, cfg config.Configuration, router *chi.Mux) {
	addr := defaultAddr
	if cfg.Addr != "" {
		addr = cfg.Addr
	}

	l.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			go func() {
				fmt.Println("############### Starting HTTP server")
				err := http.ListenAndServe(addr, router)
				if err != nil {
					panic(err)
				}
			}()
			return nil
		},
	})
}
