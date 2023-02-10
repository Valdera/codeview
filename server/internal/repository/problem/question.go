package repository

import (
	"codeview/internal/entity"
	"codeview/utils/pagination"
	"context"
	"fmt"
	"log"
	"time"
)

func (r *problemRepository) CreateQuestion(ctx context.Context, body *entity.Question) (*entity.Question, error) {
	if err := r.db.Create(body).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - CreateQuestion : %v\n", err)
		return nil, err
	}

	return body, nil
}

func (r *problemRepository) GetQuestionById(ctx context.Context, id uint) (*entity.Question, error) {
	var result entity.Question
	var total int64

	if err := r.db.Model(&entity.Question{}).
		Where("id = ?", id).
		Find(&result).
		Error; err != nil {
		log.Printf("[ERROR] Problem Repository - GetQuestionById : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("tag with id %d does not exists", id)
		log.Printf("[ERROR] Problem Repository - GetQuestionById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *problemRepository) GetQuestions(ctx context.Context, p *pagination.Pagination) ([]entity.Question, error) {
	var results []entity.Question
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Problem Repository - GetQuestions : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Problem{}).Count(&total).Offset(p.GetOffset()).Limit(p.PageSize).Find(&results).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - GetQuestions : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *problemRepository) GetQuestionsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Question, error) {
	var results []entity.Question
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Problem Repository - GetQuestionsByIds : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Problem{}).Where("id IN (?)", ids).
		Count(&total).Offset(p.GetOffset()).Limit(p.PageSize).Find(&results).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - GetQuestionsByIds : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *problemRepository) UpdateQuestionById(ctx context.Context, id uint, body *entity.Question) (*entity.Question, error) {
	var result entity.Question
	var total int64

	if err := r.db.Model(&result).
		Where("id = ?", id).
		Updates(&entity.Question{
			Content: body.Content,
		}).Count(&total).Error; err != nil {
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Problem Repository - UpdateQuestionById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *problemRepository) DeleteQuestionById(ctx context.Context, id uint) error {
	var total int64

	if err := r.db.Model(entity.Question{}).
		Where("id = ?", id).
		Updates(entity.Question{DeletedAt: time.Now()}).Count(&total).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - DeleteQuestionById : %v\n", err)
		return err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Problem Repository - DeleteQuestionById : %v\n", err)
		return err
	}

	return nil
}
