package handlers

import (
	"net/http"
	"proxy_server/internal/service"
)

type Handlers struct {
	Http  IProxyHandler
	Https IProxyHandler
}

type IProxyHandler interface {
	GetRequestService() service.IRequestService
	GetResponseService() service.IResponseService
	ServeHTTP(http.ResponseWriter, *http.Request)
}

const nodeName = "handler"

// NewHandlers
// возвращает HandlerManager со всеми хэндлерами приложения
func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		Http:  *NewHTTPHandler(services.Request, services.Response),
		Https: *NewHTTPSHandler(services.Request, services.Response),
	}
}

// NewHTTPHandler
// возвращает HTTPHandler с необходимыми сервисами
func NewHTTPHandler(reqs service.IRequestService, resps service.IResponseService) *HTTPHandler {
	return &HTTPHandler{
		client: http.Client{},
		reqs:   reqs,
		resps:  resps,
	}
}

// NewHTTPSHandler
// возвращает HTTPSHandler с необходимыми сервисами
func NewHTTPSHandler(reqs service.IRequestService, resps service.IResponseService) *HTTPSHandler {
	return &HTTPSHandler{
		client: http.Client{},
		reqs:   reqs,
		resps:  resps,
	}
}
