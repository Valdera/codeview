package response

import "time"

type Difficulty struct {
	ID        uint      `json:"id"`
	Label     string    `json:"label"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
