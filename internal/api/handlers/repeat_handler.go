package handlers

import "proxy_server/internal/service"

type RepeatHandler struct {
	rs service.IRepeatService
}

func (rh RepeatHandler) GetRepeatService() service.IRepeatService {
	return rh.rs
}
