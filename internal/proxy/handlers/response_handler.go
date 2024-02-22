package handlers

import (
	"net/http"
	"proxy_server/internal/service"
)

type ResponseHandler struct {
	rs service.IResponseService
}

func (rh ResponseHandler) GetResponseService() service.IResponseService {
	return rh.rs
}

func (rh ResponseHandler) Serve(w http.ResponseWriter, r *http.Response) {

}
