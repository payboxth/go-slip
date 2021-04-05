package endpoint_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/payboxth/go-slip/mock"
	"github.com/payboxth/go-slip/slip"
	"github.com/payboxth/go-slip/slip/endpoint"
)

func TestCreate(t *testing.T) {
	ctx := mock.Context{}
	s := mock.SlipService{
		CreateFunc: func(ctx context.Context, head *slip.Head) (id string, err error) {
			// TODO:
			// assert.Equal(t, "field1_data", head.Field2.Field1)
			return "abc", nil
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
