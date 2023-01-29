package service

import (
	"codeview/domain"
	"codeview/exception"
	"codeview/internal/entity"
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/google/uuid"
)

type problemService struct {
	imageRepo domain.ImageRepository
}

func New(imageRepo domain.ImageRepository) domain.ProblemService {
	return &problemService{
		imageRepo,
	}
}

func (s *problemService) AddProblemImage(ctx context.Context, imageFileHeader *multipart.FileHeader) (*entity.Image, error) {
	objID, _ := uuid.NewRandom()
	objName := fmt.Sprintf("problem_%s", objID.String())

	imageFile, err := imageFileHeader.Open()
	if err != nil {
		log.Printf("Failed to open image file: %v\n", err)
		return nil, exception.NewInternal()
	}

	image, err := s.imageRepo.AddImage(ctx, objName, imageFile)
	if err != nil {
		log.Printf("Unable to upload image to cloud provider: %v\n", err)
		return nil, err
	}

	return image, nil
}
