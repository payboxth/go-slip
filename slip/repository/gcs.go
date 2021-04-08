package repository

import (
	"context"

	"github.com/payboxth/go-slip/slip"
)

func NewGCS() slip.Storage {
	return &gcs{}
}

type gcs struct{}

func (gcs) SaveFile(ctx context.Context, file []byte) (path string, err error) {
	return "", nil
}
