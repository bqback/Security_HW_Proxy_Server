package dto

import (
	"net/http"
	"net/url"
)

type IncomingRequest struct {
	Method     string
	Path       string
	GetParams  url.Values
	Headers    http.Header
	Cookies    []*http.Cookie
	PostParams url.Values
}

type IncomingResponse struct {
	Code    int
	Message string
	Headers http.Header
	Body    string
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
)
