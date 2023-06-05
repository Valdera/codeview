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

func (r *tagRepository) GetTagById(ctx context.Context, id uuid.UUID) (*entity.Tag, error) {
	var result entity.Tag
	var total int64

	if err := r.db.Model(&entity.Tag{}).
		Where("id = ?", id).
		Find(&result).
		Count(&total).
		Error; err != nil {
		log.Printf("[ERROR] Tag Repository - GetTagById : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("tag with id %s does not exists", id)
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
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Find(&results).
		Count(&total).
		Error; err != nil {
		log.Printf("[ERROR] Tag Repository - GetTags : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *tagRepository) GetTagsByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]entity.Tag, error) {
	var results []entity.Tag
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Tag Repository - GetTagsByIds : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Tag{}).
		Where("id IN (?)", ids).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Find(&results).
		Count(&total).
		Error; err != nil {
		log.Printf("[ERROR] Tag Repository - GetTagsByIds : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *tagRepository) UpdateTagById(ctx context.Context, id uuid.UUID, body *entity.Tag) (*entity.Tag, error) {
	var result entity.Tag
	var total int64

	if err := r.db.Model(&result).
		Where("id = ?", id).
		Updates(&entity.Tag{
			Label: body.Label,
			Color: body.Color,
		}).
		Count(&total).
		Error; err != nil {
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %s does not exists", id)
		log.Printf("[ERROR] Tag Repository - UpdateTagById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *tagRepository) DeleteTagById(ctx context.Context, id uuid.UUID) error {
	var total int64

	if err := r.db.Model(&entity.Tag{}).
		Where("id = ?", id).
		Count(&total).
		Delete(&entity.Tag{}).
		Error; err != nil {
		log.Printf("[ERROR] Tag Repository - DeleteTagById : %v\n", err)
		return err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %s does not exists", id)
		log.Printf("[ERROR] Tag Repository - DeleteTagById : %v\n", err)
		return err
	}

	return nil
}

func (r *tagRepository) GetTagsByType(ctx context.Context, tagType entity.TagType, p *pagination.Pagination) ([]entity.Tag, error) {
	var results []entity.Tag
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Tag Repository - GetTags : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Tag{}).
		Where("tag_type = ?", tagType).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Find(&results).
		Count(&total).
		Error; err != nil {
		log.Printf("[ERROR] Tag Repository - GetTags : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}
