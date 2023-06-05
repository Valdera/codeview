package service

import (
	"codeview/internal/dto/request"
	"codeview/internal/dto/response"
	"codeview/internal/util/pagination"
	"context"

	"github.com/google/uuid"
)

//go:generate mockery --name=ProblemService --case underscore --testonly
type ProblemService interface {
	CreateProblem(ctx context.Context, req *request.ProblemCreate) (*response.Problem, error)
	GetProblems(ctx context.Context, p *pagination.Pagination) ([]response.Problem, error)
	GetProblemById(ctx context.Context, id uuid.UUID) (*response.Problem, error)
	GetProblemsByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]response.Problem, error)
	UpdateProblemById(ctx context.Context, id uuid.UUID, body *request.ProblemUpdate) (*response.Problem, error)
	DeleteProblemById(ctx context.Context, id uuid.UUID) error
}
