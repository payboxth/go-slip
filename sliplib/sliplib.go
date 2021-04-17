package sliplib

import (
	"bytes"
	"context"
	"html/template"
	"io"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/gofrs/uuid"
	"google.golang.org/api/option"
)

type SlipLib struct {
	// request
	client  *storage.Client
	bucket  string
	baseURL string //gcloud url
}

func New(_filename string, busket string, baseURL string) (e *SlipLib) {

	storageClient, err := storage.NewClient(context.Background(), option.WithCredentialsFile(_filename))
	if err != nil {
		panic(err)
	}
	e = &SlipLib{client: storageClient, bucket: busket, baseURL: baseURL}
	return
}

func (s *SlipLib) HtmlToSlipJPG(_html string, width int, format string) ([]byte, error) {
	mailTmpl, err := template.New("test").Parse(_html)
	if err != nil {

		return nil, err
	}
	buf := new(bytes.Buffer)
	err = mailTmpl.Execute(buf, nil)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat("./example.jpg"); os.IsNotExist(err) {
		f, err := os.Create("./example.jpg")
		if err != nil {
			panic(err)
		}
		defer f.Close()
	}

	d := ImageOptions{BinaryPath: "wkhtmltoimage", Input: "-", Width: width, Format: format, Output: "./example.jpg", HTML: buf.String()}
	out, err := GenerateImage(&d)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SlipLib) UploadToGCLOUD(_byte []byte) (url *string, err error) {
	ctx := context.Background()
	bh := s.client.Bucket(s.bucket)
	// Next check if the bucket exists
	if _, err := bh.Attrs(ctx); err != nil {
		return nil, err
	}

	newUUID, err := uuid.NewV4()

	if err != nil {
		return nil, err
	}
	randomKey := strings.Replace(newUUID.String(), "-", "", -1)
	filename := randomKey + ".jpg"
	obj := bh.Object(filename)
	w := obj.NewWriter(ctx)
	buf := bytes.NewBuffer(nil)
	// var b bytes.Buffer

	// Write strings to the Buffer.
	// b.WriteString(txt)
	buf.Write(_byte)
	w.ACL = append(w.ACL, storage.ACLRule{Entity: storage.AllUsers, Role: storage.RoleReader})
	w.CacheControl = "public, max-age=31536000"
	// pdfg.SetOutput(w)
	// pdfg.SetStderr(w)
	if _, err := io.Copy(w, buf); err != nil {
		return nil, err
	}
	filename = s.baseURL + "/" + filename
	return &filename, nil
}
