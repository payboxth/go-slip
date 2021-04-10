package service

import (
	"context"

	"github.com/payboxth/go-slip/slip"
)

// New creates new domain1 service
func New(repo slip.Database, storage slip.Storage) slip.Service {
	return &service{repo, storage}
}

type service struct {
	db      slip.Database
	storage slip.Storage
}

// Create is service create slip body data to repository
// and convert to slip image
// and save image to storage
func (s *service) Create(ctx context.Context, body *slip.Body) (string, string, error) {
	image, err := NewImage(body)
	if err != nil {
		return "", "", err
	}
	// Save image file to Storage
	path, err := s.storage.SaveFile(ctx, image)
	if err != nil {
		return "", "", err
	}
	// Save returned path to body.URL and insert data row to Database
	body.URL = path
	id, err := s.db.Insert(ctx, body)
	if err != nil {
		return "", "", err
	}
	return id, path, nil
}

func (s *service) FindByID(ctx context.Context, id string) (*slip.Body, error) {
	body, err := s.db.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func NewImage(body *slip.Body) ([]byte, error) {
	var image []byte
	//TODO implement this function
	// create html template and css
	// load slip data to html/template
	// generate and return image []byte in PNG format
	return image, nil
}
