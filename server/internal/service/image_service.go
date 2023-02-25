package service

import (
	"codeview/internal/dto/response"
	"context"
	"mime/multipart"
)

type ImageService interface {
	UploadImage(ctx context.Context, imageFileHeader *multipart.FileHeader) (*response.ImageResponse, error)
}
