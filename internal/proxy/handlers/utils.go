package handlers

import (
	"net/http"
)

func copyHeaders(source *http.Response, target http.ResponseWriter) {
	for headerKey, values := range source.Header {
		for _, headerValue := range values {
			target.Header().Add(headerKey, headerValue)
		}
	}
}
