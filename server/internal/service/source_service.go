package service

import (
	"codeview/internal/dto/request"
	"codeview/internal/dto/response"
	"codeview/internal/util/pagination"
	"context"
)

//go:generate mockery --name=SourceService --case underscore --testonly
type SourceService interface {
	CreateSource(ctx context.Context, req *request.SourceCreate) (*response.Source, error)
	GetSources(ctx context.Context, p *pagination.Pagination) ([]response.Source, error)
	GetSourceById(ctx context.Context, id uint) (*response.Source, error)
	GetSourcesByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]response.Source, error)
	UpdateSourceById(ctx context.Context, id uint, body *request.SourceUpdate) (*response.Source, error)
	DeleteSourceById(ctx context.Context, id uint) error
}
