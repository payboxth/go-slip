package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/storage"
	"github.com/payboxth/go-slip/slip"
)

func NewGCS() slip.Storage {
	ctx := context.Background()

	projectID := "mrtomyum"

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create gcs client: %v", err)
	}

	bucketName := "my-new-bucket"

	bucket := client.Bucket(bucketName)

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	if err := bucket.Create(ctx, projectID, nil); err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
	}
	fmt.Printf("Bucket %v created. \n", bucketName)

	return &gcs{client}
}

type gcs struct {
	Client *storage.Client
}

func (gcs) SaveFile(ctx context.Context, file []byte) (path string, err error) {

	return "", nil
}
