package dto

import "time"

type ProblemResponse struct {
	ID         uint               `json:"id"`
	Title      string             `json:"title"`
	Rating     int                `json:"rating"`
	Difficulty DifficultyResponse `json:"difficulty"`
	Tags       []TagResponse      `json:"tags"`
	Sources    []SourceResponse   `json:"sources"`
	Questions  []QuestionResponse `json:"questions"`
	Solutions  []SolutionResponse `json:"solutions"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

type TagResponse struct {
	ID        uint      `json:"id"`
	Label     string    `json:"label"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SourceResponse struct {
	ID        uint      `json:"id"`
	Label     string    `json:"label"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DifficultyResponse struct {
	ID        uint      `json:"id"`
	Label     string    `json:"label"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SolutionResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type QuestionResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
