package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/etilite/qr-coder/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestQRCodeHandler_handle(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		handler := NewQRCodeHandler()
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

		handler := NewQRCodeHandler()
		handleFunc := handler.handle()

		requestBody := strings.NewReader("")
		request := httptest.NewRequest(http.MethodPost, "/generate", requestBody)
		response := httptest.NewRecorder()

		handleFunc(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Contains(t, response.Body.String(), "failed to decode JSON: EOF")
	})

	t.Run("validation error", func(t *testing.T) {
		t.Parallel()

		handler := NewQRCodeHandler()
		handleFunc := handler.handle()

		requestBody := strings.NewReader(`{"size": 32, "content": ""}`)
		request := httptest.NewRequest(http.MethodPost, "/generate", requestBody)
		response := httptest.NewRecorder()

		handleFunc(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Contains(t, response.Body.String(), "failed to validate QR Code: content is empty")
	})

	t.Run("internal server error", func(t *testing.T) {
		t.Parallel()

		handler := NewQRCodeHandler()
		handler.encode = func(code model.QRCode) ([]byte, error) {
			return nil, errors.New("encode error")
		}
		handleFunc := handler.handle()

		requestBody := strings.NewReader(`{"size": 32, "content": "test"}`)
		request := httptest.NewRequest(http.MethodPost, "/generate", requestBody)
		response := httptest.NewRecorder()

		handleFunc(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Contains(t, response.Body.String(), "failed to generate QR-code: encode error")
	})
}
