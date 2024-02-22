package handlers

import "proxy_server/internal/service"

type Handlers struct {
	ConnectHandler
	RequestHandler
	ResponseHandler
}

const nodeName = "handler"

// NewHandlers
// возвращает HandlerManager со всеми хэндлерами приложения
func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		ConnectHandler:  *NewConnectHandler(),
		RequestHandler:  *NewRequestHandler(services.Request),
		ResponseHandler: *NewResponseHandler(services.Response),
	}
}

// NewConnectHandler
// возвращает ConnectHandler с необходимыми сервисами
func NewConnectHandler() *ConnectHandler {
	return &ConnectHandler{}
}

// NewRequestHandler
// возвращает RequestHandler с необходимыми сервисами
func NewRequestHandler(reqs service.IRequestService) *RequestHandler {
	return &RequestHandler{
		rs: reqs,
	}
}

// NewResponseHandler
// возвращает ResponseHandler с необходимыми сервисами
func NewResponseHandler(reps service.IResponseService) *ResponseHandler {
	return &ResponseHandler{
		rs: reps,
	}
}
