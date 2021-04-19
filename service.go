package slip

import "context"

type Service interface {
	Create(ctx context.Context, b *Body) (id string, url string, err error)
	FindByID(ctx context.Context, id string) (*Body, error)
}
