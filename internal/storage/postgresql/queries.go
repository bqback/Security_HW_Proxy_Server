package postgresql

var (
	allRequestInsertFields  = []string{"method", "scheme", "host", "path", "get", "headers", "cookies", "post", "body_raw", "body_text"}
	allResponseInsertFields = []string{"code", "message", "headers", "body_raw", "body_text"}
	requestResponseFields   = []string{"id_request", "id_response"}

	allRequestSelectFields = []string{"id", "method", "scheme", "host", "path", "get", "headers", "cookies", "post", "body_raw", "body_text"}
)
