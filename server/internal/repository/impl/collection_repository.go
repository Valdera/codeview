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

type collectionRepository struct {
	cfg config.AppConfig
	db  *gorm.DB
}

func NewCollectionRepository(cfg config.AppConfig, db *gorm.DB) repository.CollectionRepository {
	return &collectionRepository{
		cfg,
		db,
	}
}

func (r *collectionRepository) CreateCollection(ctx context.Context, body *entity.Collection) (*entity.Collection, error) {
	if err := r.db.Model(&entity.Collection{}).
		Create(body).
		Error; err != nil {
		log.Printf("[ERROR] Collection Repository - CreateCollection : %v\n", err)
		return nil, err
	}
	return body, nil
}

func (r *collectionRepository) GetCollectionById(ctx context.Context, id uuid.UUID) (*entity.Collection, error) {
	var result entity.Collection
	var total int64

	if err := r.db.Model(&entity.Collection{}).
		Preload("Collection").Preload("Tags").
		Where("id = ?", id).
		Find(&result).
		Count(&total).
		Error; err != nil {
		log.Printf("[ERROR] Collection Repository - GetCollectionById : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("collection with id %s does not exists", id)
		log.Printf("[ERROR] Collection Repository - GetCollectionById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *collectionRepository) GetCollections(ctx context.Context, p *pagination.Pagination) ([]entity.Collection, error) {
	var results []entity.Collection
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Collection Repository - GetCollections : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Collection{}).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Preload("Tags").
		Find(&results).
		Count(&total).
		Error; err != nil {
		log.Printf("[ERROR] Collection Repository - GetCollections : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *collectionRepository) GetCollectionsByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]entity.Collection, error) {
	var results []entity.Collection
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Collection Repository - GetCollectionsByIds : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Collection{}).
		Where("id IN (?)", ids).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Preload("Tags").Preload("Sources").
		Preload("Difficulty").
		Find(&results).
		Count(&total).
		Error; err != nil {
		log.Printf("[ERROR] Collection Repository - GetCollectionsByIds : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *collectionRepository) UpdateCollectionById(ctx context.Context, id uuid.UUID, body *entity.Collection) (*entity.Collection, error) {
	var result entity.Collection
	var total int64

	if err := r.db.Model(&result).
		Where("id = ?", id).
		Updates(&entity.Collection{
			Title: body.Title,
		}).
		Count(&total).
		Error; err != nil {
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("collection with id %s does not exists", id)
		log.Printf("[ERROR] Collection Repository - UpdateCollectionById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *collectionRepository) DeleteCollectionById(ctx context.Context, id uuid.UUID) error {
	var total int64

	if err := r.db.Model(&entity.Collection{}).
		Where("id = ?", id).
		Count(&total).
		Delete(&entity.Collection{}).
		Error; err != nil {
		log.Printf("[ERROR] Collection Repository - DeleteCollectionById : %v\n", err)
		return err
	}

	if total == 0 {
		err := fmt.Errorf("collection with id %s does not exists", id)
		log.Printf("[ERROR] Collection Repository - DeleteCollectionById : %v\n", err)
		return err
	}

	return nil
}

func (r *collectionRepository) CreateCollectionItem(ctx context.Context, body *entity.CollectionItem) (*entity.CollectionItem, error) {
	if err := r.db.Model(&entity.CollectionItem{}).
		Create(body).
		Error; err != nil {
		log.Printf("[ERROR] CollectionItem Repository - CreateCollectionItem : %v\n", err)
		return nil, err
	}

	return body, nil
}

func (r *collectionRepository) GetCollectionItemById(ctx context.Context, id uuid.UUID) (*entity.CollectionItem, error) {
	var result entity.CollectionItem
	var total int64

	if err := r.db.Model(&entity.CollectionItem{}).
		Where("id = ?", id).
		Find(&result).
		Count(&total).
		Error; err != nil {
		log.Printf("[ERROR] CollectionItem Repository - GetCollectionItemById : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("tag with id %s does not exists", id)
		log.Printf("[ERROR] CollectionItem Repository - GetCollectionItemById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *collectionRepository) GetCollectionItems(ctx context.Context, p *pagination.Pagination) ([]entity.CollectionItem, error) {
	var results []entity.CollectionItem
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] CollectionItem Repository - GetCollectionItems : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.CollectionItem{}).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Find(&results).
		Count(&total).
		Error; err != nil {
		log.Printf("[ERROR] CollectionItem Repository - GetCollectionItems : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *collectionRepository) GetCollectionItemsByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]entity.CollectionItem, error) {
	var results []entity.CollectionItem
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] CollectionItem Repository - GetCollectionItemsByIds : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.CollectionItem{}).
		Where("id IN (?)", ids).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Find(&results).
		Count(&total).
		Error; err != nil {
		log.Printf("[ERROR] CollectionItem Repository - GetCollectionItemsByIds : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *collectionRepository) UpdateCollectionItemById(ctx context.Context, id uuid.UUID, body *entity.CollectionItem) (*entity.CollectionItem, error) {
	var result entity.CollectionItem
	var total int64

	if err := r.db.Model(&result).
		Where("id = ?", id).
		Updates(&entity.CollectionItem{
			ItemID: body.ItemID,
		}).
		Count(&total).
		Error; err != nil {
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %s does not exists", id)
		log.Printf("[ERROR] CollectionItem Repository - UpdateCollectionItemById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *collectionRepository) DeleteCollectionItemById(ctx context.Context, id uuid.UUID) error {
	var total int64

	if err := r.db.Model(&entity.CollectionItem{}).
		Where("id = ?", id).
		Count(&total).
		Delete(&entity.CollectionItem{}).
		Error; err != nil {
		log.Printf("[ERROR] CollectionItem Repository - DeleteCollectionItemById : %v\n", err)
		return err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %s does not exists", id)
		log.Printf("[ERROR] CollectionItem Repository - DeleteCollectionItemById : %v\n", err)
		return err
	}

	return nil
}
