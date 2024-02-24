package handlers

import "proxy_server/internal/service"

type ScanHandler struct {
	ss service.IScanService
}

func (sh ScanHandler) GetScanService() service.IScanService {
	return sh.ss
}
