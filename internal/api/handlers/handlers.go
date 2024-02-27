package handlers

import (
	"net/http"
	"proxy_server/internal/config"
	"proxy_server/internal/service"
)

type Handlers struct {
	RequestHandler
	RepeatHandler
	ScanHandler
}

const nodeName = "handler"

// NewHandlers
// возвращает HandlerManager со всеми хэндлерами приложения
func NewHandlers(services *service.Services, config *config.Config) *Handlers {
	return &Handlers{
		RequestHandler: *NewRequestHandler(services.Request),
		RepeatHandler:  *NewRepeatHandler(services.Repeat, config),
		ScanHandler:    *NewScanHandler(services.Scan),
	}
}

// NewRequestHandler
// возвращает RequestHandler с необходимыми сервисами
func NewRequestHandler(reqs service.IRequestService) *RequestHandler {
	return &RequestHandler{
		rs: reqs,
	}
}

// NewRepeatHandler
// возвращает RepeatHandler с необходимыми сервисами
func NewRepeatHandler(reps service.IRepeatService, config *config.Config) *RepeatHandler {
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(config.Proxy.URL),
		},
	}
	return &RepeatHandler{
		rs:     reps,
		client: client,
	}
}

// NewScanHandler
// возвращает ScanHandler с необходимыми сервисами
func NewScanHandler(scans service.IScanService) *ScanHandler {
	return &ScanHandler{
		ss: scans,
	}
}
