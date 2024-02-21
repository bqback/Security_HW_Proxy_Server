package apperrors

import "errors"

var (
	ErrInvalidLoggingLevel = errors.New("Invalid logging level")
)

var (
	ErrEnvNotFound = errors.New("Unable to load .env file")
)

var (
	ErrDatabasePWMissing = errors.New("Database password is missing from env")
)
