package repository

import (
	"context"
	"fmt"
	"io/ioutil"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"google.golang.org/api/option"

	"github.com/payboxth/go-slip/slip"
)

func NewGCS(bucketName, credentialFile string) (slip.Storage, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFile))
	if err != nil {
		return nil, err
	}
	bucket := client.Bucket(bucketName)
	s := &gcs{
		bucket:  bucket,
		baseURL: "https://storage.googleapis.com/" + bucketName,
	}

	return s, nil
}

type gcs struct {
	bucket  *storage.BucketHandle
	baseURL string
}

func (s *gcs) generateName() string {
	return uuid.New().String()
}

func (s *gcs) SaveFile(ctx context.Context, file, path string) (url string, err error) {
	fileName := s.generateName()
	filePath := fmt.Sprintf("%s/%s", path, fileName)
	obj := s.bucket.Object(filePath)
	w := obj.NewWriter(ctx)
	defer func() {
		err := w.Close()
		if err != nil {
			fmt.Printf("Cannot Close *storage.Writer: %v\n", err) // TODO หาวิธีส่ง Error ไปเก็บ
		}
	}()

	w.ACL = append(w.ACL, storage.ACLRule{Entity: storage.AllUsers, Role: storage.RoleReader})
	w.CacheControl = "public, max-age=31536000"

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	_, err = w.Write(b)
	if err != nil {
		objErr := obj.Delete(ctx)
		if objErr != nil {
			return "", objErr
		}
		return "", err
	}

	return s.baseURL + "/" + filePath, nil
}
