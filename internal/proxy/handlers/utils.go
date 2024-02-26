package handlers

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"net/url"
	"proxy_server/internal/logging"
	"proxy_server/internal/pkg/dto"
)

func copyHeaders(source *http.Response, target http.ResponseWriter) {
	for headerKey, values := range source.Header {
		for _, headerValue := range values {
			target.Header().Add(headerKey, headerValue)
		}
	}
}

func setTarget(request *http.Request, target string) error {
	url := url.URL{
		Scheme:   "https",
		Host:     target,
		Path:     request.URL.Path,
		RawQuery: request.URL.RawQuery,
	}

	request.URL = &url
	request.RequestURI = ""

	return nil
}

func requestToObj(request *http.Request, logger logging.ILogger) (dto.IncomingRequest, error) {
	obj := dto.IncomingRequest{}

	rawBody, err := io.ReadAll(request.Body)
	if err != nil {
		logger.Error("Failed to read request body")
		return obj, err
	}
	defer request.Body.Close()

	if request.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		request.Body = io.NopCloser(bytes.NewReader(rawBody))
		err := request.ParseForm()
		if err != nil {
			logger.Error("Failed to parse form")
			return obj, err
		}
	}

	cookies := request.Cookies()
	request.Header.Del("Cookie")

	obj.Method = request.Method
	obj.Path = request.URL.Path
	obj.Scheme = request.URL.Scheme
	obj.Host = request.URL.Host
	obj.GetParams = request.URL.Query()
	obj.Headers = request.Header
	obj.Cookies = cookies
	obj.PostParams = request.PostForm
	obj.Body = string(rawBody)

	return obj, nil
}

func responseToObj(response *http.Response, logger logging.ILogger) (dto.IncomingResponse, error) {
	obj := dto.IncomingResponse{}

	rawBody, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Error("Failed to read response body")
		return obj, err
	}

	if !response.Uncompressed || (response.Header.Get("Content-Encoding") == "gzip") {
		gzipReader, err := gzip.NewReader(bytes.NewReader(rawBody))
		if err != nil {
			logger.Error("Failed to create a gzip reader")
			return obj, err
		}
		defer gzipReader.Close()

		rawBody, err = io.ReadAll(response.Body)
		if err != nil {
			logger.Error("Failed to read response body")
			return obj, err
		}
	}

	obj.Code = response.StatusCode
	obj.Message = http.StatusText(response.StatusCode)
	obj.Headers = response.Header
	obj.Body = string(rawBody)

	return obj, nil
}
