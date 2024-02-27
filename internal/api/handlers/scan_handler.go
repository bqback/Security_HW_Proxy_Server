package handlers

import (
	"net/http"
	"proxy_server/internal/service"
)

type ScanHandler struct {
	ss service.IScanService
}

func (sh ScanHandler) GetScanService() service.IScanService {
	return sh.ss
}

func (sh ScanHandler) ScanRequest(w http.ResponseWriter, r *http.Request) {

}
