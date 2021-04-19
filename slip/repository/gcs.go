package repository

import (
	"context"
	"fmt"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"time"

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

func encodePNG(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}

func (s *gcs) StoreFile(ctx context.Context, fileName, objectName string) (url string, err error) {
	o := s.bucket.Object(objectName)
	w := o.NewWriter(ctx)
	defer func() {
		err := w.Close()
		if err != nil {
			fmt.Printf("Cannot Close *storage.Writer: %v\n", err) // TODO หาวิธีส่ง Error ไปเก็บ
		}
	}()

	w.ACL = append(w.ACL, storage.ACLRule{Entity: storage.AllUsers, Role: storage.RoleReader})
	w.CacheControl = "public, max-age=31536000"

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	_, err = w.Write(b)
	if err != nil {
		objErr := o.Delete(ctx)
		if objErr != nil {
			return "", objErr
		}
		return "", err
	}

	return s.baseURL + "/" + objectName, nil
}

func (s *gcs) StoreByte(ctx context.Context, b []byte, objectName string) (url string, err error) {
	o := s.bucket.Object(objectName)
	w := o.NewWriter(ctx)
	defer w.Close()
	w.ACL = append(w.ACL, storage.ACLRule{Entity: storage.AllUsers, Role: storage.RoleReader})
	w.CacheControl = "public, max-age=31536000"
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	n, err := w.Write(b)
	if err != nil {
		objErr := o.Delete(ctx)
		if objErr != nil {
			fmt.Printf("Delete object error: %v", objErr)
		}
		return "", err
	}
	fmt.Printf("StoreByte write %v byte.", n)
	return s.baseURL + "/" + objectName, nil
}

func (s *gcs) StoreOriginPNG(ctx context.Context, m image.Image, objectName string) (url string, err error) {
	o := s.bucket.Object(objectName)
	w := o.NewWriter(ctx)

	defer func() {
		err := w.Close()
		if err != nil {
			fmt.Printf("Cannot Close *storage.Writer: %v\n", err) // TODO หาวิธีส่ง Error ไปเก็บ ใน Sentry
		}
	}()

	w.ACL = append(w.ACL, storage.ACLRule{Entity: storage.AllUsers, Role: storage.RoleReader})
	w.CacheControl = "public, max-age=31536000"
	err = encodePNG(w, m)
	if err != nil {
		objErr := o.Delete(ctx)
		if objErr != nil {
			return "", err
		}
		return "", err
	}

	return s.baseURL + "/" + objectName, nil
}

func (s *gcs) RemoveFile(ctx context.Context, objectName string) error {
	o := s.bucket.Object(objectName)
	w := o.NewWriter(ctx)
	defer w.Close()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	if err := o.Delete(ctx); err != nil {
		return fmt.Errorf("Object(%q).Delete: %v", objectName, err)
	}
	fmt.Fprintf(w, "Blob %v deleted.\n", objectName)
	return nil
}
