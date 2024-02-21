package scan

import (
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/storage"
)

type ScanService struct {
	storage storage.IRequestStorage
}

func NewScanService(requestStorage storage.IRequestStorage) *ScanService {
	return &ScanService{
		storage: requestStorage,
	}
}

func (rs ScanService) ScanRequest(requestID dto.RequestID) error {
	// TODO Implement
	return nil
}
