package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQRCode_ValidateSuccess(t *testing.T) {
	t.Parallel()

	code := &QRCode{
		Content: "test",
		Size:    32,
	}

	err := code.Validate()
	assert.NoError(t, err)
}

func TestQRCode_ValidateErr(t *testing.T) {
	t.Parallel()

	type fields struct {
		content string
		size    int
	}
	tests := map[string]struct {
		fields fields
		want   string
	}{
		"empty content": {
			fields: fields{
				content: "",
				size:    32,
			},
			want: "content is empty",
		},
		"size is small": {
			fields: fields{
				content: "test",
				size:    -1,
			},
			want: "invalid size: -1, must be at least 32 and at most 1024",
		},
		"size is big": {
			fields: fields{
				content: "test",
				size:    1025,
			},
			want: "invalid size: 1025, must be at least 32 and at most 1024",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			code := &QRCode{
				Content: tt.fields.content,
				Size:    tt.fields.size,
			}

			err := code.Validate()

			assert.Error(t, err)
			assert.ErrorContains(t, err, tt.want)
		})
	}
}
