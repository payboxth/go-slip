package slip

import "time"

type Head struct {
	ID          string
	DocNumber   string
	RefNumber   string
	Title       string
	DocDate     string
	URL         string
	Lines       []Line
	AccessToken string
	CreateDate  time.Time
}

type Line struct {
	LineNumber  int8
	SKU         string
	Description string
	Quantity    float32
	Price       float32
	Note        string
}
