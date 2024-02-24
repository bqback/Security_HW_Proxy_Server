package service

import (
	"context"
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/pkg/entities"
)

type IRequestService interface {
	StoreRequest(context.Context, *dto.IncomingRequest) (*dto.RequestID, error)
	GetAllRequests(context.Context) (*[]entities.Request, error)
	GetSingleRequest(context.Context, *dto.RequestID) (*entities.Request, error)
}
