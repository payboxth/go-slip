package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	bolt "go.etcd.io/bbolt"

	"github.com/payboxth/go-slip/slip"
)

func NewBolt() slip.Database {
	return &boltDB{}
}

type boltDB struct{}

func (boltDB) Insert(ctx context.Context, body *slip.Body) (id string, err error) {
	db, err := bolt.Open("slip.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return "", err
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("slips"))
		if err != nil {
			return err
		}
		id, _ := b.NextSequence()
		body.ID = fmt.Sprint(id)

		encoded, err := json.Marshal(body)
		if err != nil {
			return err
		}
		err = b.Put([]byte(body.ID), encoded)
		return err
	})
	return body.ID, nil
}

func (boltDB) FindByID(ctx context.Context, id string) (*slip.Body, error) {
	// Todo: not sure to use bolt.Open here?
	db, err := bolt.Open("slip.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	defer db.Close()
	body := &slip.Body{}
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("slips"))
		if err != nil {
			return err
		}
		v := b.Get([]byte(id))
		err := json.Unmarshal(v, body)
		if err != nil {
			return err
		}
		return nil
	})
	return body, nil
}
