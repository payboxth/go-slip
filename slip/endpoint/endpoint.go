package endpoint

import (
	"context"

	"github.com/payboxth/go-slip/slip"
)

// New creates new domain1 endpoint
func New(s slip.Service) slip.Endpoint {
	return &endpoint{s}
}

type endpoint struct {
	s slip.Service
}

func (ep *endpoint) Create(ctx context.Context, req *slip.CreateRequest) (*slip.CreateResponse, error) {
	id, url, err := ep.s.Create(ctx, &slip.Body{})
	if err != nil {
		return nil, err
	}
	return &slip.CreateResponse{
		ID:  id,
		URL: url,
	}, nil
}

func (ep *endpoint) FindByID(ctx context.Context, req *slip.FindByIDRequest) (*slip.FindByIDResponse, error) {
	h, err := ep.s.FindByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &slip.FindByIDResponse{URL: h.ImageURL}, nil
}
