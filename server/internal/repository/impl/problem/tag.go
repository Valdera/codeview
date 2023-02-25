package problemrepository

import (
	"codeview/internal/entity"
	"codeview/internal/util/pagination"
	"context"
	"fmt"
	"log"
	"time"
)

func (r *problemRepository) CreateTag(ctx context.Context, body *entity.Tag) (*entity.Tag, error) {
	if err := r.db.Create(body).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - CreateTag : %v\n", err)
		return nil, err
	}

	return body, nil
}

func (r *problemRepository) GetTagById(ctx context.Context, id uint) (*entity.Tag, error) {
	var result entity.Tag
	var total int64

	if err := r.db.Model(&entity.Tag{}).
		Where("id = ?", id).
		Find(&result).
		Error; err != nil {
		log.Printf("[ERROR] Problem Repository - GetTagById : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("tag with id %d does not exists", id)
		log.Printf("[ERROR] Problem Repository - GetTagById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *problemRepository) GetTags(ctx context.Context, p *pagination.Pagination) ([]entity.Tag, error) {
	var results []entity.Tag
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Problem Repository - GetTags : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Problem{}).Count(&total).Offset(p.GetOffset()).Limit(p.PageSize).Find(&results).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - GetTags : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *problemRepository) GetTagsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Tag, error) {
	var results []entity.Tag
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Problem Repository - GetTagsByIds : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Problem{}).Where("id IN (?)", ids).
		Count(&total).Offset(p.GetOffset()).Limit(p.PageSize).Find(&results).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - GetTagsByIds : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *problemRepository) UpdateTagById(ctx context.Context, id uint, body *entity.Tag) (*entity.Tag, error) {
	var result entity.Tag
	var total int64

	if err := r.db.Model(&result).
		Where("id = ?", id).
		Updates(&entity.Tag{
			Label: body.Label,
			Color: body.Color,
		}).Count(&total).Error; err != nil {
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Problem Repository - UpdateTagById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *problemRepository) DeleteTagById(ctx context.Context, id uint) error {
	var total int64

	if err := r.db.Model(entity.Tag{}).
		Where("id = ?", id).
		Updates(entity.Tag{DeletedAt: time.Now()}).Count(&total).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - DeleteTagById : %v\n", err)
		return err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Problem Repository - DeleteTagById : %v\n", err)
		return err
	}

	return nil
}
