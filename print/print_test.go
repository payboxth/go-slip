package print_test

import (
	"testing"

	"github.com/payboxth/go-slip/print"
	"github.com/stretchr/testify/assert"
)

func TestNewPrinter(t *testing.T) {
	printer := print.NewPrinter()
	assert.NotNil(t, printer, "NewPrinter should not nil: %v", printer)
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
