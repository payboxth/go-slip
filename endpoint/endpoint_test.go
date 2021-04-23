package endpoint_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/payboxth/goslip"
	"github.com/payboxth/goslip/endpoint"
	"github.com/payboxth/goslip/mock"
)

func TestCreate(t *testing.T) {
	ctx := mock.Context{}
	s := mock.SlipService{
		CreateFunc: func(ctx context.Context, b *slip.Body) (id string, url string, err error) {
			// TODO:
			// assert.Equal(t, "field1_data", head.Field2.Field1)
			return "abc", "path/to/slip", nil
		},
		FindByIDFunc: func(ctx context.Context, id string) (*slip.Body, error) {
			// TODO:
			b := &slip.Body{}
			return b, nil
		},
	}

	ep := endpoint.New(&s)
	resp, err := ep.Create(&ctx, &slip.CreateRequest{
		// TODO:	Field1: "field1_data",
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "abc", resp.ID)
}
