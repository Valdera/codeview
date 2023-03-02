package service

import (
	"codeview/internal/dto/request"
	"codeview/internal/dto/response"
	"codeview/internal/util/pagination"
	"context"
)

//go:generate mockery --name=TagService --case underscore --testonly
type TagService interface {
	CreateTag(ctx context.Context, req *request.TagCreate) (*response.Tag, error)
	GetTags(ctx context.Context, p *pagination.Pagination) ([]response.Tag, error)
	GetTagById(ctx context.Context, id uint) (*response.Tag, error)
	GetTagsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]response.Tag, error)
	UpdateTagById(ctx context.Context, id uint, body *request.TagUpdate) (*response.Tag, error)
	DeleteTagById(ctx context.Context, id uint) error
}
