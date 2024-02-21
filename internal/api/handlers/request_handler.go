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

func (rh RequestHandler) GetAllRequests(w http.ResponseWriter, r *http.Request) {

}

func (rh RequestHandler) GetSingleRequest(w http.ResponseWriter, r *http.Request) {

}
