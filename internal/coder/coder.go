package coder

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

const (
	minSize = 32
	maxSize = 1024
	utf8BOM = "\ufeff"
)

type QRCode struct {
	Content string `json:"content"`
	Size    int    `json:"size"`
}

func (code *QRCode) Generate() ([]byte, error) {
	if err := code.Validate(); err != nil {
		return nil, fmt.Errorf("QR Code validation error: %v", err)
	}

	qrCode, err := qrcode.Encode(utf8BOM+code.Content, qrcode.Medium, code.Size)
	if err != nil {
		return nil, fmt.Errorf("could not generate a QR code: %v", err)
	}
	return qrCode, nil
}

func (code *QRCode) Validate() error {
	if code.Content == "" {
		return fmt.Errorf("content is empty")
	}
	if code.Size < minSize || code.Size > maxSize {
		return fmt.Errorf("invalid size: %d, must be greater than %d and less than %d", code.Size, minSize, maxSize)
	}
	return nil
}
