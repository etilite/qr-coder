package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/etilite/qr-coder/internal/coder"
	"github.com/etilite/qr-coder/internal/model"
)

type QRCodeHandler struct {
	encode func(code model.QRCode) ([]byte, error)
}

func NewQRCodeHandler() *QRCodeHandler {
	return &QRCodeHandler{encode: coder.Encode}
}

func (h *QRCodeHandler) handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := r.Body.Close(); err != nil {
				log.Printf("failed to close request body %v", err)
			}
		}()

		qrCode := model.QRCode{}

		if err := json.NewDecoder(r.Body).Decode(&qrCode); err != nil {
			err = fmt.Errorf("failed to decode JSON: %w", err)
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := qrCode.Validate(); err != nil {
			err = fmt.Errorf("failed to validate QR Code: %v", err)
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		image, err := h.encode(qrCode)
		if err != nil {
			err = fmt.Errorf("failed to generate QR-code: %w", err)
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		if _, err := w.Write(image); err != nil {
			log.Print(fmt.Errorf("failed to write response: %w", err))
		}
	}
}
