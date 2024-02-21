package service

import (
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/pkg/entities"
)

type IResponseService interface {
	StoreResponse(*dto.IncomingResponse) error
	GetResponseByRequestID(*dto.RequestID) (*entities.Response, error)
	// GetAllRequests() (*[]entities.Request, error)
	// GetSingleRequest(dto.RequestID) (*entities.Request, error)
}
