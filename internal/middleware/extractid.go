package middleware

import (
	"context"
	"log"
	"net/http"
	"proxy_server/internal/apperrors"
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/utils"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

func ExtractID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := *utils.GetReqLogger(r.Context())
		if logger == nil {
			log.Fatal("Logger missing from context")
			apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		}
		logger.Info("Extracting request ID")

		funcName := "ExtractID"

		id := chi.URLParam(r, "requestID")

		rCtx := context.WithValue(r.Context(), dto.RequestIDKey, id)
		logger.DebugFmt("Added logger to context", chimw.GetReqID(rCtx), funcName, nodeName)

		logger.Info("*************** CONTEXT SET UP ***************")

		next.ServeHTTP(w, r.WithContext(rCtx))
	})
}
