package mock

import (
	"context"

	"github.com/payboxth/go-slip/slip"
)

// SlipService mocks domain1 service
type SlipService struct {
	CreateFunc func(ctx context.Context, head *slip.Head) (id string, url string, err error)
	UpdateFunc func(ctx context.Context, id string) error
}

// Create calls CreateFunc
func (s *SlipService) Create(ctx context.Context, head *slip.Head) (id string, url string, err error) {
	return s.CreateFunc(ctx, head)
}

// Update calls UpdateFunc
func (s *SlipService) FindByID(ctx context.Context, id string) error {
	return s.UpdateFunc(ctx, id)
}
