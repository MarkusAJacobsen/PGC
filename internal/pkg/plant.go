package pkg

type Plant struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name"`
	LatinName string `json:"latinName"`
	Image     string `json:"image"`
	Family    string `json:"family,omitempty"`
	Barcode   string `json:"barcode,omitempty"`
	Category  string `json:"category,omitempty"`
	GuideID   string `json:"guideID,omitempty"`
}
