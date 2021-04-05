package service

import (
	"context"

	"github.com/payboxth/go-slip/slip"
)

// New creates new domain1 service
func New(repo slip.Repository, storage slip.Storage) slip.Service {
	return &service{repo, storage}
}

type service struct {
	repo    slip.Repository
	storage slip.Storage
}

func (s *service) Create(ctx context.Context, sl *slip.Head) (string, string, error) {
	id, err := s.repo.Create(ctx, sl)
	if err != nil {
		return "", "", err
	}
	image, err := NewImage(sl)
	path, err := s.storage.SaveImage(ctx, image)
	sl.URL = path
	return id, path, nil
}

func (s *service) FindByID(ctx context.Context, id string) (*slip.Head, error) {
	sl, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return sl, nil
}

func NewImage(sl *slip.Head) ([]byte, error) {
	var image []byte
	//TODO implement this function
	// create html template and css
	// load slip data to html/template
	// generate and return image []byte in PNG format
	return image, nil
}
