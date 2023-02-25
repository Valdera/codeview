package request

type DifficultyCreate struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

type DifficultyUpdate struct {
	Label string `json:"label"`
	Color string `json:"color"`
}
