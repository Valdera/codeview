package repository

import (
	"codeview/config"
	"codeview/internal/entity"
	"codeview/internal/repository"
	"codeview/internal/util/pagination"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type questionRepository struct {
	cfg config.AppConfig
	db  *gorm.DB
}

func NewQuestionRepository(cfg config.AppConfig, db *gorm.DB) repository.QuestionRepository {
	return &questionRepository{
		cfg,
		db,
	}
}

func (r *questionRepository) WithTrx(trxHandle *gorm.DB) repository.QuestionRepository {
	return &questionRepository{
		cfg: r.cfg,
		db:  trxHandle,
	}
}

func (r *questionRepository) CreateQuestion(ctx context.Context, body *entity.Question) (*entity.Question, error) {
	if err := r.db.Model(&entity.Question{}).
		Create(body).
		Error; err != nil {
		log.Printf("[ERROR] Question Repository - CreateQuestion : %v\n", err)
		return nil, err
	}

	return body, nil
}

func (r *questionRepository) GetQuestionById(ctx context.Context, id uuid.UUID) (*entity.Question, error) {
	var result entity.Question
	var total int64

	if err := r.db.Model(&entity.Question{}).
		Where("id = ?", id).
		Find(&result).
		Count(&total).
		Error; err != nil {
		log.Printf("[ERROR] Question Repository - GetQuestionById : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("tag with id %s does not exists", id)
		log.Printf("[ERROR] Question Repository - GetQuestionById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *questionRepository) GetQuestions(ctx context.Context, p *pagination.Pagination) ([]entity.Question, error) {
	var results []entity.Question
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Question Repository - GetQuestions : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Question{}).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Find(&results).
		Count(&total).
		Error; err != nil {
		log.Printf("[ERROR] Question Repository - GetQuestions : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *questionRepository) GetQuestionsByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]entity.Question, error) {
	var results []entity.Question
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Question Repository - GetQuestionsByIds : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Question{}).
		Where("id IN (?)", ids).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Find(&results).
		Count(&total).
		Error; err != nil {
		log.Printf("[ERROR] Question Repository - GetQuestionsByIds : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *questionRepository) UpdateQuestionById(ctx context.Context, id uuid.UUID, body *entity.Question) (*entity.Question, error) {
	var result entity.Question
	var total int64

	if err := r.db.Model(&result).
		Where("id = ?", id).
		Updates(&entity.Question{
			Content: body.Content,
		}).
		Count(&total).
		Error; err != nil {
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %s does not exists", id)
		log.Printf("[ERROR] Question Repository - UpdateQuestionById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *questionRepository) DeleteQuestionById(ctx context.Context, id uuid.UUID) error {
	var total int64

	if err := r.db.Model(&entity.Question{}).
		Where("id = ?", id).
		Count(&total).
		Delete(&entity.Question{}).
		Error; err != nil {
		log.Printf("[ERROR] Question Repository - DeleteQuestionById : %v\n", err)
		return err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %s does not exists", id)
		log.Printf("[ERROR] Question Repository - DeleteQuestionById : %v\n", err)
		return err
	}

	return nil
}
