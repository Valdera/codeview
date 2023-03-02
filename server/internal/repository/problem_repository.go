package repository

import (
	"codeview/internal/entity"
	"codeview/internal/util/pagination"
	"context"
)

//go:generate mockery --name=ProblemRepository --case underscore --testonly
type ProblemRepository interface {
	CreateProblem(ctx context.Context, body *entity.Problem) (*entity.Problem, error)
	GetProblems(ctx context.Context, p *pagination.Pagination) ([]entity.Problem, error)
	GetProblemById(ctx context.Context, id uint) (*entity.Problem, error)
	GetProblemsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Problem, error)
	UpdateProblemById(ctx context.Context, id uint, body *entity.Problem) (*entity.Problem, error)
	DeleteProblemById(ctx context.Context, id uint) error
}
