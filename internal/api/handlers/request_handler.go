package handlers

import (
	"encoding/json"
	"net/http"
	"proxy_server/internal/apperrors"
	"proxy_server/internal/service"
	"proxy_server/internal/utils"

	chimw "github.com/go-chi/chi/v5/middleware"
)

type RequestHandler struct {
	rs service.IRequestService
}

func (rh RequestHandler) GetRequestService() service.IRequestService {
	return rh.rs
}

func (rh RequestHandler) GetAllRequests(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := *utils.GetReqLogger(ctx)
	requestID := chimw.GetReqID(ctx)
	funcName := "GetAllRequests"

	w.Header().Set("Content-Type", "application/json")

	requests, err := rh.rs.GetAllRequests(ctx)
	if err != nil {
		logger.Error("Failed to get requests: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("Got requests", requestID, funcName, nodeName)

	jsonResponse, err := json.Marshal(requests)
	if err != nil {
		logger.Error("Failed to marshal response: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	w.Write(jsonResponse)
	r.Body.Close()
}

func (rh RequestHandler) GetSingleRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := *utils.GetReqLogger(ctx)
	requestID := chimw.GetReqID(ctx)
	funcName := "GetAllRequests"

	proxyRequestID := utils.GetReqID(ctx)
	if proxyRequestID == nil {
		logger.Error("Failed to find request ID")
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	request, err := rh.rs.GetSingleRequest(ctx, proxyRequestID)
	if err != nil {
		logger.Error("Failed to get requests: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("Got requests", requestID, funcName, nodeName)

	jsonResponse, err := json.Marshal(request)
	if err != nil {
		logger.Error("Failed to marshal response: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	w.Write(jsonResponse)
	r.Body.Close()
}
