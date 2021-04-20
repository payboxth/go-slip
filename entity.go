package slip

type Body struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	DocDate     string  `json:"doc_date"`
	DocNumber   string  `json:"doc_number"`
	Ref         string  `json:"ref"`
	Total       float32 `json:"total"`
	ImageURL    string  `json:"image_url"`
	AccessToken string  `json:"access_token"`
	CreateDate  string  `json:"create_date"`
	Lines       []Line  `json:"lines"`
}

type Line struct {
	Seq   int8    `json:"seq"`
	SKU   string  `json:"sku"`
	Name  string  `json:"name"`
	Qty   float32 `json:"qty"`
	Price float32 `json:"price"`
	Note  string  `json:"note"`
}
