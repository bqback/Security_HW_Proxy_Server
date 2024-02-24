package storage

import (
	"context"
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/pkg/entities"
)

type IRequestStorage interface {
	StoreRequest(context.Context, *dto.IncomingRequest) (*dto.RequestID, error)
	GetRequestByID(context.Context, *dto.RequestID) (*entities.Request, error)
	GetAllRequests(context.Context) (*[]entities.Request, error)
}
