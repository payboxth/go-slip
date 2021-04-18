package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	bolt "go.etcd.io/bbolt"

	"github.com/payboxth/go-slip/slip"
)

func NewBolt(fileName string) (slip.Database, error) {
	database := new(boltDB)
	// if _, err := os.Stat(path); os.IsNotExist(err) {
	// slip.db does not exist
	config := &bolt.Options{Timeout: 1 * time.Second}
	b, err := bolt.Open(fileName, 0600, config)
	if err != nil {
		return nil, err
	}
	database.db = b
	return database, nil
}

type boltDB struct {
	db *bolt.DB
}

// Insert is method for create new record in slip.db in bucket "slips"
func (b boltDB) Insert(ctx context.Context, body *slip.Body) (id string, err error) {
	err = b.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("slips"))
		if err != nil {
			return err
		}
		id, _ := bucket.NextSequence()
		body.ID = fmt.Sprint(id)

		encoded, err := json.Marshal(body)
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(body.ID), encoded)
		return err
	})
	if err != nil {
		return "", err
	}
	return body.ID, nil
}

// FindByID is method to get slip body by slip ID
func (b boltDB) FindByID(ctx context.Context, id string) (*slip.Body, error) {
	// TODO: not sure to use bolt.Open here?
	// db, err := bolt.Open("slip.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	// if err != nil {
	// 	return nil, err
	// }
	// defer db.Close()
	body := slip.Body{}
	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("slips"))
		v := bucket.Get([]byte(id))
		err := json.Unmarshal(v, &body)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &body, nil
}
