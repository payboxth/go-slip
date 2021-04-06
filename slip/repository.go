package slip

import "context"

// Repository is the domain1 storage
type Repository interface {
	// Registers inserts given Entity1 into storage
	Create(ctx context.Context, b *Body) (id string, err error)

	// SetField3 sets field3 for Entity1
	FindByID(ctx context.Context, id string) (*Body, error)
}
