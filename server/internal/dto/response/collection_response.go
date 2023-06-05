package response

import (
	"time"

	"github.com/google/uuid"
)

type Collection struct {
	ID          uuid.UUID        `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Tags        []Tag            `json:"tags"`
	Emoji       string           `json:"emoji"`
	Items       []CollectionItem `json:"items"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

type CollectionItem struct {
	ID        uuid.UUID `json:"id"`
	ItemID    uuid.UUID `json:"item_id"`
	ItemTitle string    `json:"item_title"`
	ItemTags  []Tag     `json:"item_tags"`
	ItemEmoji string    `json:"item_emoji"`
	Position  int       `json:"position"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
