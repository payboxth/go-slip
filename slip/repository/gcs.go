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

func (s *gcs) StoreFile(ctx context.Context, fileName, object string) (url string, err error) {

	o := s.bucket.Object(object)
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

	return s.baseURL + "/" + object, nil
}

func (s *gcs) StoreOriginPNG(ctx context.Context, m image.Image, path string) (name, url string, err error) {
	fileName := s.generateName() + ".png"
	filePath := fmt.Sprintf("%s/%s", path, fileName)
	obj := s.bucket.Object(filePath)
	w := obj.NewWriter(ctx)

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
		objErr := obj.Delete(ctx)
		if objErr != nil {
			return "", "", err
		}
		return "", "", err
	}

	return fileName, s.baseURL + "/" + filePath, nil
}

func (s *gcs) RemoveFile(ctx context.Context, object string) error {
	// TODO
	return nil
}

// deleteFile removes specified object.
func deleteFile(w io.Writer, bucket, object string) error {
	// bucket := "bucket-name"
	// object := "object-name"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	o := client.Bucket(bucket).Object(object)
	if err := o.Delete(ctx); err != nil {
		return fmt.Errorf("Object(%q).Delete: %v", object, err)
	}
	fmt.Fprintf(w, "Blob %v deleted.\n", object)
	return nil
}

// copyFile copies an object into specified bucket.
func copyFile(w io.Writer, dstBucket, srcBucket, srcObject string) error {
	// dstBucket := "bucket-1"
	// srcBucket := "bucket-2"
	// srcObject := "object"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	dstObject := srcObject + "-copy"
	src := client.Bucket(srcBucket).Object(srcObject)
	dst := client.Bucket(dstBucket).Object(dstObject)

	if _, err := dst.CopierFrom(src).Run(ctx); err != nil {
		return fmt.Errorf("Object(%q).CopierFrom(%q).Run: %v", dstObject, srcObject, err)
	}
	fmt.Fprintf(w, "Blob %v in bucket %v copied to blob %v in bucket %v.\n", srcObject, srcBucket, dstObject, dstBucket)
	return nil
}
