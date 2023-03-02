package repository

import (
	"codeview/internal/entity"
	"codeview/internal/util/pagination"
	"context"
)

//go:generate mockery --name=QuestionRepository --case underscore --testonly
type QuestionRepository interface {
	CreateQuestion(ctx context.Context, body *entity.Question) (*entity.Question, error)
	GetQuestions(ctx context.Context, p *pagination.Pagination) ([]entity.Question, error)
	GetQuestionById(ctx context.Context, id uint) (*entity.Question, error)
	GetQuestionsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Question, error)
	UpdateQuestionById(ctx context.Context, id uint, body *entity.Question) (*entity.Question, error)
	DeleteQuestionById(ctx context.Context, id uint) error
}
