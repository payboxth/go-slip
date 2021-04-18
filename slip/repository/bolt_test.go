package repository_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/payboxth/go-slip/slip"
	"github.com/payboxth/go-slip/slip/repository"
	"github.com/stretchr/testify/assert"
)

var (
	// dbFile is a path for database file.
	// If you need to keep file in a folder you must create one before you call it in NewBolt(dbFile).
	dbFile string = "/slip.db"
)

func TestNewBoltDB(t *testing.T) {
	// Setup
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
		t.Errorf("Teardown error when remove dbFile: %v", err)
	}
}

// TODO:
// create slipdata.json for mock data

// TestDB_Insert()
func TestDB_Insert(t *testing.T) {
	// Setup
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("User Home Dir is: %v \nError:, %v", homeDir, err)
	}
	dbFile = homeDir + dbFile

	boltdb, err := repository.NewBolt(dbFile)
	if err != nil {
		t.Error(err)
	}
	lines := []slip.Line{
		{
			LineNumber:  1,
			SKU:         "123456",
			Description: "Product Name1",
			Quantity:    1.0,
			Price:       100,
			Note:        "",
		},
		{
			LineNumber:  2,
			SKU:         "78910",
			Description: "Product Name2",
			Quantity:    2.0,
			Price:       50,
			Note:        "Test Note",
		},
	}
	body := &slip.Body{
		DocNumber:  "101010",
		DocDate:    "93993",
		RefNumber:  "00001",
		Title:      "MakeKAFE",
		CreateDate: time.Now(),
		Lines:      lines,
	}
	ctx := context.Background()

	id, err := boltdb.Insert(ctx, body)
	if err != nil {
		t.Errorf("boltdb.Insert error: %v", err)
	}
	assert.NotEmptyf(t, id, "boltdb.Insert() return id should not empty as: %v", id)

	// Teardown
	err = os.Remove(dbFile)
	if err != nil {
		t.Errorf("Teardown error when remove dbFile: %v", err)
	}
}

// TestDB_FindByID()
