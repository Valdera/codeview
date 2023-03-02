package service

import (
	"codeview/internal/dto/response"
	"context"
	"mime/multipart"
)

//go:generate mockery --name=ImageService --case underscore --testonly
type ImageService interface {
	UploadImage(ctx context.Context, imageFileHeader *multipart.FileHeader) (*response.ImageResponse, error)
}
