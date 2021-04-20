package printer_test

import (
	"testing"

	"github.com/goccy/go-yaml/printer"
	"github.com/stretchr/testify/assert"
)

func TestNewPrinter(t *testing.T) {
	p := printer.NewPrinter()
	assert.NotNil(t, p, "NewPrinter should not nil: %v", p)
}

// Test HTMLtoSlipJPG must install wkhtml to host server by
func TestHTMLToSlipJPG(t *testing.T) {
	//setup HTML
	html := "!DOCTYPE=html"

	printer := print.NewPrinter()
	b, err := printer.HtmlToSlipJPG(html, 561, "jpg")
	if err != nil {
		t.Errorf("printer.HtmlToSlipJPG() error: %v", err)
	}
	assert.NotNil(t, b, "HTMLToSlipJPG() should not return nil")
}
