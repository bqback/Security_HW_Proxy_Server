package service

import "proxy_server/internal/pkg/dto"

type IScanService interface {
	ScanRequest(dto.RequestID) error
}
