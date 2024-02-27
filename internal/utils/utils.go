package utils

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"net/url"
	"proxy_server/internal/logging"
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/pkg/entities"

	"github.com/andybalholm/brotli"
)

func GetReqLogger(ctx context.Context) *logging.ILogger {
	if ctx == nil {
		return nil
	}
	if logger, ok := ctx.Value(dto.LoggerKey).(logging.ILogger); ok {
		return &logger
	}
	return nil
}

func GetReqID(ctx context.Context) *dto.RequestID {
	if ctx == nil {
		return nil
	}
	if id, ok := ctx.Value(dto.RequestIDKey).(uint64); ok {
		return &dto.RequestID{Value: id}
	}
	return nil
}

func ObjToRequest(obj *entities.Request) (*http.Request, error) {
	reqUrl, err := url.Parse(obj.Scheme + "://" + obj.Host + obj.Path)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(obj.Method, reqUrl.String(), bytes.NewReader(obj.RawBody))
	if err != nil {
		return nil, err
	}

	request.Header = reverseParseHeaders(obj.Headers)
	reverseParseCookies(obj.Cookies, request)
	request.URL.RawQuery = reverseParseValues(obj.GetParams).Encode()

	return request, nil
}

func reverseParseHeaders(headersMap map[string][]string) http.Header {
	headers := http.Header{}
	for header, values := range headersMap {
		for _, value := range values {
			headers.Add(header, value)
		}
	}

	return headers
}

func reverseParseCookies(cookiesMap map[string]string, request *http.Request) {
	for cookie, value := range cookiesMap {
		request.AddCookie(&http.Cookie{Name: cookie, Value: value})
	}
}

func reverseParseValues(valueMap map[string][]string) url.Values {
	urlValues := url.Values{}
	for key, values := range valueMap {
		for _, value := range values {
			urlValues.Add(key, value)
		}
	}

	return urlValues
}

func DecodeResponse(response *http.Response) ([]byte, error) {
	decodedBody := []byte{}
	var err error
	var reader io.Reader

	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			return decodedBody, err
		}
	case "br":
		reader = brotli.NewReader(response.Body)
	default:
		reader = response.Body
	}

	decodedBody, err = io.ReadAll(reader)
	if err != nil {
		return decodedBody, err
	}

	return decodedBody, nil
}
