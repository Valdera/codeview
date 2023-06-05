package entity

import (
	"time"

	"github.com/google/uuid"
)

type Problem struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Title        string
	Rating       *int
	Emoji        string
	DifficultyID uuid.UUID
	Difficulty   Difficulty `gorm:"foreignKey:DifficultyID"`
	Tags         []Tag      `gorm:"many2many:problems_tags"`
	Sources      []Source   `gorm:"many2many:problems_sources"`
	Questions    []Question `gorm:"foreignKey:ProblemID"`
	Solutions    []Solution `gorm:"foreignKey:ProblemID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
type Source struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Label     string
	Color     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Difficulty struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Label     string
	Color     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Solution struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ProblemID uuid.UUID
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Question struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ProblemID uuid.UUID
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
