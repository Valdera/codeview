package repository

import (
	"codeview/config"
	"codeview/internal/entity"
	"codeview/internal/repository"
	"codeview/internal/util/pagination"
	"context"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type noteRepository struct {
	cfg config.AppConfig
	db  *gorm.DB
}

func NewNoteRepository(cfg config.AppConfig, db *gorm.DB) repository.NoteRepository {
	return &noteRepository{
		cfg,
		db,
	}
}

func (r *noteRepository) CreateNote(ctx context.Context, body *entity.Note) (*entity.Note, error) {
	if err := r.db.Model(&entity.Note{}).
		Create(body).
		Error; err != nil {
		log.Printf("[ERROR] Note Repository - CreateNote : %v\n", err)
		return nil, err
	}

	return body, nil
}

func (r *noteRepository) GetNoteById(ctx context.Context, id uint) (*entity.Note, error) {
	var result entity.Note
	var total int64

	if err := r.db.Model(&entity.Note{}).
		Count(&total).
		Preload("NoteItem").Preload("Tags").
		Where("id = ?", id).
		Find(&result).
		Error; err != nil {
		log.Printf("[ERROR] Note Repository - GetNoteById : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("note with id %d does not exists", id)
		log.Printf("[ERROR] Note Repository - GetNoteById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *noteRepository) GetNotes(ctx context.Context, p *pagination.Pagination) ([]entity.Note, error) {
	var results []entity.Note
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Note Repository - GetNotes : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Note{}).
		Count(&total).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Preload("Tags").
		Find(&results).
		Error; err != nil {
		log.Printf("[ERROR] Note Repository - GetNotes : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *noteRepository) GetNotesByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Note, error) {
	var results []entity.Note
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] Note Repository - GetNotesByIds : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.Note{}).
		Count(&total).
		Where("id IN (?)", ids).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Preload("Tags").Preload("Sources").
		Preload("Difficulty").
		Find(&results).
		Error; err != nil {
		log.Printf("[ERROR] Note Repository - GetNotesByIds : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *noteRepository) UpdateNoteById(ctx context.Context, id uint, body *entity.Note) (*entity.Note, error) {
	var result entity.Note
	var total int64

	if err := r.db.Model(&result).
		Count(&total).
		Where("id = ?", id).
		Updates(&entity.Note{
			Title: body.Title,
		}).
		Error; err != nil {
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("note with id %d does not exists", id)
		log.Printf("[ERROR] Note Repository - UpdateNoteById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *noteRepository) DeleteNoteById(ctx context.Context, id uint) error {
	var total int64

	if err := r.db.Model(&entity.Note{}).
		Count(&total).
		Where("id = ?", id).
		Delete(&entity.Note{}).
		Error; err != nil {
		log.Printf("[ERROR] Note Repository - DeleteNoteById : %v\n", err)
		return err
	}

	if total == 0 {
		err := fmt.Errorf("note with id %d does not exists", id)
		log.Printf("[ERROR] Note Repository - DeleteNoteById : %v\n", err)
		return err
	}

	return nil
}

func (r *noteRepository) CreateNoteItem(ctx context.Context, body *entity.NoteItem) (*entity.NoteItem, error) {
	if err := r.db.Model(&entity.NoteItem{}).
		Create(body).
		Error; err != nil {
		log.Printf("[ERROR] NoteItem Repository - CreateNoteItem : %v\n", err)
		return nil, err
	}

	return body, nil
}

func (r *noteRepository) GetNoteItemById(ctx context.Context, id uint) (*entity.NoteItem, error) {
	var result entity.NoteItem
	var total int64

	if err := r.db.Model(&entity.NoteItem{}).
		Count(&total).
		Where("id = ?", id).
		Find(&result).
		Error; err != nil {
		log.Printf("[ERROR] NoteItem Repository - GetNoteItemById : %v\n", err)
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("tag with id %d does not exists", id)
		log.Printf("[ERROR] NoteItem Repository - GetNoteItemById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *noteRepository) GetNoteItems(ctx context.Context, p *pagination.Pagination) ([]entity.NoteItem, error) {
	var results []entity.NoteItem
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] NoteItem Repository - GetNoteItems : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.NoteItem{}).
		Count(&total).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Find(&results).
		Error; err != nil {
		log.Printf("[ERROR] NoteItem Repository - GetNoteItems : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *noteRepository) GetNoteItemsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.NoteItem, error) {
	var results []entity.NoteItem
	var total int64

	if p == nil {
		err := fmt.Errorf("pagination must not empty")
		log.Printf("[ERROR] NoteItem Repository - GetNoteItemsByIds : %v\n", err)
		return nil, err
	}

	if err := r.db.Model(&entity.NoteItem{}).
		Count(&total).
		Where("id IN (?)", ids).
		Offset(p.GetOffset()).
		Limit(p.PageSize).
		Find(&results).
		Error; err != nil {
		log.Printf("[ERROR] NoteItem Repository - GetNoteItemsByIds : %v\n", err)
		return nil, err
	}

	p.Total = int(total)
	p.SetPagination()

	return results, nil
}

func (r *noteRepository) UpdateNoteItemById(ctx context.Context, id uint, body *entity.NoteItem) (*entity.NoteItem, error) {
	var result entity.NoteItem
	var total int64

	if err := r.db.Model(&result).
		Count(&total).
		Where("id = ?", id).
		Updates(&entity.NoteItem{
			Content: body.Content,
		}).
		Error; err != nil {
		return nil, err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] NoteItem Repository - UpdateNoteItemById : %v\n", err)
		return nil, err
	}

	return &result, nil
}

func (r *noteRepository) DeleteNoteItemById(ctx context.Context, id uint) error {
	var total int64

	if err := r.db.Model(&entity.NoteItem{}).
		Count(&total).
		Where("id = ?", id).
		Delete(&entity.NoteItem{}).
		Error; err != nil {
		log.Printf("[ERROR] NoteItem Repository - DeleteNoteItemById : %v\n", err)
		return err
	}

	if total == 0 {
		err := fmt.Errorf("problem with id %d does not exists", id)
		log.Printf("[ERROR] NoteItem Repository - DeleteNoteItemById : %v\n", err)
		return err
	}

	return nil
}
