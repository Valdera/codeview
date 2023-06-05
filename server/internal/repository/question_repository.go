package repository

import (
	"codeview/internal/entity"
	"codeview/internal/util/pagination"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//go:generate mockery --name=QuestionRepository --case underscore --testonly
type QuestionRepository interface {
	WithTrx(trxHandle *gorm.DB) QuestionRepository

	CreateQuestion(ctx context.Context, body *entity.Question) (*entity.Question, error)
	GetQuestions(ctx context.Context, p *pagination.Pagination) ([]entity.Question, error)
	GetQuestionById(ctx context.Context, id uuid.UUID) (*entity.Question, error)
	GetQuestionsByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]entity.Question, error)
	UpdateQuestionById(ctx context.Context, id uuid.UUID, body *entity.Question) (*entity.Question, error)
	DeleteQuestionById(ctx context.Context, id uuid.UUID) error
}
