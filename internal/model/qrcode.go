package model

import "fmt"

const (
	minSize = 32
	maxSize = 1024
)

type QRCode struct {
	Content string `json:"content"`
	Size    int    `json:"size"`
}

func (code *QRCode) Validate() error {
	if code.Content == "" {
		return fmt.Errorf("content is empty")
	}
	if code.Size < minSize || code.Size > maxSize {
		return fmt.Errorf("invalid size: %d, must be at least %d and at most %d", code.Size, minSize, maxSize)
	}
	return nil
}
