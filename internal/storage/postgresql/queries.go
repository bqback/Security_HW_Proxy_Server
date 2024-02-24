package postgresql

var (
	allRequestInsertFields  = []string{"method", "path", "get", "headers", "cookies", "post"}
	allResponseInsertFields = []string{"code", "message", "headers", "body"}
	requestResponseFields   = []string{"id_request", "id_response"}
)
