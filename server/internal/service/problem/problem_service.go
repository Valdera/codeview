package service

import (
	"codeview/domain"
)

type problemService struct {
	problemRepo domain.ProblemRepository
}

func New(problemRepo domain.ProblemRepository) domain.ProblemService {
	return &problemService{
		problemRepo,
	}
}
