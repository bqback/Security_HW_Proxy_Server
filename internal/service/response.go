package service

import (
	"context"
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/pkg/entities"
)

type IResponseService interface {
	StoreResponse(context.Context, *dto.IncomingResponse, *dto.RequestID) error
	GetResponseByRequestID(context.Context, *dto.RequestID) (*entities.Response, error)
	GetResponseByResponseID(context.Context, *dto.ResponseID) (*entities.Response, error)
}
