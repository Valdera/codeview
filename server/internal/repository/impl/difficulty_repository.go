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

type difficultyRepository struct {
	cfg config.AppConfig
	db  *gorm.DB
}

func NewDifficultyRepository(cfg config.AppConfig, db *gorm.DB) repository.DifficultyRepository {
	return &difficultyRepository{
		cfg,
		db,
	}
}

func (r *difficultyRepository) CreateDifficulty(ctx context.Context, body *entity.Difficulty) (*entity.Difficulty, error) {
	if err := r.db.Model(&entity.Difficulty{}).
		Create(body).
		Error; err != nil {
		log.Printf("[ERROR] Difficulty Repository - CreateDifficulty : %v\n", err)
		return nil, err
	}

	return body, nil
}

func (r *difficultyRepository) GetDifficultyById(ctx context.Context, id uuid.UUID) (*entity.Difficulty, error) {
	var result entity.Difficulty
	var total int64

	if err := r.db.Model(&entity.Difficulty{}).
		Where("id = ?", id).
		Find(&result).
		Count(&total).
		Error; err != nil {
		log.Printf("[ERROR] Difficulty Repository - GetDifficultyById : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("difficulty with id %s does not exists", id)
		log.Printf("[ERROR] Difficulty Repository - GetDifficultyById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *difficultyRepository) GetDifficulties(ctx context.Context, p *pagination.Pagination) ([]entity.Difficulty, error) {
	var results []entity.Difficulty
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Difficulty Repository - GetDifficulties : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Difficulty{}).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Find(&results).
		Count(&total).
		Error; err != nil {
		log.Printf("[ERROR] Difficulty Repository - GetDifficulties : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *difficultyRepository) GetDifficultiesByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]entity.Difficulty, error) {
	var results []entity.Difficulty
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Difficulty Repository - GetDifficultiesByIds : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Difficulty{}).
		Where("id IN (?)", ids).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Find(&results).
		Count(&total).
		Error; err != nil {
		log.Printf("[ERROR] Difficulty Repository - GetDifficultiesByIds : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *difficultyRepository) UpdateDifficultyById(ctx context.Context, id uuid.UUID, body *entity.Difficulty) (*entity.Difficulty, error) {
	var result entity.Difficulty
	var total int64

	if err := r.db.Model(&result).
		Where("id = ?", id).
		Updates(&entity.Difficulty{
			Label: body.Label,
			Color: body.Color,
		}).
		Count(&total).
		Error; err != nil {
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %s does not exists", id)
		log.Printf("[ERROR] Difficulty Repository - UpdateDifficultyById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *difficultyRepository) DeleteDifficultyById(ctx context.Context, id uuid.UUID) error {
	var total int64

	if err := r.db.Model(&entity.Difficulty{}).
		Where("id = ?", id).
		Count(&total).
		Delete(&entity.Difficulty{}).
		Error; err != nil {
		log.Printf("[ERROR] Difficulty Repository - DeleteDifficultyById : %v\n", err)
		return err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %s does not exists", id)
		log.Printf("[ERROR] Difficulty Repository - DeleteDifficultyById : %v\n", err)
		return err
	}

	return nil
}
