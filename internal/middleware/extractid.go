package middleware

import (
	"context"
	"log"
	"net/http"
	"proxy_server/internal/apperrors"
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/utils"
	"strconv"

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

		id, err := strconv.ParseUint(chi.URLParam(r, "requestID"), 10, 64)
		if err != nil {
			logger.Error("Invalid request ID in url")
			apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
			return
		}

		rCtx := context.WithValue(r.Context(), dto.RequestIDKey, id)
		logger.DebugFmt("Added request ID to context", chimw.GetReqID(rCtx), funcName, nodeName)

		next.ServeHTTP(w, r.WithContext(rCtx))
	})
}
