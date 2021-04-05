package slip

import (
	"net/url"
)

type Service interface {
	NewSlip(*Slip) url.URL
	FindSlipByID(id string) url.URL
}
