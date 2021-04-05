package slip

import "context"

type Service interface {
	Create(ctx context.Context, h *Head) (id string, url string, err error)
	FindByID(ctx context.Context, id string) (*Head, error)
}
