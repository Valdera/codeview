package repository

import (
	"codeview/internal/entity"
	"codeview/internal/util/pagination"
	"context"
)

//go:generate mockery --name=TagRepository --case underscore --testonly
type TagRepository interface {
	CreateTag(ctx context.Context, body *entity.Tag) (*entity.Tag, error)
	GetTags(ctx context.Context, p *pagination.Pagination) ([]entity.Tag, error)
	GetTagById(ctx context.Context, id uint) (*entity.Tag, error)
	GetTagsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Tag, error)
	UpdateTagById(ctx context.Context, id uint, body *entity.Tag) (*entity.Tag, error)
	DeleteTagById(ctx context.Context, id uint) error
}
