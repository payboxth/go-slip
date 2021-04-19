package slip

type Body struct {
	ID          string
	DocDate     string
	DocNumber   string
	RefNumber   string
	Title       string
	Total       float32
	ImageURL    string
	AccessToken string
	CreateDate  string
	Lines       []Line
}

type Line struct {
	LineNumber  int8
	SKU         string
	Description string
	Quantity    float32
	Price       float32
	Note        string
}
