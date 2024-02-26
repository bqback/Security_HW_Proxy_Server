package handlers

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
	"proxy_server/internal/apperrors"
	"proxy_server/internal/config"
	proxytls "proxy_server/internal/proxy/tls"
	"proxy_server/internal/service"
	"proxy_server/internal/utils"

	chimw "github.com/go-chi/chi/v5/middleware"
)

type HTTPSHandler struct {
	client    http.Client
	ca        tls.Certificate
	tlsConfig *config.TLSConfig
	reqs      service.IRequestService
	resps     service.IResponseService
}

func (h HTTPSHandler) GetRequestService() service.IRequestService {
	return h.reqs
}

func (h HTTPSHandler) GetResponseService() service.IResponseService {
	return h.resps
}

func (h HTTPSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := *utils.GetReqLogger(ctx)
	requestID := chimw.GetReqID(ctx)
	funcName := "HTTPS Handler"

	r.Header.Del("Proxy-Connection")
	r.RequestURI = ""

	hj, ok := w.(http.Hijacker)
	if !ok {
		logger.Error("Failed to cast connection to hijacker")
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	conn, _, err := hj.Hijack()
	if err != nil {
		logger.Error("Failed to hijack connection: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("Connection hijacked", requestID, funcName, nodeName)

	_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	if err != nil {
		logger.Error("Failed to return 200 on CONNECT: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("Connection established", requestID, funcName, nodeName)

	tlsCert, err := proxytls.GenCert(r.URL.Hostname(), h.tlsConfig, logger)
	if err != nil {
		logger.Error("Failed to generate TLS certificate for host" + r.URL.Hostname() + ": " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("Certificate generated", requestID, funcName, nodeName)

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
		MinVersion:       tls.VersionTLS12,
		Certificates:     []tls.Certificate{tlsCert},
	}
	tlsConn := tls.Server(conn, tlsConfig)
	defer tlsConn.Close()

	err = tlsConn.Handshake()
	if err != nil {
		logger.Error("Failed to handshake: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	logger.DebugFmt("Handshake done", requestID, funcName, nodeName)

	tlsConnReader := bufio.NewReader(tlsConn)
	tlsRequest, err := http.ReadRequest(tlsConnReader)
	if err == io.EOF {
		return
	} else if err != nil {
		logger.Error("Failed to read request from client: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	reqBody, _ := io.ReadAll(r.Body)
	tlsRequest.Body = io.NopCloser(bytes.NewReader(reqBody))

	err = setTarget(tlsRequest, r.Host)
	if err != nil {
		logger.Error("Failed to modify request: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	response, err := http.DefaultClient.Do(tlsRequest)
	if err != nil {
		logger.Error("Failed to send request to server: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	logger.DebugFmt("Got response", requestID, funcName, nodeName)
	respBody, _ := io.ReadAll(response.Body)
	response.Body = io.NopCloser(bytes.NewReader(respBody))

	reqObj, err := requestToObj(tlsRequest, logger)
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

	rawResponseBody, _ := io.ReadAll(response.Body)
	response.Body = io.NopCloser(bytes.NewReader(rawResponseBody))

	respObj, err := responseToObj(response, logger)
	if err != nil {
		logger.Error("Failed to parse response into object: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
	err = h.resps.StoreResponse(ctx, &respObj, reqID)
	if err != nil {
		logger.Error("Failed to store response: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}

	response.Body = io.NopCloser(bytes.NewReader([]byte(rawResponseBody)))

	err = response.Write(tlsConn)
	if err != nil {
		logger.Error("Failed to return response to client: " + err.Error())
		apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)
		return
	}
}
