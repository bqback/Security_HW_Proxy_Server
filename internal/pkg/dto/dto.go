package dto

import (
	"net/http"
	"net/url"
)

type IncomingRequest struct {
	Method     string
	Path       string
	Scheme     string
	Host       string
	GetParams  url.Values
	Headers    http.Header
	Cookies    []*http.Cookie
	PostParams url.Values
	Body       string
}

type IncomingResponse struct {
	Code     int
	Message  string
	Headers  http.Header
	RawBody  []byte
	TextBody string
}

type RequestID struct {
	Value uint64
}

type ResponseID struct {
	Value uint64
}

type key int

const (
	ErrorKey key = iota
	LoggerKey
	RequestIDKey
)
