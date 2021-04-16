// +build integrate
package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	sliprepository "github.com/payboxth/go-slip/slip/repository"
)

func TestNewGCSClient(t *testing.T) {
	_, err := sliprepository.NewGCS("paybox_slip", "/Users/tom/secret/paybox_slip.json")
	if err != nil {
		t.Errorf("Repository cannot create Storage Client: %v", err)
	}
}

func TestSaveFile_URLmustContainPath(t *testing.T) {
	s, err := sliprepository.NewGCS("paybox_slip", "/Users/tom/secret/paybox_slip.json")
	if err != nil {
		t.Errorf("Repository cannot create Storage Client: %v", err)
	}
	url, err := s.SaveFile("test_slip.png")
	if err != nil {
		t.Fatalf("Error on s.SaveFile: %v", err)
	}
	assert.Containsf(t, url, "https://storage.googleapis.com/paybox_slip/public/", "Success storage save file and return URL = %v", url)
	// TODO Teardown by delete saved file

}
