package entity

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Title     string
	Tags      []Tag      `gorm:"many2many:notes_tags"`
	NoteItem  []NoteItem `gorm:"foreignKey:NoteID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type NoteItem struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	NoteID    uuid.UUID
	Header    string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
