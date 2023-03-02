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

type tagRepository struct {
	cfg config.AppConfig
	db  *gorm.DB
}

func NewTagRepository(cfg config.AppConfig, db *gorm.DB) repository.TagRepository {
	return &tagRepository{
		cfg,
		db,
	}
}

func (r *tagRepository) CreateTag(ctx context.Context, body *entity.Tag) (*entity.Tag, error) {
	if err := r.db.Model(&entity.Tag{}).
		Create(body).
		Error; err != nil {
		log.Printf("[ERROR] Tag Repository - CreateTag : %v\n", err)
		return nil, err
	}

	return body, nil
}

func (r *tagRepository) GetTagById(ctx context.Context, id uint) (*entity.Tag, error) {
	var result entity.Tag
	var total int64

	if err := r.db.Model(&entity.Tag{}).
		Count(&total).
		Where("id = ?", id).
		Find(&result).
		Error; err != nil {
		log.Printf("[ERROR] Tag Repository - GetTagById : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("tag with id %d does not exists", id)
		log.Printf("[ERROR] Tag Repository - GetTagById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *tagRepository) GetTags(ctx context.Context, p *pagination.Pagination) ([]entity.Tag, error) {
	var results []entity.Tag
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Tag Repository - GetTags : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Tag{}).
		Count(&total).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Find(&results).
		Error; err != nil {
		log.Printf("[ERROR] Tag Repository - GetTags : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *tagRepository) GetTagsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Tag, error) {
	var results []entity.Tag
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Tag Repository - GetTagsByIds : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Tag{}).
		Count(&total).
		Where("id IN (?)", ids).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Find(&results).
		Error; err != nil {
		log.Printf("[ERROR] Tag Repository - GetTagsByIds : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *tagRepository) UpdateTagById(ctx context.Context, id uint, body *entity.Tag) (*entity.Tag, error) {
	var result entity.Tag
	var total int64

	if err := r.db.Model(&result).
		Count(&total).
		Where("id = ?", id).
		Updates(&entity.Tag{
			Label: body.Label,
			Color: body.Color,
		}).
		Error; err != nil {
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Tag Repository - UpdateTagById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *tagRepository) DeleteTagById(ctx context.Context, id uint) error {
	var total int64

	if err := r.db.Model(&entity.Tag{}).
		Count(&total).
		Where("id = ?", id).
		Updates(&entity.Tag{
			DeletedAt: time.Now(),
		}).
		Error; err != nil {
		log.Printf("[ERROR] Tag Repository - DeleteTagById : %v\n", err)
		return err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Tag Repository - DeleteTagById : %v\n", err)
		return err
	}

	return nil
}
