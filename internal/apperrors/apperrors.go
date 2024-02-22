package apperrors

import (
	"errors"
	"net/http"
)

// const nodeName string = "error"

var (
	ErrInvalidLoggingLevel = errors.New("invalid logging level")
)

var (
	ErrEnvNotFound = errors.New("unable to load .env file")
)

var (
	ErrDatabasePWMissing = errors.New("database password is missing from env")
)

type ErrorResponse struct {
	Code    int
	Message string
}

var InternalServerErrorResponse = ErrorResponse{
	Code:    http.StatusInternalServerError,
	Message: "Internal error",
}

func ReturnError(err ErrorResponse, w http.ResponseWriter, r *http.Request) {
	// rCtx := r.Context()
	// logger := *utils.GetReqLogger(rCtx)
	// requestID := chimw.GetReqID(rCtx)
	// funcName := "ReturnError"

	w.WriteHeader(err.Code)
	// logger.DebugFmt(fmt.Sprintf("Wrote error code %v", err.Code), requestID, funcName, nodeName)
	_, _ = w.Write([]byte(err.Message))
	// logger.DebugFmt(fmt.Sprintf("Wrote message %v", err.Code), requestID, funcName, nodeName)
	r.Body.Close()
	// logger.DebugFmt("Request body closed", requestID, funcName, nodeName)
}
