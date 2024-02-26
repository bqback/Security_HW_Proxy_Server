package postgresql

var (
	allRequestInsertFields  = []string{"method", "scheme", "host", "path", "get", "headers", "cookies", "post"}
	allResponseInsertFields = []string{"code", "message", "headers", "body_raw", "body_text"}
	requestResponseFields   = []string{"id_request", "id_response"}
)
