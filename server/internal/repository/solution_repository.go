package repository

import (
	"codeview/internal/entity"
	"codeview/internal/util/pagination"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//go:generate mockery --name=SolutionRepository --case underscore --testonly
type SolutionRepository interface {
	WithTrx(trxHandle *gorm.DB) SolutionRepository

	CreateSolution(ctx context.Context, body *entity.Solution) (*entity.Solution, error)
	GetSolutions(ctx context.Context, p *pagination.Pagination) ([]entity.Solution, error)
	GetSolutionById(ctx context.Context, id uuid.UUID) (*entity.Solution, error)
	GetSolutionsByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]entity.Solution, error)
	UpdateSolutionById(ctx context.Context, id uuid.UUID, body *entity.Solution) (*entity.Solution, error)
	DeleteSolutionById(ctx context.Context, id uuid.UUID) error
}
