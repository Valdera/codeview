package dto

type ProblemRequest struct {
	ID uint `json:"id"`
}

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

type TagRequest struct {
	ID uint `json:"id"`
}

type TagCreate struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

type TagUpdate struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

type SourceRequest struct {
	ID uint `json:"id"`
}

type SourceCreate struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

type SourceUpdate struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

type DifficultyRequest struct {
	ID uint `json:"id"`
}

type DifficultyCreate struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

type DifficultyUpdate struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

type SolutionRequest struct {
	ID uint `json:"id"`
}

type SolutionCreate struct {
	Content string `json:"content"`
}

type SolutionUpdate struct {
	Content string `json:"content"`
}
type QuestionRequest struct {
	ID uint `json:"id"`
}

type QuestionCreate struct {
	Content string `json:"content"`
}

type QuestionUpdate struct {
	Content string `json:"content"`
}
