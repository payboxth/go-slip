package slip

type Printer interface {
	// HtmlToSlipJPG function create jpeg file image from -> input HTML string
	// -> slip width in pixel
	// -> input file format JPG or PNG
	HtmlToSlipJPG(html string, width int, format string) ([]byte, error)
}
