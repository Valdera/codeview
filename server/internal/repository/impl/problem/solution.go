package problemrepository

import (
	"codeview/internal/entity"
	"codeview/utils/pagination"
	"context"
	"fmt"
	"log"
	"time"
)

func (r *problemRepository) CreateSolution(ctx context.Context, body *entity.Solution) (*entity.Solution, error) {
	if err := r.db.Create(body).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - CreateSolution : %v\n", err)
		return nil, err
	}

	return body, nil
}

func (r *problemRepository) GetSolutionById(ctx context.Context, id uint) (*entity.Solution, error) {
	var result entity.Solution
	var total int64

	if err := r.db.Model(&entity.Solution{}).
		Where("id = ?", id).
		Find(&result).
		Error; err != nil {
		log.Printf("[ERROR] Problem Repository - GetSolutionById : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("tag with id %d does not exists", id)
		log.Printf("[ERROR] Problem Repository - GetSolutionById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *problemRepository) GetSolutions(ctx context.Context, p *pagination.Pagination) ([]entity.Solution, error) {
	var results []entity.Solution
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Problem Repository - GetSolutions : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Problem{}).Count(&total).Offset(p.GetOffset()).Limit(p.PageSize).Find(&results).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - GetSolutions : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *problemRepository) GetSolutionsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Solution, error) {
	var results []entity.Solution
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Problem Repository - GetSolutionsByIds : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Problem{}).Where("id IN (?)", ids).
		Count(&total).Offset(p.GetOffset()).Limit(p.PageSize).Find(&results).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - GetSolutionsByIds : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *problemRepository) UpdateSolutionById(ctx context.Context, id uint, body *entity.Solution) (*entity.Solution, error) {
	var result entity.Solution
	var total int64

	if err := r.db.Model(&result).
		Where("id = ?", id).
		Updates(&entity.Solution{
			Content: body.Content,
		}).Count(&total).Error; err != nil {
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Problem Repository - UpdateSolutionById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *problemRepository) DeleteSolutionById(ctx context.Context, id uint) error {
	var total int64

	if err := r.db.Model(entity.Solution{}).
		Where("id = ?", id).
		Updates(entity.Solution{DeletedAt: time.Now()}).Count(&total).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - DeleteSolutionById : %v\n", err)
		return err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Problem Repository - DeleteSolutionById : %v\n", err)
		return err
	}

	return nil
}
