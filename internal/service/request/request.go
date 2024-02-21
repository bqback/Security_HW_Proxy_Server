package request

import (
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/pkg/entities"
	"proxy_server/internal/storage"
)

type RequestService struct {
	storage storage.IRequestStorage
}

func NewRequestService(requestStorage storage.IRequestStorage) *RequestService {
	return &RequestService{
		storage: requestStorage,
	}
}

func (rs RequestService) StoreRequest(*dto.IncomingRequest) error {
	// TODO Implement
	return nil
}

func (rs RequestService) GetAllRequests() (*[]entities.Request, error) {
	// TODO Implement
	return nil, nil
}

func (rs RequestService) GetSingleRequest(requestID dto.RequestID) (*entities.Request, error) {
	// TODO Implement
	return nil, nil
}
