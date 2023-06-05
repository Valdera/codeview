package request

type TagCreate struct {
	Label   string `json:"label"`
	Color   string `json:"color"`
	TagType string `json:"tag_type"`
}

type TagUpdate struct {
	Label string `json:"label"`
	Color string `json:"color"`
}
