package repository

import (
	"codeview/config"
	"codeview/internal/entity"
	"codeview/internal/repository"
	"codeview/internal/util/pagination"
	"context"
	"fmt"
	"log"

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

func (r *questionRepository) CreateQuestion(ctx context.Context, body *entity.Question) (*entity.Question, error) {
	if err := r.db.Model(&entity.Question{}).
		Create(body).
		Error; err != nil {
		log.Printf("[ERROR] Question Repository - CreateQuestion : %v\n", err)
		return nil, err
	}

	return body, nil
}

func (r *questionRepository) GetQuestionById(ctx context.Context, id uint) (*entity.Question, error) {
	var result entity.Question
	var total int64

	if err := r.db.Model(&entity.Question{}).
		Count(&total).
		Where("id = ?", id).
		Find(&result).
		Error; err != nil {
		log.Printf("[ERROR] Question Repository - GetQuestionById : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("tag with id %d does not exists", id)
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
		Count(&total).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Find(&results).
		Error; err != nil {
		log.Printf("[ERROR] Question Repository - GetQuestions : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *questionRepository) GetQuestionsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Question, error) {
	var results []entity.Question
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Question Repository - GetQuestionsByIds : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Question{}).
		Count(&total).
		Where("id IN (?)", ids).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Find(&results).
		Error; err != nil {
		log.Printf("[ERROR] Question Repository - GetQuestionsByIds : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *questionRepository) UpdateQuestionById(ctx context.Context, id uint, body *entity.Question) (*entity.Question, error) {
	var result entity.Question
	var total int64

	if err := r.db.Model(&result).
		Count(&total).
		Where("id = ?", id).
		Updates(&entity.Question{
			Content: body.Content,
		}).
		Error; err != nil {
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Question Repository - UpdateQuestionById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *questionRepository) DeleteQuestionById(ctx context.Context, id uint) error {
	var total int64

	if err := r.db.Model(&entity.Question{}).
		Count(&total).
		Where("id = ?", id).
		Delete(&entity.Question{}).
		Error; err != nil {
		log.Printf("[ERROR] Question Repository - DeleteQuestionById : %v\n", err)
		return err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Question Repository - DeleteQuestionById : %v\n", err)
		return err
	}

	return nil
}
