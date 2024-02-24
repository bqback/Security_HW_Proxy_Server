package response

import (
	"context"
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

func (rs ResponseService) StoreResponse(ctx context.Context, response *dto.IncomingResponse, requestID *dto.RequestID) error {
	return rs.storage.StoreResponse(ctx, response, requestID)
}

func (rs ResponseService) GetResponseByRequestID(ctx context.Context, requestID *dto.RequestID) (*entities.Response, error) {
	// TODO Implement
	return nil, nil
}

func (rs ResponseService) GetResponseByResponseID(ctx context.Context, responseID *dto.ResponseID) (*entities.Response, error) {
	// TODO Implement
	return nil, nil
}
