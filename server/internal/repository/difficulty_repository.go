package repository

import (
	"codeview/internal/entity"
	"codeview/internal/util/pagination"
	"context"

	"github.com/google/uuid"
)

//go:generate mockery --name=DifficultyRepository --case underscore --testonly
type DifficultyRepository interface {
	CreateDifficulty(ctx context.Context, body *entity.Difficulty) (*entity.Difficulty, error)
	GetDifficulties(ctx context.Context, p *pagination.Pagination) ([]entity.Difficulty, error)
	GetDifficultyById(ctx context.Context, id uuid.UUID) (*entity.Difficulty, error)
	GetDifficultiesByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]entity.Difficulty, error)
	UpdateDifficultyById(ctx context.Context, id uuid.UUID, body *entity.Difficulty) (*entity.Difficulty, error)
	DeleteDifficultyById(ctx context.Context, id uuid.UUID) error
}
