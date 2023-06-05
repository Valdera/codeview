package response

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID        uuid.UUID  `json:"id"`
	Title     string     `json:"title"`
	Tags      []Tag      `json:"tags"`
	Emoji     string     `json:"emoji"`
	Items     []NoteItem `json:"items"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type NoteItem struct {
	ID        uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	Position  int       `json:"position"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
