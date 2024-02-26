package api

import (
	"net/http"
	"proxy_server/internal/api/handlers"
	"proxy_server/internal/config"
	"proxy_server/internal/logging"
	"proxy_server/internal/middleware"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

// GetChiMux
// возвращает mux, реализованный с помощью модуля chi
func GetChiMux(manager handlers.Handlers, config config.Config, logger logging.ILogger) (http.Handler, error) {
	mux := chi.NewRouter()

	mux.Use(chimw.RequestID)
	mux.Use(middleware.SetContext(logger))
	mux.Use(middleware.PanicRecovery)
	// mux.Use(middleware.JsonHeader)

	mux.Route("/", func(r chi.Router) {
		r.Route("/request", func(r chi.Router) {
			r.Get("/", manager.RequestHandler.GetAllRequests)
			r.Route("/{requestID}", func(r chi.Router) {
				r.Use(middleware.ExtractID)
				r.Get("/", manager.RequestHandler.GetSingleRequest)
			})
		})
		r.Route("/repeat/{requestID}", func(r chi.Router) {
			r.Use(middleware.ExtractID)
			r.Get("/", manager.RepeatHandler.RepeatRequest)
		})
		r.Route("/scan/{requestID}", func(r chi.Router) {
			r.Use(middleware.ExtractID)
			r.Get("/", manager.ScanHandler.ScanRequest)
		})
	})

	return mux, nil
}
