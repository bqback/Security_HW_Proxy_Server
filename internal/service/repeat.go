package service

import (
	"context"
	"net/http"
	"proxy_server/internal/pkg/dto"
)

type IRepeatService interface {
	RepeatRequest(context.Context, *dto.RequestID, http.Client) (*http.Response, error)
}
