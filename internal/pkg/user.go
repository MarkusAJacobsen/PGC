package pkg

type User struct {
	IdToken string `json:"idToken"`
	Name    string `json:"name"`
	Origin  string `json:"origin,omitempty"`
	Email   string `json:"email,omitempty"`
	Area    string `json:"area,omitempty"`
}
