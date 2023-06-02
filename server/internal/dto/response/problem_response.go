package response

import (
	"time"

	"github.com/google/uuid"
)

type Problem struct {
	ID         uuid.UUID  `json:"id"`
	Title      string     `json:"title"`
	Rating     int        `json:"rating"`
	Difficulty Difficulty `json:"difficulty"`
	Tags       []Tag      `json:"tags"`
	Sources    []Source   `json:"sources"`
	Questions  []Question `json:"questions"`
	Solutions  []Solution `json:"solutions"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
