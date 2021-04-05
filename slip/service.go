package slip

import (
	"net/url"
	"time"
)

type Service interface {
	NewSlip(*Slip) url.URL
	FindSlipByID(id string) url.URL
}

type Slip struct {
	ID          string
	DocNumber   string
	RefNumber   string
	Title       string
	DocDate     string
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
