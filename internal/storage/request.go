package storage

import (
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/pkg/entities"
)

type IRequestStorage interface {
	StoreRequest(*dto.IncomingRequest) error
	GetRequestByID(*dto.RequestID) (*entities.Request, error)
	GetAllRequests() (*[]entities.Request, error)
}
