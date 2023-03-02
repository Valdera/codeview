package repository

import (
	"codeview/internal/entity"
	"codeview/internal/util/pagination"
	"context"
)

//go:generate mockery --name=SourceRepository --case underscore --testonly
type SourceRepository interface {
	CreateSource(ctx context.Context, body *entity.Source) (*entity.Source, error)
	GetSources(ctx context.Context, p *pagination.Pagination) ([]entity.Source, error)
	GetSourceById(ctx context.Context, id uint) (*entity.Source, error)
	GetSourcesByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Source, error)
	UpdateSourceById(ctx context.Context, id uint, body *entity.Source) (*entity.Source, error)
	DeleteSourceById(ctx context.Context, id uint) error
}
