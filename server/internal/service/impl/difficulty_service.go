package service

import (
	"codeview/config"
	"codeview/internal/dto/request"
	"codeview/internal/dto/response"
	"codeview/internal/entity"
	"codeview/internal/mapper"
	"codeview/internal/repository"
	"codeview/internal/service"
	"codeview/internal/util/pagination"
	"context"
	"log"

	"github.com/google/uuid"
)

type difficultyService struct {
	cfg            config.AppConfig
	difficultyRepo repository.DifficultyRepository
}

func NewDifficultyService(cfg config.AppConfig, difficultyRepo repository.DifficultyRepository) service.DifficultyService {
	return &difficultyService{
		cfg,
		difficultyRepo,
	}
}

func (s *difficultyService) CreateDifficulty(ctx context.Context, body *request.DifficultyCreate) (*response.Difficulty, error) {
	difficulty, err := s.difficultyRepo.CreateDifficulty(ctx, &entity.Difficulty{
		Label: body.Label,
		Color: body.Color,
	})
	if err != nil {
		log.Printf("[ERROR] Difficulty Service - CreateDifficulty : %v\n", err)
		return nil, err
	}

	response, err := mapper.DifficultyToDTO(difficulty)
	if err != nil {
		log.Printf("[ERROR] Difficulty Service - CreateDifficulty : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *difficultyService) GetDifficulties(ctx context.Context, p *pagination.Pagination) ([]response.Difficulty, error) {
	difficulties, err := s.difficultyRepo.GetDifficulties(ctx, p)
	if err != nil {
		log.Printf("[ERROR] Difficulty Service - GetDifficulties : %v\n", err)
		return nil, err
	}

	responses := make([]response.Difficulty, len(difficulties))
	for i, difficulty := range difficulties {
		response, err := mapper.DifficultyToDTO(&difficulty)
		if err != nil {
			log.Printf("[ERROR] Difficulty Service - GetDifficulties : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *difficultyService) GetDifficultyById(ctx context.Context, id uuid.UUID) (*response.Difficulty, error) {
	difficulty, err := s.difficultyRepo.GetDifficultyById(ctx, id)
	if err != nil {
		log.Printf("[ERROR] Difficulty Service - GetDifficultyById : %v\n", err)
		return nil, err
	}

	response, err := mapper.DifficultyToDTO(difficulty)
	if err != nil {
		log.Printf("[ERROR] Difficulty Service - GetDifficultyById : %v\n", err)
		return nil, err
	}

	return response, nil
}

func (s *difficultyService) GetDifficultiesByIds(ctx context.Context, ids []uuid.UUID, p *pagination.Pagination) ([]response.Difficulty, error) {
	difficulties, err := s.difficultyRepo.GetDifficultiesByIds(ctx, ids, p)
	if err != nil {
		log.Printf("[ERROR] Difficulty Service - GetDifficultiesByIds : %v\n", err)
		return nil, err
	}

	responses := make([]response.Difficulty, len(difficulties))
	for i, difficulty := range difficulties {
		response, err := mapper.DifficultyToDTO(&difficulty)
		if err != nil {
			log.Printf("[ERROR] Difficulty Service - GetDifficultiesByIds : %v\n", err)
			return nil, err
		}
		responses[i] = *response
	}

	return responses, nil
}

func (s *difficultyService) UpdateDifficultyById(ctx context.Context, id uuid.UUID, body *request.DifficultyUpdate) (*response.Difficulty, error) {
	difficulty, err := s.difficultyRepo.UpdateDifficultyById(ctx, id, &entity.Difficulty{
		Label: body.Label,
		Color: body.Color,
	})
	if err != nil {
		log.Printf("[ERROR] Difficulty Service - UpdateDifficultyById : %v\n", err)
		return nil, err
	}

	response, err := mapper.DifficultyToDTO(difficulty)
	if err != nil {
		log.Printf("[ERROR] Difficulty Service - UpdateDifficultyById : %v\n", err)
		return nil, err
	}

	return response, err
}

func (s *difficultyService) DeleteDifficultyById(ctx context.Context, id uuid.UUID) error {
	if err := s.difficultyRepo.DeleteDifficultyById(ctx, id); err != nil {
		log.Printf("[ERROR] Difficulty Service - DeleteDifficultyById : %v\n", err)
		return err
	}

	return nil
}
