package utils

import (
	"context"
	"proxy_server/internal/logging"
	"proxy_server/internal/pkg/dto"
)

func GetReqLogger(ctx context.Context) *logging.ILogger {
	if ctx == nil {
		return nil
	}
	if logger, ok := ctx.Value(dto.LoggerKey).(logging.ILogger); ok {
		return &logger
	}
	return nil
}

func GetReqID(ctx context.Context) *dto.RequestID {
	if ctx == nil {
		return nil
	}
	if id, ok := ctx.Value(dto.RequestIDKey).(uint64); ok {
		return &dto.RequestID{Value: id}
	}
	return nil
}
