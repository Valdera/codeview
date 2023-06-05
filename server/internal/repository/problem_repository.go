package repository

import (
	"codeview/internal/entity"
	"codeview/internal/util/pagination"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//go:generate mockery --name=ProblemRepository --case underscore --testonly
type ProblemRepository interface {
	WithTrx(trxHandle *gorm.DB) ProblemRepository

	CreateProblem(ctx context.Context, body *entity.Problem) (*entity.Problem, error)
	GetProblems(ctx context.Context, p *pagination.Pagination) ([]entity.Problem, error)
	GetProblemById(ctx context.Context, id uuid.UUID) (*entity.Problem, error)
	GetProblemsByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]entity.Problem, error)
	UpdateProblemById(ctx context.Context, id uuid.UUID, body *entity.Problem) (*entity.Problem, error)
	DeleteProblemById(ctx context.Context, id uuid.UUID) error
}
