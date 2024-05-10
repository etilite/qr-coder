package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/etilite/qr-coder/internal/coder"
)

type QRCodeHandler struct {
}

func (h *QRCodeHandler) handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := r.Body.Close(); err != nil {
				log.Printf("failed to close request body %v", err)
			}
		}()

		qrCode := coder.QRCode{}

		if err := json.NewDecoder(r.Body).Decode(&qrCode); err != nil {
			err = fmt.Errorf("failed to decode JSON: %w", err)
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		image, err := qrCode.Generate()
		if err != nil {
			err = fmt.Errorf("failed to generate QR-code: %w", err)
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		if _, err := w.Write(image); err != nil {
			err = fmt.Errorf("failed to decode JSON: %w", err)
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
