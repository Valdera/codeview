package repository

import (
	"codeview/internal/entity"
	"codeview/internal/util/pagination"
	"context"
)

//go:generate mockery --name=NoteRepository --case underscore --testonly
type NoteRepository interface {
	CreateNote(ctx context.Context, body *entity.Note) (*entity.Note, error)
	GetNotes(ctx context.Context, p *pagination.Pagination) ([]entity.Note, error)
	GetNoteById(ctx context.Context, id uint) (*entity.Note, error)
	GetNotesByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.Note, error)
	UpdateNoteById(ctx context.Context, id uint, body *entity.Note) (*entity.Note, error)
	DeleteNoteById(ctx context.Context, id uint) error

	CreateNoteItem(ctx context.Context, body *entity.NoteItem) (*entity.NoteItem, error)
	GetNoteItems(ctx context.Context, p *pagination.Pagination) ([]entity.NoteItem, error)
	GetNoteItemById(ctx context.Context, id uint) (*entity.NoteItem, error)
	GetNoteItemsByIds(ctx context.Context, ids []uint, p *pagination.Pagination) ([]entity.NoteItem, error)
	UpdateNoteItemById(ctx context.Context, id uint, body *entity.NoteItem) (*entity.NoteItem, error)
	DeleteNoteItemById(ctx context.Context, id uint) error
}
