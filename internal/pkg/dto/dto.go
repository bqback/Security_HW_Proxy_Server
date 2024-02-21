package dto

type IncomingRequest struct {
	Method     string
	Path       string
	GetParams  map[string]interface{}
	Headers    map[string]string
	Cookies    map[string]interface{}
	PostParams map[string]interface{}
}

type IncomingResponse struct {
	Code    int
	Message string
	Headers map[string]string
	Body    string
}

type RequestID struct {
	Value uint64
}

type ResponseID struct {
	Value uint64
}
