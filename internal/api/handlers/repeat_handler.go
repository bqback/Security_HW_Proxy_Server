package handlers

import (
	"bytes"
	"io"
	"net/http"
	"proxy_server/internal/apperrors"
	"proxy_server/internal/service"
	"proxy_server/internal/utils"

	chimw "github.com/go-chi/chi/v5/middleware"
)

type RepeatHandler struct {
	rs     service.IRepeatService
	client http.Client
}

func (rh RepeatHandler) GetRepeatService() service.IRepeatService {
	return rh.rs
}

func (rh RepeatHandler) RepeatRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := *utils.GetReqLogger(ctx)
	requestID := chimw.GetReqID(ctx)
	funcName := "RepeatRequest"

	proxyRequestID := utils.GetReqID(ctx)
	if proxyRequestID == nil {
		logger.Error("Failed to find request ID")
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	response, err := rh.rs.RepeatRequest(ctx, proxyRequestID, rh.client)
	if err != nil {
		logger.Error("Failed to repeat request: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("Request repeated", requestID, funcName, nodeName)

	decodedBody, err := utils.DecodeResponse(response)
	if err != nil {
		logger.Error("Failed to repeat request: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	response.ContentLength = int64(len(decodedBody))
	response.Body = io.NopCloser(bytes.NewReader(decodedBody))

	err = response.Write(w)
	if err != nil {
		logger.Error("Failed to return response: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
}
