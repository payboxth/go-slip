package repository_test

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/payboxth/goslip"
	"github.com/payboxth/goslip/repository"
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

	bolt, err := repository.NewBolt(dbFile)
	if err != nil {
		t.Error(err)
	}

	body := &slip.Body{
		Title:      "MakeKAFE",
		DocDate:    "2021-04-02T15:04:05+07:00", // time formated in RFC3339
		DocNumber:  "000001",
		Ref:        "202341",
		CreateDate: time.Now().Format(time.RFC3339),
		Lines: []slip.Line{
			{
				Seq:   1,
				SKU:   "123456",
				Name:  "Product Name1",
				Qty:   1.0,
				Price: 100,
				Note:  "",
			},
			{
				Seq:   2,
				SKU:   "78910",
				Name:  "Product Name2",
				Qty:   2.0,
				Price: 50,
				Note:  "Test Note",
			},
		},
	}
	ctx := context.Background()

	id, err := bolt.Insert(ctx, body)
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
func TestDB_FindByID(t *testing.T) {
	// Setup
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("User Home Dir is: %v \nError:, %v", homeDir, err)
	}
	dbFile = homeDir + dbFile

	bolt, err := repository.NewBolt(dbFile)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Open boltDB as: %v", bolt)

	lines := []slip.Line{
		{
			Seq:   1,
			SKU:   "123456",
			Name:  "Product Name1",
			Qty:   1.0,
			Price: 100,
			Note:  "",
		},
		{
			Seq:   2,
			SKU:   "78910",
			Name:  "Product Name2",
			Qty:   2.0,
			Price: 50,
			Note:  "Test Note",
		},
	}
	bodyIn := &slip.Body{
		DocNumber:  "101010",
		DocDate:    "93993",
		Ref:        "00001",
		Title:      "MakeKAFE",
		CreateDate: time.Now().Format(time.RFC3339),
		Lines:      lines,
	}
	ctx := context.Background()

	id, err := bolt.Insert(ctx, bodyIn)
	if err != nil {
		t.Errorf("boltdb.Insert error: %v", err)
	}
	assert.NotEmptyf(t, id, "boltdb.Insert() return id should not empty as: %v", id)
	t.Logf("id = %v", id)

	bodyOut, err := bolt.FindByID(ctx, id)
	if err != nil {
		t.Errorf("boltdb.FindByID error: %v", err)
	}
	if !reflect.DeepEqual(bodyIn, bodyOut) {
		t.Errorf("boltdb.FindByID\n   bodyIn : %v\n !=bodyOut: %v", bodyIn, bodyOut)
	}

	// Teardown
	err = os.Remove(dbFile)
	if err != nil {
		t.Errorf("Teardown error when remove dbFile: %v", err)
	}
}
