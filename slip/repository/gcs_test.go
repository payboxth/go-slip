package repository_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	sliprepository "github.com/payboxth/go-slip/slip/repository"
)

var (
	bucketName string = "paybox_slip"
	secretPath string = "/Users/tom/secret/paybox_slip.json"
	fileName   string = "test_slip.png"
	folderName string = "test"
)

func TestNewGCSClient(t *testing.T) {
	_, err := sliprepository.NewGCS("paybox_slip", "/Users/tom/secret/paybox_slip.json")
	if err != nil {
		t.Errorf("Repository cannot create Storage Client: %v", err)
	}
	t.Log("Create new GCS Client: OK")
}

// ตอนรันเทส ต้องเอา paybox_slip.json ไปใส่ด้วยนะ อยู่ใน slip/repository/secret
// ตอนนี้ยังเอาไว้ทำ integration test
// TODO แต่ถ้าปล่อย package นี้เป็น lib opensource จริงคงต้องแยกทำ mock ไว้ เทสกันด้วย
func TestStoreFile_URLMustContainPath(t *testing.T) {
	generateName := uuid.New().String()
	objectName := fmt.Sprintf("%s/%s", folderName, generateName)
	t.Logf("object = %v", objectName)

	s, err := sliprepository.NewGCS(bucketName, secretPath)
	if err != nil {
		t.Errorf("Repository cannot create Storage Client: %v", err)
	}

	ctx := context.Background()

	url, err := s.StoreFile(ctx, fileName, objectName)
	if err != nil {
		t.Fatalf("Error on s.SaveFile: %v", err)
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
}
