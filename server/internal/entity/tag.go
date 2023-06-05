package entity

import (
	"time"

	"github.com/google/uuid"
)

type TagType string

const (
	TagTypeNote       TagType = "NOTE"
	TagTypeProblem    TagType = "PROBLEM"
	TagTypeCollection TagType = "COLLECTION"
)

type Tag struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Label     string
	Color     string
	TagType   TagType
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
