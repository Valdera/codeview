package request

type SourceCreate struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

type SourceUpdate struct {
	Label string `json:"label"`
	Color string `json:"color"`
}
