package service

import (
	"codeview/internal/dto/request"
	"codeview/internal/dto/response"
	"codeview/internal/util/pagination"
	"context"
)

//go:generate mockery --name=DifficultyService --case underscore --testonly
type DifficultyService interface {
	CreateDifficulty(ctx context.Context, req *request.DifficultyCreate) (*response.Difficulty, error)
	GetDifficulties(ctx context.Context, p *pagination.Pagination) ([]response.Difficulty, error)
	GetDifficultyById(ctx context.Context, id uint) (*response.Difficulty, error)
	GetDifficultiesByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]response.Difficulty, error)
	UpdateDifficultyById(ctx context.Context, id uint, body *request.DifficultyUpdate) (*response.Difficulty, error)
	DeleteDifficultyById(ctx context.Context, id uint) error
}
