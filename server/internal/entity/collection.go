package entity

import (
	"time"

	"github.com/google/uuid"
)

type CollectionType string

const (
	CollectionTypeProblem CollectionType = "DRAFT"
	CollectionTypeNote    CollectionType = "PUBLISHED"
)

type Collection struct {
	ID             uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Title          string
	Description    string
	Emoji          string
	Tags           []Tag            `gorm:"many2many:collections_tags"`
	CollectionItem []CollectionItem `gorm:"foreignKey:CollectionID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

type CollectionItem struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CollectionID uuid.UUID
	ItemID       uuid.UUID
	Position     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
