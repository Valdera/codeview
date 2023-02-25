package problemrepository

import (
	"codeview/internal/entity"
	"codeview/utils/pagination"
	"context"
	"fmt"
	"log"
)

func (r *problemRepository) CreateProblem(ctx context.Context, body *entity.Problem) (*entity.Problem, error) {
	if err := r.db.Model(&entity.Problem{}).Create(body).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - CreateProblem : %v\n", err)
		return nil, err
	}

	return body, nil
}

func (r *problemRepository) GetProblemById(ctx context.Context, id uint) (*entity.Problem, error) {
	var result entity.Problem
	var total int64

	if err := r.db.Model(&entity.Problem{}).Count(&total).
		Preload("Tags").Preload("Sources").Preload("Difficulty").
		Preload("Solutions").Preload("Questions").
		Where("id = ?", id).
		Find(&result).
		Error; err != nil {
		log.Printf("[ERROR] Problem Repository - GetProblemById : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Problem Repository - GetProblemById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *problemRepository) GetProblems(ctx context.Context, p *pagination.Pagination) ([]entity.Problem, error) {
	var results []entity.Problem
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Problem Repository - GetProblems : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Problem{}).Count(&total).Offset(p.GetOffset()).Limit(p.PageSize).
		Preload("Tags").Preload("Sources").Preload("Difficulty").
		Find(&results).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - GetProblems : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *problemRepository) GetProblemsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Problem, error) {
	var results []entity.Problem
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Problem Repository - GetProblemsByIds : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Problem{}).Where("id IN (?)", ids).
		Count(&total).Offset(p.GetOffset()).Limit(p.PageSize).
		Preload("Tags").Preload("Sources").Preload("Difficulty").
		Find(&results).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - GetProblemsByIds : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *problemRepository) UpdateProblemById(ctx context.Context, id uint, body *entity.Problem) (*entity.Problem, error) {
	var result entity.Problem
	var total int64

	if err := r.db.Model(&result).
		Where("id = ?", id).
		Updates(&entity.Problem{
			Title:        body.Title,
			Rating:       body.Rating,
			DifficultyID: body.DifficultyID,
		}).Count(&total).Error; err != nil {
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Problem Repository - UpdateProblemById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *problemRepository) DeleteProblemById(ctx context.Context, id uint) error {
	var total int64

	if err := r.db.Model(entity.Problem{}).
		Where("id = ?", id).
		Delete(&entity.Problem{}).Count(&total).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - DeleteProblemById : %v\n", err)
		return err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Problem Repository - DeleteProblemById : %v\n", err)
		return err
	}

	return nil
}
