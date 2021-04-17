package service_test

import (
	"context"
	"image"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/payboxth/go-slip/mock"
	"github.com/payboxth/go-slip/slip"
	"github.com/payboxth/go-slip/slip/service"
)

func TestServiceCreateDatabase(t *testing.T) {
	ctx := mock.Context{}
	db := mock.SlipDatabase{
		InsertFunc: func(ctx context.Context, body *slip.Body) (string, error) {
			assert.NotNil(t, body)
			return "abc", nil
		},
		FindByIDFunc: func(ctx context.Context, id string) (*slip.Body, error) {
			assert.NotNil(t, id)
			return nil, nil
		},
	}
	storage := mock.SlipStorage{
		StoreFileFunc: func(ctx context.Context, fileName, objectName string) (string, error) {
			assert.NotNil(t, fileName)
			return "path/to/image", nil
		},
		StoreByteFunc: func(ctx context.Context, b []byte, objectName string) (string, error) {

			return "path/to/image", nil
		},
		StoreOriginPNGFunc: func(ctx context.Context, m image.Image, objectName string) (url string, err error) {
			return "path/to/image", nil
		},
		RemoveFileFunc: func(ctx context.Context, objectName string) error {
			return nil
		},
	}

	s := service.New(&db, &storage)

	id, url, err := s.Create(&ctx, &slip.Body{})
	assert.NoError(t, err)
	assert.Equal(t, "abc", id)
	assert.Equal(t, "path/to/image", url)
}
