package pkg

type Project struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	StartDate string `json:"startDate"`
	Climate   string `json:"climate"`
}

type ProjectLink struct {
	Project Project `json:"project"`
	UserId  string  `json:"idToken"`
	PlantId string  `json:"pId"`
}
