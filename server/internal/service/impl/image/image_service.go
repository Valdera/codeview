package imageservice

import (
	"codeview/config"
	"codeview/internal/repository"
	"codeview/internal/service"

	"codeview/internal/dto/response"
	"codeview/internal/util/exception"
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/google/uuid"
)

type imageService struct {
	config    config.AppConfig
	imageRepo repository.ImageRepository
}

func New(config config.AppConfig, imageRepo repository.ImageRepository) service.ImageService {
	return &imageService{
		config,
		imageRepo,
	}
}

func (s *imageService) UploadImage(ctx context.Context, imageFileHeader *multipart.FileHeader) (*response.ImageResponse, error) {
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

	return &response.ImageResponse{
		URL: image.URL,
	}, nil
}
