package entities

import "proxy_server/internal/pkg/dto"

type Request struct {
	ID uint64 `json:"id"`
	dto.IncomingRequest
	Response Response `json:"response"`
}

type Response struct {
	ID        uint64 `json:"id"`
	RequestID uint64 `json:"request_id"`
	dto.IncomingResponse
}
