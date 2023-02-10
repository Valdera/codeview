package entity

import (
	"time"
)

type Problem struct {
	ID           uint `gorm:"primaryKey"`
	Title        string
	Rating       *int
	DifficultyID uint
	Difficulty   Difficulty `gorm:"foreignKey:DifficultyID"`
	Tags         []Tag      `gorm:"many2many:problems_tags"`
	Sources      []Source   `gorm:"many2many:problems_sources"`
	Questions    []Question `gorm:"foreignKey:ProblemID"`
	Solutions    []Solution `gorm:"foreignKey:ProblemID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

type Tag struct {
	ID        uint `gorm:"primaryKey"`
	Label     string
	Color     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Source struct {
	ID        uint `gorm:"primaryKey"`
	Label     string
	Color     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Difficulty struct {
	ID        uint `gorm:"primaryKey"`
	Label     string
	Color     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Solution struct {
	ID        uint `gorm:"primaryKey"`
	ProblemID uint
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Question struct {
	ID        uint `gorm:"primaryKey"`
	ProblemID uint
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
