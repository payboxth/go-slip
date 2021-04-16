package slip

import (
	"context"
)

// Database is the slip database interface
type Database interface {
	// Registers inserts given Entity1 into storage
	Insert(ctx context.Context, b *Body) (id string, err error)

	// SetField3 sets field3 for Entity1
	FindByID(ctx context.Context, id string) (*Body, error)
}

// Storage is the slip storage interface for save image file
type Storage interface {
	SaveFile(file, path string) (url string, err error)
}
