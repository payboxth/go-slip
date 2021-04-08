package repository

import (
	"context"

	"github.com/payboxth/go-slip/slip"
)

func NewSqlite() slip.Database {
	return &sqlite{}
}

type sqlite struct{}

func (sqlite) Insert(ctx context.Context, b *slip.Body) (id string, err error) {
	return "", nil
}

func (sqlite) FindByID(ctx context.Context, id string) (*slip.Body, error) {
	return nil, nil
}
