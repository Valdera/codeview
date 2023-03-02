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

type sourceRepository struct {
	cfg config.AppConfig
	db  *gorm.DB
}

func NewSourceRepository(cfg config.AppConfig, db *gorm.DB) repository.SourceRepository {
	return &sourceRepository{
		cfg,
		db,
	}
}

func (r *sourceRepository) CreateSource(ctx context.Context, body *entity.Source) (*entity.Source, error) {
	if err := r.db.Model(&entity.Source{}).
		Create(body).
		Error; err != nil {
		log.Printf("[ERROR] Source Repository - CreateSource : %v\n", err)
		return nil, err
	}

	return body, nil
}

func (r *sourceRepository) GetSourceById(ctx context.Context, id uint) (*entity.Source, error) {
	var result entity.Source
	var total int64

	if err := r.db.Model(&entity.Source{}).
		Count(&total).
		Where("id = ?", id).
		Find(&result).
		Error; err != nil {
		log.Printf("[ERROR] Source Repository - GetSourceById : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("tag with id %d does not exists", id)
		log.Printf("[ERROR] Source Repository - GetSourceById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *sourceRepository) GetSources(ctx context.Context, p *pagination.Pagination) ([]entity.Source, error) {
	var results []entity.Source
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Source Repository - GetSources : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Source{}).
		Count(&total).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Find(&results).
		Error; err != nil {
		log.Printf("[ERROR] Source Repository - GetSources : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *sourceRepository) GetSourcesByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Source, error) {
	var results []entity.Source
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Source Repository - GetSourcesByIds : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Source{}).
		Count(&total).
		Where("id IN (?)", ids).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Find(&results).
		Error; err != nil {
		log.Printf("[ERROR] Source Repository - GetSourcesByIds : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *sourceRepository) UpdateSourceById(ctx context.Context, id uint, body *entity.Source) (*entity.Source, error) {
	var result entity.Source
	var total int64

	if err := r.db.Model(&result).
		Count(&total).
		Where("id = ?", id).
		Updates(&entity.Source{
			Label: body.Label,
			Color: body.Color,
		}).
		Error; err != nil {
		log.Printf("[ERROR] Source Repository - UpdateSourceById : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Source Repository - UpdateSourceById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *sourceRepository) DeleteSourceById(ctx context.Context, id uint) error {
	var total int64

	if err := r.db.Model(&entity.Source{}).
		Count(&total).
		Where("id = ?", id).
		Updates(&entity.Source{
			DeletedAt: time.Now(),
		}).
		Error; err != nil {
		log.Printf("[ERROR] Source Repository - DeleteSourceById : %v\n", err)
		return err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Source Repository - DeleteSourceById : %v\n", err)
		return err
	}

	return nil
}
