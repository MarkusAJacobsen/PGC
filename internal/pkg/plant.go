package pkg

type Plant struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name"`
	LatinName string `json:"latinName"`
	Family    string `json:"family,omitempty"`
	Barcode   string `json:"barcode,omitempty"`
	Category  uint8  `json:"category,omitempty"`
}
