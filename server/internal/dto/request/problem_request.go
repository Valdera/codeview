package request

type ProblemCreate struct {
	Title        string `json:"title"`
	Rating       int    `json:"rating"`
	DifficultyID string `json:"difficulty_id"`
	Emoji        string `json:"emoji"`
}

type ProblemUpdate struct {
	Title        string `json:"title"`
	Rating       int    `json:"rating"`
	DifficultyID string `json:"difficulty_id"`
	Emoji        string `json:"emoji"`
}
