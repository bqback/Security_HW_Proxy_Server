package service

import (
	"context"
	"proxy_server/internal/pkg/dto"
)

type IScanService interface {
	ScanRequest(context.Context, dto.RequestID) error
}
