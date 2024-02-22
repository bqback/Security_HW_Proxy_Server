package proxy

import (
	"net/http"
	"proxy_server/internal/config"
	"proxy_server/internal/logging"
	"proxy_server/internal/middleware"
	"proxy_server/internal/proxy/handlers"

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
		r.Connect("/", manager.ConnectHandler.Connect)
	})
	return mux, nil
}
