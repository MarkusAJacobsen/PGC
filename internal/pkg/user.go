package pkg

type User struct {
	IdToken string `json:"idToken"`
	Name    string `json:"name"`
	Email   string `json:"email,omitempty"`
	Area    string `json:"area,omitempty"`
}
