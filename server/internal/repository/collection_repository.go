package repository

import (
	"codeview/internal/entity"
	"codeview/internal/util/pagination"
	"context"

	"github.com/google/uuid"
)

//go:generate mockery --name=CollectionRepository --case underscore --testonly
type CollectionRepository interface {
	CreateCollection(ctx context.Context, body *entity.Collection) (*entity.Collection, error)
	GetCollections(ctx context.Context, p *pagination.Pagination) ([]entity.Collection, error)
	GetCollectionById(ctx context.Context, id uuid.UUID) (*entity.Collection, error)
	GetCollectionsByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]entity.Collection, error)
	UpdateCollectionById(ctx context.Context, id uuid.UUID, body *entity.Collection) (*entity.Collection, error)
	DeleteCollectionById(ctx context.Context, id uuid.UUID) error

	CreateCollectionItem(ctx context.Context, body *entity.CollectionItem) (*entity.CollectionItem, error)
	GetCollectionItems(ctx context.Context, p *pagination.Pagination) ([]entity.CollectionItem, error)
	GetCollectionItemById(ctx context.Context, id uuid.UUID) (*entity.CollectionItem, error)
	GetCollectionItemsByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]entity.CollectionItem, error)
	UpdateCollectionItemById(ctx context.Context, id uuid.UUID, body *entity.CollectionItem) (*entity.CollectionItem, error)
	DeleteCollectionItemById(ctx context.Context, id uuid.UUID) error
}
