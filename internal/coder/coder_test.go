package coder

import (
	"testing"

	"github.com/etilite/qr-coder/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestQRCode_Generate(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		code := model.QRCode{
			Content: "test",
			Size:    32,
		}

		got, err := Encode(code)

		assert.NoError(t, err)
		assert.NotEmpty(t, got)
	})

	t.Run("encode error", func(t *testing.T) {
		t.Parallel()

		code := model.QRCode{
			Content: tooBigContent,
			Size:    32,
		}

		got, err := Encode(code)

		assert.Error(t, err)
		assert.Empty(t, got)
	})
}

const tooBigContent = `this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_
this_content_is_too_big_this_content_is_too_big_this_content_is_too_big_`
