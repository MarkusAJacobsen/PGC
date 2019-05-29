package pkg

type Project struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name"`
	StartDate string `json:"startDate,omitempty"`
	Climate   string `json:"climate"`
	Image     string `json:"image"`
	Status    int64  `json:"status"`
}

type ProjectLink struct {
	Project Project `json:"project"`
	UserId  string  `json:"idToken"`
	PlantId string  `json:"pId"`
}
