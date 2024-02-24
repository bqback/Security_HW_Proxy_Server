package service

import (
	"context"
	"proxy_server/internal/pkg/dto"
)

type IRepeatService interface {
	RepeatRequest(context.Context, dto.RequestID) error
}
