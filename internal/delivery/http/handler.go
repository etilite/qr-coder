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
	return func(writer http.ResponseWriter, request *http.Request) {
		qrCode := coder.QRCode{}
		err := json.NewDecoder(request.Body).Decode(&qrCode)

		//writer.Header().Set("Content-Type", "application/json")

		if err != nil {
			err = fmt.Errorf("failed to decode JSON: %w", err)
			log.Print(err)
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		//var image []byte
		image, err := qrCode.Generate()
		if err != nil {
			writer.WriteHeader(400)
			json.NewEncoder(writer).Encode(
				fmt.Sprintf("Could not generate QR code. %v", err),
			)
			return
		}

		writer.Header().Set("Content-Type", "image/png")
		writer.Write(image)
	}
}
