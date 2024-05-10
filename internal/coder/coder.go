package coder

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

const utf8BOM = "\ufeff"

type QRCode struct {
	Content string `json:"content"`
	Size    int    `json:"size"`
}

func (code *QRCode) Generate() ([]byte, error) {
	qrCode, err := qrcode.Encode(utf8BOM+code.Content, qrcode.Medium, code.Size)
	if err != nil {
		return nil, fmt.Errorf("could not generate a QR code: %v", err)
	}
	return qrCode, nil
}
