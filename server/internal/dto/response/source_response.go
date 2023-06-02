package response

import (
	"time"

	"github.com/google/uuid"
)

type Source struct {
	ID        uuid.UUID `json:"id"`
	Label     string    `json:"label"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
