package service

import (
	"context"
	"image"

	"github.com/payboxth/goslip"
)

type service struct {
	db      slip.Database
	storage slip.Storage
}

// New creates new domain1 service
func New(repo slip.Database, storage slip.Storage) slip.Service {
	return &service{repo, storage}
}

// Create is service create slip body data to repository
// and convert to slip image
// and save image to storage
func (s *service) Create(ctx context.Context, body *slip.Body) (string, string, error) {
	imageByte, err := NewImage(body)
	if err != nil {
		return "", "", err
	}
	// Save image file to Storage
	path := "image"
	url, err := s.storage.StoreOriginPNG(ctx, imageByte, path)
	if err != nil {
		return "", "", err
	}
	// Save returned path to body.URL and insert data row to Database
	body.ImageURL = url
	id, err := s.db.Insert(ctx, body)
	if err != nil {
		return "", "", err
	}
	return id, url, nil
}

// FindByID function to find slip by id. Return slip.Body and error
func (s *service) FindByID(ctx context.Context, id string) (*slip.Body, error) {
	body, err := s.db.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func NewImage(body *slip.Body) (image.Image, error) {
	var m image.Image
	//TODO implement this function
	// create html template and css
	// load slip data to html/template
	// generate and return image []byte in PNG format
	return m, nil
}
