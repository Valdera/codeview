package repository

import (
	"codeview/internal/entity"
	"codeview/internal/util/pagination"
	"context"

	"github.com/google/uuid"
)

//go:generate mockery --name=TagRepository --case underscore --testonly
type TagRepository interface {
	CreateTag(ctx context.Context, body *entity.Tag) (*entity.Tag, error)
	GetTags(ctx context.Context, p *pagination.Pagination) ([]entity.Tag, error)
	GetTagById(ctx context.Context, id uuid.UUID) (*entity.Tag, error)
	GetTagsByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]entity.Tag, error)
	UpdateTagById(ctx context.Context, id uuid.UUID, body *entity.Tag) (*entity.Tag, error)
	DeleteTagById(ctx context.Context, id uuid.UUID) error

	GetTagsByType(ctx context.Context, tagType entity.TagType, p *pagination.Pagination) ([]entity.Tag, error)
}
