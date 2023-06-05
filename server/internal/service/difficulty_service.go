package service

import (
	"codeview/internal/dto/request"
	"codeview/internal/dto/response"
	"codeview/internal/util/pagination"
	"context"

	"github.com/google/uuid"
)

//go:generate mockery --name=DifficultyService --case underscore --testonly
type DifficultyService interface {
	CreateDifficulty(ctx context.Context, req *request.DifficultyCreate) (*response.Difficulty, error)
	GetDifficulties(ctx context.Context, p *pagination.Pagination) ([]response.Difficulty, error)
	GetDifficultyById(ctx context.Context, id uuid.UUID) (*response.Difficulty, error)
	GetDifficultiesByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]response.Difficulty, error)
	UpdateDifficultyById(ctx context.Context, id uuid.UUID, body *request.DifficultyUpdate) (*response.Difficulty, error)
	DeleteDifficultyById(ctx context.Context, id uuid.UUID) error
}
