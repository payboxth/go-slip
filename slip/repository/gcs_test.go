package repository_test

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/png"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/payboxth/go-slip/slip/repository"
)

var (
	bucketName     string = "paybox_slip"
	credentialFile string = "/secret/paybox_slip_key.json"
	fileName       string = "test_slip.png"
	folderName     string = "test"
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("User Home Dir is: %v \nError:, %v", homeDir, err)
	}
	credentialFile = homeDir + credentialFile
}

func TestNewGCSClient(t *testing.T) {
	_, err := repository.NewGCS(bucketName, credentialFile)
	if err != nil {
		t.Errorf("Repository cannot create Storage Client: %v", err)
	}
	t.Log("Create new GCS Client: OK")
}

// ตอนรันเทส ต้องเอา paybox_slip.json ไปใส่ด้วยนะ อยู่ใน slip/repository/secret
// ตอนนี้ยังเอาไว้ทำ integration test
// TODO แต่ถ้าปล่อย package นี้เป็น lib opensource จริงคงต้องแยกทำ mock ไว้ เทสกันด้วย
func TestStoreFile(t *testing.T) {
	generateName := uuid.New().String()
	objectName := fmt.Sprintf("%s/%s", folderName, generateName)
	t.Logf("object = %v", objectName)

	s, err := repository.NewGCS(bucketName, credentialFile)
	if err != nil {
		t.Errorf("Repository cannot create Storage Client: %v", err)
	}

	ctx := context.Background()

	url, err := s.StoreFile(ctx, fileName, objectName)
	if err != nil {
		t.Fatalf("Error on s.StoreFile: %v", err)
	}

	expected := "https://storage.googleapis.com/paybox_slip/test/"
	assert.Containsf(t, url, expected, "Return URL does not contain ecpected = %v", url)
	assert.NotZerof(t, url, "URL is not empty as: %v", url)
	t.Logf("Success storage save file and return fileName = %v URL = %v", fileName, url)

	// Teardown by delete saved file
	err = s.RemoveFile(ctx, objectName)
	if err != nil {
		t.Errorf("cannot teardown by delete saved file: %v", err)
	}
	t.Logf("Teardown by delete file: %v", objectName)
}

func TestStoreByte(t *testing.T) {
	s, err := repository.NewGCS(bucketName, credentialFile)
	if err != nil {
		t.Errorf("Repository cannot create Storage Client: %v", err)
	}

	ctx := context.Background()
	inputFile, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	m, format, err := image.Decode(inputFile)
	if err != nil {
		t.Errorf("image.Decode error = %v", err)
	}
	if format != "png" {
		t.Errorf("file format is not png but: %v", format)
	}
	t.Logf("file format is: %v", format)
	buf := new(bytes.Buffer)
	err = png.Encode(buf, m)
	if err != nil {
		t.Errorf("Error while png.Encode() = %v", err)
	}
	b := buf.Bytes()
	generateName := uuid.New().String()
	objectName := fmt.Sprintf("%s/%s.%s", folderName, generateName, "png")
	t.Logf("objectName = %v", objectName)
	url, err := s.StoreByte(ctx, b, objectName)
	if err != nil {
		assert.NotNilf(t, err, "Error on s.StoreByte(): %v", err)
	}
	t.Logf("Success storage save byte and return URL = %v", url)

	// Teardown by delete saved file
	err = s.RemoveFile(ctx, objectName)
	if err != nil {
		t.Errorf("cannot teardown by delete saved file: %v", err)
	}
	t.Logf("Teardown by delete file: %v", objectName)
}
