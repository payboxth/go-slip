package mock

import (
	"context"

	"github.com/payboxth/go-slip/slip"
)

// SlipRepository is the mock struct for domain1 repository
type SlipRepository struct {
	CreateFunc   func(ctx context.Context, b *slip.Body) (string, error)
	FindByIDFunc func(ctx context.Context, id string) (*slip.Body, error)
}

// Register calls RegisterFunc
func (r *SlipRepository) Create(ctx context.Context, b *slip.Body) (string, error) {
	return r.CreateFunc(ctx, b)
}

// FindByID calls FindByID func
func (r *SlipRepository) FindByID(ctx context.Context, id string) (*slip.Body, error) {
	return r.FindByIDFunc(ctx, id)
}
