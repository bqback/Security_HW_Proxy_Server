package service

import (
	"proxy_server/internal/service/repeat"
	"proxy_server/internal/service/request"
	"proxy_server/internal/service/response"
	"proxy_server/internal/service/scan"
	"proxy_server/internal/storage"
)

type Services struct {
	Repeat   IRepeatService
	Request  IRequestService
	Response IResponseService
	Scan     IScanService
}

func NewServices(storages *storage.Storages) *Services {
	return &Services{
		Repeat:   repeat.NewRepeatService(storages.Request),
		Request:  request.NewRequestService(storages.Request),
		Response: response.NewResponseService(storages.Response),
		Scan:     scan.NewScanService(storages.Request),
	}
}
