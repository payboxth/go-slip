package render_test

import (
	"testing"

	"github.com/payboxth/go-slip/render"
	"github.com/stretchr/testify/assert"
)

func TestNewRender(t *testing.T) {
	r := render.New()
	assert.NotNil(t, r, "NewPrinter should not nil: %v", r)
}

// Test HTMLtoSlipJPG must install wkhtml to host server by
func TestHTMLToSlipJPG(t *testing.T) {
	//setup HTML
	html := "!DOCTYPE=html"

	r := render.New()
	b, err := r.HtmlToSlipJPG(html, 561, "jpg")
	if err != nil {
		t.Errorf("printer.HtmlToSlipJPG() error: %v", err)
	}
	assert.NotNil(t, b, "HTMLToSlipJPG() should not return nil")
}
