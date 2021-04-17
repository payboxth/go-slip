package mock

import (
	"context"
	"image"
)

// SlipRepository is the mock struct for domain1 repository
type SlipStorage struct {
	StoreFileFunc      func(ctx context.Context, fileName, objectName string) (string, error)
	StoreByteFunc      func(ctx context.Context, b []byte, objectName string) (string, error)
	StoreOriginPNGFunc func(ctx context.Context, m image.Image, objectName string) (url string, err error)
	RemoveFileFunc     func(ctx context.Context, objectName string) error
}

// StoreFile calls SaveImageFunc
func (s *SlipStorage) StoreFile(ctx context.Context, fileName, objectName string) (string, error) {
	return s.StoreFileFunc(ctx, fileName, objectName)
}

// StoreByte calls StoreByteFunc
func (s *SlipStorage) StoreByte(ctx context.Context, b []byte, objectName string) (url string, err error) {
	return s.StoreByteFunc(ctx, b, objectName)
}

// StoreOriginPNG calls StoreOriginPNGFunc
func (s *SlipStorage) StoreOriginPNG(ctx context.Context, m image.Image, objectName string) (url string, err error) {
	return s.StoreOriginPNGFunc(ctx, m, objectName)
}

// RemoveFile call RemoveFileFunc
func (s *SlipStorage) RemoveFile(ctx context.Context, objectName string) error {
	return s.RemoveFileFunc(ctx, objectName)
}
