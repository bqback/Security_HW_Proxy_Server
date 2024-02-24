package handlers

import (
	"net/http"
	"proxy_server/internal/service"
)

type HTTPSHandler struct {
	client http.Client
	reqs   service.IRequestService
	resps  service.IResponseService
}

func (h HTTPSHandler) GetRequestService() service.IRequestService {
	return h.reqs
}

func (h HTTPSHandler) GetResponseService() service.IResponseService {
	return h.resps
}

func (h HTTPSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
