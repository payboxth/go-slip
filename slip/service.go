package slip

import (
	"net/url"
)

type Service interface {
	NewSlip(*Slip) (id string, url url.URL)
	FindSlipByID(id string) *Slip
}
