package domain

import (
	"codeview/internal/entity"
	"context"
	"mime/multipart"
)

type ImageRepository interface {
	AddImage(ctx context.Context, objName string, imageFile multipart.File) (*entity.Image, error)
}
