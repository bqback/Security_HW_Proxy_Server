package apperrors

import (
	"errors"
	"net/http"
)

// const nodeName string = "error"

var (
	ErrInvalidLoggingLevel      = errors.New("invalid logging level")
	ErrLoggerMissingFromContext = errors.New("logger missing from context")
)

var (
	ErrEnvNotFound       = errors.New("unable to load .env file")
	ErrDatabasePWMissing = errors.New("database password is missing from env")
)

var (
	ErrCouldNotBuildQuery       = errors.New("failed to build SQL query")
	ErrCouldNotBeginTransaction = errors.New("failed to start DB transaction")
	ErrCouldNotRollback         = errors.New("failed to roll back after a failed query")
	ErrCouldNotCommit           = errors.New("failed to commit DB transaction changes")
)

var (
	ErrCouldNotGetRequest = errors.New("could not get request from DB")
)

type ErrorResponse struct {
	Code    int
	Message string
}

var BadRequestResponse = ErrorResponse{
	Code:    http.StatusBadRequest,
	Message: "Bad request",
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
