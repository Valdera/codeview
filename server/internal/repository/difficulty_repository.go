package repository

import (
	"codeview/internal/entity"
	"codeview/internal/util/pagination"
	"context"
)

//go:generate mockery --name=DifficultyRepository --case underscore --testonly
type DifficultyRepository interface {
	CreateDifficulty(ctx context.Context, body *entity.Difficulty) (*entity.Difficulty, error)
	GetDifficulties(ctx context.Context, p *pagination.Pagination) ([]entity.Difficulty, error)
	GetDifficultyById(ctx context.Context, id uint) (*entity.Difficulty, error)
	GetDifficultiesByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Difficulty, error)
	UpdateDifficultyById(ctx context.Context, id uint, body *entity.Difficulty) (*entity.Difficulty, error)
	DeleteDifficultyById(ctx context.Context, id uint) error
}
