package mock

import (
	"context"

	"github.com/payboxth/go-slip/slip"
)

// SlipService mocks domain1 service
type SlipService struct {
	CreateFunc   func(ctx context.Context, b *slip.Body) (id string, url string, err error)
	FindByIDFunc func(ctx context.Context, id string) (*slip.Body, error)
}

// Create calls CreateFunc
func (s *SlipService) Create(ctx context.Context, b *slip.Body) (id string, url string, err error) {
	return s.CreateFunc(ctx, b)
}

// Update calls UpdateFunc
func (s *SlipService) FindByID(ctx context.Context, id string) (*slip.Body, error) {
	return s.FindByIDFunc(ctx, id)
}
