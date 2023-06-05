package repository

import (
	"codeview/internal/entity"
	"codeview/internal/util/pagination"
	"context"

	"github.com/google/uuid"
)

//go:generate mockery --name=SourceRepository --case underscore --testonly
type SourceRepository interface {
	CreateSource(ctx context.Context, body *entity.Source) (*entity.Source, error)
	GetSources(ctx context.Context, p *pagination.Pagination) ([]entity.Source, error)
	GetSourceById(ctx context.Context, id uuid.UUID) (*entity.Source, error)
	GetSourcesByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]entity.Source, error)
	UpdateSourceById(ctx context.Context, id uuid.UUID, body *entity.Source) (*entity.Source, error)
	DeleteSourceById(ctx context.Context, id uuid.UUID) error
}
