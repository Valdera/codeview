package request

type QuestionCreate struct {
	Content string `json:"content"`
}

type QuestionUpdate struct {
	Content string `json:"content"`
}
