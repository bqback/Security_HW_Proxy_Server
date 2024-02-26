package entities

import "proxy_server/internal/pkg/dto"

type Request struct {
	ID        uint64     `json:"id"`
	Responses []Response `json:"responses"`
	dto.IncomingRequest
}

type Response struct {
	ID        uint64 `json:"id"`
	RequestID uint64 `json:"request_id"`
	dto.IncomingResponse
}
