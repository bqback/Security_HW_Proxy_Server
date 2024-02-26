package middleware

import (
	"context"
	"net/http"
	"proxy_server/internal/logging"
	"proxy_server/internal/pkg/dto"

	chimw "github.com/go-chi/chi/v5/middleware"
)

const nodeName = "middleware"

func SetContext(logger logging.ILogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("*************** SETTING UP CONTEXT ***************")

			funcName := "SetLogger"

			rCtx := context.WithValue(r.Context(), dto.LoggerKey, logger)
			logger.DebugFmt("Added logger to context", chimw.GetReqID(rCtx), funcName, nodeName)

			logger.Info("*************** CONTEXT SET UP ***************")

			next.ServeHTTP(w, r.WithContext(rCtx))
		})
	}
}
