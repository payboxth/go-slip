package slip

type Service interface {
	NewSlip(*Head) (id string, url string, error)
	FindByID(id string) (*Slip, error)
}
