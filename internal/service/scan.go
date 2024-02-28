package service

import (
	"context"
	"net/http"
	"proxy_server/internal/pkg/dto"
)

type IScanService interface {
	ScanRequest(context.Context, *dto.RequestID, http.Client, string) (map[string]int, error)
}
