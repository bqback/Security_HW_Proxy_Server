package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	chimw "github.com/go-chi/chi/v5/middleware"

	"proxy_server/internal/apperrors"
	"proxy_server/internal/pkg/dto"
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

	logger.DebugFmt("New request", requestID, funcName, nodeName)
	reqObj := dto.IncomingRequest{
		Method:     r.Method,
		Path:       r.URL.Path,
		GetParams:  r.URL.Query(),
		Headers:    r.Header,
		Cookies:    r.Cookies(),
		PostParams: r.PostForm,
	}
	reqID, err := h.reqs.StoreRequest(ctx, &reqObj)
	if err != nil {
		logger.Error("Failed to store request: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	r.Header.Del("Proxy-Connection")
	r.RequestURI = ""

	response, err := h.client.Do(r)
	if err != nil {
		logger.Error("Failed to send request to server: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	fmt.Println(r)

	logger.DebugFmt("Got response", requestID, funcName, nodeName)

	rawBody, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Error("Failed to read response body: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	respObj := dto.IncomingResponse{
		Code:    response.StatusCode,
		Message: http.StatusText(response.StatusCode),
		Headers: response.Header,
		Body:    string(rawBody),
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
