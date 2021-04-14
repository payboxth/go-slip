package repository_test

import (
	"testing"

	"github.com/payboxth/go-slip/slip/repository"
)

func TestNewGCSClient(t *testing.T) {
	gcs := repository.NewGCS()
	if gcs == nil {
		t.Errorf("Repository cannot create Storage Client: %v", gcs)
	}
}
