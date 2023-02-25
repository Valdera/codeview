package problemrepository

import (
	"codeview/internal/entity"
	"codeview/utils/pagination"
	"context"
	"fmt"
	"log"
	"time"
)

func (r *problemRepository) CreateDifficulty(ctx context.Context, body *entity.Difficulty) (*entity.Difficulty, error) {
	if err := r.db.Create(body).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - CreateDifficulty : %v\n", err)
		return nil, err
	}

	return body, nil
}

func (r *problemRepository) GetDifficultyById(ctx context.Context, id uint) (*entity.Difficulty, error) {
	var result entity.Difficulty
	var total int64

	if err := r.db.Model(&entity.Difficulty{}).
		Count(&total).
		Where("id = ?", id).
		Find(&result).
		Error; err != nil {
		log.Printf("[ERROR] Problem Repository - GetDifficultyById : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("difficulty with id %d does not exists", id)
		log.Printf("[ERROR] Problem Repository - GetDifficultyById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *problemRepository) GetDifficulties(ctx context.Context, p *pagination.Pagination) ([]entity.Difficulty, error) {
	var results []entity.Difficulty
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Problem Repository - GetDifficulties : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Problem{}).Count(&total).Offset(p.GetOffset()).Limit(p.PageSize).Find(&results).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - GetDifficulties : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *problemRepository) GetDifficultiesByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Difficulty, error) {
	var results []entity.Difficulty
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Problem Repository - GetDifficultiesByIds : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Problem{}).Where("id IN (?)", ids).
		Count(&total).Offset(p.GetOffset()).Limit(p.PageSize).Find(&results).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - GetDifficultiesByIds : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *problemRepository) UpdateDifficultyById(ctx context.Context, id uint, body *entity.Difficulty) (*entity.Difficulty, error) {
	var result entity.Difficulty
	var total int64

	if err := r.db.Model(&result).
		Where("id = ?", id).
		Updates(&entity.Difficulty{
			Label: body.Label,
			Color: body.Color,
		}).Count(&total).Error; err != nil {
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Problem Repository - UpdateDifficultyById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *problemRepository) DeleteDifficultyById(ctx context.Context, id uint) error {
	var total int64

	if err := r.db.Model(entity.Difficulty{}).
		Where("id = ?", id).
		Updates(entity.Difficulty{DeletedAt: time.Now()}).Count(&total).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - DeleteDifficultyById : %v\n", err)
		return err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Problem Repository - DeleteDifficultyById : %v\n", err)
		return err
	}

	return nil
}
