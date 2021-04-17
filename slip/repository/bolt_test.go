package repository_test

import (
	"os"
	"testing"

	"github.com/payboxth/go-slip/slip/repository"
	"github.com/stretchr/testify/assert"
)

func TestNewBoltDB(t *testing.T) {
	boltdb, err := repository.NewBolt("slip.db")
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, boltdb, "boltdb should not nil")
	assert.FileExists(t, "slip.db", "Slip database file should be exists")
	// Teardown
	err = os.Remove("slip.db")
	if err != nil {
		t.Error(err)
	}
}

// TODO:
// create slipdata.json for mock data
