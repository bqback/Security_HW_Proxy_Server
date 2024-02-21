package service

import "proxy_server/internal/pkg/dto"

type IRepeatService interface {
	RepeatRequest(dto.RequestID) error
}
