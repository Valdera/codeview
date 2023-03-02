package service

import (
	"codeview/internal/dto/request"
	"codeview/internal/dto/response"
	"codeview/internal/util/pagination"
	"context"
)

//go:generate mockery --name=QuestionService --case underscore --testonly
type QuestionService interface {
	CreateQuestion(ctx context.Context, req *request.QuestionCreate) (*response.Question, error)
	GetQuestions(ctx context.Context, p *pagination.Pagination) ([]response.Question, error)
	GetQuestionById(ctx context.Context, id uint) (*response.Question, error)
	GetQuestionsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]response.Question, error)
	UpdateQuestionById(ctx context.Context, id uint, body *request.QuestionUpdate) (*response.Question, error)
	DeleteQuestionById(ctx context.Context, id uint) error
}
