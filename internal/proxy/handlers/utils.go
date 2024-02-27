package handlers

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"proxy_server/internal/logging"
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/utils"
	"strings"
)

func copyHeaders(source *http.Response, target http.ResponseWriter) {
	for headerKey, values := range source.Header {
		for _, headerValue := range values {
			target.Header().Add(headerKey, headerValue)
		}
	}
}

func parseValues(urlValues url.Values) map[string][]string {
	result := map[string][]string{}
	for key, values := range urlValues {
		result[key] = values
	}

	return result
}

func parseHeaders(headers http.Header) map[string][]string {
	result := map[string][]string{}
	for headerKey, values := range headers {
		result[headerKey] = values
	}

	return result
}

func parseCookies(cookies []*http.Cookie) map[string]string {
	result := map[string]string{}
	for _, cookie := range cookies {
		result[cookie.Name] = cookie.Value
	}

	return result
}

func setTarget(request *http.Request, target string) error {
	url := url.URL{
		Scheme:   "https",
		Host:     target,
		Path:     request.URL.Path,
		RawQuery: request.URL.RawQuery,
		Fragment: request.URL.Fragment,
	}

	request.URL = &url
	request.RequestURI = ""

	return nil
}

func checkType(header string) string {
	if strings.HasPrefix(header, "text/") ||
		(strings.HasPrefix(header, "application/") && !strings.HasSuffix(header, "octet-stream")) {
		return "text"
	} else {
		return "file"
	}
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

		obj.PostParams = parseValues(request.PostForm)
	}

	cookies := request.Cookies()
	request.Header.Del("Cookie")

	obj.Method = request.Method
	obj.Path = request.URL.Path
	obj.Scheme = request.URL.Scheme
	obj.Host = request.URL.Host
	obj.GetParams = parseValues(request.URL.Query())
	obj.Headers = parseHeaders(request.Header)
	obj.Cookies = parseCookies(cookies)

	switch checkType(request.Header.Get("Content-Type")) {
	case "text":
		obj.RawBody = rawBody
		obj.TextBody = string(rawBody)
	case "file":
		obj.RawBody = rawBody
	}

	return obj, nil
}

func responseToObj(response *http.Response, logger logging.ILogger) (dto.IncomingResponse, error) {
	obj := dto.IncomingResponse{}

	decodedBody, err := utils.DecodeResponse(response)
	if err != nil {
		logger.Error("Failed to decode response")
		return obj, err
	}

	obj.Code = response.StatusCode
	obj.Message = http.StatusText(response.StatusCode)
	obj.Headers = parseHeaders(response.Header)

	switch checkType(response.Header.Get("Content-Type")) {
	case "text":
		obj.RawBody = decodedBody
		obj.TextBody = string(decodedBody)
	case "file":
		obj.RawBody = decodedBody
	}

	return obj, nil
}
