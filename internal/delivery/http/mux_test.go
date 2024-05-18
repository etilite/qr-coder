package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMux(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		method string
		path   string
	}{
		"generate": {
			method: http.MethodPost,
			path:   "/generate",
		},
	}
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mux := newMux()

			requestBody := strings.NewReader(`{"size": 32, "content": "test"}`)
			request := httptest.NewRequest(tt.method, tt.path, requestBody)
			response := httptest.NewRecorder()

			mux.ServeHTTP(response, request)

			assert.Equal(t, http.StatusOK, response.Code)
		})
	}
}
