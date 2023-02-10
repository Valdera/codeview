package repository

import (
	"codeview/domain"

	"gorm.io/gorm"
)

type problemRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ProblemRepository {
	return &problemRepository{
		db,
	}
}
