package request

type ProblemCreate struct {
	Title        string `json:"title"`
	Rating       int    `json:"rating"`
	DifficultyID int    `json:"difficulty_id"`
}

type ProblemUpdate struct {
	Title      string `json:"title"`
	Rating     int    `json:"rating"`
	Difficulty int    `json:"difficulty"`
}
