package mock

import (
	"context"

	"github.com/payboxth/go-slip/slip"
)

// SlipDatabase is the mock struct for domain1 repository
type SlipDatabase struct {
	InsertFunc   func(ctx context.Context, body *slip.Body) (string, error)
	FindByIDFunc func(ctx context.Context, id string) (*slip.Body, error)
}

// Register calls RegisterFunc
func (r *SlipDatabase) Insert(ctx context.Context, body *slip.Body) (string, error) {
	return r.InsertFunc(ctx, body)
}

// FindByID calls FindByID func
func (r *SlipDatabase) FindByID(ctx context.Context, id string) (*slip.Body, error) {
	return r.FindByIDFunc(ctx, id)
}
