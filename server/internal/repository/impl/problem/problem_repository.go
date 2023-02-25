package problemrepository

import (
	"codeview/config"
	"codeview/internal/repository"

	"gorm.io/gorm"
)

type problemRepository struct {
	config config.AppConfig
	db     *gorm.DB
}

func New(config config.AppConfig, db *gorm.DB) repository.ProblemRepository {
	return &problemRepository{
		config,
		db,
	}
}
