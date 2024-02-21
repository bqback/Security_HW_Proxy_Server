package storage

import (
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/pkg/entities"
)

type IResponseStorage interface {
	StoreResponse(*dto.IncomingResponse) error
	GetResponseByRequestID(*dto.RequestID) (*entities.Response, error)
	GetResponseByResponseID(*dto.ResponseID) (*entities.Response, error)
}
