package handlers

import "net/http"

type ConnectHandler struct {
}

func (ch ConnectHandler) Connect(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
