package slip

import "context"

// Repository is the domain1 storage
type Repository interface {
	// Registers inserts given Entity1 into storage
	Create(ctx context.Context, slip *Head) (id string, err error)

	// SetField3 sets field3 for Entity1
	FindSlipByID(ctx context.Context, ID string) (*Head, error)
}

type Storage interface {
	StoreImage(ctx context.Context, file []byte) (path string, err error)
}
