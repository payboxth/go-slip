package repository

import (
	"context"

	"github.com/payboxth/go-slip/slip"
)

func NewBolt() slip.Repository {
	return &bolt{}
}

type bolt struct{}

func (bolt) Create(ctx context.Context, b *slip.Body) (id string, err error) {
	return "", nil
}

func (bolt) FindByID(ctx context.Context, id string) (*slip.Body, error) {
	return nil, nil
}
