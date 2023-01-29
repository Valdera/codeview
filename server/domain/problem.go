package domain

import (
	"codeview/internal/entity"
	"context"
	"mime/multipart"
)

type ProblemRepository interface {
}

type ProblemService interface {
	AddProblemImage(ctx context.Context, imageFileHeader *multipart.FileHeader) (*entity.Image, error)
}
