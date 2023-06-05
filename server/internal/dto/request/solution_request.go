package request

type SolutionCreate struct {
	ProblemID string `json:"problem_id"`
	Content   string `json:"content"`
}

type SolutionUpdate struct {
	Content string `json:"content"`
}
