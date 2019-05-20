package pkg

type Guide struct {
	Id            string  `json:"id"`
	Title         string  `json:"title,omitempty"`
	ChapterTitles []string  `json:"chapterTitles"`
	Stages        []Stage `json:"stages"`
}

type Stage struct {
	Id        string   `json:"id"`
	Title     string   `json:"title"`
	PageNr    int64    `json:"pageNr"`
	ChapterNr int64    `json:"chapterNr"`
	Filter    string   `json:"filter"`
	Text      string   `json:"text"`
	Images    []string `json:"images"`
}

type GuideComposite struct {
	Guide  Guide   `json:"guide"`
	Stages []Stage `json:"stages"`
}
