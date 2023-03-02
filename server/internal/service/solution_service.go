package service

import (
	"codeview/internal/dto/request"
	"codeview/internal/dto/response"
	"codeview/internal/util/pagination"
	"context"
)

//go:generate mockery --name=SolutionService --case underscore --testonly
type SolutionService interface {
	CreateSolution(ctx context.Context, req *request.SolutionCreate) (*response.Solution, error)
	GetSolutions(ctx context.Context, p *pagination.Pagination) ([]response.Solution, error)
	GetSolutionById(ctx context.Context, id uint) (*response.Solution, error)
	GetSolutionsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]response.Solution, error)
	UpdateSolutionById(ctx context.Context, id uint, body *request.SolutionUpdate) (*response.Solution, error)
	DeleteSolutionById(ctx context.Context, id uint) error
}
