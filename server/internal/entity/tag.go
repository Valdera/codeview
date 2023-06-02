package entity

import (
	"time"

	"github.com/google/uuid"
)

type TagType string

const (
	TagTypeNote    TagType = "NOTE"
	TagTypeProblem TagType = "PROBLEM"
)

type Tag struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Label     string
	Color     string
	TagType   TagType
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
