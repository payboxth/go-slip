package mock

import (
	"context"
)

// SlipRepository is the mock struct for domain1 repository
type SlipStorage struct {
	SaveImageFunc func(ctx context.Context, image []byte) (string, error)
}

// SaveImage calls SaveImageFunc
func (r *SlipStorage) SaveImage(ctx context.Context, image []byte) (string, error) {
	return r.SaveImageFunc(ctx, image)
}
