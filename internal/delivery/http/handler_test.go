package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQRCodeHandler_handle(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		handler := &QRCodeHandler{}
		handleFunc := handler.handle()

		requestBody := strings.NewReader(`{"size": 32, "content": "test"}`)

		request := httptest.NewRequest(http.MethodPost, "/generate", requestBody)
		response := httptest.NewRecorder()

		handleFunc(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "image/png", response.Header().Get("Content-Type"))
	})

	t.Run("bad request", func(t *testing.T) {
		t.Parallel()

		handler := &QRCodeHandler{}
		handleFunc := handler.handle()

		requestBody := strings.NewReader("")
		request := httptest.NewRequest(http.MethodPost, "/generate", requestBody)
		response := httptest.NewRecorder()

		handleFunc(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Contains(t, response.Body.String(), "failed to decode JSON: EOF")
	})

	t.Run("internal server error", func(t *testing.T) {
		t.Parallel()

		handler := &QRCodeHandler{}
		handleFunc := handler.handle()

		requestBody := strings.NewReader(`{"size": 1000000, "content": ""}`)
		request := httptest.NewRequest(http.MethodPost, "/generate", requestBody)
		response := httptest.NewRecorder()

		handleFunc(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Contains(t, response.Body.String(), "failed to generate QR-code: QR Code validation error: content is empty")
	})
}
