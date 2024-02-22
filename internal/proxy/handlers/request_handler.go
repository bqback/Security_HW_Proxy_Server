package handlers

import (
	"net/http"
	"proxy_server/internal/service"
)

type RequestHandler struct {
	rs service.IRequestService
}

func (rh RequestHandler) GetRequestService() service.IRequestService {
	return rh.rs
}

func (rh RequestHandler) Serve(w http.ResponseWriter, r *http.Request) {

}
