package pkg

type Plant struct {
	Name      string `json:"name"`
	LatinName string `json:"latinName"`
	Family    string `json:"family,omitempty"`
}
