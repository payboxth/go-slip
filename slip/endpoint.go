package slip

import "context"

// Endpoint is the slip endpoint
type Endpoint interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	FindByID(context.Context, *FindByIDRequest) (*FindByIDResponse, error)
}

// Create
type (
	// CreateRequest is the request for create endpoint
	CreateRequest struct {
		DocNumber   string `json:"doc_number"`
		RefNumber   string `json:"ref_number"`
		Title       string `json:"title"`
		DocDate     string `json:"doc_date"`
		SlipLines   []Line
		AccessToken string `json:"access_token"`
	}

	// CreateResponse is the response for create endpoint
	CreateResponse struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	}
)

// Find
type (
	// FindRequest is the request for find endpoint
	FindByIDRequest struct {
		ID string `json:"id"`
	}

	// FindByIDRespose is the response for find by id endpoint
	FindByIDResponse struct {
		URL string `json:"id"`
	}
)
