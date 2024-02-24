package repeat

import (
	"context"
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/storage"
)

type RepeatService struct {
	storage storage.IRequestStorage
}

func NewRepeatService(requestStorage storage.IRequestStorage) *RepeatService {
	return &RepeatService{
		storage: requestStorage,
	}
}

func (rs RepeatService) RepeatRequest(ctx context.Context, requestID dto.RequestID) error {
	// TODO Implement
	return nil
}
