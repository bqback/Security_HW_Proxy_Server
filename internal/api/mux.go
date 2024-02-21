package api

import (
	"net/http"
	"proxy_server/internal/api/handlers"
	"proxy_server/internal/config"
	"proxy_server/internal/logging"

	"github.com/go-chi/chi/v5"
)

// GetChiMux
// возвращает mux, реализованный с помощью модуля chi
func GetChiMux(manager handlers.Handlers, config config.Config, logger logging.ILogger) (http.Handler, error) {
	mux := chi.NewRouter()

	// mux.Use(middleware.SetContext(*config.Server, logger))
	// mux.Use(middleware.PanicRecovery)
	// mux.Use(middleware.JsonHeader)

	mux.Route("/requests", func(r chi.Router) {
		r.Get("/", manager.RequestHandler.GetAllRequests)
		r.Get("/{id}", manager.RequestHandler.GetSingleRequest)
	})
	return mux, nil
}
