package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type IncomingRequest struct {
	Method     string    `json:"method"`
	Path       string    `json:"path"`
	Scheme     string    `json:"scheme"`
	Host       string    `json:"host"`
	GetParams  SliceMap  `json:"get_params"`
	Headers    SliceMap  `json:"headers"`
	Cookies    StringMap `json:"cookies"`
	PostParams SliceMap  `json:"post_params"`
	Body
}

type SliceMap map[string][]string

func (m SliceMap) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *SliceMap) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &m)
}

type StringMap map[string]string

func (m StringMap) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *StringMap) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &m)
}

type IncomingResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Headers SliceMap `json:"headers"`
	Body
}

type Body struct {
	RawBody  []byte `json:"raw_body"`
	TextBody string `json:"text_body"`
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
