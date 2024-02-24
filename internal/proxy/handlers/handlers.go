package handlers

import (
	"crypto/tls"
	"net/http"
	"proxy_server/internal/config"
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
func NewHandlers(services *service.Services, ca tls.Certificate, config *config.TLSConfig) *Handlers {
	return &Handlers{
		Http:  *NewHTTPHandler(services.Request, services.Response),
		Https: *NewHTTPSHandler(services.Request, services.Response, ca, config),
	}
}

// NewHTTPHandler
// возвращает HTTPHandler с необходимыми сервисами
func NewHTTPHandler(reqs service.IRequestService, resps service.IResponseService) *HTTPHandler {
	return &HTTPHandler{
		client: http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
		reqs:  reqs,
		resps: resps,
	}
}

// NewHTTPSHandler
// возвращает HTTPSHandler с необходимыми сервисами
func NewHTTPSHandler(reqs service.IRequestService, resps service.IResponseService, ca tls.Certificate, config *config.TLSConfig) *HTTPSHandler {
	return &HTTPSHandler{
		ca:        ca,
		tlsConfig: config,
		client: http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
		reqs:  reqs,
		resps: resps,
	}
}
