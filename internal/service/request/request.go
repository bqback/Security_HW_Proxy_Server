package request

import (
	"context"
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

func (rs RequestService) StoreRequest(ctx context.Context, request *dto.IncomingRequest) (*dto.RequestID, error) {
	return rs.storage.StoreRequest(ctx, request)
}

func (rs RequestService) GetAllRequests(ctx context.Context) (*[]entities.Request, error) {
	return rs.storage.GetAllRequests(ctx)
}

func (rs RequestService) GetSingleRequest(ctx context.Context, requestID *dto.RequestID) (*entities.Request, error) {
	return rs.storage.GetRequestByID(ctx, requestID)
}
