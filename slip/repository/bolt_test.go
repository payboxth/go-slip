package repository_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/payboxth/go-slip/slip/repository"
	"github.com/stretchr/testify/assert"
)

var (
	dbFile string = "/slip.db"
)

func TestNewBoltDB(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("User Home Dir is: %v \nError:, %v", homeDir, err)
	}
	dbFile = homeDir + dbFile
	boltdb, err := repository.NewBolt(dbFile)
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, boltdb, "boltdb should not nil")
	assert.FileExists(t, dbFile, "Slip database file should be exists")
	// Teardown
	err = os.Remove(dbFile)
	if err != nil {
		t.Error(err)
	}
}

// TODO:
// create slipdata.json for mock data
