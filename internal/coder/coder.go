package coder

import (
	"fmt"

	"github.com/etilite/qr-coder/internal/model"
	"github.com/skip2/go-qrcode"
)

const (
	utf8BOM = "\ufeff"
)

func Encode(code model.QRCode) ([]byte, error) {
	qrCode, err := qrcode.Encode(utf8BOM+code.Content, qrcode.Medium, code.Size)
	if err != nil {
		return nil, fmt.Errorf("could not generate a QR code: %v", err)
	}
	return qrCode, nil
}
