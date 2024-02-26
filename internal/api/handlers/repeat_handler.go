package handlers

import (
	"net/http"
	"proxy_server/internal/service"
)

type RepeatHandler struct {
	rs service.IRepeatService
}

func (rh RepeatHandler) GetRepeatService() service.IRepeatService {
	return rh.rs
}

func (rh RepeatHandler) RepeatRequest(w http.ResponseWriter, r *http.Request) {

}
