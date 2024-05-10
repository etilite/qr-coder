package http

import (
	"net/http"
)

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	handler := &QRCodeHandler{}

	mux.Handle("/generate", handler.handle())

	return mux
}
