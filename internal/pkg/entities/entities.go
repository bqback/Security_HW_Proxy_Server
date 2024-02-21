package entities

import "proxy_server/internal/pkg/dto"

type Request struct {
	ID dto.RequestID
	dto.IncomingRequest
}

type Response struct {
	ID        dto.ResponseID
	RequestID dto.RequestID
	dto.IncomingResponse
}
