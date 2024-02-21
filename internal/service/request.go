package service

import (
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/pkg/entities"
)

type IRequestService interface {
	StoreRequest(*dto.IncomingRequest) error
	GetAllRequests() (*[]entities.Request, error)
	GetSingleRequest(dto.RequestID) (*entities.Request, error)
}
