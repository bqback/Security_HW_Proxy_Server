package response

import (
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/pkg/entities"
	"proxy_server/internal/storage"
)

type ResponseService struct {
	storage storage.IResponseStorage
}

func NewResponseService(ResponseStorage storage.IResponseStorage) *ResponseService {
	return &ResponseService{
		storage: ResponseStorage,
	}
}

func (rs ResponseService) StoreResponse(response *dto.IncomingResponse) error {
	// TODO Implement
	return nil
}

func (rs ResponseService) GetResponseByRequestID(requestID *dto.RequestID) (*entities.Response, error) {
	// TODO Implement
	return nil, nil
}
