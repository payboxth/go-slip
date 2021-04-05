package slip

import "context"

// Repository is the domain1 storage
type Repository interface {
	// Registers inserts given Entity1 into storage
	CreateSlip(ctx context.Context, entity *Slip) (ID string, err error)

	// SetField3 sets field3 for Entity1
	FindSlipByID(ctx context.Context, ID string) error
}
