package handlers

import "proxy_server/internal/service"

type Handlers struct {
	RequestHandler
	RepeatHandler
	ScanHandler
}

const nodeName = "handler"

// NewHandlers
// возвращает HandlerManager со всеми хэндлерами приложения
func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		RequestHandler: *NewRequestHandler(services.Request),
		RepeatHandler:  *NewRepeatHandler(services.Repeat),
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
func NewRepeatHandler(reps service.IRepeatService) *RepeatHandler {
	return &RepeatHandler{
		rs: reps,
	}
}

// NewScanHandler
// возвращает ScanHandler с необходимыми сервисами
func NewScanHandler(scans service.IScanService) *ScanHandler {
	return &ScanHandler{
		ss: scans,
	}
}
