package repository

import (
	"codeview/internal/entity"
	"context"
	"mime/multipart"
)

type ImageRepository interface {
	UploadImage(ctx context.Context, objName string, imageFile multipart.File) (*entity.Image, error)
}
