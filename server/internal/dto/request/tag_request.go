package request

type TagCreate struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

type TagUpdate struct {
	Label string `json:"label"`
	Color string `json:"color"`
}
