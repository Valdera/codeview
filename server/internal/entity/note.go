package entity

import (
	"time"

	"github.com/google/uuid"
)

type NoteStatus string

const (
	NoteStatusDraft     NoteStatus = "DRAFT"
	NoteStatusPublished NoteStatus = "PUBLISHED"
)

type Note struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Title     string
	Emoji     string
	Tags      []Tag      `gorm:"many2many:notes_tags"`
	NoteItem  []NoteItem `gorm:"foreignKey:NoteID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type NoteItem struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	NoteID    uuid.UUID
	Position  int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
