package service

import (
	"codeview/domain"
	"codeview/exception"
	"codeview/internal/dto"
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/google/uuid"
)

type imageService struct {
	imageRepo domain.ImageRepository
}

func New(imageRepo domain.ImageRepository) domain.ImageService {
	return &imageService{
		imageRepo,
	}
}

func (s *imageService) UploadImage(ctx context.Context, imageFileHeader *multipart.FileHeader) (*dto.ImageResponse, error) {
	objID, _ := uuid.NewRandom()
	objName := fmt.Sprintf("problem_%s", objID.String())

	imageFile, err := imageFileHeader.Open()
	if err != nil {
		log.Printf("Failed to open image file: %v\n", err)
		return nil, exception.NewInternal()
	}

	image, err := s.imageRepo.UploadImage(ctx, objName, imageFile)
	if err != nil {
		log.Printf("Unable to upload image to cloud provider: %v\n", err)
		return nil, err
	}

	return &dto.ImageResponse{
		URL: image.URL,
	}, nil
}
