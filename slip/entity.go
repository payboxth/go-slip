package slip

import "time"

type Slip struct {
	ID          string
	DocNumber   string
	RefNumber   string
	Title       string
	DocDate     string
	URL         string
	SlipLines   []SlipLine
	AccessToken string
	CreateDate  time.Time
}

type SlipLine struct {
	LineNumber  int8
	SKU         string
	Description string
	Quantity    float32
	Price       float32
	Note        string
}
