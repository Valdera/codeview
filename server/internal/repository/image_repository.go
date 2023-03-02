package repository

import (
	"codeview/internal/entity"
	"context"
	"mime/multipart"
)

//go:generate mockery --name=ImageRepository --case underscore --testonly
type ImageRepository interface {
	UploadImage(ctx context.Context, objName string, imageFile multipart.File) (*entity.Image, error)
}
