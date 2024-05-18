package http

import (
	"net/http"
)

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	handler := NewQRCodeHandler()

	mux.Handle("POST /generate", handler.handle())

	return mux
}
