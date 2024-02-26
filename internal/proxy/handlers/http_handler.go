package handlers

import (
	"bytes"
	"io"
	"net/http"

	chimw "github.com/go-chi/chi/v5/middleware"

	"proxy_server/internal/apperrors"
	"proxy_server/internal/service"
	"proxy_server/internal/utils"
)

type HTTPHandler struct {
	client http.Client
	reqs   service.IRequestService
	resps  service.IResponseService
}

func (h HTTPHandler) GetRequestService() service.IRequestService {
	return h.reqs
}

func (h HTTPHandler) GetResponseService() service.IResponseService {
	return h.resps
}

func (h HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := *utils.GetReqLogger(ctx)
	requestID := chimw.GetReqID(ctx)
	funcName := "HTTP Handler"

	r.Header.Del("Proxy-Connection")
	r.RequestURI = ""

	logger.DebugFmt("New request", requestID, funcName, nodeName)

	response, err := h.client.Do(r)
	if err != nil {
		logger.Error("Failed to send request to server: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	logger.DebugFmt("Got response", requestID, funcName, nodeName)

	reqObj, err := requestToObj(r, logger)
	if err != nil {
		logger.Error("Failed to parse request into object: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	reqID, err := h.reqs.StoreRequest(ctx, &reqObj)
	if err != nil {
		logger.Error("Failed to store request: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	respObj, err := responseToObj(response, logger)
	if err != nil {
		logger.Error("Failed to parse request into object: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	err = h.resps.StoreResponse(ctx, &respObj, reqID)
	if err != nil {
		logger.Error("Failed to store response: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	w.WriteHeader(respObj.Code)
	copyHeaders(response, w)
	_, err = io.Copy(w, bytes.NewReader([]byte(respObj.Body)))
	if err != nil {
		logger.Error("Failed to copy response body: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
}
