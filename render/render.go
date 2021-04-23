package render

import (
	"bytes"
	"html/template"
	"os"

	"github.com/payboxth/goslip"
)

type render struct{}

func New() slip.Render {
	p := &render{}
	return p
}

// Jpg is a slip rendering to binary JPEG file function.
func (r *render) Bytes(html string, width int, format string) ([]byte, error) {
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

func (r *render) Bytes2(body *slip.Body, html string, width int, format string) ([]byte, error) {
	tmpl, err := template.New("slip").Parse(html)
	if err != nil {
		return nil, err
	}
	slip := &slip.Body{}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, slip)
	if err != nil {
		return nil, err
	}
	var b []byte
	// TODO implement how to return image byte.

	return b, nil
}
