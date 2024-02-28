package handlers

import (
	"encoding/json"
	"net/http"
	"proxy_server/internal/apperrors"
	"proxy_server/internal/service"
	"proxy_server/internal/utils"

	chimw "github.com/go-chi/chi/v5/middleware"
)

type ScanHandler struct {
	ss           service.IScanService
	client       http.Client
	dictLocation string
}

func (sh ScanHandler) GetScanService() service.IScanService {
	return sh.ss
}

func (sh ScanHandler) ScanRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := *utils.GetReqLogger(ctx)
	requestID := chimw.GetReqID(ctx)
	funcName := "ScanRequest"

	proxyRequestID := utils.GetReqID(ctx)
	if proxyRequestID == nil {
		logger.Error("Failed to find request ID")
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	files, err := sh.ss.ScanRequest(ctx, proxyRequestID, sh.client, sh.dictLocation)
	if err != nil {
		logger.Error("Failed to repeat request: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("Request repeated", requestID, funcName, nodeName)

	jsonResponse, err := json.Marshal(files)
	if err != nil {
		logger.Error("Failed to marshal response: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.Error("Failed to return response: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
}
