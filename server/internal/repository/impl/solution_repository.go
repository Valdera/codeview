package repository

import (
	"codeview/config"
	"codeview/internal/entity"
	"codeview/internal/repository"
	"codeview/internal/util/pagination"
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type solutionRepository struct {
	cfg config.AppConfig
	db  *gorm.DB
}

func NewSolutionRepository(cfg config.AppConfig, db *gorm.DB) repository.SolutionRepository {
	return &solutionRepository{
		cfg,
		db,
	}
}

func (r *solutionRepository) CreateSolution(ctx context.Context, body *entity.Solution) (*entity.Solution, error) {
	if err := r.db.Model(&entity.Solution{}).
		Create(body).
		Error; err != nil {
		log.Printf("[ERROR] Solution Repository - CreateSolution : %v\n", err)
		return nil, err
	}

	return body, nil
}

func (r *solutionRepository) GetSolutionById(ctx context.Context, id uint) (*entity.Solution, error) {
	var result entity.Solution
	var total int64

	if err := r.db.Model(&entity.Solution{}).
		Count(&total).
		Where("id = ?", id).
		Find(&result).
		Error; err != nil {
		log.Printf("[ERROR] Solution Repository - GetSolutionById : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("tag with id %d does not exists", id)
		log.Printf("[ERROR] Solution Repository - GetSolutionById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *solutionRepository) GetSolutions(ctx context.Context, p *pagination.Pagination) ([]entity.Solution, error) {
	var results []entity.Solution
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Solution Repository - GetSolutions : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Solution{}).
		Count(&total).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Find(&results).
		Error; err != nil {
		log.Printf("[ERROR] Solution Repository - GetSolutions : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *solutionRepository) GetSolutionsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Solution, error) {
	var results []entity.Solution
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Solution Repository - GetSolutionsByIds : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Solution{}).
		Count(&total).
		Where("id IN (?)", ids).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Find(&results).
		Error; err != nil {
		log.Printf("[ERROR] Solution Repository - GetSolutionsByIds : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *solutionRepository) UpdateSolutionById(ctx context.Context, id uint, body *entity.Solution) (*entity.Solution, error) {
	var result entity.Solution
	var total int64

	if err := r.db.Model(&result).
		Count(&total).
		Where("id = ?", id).
		Updates(&entity.Solution{
			Content: body.Content,
		}).
		Error; err != nil {
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Solution Repository - UpdateSolutionById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *solutionRepository) DeleteSolutionById(ctx context.Context, id uint) error {
	var total int64

	if err := r.db.Model(&entity.Solution{}).
		Count(&total).
		Where("id = ?", id).
		Updates(&entity.Solution{
			DeletedAt: time.Now(),
		}).
		Error; err != nil {
		log.Printf("[ERROR] Solution Repository - DeleteSolutionById : %v\n", err)
		return err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Solution Repository - DeleteSolutionById : %v\n", err)
		return err
	}

	return nil
}
