package repository

import (
	"codeview/internal/entity"
	"codeview/internal/util/pagination"
	"context"
)

//go:generate mockery --name=SolutionRepository --case underscore --testonly
type SolutionRepository interface {
	CreateSolution(ctx context.Context, body *entity.Solution) (*entity.Solution, error)
	GetSolutions(ctx context.Context, p *pagination.Pagination) ([]entity.Solution, error)
	GetSolutionById(ctx context.Context, id uint) (*entity.Solution, error)
	GetSolutionsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Solution, error)
	UpdateSolutionById(ctx context.Context, id uint, body *entity.Solution) (*entity.Solution, error)
	DeleteSolutionById(ctx context.Context, id uint) error
}
