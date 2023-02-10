package domain

import (
	"codeview/internal/dto"
	"codeview/internal/entity"

	"context"
	"mime/multipart"
)

type ImageRepository interface {
	UploadImage(ctx context.Context, objName string, imageFile multipart.File) (*entity.Image, error)
}

type ImageService interface {
	UploadImage(ctx context.Context, imageFileHeader *multipart.FileHeader) (*dto.ImageResponse, error)
}
