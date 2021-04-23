package slip

type Render interface {
	// HtmlToSlipJPG function create jpeg file image from -> input HTML string
	// -> slip width in pixel
	// -> input file format JPG or PNG
	Bytes(html string, width int, format string) ([]byte, error)
	Bytes2(body *Body, html string, width int, format string) ([]byte, error)
}
