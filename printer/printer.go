package printer

import (
	"bytes"
	"html/template"
	"os"

	"github.com/payboxth/go-slip"
)

type printer struct{}

func NewPrinter() slip.Printer {
	p := &printer{}
	return p
}

// H
func (p *printer) HtmlToSlipJPG(html string, width int, format string) ([]byte, error) {
	mailTmpl, err := template.New("test").Parse(html)
	if err != nil {

		return nil, err
	}
	buf := new(bytes.Buffer)
	err = mailTmpl.Execute(buf, nil)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat("./slip.jpg"); os.IsNotExist(err) {
		f, err := os.Create("./slip.jpg")
		if err != nil {
			panic(err)
		}
		defer f.Close()
	}

	d := ImageOptions{BinaryPath: "wkhtmltoimage", Input: "-", Width: width, Format: format, Output: "./slip.jpg", HTML: buf.String()}
	out, err := GenerateImage(&d)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// func (s *printer) UploadToGCLOUD(_byte []byte) (url *string, err error) {
// 	ctx := context.Background()
// 	bh := s.client.Bucket(s.bucket)
// 	// Next check if the bucket exists
// 	if _, err := bh.Attrs(ctx); err != nil {
// 		return nil, err
// 	}

// 	newUUID, err := uuid.NewV4()
// 	if err != nil {
// 		return nil, err
// 	}
// 	filename := strings.Replace(newUUID.String(), "-", "", -1) + ".jpg"
// 	obj := bh.Object(filename)
// 	w := obj.NewWriter(ctx)
// 	buf := bytes.NewBuffer(nil)

// 	buf.Write(_byte)
// 	w.ACL = append(w.ACL, storage.ACLRule{Entity: storage.AllUsers, Role: storage.RoleReader})
// 	w.CacheControl = "public, max-age=31536000"
// 	if _, err := io.Copy(w, buf); err != nil {
// 		return nil, err
// 	}
// 	filename = s.baseURL + "/" + filename
// 	return &filename, nil
// }
