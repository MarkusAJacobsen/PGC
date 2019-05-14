package pkg

type Guide struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Stages []Stage `json:"stages"`
}

type Stage struct {
	Id     string   `json:"id"`
	PageNr int      `json:"pageNr"`
	Text   string   `json:"text"`
	Images []string `json:"images"`
}

type GuideComposite struct {
	Guide Guide `json:"guide"`
	Stages []Stage `json:"stages"`
}
