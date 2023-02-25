package problemservice

import (
	"codeview/config"
	"codeview/internal/repository"
	"codeview/internal/service"
)

type problemService struct {
	config      config.AppConfig
	problemRepo repository.ProblemRepository
}

func New(config config.AppConfig, problemRepo repository.ProblemRepository) service.ProblemService {
	return &problemService{
		config,
		problemRepo,
	}
}
