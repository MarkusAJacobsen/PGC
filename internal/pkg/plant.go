package pkg

type Plant struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	LatinName string `json:"latinName"`
	Family    string `json:"family,omitempty"`
}
