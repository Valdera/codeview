package problemrepository

import (
	"codeview/internal/entity"
	"codeview/internal/util/pagination"
	"context"
	"fmt"
	"log"
	"time"
)

func (r *problemRepository) CreateSource(ctx context.Context, body *entity.Source) (*entity.Source, error) {
	if err := r.db.Create(body).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - CreateSource : %v\n", err)
		return nil, err
	}

	return body, nil
}

func (r *problemRepository) GetSourceById(ctx context.Context, id uint) (*entity.Source, error) {
	var result entity.Source
	var total int64

	if err := r.db.Model(&entity.Source{}).
		Where("id = ?", id).
		Find(&result).
		Error; err != nil {
		log.Printf("[ERROR] Problem Repository - GetSourceById : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("tag with id %d does not exists", id)
		log.Printf("[ERROR] Problem Repository - GetSourceById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *problemRepository) GetSources(ctx context.Context, p *pagination.Pagination) ([]entity.Source, error) {
	var results []entity.Source
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Problem Repository - GetSources : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Problem{}).Count(&total).Offset(p.GetOffset()).Limit(p.PageSize).Find(&results).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - GetSources : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *problemRepository) GetSourcesByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Source, error) {
	var results []entity.Source
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Problem Repository - GetSourcesByIds : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Problem{}).Where("id IN (?)", ids).
		Count(&total).Offset(p.GetOffset()).Limit(p.PageSize).Find(&results).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - GetSourcesByIds : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *problemRepository) UpdateSourceById(ctx context.Context, id uint, body *entity.Source) (*entity.Source, error) {
	var result entity.Source
	var total int64

	if err := r.db.Model(&result).
		Where("id = ?", id).
		Updates(&entity.Source{
			Label: body.Label,
			Color: body.Color,
		}).Count(&total).Error; err != nil {
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Problem Repository - UpdateSourceById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *problemRepository) DeleteSourceById(ctx context.Context, id uint) error {
	var total int64

	if err := r.db.Model(entity.Source{}).
		Where("id = ?", id).
		Updates(entity.Source{DeletedAt: time.Now()}).Count(&total).Error; err != nil {
		log.Printf("[ERROR] Problem Repository - DeleteSourceById : %v\n", err)
		return err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] Problem Repository - DeleteSourceById : %v\n", err)
		return err
	}

	return nil
}
